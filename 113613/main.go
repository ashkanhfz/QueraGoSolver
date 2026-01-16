package main

import (
	"sync"
)

func Solution(f func(uint8) uint8, inp uint32) uint32 {
	b0 := uint8(inp & 0xFF)
	b1 := uint8((inp >> 8) & 0xFF)
	b2 := uint8((inp >> 16) & 0xFF)
	b3 := uint8((inp >> 24) & 0xFF)


	inputs := [4]uint8{b0, b1, b2, b3}
	results := [4]uint8{}

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(index int, inputByte uint8) {
			defer wg.Done()
			results[index] = f(inputByte)
		}(i, inputs[i])
	}
	wg.Wait()
	var out uint32
	out |= uint32(results[0])
	out |= uint32(results[1])<<8
	out |= uint32(results[2])<<16
	out |= uint32(results[3])<<24

	return out
	
}
