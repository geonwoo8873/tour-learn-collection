package main

import (
	"os"   // https://pkg.go.dev/os#File = OS file type
	"time" // https://pkg.go.dev/time#Time = Time type for timestamps
)

type FileSystem struct {
	path      string
	file      *os.File
	timestamp time.Time
}

func main() {
	fs, err := FileSystem{
		path:      "",
		file:      nil,
		timestamp: time.Now(),
	}
	if err != nil {
		panic(err)
	}
}
