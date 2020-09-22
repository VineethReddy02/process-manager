package process

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Metadata required to create a process
type Process struct {
	Name     string
	Args     []string
	Envs     []string
	Command  *exec.Cmd
	Complete chan struct{}
	Crashed  chan struct{}
	Restarts int
}

// New process instance
func New(p *Process) Process {
	return Process{
		Name:     p.Name,
		Args:     p.Args,
		Envs:     p.Envs,
		Command:  exec.Command(p.Name, p.Args...),
		Complete: make(chan struct{}, 1),
		Crashed:  make(chan struct{}, 1),
		Restarts: 0,
	}
}

// Start the process
func (p *Process) Start() {
	var stdout, stderr bytes.Buffer
	p.Command.Stdout = &stdout
	p.Command.Stderr = &stderr
	err := p.Command.Start()
	if err != nil {
		// we failed to start the process
		fmt.Println(fmt.Sprintf("Failed to create the process for %s", p.Name))
		return
	}
	fmt.Println(fmt.Sprintf("Created the process for %s with args as %v PID as %v", p.Name, p.Args, p.Command.Process.Pid))

	err = p.Command.Wait()
	if err != nil {
		// process is killed
		fmt.Println(fmt.Sprintf("Process %s is crashed", p.Name))
		p.Crashed <- struct{}{}
		return
	}

	fmt.Println(fmt.Sprintf("\nProcess %s is successfully completed", p.Name))
	p.Complete <- struct{}{}

}

// Stop the process
func (p *Process) Stop() {
	err := p.Command.Process.Kill()
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to kill the process %s", err))
	}
	fmt.Println(fmt.Sprintf("Successfully killed process %s with args %v", p.Name, p.Args))
}
