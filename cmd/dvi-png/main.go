// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command dvi-png converts a DVI document into a (set of) PNG file(s).
package main // import "star-tex.org/x/tex/cmd/dvi-png"

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"star-tex.org/x/tex/dvi"
	"star-tex.org/x/tex/kpath"
)

func main() {
	log.SetPrefix("dvi-png: ")
	log.SetFlags(0)

	var (
		texmf = flag.String("texmf", "", "path to TexMF root")
	)

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		log.Fatalf("missing DVI input file")
	}

	xmain(log.Writer(), flag.Arg(0), *texmf)
}

func xmain(stdout io.Writer, fname, texmf string) {
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("could not open DVI file %q: %+v", flag.Arg(0), err)
	}
	defer f.Close()

	ctx := kpath.New()
	if texmf != "" {
		ctx, err = kpath.NewFromFS(os.DirFS(texmf))
		if err != nil {
			log.Fatalf("could not create kpath context: %+v", err)
		}
	}

	err = interp(ctx, stdout, f)
	if err != nil {
		log.Fatalf("could not process DVI file %q: %+v", fname, err)
	}
}

func interp(ctx kpath.Context, stdout io.Writer, r io.Reader) error {
	render := pngRenderer{
		name: "out.png",
	}
	vm := dvi.NewMachine(
		dvi.WithContext(ctx),
		// dvi.WithLogOutput(stdout),
		dvi.WithRenderer(&render),
	)

	raw, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("could not read DVI program file: %w", err)
	}

	prog, err := dvi.Compile(raw)
	if err != nil {
		return fmt.Errorf("could not compile DVI program: %w", err)
	}

	err = vm.Run(prog)
	if err != nil {
		return fmt.Errorf("could not interpret DVI program: %w", err)
	}

	return nil
}
