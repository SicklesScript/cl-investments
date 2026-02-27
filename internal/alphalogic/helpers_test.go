package alphalogic_test

import (
	"testing"

	"github.com/SicklesScript/cl-investments/internal/alphalogic"
)

/*
Fills dividend data struct with sample data points.
Calculate div uses the data to extract ttm div and prev div values.
*/
func TestCalculateDiv(t *testing.T) {
	// Create 8 sample data points to fill divdata slice of dividned data struct
	d1 := alphalogic.Data{
		ExDivDate: "2026-02-10",
		Amount:    "1.68",
	}
	d2 := alphalogic.Data{
		ExDivDate: "2025-11-10",
		Amount:    "1.68",
	}
	d3 := alphalogic.Data{
		ExDivDate: "2025-08-08",
		Amount:    "1.68",
	}
	d4 := alphalogic.Data{
		ExDivDate: "2025-05-09",
		Amount:    "1.68",
	}
	d5 := alphalogic.Data{
		ExDivDate: "2025-02-10",
		Amount:    "1.67",
	}
	d6 := alphalogic.Data{
		ExDivDate: "2024-11-12",
		Amount:    "1.67",
	}
	d7 := alphalogic.Data{
		ExDivDate: "2024-08-09",
		Amount:    "1.67",
	}
	d8 := alphalogic.Data{
		ExDivDate: "2024-05-09",
		Amount:    "1.67",
	}
	// Fill dd struct with data
	dd := alphalogic.DividendData{
		Symbol:  "IBM",
		DivData: []alphalogic.Data{d1, d2, d3, d4, d5, d6, d7, d8},
	}
	// Get ttm and prev div from dd struct (3 shares)
	ttm, prev := dd.CalculateDiv("3")

	// Check ttm data matches
	expectedTTMDiv := 3 * 4 * 1.68
	if ttm != expectedTTMDiv {
		t.Errorf("TTM div values do not match. Expected %.2f, got %.2f", expectedTTMDiv, ttm)
	}

	// Check prev data matches
	expectedPrevDiv := 3 * 4 * 1.67
	if prev != expectedPrevDiv {
		t.Errorf("Prev div values do not match. Expected %.2f, got %.2f", expectedPrevDiv, prev)
	}
}
