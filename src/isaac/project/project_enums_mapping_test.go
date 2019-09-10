package project_test

import (
	"bufio"
	"io"
	"isaac/file"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readFile2(file string) []string {
	fd, err := os.Open(file)
	defer fd.Close()
	if err != nil {
		println("read error:", err)
	}
	buff := bufio.NewReader(fd)

	var rows []string
	index := 1
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}

		rows = append(rows, strconv.Itoa(index)+" "+string(data))

		index = index + 1
	}

	return rows
}

func TestEnums(t *testing.T) {
	fromProject := "E:\\app\\unifiedcenterbraches\\unifiedcenter-1.3.33"
	includePath := []string{"com\\yunma\\unifiedcenter\\common\\enums"}
	includeFile := []string{".java"}
	excludePath := []string{".svn", ".idea", "unifiedcenter-web-manager", "target", "unifiedcenter-generate"}
	excludeFile := []string{"Test.java"}
	//end := "\r\n"

	reportFile := "F:\\dir\\enums.txt"
	if _, err := os.Stat(reportFile); err != nil {
		if os.IsExist(err) {
			os.Remove(reportFile)
		}
	}
	currFile, err := os.Create(reportFile)
	defer currFile.Close()
	if err != nil {
		println(err)
		return
	}

	fromFiles, err := file.GetProjectFiles(fromProject, includePath, includeFile, excludePath, excludeFile)
	if err != nil {
		println(err)
		return
	}

	includeFile = []string{".html"}
	includePath = []string{"unifiedcenter-web-manager"}
	excludePath = []string{".svn", ".idea", "target", "unifiedcenter-generate"}
	targetFiles, err := file.GetProjectFiles(fromProject, includePath, includeFile, excludePath, nil)

	for _, fromFile := range fromFiles {

		className := fromFile[strings.LastIndex(fromFile, "\\")+1 : strings.LastIndex(fromFile, ".")]
		fullClassName := fromFile[strings.LastIndex(fromFile, "java\\")+5 : strings.LastIndex(fromFile, ".")]
		fullClassName = strings.ReplaceAll(fullClassName, "\\", ".")

		for _, targetFile := range targetFiles {
			targetRows := readFile2(targetFile)
			for _, row := range targetRows {
				if strings.Contains(row, "."+className) && !strings.Contains(row, fullClassName) {
					println("")
					println(targetFile)
					println("______________________________________________")
					println(row)
					println(className)
				}
			}
		}

	}

}
