// See https://code.google.com/p/go-wiki/wiki/Ubuntu
// To install go you can do either
//   apt-get install golang # for v. 1.1
// or
//   sudo add-apt-repository ppa:gophers/go
//   sudo apt-get update
//   sudo apt-get install golang-stable
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	// "html/template"
	"net/http"
	"text/template"
)

type Link struct {
	Name   string
	Target string
}

type Files struct {
	Files []Link
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.Path[1:] = \"%s\"", r.URL.Path[1:])
}

var fake_links = Files{
	[]Link{
		Link{"N1", "L1"},
		Link{"N2", "L2"},
		Link{"N3", "L3"},
	}}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":80", nil)

	// Alternatively, os.Readdir(".")
	cmd := exec.Command("ls")
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		panic(e)
	}
	if e := cmd.Run(); e != nil {
		panic(e)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	files := Split(buf.String(), "\n")
	links = make([]Link, len(files))
	for i := range len(files) {
		links[i] = Link{files[i], files[i]}
	}

	if t, e := template.ParseFiles("index.html"); e != nil {
		panic(e)
	}

	// put this in the handler above, and give it to w
	// (http.ResponseWriter) instead of os.Stdout
	fake_links := Files{links}
	t.Execute(os.Stdout, fake_links)
}
