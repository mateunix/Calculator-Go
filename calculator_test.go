package main

import "testing"

func Test_Sum_Correct(t *testing.T) {
	test := sum(3, 2, 1)
	result := 6

	if test != result {
		t.Error("Expected result:", result, "Returned value:", test)
	}
}

func Test_Product_Correct(t *testing.T) {
	test := product(10, 10)
	result := 100

	if test != result {
		t.Error("Expected result:", result, "Returned value:", test)
	}
}

func Test_Subtraction_Incorrect(t *testing.T) {
	test := subtraction(10, 3, 2)
	result := 5

	if test != result {
		t.Error("Expected result:", result, "Returned value:", test)
	}
}

func Test_Division_Incorrect(t *testing.T) {
	test := division(4, 2)
	result := 2

	if test != result {
		t.Error("Expected result:", result, "Returned value:", test)
	}
}
