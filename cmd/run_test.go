// Package cmd
/**
* @Project : GenericGo
* @File    : run_test.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/12 17:24
**/

package cmd

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ls -l ./",
			args: args{
				ctx:  context.Background(),
				name: "ls",
				args: []string{"-l", "./"},
			},
		},
		{
			name: "nonexistentcommand",
			args: args{
				ctx:  context.Background(),
				name: "nonexistentcommand",
				args: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdStr := tt.args.name + " " + strings.Join(tt.args.args, " ")
			fmt.Printf("Run cmd: \n%s\n", cmdStr)
			res := Run(tt.args.ctx, tt.args.name, tt.args.args...)
			fmt.Println(res2Str(res))
		})
	}
}

func res2Str(res Result) string {
	return fmt.Sprintf("Stdout: \n%s\nErr: \n%v", res.Stdout, res.Err)
}
