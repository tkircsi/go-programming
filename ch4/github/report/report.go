package report

import (
	"io"
	"log"
	"text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------------------------------------------
Number:	{{.Number}}
User:	{{.User.Login}}
Title:	{{.Title}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

const htmlTempl = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func PrintTextReport(w io.Writer, data interface{}) {
	var report = template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	if err := report.Execute(w, data); err != nil {
		log.Fatalf("Report error: %s\n", err)
	}
}

func PrintHTMLReport(w io.Writer, data interface{}) {
	var report = template.Must(template.New("issuelist").
		Parse(htmlTempl))
	if err := report.Execute(w, data); err != nil {
		log.Fatalf("Report error: %s\n", err)
	}
}
