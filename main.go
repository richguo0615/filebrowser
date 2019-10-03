package main

import (
	"runtime"

	"github.com/richguo0615/filebrowser/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
