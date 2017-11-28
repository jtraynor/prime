package prime

import "math"

type sieve []uint64

func newSieve(size uint64) sieve {
	return make([]uint64, (size/64)+1)
}

func (s sieve) check(n uint64) bool {
	return ((s[n/64] >> (n % 64)) & 1) > 0
}

func (s sieve) set(n uint64) {
	s[n/64] |= 1 << (n % 64)
}

// Generate returns all of the prime numbers equal to or less than the target.
func Generate(out chan<- uint64, target uint64) {
	sieve := newSieve(target)

	i := uint64(0)
	j := uint64(0)

	sqrt := uint64(math.Sqrt(float64(target)))
	for i = 2; i <= sqrt; i++ {
		if !sieve.check(i) {
			for j = i * i; j <= target; j += i {
				sieve.set(j)
			}
			out <- i
		}
	}

	for ; i <= target; i++ {
		if !sieve.check(i) {
			out <- i
		}
	}

	close(out)
}

// IsPrime returns if the provided number is a prime number or not.
func IsPrime(candidate uint64) bool {
	if candidate <= 1 {
		return false
	}

	sqrt := uint64(math.Sqrt(float64(candidate)))
	for i := uint64(2); i <= sqrt; i++ {
		if (candidate % i) == 0 {
			return false
		}
	}

	return true
}
