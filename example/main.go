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
	}

	m := manager.New(processes, 3)

	go func() {
		m.StartProcesses()
	}()

	counter := 0
	for {
		// sleep for 3 secs before killing the process
		time.Sleep(2 * time.Second)
		m.StopProcesses()
		if counter == 5 {
			break
		}
		counter++
	}
}
