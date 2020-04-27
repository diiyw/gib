package gos

import (
	"io/ioutil"
)

// FilePath the filepath tree node
type FileInfo struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"is_dir"`
	Size     int64  `json:"size"`
	ModeTime string `json:"mode_time"`
	Mode     string `json:"mode"`
}

func Dirs(dirPath string) (files []FileInfo, err error) {
	fi, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	for _, f := range fi {
		var file FileInfo
		file.Name = f.Name()
		file.Size = f.Size()
		file.ModeTime = f.ModTime().Format(DateTimeFormat)
		file.Mode = f.Mode().String()
		file.IsDir = f.IsDir()
		files = append(files, file)
	}
	return
}
