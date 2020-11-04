package klocust

import "fmt"

type FileExistsError struct {
	Filename string
}

func (e FileExistsError) Error() string {
	return fmt.Sprintf("`%s` file is already exists.", e.Filename)
}

func NewFileExistsError(filename string) FileExistsError {
	return FileExistsError{Filename: filename}
}
