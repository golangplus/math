// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathp

import (
	"fmt"
	"testing"

	"github.com/golangplus/testing/assert"
)

func TestMin(t *testing.T) {
	data := []int{5, 7, 1, 9, 3}
	assert.Equal(t, "min", Min(len(data), func(i, j int) bool {
		return data[i] < data[j]
	}), 2)
}

func ExampleMinStrings() {
	names := []string{
		"David",
		"Richard",
		"Sophie",
		"Mengxi",
	}
	minIndex := Min(len(names), func(i, j int) bool {
		return names[i] < names[j]
	})
	fmt.Println(minIndex, names[minIndex])

	// Output:
	// 0 David
}

func TestMax(t *testing.T) {
	data := []int{5, 7, 1, 9, 3}
	assert.Equal(t, "max", Max(len(data), func(i, j int) bool {
		return data[i] < data[j]
	}), 3)
}

func ExampleMaxFloat() {
	floats := []float64{
		2.5, 3.1, 1.7, 0.5,
	}
	maxIndex := Max(len(floats), func(i, j int) bool {
		return floats[i] < floats[j]
	})
	fmt.Println(maxIndex)

	// Output:
	// 1
}

func TestMaxI(t *testing.T) {
	if MaxI(2, 3) != 3 {
		t.Errorf("MaxI(2, 3) returns %v instead of 3", MaxI(2, 3))
	}
	if MaxI(3, 2) != 3 {
		t.Errorf("MaxI(3, 2) returns %v instead of 3", MaxI(3, 2))
	}
}

func ExampleMaxI() {
	fmt.Println(MaxI(2, 3))
	// OUTPUT:
	// 3
}

func ExampleMaxI32() {
	fmt.Println(MaxI32(2, 3))
	// OUTPUT:
	// 3
}

func ExampleMaxI64() {
	fmt.Println(MaxI64(2, 3))
	// OUTPUT:
	// 3
}

func ExampleMinI() {
	fmt.Println(MinI(2, 3))
	// OUTPUT:
	// 2
}

func ExampleMinI32() {
	fmt.Println(MinI32(2, 3))
	// OUTPUT:
	// 2
}

func ExampleMinI64() {
	fmt.Println(MinI64(2, 3))
	// OUTPUT:
	// 2
}
