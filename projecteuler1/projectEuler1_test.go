package projecteuler1

import "testing"

func TestProjectEulerChallenges(t *testing.T) {
	testCases := []struct {
		desc           string
		function       func() int
		expectedResult int
	}{
		{"1a: Multiples of 3 and 5 (Method 1)", MultiplesOf3And5MethodOne, 233168},
		{"1b: Multiples of 3 and 5 (Method 2)", MultiplesOf3And5MethodTwo, 233168},
		{"2: Even Fibonacci Numbers", EvenFibonacciNumbers, 4613732},
		{"3: Largest Prime Factor", LargestPrimeFactor, 6857},
		{"4: Largest Palindrome Product", LargestPalindromeProduct, 906609},
		{"5: Smallest Multiple", SmallestMultiple, 232792560},
		{"6: Sum Square Difference", SumSquareDifference, 25164150},
		{"7: 10001st Prime Number", Position10001Prime, 104743},
		{"8: Largest Product in a Series", LargestProductInASeries, 23514624000},
		{"9: Special Pythagorean triplet", SpecialPythagoreanTriplet, 31875000},
		{"10: Summation of Primes", SummationOfPrimes, 142913828922},
		{"11: Largest Product in a Grid", LargestProductInAGrid, 70600674},
		{"12: Highly Divisible Triangular Number", HighlyDivisibleTriangularNumber, 76576500},
		{"13: Large Sum", LargeSum, 5537376230},
		{"14: Longest Collatz Sequence", LongestCollatzSequence, 837799},
		{"15: Lattice Paths", LatticePaths, 137846528820},
		{"16: Power Digit Sum", PowerDigitSum, 1366},
		{"17: Number Letter Counts", NumberLetterCounts, 21124},
		{"18: Maximum Path Sum One", MaximumPathSumOne, 1074},
		{"19: Counting Sundays", CountingSundays, 171},
		{"20: Factorial Digit Sum", FactorialDigitSum, 648},
	}
	for _, tC := range testCases {
		if actualResult := tC.function(); tC.expectedResult != actualResult {
			t.Errorf("Failure. Test case: %v. Expected result: %v. Actual result: %v", tC.desc, tC.expectedResult, actualResult)
		}
	}
}