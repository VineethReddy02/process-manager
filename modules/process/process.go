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
func (p *Process) Start() error {
	var stdout, stderr bytes.Buffer
	p.Command.Stdout = &stdout
	p.Command.Stderr = &stderr
	err := p.Command.Start()
	if err != nil {
		// we failed to start the process
		return err
	}
	fmt.Printf("Created the process for %s with args as %v PID as %v and the restart count as %v\n",
		p.Name, p.Args, p.Command.Process.Pid, p.Restarts)

	err = p.Command.Wait()
	if err != nil {
		// process is killed
		fmt.Printf("Process %s is crashed\n", p.Name)
		p.Crashed <- struct{}{}
		return nil
	}

	fmt.Printf("Process %s is successfully completed\n", p.Name)
	p.Complete <- struct{}{}
	return err
}

// Stop the process
func (p *Process) Stop() error {
	err := p.Command.Process.Kill()
	if err != nil {
		return err
	}
	fmt.Printf("Successfully killed process %s with args %v\n", p.Name, p.Args)
	return err
}
