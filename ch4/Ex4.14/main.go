// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"Ex4.14/github"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>Age</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td>{{.CreatedAt | daysAgo}} days</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

var issueList = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		query := []string{""}
		if q.Has("query") {
			query = q["query"]
		}
		fmt.Println(q)
		printQuery(w, query)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func printQuery(out io.Writer, query []string) {
	if len(query) == 1 && query[0] == "" {
		fmt.Fprintf(out, "%s", "")
		return
	}
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(out, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
