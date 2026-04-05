package modules

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func Ram() string {

	ram, _ := mem.VirtualMemory()

	return fmt.Sprintf("%vMiB / %vMiB", ram.Used/1024/1024, ram.Total/1024/1024)

}

func Cpu() string {

	cpu, _ := cpu.Info()

	return fmt.Sprintf("%s", cpu[0].ModelName)

}
