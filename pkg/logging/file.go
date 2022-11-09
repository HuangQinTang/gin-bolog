package logging

import (
	"blog/pkg/setting"
	"fmt"
	"log"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

// 获取日志文件路径及文件名
func getLogFileFullPath() string {
	//fmt.Println("----")
	//fmt.Println("LogSavePath", LogSavePath)
	//fmt.Println("LogSaveName", LogSaveName)
	//fmt.Println("LogFileExt", LogFileExt)
	//fmt.Println("TimeFormat", TimeFormat)
	//fmt.Println("----")
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
