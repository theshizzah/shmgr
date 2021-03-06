package main

import (
	"os"
	"path/filepath"

	profilemanager "github.com/theshizzah/shmgr/profile/manager"
)

var baseDir = filepath.Join(os.Getenv("HOME"), ".shellmgr")

func usage() {
	println("./shellmgr init | env ENV_NAME | use ENV_NAME")
	os.Exit(1)
}

func doInit() {
	println("initializing shellmgr!")
	if err := os.Mkdir(baseDir, os.ModePerm); err != nil {
		println("error creating directory: ", err.Error())
	}
}

func make(name string) {
	println("making env: ", name)
	checkName(name)
	profilePath := filepath.Join(baseDir, name)
	if err := os.Mkdir(profilePath, os.ModePerm); err != nil {
		println("error creating subdir: ", err.Error())
	}
	_ = profilemanager.New(profilePath)
}

func use(name string) {
	checkName(name)
	println("using env: ", name)
}

func checkName(name string) {
	if name == "" {
		usage()
	}
}

func main() {
	println("welcome to shellmgr!")

	switch arg := os.Args[1]; arg {
	case "init":
		doInit()
	case "env":
		make(os.Args[2])
	case "use":
		use(os.Args[2])
	default:
		usage()
	}
}
