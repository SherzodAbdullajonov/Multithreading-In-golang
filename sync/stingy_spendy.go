package main

import "time"

var money = 100

func stringy() {
	for i := 1; i <= 1000; i++ {
		money += 10
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}
func spendy() {
	for i := 1; i < 1000; i++ {
		money -= 10
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}
func main() {
	go stringy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
