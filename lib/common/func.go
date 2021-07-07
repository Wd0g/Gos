package common

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func Shell(command string, cmdPath string) (res string, err error) {
	var execCmd *exec.Cmd
	if runtime.GOOS == "windows" {
		execCmd = exec.Command(cmdPath, "/c", command)
	} else {
		execCmd = exec.Command(cmdPath, "-c", command)
	}
	var (
		stdout bytes.Buffer
	)

	execCmd.Stdout = &stdout
	execCmd.Stderr = &stdout
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

func CopyFile(srcFile, dstFile string) (err error) {
	sinfo, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	if sinfo.IsDir() {
		return errors.New("copy file error, src file can not be dir")
	}

	dinfo, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	srcData, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return err
	}

	err = dinfo.Chmod(sinfo.Mode())
	if err != nil {
		return err
	}
	_, err = dinfo.Write(srcData)
	if err != nil {
		return err
	}
	err = dinfo.Close()
	if err != nil {
		return err
	}
	err = os.Chtimes(dstFile, sinfo.ModTime(), sinfo.ModTime())
	if err != nil {
		return err
	}

	return nil
}

func GetDataFromUrl(srcUrl string) (res []byte, err error) {
	resp, err := http.Get(srcUrl)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func GetDiskList() (res string) {
	rawList := []string{"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	diskList := ""
	for _, d := range rawList {
		if _, err := os.Stat(fmt.Sprintf("%s:/", d)); !os.IsNotExist(err) {
			diskList += fmt.Sprintf("%s:", d)
		}
	}
	return diskList
}
