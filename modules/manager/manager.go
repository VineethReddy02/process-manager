package manager

import (
	"fmt"

	"github.com/VineethReddy02/process-manager/modules/process"
)

// Manages the lifecycle of processes
type Manager struct {
	// set of processes
	Processes []*process.Process
	// number of restarts configured
	Restarts int
}

// Create new manager instance
func New(processes []*process.Process, restarts int) Manager {
	return Manager{
		Processes: processes,
		Restarts:  restarts,
	}
}

// Start set of processes and manage the lifecycle
func (c *Manager) StartProcesses() {
	var err error
	// Start the processes
	for i, p := range c.Processes {
		proc := p
		index := i
		go func() {
			*proc = process.New(proc)
			err = proc.Start()
			if err != nil {
				c.dropTheProcess(index)
				fmt.Printf("Failed to create the process for %s\n", proc.Name)
			}
		}()
	}

	// Manage the lifecycle
	c.manageProcesses()
}

// Stop the set of processes
func (c *Manager) StopProcesses() {
	var err error
	for _, proc := range c.Processes {
		err = proc.Stop()
		if err != nil {
			fmt.Printf("Failed to kill the process %s\n", err)
		}
	}
}

// Manages the lifecycle of processes
func (c *Manager) manageProcesses() {
	var err error
	// keep iterating over processes, until all the processes are completed.
	// or till the processes reach maximum restarts.
	for {
		// iterate over the running processes to check their state.
		for index, runningProcess := range c.Processes {
			select {
			case <-runningProcess.Complete:
				fmt.Println("Output:")
				fmt.Println(runningProcess.Command.Stdout)
				c.dropTheProcess(index)
			case <-runningProcess.Crashed:
				if runningProcess.Restarts < c.Restarts {
					restarts := runningProcess.Restarts
					*runningProcess = process.New(runningProcess)
					runningProcess.Restarts = restarts
					err = runningProcess.Start()
					if err != nil {
						c.dropTheProcess(index)
						fmt.Printf("Failed to create the process for %s\n", runningProcess.Name)
					}
					runningProcess.Restarts++
				} else {
					fmt.Printf("Process %s reached maximum restarts. Ignoring the process for restart.\n",
						runningProcess.Name)
					c.dropTheProcess(index)
				}
			default:
				// to avoid deadlock
			}
		}

		if len(c.Processes) == 0 {
			break
		}
	}
}

// Drops the process from list of managing processes.
func (c *Manager) dropTheProcess(index int) {
	if len(c.Processes) > index {
		c.Processes = append(c.Processes[:index], c.Processes[index+1:]...)
	}
}
