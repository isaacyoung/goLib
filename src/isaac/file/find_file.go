package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
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

func ReadFile(file string) ([]string, error) {
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
