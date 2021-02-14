package filex

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func OpenFile(filename string) (*os.File, error) {
	if !FileExists(filename) {
		err := os.MkdirAll(filepath.Dir(filename), 0755)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	if f == nil {
		return nil, errors.New("open file must success")
	}

	return f, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func FileEmpty(name string) bool {
	stat, err := os.Stat(name)
	return os.IsNotExist(err) || stat.Size() <= 0
}

func ReadFile(path string) string {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return string(buff)
}
