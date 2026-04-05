package modules

import "syscall"

func Kernel() string {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		return "Unknown"
	}

	b := make([]byte, 0, len(uname.Release))
	for _, v := range uname.Release {
		if v == 0 {
			break
		}
		b = append(b, byte(v))
	}

	return string(b)
}
