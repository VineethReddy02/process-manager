package manager

import (
	"testing"
	"time"

	"github.com/VineethReddy02/process-manager/modules/process"
)

func TestManager_Start_Stop_Processes(t *testing.T) {
	type fields struct {
		Processes []*process.Process
		Restarts  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "EXAMPLE_TESTCASES",
			fields: fields{
				Processes: []*process.Process{
					{
						Name: "sleep",
						Args: []string{"5"},
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
				},
				Restarts: 2,
			},
		},
	}
	for _, tt := range tests {
		c := &Manager{}
		counter := 0
		t.Run(tt.name, func(t *testing.T) {
			c = &Manager{
				Processes: tt.fields.Processes,
				Restarts:  tt.fields.Restarts,
			}
			go func() {
				c.StartProcesses()
			}()
			for {
				// sleep for 3 secs before killing the process
				time.Sleep(2 * time.Second)
				c.StopProcesses()
				if counter == 5 {
					break
				}
				counter++
			}
		})
		if len(c.Processes) != 0 && c.Restarts == 2 {
			t.Errorf("Failed to complete processes %v", len(c.Processes))
		}

	}
}
