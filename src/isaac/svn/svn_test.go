package svn_test

import (
	"github.com/stretchr/testify/assert"
	"isaac/svn"
	"testing"
)

func TestChangedFiles(t *testing.T) {
	remoteUrl := "http://192.168.68.12/svn/app/unifiedcenter/trunk"
	localUrl := "E:\\app\\unifiedcenter"
	fromDate := "2019-9-4"
	includeStr := []string{"unifiedcenter", "jiaofei"}
	excludeStr := []string{"unifiedcenter-web-manager"}
	logs, err := svn.GetChangeFiles(remoteUrl, localUrl, fromDate, includeStr, excludeStr)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, logs)
	for _, log := range logs {
		println(log)
	}
}
