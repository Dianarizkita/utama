package main

import "fmt"

func primeNumber(bilangan int) bool {
	BilPrima := true
	if bilangan == 0 || bilangan == 1 {
		fmt.Printf("%d Bukan Bilangan Prima \n", bilangan)
	} else {
		for i := 2; i <= bilangan/2; i++ {
			if bilangan%i == 0 {
				fmt.Printf(" %d Bukan Bilangan Prima \n", bilangan)
				BilPrima = false
				break
			}
		}
		if BilPrima == true {
			fmt.Printf(" %d Bilangan Prima \n", bilangan)
		}
	}
	return BilPrima
}
func main() {
	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
