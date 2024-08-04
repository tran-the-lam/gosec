// Using a variable to initialize could potentially be dangerous. Under the
// current model this will likely produce some false positives.
package main

import (
	"html/template"
	"os"
)

func main() {
	a := "<b>something from another place<b>"
	tmpl := `<h1>{{.Title}}</h1><p>{{.Body}}</p>`
	t := template.Must(template.New("ex").Parse(tmpl))
	v := map[string]interface{}{
		"Title": "Test <b>World</b>",
		"Body":  a,
	}
	t.Execute(os.Stdout, v)
}
