package image

import (
	"path"
	"runtime"
)

func ImagePath() string {
	_, file, _, _ := runtime.Caller(0)
	p := path.Dir(file)
	return p + "/"
}
