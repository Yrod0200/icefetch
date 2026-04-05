package modules

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func getOs() string {

	os := runtime.GOOS

	return os

}

func getUsername() string {

	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	username := u.Username

	return username

}

func getHost() string {
	content, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		fmt.Println("Error reading host:", err)
	}

	fileString := strings.Replace(string(content), "\n", "", 1)
	return fileString
}

func getLinuxDistro() string {

	file, _ := os.Open("/etc/os-release")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "PRETTY_NAME=") {
			name := strings.TrimPrefix(line, "PRETTY_NAME=")
			name = strings.Trim(name, "\"")
			return name
		}
	}

	return "Unknown Distro / Error"

}

func getHostname() string {
	hostname, _ := os.Hostname()

	return hostname
}

func HostAndUser() string {

	fullname := fmt.Sprintf("%s@%s", getUsername(), getHostname())

	return fullname

}

func Host() string {

	return getHost()

}

func FullOSName() string {

	os := getOs()

	if os == "windows" {
		return "Windows 7/8.x/10/11"
	}

	if os == "darwin" {
		return "Mac OS"
	}

	if os == "linux" {
		return getLinuxDistro()
	}

	if os == "freebsd" {
		return "FreeBSD"
	}

	return "Unknown OS"

}
