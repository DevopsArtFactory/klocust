package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"k8s.io/klog/v2"
)

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsDirExists(dirName string) bool {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateDir(dirName string) error {
	if IsDirExists(dirName) {
		return errors.New(fmt.Sprintf("`%s` directory exists already.", dirName))
	}
	if err := os.MkdirAll(dirName, 0700); err != nil {
		return err
	}
	return nil
}

func DownloadFile(url string, dstPath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		err = response.Body.Close()
	}()

	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code %v", response.StatusCode))
	}

	file, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return err
}

func PrintFile(filename string, isPrintFilename bool) error {
	if isPrintFilename {
		klog.Infof("> Filename: %s\n", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)
	if _, err = io.CopyBuffer(os.Stdout, file, buffer); err != nil {
		return err
	}
	return nil
}

func CopyFile(src, dst string, bufferSize int64) (string, error) {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return "", err
	}

	if !srcFileStat.Mode().IsRegular() {
		return "", errors.New(fmt.Sprintf("%s is not a regular file.", src))
	}

	source, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer func() {
		err = source.Close()
	}()

	_, err = os.Stat(dst)
	if err == nil {
		return "", errors.New(fmt.Sprintf("File %s already exists.", src))
	}

	destination, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer func() {
		err = destination.Close()
	}()

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return "", err
		}
	}

	return dst, err
}
