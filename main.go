package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"text/template"
)

func main() {
	args := os.Args[1:]
	r := regexp.MustCompile(args[1])

	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("substr").Parse(args[2])
	if err != nil {
		panic(err)
	}

	match := r.FindStringSubmatch(string(file))

	data := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			data[name] = match[i]
		}
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
