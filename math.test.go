package main

import "testing"

func TestSum(t *testing.T) {

	total := sum(15, 15)

	if total != 30 {
		t.ErrorF("Invalid sum result: Expected -> %d Received -> %d", total, 30)
	}
}
