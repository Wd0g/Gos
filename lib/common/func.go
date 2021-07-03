package common

import (
	"bytes"
	"encoding/base64"
	"os/exec"
)

func Shell(cmd string) (res string, err error) {
	var execCmd *exec.Cmd
	execCmd = exec.Command("sh", "-c", cmd)
	var (
		stdout bytes.Buffer
	)

	execCmd.Stdout = &stdout
	err = execCmd.Run()
	if err != nil {
		return "", err
	} else {
		return stdout.String(), nil
	}
}

func Base64Encode(src string) (res string) {
	srcBytes := []byte(src)
	return base64.StdEncoding.EncodeToString(srcBytes)
}

func Base64Decode(src string) (res string) {
	srcBytes, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	} else {
		return string(srcBytes)
	}
}
