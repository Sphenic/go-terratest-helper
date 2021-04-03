package terratest

import (
	"runtime"
)

func GetTestFilePath() string {

	_,file,_,_ := runtime.Caller(1)
	return file
}
