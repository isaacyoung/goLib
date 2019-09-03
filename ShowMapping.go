package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func doShow(w http.ResponseWriter, r *http.Request) {
	showHtmlHead(w)

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
	fmt.Fprintf(w, data+"<br/>")
}

func main() {
	port := "9091"
	println("ready go http://localhost:" + port)

	http.HandleFunc("/", doShow)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
