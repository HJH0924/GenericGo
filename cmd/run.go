// Package cmd
/**
* @Project : GenericGo
* @File    : run.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/10/12 16:10
**/

package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Result struct {
	Stdout string // 标准输出
	Stderr string // 标准错误
	Code   int    // 错误码
	Err    error  // 包装错误
}

// Run 在给定的上下文中执行指定的命令，并返回执行结果。
// 它使用 context.Context 来支持取消操作，允许从运行的命令中传递信号。
// 此函数封装了 os/exec 包的 CommandContext 方法，简化了命令执行的过程。
func Run(ctx context.Context, name string, args ...string) Result {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
		err    error
	)

	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()

	stdoutStr := stdout.String()
	stderrStr := stderr.String()
	code := cmd.ProcessState.ExitCode()
	if code != 0 {
		err = fmt.Errorf("received error code %d for stderr `%s`", code, strings.TrimRight(stderrStr, "\n"))
	}

	return Result{
		Stdout: stdoutStr,
		Stderr: stderrStr,
		Code:   code,
		Err:    err,
	}
}
