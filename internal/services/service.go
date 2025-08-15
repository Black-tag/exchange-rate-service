package services

import (
	"fmt"
	"strings"

	
)


func ConvertCurrency(From, To string, exchange map[string]float64 ) (amount float64) {
	fmt.Print(exchange)
	from := strings.ToUpper(From)
	fmt.Printf("from: %s", from)
	to := strings.ToUpper(To)
	fmt.Printf("from: %s", to)
	// exchangeStr := from + to
	if from == to {
		fmt.Printf("form is equal to to")
		return 1
	}
	if from == "USD" { 
		key := "USD" + to
		fmt.Printf("key: %s, rate: %f\n", key, exchange[key])
		return exchange["USD"+to]
	}
	if to == "USD" {
		key := "USD" + from
		rate := exchange[key]
		fmt.Printf("key: %s, rate: %f\n", key, rate)
		if rate == 0 {
			fmt.Printf("rate: %f", rate)
			return 0
		}

		return 1 / rate

	}
	convertFrom := exchange["USD"+from]
	fmt.Printf("convertfrom: %f", convertFrom)
	convertTo := exchange["USD"+to]
	fmt.Printf("convertTo: %f", convertTo)
	if convertFrom == 0 {
		return 0
	}
	return convertTo / convertFrom
	
}	


// 		if convertFrom == 0 {
//     		return 0
// 		}
// 		exhangeRate := convertTo / convertFrom
// 		return exhangeRate