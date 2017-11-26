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

func (s sieve) clear(n uint64) {
	s[n/64] &^= 1 << (n % 64)
}

var wheel = []uint64{1, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 53, 59}

// Generate returns all of the prime numbers equal to or less than the target.
func Generate(out chan<- uint64, target uint64) {
	sieve := newSieve(target)

	x, xx, y := uint64(0), uint64(0), uint64(0)
	n, r := uint64(0), uint64(0)

	for x = 1; ; x++ {
		xx = 4 * x * x
		if xx > target {
			break
		}

		for y = 1; ; y += 2 {
			n = xx + (y * y)

			if n > target {
				break
			}

			r = n % 60
			if r == 1 || r == 13 || r == 17 || r == 29 || r == 37 || r == 41 || r == 49 || r == 53 {
				sieve.set(n)
			}
		}
	}

	for x = 1; ; x += 2 {
		xx = 3 * x * x
		if xx > target {
			break
		}

		for y = 2; ; y += 2 {
			n = xx + (y * y)

			if n > target {
				break
			}

			r = n % 60
			if r == 7 || r == 19 || r == 31 || r == 43 {
				sieve.set(n)
			}
		}
	}

	for x = 2; x <= target; x++ {
		xx = 3 * x * x

		for y = x - 1; y > 0; y -= 2 {
			n = xx - (y * y)

			if n > target {
				break
			}

			r = n % 60
			if r == 11 || r == 23 || r == 47 || r == 59 {
				sieve.set(n)
			}
		}
	}

	out <- 2
	out <- 3
	out <- 5

	for x = 0; ; x += 60 {
		for _, y = range wheel {
			n = x + y

			if n > target {
				goto END
			}

			if sieve.check(n) {
				out <- n

				for xx = n * n; xx <= target; xx += 2 * n {
					sieve.clear(xx)
				}
			}
		}
	}

END:
	close(out)
}

// IsPrime returns if the provided number is a prime number or not.
func IsPrime(candidate uint64) bool {
	if candidate <= 1 {
		return false
	}

	limit := uint64(math.Sqrt(float64(candidate)))
	for i := uint64(2); i <= limit; i++ {
		if (candidate % i) == 0 {
			return false
		}
	}

	return true
}
