package main

import (
	"io/ioutil"
	"path/filepath"
)

type FileList struct {
	Dir             string
	CurrentFileList []string
	UnderFileList   []*FileList
}

func ListFiles(path string) (*FileList, error) {
	return listFiles(path, path, path)
}

func listFiles(rootPath string, searchPath string, currentPath string) (*FileList, error) {
	fis, err := ioutil.ReadDir(searchPath)
	if err != nil {
		return nil, err
	}

	fileList := &FileList{}
	fileList.Dir = currentPath
	for _, fi := range fis {
		if isSkip(fi.Name()) {
			continue
		}

		fullPath := filepath.Join(searchPath, fi.Name())
		if fi.IsDir() {
			ulf, err := listFiles(rootPath, fullPath, fi.Name())
			if err != nil {
				return nil, err
			}
			fileList.UnderFileList = append(fileList.UnderFileList, ulf)
		} else {
			fileList.CurrentFileList = append(fileList.CurrentFileList, fi.Name())
		}
	}
	return fileList, nil
}

func isSkip(fileName string) bool {
	switch fileName {
	case ".idea", ".git":
		return true
	default:
		return false
	}
}
