package sh

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// Cmd represents a single shell command.
type Cmd struct {
	Path    string // path to the binary you want to run
	Args    []string
	Timeout time.Duration
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Env     []string
}

// Command initializes a Cmd. It will panic if timeout is negative.
func Command(timeout time.Duration, name string, args ...string) *Cmd {
	if timeout < 0 {
		msg := fmt.Sprintf("negative timeout: %v", timeout)
		panic(msg)
	}

	return &Cmd{
		Path:    name,
		Args:    args,
		Timeout: timeout,
	}
}

// Run executes the shell command.
func (c *Cmd) Run() error {
	if c.Timeout < 0 {
		return fmt.Errorf("negative timeout: %v", c.Timeout)
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c.Path, c.Args...)
	cmd.Stdin = os.Stdin
	if c.Stdin != nil {
		cmd.Stdin = c.Stdin
	}

	cmd.Stdout = os.Stdout
	if c.Stdout != nil {
		cmd.Stdout = c.Stdout
	}

	cmd.Stderr = os.Stderr
	if c.Stderr != nil {
		cmd.Stderr = c.Stderr
	}

	cmd.Env = c.Env

	return cmd.Run()
}

// RedirectStdout will copy the stdout output from the command to the given
// writer. mycmd.RedirectStdout(myFile) is equivalent to ./mycmd > myfile.
func (c *Cmd) RedirectStdout(w io.Writer) {
	if w == nil {
		panic("passed a nil writer to sh.Cmd.RedirectStdout()")
	}

	c.Stdout = w
}

// RedirectStderr will copy the stderr output from the command to the given
// writer. mycmd.RedirectStderr(myFile) is equivalent to ./mycmd 2> myfile.
func (c *Cmd) RedirectStderr(w io.Writer) {
	if w == nil {
		panic("passed a nil writer to sh.Cmd.RedirectStderr()")
	}

	c.Stderr = w
}
