package prime_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/jtraynor/prime"
)

var checksumTT = map[string]struct {
	limit    uint64
	checksum string
	isShort  bool
}{
	"1 Million": {
		limit:    1000000,
		checksum: "c13929ee9d2aea8f83aa076236079e94",
		isShort:  true,
	},
	"10 Million": {
		limit:    10000000,
		checksum: "60e34d268bad671a5f299e1ecc988ff6",
		isShort:  true,
	},
	"100 Million": {
		limit:    100000000,
		checksum: "4e2b0027288a27e9c99699364877c9db",
		isShort:  true,
	},
	"1 Billion": {
		limit:    1000000000,
		checksum: "92c178cc5bb85e06366551c0ae7e18f6",
	},
	"10 Billion": {
		limit:    10000000000,
		checksum: "95ded65c9ddca18e1df966ba2a2b373a",
	},
}

func TestGenerate(t *testing.T) {
	short := testing.Short()
	if !short {
		fmt.Println("This test may take several minutes to complete. To skip the long test cases, use the -short flag.")
	}

	for name, tc := range checksumTT {
		if short && !tc.isShort {
			continue
		}

		t.Run(name, func(t *testing.T) {
			var i int
			size := 32
			digits := make([]byte, size)
			digits[size-1] = '\n'
			hasher := md5.New()

			primes := make(chan uint64, tc.limit/1000)
			go prime.Generate(primes, tc.limit)

			for number := range primes {
				for i = size - 2; i > 0; i-- {
					digits[i] = '0' + byte(number%10)
					if number < 10 {
						break
					}
					number /= 10
				}
				if _, err := hasher.Write(digits[i:]); err != nil {
					t.Errorf("ERROR: %+v", err)
				}
			}

			checksum := hex.EncodeToString(hasher.Sum(nil))

			if checksum != tc.checksum {
				t.Errorf("Expected: %s. Result: %s.", tc.checksum, checksum)
			}
		})
	}
}

var isPrimeTT = map[string]struct {
	candidate uint64
	isPrime   bool
}{
	"1": {
		candidate: 1,
		isPrime:   false,
	},
	"14": {
		candidate: 14,
		isPrime:   false,
	},
	"1,000th Prime": {
		candidate: 7919,
		isPrime:   true,
	},
	"1,000,000th Prime": {
		candidate: 15485863,
		isPrime:   true,
	},
}

func TestIsPrime(t *testing.T) {
	for name, tc := range isPrimeTT {
		t.Run(name, func(t *testing.T) {
			result := prime.IsPrime(tc.candidate)
			if tc.isPrime != result {
				t.Errorf("Expected: %t. Result: %t.", tc.isPrime, result)
			}
		})
	}
}
