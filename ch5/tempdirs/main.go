package main

import (
	"log"
	"os"
	"time"
)

func main() {

	dirs := os.Args[1:]
	if len(dirs) < 1 {
		log.Fatal("At least one directory is needed.")
		os.Exit(0)
	}
	actions := createDirs(dirs)
	time.Sleep(3 * time.Second)
	removeDirs(actions)
}

func createDirs(dirs []string) []func() {
	var actions []func()
	for _, dir := range dirs {
		d := dir
		err := os.MkdirAll(d, 0755)
		if err == nil {
			actions = append(actions, func() { os.RemoveAll(d) })
		}
	}
	return actions
}

func removeDirs(actions []func()) {
	for _, f := range actions {
		f()
	}
}
