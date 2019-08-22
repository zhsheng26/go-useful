package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// 同步文件夹,保证对应的两个文件夹的内容一样
func syncDirs(dirKv map[string]string) {

}

func copyDirFile(dirSrc string, dirDest string) {
	for _, file := range walkDir(dirSrc) {
		src, err := os.Open(dirSrc + string(filepath.Separator) + file)
		if err != nil {
			return
		}
		dst, err := os.OpenFile(dirDest+string(filepath.Separator)+file, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		_, err = io.Copy(dst, src)
		_ = src.Close()
		_ = dst.Close()
	}
}

// recursion
// return file absolute path
func readListFile(folder string) []string {
	var list []string
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			readListFile(folder + "/" + file.Name())
		} else {
			list = append(list, folder+"/"+file.Name())
		}
	}
	return list
}

// return file relative path
func walkDir(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err.Error())
	}
	return files
}

//if filepath.Ext(path) == ".dat" {
//	return nil
//}
//if info.IsDir() {
//	return nil
//}
func walkDirFilter(dir string, skip func(path string, info os.FileInfo) bool) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if skip(path, info) {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Println(err.Error())
	}
	return files
}
