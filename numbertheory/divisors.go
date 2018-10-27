package numbertheory

import (
	"sort"
)

// GetNumberOfDivisors returns the number of divisors of number.
func GetNumberOfDivisors(number uint) int {

	if number <= 1 {
		panic("The number must be larger than 1")
	}

	primeFactorChannel := GetPrimeFactorisation(number)

	// Calculate the number of divisors.
	power := 0
	var previousPrime uint
	divisors := 1
	for val := range primeFactorChannel {
		if val == previousPrime {
			power++
		} else {
			divisors *= power + 1 // The first time this is OK because power = 0.
			power = 1
			previousPrime = val
		}
	}

	divisors *= power + 1

	return divisors
}

// GetDivisorsOfANumber fills a channel with all the divisors of a number.
// Syncing and safely exiting this function can be done through flushing the divisor channel.
func GetDivisorsOfANumber(number uint) chan uint {

	if number < 1 {
		panic("The number must be larger than 0")
	}

	divisorChannel := make(chan uint, 100)
	go func() {
		divisorChannel <- 1
		if number != 1 {
			primeFactorChannel := GetPrimeFactorisation(number)

			// A record of all the divisors.
			allDivisors := []uint{1}

			// A record of the new divisors to be added. This slice is new generated for each distinct prime factor.
			newDivisors := make([]uint, 0)

			// The previous prime factor in the loop to check whether a new distinct prime factor has come through.
			var previousPrime uint

			// The prime factor raised to each power from 1 to the number of times it appears in the factorisation
			// which is used to calculate the new divisors.
			var currentPrimeMultiple uint

			for currentPrime := range primeFactorChannel {
				if currentPrime == previousPrime {
					currentPrimeMultiple *= currentPrime
				} else {
					// A new distinct factor is found. Add all the divisors generated from the previous factor and generate the new ones.
					allDivisors = append(allDivisors, newDivisors...)
					newDivisors = []uint{}
					previousPrime = currentPrime
					currentPrimeMultiple = currentPrime
				}
				for _, divisor := range allDivisors {
					divisor := currentPrimeMultiple * divisor
					divisorChannel <- divisor
					newDivisors = append(newDivisors, divisor)
				}
			}
		}

		close(divisorChannel)
	}()

	return divisorChannel
}

// GetDivisorsOfANumberInASlice returns a sorted slice with all the divisors of a number.
func GetDivisorsOfANumberInASlice(number uint) []uint {

	if number < 1 {
		panic("The number must be larger than 0")
	}
	// A record of all the divisors.
	allDivisors := []uint{1}

	if number != 1 {
		primeFactorChannel := GetPrimeFactorisation(number)

		// A record of the new divisors to be added. This slice is new generated for each distinct prime factor.
		newDivisors := make([]uint, 0)

		// The previous prime factor in the loop to check whether a new distinct prime factor has come through.
		var previousPrime uint

		// The prime factor raised to each power from 1 to the number of times it appears in the factorisation
		// which is used to calculate the new divisors.
		var currentPrimeMultiple uint

		for currentPrime := range primeFactorChannel {
			if currentPrime == previousPrime {
				currentPrimeMultiple *= currentPrime
			} else {
				// A new distinct factor is found. Add all the divisors generated from the previous factor and generate the new ones.
				allDivisors = append(allDivisors, newDivisors...)
				newDivisors = []uint{}
				previousPrime = currentPrime
				currentPrimeMultiple = currentPrime
			}
			for _, divisor := range allDivisors {
				newDivisors = append(newDivisors, currentPrimeMultiple*divisor)
			}
		}

		allDivisors = append(allDivisors, newDivisors...)
		sort.Slice(allDivisors, func(i, j int) bool { return allDivisors[i] < allDivisors[j] })
	}

	return allDivisors
}