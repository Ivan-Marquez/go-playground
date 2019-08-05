package ch02

// pc[i]is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
 * PopCount returns the number of set bits, that is, bits whose value
 * is 1, in a uint64 value, which is called its population count. It
 * uses an init function to precompute a table of results, pc, for
 * each possible 8-bit value so that the PopCount function needn’t
 * take 64 steps but can just return the sum of eight table lookups.
 * (This is definitely not the fastest algorithm for counting bits,
 * but it’s convenient for illustrating init functions, and for
 * showing how to precompute a table of values, which is often a
 * useful programming technique.)
 */
