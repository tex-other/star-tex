// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package type1

import (
	"io/fs"
	"testing"

	"star-tex.org/x/tex/kpath"
)

func TestParse(t *testing.T) {
	ctx := kpath.New()
	pfbName, err := ctx.Find("cmr10.pfb")
	if err != nil {
		t.Fatalf("could not find type1 font file: %+v", err)
	}

	pfb, err := fs.ReadFile(ctx.FS(), pfbName)
	if err != nil {
		t.Fatalf("could not read file %q: %+v", pfbName, err)
	}

	afmName, err := ctx.Find("cmr10.afm")
	if err != nil {
		t.Fatalf("could not find type1 metrics font file: %+v", err)
	}

	afm, err := fs.ReadFile(ctx.FS(), afmName)
	if err != nil {
		t.Fatalf("could not read file %q: %+v", afmName, err)
	}

	fnt, err := ParsePFB(pfb, afm)
	if err != nil {
		t.Fatalf("could not parse file %q: %+v", pfbName, err)
	}

	t.Fatalf("font: %+v", fnt.fnt)
}
