package main

import "os"

func usage() {
	println("./shellmgr init | env ENV_NAME | use ENV_NAME")
	os.Exit(1)
}

func doInit() {
	println("initializing shellmgr!")
	if err := os.Mkdir(os.Getenv("HOME")+"/.shellmgr", os.ModePerm); err != nil {
		println("error creating directory: ", err.Error())
	}
}

func make(name string) {
	checkName(name)
	println("making env: ", name)
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
