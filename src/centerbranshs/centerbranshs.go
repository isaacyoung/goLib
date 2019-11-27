package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	projectPath := "http://192.168.68.12/svn/app/unifiedcenter/branches/"
	branshNum := 15

	out, err := exec.Command("svn", "list", projectPath).Output()
	if err != nil {
		log.Fatal(err)
	}

	branshs := strings.Split(string(out), "\r\n")
	branshNum = len(branshs) - branshNum
	if branshNum < 0 {
		branshNum = 0
	}
	branshs = branshs[branshNum:]
	for _, bransh := range branshs {
		branshFullPath := projectPath + bransh

		out2, err := exec.Command("svn", "log", "--stop-on-copy", branshFullPath).Output()
		if err != nil {
			log.Fatal(err)
		}

		outStr, err := simplifiedchinese.GBK.NewDecoder().String(string(out2))

		groups := strings.Split(outStr, "------------------------------------------------------------------------")
		block := ""
		for _, group := range groups {
			if strings.ReplaceAll(group, "\r\n", "") != "" {
				block = group
			}
		}

		rows := strings.Split(block, "\r\n")
		msg := ""
		for _, row := range rows {
			if strings.Trim(row, " ") != "" {
				msg = row
			}
		}

		if strings.Contains(msg, "|") {
			msg = ""
		}

		fmt.Printf("%s %s\r\n", bransh, msg)

	}

	time.Sleep(10000 * time.Millisecond)
}
