package main

import {
	"fmt"
	"os"
	"os/exec"
	"syscall"
}

func main(){
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("No valid arguments provided.")
	}
}

func parent(){

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	// Runs an in-memory image of the current executable, passing in child as the first argument
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.Clone_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS
	}
	// Run the child program in the UTS, PID and MNT namespaces

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Prinln("Error: ", err)
		os.Exit(1)
	}
}

func child(){

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Prinln("Error: ", err)
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}