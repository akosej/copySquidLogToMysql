package system

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

// Convert a shell command with a series of pipes into
// correspondingly piped list of *exec.Cmd
// If an arg has spaces, this will fail
func RunString(s string) (string, error) {
	buf := bytes.NewBuffer([]byte{})
	sp := strings.Split(s, "|")
	cmds := make([]*exec.Cmd, len(sp))
	// create the commands
	for i, c := range sp {
		cs := strings.Split(strings.TrimSpace(c), " ")
		cmd := cmdFromStrings(cs)
		cmds[i] = cmd
	}

	cmds = AssemblePipes(cmds, nil, buf)
	if err := RunCmds(cmds); err != nil {
		return "", err
	}

	b := buf.Bytes()
	return string(b), nil
}

func cmdFromStrings(cs []string) *exec.Cmd {
	if len(cs) == 1 {
		return exec.Command(cs[0])
	} else if len(cs) == 2 {
		return exec.Command(cs[0], cs[1])
	}
	return exec.Command(cs[0], cs[1:]...)
}

// AssemblePipes Pipe stdout of each command into stdin of next
func AssemblePipes(cmds []*exec.Cmd, stdin io.Reader, stdout io.Writer) []*exec.Cmd {
	cmds[0].Stdin = stdin
	cmds[0].Stderr = stdout
	// assemble pipes
	for i, c := range cmds {
		if i < len(cmds)-1 {
			cmds[i+1].Stdin, _ = c.StdoutPipe()
			cmds[i+1].Stderr = stdout
		} else {
			c.Stdout = stdout
			c.Stderr = stdout
		}
	}
	return cmds
}

// RunCmds Run series of piped commands
func RunCmds(cmds []*exec.Cmd) error {
	// start processes in descending order
	for i := len(cmds) - 1; i > 0; i-- {
		if err := cmds[i].Start(); err != nil {
			return err
		}
	}
	// run the first process
	if err := cmds[0].Run(); err != nil {
		return err
	}
	// wait on processes in ascending order
	for i := 1; i < len(cmds); i++ {
		if err := cmds[i].Wait(); err != nil {
			return err
		}
	}
	return nil
}

func Run(arg string) {
	cmdArgs := strings.Fields(arg)
	run := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	_ = run.Run()
}

func RunCMD(arg ...string) {
	head := arg[0]
	parts := arg[1:]
	run := exec.Command(head, parts...)
	_ = run.Run()
}