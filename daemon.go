/*
Makefile: replace all "tidb-server/main.go" to "tidb-server/*.go"
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	daemon = flagBoolean("d", false, "run in daemon mode")
)

func init() {
	flag.Parse()
	Daemon()
}

func Daemon() {
	if *daemon {
		Self, err := os.Executable()
		if err != nil {
			panic(err)
		}
		os.Args = DeleteKeyInSlice(0, os.Args)
		os.Args = DeleteInSlice("-d", os.Args)
		cmd := exec.Command(Self, os.Args...)
		if err = cmd.Start(); err != nil {
			panic(err)
		}
		fmt.Println("Running application in daemon mode", cmd.Process.Pid)
		os.Exit(1)
	}
}

func DeleteInSlice(str string, list []string) (_list []string) {
	for _, s := range list {
		if s != str {
			_list = append(_list, s)
		}
	}
	return _list
}

func DeleteKeyInSlice(key int, list []string) (_list []string) {
	for i, s := range list {
		if i != key {
			_list = append(_list, s)
		}
	}
	return _list
}
