package main

func primes(n int) {
  var i int
  for i <= n {
    if(isPrime(i)) {
      // println(i)
    }
    i++
  }
}
func primesOld(n int) {
  var i int
  for i <= n {
    if(isPrimeOld(i)) {
      // println(i)
    }
    i++
  }
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
  if(num > 6 && !((num-1) % 6 == 0 || (num+1) % 6 == 0)) {
    return false
  }
	i := 2
	for i < num+1 {
		if num != i && num%i == 0 {
			return false
		}
		i++
	}
	return true

}

func isPrimeOld(num int) bool {
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
	return true

}
