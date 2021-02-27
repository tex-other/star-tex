// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dvi

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	w io.Writer

	buf   []byte
	state state
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:   w,
		buf: make([]byte, 4),
	}
}

func (w *Writer) op(o opCode) {
	w.buf[0] = uint8(o)
	w.w.Write(w.buf[:1])
}

func (w *Writer) writeU8(v uint8) {
	w.buf[0] = v
	w.w.Write(w.buf[:1])
}

func (w *Writer) writeU16(v uint16) {
	binary.BigEndian.PutUint16(w.buf, v)
	w.w.Write(w.buf[:2])
}

func (w *Writer) writeU24(v uint32) {
	w.buf[0] = uint8(v >> 16)
	w.buf[1] = uint8(v >> 8)
	w.buf[2] = uint8(v)
	w.w.Write(w.buf[:3])
}

func (w *Writer) writeI24(v int32) {
	w.buf[0] = uint8(v >> 16)
	w.buf[1] = uint8(v >> 8)
	w.buf[2] = uint8(v)
	w.w.Write(w.buf[:3])
}

func (w *Writer) writeU32(v uint32) {
	binary.BigEndian.PutUint32(w.buf, v)
	w.w.Write(w.buf[:4])
}

func (w *Writer) writeBuf(v []byte) {
	w.w.Write(v)
}
