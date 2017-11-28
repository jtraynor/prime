# Prime

Generates prime numbers using the [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes) algorithm.
The prime numbers are streamed back to the caller via a channel.

## Performance

While this generator definitely isn't the fastest available, it is very memory efficient. It achieves this by utilising an
array of uint64 [bit fields](https://en.wikipedia.org/wiki/Bit_field) rather than the traditional array of booleans. The
benefit of this is that the number of primes that can be generated is impacted much less by memory constraints.

	--- PASS: TestGenerate
		--- PASS: TestGenerate/1_Million (0.01s)
		--- PASS: TestGenerate/10_Million (0.13s)
		--- PASS: TestGenerate/100_Million (1.40s)
		--- PASS: TestGenerate/1_Billion (17.92s)
		--- PASS: TestGenerate/10_Billion (190.90s)

Run on an Intel Xeon E3-1241 CPU (3.5GHz, 8 logical cores).

## Example

The following code will calculate the sum of all of the prime numbers between 1 and 1,000,000.

	primes := make(chan uint64)
	go prime.Generate(primes, 1000000)

	sum := uint64(0)
	for number := range primes {
		sum += number
	}
