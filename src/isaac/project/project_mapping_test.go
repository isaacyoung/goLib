package project_test

import (
	"isaac/file"
	"os"
	"regexp"
	"strings"
	"testing"
)

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

func TestMapping(t *testing.T) {
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

	projectFiles, err := file.GetProjectFiles(projectPath, includePath, includeFile, excludePath, excludeFile)
	if err != nil {
		println(err)
		return
	}

	for _, proFile := range projectFiles {
		currFile.WriteString(end)
		rows, _ := file.ReadFile(proFile)

		classPath := getClassPath(rows)
		currFile.WriteString(classPath + end)
		mapping := getMapping(rows)
		for _, m := range mapping {
			currFile.WriteString(m + end)
		}
	}

}
