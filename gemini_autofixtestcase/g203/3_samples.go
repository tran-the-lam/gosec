package main

import (
	"html/template"
	"os"
)

const tmpl = ""

func main() {
	a := "something from another place"
	t := template.Must(template.New("ex").Parse(tmpl))
	v := map[string]interface{}{
		"Title": "Test <b>World</b>",
		"Body":  template.URL(a),
	}
	t.Execute(os.Stdout, v)
}
