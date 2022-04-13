package models

import (
	"fmt"
	"os"
)

func CreateMutiDir(filePath string) error {
	if !JudgeDirExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("Create folder failed ,error info:", err)
			return err
		}
		os.Chmod(filePath, 0777)
		return err
	}
	return nil
}
