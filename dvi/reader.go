// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dvi

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type Reader struct {
	r   *bufio.Reader
	buf []byte
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r:   bufio.NewReader(r),
		buf: make([]byte, 4),
	}
}

func (r *Reader) Read(f func(cmd Cmd) error) error {
	var err error

loop:
	for {
		_, err = io.ReadFull(r.r, r.buf[:1])
		if err != nil {
			if err == io.EOF {
				err = nil
				break loop
			}
			return fmt.Errorf("dvi: could not read opcode: %w", err)
		}

		err = r.r.UnreadByte()
		if err != nil {
			return fmt.Errorf("dvi: could not unread opcode: %w", err)
		}

		var (
			op  = opCode(r.buf[0])
			cmd = op.cmd()
		)
		if cmd == nil {
			return fmt.Errorf("dvi: unknown opcode %v (v=0x%x)", op, op)
		}

		cmd.read(r)
		if r.err != nil {
			return fmt.Errorf("dvi: could not unmarshal dvi cmd 0x%x: %w", op, r.err)
		}

		err = f(cmd)
		if err != nil {
			return fmt.Errorf("dvi: could not call user provided function: %w", err)
		}

		if cmd.opcode() == opPostPost {
			break
		}
	}

	return err
}

func (r *Reader) tailU8() (uint8, error) {
	r.buf[0] = 0
	_, err := io.ReadFull(r.r, r.buf[:1])
	return r.buf[0], err
}

func (r *Reader) op() uint8 {
	if r.err != nil {
		return 0
	}
	_, r.err = r.r.Read(r.buf[:1])
	return r.buf[0]
}

func (r *Reader) readU8() uint8 {
	return r.op()
}

func (r *Reader) readU16() uint16 {
	if r.err != nil {
		return 0
	}
	_, r.err = r.r.Read(r.buf[:2])
	return binary.BigEndian.Uint16(r.buf[:2])
}

func (r *Reader) readU24() uint32 {
	if r.err != nil {
		return 0
	}
	_, r.err = r.r.Read(r.buf[:3])
	return uint32(r.buf[0])<<16 | uint32(r.buf[1])<<8 | uint32(r.buf[2])
}

func (r *Reader) readI24() int32 {
	if r.err != nil {
		return 0
	}
	_, r.err = r.r.Read(r.buf[:3])
	if r.buf[0] < 128 {
		return int32(uint32(r.buf[0])<<16 | uint32(r.buf[1])<<8 | uint32(r.buf[2]))
	}
	return int32((uint32(r.buf[0])-256)<<16 | uint32(r.buf[1])<<8 | uint32(r.buf[2]))

}

func (r *Reader) readU32() uint32 {
	if r.err != nil {
		return 0
	}
	_, r.err = r.r.Read(r.buf[:4])
	return binary.BigEndian.Uint32(r.buf[:4])
}

func (r *Reader) readBuf(n int) []byte {
	if r.err != nil {
		return nil
	}
	buf := make([]byte, n)
	_, r.err = r.r.Read(buf)
	return buf
}
