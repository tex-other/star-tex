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
	Size float64 // Size is the font size in DVI points.
	DPI  float64 // DPI is the dots per inch resolution
}

func defaultFaceOptions(font *tfm.Font) *FaceOptions {
	return &FaceOptions{
		Size: font.DesignSize().Float64(),
		DPI:  72,
	}
}

// Units are an integral number of abstract, scalable "font units". The em
// square is typically 1000 or 2048 "font units". This would map to a certain
// number (e.g. 30 pixels) of physical pixels, depending on things like the
// display resolution (DPI) and font size (e.g. a 12 point font).
type units int32

// scale returns x divided by unitsPerEm, rounded to the nearest fixed.Int12_20
// value (1/1048576th of a pixel).
func scale(x fixed.Int12_20, unitsPerEm units) fixed.Int12_20 {
	if x >= 0 {
		x += fixed.Int12_20(unitsPerEm) / 2
	} else {
		x -= fixed.Int12_20(unitsPerEm) / 2
	}
	return x / fixed.Int12_20(unitsPerEm)
}

func NewFace(font *Font, metrics *tfm.Font, opts *FaceOptions) *Face {
	if opts == nil {
		opts = defaultFaceOptions(metrics)
	}
	log.Printf("design: %v", opts.Size)
	return &Face{
		font:  font,
		tfm:   metrics,
		scale: fixed.Int12_20(0.5 + (opts.Size * opts.DPI * 64 / 72)),
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
		X: int(dot.X) - int(g.width)/2,
		Y: int(dot.Y) - int(g.height)/2,
	}

	dr = image.Rectangle{
		Min: p,
		Max: p.Add(image.Point{
			X: int(g.width),
			Y: int(g.height),
		}),
	}
	if false {
		maskp.X -= int(g.xoff)
		maskp.Y -= int(g.yoff) // FIXME(sbinet): y-axis direction.
	}

	if r == 'T' || r == 'a' || true {
		log.Printf(
			"glyph: %q h=%d w=%d dot=%+v --> p=%+v (off=(%d, %d), size=(%d,%d), d(%d,%d) tfmw: %d flag=%d",
			r,
			g.height, g.width,
			dot,
			p,
			g.xoff, g.yoff,
			g.width, g.height,
			g.dx, g.dy,
			g.wtfm,
			g.flag,
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
	advance = fixed.Int12_20((int64(g.width) * int64(face.scale)) >> 20)
	ok = true
	return
}

// GlyphBounds returns the bounding box of r's glyph, drawn at a dot equal
// to the origin, and that glyph's advance width.
//
// It returns !ok if the face does not contain a glyph for r.
//
// The glyph's ascent and descent are equal to -bounds.Min.Y and
// +bounds.Max.Y. The glyph's left-side and right-side bearings are equal
// to bounds.Min.X and advance-bounds.Max.X. A visual depiction of what
// these metrics are is at
// https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Art/glyphterms_2x.png
func (face *Face) GlyphBounds(r rune) (bounds fixed.Rectangle12_20, advance fixed.Int12_20, ok bool) {

	g, ok := face.font.gidx(r)
	if !ok {
		return
	}

	bounds = fixed.Rectangle12_20{
		Min: fixed.Point12_20{
			X: 0,
			Y: 0,
		},
		Max: fixed.Point12_20{
			X: fixed.Int12_20(g.width),
			Y: fixed.Int12_20(g.height),
		},
	}

	advance = fixed.Int12_20((int64(g.width) * int64(face.scale)) >> 20)
	ok = true
	return
}

var (
	_ font.Face = (*Face)(nil)
)

func sp2px(v fixed.Int12_20, dpi float64) int {
	return 0
}
