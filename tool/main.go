package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TupleData struct {
	Package string
	N       int
	Fields  []FieldData
}

type FieldData struct {
	Index int
	Name  string
	Type  string
}

const tupleTemplate = `package {{.Package}}

type T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}} any{{end}}] struct {
{{range .Fields}}	{{.Name}} {{.Type}}
{{end}}}

// NewT{{.N}} creates a new T{{.N}} tuple
func NewT{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}} any{{end}}]({{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Name | lower}} {{$f.Type}}{{end}}) T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}] {
	return T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}]{
{{range .Fields}}		{{.Name}}: {{.Name | lower}},
{{end}}	}
}

// First returns the first field of the tuple
func (t T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}]) First() {{(index .Fields 0).Type}} {
	return t.{{(index .Fields 0).Name}}
}

// Last returns the last field of the tuple
func (t T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}]) Last() {{(index .Fields (sub (len .Fields) 1)).Type}} {
	return t.{{(index .Fields (sub (len .Fields) 1)).Name}}
}
{{if gt .N 1}}
// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}]) CutLast() T{{sub .N 1}}[{{range $i, $f := .Fields}}{{if lt $i (sub (len $.Fields) 1)}}{{if $i}}, {{end}}{{$f.Type}}{{end}}{{end}}] {
	return T{{sub .N 1}}[{{range $i, $f := .Fields}}{{if lt $i (sub (len $.Fields) 1)}}{{if $i}}, {{end}}{{$f.Type}}{{end}}{{end}}]{
{{range $i, $f := .Fields}}{{if lt $i (sub (len $.Fields) 1)}}		{{$f.Name}}: t.{{$f.Name}},
{{end}}{{end}}	}
}
{{end}}
// Get returns all tuple field values
func (t T{{.N}}[{{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}]) Get() ({{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Type}}{{end}}) {
	return {{range $i, $f := .Fields}}{{if $i}}, {{end}}t.{{$f.Name}}{{end}}
}
`

func main() {
	var packageName string
	var outputPath string
	var n int

	flag.StringVar(&packageName, "package", "", "the package to generate code in")
	flag.StringVar(&outputPath, "out", "", "the path to write generated code to")
	flag.IntVar(&n, "N", 0, "the number of tuple fields to generate")

	flag.Parse()

	if packageName == "" {
		fmt.Fprintf(os.Stderr, "Error: package name is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if outputPath == "" {
		fmt.Fprintf(os.Stderr, "Error: output path is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if n <= 0 {
		fmt.Fprintf(os.Stderr, "Error: N must be a positive integer\n")
		flag.Usage()
		os.Exit(1)
	}

	// Generate field data
	var fields []FieldData
	for i := 0; i < n; i++ {
		fields = append(fields, FieldData{
			Index: i,
			Name:  fmt.Sprintf("F%d", i),
			Type:  fmt.Sprintf("F%d", i),
		})
	}

	data := TupleData{
		Package: packageName,
		N:       n,
		Fields:  fields,
	}

	// Create custom template functions
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
		"lt": func(a, b int) bool {
			return a < b
		},
	}

	// Parse template
	tmpl, err := template.New("tuple").Funcs(funcMap).Parse(tupleTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
		os.Exit(1)
	}

	// Create output directory if it doesn't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Execute template
	if err := tmpl.Execute(file, data); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated tuple struct T%d in package %s at %s\n", n, packageName, outputPath)
}
