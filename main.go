package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var LogsDir string

func main() {
	logsDir := flag.String("d", "", "Logs directory")
	age := flag.Float64("age", 3, "Files greater than the age will be uploaded to s3")
	flag.Parse()
	fmt.Println(*logsDir, *age)
	if *logsDir == "" {
		log.Fatalln("flag missing ===> -dir=<logs-directory>")
	}
	if *age == 0 {
		log.Fatalln("flag missing ===> -age=<number in days>")
	}

	files, err := ioutil.ReadDir(*logsDir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, file := range files {
		fileInfo := NewFileInfo(*logsDir + file.Name())
		createdTime, err := fileInfo.GetCreatedTime()
		if err != nil {
			log.Fatalln(err.Error())
		}

		now := time.Now()
		fileAge := now.Sub(createdTime).Hours() / 24

		if fileAge > *age {
			s3Uploader := NewS3Uploader()
			s3Uploader.AddFileToS3(fileInfo.FilePath)
		}
	}
}
