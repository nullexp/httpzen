package misc

import (
	"path/filepath"
	"runtime"
)

func GetRootPrjPath() string {
	_, b, _, _ := runtime.Caller(0)

	dir := filepath.Dir(b)

	path := filepath.Join(dir, "../../..")

	return path
}

func GetPathFromRoot(path string) string {
	return filepath.Join(GetRootPrjPath(), path)
}
