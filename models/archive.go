package models

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Package as a ZIP file
func Zip(src_dir string, zip_file_name string) {

	// Prevention: Old files cannot be overwritten
	os.RemoveAll(zip_file_name)

	// Create :zip file
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()

	// Open the :zip file
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// Path information was traversed
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {

		// If it is a source path, perform the next traversal in advance
		if path == src_dir {
			return nil
		}

		// Get: file header information
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, src_dir+`\`)

		//Judgment: the file is not a folder
		if info.IsDir() {
			header.Name += `/`
		} else {
			// Set: Zip file compression algorithm
			header.Method = zip.Deflate
		}

		// Create: Compress header information
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}
