package html

import (
	"fmt"
	"net/http"
)

func ShowHtmlHead(w http.ResponseWriter) {
	fmt.Fprintf(w, "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">")
	fmt.Fprintf(w, "<html xmlns=\"http://www.w3.org/1999/xhtml\">")
	fmt.Fprintf(w, "<body>")
}

func ShowHtmlFloor(w http.ResponseWriter) {
	fmt.Fprintf(w, "</body>")
}
