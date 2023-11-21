package command

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func output(p string, reader io.ReadCloser) error {
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if num > 0 {
			fmt.Printf("%s%s", p, string(buf[:num]))
		}
	}
	//return nil
}

func Exec(e string, arg ...string) error {

	return ExecPrefix("", e, arg...)
}

func ExecPrefix(p, e string, arg ...string) error {
	cmd := exec.Command(e, arg...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		log.Printf("Error executing command: %s......\n", err.Error())
		return err
	}

	go output(p, stdout)
	go output(p, stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......\n", err.Error())
		return err
	}

	log.Printf("done")

	return nil
}
