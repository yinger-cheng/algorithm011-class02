package main

import "fmt"

func main(){
	x := 2.00000
	n := 10
	fmt.Println(myPow(x,n))
}

func myPow(x float64, n int) float64 {
	var result float64
	if n < 0 {
		x = 1 / x
		result = - result
	}
	return myPowHalf(x , n)
}

func myPowHalf(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	result := myPowHalf(x, n/2 )
	if n % 2 == 0 {
		result = result * result
	}else{
		result = result * result * x
	}
	return result
}