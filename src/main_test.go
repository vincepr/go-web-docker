package main

import "testing"

func TestSum(t *testing.T){
	if add(2,5) != 7{
		t.Fail()
	}
	if add(2,99) != 101{
		t.Fail()
	}
	if add(87,111) != 198{
		t.Fail()
	}
}
func TestProduct(t *testing.T){
	if multiply(2,5) != 10{
		t.Fail()
	}
	if multiply(2,60) != 120{
		t.Fail()
	}
	if multiply(8,111) != 888{
		t.Fail()
	}
}