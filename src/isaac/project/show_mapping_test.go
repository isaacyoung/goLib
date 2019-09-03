package project_test

import (
	"bufio"
	"fmt"
	"io"
	"isaac/html"
	"log"
	"net/http"
	"os"
	"testing"
)

func doShowMapping(w http.ResponseWriter, r *http.Request) {
	html.ShowHtmlHead(w)

	reportFile := "F:\\dir\\mapping.txt"
	fd, err := os.Open(reportFile)
	defer fd.Close()
	if err != nil {
		println("read error:", err)
	}
	buff := bufio.NewReader(fd)
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}

		showDataRows(w, string(data))
	}

	html.ShowHtmlFloor(w)

}

func showDataRows(w http.ResponseWriter, data string) {
	fmt.Fprintf(w, data+"<br/>")
}

func TestShowMapping(t *testing.T) {
	port := "9091"
	println("ready go http://localhost:" + port)

	http.HandleFunc("/", doShowMapping)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
