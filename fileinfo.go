package main

import (
	"log"
	"os"
	"syscall"
	"time"
)

type FileInfo struct {
	FilePath string
}

func NewFileInfo(fileName string) *FileInfo {
	return &FileInfo{
		FilePath: fileName,
	}
}

func (f *FileInfo) GetCreatedTime() (ctime time.Time, err error) {
	fi, err := os.Stat(f.FilePath)
	if err != nil {
		return
	}
	stat := fi.Sys().(*syscall.Stat_t)
	ctime = time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	log.Printf("\nName = %s\nCreatedDate = %vth %v %v\n--------------------", f.FilePath, ctime.Day(), ctime.Month(), ctime.Year())
	return
}
