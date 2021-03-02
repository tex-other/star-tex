// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ctex

//
import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	cEOF = -1
)

//export ctex_file
type ctex_file int

type iofile struct {
	f   *os.File
	err error
}

func (f *iofile) String() string {
	if f == nil {
		return "(*iofile)(nil)"
	}
	str := ""
	if f.f != nil {
		str = fmt.Sprintf(" (name=%s)", f.f.Name())
	}
	return fmt.Sprintf("&iofile{f: %v%s, err: %v}", f.f, str, f.err)
}

func (f *iofile) Read(p []byte) (int, error) {
	n, err := f.f.Read(p)
	if err != nil && f.err == nil {
		f.err = err
	}
	return n, err
}

func (f *iofile) ReadAt(p []byte, off int64) (int, error) {
	n, err := f.f.ReadAt(p, off)
	if err != nil && f.err == nil {
		f.err = err
	}
	return n, err
}

func (f *iofile) Write(p []byte) (int, error) {
	n, err := f.f.Write(p)
	if err != nil && f.err == nil {
		f.err = err
	}
	return n, err
}

func (f *iofile) Seek(offset int64, whence int) (int64, error) {
	n, err := f.f.Seek(offset, whence)
	if err != nil && f.err == nil {
		f.err = err
	}
	return n, err
}

func (f *iofile) Close() error {
	err := f.f.Close()
	if err != nil && f.err == nil {
		f.err = err
	}
	return err
}

var (
	ioMap = map[int]*iofile{-1: nil, 0: nil}

	_ io.Reader   = (*iofile)(nil)
	_ io.ReaderAt = (*iofile)(nil)
	_ io.Writer   = (*iofile)(nil)
	_ io.Seeker   = (*iofile)(nil)
	_ io.Closer   = (*iofile)(nil)
)

func (f ctex_file) get() *iofile {
	return ioMap[int(f)]
}

//export ctex_io_fopen
func ctex_io_fopen(name *C.char) ctex_file {
	var (
		fname = C.GoString(name)
		hdl   = len(ioMap)
	)
	f, err := os.Open(fname)
	if err != nil {
		return 0
	}
	ioMap[hdl] = &iofile{f: f}
	return ctex_file(hdl)
}

//export ctex_io_fcreate
func ctex_io_fcreate(name *C.char) ctex_file {
	var (
		fname = C.GoString(name)
		hdl   = len(ioMap)
	)
	f, err := os.Create(fname)
	if err != nil {
		log.Panicf("ctex: could not create %q: %+v", fname, err)
	}
	ioMap[hdl] = &iofile{f: f}
	return ctex_file(hdl)
}

//export ctex_io_fflush
func ctex_io_fflush(self ctex_file) C.int {
	f := self.get()
	err := f.f.Sync()
	if err != nil {
		if f.err == nil {
			f.err = err
		}
		return 1
	}
	return 0
}

//export ctex_io_fclose
func ctex_io_fclose(self ctex_file) C.int {
	var (
		rc C.int
		f  = self.get()
	)
	ioMap[int(self)] = nil

	if f == nil || f.f == nil {
		return 0
	}

	err := f.Close()
	if err != nil {
		rc = 1
		log.Printf("ctex: could not close %q: %+v", f.f.Name(), err)
		return rc
	}

	return rc
}

//export ctex_io_ferror
func ctex_io_ferror(self ctex_file) C.int {
	f := self.get()
	if f.err != nil {
		return 1
	}
	return 0
}

//export ctex_io_feof
func ctex_io_feof(self ctex_file) C.int {
	f := self.get()
	if f.err == io.EOF {
		return 1
	}
	return 0
}

//export ctex_io_fgetc
func ctex_io_fgetc(self ctex_file) C.int {
	return C.int(ctex_io_readU8(self))
}

//export ctex_io_fpeek
func ctex_io_fpeek(self ctex_file) C.int {
	f := self.get()
	if f.err == io.EOF {
		return cEOF
	}
	pos, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return cEOF // FIXME(sbinet)
	}
	buf := make([]byte, 1)
	_, err = f.ReadAt(buf, pos)
	if err != nil {
		return cEOF // FIXME(sbinet)
	}

	return C.int(buf[0])
}

//export ctex_io_fprintf
func ctex_io_fprintf(self ctex_file, str *C.char) {
	v := C.GoString(str)
	_, _ = self.get().Write([]byte(v))
}

//export ctex_io_fprintfU8
func ctex_io_fprintfU8(self ctex_file, format *C.char, v uint8) {
	_, _ = fmt.Fprintf(self.get(), C.GoString(format), uint8(v))
}

//export ctex_io_readU8
func ctex_io_readU8(self ctex_file) uint8 {
	var (
		f   = self.get()
		buf = make([]byte, 1)
	)
	_, err := io.ReadFull(f, buf)
	if err != nil {
		return 0
	}

	return uint8(buf[0])
}

//export ctex_io_readU32
func ctex_io_readU32(self ctex_file) uint32 {
	var (
		f   = self.get()
		buf = make([]byte, 4)
	)
	_, err := io.ReadFull(f, buf)
	if err != nil {
		return 0
	}

	return binary.BigEndian.Uint32(buf)
}

//export ctex_io_writeU32
func ctex_io_writeU32(self ctex_file, v uint32) {
	var (
		f   = self.get()
		buf = make([]byte, 4)
	)
	binary.BigEndian.PutUint32(buf, v)

	_, err := f.Write(buf)
	if err != nil {
		return
	}
}
