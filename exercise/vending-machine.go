package main

import "fmt"

type VendingMachine struct {
	coins []string
	items string
	money int
}

func (vm VendingMachine) InsertCoin(coin string) {
	vm.coins = append(vm.coins, coin)
	if coin=="T" {
		vm.money = vm.money+10
	} else if coin=="F" {
		vm.money = vm.money+5
	} else if coin=="TW" {
		vm.money = vm.money+2
	} else if coin=="O" {
		vm.money = vm.money+1
	}
}

func (vm VendingMachine) GetInsertedMoney() int {
	return vm.money
}

func (vm VendingMachine) SelectSD() string {
	vm.items = "SD"
	return vm.items
}

func NewVendingMachine(coins []string, items string, money int) *VendingMachine {
	return &VendingMachine{coins, items, money}
}

func main() {

	money := 0
	items := ""
	coins := make([]string, 5)
	vm := NewVendingMachine(coins, items, money)

	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney())
	can := vm.SelectSD()
	fmt.Println(can) // SD
	/*
	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	can = vm.SelectCC()
	fmt.Println(can)

	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	change := vm.CoinReturn()
	fmt.Println(change)
	*/
}


