package mocks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

const (
	testDataFolder = "testdata"
)

func readSample(pathParts ...string) ([]byte, error) {
	_, caller, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to determine the caller file")
	}
	pathParts = append([]string{filepath.Dir(caller), testDataFolder}, pathParts...)
	return ioutil.ReadFile(filepath.Join(pathParts...))
}
