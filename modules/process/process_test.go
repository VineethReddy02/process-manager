package process

import (
	"os/exec"
	"testing"
	"time"
)

func TestProcess_Start(t *testing.T) {
	type fields struct {
		Name     string
		Args     []string
		Envs     []string
		Command  *exec.Cmd
		Complete chan struct{}
		Crashed  chan struct{}
		Restarts int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success-case-on-ls",
			fields: fields{
				Name:     "ls",
				Complete: make(chan struct{}, 1),
				Crashed:  make(chan struct{}, 1),
			},
			wantErr: false,
		},
		{
			name: "fail-case-on-abc",
			fields: fields{
				Name:     "abc",
				Complete: make(chan struct{}, 1),
				Crashed:  make(chan struct{}, 1),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				Name:     tt.fields.Name,
				Args:     tt.fields.Args,
				Envs:     tt.fields.Envs,
				Command:  exec.Command(tt.fields.Name, tt.fields.Args...),
				Complete: tt.fields.Complete,
				Crashed:  tt.fields.Crashed,
				Restarts: tt.fields.Restarts,
			}
			if err := p.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProcess_Stop(t *testing.T) {
	type fields struct {
		Name     string
		Args     []string
		Envs     []string
		Command  *exec.Cmd
		Complete chan struct{}
		Crashed  chan struct{}
		Restarts int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success-sleep",
			fields: fields{
				Name:     "sleep",
				Args:     []string{"5"},
				Complete: make(chan struct{}, 1),
				Crashed:  make(chan struct{}, 1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Process{
				Name:     tt.fields.Name,
				Args:     tt.fields.Args,
				Envs:     tt.fields.Envs,
				Command:  exec.Command(tt.fields.Name, tt.fields.Args...),
				Complete: tt.fields.Complete,
				Crashed:  tt.fields.Crashed,
				Restarts: tt.fields.Restarts,
			}
			// create the process, before stopping it.
			go func() {
				err := p.Start()
				if err != nil {
					t.Errorf("Failed to start the process %v with error %v", p.Name, err)
				}
			}()
			time.Sleep(1 * time.Second)
			if err := p.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
