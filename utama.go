package main

import (
	"fmt"
	"math"
)

const Nmax = 100000

var count = 0

type borrower struct {
	name           string
	loanAmount     float64
	loanTerm       int
	interestRate   float64
	monthlyPayment float64
}

type Borrow [Nmax]borrower

func main() {
	var B Borrow
	var choice int
	menu()
	fmt.Scan(&choice)
	for choice != 0 {
		switch choice {
		case 1:
			addBorrower(&B, &count)
			menu()
			fmt.Scan(&choice)
		case 2:
			removeBorrower(&B, &count)
			menu()
			fmt.Scan(&choice)
		case 3:
			editBorrower(&B, &count)
			menu()
			fmt.Scan(&choice)
		case 4:
			searchBorrower(B, count)
			menu()
			fmt.Scan(&choice)
		case 5:
			report(B, count)
			menu()
			fmt.Scan(&choice)
		case 6:
			sortBorrowers(&B, count)
			menu()
			fmt.Scan(&choice)
		default:
			fmt.Println("Invalid option.")
			menu()
			fmt.Scan(&choice)
		}
	}
}

func menu() {
	fmt.Println("___________________________________________")
	fmt.Println("1.Add Borrower")
	fmt.Println("2.Remove Borrower")
	fmt.Println("3.Edit Borrower")
	fmt.Println("4.Search Borrower")
	fmt.Println("5.All Borrower Report")
	fmt.Println("6.Sort Borrowers")
	fmt.Println("0.Exit")
	fmt.Println("___________________________________________")
	fmt.Print("Choose: ")
}

func addBorrower(B *Borrow, N *int) {
	var name string
	var amount float64
	var term int
	var rate float64

	fmt.Print("Name: ")
	fmt.Scan(&name)
	B[*N].name = name

	fmt.Print("Loan amount: ")
	fmt.Scan(&amount)
	B[*N].loanAmount = amount

	fmt.Print("Loan term (months): ")
	fmt.Scan(&term)
	B[*N].loanTerm = term

	fmt.Print("Annual interest rate (%): ")
	fmt.Scan(&rate)
	B[*N].interestRate = rate

	B[*N].monthlyPayment = CalculatemonthlyPayment(amount, rate, term)

	*N++
	fmt.Println("Borrower added successfully")
}

func CalculatemonthlyPayment(amount, rate float64, term int) float64 {
	if rate == 0 {
		return amount / float64(term)
	}
	r := rate / 12 / 100
	return amount * r * math.Pow(1+r, float64(term)) / (math.Pow(1+r, float64(term)) - 1)
}

func removeBorrower(B *Borrow, N *int) {
	var name string
	fmt.Print("Enter name to remove: ")
	fmt.Scan(&name)

	found := false
	for i := 0; i < *N && !found; i++ {
		if B[i].name == name {
			for j := i; j < *N-1; j++ {
				B[j] = B[j+1]
			}
			*N--
			fmt.Println("Borrower removed.")
			found = true
		}
	}
	if !found {
		fmt.Println("Borrower not found.")
	}
}

func editBorrower(B *Borrow, N *int) {
	var name string
	fmt.Print("Enter borrower name to edit: ")
	fmt.Scan(&name)

	found := false
	for i := 0; i < *N; i++ {
		if B[i].name == name {
			fmt.Print("New Loan Amount: ")
			fmt.Scan(&B[i].loanAmount)

			fmt.Print("New Loan Term (months): ")
			fmt.Scan(&B[i].loanTerm)

			fmt.Print("New Interest Rate (%): ")
			fmt.Scan(&B[i].interestRate)

			B[i].monthlyPayment = CalculatemonthlyPayment(B[i].loanAmount, B[i].interestRate, B[i].loanTerm)
			fmt.Println("Borrower updated.")
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Borrower not found.")
	}
}

func report(B Borrow, N int) {
	fmt.Println("All Borrowers:")
	for i := 0; i < N; i++ {
		fmt.Printf("Borrower %d: %+v\n", i+1, B[i])
	}
}

func sortBorrowers(B *Borrow, n int) {
	var option int
	fmt.Println("Sort Options:")
	fmt.Println("1. Name (A-Z)")
	fmt.Println("2. Name (Z-A)")
	fmt.Println("3. Loan Amount (High to Low)")
	fmt.Println("4. Interest Rate (High to Low)")
	fmt.Print("Choose sort option: ")
	fmt.Scan(&option)

	var pass, idx, i int
	var temp borrower

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			switch option {
			case 1:
				if B[i].name < B[idx].name {
					idx = i
				}
			case 2:
				if B[i].name > B[idx].name {
					idx = i
				}
			case 3:
				if B[i].loanAmount > B[idx].loanAmount {
					idx = i
				}
			case 4:
				if B[i].interestRate > B[idx].interestRate {
					idx = i
				}
			default:
				fmt.Println("Invalid option.")
				return
			}
		}
		temp = B[pass-1]
		B[pass-1] = B[idx]
		B[idx] = temp
	}

	fmt.Printf("%10s \n", "Sorted Borrowers")
	fmt.Printf("%10s %10s %10s %10s %10s \n", "Name", "LoanAmount", "LoanTerm", "InterestRate", "MonthlyPayment")
	for i = 0; i < n; i++ {
		fmt.Printf("%10s %10.2f %10d %10.2f %10.2f\n", B[i].name, B[i].loanAmount, B[i].loanTerm, B[i].interestRate, B[i].monthlyPayment)
	}
	fmt.Println()
}

func printPaymentSchedule(b borrower) {
	fmt.Printf("Payment Schedule for %s\n", b.name)
	fmt.Printf("Loan Amount: %.2f, Term: %d months, Interest Rate: %.2f%%\n", b.loanAmount, b.loanTerm, b.interestRate)
	fmt.Println("Month | Payment | Interest | Principal | Remaining Balance")

	monthlyRate := b.interestRate / 12 / 100
	remaining := b.loanAmount

	for month := 1; month <= b.loanTerm; month++ {
		interest := remaining * monthlyRate
		principal := b.monthlyPayment - interest
		remaining -= principal
		if remaining < 0 {
			remaining = 0
		}
		fmt.Printf("%5d | %7.2f | %8.2f | %9.2f | %17.2f\n", month, b.monthlyPayment, interest, principal, remaining)
	}
	fmt.Println()
}

func sortForBinarySearch(B *Borrow, N int) {
	var idx int
	var temp borrower

	for i := 0; i < N-1; i++ {
		idx = i
		for j := i + 1; j < N; j++ {
			if B[j].name < B[idx].name {
				idx = j
			}
		}
		temp = B[i]
		B[i] = B[idx]
		B[idx] = temp
	}
}

func searchBorrower(B Borrow, N int) {
	var method int
	var name string
	fmt.Println("1.Binary Search")
	fmt.Println("2.Sequential Search")
	fmt.Print("Choose: ")
	fmt.Scan(&method)
	if method == 1 {
		fmt.Print("Enter borrower name to search: ")
		fmt.Scan(&name)
		binarySearch(B, N, name)
	} else if method == 2 {
		fmt.Print("Enter borrower name to search: ")
		fmt.Scan(&name)
		sequentialSearch(B, N, name)
	} else {
		fmt.Println("Invalid choice")
	}
}
func binarySearch(B Borrow, N int, name string) {
	var left, right, mid int
	var found bool
	sortForBinarySearch(&B, N)
	left = 0
	right = N - 1
	found = false
	for left <= right {
		mid = (left + right) / 2
		if B[mid].name == name {
			found = true
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("Borrower Found")
			fmt.Println("Name: ", B[mid].name)
			fmt.Println("Loan Amount: ", B[mid].loanAmount)
			fmt.Println("Loan Term: ", B[mid].loanTerm)
			fmt.Println("Interest Rate: ", B[mid].interestRate)
			fmt.Println("Monthly Payment: ", B[mid].monthlyPayment)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			printPaymentSchedule(B[mid])
			left = right + 1
		} else if B[mid].name < name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("Borrower not found.")
	}
}
func sequentialSearch(B Borrow, N int, name string) {
	var found bool
	for i := 0; i < N; i++ {
		if B[i].name == name {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println("Borrower Found")
			fmt.Println("Name: ", B[i].name)
			fmt.Println("Loan Amount: ", B[i].loanAmount)
			fmt.Println("Loan Term: ", B[i].loanTerm)
			fmt.Println("Interest Rate: ", B[i].interestRate)
			fmt.Println("Monthly Payment: ", B[i].monthlyPayment)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			found = true
			printPaymentSchedule(B[i]) // ここで支払スケジュール表示を追加
		}
	}
	if !found {
		fmt.Println("Borrower not found.")
	}
}
