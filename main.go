package main

import "os"

func usage() {
	println("./shellmgr init | env ENV_NAME | use ENV_NAME")
	os.Exit(1)
}

func main() {
	println("welcome to shellmgr!")

	if os.Args[1] == "init" {
		println("initializing shellmgr!")
		if err := os.Mkdir(os.Getenv("HOME")+"/.shellmgr", os.ModePerm); err != nil {
			println("error creating directory: ", err.Error())
		}
	} else if os.Args[1] == "env" {
		if "" == os.Args[2] {
			usage()
		}
		print("creating shell environment named ", os.Args[2])
	} else if os.Args[2] == "use" {
		if "" == os.Args[2] {
			usage()
		}
	}
}
