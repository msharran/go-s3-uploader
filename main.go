package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

var LogsDir string

func main() {
	if len(os.Args) > 1 {
		LogsDir = os.Args[1]
	} else {
		log.Fatalln("Arg missing error: File path should be provided as 1st argument")
	}

	files, err := ioutil.ReadDir(LogsDir)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, file := range files {
		fileInfo := NewFileInfo(LogsDir + "/" + file.Name())
		createdTime, err := fileInfo.GetCreatedTime()
		if err != nil {
			log.Fatalln(err.Error())
		}

		now := time.Now()
		fileAge := now.Sub(createdTime).Hours() / 24
		log.Printf("\nName = %s\nAge = %.2f days old (%vth %v %v)\n--------------------", file.Name(), fileAge, createdTime.Day(), createdTime.Month(), createdTime.Year())

		if fileAge > 3 {
			log.Println("Uploading " + file.Name() + " to s3...")
			s3Uploader := NewS3Uploader()
			s3Uploader.AddFileToS3(fileInfo.FilePath)
		}
	}
}
