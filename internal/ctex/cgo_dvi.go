// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ctex

//#include "ctex-consts.h"
//#include <stdio.h>
import "C"

import (
	"unsafe"
)

type ctex_dvi_t int

func (dvi ctex_dvi_t) get() *DVI {
	return dviMap[int(dvi)]
}

type DVI struct {
	pages  int32
	offset int32
	gone   int32 // the number of bytes already output to file.

	h int32 // current horizontal position
	v int32 // current vertical position

	f *C.FILE

	buf      [C.dvi_buf_size + 1]uint8
	half_buf int32
	limit    int32
	ptr      int32

	font int32 // font index.
}

func (dvi *DVI) write(beg, end int32) {
	buf := unsafe.Pointer(&dvi.buf[beg])
	C.fwrite(buf, 1, C.size_t(end+1-beg), dvi.f)
}

func (dvi *DVI) swap() {
	switch {
	case dvi.limit == C.dvi_buf_size:
		dvi.write(0, dvi.half_buf-1)
		dvi.limit = dvi.half_buf
		dvi.offset += int32(C.dvi_buf_size)
		dvi.ptr = 0
	default:
		dvi.write(dvi.half_buf, C.dvi_buf_size-1)
		dvi.limit = int32(C.dvi_buf_size)

	}
	dvi.gone += dvi.half_buf
}

func (dvi *DVI) writeU8(v uint8) {
	dvi.buf[dvi.ptr] = v
	dvi.ptr++

	if dvi.ptr == dvi.limit {
		dvi.swap()
	}
}

func (dvi *DVI) writeU32(x uint32) {
	dvi.writeU8(uint8(x >> 24))
	dvi.writeU8(uint8(x >> 16))
	dvi.writeU8(uint8(x >> 8))
	dvi.writeU8(uint8(x))
}

var (
	dviMap = make(map[int]*DVI)
)

//export ctex_dvi_new
func ctex_dvi_new() ctex_dvi_t {
	var (
		dvi = DVI{
			half_buf: int32(C.dvi_buf_size / 2),
			limit:    int32(C.dvi_buf_size),
		}
		hdl = len(dviMap)
	)
	dviMap[hdl] = &dvi
	return ctex_dvi_t(hdl)
}

//export ctex_dvi_at
func ctex_dvi_at(self ctex_dvi_t, i C.int) uint8 {
	return self.get().buf[i]
}

//export ctex_dvi_set
func ctex_dvi_set(self ctex_dvi_t, i C.int, v uint8) {
	self.get().buf[i] = v
}

//export ctex_dvi_file
func ctex_dvi_file(self ctex_dvi_t) *C.FILE {
	return self.get().f
}

//export ctex_dvi_set_file
func ctex_dvi_set_file(self ctex_dvi_t, f *C.FILE) {
	self.get().f = f
}

//export ctex_dvi_fclose
func ctex_dvi_fclose(self ctex_dvi_t) int {
	var (
		rc  int
		dvi = self.get()
	)
	if dvi.f != nil {
		rc = int(C.fclose(dvi.f))
		dvi.f = nil
	}
	return rc
}

//export ctex_dvi_add_page
func ctex_dvi_add_page(self ctex_dvi_t) {
	dvi := self.get()
	dvi.pages++
}

//export ctex_dvi_pages
func ctex_dvi_pages(self ctex_dvi_t) int32 {
	return self.get().pages
}

//export ctex_dvi_offset
func ctex_dvi_offset(self ctex_dvi_t) int32 {
	return self.get().offset
}

//export ctex_dvi_gone
func ctex_dvi_gone(self ctex_dvi_t) int32 {
	return self.get().gone
}

//export ctex_dvi_flush
func ctex_dvi_flush(self ctex_dvi_t) {
	dvi := self.get()
	if dvi.limit == dvi.half_buf {
		dvi.write(dvi.half_buf, C.dvi_buf_size-1)
	}
	if dvi.ptr > 0 {
		dvi.write(0, dvi.ptr-1)
	}
}

//export ctex_dvi_wU8
func ctex_dvi_wU8(self ctex_dvi_t, v uint8) {
	self.get().writeU8(v)
}

//export ctex_dvi_four
func ctex_dvi_four(self ctex_dvi_t, x int32) {
	dvi := self.get()
	dvi.writeU8(uint8(x >> 24))
	dvi.writeU8(uint8(x >> 16))
	dvi.writeU8(uint8(x >> 8))
	dvi.writeU8(uint8(x))
}

//export ctex_dvi_pop
func ctex_dvi_pop(self ctex_dvi_t, l int32) {
	dvi := self.get()
	if l == dvi.offset+dvi.ptr && dvi.ptr > 0 {
		dvi.ptr--
		return
	}
	const dvi_cmd_pop = 142
	dvi.writeU8(dvi_cmd_pop)
}

//export ctex_dvi_pos
func ctex_dvi_pos(self ctex_dvi_t) int32 {
	dvi := self.get()
	return dvi.offset + dvi.ptr
}

//export ctex_dvi_cap
func ctex_dvi_cap(self ctex_dvi_t) int32 {
	dvi := self.get()
	return C.dvi_buf_size - dvi.ptr
}

//export ctex_dvi_set_font
func ctex_dvi_set_font(self ctex_dvi_t, f int32) {
	self.get().font = f
}

//export ctex_dvi_get_font
func ctex_dvi_get_font(self ctex_dvi_t) int32 {
	return self.get().font
}

//export ctex_dvi_set_h
func ctex_dvi_set_h(self ctex_dvi_t, h int32) {
	self.get().h = h
}

//export ctex_dvi_set_v
func ctex_dvi_set_v(self ctex_dvi_t, v int32) {
	self.get().v = v
}

//export ctex_dvi_get_h
func ctex_dvi_get_h(self ctex_dvi_t) int32 {
	return self.get().h
}

//export ctex_dvi_get_v
func ctex_dvi_get_v(self ctex_dvi_t) int32 {
	return self.get().v
}

//export ctex_dvi_font_def
func ctex_dvi_font_def(self ctex_dvi_t,
	fid int32, chksum uint32,
	size, dsize int32,
	areasz C.size_t, area *C.char,
	namesz C.size_t, name *C.char) {

	const fnt_def1 = 243
	dvi := self.get()
	dvi.writeU8(fnt_def1)
	dvi.writeU8(uint8(fid))
	dvi.writeU32(chksum)
	dvi.writeU32(uint32(size))
	dvi.writeU32(uint32(dsize))
	dvi.writeU8(uint8(areasz))
	dvi.writeU8(uint8(namesz))
	for _, v := range C.GoStringN(area, C.int(areasz)) {
		dvi.writeU8(uint8(v))
	}
	for _, v := range C.GoStringN(name, C.int(namesz)) {
		dvi.writeU8(uint8(v))
	}
}

//export ctex_dvi_wcmd
func ctex_dvi_wcmd(self ctex_dvi_t, cmd uint8, v int32) {
	u := uint32(v)
	if v < 0 {
		u = uint32(-v)
	}

	dvi := self.get()
	switch {
	case u >= 1<<23:
		dvi.writeU8(cmd + 3)
		dvi.writeU32(uint32(v))
	case u >= 1<<15:
		dvi.writeU8(cmd + 2)
		dvi.writeU8(uint8(v >> 16))
		dvi.writeU8(uint8(v >> 8))
		dvi.writeU8(uint8(v))
	case u >= 1<<7:
		dvi.writeU8(cmd + 1)
		dvi.writeU8(uint8(v >> 8))
		dvi.writeU8(uint8(v))
	default:
		dvi.writeU8(cmd)
		dvi.writeU8(uint8(v))
	}
}
