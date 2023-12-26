package Util

import (
	"fmt"
	"testing"
)

func TestCompareMax(t *testing.T) {
	i := 10
	j := 5
	fmt.Println(Max(i, j))
}

func TestCompareMin(t *testing.T) {
	i := 10
	j := 5
	fmt.Println(Min(i, j))
}

func TestCompareStrToFloatMin(t *testing.T) {
	i := "5.82"
	j := "5.83"
	fmt.Println(StrToFloatMin(i, j))
}

func TestCompareStrToFloatMax(t *testing.T) {
	i := "5.82"
	j := "5.83"
	fmt.Println(StrToFloatMax(i, j))
}
