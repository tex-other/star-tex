// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dvi

import (
	"image"
	"os"

	"star-tex.org/x/tex/font"
	"star-tex.org/x/tex/font/fixed"
	"star-tex.org/x/tex/font/pkf"
	"star-tex.org/x/tex/font/tfm"
	"star-tex.org/x/tex/kpath"
)

// Font describes a DVI font, with TeX Font Metrics and its
// associated font glyph data.
type Font struct {
	font *tfm.Font
	face *tfm.Face
}

func (fnt *Font) Face() font.Face {
	// return fnt.face
	return fakeFace{}
}

type fakeFace struct{}

func (fakeFace) Glyph(dot fixed.Point12_20, r rune) (
	dr image.Rectangle, mask image.Image, maskp image.Point, advance fixed.Int12_20, ok bool) {
	return pkFace.Glyph(dot, r)
}

func (fakeFace) GlyphBounds(r rune) (bounds fixed.Rectangle12_20, advance fixed.Int12_20, ok bool) {
	return pkFace.GlyphBounds(r)
}

var pkFace font.Face

func init() {
	ktx := kpath.New()
	fname, err := ktx.Find("cmr10.pk")
	if err != nil {
		panic(err)
	}
	fpkf, err := os.Open("/home/binet/.texlive/texmf-var/fonts/pk/ljfour/public/cm/cmr10.600pk")
	//fpkf, err := ktx.Open(fname)
	//if err != nil {
	//	panic(err)
	//}
	defer fpkf.Close()

	fname, err = ktx.Find("cmr10.tfm")
	if err != nil {
		panic(err)
	}
	ftfm, err := ktx.Open(fname)
	if err != nil {
		panic(err)
	}
	defer ftfm.Close()

	pfnt, err := pkf.Parse(fpkf)
	if err != nil {
		panic(err)
	}

	tfnt, err := tfm.Parse(ftfm)
	if err != nil {
		panic(err)
	}

	pkFace = pkf.NewFace(pfnt, &tfnt, nil)
}
