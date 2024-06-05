package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name  string
	Brand string
	Type  string
	Price float64
	Stock int
}

type Inventory struct {
	Products []Product
}

type Transaction struct {
	ProductName string
	Quantity    int
}

func (inv *Inventory) AddProduct(p Product) {
	inv.Products = append(inv.Products, p)
}

func (inv *Inventory) FindProduct(name string) *Product {
	for i, p := range inv.Products {
		if p.Name == name {
			return &inv.Products[i]
		}
	}
	return nil
}

func (inv *Inventory) UpdateProduct(name string, newProduct Product) bool {
	for i, p := range inv.Products {
		if p.Name == name {
			inv.Products[i] = newProduct
			return true
		}
	}
	return false
}

func (inv *Inventory) DeleteProduct(name string) bool {
	for i, p := range inv.Products {
		if p.Name == name {
			inv.Products = append(inv.Products[:i], inv.Products[i+1:]...)
			return true
		}
	}
	return false
}

func (inv *Inventory) SortProductsBy(criteria string) {
	BubbleSort(inv.Products, criteria)
}

func BubbleSort(products []Product, criteria string) {
	n := len(products)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			swap := false
			switch criteria {
			case "price":
				if products[j].Price > products[j+1].Price {
					swap = true
				}
			case "name":
				if products[j].Name > products[j+1].Name {
					swap = true
				}
			case "brand":
				if products[j].Brand > products[j+1].Brand {
					swap = true
				}
			}
			if swap {
				products[j], products[j+1] = products[j+1], products[j]
			}
		}
	}
}

func (inv *Inventory) DisplayProducts() {
	headers := []string{"Name", "Brand", "Type", "Price", "Stock"}
	rows := [][]string{}

	for _, p := range inv.Products {
		row := []string{p.Name, p.Brand, p.Type, fmt.Sprintf("%.2f", p.Price), fmt.Sprintf("%d", p.Stock)}
		rows = append(rows, row)
	}

	printTable(headers, rows)
}

func (inv *Inventory) RecordTransaction(t Transaction) bool {
	product := inv.FindProduct(t.ProductName)
	if product == nil {
		fmt.Println("Product not found")
		return false
	}
	if product.Stock < t.Quantity {
		fmt.Println("Insufficient stock")
		return false
	}
	product.Stock -= t.Quantity
	return true
}

func printTable(headers []string, rows [][]string) {
	// Calculate column widths
	columnWidths := make([]int, len(headers))
	for i, header := range headers {
		columnWidths[i] = len(header)
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > columnWidths[i] {
				columnWidths[i] = len(cell)
			}
		}
	}

	// Print headers with lines
	printLine(columnWidths, '┌', '┬', '┐')
	printRow(headers, columnWidths)
	printLine(columnWidths, '├', '┼', '┤')

	// Print rows
	for _, row := range rows {
		printRow(row, columnWidths)
	}
	printLine(columnWidths, '└', '┴', '┘')
}

func printRow(row []string, columnWidths []int) {
	for i, cell := range row {
		fmt.Printf("│ %-*s ", columnWidths[i], cell)
	}
	fmt.Println("│")
}

func printLine(columnWidths []int, left, mid, right rune) {
	fmt.Printf("%c", left)
	for i, width := range columnWidths {
		fmt.Printf("%s", strings.Repeat("─", width+2))
		if i < len(columnWidths)-1 {
			fmt.Printf("%c", mid)
		}
	}
	fmt.Printf("%c\n", right)
}

func main() {
	inv := Inventory{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nCommands: add, find, update, delete, sort, transaction, display, exit")
		fmt.Print("Enter command: ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			var name, brand, typeProduct string
			var price float64
			var stock int

			fmt.Print("Enter product name: ")
			scanner.Scan()
			name = scanner.Text()

			fmt.Print("Enter product brand: ")
			scanner.Scan()
			brand = scanner.Text()

			fmt.Print("Enter product type: ")
			scanner.Scan()
			typeProduct = scanner.Text()

			fmt.Print("Enter product price: ")
			scanner.Scan()
			price, _ = strconv.ParseFloat(scanner.Text(), 64)

			fmt.Print("Enter product stock: ")
			scanner.Scan()
			stock, _ = strconv.Atoi(scanner.Text())

			inv.AddProduct(Product{Name: name, Brand: brand, Type: typeProduct, Price: price, Stock: stock})
			fmt.Println("Product added successfully.")

		case "find":
			fmt.Print("Enter product name to find: ")
			scanner.Scan()
			name := scanner.Text()

			if p := inv.FindProduct(name); p != nil {
				fmt.Printf("Product found: %v\n", *p)
			} else {
				fmt.Println("Product not found.")
			}

		case "update":
			var name, newName, newBrand, newType string
			var newPrice float64
			var newStock int

			fmt.Print("Enter product name to update: ")
			scanner.Scan()
			name = scanner.Text()

			fmt.Print("Enter new product name: ")
			scanner.Scan()
			newName = scanner.Text()

			fmt.Print("Enter new product brand: ")
			scanner.Scan()
			newBrand = scanner.Text()

			fmt.Print("Enter new product type: ")
			scanner.Scan()
			newType = scanner.Text()

			fmt.Print("Enter new product price: ")
			scanner.Scan()
			newPrice, _ = strconv.ParseFloat(scanner.Text(), 64)

			fmt.Print("Enter new product stock: ")
			scanner.Scan()
			newStock, _ = strconv.Atoi(scanner.Text())

			if inv.UpdateProduct(name, Product{Name: newName, Brand: newBrand, Type: newType, Price: newPrice, Stock: newStock}) {
				fmt.Println("Product updated successfully.")
			} else {
				fmt.Println("Product not found.")
			}

		case "delete":
			fmt.Print("Enter product name to delete: ")
			scanner.Scan()
			name := scanner.Text()

			if inv.DeleteProduct(name) {
				fmt.Println("Product deleted successfully.")
			} else {
				fmt.Println("Product not found.")
			}

		case "sort":
			fmt.Print("Enter criteria to sort by (price, name, brand): ")
			scanner.Scan()
			criteria := scanner.Text()

			inv.SortProductsBy(criteria)
			fmt.Println("Products sorted successfully.")

		case "transaction":
			fmt.Print("Enter product name for transaction: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter quantity: ")
			scanner.Scan()
			quantity, _ := strconv.Atoi(scanner.Text())

			if inv.RecordTransaction(Transaction{ProductName: name, Quantity: quantity}) {
				fmt.Println("Transaction recorded successfully.")
			} else {
				fmt.Println("Transaction failed.")
			}

		case "display":
			inv.DisplayProducts()

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid command.")
		}
	}
}
