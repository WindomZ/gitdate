package gitdate

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func pipeCommand(name string, arg ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	if err := cmd.Wait(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), nil
}

func execCommand(name string, arg ...string) (string, error) {
	stdout, stderr, err := pipeCommand(name, arg...)
	os.Stdout.WriteString(stdout)
	os.Stderr.WriteString(stderr)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return strings.TrimSpace(stdout + stderr), err
	} else if len(stderr) != 0 {
		return stdout, errors.New(stderr)
	}
	return stdout, nil
}

func execCommandStatement(statement string) (string, error) {
	return execCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

func execCommandSplit(statement string) (string, error) {
	strs := regexp.MustCompile(`([^\\]"[\w\s\S]+?[^\\]")|([^\s\\]+)`).
		FindAllString(statement, -1)
	for i, str := range strs {
		strs[i] = strings.TrimSpace(str)
	}
	if len(strs) <= 0 {
		return execCommand("")
	} else if len(strs) > 1 {
		return execCommand(strs[0], strs[1:]...)
	}
	return execCommand(strs[0])
}
