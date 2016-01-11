// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathp

import (
	"fmt"
	"testing"

	"github.com/golangplus/strings"
	"github.com/golangplus/testing/assert"
)

func TestMin(t *testing.T) {
	data := []string{"5", "7", "1", "9", "3"}
	assert.Equal(t, "min", data[Min(len(data), stringsp.LessFunc(data))], "1")
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
	assert.Equal(t, "MaxI", MaxI(2, 3), 3)
	assert.Equal(t, "MaxI", MaxI(3, 2), 3)
}

func TestMaxI32(t *testing.T) {
	assert.Equal(t, "MaxI", MaxI32(2, 3), int32(3))
	assert.Equal(t, "MaxI", MaxI32(3, 2), int32(3))
}

func TestMaxI64(t *testing.T) {
	assert.Equal(t, "MaxI", MaxI64(2, 3), int64(3))
	assert.Equal(t, "MaxI", MaxI64(3, 2), int64(3))
}

func TestMinI(t *testing.T) {
	assert.Equal(t, "MinI", MinI(2, 3), 2)
	assert.Equal(t, "MinI", MinI(3, 2), 2)
}

func TestMinI32(t *testing.T) {
	assert.Equal(t, "MinI", MinI32(2, 3), int32(2))
	assert.Equal(t, "MinI", MinI32(3, 2), int32(2))
}

func TestMinI64(t *testing.T) {
	assert.Equal(t, "MinI", MinI64(2, 3), int64(2))
	assert.Equal(t, "MinI", MinI64(3, 2), int64(2))
}
