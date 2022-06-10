package main

import "testing"

func TestSum(t *testing.T) {

	result := soma(2,3)

	if result != 5 {
		t.Error("The result must be 5")
	}

}