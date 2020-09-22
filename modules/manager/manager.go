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
	// Start the processes
	for _, p := range c.Processes {
		proc := p
		go func() {
			*proc = process.New(proc)
			proc.Start()
		}()
	}

	// Manage the lifecycle
	c.manageProcesses()
}

// Stop the set of processes
func (c *Manager) StopProcesses() {
	for _, proc := range c.Processes {
		proc.Stop()
	}
}

// Manages the lifecycle of processes
func (c *Manager) manageProcesses() {
	for {
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
					runningProcess.Start()
					runningProcess.Restarts++
				} else {
					fmt.Println(fmt.Sprintf("Process %s reached maximum restarts. Ignoring the process for restart.",
						runningProcess.Name))
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
	c.Processes = append(c.Processes[:index], c.Processes[index+1:]...)
}
