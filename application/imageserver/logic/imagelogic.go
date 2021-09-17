package logic

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

var Enter = "\n"

type (
	ImageLogic struct {
		FilePath string `json:"file_path"`
	}
	RegisterRequest struct {
		GidList []string `json:"gid_list" binding:"required"`
	}
	RegisterResponse struct {
		Message string `json:"message"`
	}
)

func New(filePath string) *ImageLogic {

	return &ImageLogic{FilePath: filePath}
}

func (l *ImageLogic) Register(r *RegisterRequest) (*RegisterResponse, error) {
	flagInt := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	var existGidString string
	if l.Exists(l.FilePath) {
		flagInt = os.O_WRONLY | os.O_APPEND
		bs, err := ioutil.ReadFile(l.FilePath)
		if err != nil {
			return nil, err
		}
		existGidString = string(bs)

	}
	file, err := os.OpenFile(l.FilePath, flagInt, 0666)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
		}
	}()
	write := bufio.NewWriter(file)
	for _, s := range r.GidList {
		if strings.Contains(existGidString, s) {
			continue
		}
		_, err = write.WriteString(s + Enter)
		if err != nil {
			return nil, err
		}
	}
	if err = write.Flush(); err != nil {
		return nil, err
	}
	return &RegisterResponse{
		Message: "SUCCESS",
	}, nil
}

func (l *ImageLogic) Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
