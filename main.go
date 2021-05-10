package main

/* Imporing the libraries */
import (
	"bufio"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/* Declaring global variables */
var vegPrice = []int{80, 120, 100, 130}
var fruitPrice = []int{140, 120, 90, 180}
var fruits = []string{"apple", "orange", "banana", "strawberry"}
var vegetables = []string{"tomato", "carrot", "cucumber", "cauliflower"}

var payment int
var product string
var subtractedAmount int

/* Showing the list of items */
func listItems() {
	fmt.Println("Here is the list of items")
	for i := 0; i < 4; i++ {
		fmt.Println("Vegetables: " + vegetables[i] + " " + strconv.Itoa(vegPrice[i]))
	}
	for i := 0; i < 4; i++ {
		fmt.Println("Fruits: " + fruits[i] + " " + strconv.Itoa(fruitPrice[i]))
	}
}

/* Accepting string with spaces */
var combined = make(map[string]string)

func oneLine(key string) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result := scanner.Text()
		combined[key] = result + ", "
	}
}

/* Adding customer */
func addCustomer() {
	fmt.Println("-------------------------")
	fmt.Println("Enter your details to buy")
	fmt.Println("-------------------------")

	fmt.Printf("Enter your name: ")
	oneLine("name")

	fmt.Printf("Enter your address: ")
	oneLine("address")

	fmt.Printf("Enter your phone number: ")
	oneLine("phone_number")

	fmt.Printf("Enter your product: ")
	fmt.Scanln(&product)

	for i := 0; i < 4; i++ {
		if product == vegetables[i] {
			oneLine(vegetables[i])
			break
		} else if product == fruits[i] {
			oneLine(fruits[i])
			break
		}
	}

	fmt.Printf("Enter your payment: ")
	fmt.Scanln(&payment)

	if payment <= 0 {
		fmt.Println("Incorrect Payment")
		os.Exit(3)
	} else {
		for i := 0; i < 4; i++ {
			if product == vegetables[i] {
				subtractedAmount = payment - vegPrice[i]
				fmt.Println("Take your arrears: " + strconv.Itoa(subtractedAmount))
				fmt.Println("Thank you!")
			} else if product == fruits[i] {
				subtractedAmount = payment - fruitPrice[i]
				fmt.Println("Take your arrears: " + strconv.Itoa(subtractedAmount))
				fmt.Println("Thank you!")
			}
		}
	}
}

/* Writing customer data into file */
func writeFile() {
	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	presentTime := time.Now()
	strArray := strings.Fields(presentTime.String())
	oneArray := []string{combined["name"], combined["address"], combined["phone_number"], product + ", ", strconv.Itoa(payment) + ", ", strconv.Itoa(subtractedAmount) + ", ", strArray[0] + ", ", strArray[1] + "\n"}

	for i := 0; i < 8; i++ {
		if _, err := file.Write([]byte(oneArray[i])); err != nil {
			log.Fatal(err)
		}
	}

	if file.Close(); err != nil {
		log.Fatal(err)
	}
}

/* Accessing file data */
func showData() {
	var searchData string
	fmt.Printf("Enter your name or phone number: ")
	fmt.Scanln(&searchData)

	content, err := ioutil.ReadFile("data.csv")

	if err != nil {
		log.Fatal(err)
	}

	suffix := suffixarray.New(content)
	indexList := suffix.Lookup([]byte(searchData), -1)

	if len(indexList) == 0 {
		fmt.Println("Data is not found")
		return
	}

	data := string(content)
	for _, idx := range indexList {
		start := idx
		for start >= 0 && data[start] != '\n' {
			start--
		}
		end := idx
		for end < len(data) && data[end] != '\n' {
			end++
		}
		fmt.Println(string(data[start+1 : end]))
	}

}

/* Calling all the functions */
func main() {
	var choice string
	fmt.Println("------------------------")
	fmt.Println("Welcome to the XYZ shop")
	fmt.Println("------------------------")

	fmt.Println("1. Start shopping")
	fmt.Println("2. Check the shopping record")
	fmt.Println("3. Exit")

	fmt.Printf("Enter your choice: ")
	fmt.Scanln(&choice)

	if choice == "1" {
		listItems()
		addCustomer()
		writeFile()
	} else if choice == "2" {
		showData()
	} else if choice == "3" {
		os.Exit(1)
	} else {
		fmt.Println("Enter the correct choice")
		os.Exit(2)
	}
}
