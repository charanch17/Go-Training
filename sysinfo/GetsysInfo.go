package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func main() {
	memory, _ := ghw.Memory()
	phys := memory.TotalPhysicalBytes
	fmt.Printf(" %d bytes of RAM\n", phys)
	cpu, _ := ghw.CPU()
	fmt.Printf("processor %+v", cpu.Processors)
}
