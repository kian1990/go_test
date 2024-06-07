package controllers

import (
    "os"
    "path/filepath"
)

func GetFile(folderPath string) ([]string, error) {
    var file []string
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            file = append(file, info.Name())
        }
        return nil
    })
    return file, err
}