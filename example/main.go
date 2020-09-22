package main

import (
	"time"

	"github.com/VineethReddy02/process-manager/modules/manager"
	"github.com/VineethReddy02/process-manager/modules/process"
)

func main() {
	processes := []*process.Process{
		{
			Name: "sleep",
			Args: []string{"10"},
		},
		{
			Name: "ls",
		},
		{
			Name: "go",
			Args: []string{"version"},
		},
		{
			Name: "echo",
			Args: []string{"abc"},
		},
		{
			Name: "abc",
		},
	}

	m := manager.New(processes, 5)

	go func() {
		m.StartProcesses()
	}()

	for i:=0; i<5; i++ {
		// sleep for 2 secs before killing the process
		time.Sleep(2 * time.Second)
		m.StopProcesses()
	}
}
