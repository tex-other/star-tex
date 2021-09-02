// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package font defines an interface for font faces, for drawing text on an
// image.
package font // import "star-tex.org/x/tex/font"

import (
	"image"

	"star-tex.org/x/tex/font/fixed"
)

// Face is a font face. Its glyphs are often derived from a font file, such as
// "cmr10.pfb", but a face has a specific size, style, weight and
// hinting. For example, the 12pt and 18pt versions of Computer Modern are two
// different faces, even if derived from the same font file.
//
// A Face is not safe for concurrent use by multiple goroutines, as its methods
// may re-use implementation-specific caches and mask image buffers.
//
// To create a Face, look to other packages that implement specific font file
// formats.
type Face interface {
	// Glyph returns the draw.DrawMask parameters (dr, mask, maskp) to draw r's
	// glyph at the sub-pixel destination location dot, and that glyph's
	// advance width.
	//
	// It returns !ok if the face does not contain a glyph for r.
	//
	// The contents of the mask image returned by one Glyph call may change
	// after the next Glyph call. Callers that want to cache the mask must make
	// a copy.
	Glyph(dot fixed.Point12_20, r rune) (
		dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int12_20, ok bool)
}
