package main

import (
  "time"
)

func timeIt() {
  start := time.Now()
  primes(10000)
  elapsed := time.Since(start)

  startOther := time.Now()
  primesOld(10000)
  elapsedOther := time.Since(startOther)

  println("efficient method: ", elapsed/1000/1000, "ms")
  println("other method: ", elapsedOther/1000/1000, "ms")
}
