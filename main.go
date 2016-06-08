package main

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/NerdsvilleCEO/customer-ss/customer"
	. "github.com/NerdsvilleCEO/customer-ss/marketing_utils"
	. "github.com/NerdsvilleCEO/customer-ss/product"
	"os"
	"strconv"
)

var (
	customers *string = flag.String("c", "./customers", "The file holding the customers")
	products  *string = flag.String("p", "./products", "The file holding the products")
)

func init() {
	flag.Parse()
}

func main() {
	custScores := make(map[*Customer][]SustainabilityScore)
	custFile, _ := os.Open(*customers)
	defer custFile.Close()
	scanner := bufio.NewScanner(custFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		productsFile, _ := os.Open(*products)
		defer productsFile.Close()
		prodScan := bufio.NewScanner(productsFile)
		prodScan.Split(bufio.ScanLines)
		customer := &Customer{Name: scanner.Text()}
		for prodScan.Scan() {
			product := &Product{Name: prodScan.Text()}
			ss := SustainabilityScore{
				Customer: customer,
				Product:  product,
			}
			ss.Score = ss.Find()
			custScores[customer] = append(custScores[customer], ss)
		}
	}

	for customer := range custScores {
		var maxScore float64
		var winner SustainabilityScore
		for _, ss := range custScores[customer] {
			if ss.Score > maxScore {
				maxScore = ss.Score
				winner = ss
			}
		}
		fmt.Println(customer.Name + " " + winner.Product.Name + " matched with " + strconv.FormatFloat(winner.Score, 'f', -1, 64) + " sustainability score")
	}
}
