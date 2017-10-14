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
		var err error
		var cmd *exec.Cmd

		Self, err := os.Executable()
		if err != nil {
			panic(err)
		}
		os.Args = DeleteInSlice(Self, os.Args)
		os.Args = DeleteInSlice("-d", os.Args)
		cmd = exec.Command(Self, os.Args...)
		if err = cmd.Start(); err != nil {
			panic(err)
		}
		fmt.Println("Running application in daemon mode", cmd.Process.Pid)
		os.Exit(1)
	}
}

func DeleteInSlice(str string, list []string) []string {
	var newlist []string
	for _, s := range list {
		if s != str {
			newlist = append(newlist, s)
		}
	}
	return newlist
}
