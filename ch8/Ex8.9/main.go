package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RootSizePiece struct {
	root string
	size int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan RootSizePiece)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	nFilesGroup := make(map[string]int64)
	nBytesGroup := make(map[string]int64)
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nFilesGroup[size.root]++
			nBytesGroup[size.root] += size.size
		case <-tick:
			for _, root := range roots {
				printDiskUsage(nFilesGroup[root], nBytesGroup[root], root)
			}
			fmt.Printf("\n")
		}
	}

	nfiles := int64(0)
	nbytes := int64(0)
	for _, root := range roots {
		nfiles += nFilesGroup[root]
		nbytes += nBytesGroup[root]
		printDiskUsage(nFilesGroup[root], nBytesGroup[root], root)
	}
	fmt.Printf("\n")
	printTotalDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64, root string) {
	fmt.Printf("Root dir %s disk usage %d files %.1f GB\n", root, nfiles, float64(nbytes)/1e9)
}
func printTotalDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("Total disk usage %d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir, root string, n *sync.WaitGroup, fileSizes chan<- RootSizePiece) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, root, n, fileSizes)
		} else {
			// get the size of the entry file
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
				// sending 0 to avoid not counting this file
				fileSizes <- RootSizePiece{root: root, size: 0}
				continue
			}
			fileSizes <- RootSizePiece{root: root, size: info.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
