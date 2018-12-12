package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name")
	text, _ := reader.ReadString('\n')
	fmt.Println("text", text)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name")
	scanner.Scan()
	text = scanner.Text()
	fmt.Println("text", text)

	fmt.Println("Enter Text")
	text2 := ""
	fmt.Scan(&text2)
	fmt.Println(text2)
	fmt.Scanln(&text2)
	fmt.Println(text2)

	fmt.Println("Enter value for a and b")
	var number1 int
	number2 := 0
	fmt.Scanf("%d \n %d", &number1 , &number2) //c printf......& is compulsory
	fmt.Println("Num1" , number1 , "NUm2" , number2)
	//for scanning infinite times
	//fmt.Println("Enter your name")
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}

}
