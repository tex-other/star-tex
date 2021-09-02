// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"

	"star-tex.org/x/tex/dvi"
	"star-tex.org/x/tex/font/fixed"
)

type pngRenderer struct {
	name string
	page int

	bkg color.Color

	pre   dvi.CmdPre
	post  dvi.CmdPost
	conv  float32 // converts DVI units to pixels
	tconv float32 // converts unmagnified DVI units to pixels

	img draw.Image
	err error
}

func (pr *pngRenderer) Init(pre *dvi.CmdPre, post *dvi.CmdPost) {
	pr.pre = *pre
	pr.post = *post
	res := float32(300)
	conv := float32(pr.pre.Num) / 254000.0 * (res / float32(pr.pre.Den))
	pr.tconv = conv
	pr.conv = conv * float32(pr.pre.Mag) / 1000.0

	if pr.bkg == nil {
		pr.bkg = color.White
	}
}

func (pr *pngRenderer) BOP(bop *dvi.CmdBOP) {
	if pr.err != nil {
		return
	}

	pr.page = int(bop.C0)

	bnd := image.Rect(0, 0,
		int(pr.pixels(int32(pr.post.Width))),
		int(pr.pixels(int32(pr.post.Height))),
	)
	//pr.img = image.NewRGBA(bnd)
	pr.img = image.NewPaletted(bnd, color.Palette{pr.bkg, color.Black})
	//draw.Draw(pr.img, bnd, image.NewUniform(pr.bkg), image.Point{}, draw.Over)
}

func (pr *pngRenderer) EOP() {
	if pr.err != nil {
		return
	}

	name := pr.name[:len(pr.name)-len(".png")] + "_" + strconv.Itoa(pr.page) + ".png"
	f, err := os.Create(name)
	if err != nil {
		if pr.err == nil {
			pr.err = fmt.Errorf("could not create output PNG file: %w", err)
		}
		return
	}
	defer f.Close()

	err = png.Encode(f, pr.img)
	if err != nil {
		if pr.err == nil {
			pr.err = fmt.Errorf("could not encode PNG image: %w", err)
		}
		return
	}

	err = f.Close()
	if err != nil {
		if pr.err == nil {
			pr.err = fmt.Errorf("could not close output PNG file %q: %w", name, err)
		}
		return
	}
}

func (pr *pngRenderer) DrawGlyph(x, y int32, font dvi.Font, glyph rune, c color.Color) {
	dot := fixed.Point12_20{
		X: fixed.Int12_20(pr.pixels(x)),
		Y: fixed.Int12_20(pr.pixels(y)),
	}
	dr, mask, maskp, adv, ok := font.Face().Glyph(dot, glyph)
	if !ok {
		if pr.err != nil {
			return
		}
		pr.err = fmt.Errorf("could not find glyph 0x%02x", glyph)
		return
	}
	_ = adv

	draw.DrawMask(
		pr.img,
		dr, image.NewUniform(c),
		maskp, mask,
		dr.Min, draw.Over,
	)

	//		log.Printf(
	//			"draw-glyph(%d, %d, %v, %q, %v)...",
	//			x, y, font, glyph, c,
	//		)
}

func (pr *pngRenderer) DrawRule(x, y, w, h int32, c color.Color) {
	log.Printf(
		"draw-rule(%d, %d, %d, %d, %v)...",
		x, y, w, h, c,
	)
	r := image.Rect(
		int(pr.pixels(x+0)), int(pr.pixels(y+0)),
		int(pr.pixels(x+w)), int(pr.pixels(y+h)),
	)

	draw.Draw(pr.img, r, image.NewUniform(c), image.Point{}, draw.Over)
}

func maxI32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func roundF32(v float32) int32 {
	if v > 0 {
		return int32(v + 0.5)
	}
	return int32(v - 0.5)
}

func (pr *pngRenderer) pixels(v int32) int32 {
	x := pr.conv * float32(v)
	return roundF32(x / 3)
	//return roundF32(x)
}

func (pr *pngRenderer) rulepixels(v int32) int32 {
	x := int32(pr.conv * float32(v))
	if float32(x) < pr.conv*float32(v) {
		return x + 1
	}
	return x
}

var (
	_ dvi.Renderer = (*pngRenderer)(nil)
)
