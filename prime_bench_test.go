package prime_test

import (
	"testing"

	"github.com/jtraynor/prime"
)

func BenchmarkGenerate(b *testing.B) {
	for name, tc := range checksumTT {
		if !tc.isShort {
			continue
		}

		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				primes := make(chan uint64, tc.limit/1000)
				go prime.Generate(primes, tc.limit)

				sum := uint64(0)
				for number := range primes {
					sum += number
				}
			}
		})
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for name, tc := range isPrimeTT {
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				prime.IsPrime(tc.candidate)
			}
		})
	}
}
