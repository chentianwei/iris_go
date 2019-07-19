package util

import (
	"github.com/kataras/golog"
	"os"
)

func init() {
	ReadEnv()
	f := LogFile()
	golog.AddOutput(f)
	os.Stdout = f
	os.Stderr = f
}

func LogFile() *os.File {
	filename := os.Getenv("LOG_PATH") + "server-error.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
