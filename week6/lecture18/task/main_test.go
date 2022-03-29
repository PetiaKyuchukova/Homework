package primes

import "testing"

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		primesAndSleep(100, 0)
	}
}
func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		primesAndSleep(100, 5)
	}
}
func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		primesAndSleep(100, 10)
	}
}
func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goPrimesAndSleep(100, 0)
	}
}
func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goPrimesAndSleep(100, 5)
	}
}
func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		goPrimesAndSleep(100, 10)
	}
}
