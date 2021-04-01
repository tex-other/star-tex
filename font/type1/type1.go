// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package type1 implements a decoder for Type1 PostScript fonts.
package type1 // import "star-tex.org/x/tex/font/type1"

import (
	"bytes"

	"github.com/speedata/fonts/type1"
	"golang.org/x/image/font"
)

type Font struct {
	pfb []byte
	afm []byte

	fnt *type1.Type1
}

func (f *Font) Metrics() (font.Metrics, error) {
	panic("not implemented")
}

func ParsePFB(pfb, afm []byte) (*Font, error) {
	var (
		err error
		fnt type1.Type1
	)

	err = fnt.ParsePFB(bytes.NewReader(pfb))
	if err != nil {
		return nil, err
	}

	err = fnt.ParseAFM(bytes.NewReader(afm))
	if err != nil {
		return nil, err
	}

	return &Font{
		pfb: pfb,
		afm: afm,
		fnt: &fnt,
	}, nil
}
