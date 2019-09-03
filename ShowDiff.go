package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func doShow(w http.ResponseWriter, r *http.Request) {
	showHtmlHead(w)

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

	showHtmlFloor(w)

}

func showHtmlHead(w http.ResponseWriter) {
	fmt.Fprintf(w, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">")
	fmt.Fprintf(w, "<html xmlns=\"http://www.w3.org/1999/xhtml\">")
	fmt.Fprintf(w, "<body>")
}

func showHtmlFloor(w http.ResponseWriter) {
	fmt.Fprintf(w, "</body>")
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

func main() {
	port := "9090"
	println("ready go http://localhost:" + port)

	http.HandleFunc("/", doShow)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
