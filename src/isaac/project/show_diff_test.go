package project_test

import (
	"bufio"
	"fmt"
	"io"
	"isaac/html"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func doShow(w http.ResponseWriter, r *http.Request) {
	html.ShowHtmlHead(w)

	reportFile := "F:\\dir\\dif.txt"
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

		showData(w, string(data))
	}

	html.ShowHtmlFloor(w)

}

func showData(w http.ResponseWriter, data string) {
	var html string
	if strings.HasPrefix(data, "#########") || strings.HasPrefix(data, "E:\\app\\unifiedcenter") {
		html = "<span style=\"color:red;\"><strong>" + data + "</strong></span><br/>"
	} else if strings.HasPrefix(data, ">>>>>>>>>>>>>") {
		html = "<span style=\"color:#0000FF;\">" + data + "</span><br/>"
	} else {
		html = data + "<br/>"
	}
	fmt.Fprintf(w, html)
}

func TestShowDiff(t *testing.T) {
	port := "9090"
	println("ready go http://localhost:" + port)

	http.HandleFunc("/", doShow)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
