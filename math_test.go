// Copyright 2015 The Golang Plus Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathp

import (
	"fmt"
)

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
