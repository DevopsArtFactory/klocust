/*
Copyright 2020 The klocust Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
		return fmt.Errorf("`%s` directory exists already", dirName)
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

	dir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("status code: %v", response.StatusCode)
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
		return "", fmt.Errorf("%s is not a regular file", src)
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
		return "", fmt.Errorf("file %s already exists", src)
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

func GetSha256Checksum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func DeleteFile(filename string) error {
	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}

func GetFileSize(filename string) (int64, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}
