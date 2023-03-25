package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Clicks map[int][]string

func (x Clicks) Len() int           { return len(x) }
func (x Clicks) Less(i, j int) bool { return i < j }
func (x Clicks) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type multiTierSort struct {
	t              []*Track
	clickedColumns Clicks
	less           func(x, y *Track, columnClicks Clicks) bool
}

func (x multiTierSort) Len() int           { return len(x.t) }
func (x multiTierSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j], x.clickedColumns) }
func (x multiTierSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	clickColumns := Clicks(map[int][]string{
		2: {"Artist"},
		3: {"Album"},
		0: {"Year"},
		5: {"Title", "Length"},
	})

	byClicks := func(x, y *Track, columnClicks Clicks) bool {
		sort.Sort(columnClicks)
		sort.Sort(sort.Reverse(columnClicks)) // sorting in ascending order
		for _, fields := range columnClicks {
			for _, field := range fields {
				switch field {
				case "Title":
					if x.Title == y.Title {
						continue
					}
					return x.Title < y.Title
				case "Artist":
					if x.Artist == y.Artist {
						continue
					}
					return x.Artist < y.Artist
				case "Album":
					if x.Album == y.Album {
						continue
					}
					return x.Album < y.Album
				case "Year":
					if x.Year == y.Year {
						continue
					}
					return x.Year < y.Year
				case "Length":
					if int(x.Length) == int(y.Length) {
						continue
					}
					return int(x.Length) < int(y.Length)
				default:
					return x.Title < y.Title
				}
			}
		}
		return false
	}

	fmt.Printf("Before sorting:\n")
	printTracks(tracks)
	sort.Sort(multiTierSort{t: tracks, less: byClicks, clickedColumns: clickColumns})
	fmt.Printf("\n\nAfter sorting by Clicks:\n")
	printTracks(tracks)
}
