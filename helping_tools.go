package IpMasks

import (
	"math"
)

func numberOfBits(n int) int {
	if n == 0 {
		return 0
	}
	return int(math.Ceil(math.Log2(float64(n + 1))))
}

func AddBits(ip IPv4, bits_add int) IPv4 {
	// Find the position of the first '0' bit in the mask
	mask := IPv4(0b10000000_00000000_00000000_00000000)
	for i := 0; i < 32; i++ {
		if ip&mask == 0 {
			break
		}
		mask = mask >> 1
	}

	// Add '1' bits after the end of the '11....11' part
	for i := 0; i < bits_add; i++ {
		ip = ip | mask
		mask = mask >> 1
	}

	return ip
}

func MaxNumberWithBits(num_bits int) int {
	if num_bits <= 0 {
		return 0
	}
	return (1 << num_bits) - 1
}
