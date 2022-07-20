package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// 获取文件的md5码
func getFileMd5(filename string) (string, error) {
	// 文件全路径名
	path := fmt.Sprintf("%s", filename)
	pFile, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("打开文件失败，filename=%v, err=%v", filename, err)
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func main() {
	// 当前目录的csv配置文件为例
	fileName := os.Args[1]
	if md5Val, err := getFileMd5(fileName); err == nil {
		fmt.Println("配置文件的md5值", md5Val)
	} else {
		fmt.Println(err.Error())
	}

}
