package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func getFiles(path string, includeFile_ []string, excludePath_ []string, excludeFile_ []string) (files []string, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		println(err)
		return files, err
	}

	sep := string(os.PathSeparator)

	for _, file := range dir {
		p := path + sep + file.Name()
		if file.IsDir() {
			if !excludePath(p, excludePath_) {
				files2, err := getFiles(p, includeFile_, excludePath_, excludeFile_)
				if err != nil {
					println(err)
					return files, err
				}
				files = append(files, files2...)
			}

		} else {
			if includeFile(file.Name(), includeFile_) && !excludeFile(file.Name(), excludeFile_) {
				files = append(files, p)
			}

		}
	}

	return files, nil
}

func includePath(path string, includePath []string) bool {
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

func includeFile(file string, includeFile []string) bool {
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

func excludePath(path string, excludePath []string) bool {
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

func excludeFile(file string, excludeFile []string) bool {
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

func getProjectFiles(path string, includePath_ []string, includeFile []string, excludePath []string, excludeFile []string) ([]string, error) {
	files, err := getFiles(path, includeFile, excludePath, excludeFile)
	if err != nil {
		return nil, err
	}

	var fileList []string
	for _, f := range files {
		if includePath(f, includePath_) {
			fileList = append(fileList, f)
		}
	}
	return fileList, nil
}

func readFile(file string) ([]string, error) {
	fd, err := os.Open(file)
	defer fd.Close()
	if err != nil {
		println("read error:", err)
		return nil, err
	}

	var rows []string
	buff := bufio.NewReader(fd)

	for {
		data, _, eof := buff.ReadLine()
		if eof == io.EOF {
			break
		}

		rows = append(rows, string(data))
	}
	return rows, nil
}

func getClassPath(rows []string) string {
	var pkg string
	var className string
	for _, row := range rows {
		if strings.HasPrefix(row, "package") {
			pkg = row[8 : len(row)-1]
		} else if strings.HasPrefix(row, "public class") {
			className = row[13:]
			className = className[:strings.Index(className, " ")]
			break
		}
	}
	return pkg + "." + className
}

func getMapping(rows []string) []string {
	var classMapping string
	var headEnd = false
	var result []string
	mapStr := `RequestMapping\("(.+)"\)`
	mapReg := regexp.MustCompile(mapStr)
	mapStr2 := `value = "(.+)"`
	mapReg2 := regexp.MustCompile(mapStr2)

	for _, row := range rows {
		if strings.HasPrefix(strings.TrimLeft(row, " "), "public class ") {
			headEnd = true
		}
		if strings.HasPrefix(strings.TrimLeft(row, " "), "@RequestMapping") {
			var arr []string
			arr = mapReg.FindStringSubmatch(row)
			if arr != nil {
				if !headEnd {
					classMapping = arr[len(arr)-1]
				} else {
					result = append(result, classMapping+arr[len(arr)-1])
				}

				continue
			}

			arr = mapReg2.FindStringSubmatch(row)
			if arr != nil {
				if !headEnd {
					classMapping = arr[len(arr)-1]
				} else {
					result = append(result, classMapping+arr[len(arr)-1])
				}
				continue
			}
		}
	}
	return result
}

func main() {
	projectPath := "E:\\app\\unifiedcenter"
	includePath := []string{"controller"}
	includeFile := []string{".java"}
	excludePath := []string{".svn", ".idea", "target", "unifiedcenter-generate"}
	excludeFile := []string{"Test.java"}
	end := "\r\n"

	reportFile := "F:\\dir\\mapping.txt"
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

	projectFiles, err := getProjectFiles(projectPath, includePath, includeFile, excludePath, excludeFile)
	if err != nil {
		println(err)
		return
	}

	for _, proFile := range projectFiles {
		currFile.WriteString(end)
		rows, _ := readFile(proFile)

		classPath := getClassPath(rows)
		currFile.WriteString(classPath + end)
		mapping := getMapping(rows)
		for _, m := range mapping {
			currFile.WriteString(m + end)
		}
	}

}
