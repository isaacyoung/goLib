package svn

import (
	"github.com/Masterminds/vcs"
	"github.com/axgle/mahonia"
	"strings"
)

func GetSvnLog(remoteUrl string, localUrl string, fromDate string) ([]string, error) {
	repo, _ := vcs.NewRepo(remoteUrl, localUrl)
	out, err := repo.RunFromDir("svn", "log", "-r", "{"+fromDate+"}:HEAD", "-v")
	if err != nil {
		return nil, err
	}
	enc := mahonia.NewDecoder("gbk")
	goStr := enc.ConvertString(string(out))
	arr := strings.Split(goStr, "\r\n")

	return arr, nil
}

func GetChangeFiles(remoteUrl string, localUrl string, fromDate string, containStr []string, excludeStr []string) ([]string, error) {
	logs, err := GetSvnLog(remoteUrl, localUrl, fromDate)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, log := range logs {
		if notContain(log, excludeStr) {
			continue
		}

		if contain(log, containStr) {
			str := strings.Trim(log, " ")[2:]
			if !hasItem(result, str) {
				result = append(result, str)
			}
		}
	}
	return result, nil
}

func contain(content string, arr []string) bool {
	if arr == nil {
		return true
	}
	for _, a := range arr {
		if !strings.Contains(content, a) {
			return false
		}
	}
	return true
}

func notContain(content string, arr []string) bool {
	if arr == nil {
		return false
	}
	for _, a := range arr {
		if strings.Contains(content, a) {
			return true
		}
	}
	return false
}

func hasItem(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(a, str) {
			return true
		}
	}
	return false
}
