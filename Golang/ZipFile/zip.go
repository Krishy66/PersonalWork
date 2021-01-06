package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) == ".pdf" {
			*files = append(*files, path)
		}
		return nil

	}

}

func zippy() {

	var files []string

	outFile, err := os.Create("./zipsss/All_Students_Reports" + ".zip")

	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	// Create a new zip archive.
	wr := zip.NewWriter(outFile)

	root := "../Bl's/pdfs"

	err = filepath.Walk(root, visit(&files))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		full := strings.Split(string(file), string("\\"))

		last := strings.Join(full[1:], "/")

		f, err := wr.Create(last)

		if err != nil {
			log.Fatal(err)
		}

		data, _ := ioutil.ReadFile(file)

		_, err = f.Write(data)

		if err != nil {
			log.Fatal(err)
		}

	}
	defer wr.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", foo)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func foo(w http.ResponseWriter, r *http.Request) {

	zippy()
	zipp, _ := os.Open("./zipsss/All_Students_Reports.zip")

	w.Header().Set("Content-Disposition", "attachment; filename=filee.zip")

	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, zipp)

	defer os.Remove("./zipsss/All_Students_Reports" + ".zip")
	zipp.Close()
}
