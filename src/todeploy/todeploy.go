package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("explorer", "F:\\LIANLIAN_DAYLY\\生产环境\\升级申请\\升级文件列表")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
