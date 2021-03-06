package primenumbers

// CalculatePrimes will calculate prime numbers between start and stop.
func CalculatePrimes(start uint64, stop uint64) []uint64 {
	primes := make([]uint64, 0)

	for number := start; number <= stop; number++ {
		if isPrimeToTen(number) {
			primes = append(primes, number)
		} else if number == 1 {
			continue
		}

		isPrime := true

		for _, divisor := range [4]uint64{2, 3, 5, 7} {
			if number%divisor == 0 {
				isPrime = false

				break
			}
		}

		if isPrime {
			primes = append(primes, number)
		}
	}

	return primes
}

func isPrimeToTen(prime uint64) bool {
	return prime == 2 || prime == 3 || prime == 5 || prime == 7
}
