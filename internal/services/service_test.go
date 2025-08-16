package services



import "testing"


func TestConvertCurrency(t *testing.T) {
	expecteRates := map[string]float64{
		"USDINR": 83.0,
		"USDEUR": 1.1,
	}
	tests := []struct {
		name string
		from string
		to   string
		rates map[string]float64
		expect float64
	}{
		{"conversion from USD-INR, Valid", "USD", "INR", expecteRates, 83.0},
		{"conversion from EUR-USD, valid", "EUR", "USD", expecteRates, 1/1.1},
		{"conversion from ZMW-GBP, invalid", "ZMW", "GBP", expecteRates, 0},

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ConvertCurrency(test.from, test.to, test.rates)
			if got != test.expect {
				t.Errorf("ConvertCurrency(%s,%s) = %v; expect %v", test.from, test.to, got, test.expect)
			}
		})
	}
}