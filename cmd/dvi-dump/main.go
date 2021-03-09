// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main // import "git.sr.ht/~sbinet/star-tex/cmd/dvi-dump"

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"git.sr.ht/~sbinet/star-tex/dvi"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("dvi-dump: ")

	var (
		doJSON = flag.Bool("json", false, "enable JSON output")
	)

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatalf("missing input dvi file")
	}

	xmain(flag.Arg(0), *doJSON)
}

func xmain(fname string, doJSON bool) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("could not open DVI file %q: %+v", fname, err)
	}
	defer f.Close()

	switch {
	case doJSON:
		err = jsonProcess(f)
	default:
		err = interp(f)
	}
	if err != nil {
		log.Fatalf("could not process DVI file %q: %+v", fname, err)
	}
}

func jsonProcess(f *os.File) error {
	r := dvi.NewReader(f)
	w := os.Stdout
	o := json.NewEncoder(w)

	err := r.Read(func(cmd dvi.Cmd) error {
		var v struct {
			Cmd  string  `json:"cmd"`
			Args dvi.Cmd `json:"args,omitempty"`
		}
		v.Cmd = cmd.Name()
		v.Args = cmd
		return o.Encode(v)
	})

	if err != nil {
		return fmt.Errorf("could not read DVI file: %w", err)
	}

	return nil
}

func interp(f *os.File) error {
	var (
		r  = dvi.NewReader(f)
		w  = os.Stdout
		in = dvi.NewInterpreter(w)
	)

	err := r.Read(func(cmd dvi.Cmd) error {
		err := in.Run(cmd)
		if err != nil {
			return fmt.Errorf("could not interpret command %v: %w", cmd, err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("could not read DVI file: %w", err)
	}

	return nil
}
