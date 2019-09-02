package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func GetFiles(path string, includeFile []string, excludePath []string, excludeFile []string) (files []string, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		println(err)
		return files, err
	}

	sep := string(os.PathSeparator)

	for _, file := range dir {
		p := path + sep + file.Name()
		if file.IsDir() {
			if !ExcludePath(p, excludePath) {
				files2, err := GetFiles(p, includeFile, excludePath, excludeFile)
				if err != nil {
					println(err)
					return files, err
				}
				files = append(files, files2...)
			}

		} else {
			if IncludeFile(file.Name(), includeFile) && !ExcludeFile(file.Name(), excludeFile) {
				files = append(files, p)
			}

		}
	}

	return files, nil
}

func IncludePath(path string, includePath []string) bool {
	if includePath == nil {
		return true
	}

	for _, p := range includePath {
		if strings.Contains(path, string(os.PathSeparator)+p) {
			return true
		}
	}
	return false
}

func IncludeFile(file string, includeFile []string) bool {
	if includeFile == nil {
		return true
	}

	for _, ext := range includeFile {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

func ExcludePath(path string, excludePath []string) bool {
	if excludePath == nil {
		return false
	}

	for _, p := range excludePath {
		if strings.Contains(path, string(os.PathSeparator)+p) {
			return true
		}
	}
	return false
}

func ExcludeFile(file string, excludeFile []string) bool {
	if excludeFile == nil {
		return false
	}

	for _, ext := range excludeFile {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

func GetProjectFiles(path string, includePath []string, includeFile []string, excludePath []string, excludeFile []string) ([]string, error) {
	files, err := GetFiles(path, includeFile, excludePath, excludeFile)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, f := range files {
		if IncludePath(f, includePath) {
			fileList = append(fileList, f)
		}
	}
	return fileList, nil
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

func ReadFile(file string) []string {
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

func main() {
	fromProject := "E:\\app\\unifiedcenter"
	targetProject := "E:\\app\\paymenthall"
	includePath := []string{"jiaofei"}
	includeFile := []string{".java"}
	excludePath := []string{".svn", ".idea", "unifiedcenter-web-manager", "target", "unifiedcenter-generate"}
	excludeFile := []string{"Test.java"}

	fromFiles, err := GetProjectFiles(fromProject, includePath, includeFile, excludePath, excludeFile)
	if err != nil {
		println(err)
		return
	}

	targetFiles, err := GetProjectFiles(targetProject, nil, includeFile, excludePath, nil)
	if err != nil {
		println(err)
		return
	}

	println("not found_______________________________________________________________________")
	fileMap := make(map[string]string)
	for _, fromFile := range fromFiles {
		targetFile := getTargetFile(fromFile, targetFiles)
		if targetFile == "" {
			println(fromFile)
		} else {
			fileMap[fromFile] = targetFile
		}
	}

	println("_______________________________________________________________________")

	for k, v := range fileMap {
		fromRows := ReadFile(k)
		targetRows := ReadFile(v)

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
			println()
			println("#################################################")
			println(k)

			for _, block := range errorList {
				println(block)
			}
		}
	}
}
