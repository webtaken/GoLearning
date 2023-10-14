package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracksList = template.Must(template.New("trackslist").Parse(`
<table>
<tr style='text-align: left'>
<th>#</th>
<th>Title</th>
<th>Artist</th>
<th>Album</th>
<th>Year</th>
<th>Length</th>
</tr>
{{range $index, $track := .Items}}
<tr>
<td>{{$index}}</td>
<td>{{$track.Title}}</td>
<td>{{$track.Artist}}</td>
<td>{{$track.Album}}</td>
<td>{{$track.Year}}</td>
<td>{{$track.Length}}</td>
</tr>
{{end}}
</table>
`))

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

type TracksHTML struct {
	Items []*Track
}

func main() {
	tracksHTML := TracksHTML{Items: tracks}
	fmt.Printf("Before sorting:\n")
	if err := tracksList.Execute(os.Stdout, tracksHTML); err != nil {
		log.Fatal(err)
	}
	// printTracks(tracks)
	sort.Sort(byArtist(tracks))
	fmt.Printf("\n\nAfter sorting:\n")
	if err := tracksList.Execute(os.Stdout, tracksHTML); err != nil {
		log.Fatal(err)
	}
	// printTracks(tracks)
	sort.Sort(sort.Reverse(byArtist(tracks)))
	fmt.Printf("\n\nAfter sorting again:\n")
	if err := tracksList.Execute(os.Stdout, tracksHTML); err != nil {
		log.Fatal(err)
	}
	// printTracks(tracks)
}
