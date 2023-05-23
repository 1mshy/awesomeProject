package main

func primes() {

}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	i := 2
	for i < num+1 {
		if num != i && num%i == 0 {
			return false
		}
		i++
	}
	println(num)
	return true

}
