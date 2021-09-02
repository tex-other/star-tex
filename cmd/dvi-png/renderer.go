// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"

	"star-tex.org/x/tex/dvi"
)

type pngRenderer struct {
	name string
	page int

	xmax int32
	ymax int32

	pre   dvi.CmdPre
	conv  float32 // converts DVI units to pixels
	tconv float32 // converts unmagnified DVI units to pixels

	cmds []func()
	err  error
}

func (pr *pngRenderer) Init(pre *dvi.CmdPre) {
	pr.pre = *pre
	res := float32(300)
	conv := float32(pr.pre.Num) / 254000.0 * (res / float32(pr.pre.Den))
	pr.tconv = conv
	pr.conv = conv * float32(pr.pre.Mag) / 1000.0
}

func (pr *pngRenderer) BOP(bop *dvi.CmdBOP) {
	if pr.err != nil {
		return
	}

	pr.page = int(bop.C0)
	log.Printf(">>> bop: %+v", bop)

	pr.xmax = 0
	pr.ymax = 0
	pr.cmds = pr.cmds[:0]
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

	for _, cmd := range pr.cmds {
		cmd()
	}

	var (
		xmax   = int(pr.pixels(pr.xmax))
		ymax   = int(pr.pixels(pr.ymax))
		bounds = image.Rect(0, 0, xmax, ymax)
	)
	log.Printf("==> %+v", bounds)
	img := image.NewRGBA(bounds)
	err = png.Encode(f, img)
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
	pr.xmax = maxI32(x, pr.xmax) // FIXME(sbinet): add glyph extent to x
	pr.ymax = maxI32(y, pr.ymax) // FIXME(sbinet): add glyph extent to y
	pr.cmds = append(pr.cmds, func() {
		log.Printf(
			"draw-glyph(%d, %d, %v, %q, %v)...",
			x, y, font, glyph, c,
		)
	})
}

func (pr *pngRenderer) DrawRule(x, y, w, h int32, c color.Color) {
	pr.xmax = maxI32(x+w, pr.xmax)
	pr.ymax = maxI32(y+h, pr.ymax)
	pr.cmds = append(pr.cmds, func() {
		log.Printf(
			"draw-rule(%d, %d, %d, %d, %v)...",
			x, y, w, h, c,
		)
	})
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
	return roundF32(x)
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
