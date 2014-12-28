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

// Quick, ad-hoc test. I just kept this as reference for writing slice literals
// TODO: figure out how to put this in a real test
// var fake_links = Files{
// 	[]Link{
// 		Link{"N1", "L1"},
// 		Link{"N2", "L2"},
// 		Link{"N3", "L3"},
// 	}}

func ToStringSlice(words [][]byte) []string {
	result := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		result[i] = string(words[i])
	}
	return result
}

// Gets the contents of the current dir by fork/executing "ls"
func PrintCurrentDirs1() (filenames []string, err error) {
	cmd := exec.Command("ls")
	stdout, e := cmd.StdoutPipe()
	if e != nil {
		return nil, e
	}
	if e := cmd.Start(); e != nil {
		return nil, e
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	if e := cmd.Wait(); e != nil {
		return nil, e
	}
	return ToStringSlice(bytes.Split(buf.Bytes(), []byte{'\n'})), nil
}

// Gets the contents of the current dir with os.Readdir(".") (presumably
// implemented as a system call)
func PrintCurrentDirs2() (filenames []string, err error) {
	curdir, e := os.Open(".")
	if e != nil {
		return nil, e
	}
	result, e := curdir.Readdirnames(-1)
	return result, e
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":80", nil)

	// Alternatively, os.Readdir(".")
	files, e := PrintCurrentDirs2()
	if e != nil {
		panic(e)
	}
	links := make([]Link, len(files))
	for i := 1; i < len(files); i++ {
		links[i] = Link{string(files[i]), string(files[i])}
	}

	t, e := template.ParseFiles("index.html")
	if e != nil {
		panic(e)
	}

	// put this in the handler above, and give it to w
	// (http.ResponseWriter) instead of os.Stdout
	fake_links := Files{links}
	t.Execute(os.Stdout, fake_links)
}
