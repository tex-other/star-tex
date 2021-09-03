// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fixed

// Point12_20 is a 12.20 fixed-point coordinate pair.
//
// It is analogous to the image.Point type in the standard library.
type Point12_20 struct {
	X, Y Int12_20
}

type Rectangle12_20 struct {
	Min Point12_20
	Max Point12_20
}
