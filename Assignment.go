package main

import (
 "fmt"
 "os"
 "strings"
)

func main() {
 var or string
 var q int
 var sales = []int{}
 counter := 0
 for true {
  fmt.Println("\t\t\t\t\t______________MENU______________")
  fmt.Print("\t\t\t\t\t C: coffee \t\t40rs\n\t\t\t\t\t D: dosa   \t\t80rs\n\t\t\t\t\t T: tomato soup \t20rs\n\t\t\t\t\t I: idli   \t\t20rs\n\t\t\t\t\t V: vada \t\t25rs \n\t\t\t\t\t B: bhature &chhole \t30rs\n\t\t\t\t\t P: paneer pakoda   \t30rs\n\t\t\t\t\t M: manchurian  \t90rs\n\t\t\t\t\t H: hakka noodle   \t70rs \n\t\t\t\t\t F: French fries \t30rs\n\t\t\t\t\t J: jalebi \t\t30rs\n\t\t\t\t\t L: lemonade \t\t15rs\n\t\t\t\t\t S: spring roll \t20rs\n")
  fmt.Println("_____________________________________________________________")
  fmt.Println("Press END For closing the day.")
  fmt.Println("_____________________________________________________________")
  fmt.Println("Please place the order type: ")
  fmt.Scan(&or)
  or = strings.ToUpper(or)

  if or == "END" {
   total_income(sales)
  } else {
      fmt.Println("Please place the order Quantity: ")
      fmt.Scan(&q)
  }
  bill := bill(q, or)
  fmt.Println("Total bill: ", bill)
  sales = append(sales, bill)
  counter++

 }
}

func total_income(sale []int) {
 sum :=0

 for j := 0; j < len(sale); j++ {
  sum = sale[j] + sum
 }
 fmt.Println("Total Income for the day is : ", sum)
 os.Exit(0)
}

func bill(q int, or string) int {
 m := map[string]int{
  "C": 40,
  "D": 80,
  "T": 20,
  "I": 20,
  "V": 25,
  "B": 30,
  "P": 30,
  "M": 90,
  "H": 70,
  "F": 30,
  "J": 30,
  "L": 15,
  "S": 20,
 }
 bills := q * m[or]
 return bills
}//or=order
//q=quantity