package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{1000.0, 10.0},
		{100.0, 5.0},
		{500.0, 5.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-10, -9.0, -1.0, 0.0, 1.0, 2.0, 10.0, 100.0, 1000.99, 10000.01}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0 but got %f", result)
		}
		if amount >= 1000 && result != 10.0 {
			t.Errorf("Expected 10 but got %f", result)
		}
		if (amount > 0 && amount < 1000) && result != 5.0 {
			t.Errorf("Expected 5 but got %f", result)
		}
	})
}
