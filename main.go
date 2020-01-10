package main

import "fmt"

var (
	Red    = Color("\033[1;31m%s\033[0m")
	Yellow = Color("\033[1;33m%s\033[0m")
	Teal   = Color("\033[1;36m%s\033[0m")
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

// Color func
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// ChallengeWeek struct
type ChallengeWeek struct {
	Amount int
	Total  int
	Weeks  int
}

// NewChallengeWeek func
func NewChallengeWeek() *ChallengeWeek {
	return &ChallengeWeek{
		Weeks: 53,
	}
}

func (cw *ChallengeWeek) getTotal() int {
	for week := 1; week < cw.Weeks; week++ {
		amountPerWeek := week * cw.Amount
		cw.Total = (cw.Total + amountPerWeek)
		fmt.Printf(Info("Week %d: $%d \n"), week, amountPerWeek)
	}

	return cw.Total
}

func main() {
	cw := NewChallengeWeek()

	fmt.Println(Fata("How much do you want to save?"))
	fmt.Scanf("%d", &cw.Amount)

	total := cw.getTotal()
	fmt.Printf(Warn("Total: $%d \n"), total)
}
