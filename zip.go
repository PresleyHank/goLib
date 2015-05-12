/**
 *  @usage
 *      zf := goLib.Zip("path/to/myFile.zip")
 *      zf.Extract("path/to/dest_folder/")
 *
 *  @author
 *      kilfu0701
 */
package goLib

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func cloneZipItem(f *zip.File, dest string) {
	// Create full directory path
	path := filepath.Join(dest, f.Name)
	fmt.Println("Creating", path)
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir|os.ModePerm)
	checkError(err)

	// Clone if item is a file
	rc, err := f.Open()
	checkError(err)
	if !f.FileInfo().IsDir() {
		// Use os.Create() since Zip don't store file permissions.
		fileCopy, err := os.Create(path)
		checkError(err)
		_, err = io.Copy(fileCopy, rc)
		fileCopy.Close()
		checkError(err)
	}
	rc.Close()
}

type ZipFile struct {
	path string
}

func Zip(zip_path string) *ZipFile {
	z := &ZipFile{path: zip_path}
	return z
}

func (z *ZipFile) Extract(dest string) {
	r, err := zip.OpenReader(z.path)
	checkError(err)
	defer r.Close()
	for _, f := range r.File {
		cloneZipItem(f, dest)
	}
}
