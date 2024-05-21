package main

import (
	"fmt"
	//"math/rand"
	"strings"
	"time"
)

// just beautifying my code :P
func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	}
}

// prints item name, price, quantity and total price and sub total amount.
func generateBill() {

	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//prints the details of the food item that you've ordered.
	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//print sub total amount in a cool way!
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): ₹" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)

}

// prints the data of the items that you ordered.
func printOrderData() {
	for key := range customerOrder {
		//key -> it is the key values
		for _, element := range menu {
			if key == element.itemName {
				//customerOrder[key] -> will provide the no. of plates of that item
				//float64(customerOrder[key])*element.itemPrice -> this will result in the cost of each item
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}
