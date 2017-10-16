package main

import "fmt"

func main() {
	length := 199999

	// Primes.
	primes := make([]int, length)
	for ind := range primes {
		primes[ind] = ind + 2
	}

	// Euclidean seive.
	for ind, val := range primes {
		if val != 1 {
			for j := ind + val; j < length; j += val {
				if primes[j] != 1 {
					primes[j] = 1
				}
			}
		}
	}

	tri := 1
	divisors := 1

	for n := 2; divisors < 501; n++ {
		tri += n

		temp := tri
		factors := make([]int, 0)

		// Make factors.
		for i := 0; i < length; i++ {
			if primes[i] != 1 && temp%primes[i] == 0 {
				factors = append(factors, primes[i])
				temp /= primes[i]
				i--
				if temp == 1 {
					break
				}
			}
		}
		if temp != 1 {
			fmt.Println("Not enough primes!")
			break
		}

		power := 0
		check := factors[0]
		divisors = 1

		for _, val := range factors {
			if val == check {
				power++
			} else {
				divisors *= power + 1
				power = 1
				check = val
			}
		}
		divisors *= power + 1
	}

	fmt.Println("Tri: ", tri)
}