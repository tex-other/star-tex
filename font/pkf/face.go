// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkf

import (
	"image"
	"log"

	"star-tex.org/x/tex/font"
	"star-tex.org/x/tex/font/fixed"
	"star-tex.org/x/tex/font/tfm"
)

// Face implements the font.Face interface for PK fonts.
type Face struct {
	font  *Font
	tfm   *tfm.Font
	scale fixed.Int12_20
}

// FaceOptions describes the possible options given to NewFace when
// creating a new Face from a Font.
type FaceOptions struct {
	Size fixed.Int12_20 // Size is the font size in DVI points.
}

func defaultFaceOptions(font *tfm.Font) *FaceOptions {
	return &FaceOptions{
		Size: font.DesignSize(),
	}
}

func NewFace(font *Font, metrics *tfm.Font, opts *FaceOptions) *Face {
	if opts == nil {
		opts = defaultFaceOptions(metrics)
	}
	return &Face{
		font:  font,
		tfm:   metrics,
		scale: opts.Size,
	}
}

// Name returns the name of the font face.
func (face *Face) Name() string {
	return face.tfm.Name()
}

func (face *Face) Glyph(dot fixed.Point12_20, r rune) (
	dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int12_20, ok bool) {

	g, ok := face.font.gidx(r)
	if !ok {
		return
	}

	g.unpack()

	p := image.Point{
		X: int(dot.X),
		Y: int(dot.Y),
	}

	dr = image.Rectangle{
		Min: p,
		Max: p.Add(image.Point{
			X: int(g.width),
			Y: int(g.height),
		}),
	}
	if true {
		maskp.X = int(g.xoff)
		maskp.Y = int(g.yoff) // FIXME(sbinet): y-axis direction.
	}

	if r == 'T' || r == 'a' {
		log.Printf(
			"glyph: %q h=%d w=%d dot=%+v --> p=%+v",
			r,
			g.height, g.width,
			dot,
			p,
		)
	}

	msk := g.Mask()
	mask = &msk
	// w := len(g.mask) / int(g.height)
	//	mask = &image.Alpha{
	//		Stride: w,
	//		Pix:    g.mask,
	//		Rect:   image.Rect(0, 0, w, int(g.height)),
	//	}
	advance = fixed.Int12_20(g.width)
	ok = true
	return
}

var (
	_ font.Face = (*Face)(nil)
)
