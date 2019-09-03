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

func readFile(file string) []string {
	fd, err := os.Open(file)
	defer fd.Close()
	if err != nil {
		println("read error:", err)
	}
	buff := bufio.NewReader(fd)

	fileName := file[strings.LastIndex(file, string(os.PathSeparator))+1 : strings.LastIndex(file, ".")]

	var rows []string
	index := 1
	begin := false
	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}

		if strings.HasPrefix(string(data), "public") && strings.Contains(string(data), fileName) {
			begin = true
		}

		if begin {
			rows = append(rows, strconv.Itoa(index)+" "+string(data))
		}
		index = index + 1
	}

	return rows
}

func getTargetFile(fromFile string, targetFiles []string) string {
	for _, file := range targetFiles {
		name := file[strings.LastIndex(file, string(os.PathSeparator)):]
		if strings.HasSuffix(fromFile, name) {
			return file
		}
	}
	return ""
}

func TestDiff(t *testing.T) {
	fromProject := "E:\\app\\unifiedcenter"
	targetProject := "E:\\app\\paymenthall"
	includePath := []string{"jiaofei"}
	includeFile := []string{".java"}
	excludePath := []string{".svn", ".idea", "unifiedcenter-web-manager", "target", "unifiedcenter-generate"}
	excludeFile := []string{"Test.java"}
	end := "\r\n"

	reportFile := "F:\\dir\\dif.txt"
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

	targetFiles, err := file.GetProjectFiles(targetProject, nil, includeFile, excludePath, nil)
	if err != nil {
		println(err)
		return
	}

	currFile.WriteString("not found_______________________________________________________________________" + end)
	fileMap := make(map[string]string)
	for _, fromFile := range fromFiles {
		targetFile := getTargetFile(fromFile, targetFiles)
		if targetFile == "" {
			currFile.WriteString(fromFile + end)
		} else {
			fileMap[fromFile] = targetFile
		}
	}

	currFile.WriteString("_______________________________________________________________________" + end)

	for k, v := range fileMap {
		fromRows := readFile(k)
		targetRows := readFile(v)

		var errorList []string
		var errorFrom []string
		var errorTo []string

		diff := false
		errorIndex := -2
		for i, row := range fromRows {
			if i >= len(targetRows) {
				break
			}
			if row[strings.Index(row, " "):] != targetRows[i][strings.Index(targetRows[i], " "):] {
				diff = true

				if errorIndex == -2 || i == errorIndex+1 {
					errorFrom = append(errorFrom, row)
					errorTo = append(errorTo, targetRows[i])
				} else {
					errorList = append(errorList, errorFrom...)
					errorList = append(errorList, ">>>>>>>>>>>>>>>")
					errorList = append(errorList, errorTo...)
					errorList = append(errorList, "__________________________________________")
					errorFrom = nil
					errorTo = nil
				}
				errorIndex = i
			}
		}

		if errorFrom != nil {
			errorList = append(errorList, errorFrom...)
			errorList = append(errorList, ">>>>>>>>>>>>>>>")
			errorList = append(errorList, errorTo...)
			errorList = append(errorList, "__________________________________________")
		}

		if diff {

			currFile.WriteString(end)
			currFile.WriteString("#################################################" + end)
			currFile.WriteString(k + end)

			for _, block := range errorList {
				currFile.WriteString(block + end)
			}
		}
	}
}
