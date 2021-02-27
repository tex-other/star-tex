// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dvi

import (
	"fmt"
)

// Cmd is a DVI command.
type Cmd interface {
	opcode() opCode
	Name() string

	write(w *Writer)
	read(r *Reader)
}

type CmdSetChar struct {
	Value uint8 `json:"-"`
}

func (c CmdSetChar) opcode() opCode  { return opCode(c.Value) }
func (c CmdSetChar) Name() string    { return fmt.Sprintf("set_char_%d", c.Value) }
func (c CmdSetChar) write(w *Writer) { w.op(c.opcode()) }
func (c *CmdSetChar) read(r *Reader) { c.Value = r.op() }

type CmdSet1 struct {
	Value uint32 `json:"v"`
}

func (CmdSet1) opcode() opCode { return opSet1 }
func (CmdSet1) Name() string   { return "set1" }
func (c CmdSet1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdSet1) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU8())
}

type CmdSet2 struct {
	Value uint32 `json:"v"`
}

func (CmdSet2) opcode() opCode { return opSet2 }
func (CmdSet2) Name() string   { return "set2" }
func (c CmdSet2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdSet2) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU16())
}

type CmdSet3 struct {
	Value uint32 `json:"v"`
}

func (CmdSet3) opcode() opCode { return opSet3 }
func (CmdSet3) Name() string   { return "set3" }
func (c CmdSet3) write(w *Writer) {
	w.op(c.opcode())
	w.writeU24(uint32(c.Value))
}
func (c *CmdSet3) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU24())
}

type CmdSet4 struct {
	Value int32 `json:"v"`
}

func (CmdSet4) opcode() opCode { return opSet4 }
func (CmdSet4) Name() string   { return "set4" }
func (c CmdSet4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdSet4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdSetRule struct {
	Height int32 `json:"h"`
	Width  int32 `json:"w"`
}

func (CmdSetRule) opcode() opCode { return opSetRule }
func (CmdSetRule) Name() string   { return "set_rule" }
func (c CmdSetRule) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Height))
	w.writeU32(uint32(c.Width))
}
func (c *CmdSetRule) read(r *Reader) {
	_ = r.op()
	c.Height = int32(r.readU32())
	c.Width = int32(r.readU32())
}

type CmdPut1 struct {
	Value uint32 `json:"v"`
}

func (CmdPut1) opcode() opCode { return opPut1 }
func (CmdPut1) Name() string   { return "put1" }
func (c CmdPut1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdPut1) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU8())
}

type CmdPut2 struct {
	Value uint32 `json:"v"`
}

func (CmdPut2) opcode() opCode { return opPut2 }
func (CmdPut2) Name() string   { return "put2" }
func (c CmdPut2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdPut2) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU16())
}

type CmdPut3 struct {
	Value uint32 `json:"v"`
}

func (CmdPut3) opcode() opCode { return opPut3 }
func (CmdPut3) Name() string   { return "put3" }
func (c CmdPut3) write(w *Writer) {
	w.op(c.opcode())
	w.writeU24(uint32(c.Value))
}
func (c *CmdPut3) read(r *Reader) {
	_ = r.op()
	c.Value = uint32(r.readU24())
}

type CmdPut4 struct {
	Value int32 `json:"v"`
}

func (CmdPut4) opcode() opCode { return opPut4 }
func (CmdPut4) Name() string   { return "put4" }
func (c CmdPut4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdPut4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdPutRule struct {
	Height int32 `json:"h"`
	Width  int32 `json:"w"`
}

func (CmdPutRule) opcode() opCode { return opPutRule }
func (CmdPutRule) Name() string   { return "put_rule" }
func (c CmdPutRule) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Height))
	w.writeU32(uint32(c.Width))
}
func (c *CmdPutRule) read(r *Reader) {
	_ = r.op()
	c.Height = int32(r.readU32())
	c.Width = int32(r.readU32())
}

type CmdNOP struct{}

func (CmdNOP) opcode() opCode  { return opNOP }
func (CmdNOP) Name() string    { return "nop" }
func (CmdNOP) write(w *Writer) { w.op(opNOP) }
func (CmdNOP) read(r *Reader)  { _ = r.op() }

type CmdBOP struct {
	C0   int32 `json:"c0"`
	C1   int32 `json:"c1"`
	C2   int32 `json:"c2"`
	C3   int32 `json:"c3"`
	C4   int32 `json:"c4"`
	C5   int32 `json:"c5"`
	C6   int32 `json:"c6"`
	C7   int32 `json:"c7"`
	C8   int32 `json:"c8"`
	C9   int32 `json:"c9"`
	Prev int32 `json:"prev"`
}

func (CmdBOP) opcode() opCode { return opBOP }
func (CmdBOP) Name() string   { return "bop" }
func (c CmdBOP) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.C0))
	w.writeU32(uint32(c.C1))
	w.writeU32(uint32(c.C2))
	w.writeU32(uint32(c.C3))
	w.writeU32(uint32(c.C4))
	w.writeU32(uint32(c.C5))
	w.writeU32(uint32(c.C6))
	w.writeU32(uint32(c.C7))
	w.writeU32(uint32(c.C8))
	w.writeU32(uint32(c.C9))
	w.writeU32(uint32(c.Prev))
}

func (c *CmdBOP) read(r *Reader) {
	_ = r.op()
	c.C0 = int32(r.readU32())
	c.C1 = int32(r.readU32())
	c.C2 = int32(r.readU32())
	c.C3 = int32(r.readU32())
	c.C4 = int32(r.readU32())
	c.C5 = int32(r.readU32())
	c.C6 = int32(r.readU32())
	c.C7 = int32(r.readU32())
	c.C8 = int32(r.readU32())
	c.C9 = int32(r.readU32())
	c.Prev = int32(r.readU32())
}

type CmdEOP struct{}

func (CmdEOP) opcode() opCode    { return opEOP }
func (CmdEOP) Name() string      { return "eop" }
func (c CmdEOP) write(w *Writer) { w.op(c.opcode()) }
func (CmdEOP) read(r *Reader)    { _ = r.op() }

type CmdPush struct{}

func (CmdPush) opcode() opCode    { return opPush }
func (CmdPush) Name() string      { return "push" }
func (c CmdPush) write(w *Writer) { w.op(c.opcode()) }
func (CmdPush) read(r *Reader)    { _ = r.op() }

type CmdPop struct{}

func (CmdPop) opcode() opCode    { return opPop }
func (CmdPop) Name() string      { return "pop" }
func (c CmdPop) write(w *Writer) { w.op(c.opcode()) }
func (CmdPop) read(r *Reader)    { _ = r.op() }

type CmdRight1 struct {
	Value int32 `json:"v"`
}

func (CmdRight1) opcode() opCode { return opRight1 }
func (CmdRight1) Name() string   { return "right1" }
func (c CmdRight1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdRight1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdRight2 struct {
	Value int32 `json:"v"`
}

func (CmdRight2) opcode() opCode { return opRight2 }
func (CmdRight2) Name() string   { return "right2" }
func (c CmdRight2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdRight2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdRight3 struct {
	Value int32 `json:"v"`
}

func (CmdRight3) opcode() opCode { return opRight3 }
func (CmdRight3) Name() string   { return "right3" }
func (c CmdRight3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdRight3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdRight4 struct {
	Value int32 `json:"v"`
}

func (CmdRight4) opcode() opCode { return opRight4 }
func (CmdRight4) Name() string   { return "right4" }
func (c CmdRight4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdRight4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdW0 struct{}

func (CmdW0) opcode() opCode { return opW0 }
func (CmdW0) Name() string   { return "w0" }
func (c CmdW0) write(w *Writer) {
	w.op(c.opcode())
}
func (c *CmdW0) read(r *Reader) {
	_ = r.op()
}

type CmdW1 struct {
	Value int32 `json:"v"`
}

func (CmdW1) opcode() opCode { return opW1 }
func (CmdW1) Name() string   { return "w1" }
func (c CmdW1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdW1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdW2 struct {
	Value int32 `json:"v"`
}

func (CmdW2) opcode() opCode { return opW2 }
func (CmdW2) Name() string   { return "w2" }
func (c CmdW2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdW2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdW3 struct {
	Value int32 `json:"v"`
}

func (CmdW3) opcode() opCode { return opW3 }
func (CmdW3) Name() string   { return "w3" }
func (c CmdW3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdW3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdW4 struct {
	Value int32 `json:"v"`
}

func (CmdW4) opcode() opCode { return opW4 }
func (CmdW4) Name() string   { return "w4" }
func (c CmdW4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdW4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdX0 struct{}

func (CmdX0) opcode() opCode { return opX0 }
func (CmdX0) Name() string   { return "x0" }
func (c CmdX0) write(w *Writer) {
	w.op(c.opcode())
}
func (c *CmdX0) read(r *Reader) {
	_ = r.op()
}

type CmdX1 struct {
	Value int32 `json:"v"`
}

func (CmdX1) opcode() opCode { return opX1 }
func (CmdX1) Name() string   { return "x1" }
func (c CmdX1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdX1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdX2 struct {
	Value int32 `json:"v"`
}

func (CmdX2) opcode() opCode { return opX2 }
func (CmdX2) Name() string   { return "x2" }
func (c CmdX2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdX2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdX3 struct {
	Value int32 `json:"v"`
}

func (CmdX3) opcode() opCode { return opX3 }
func (CmdX3) Name() string   { return "x3" }
func (c CmdX3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdX3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdX4 struct {
	Value int32 `json:"v"`
}

func (CmdX4) opcode() opCode { return opX4 }
func (CmdX4) Name() string   { return "x4" }
func (c CmdX4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdX4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdY0 struct{}

func (CmdY0) opcode() opCode { return opY0 }
func (CmdY0) Name() string   { return "y0" }
func (c CmdY0) write(w *Writer) {
	w.op(c.opcode())
}
func (c *CmdY0) read(r *Reader) {
	_ = r.op()
}

type CmdY1 struct {
	Value int32 `json:"v"`
}

func (CmdY1) opcode() opCode { return opY1 }
func (CmdY1) Name() string   { return "y1" }
func (c CmdY1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdY1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdY2 struct {
	Value int32 `json:"v"`
}

func (CmdY2) opcode() opCode { return opY2 }
func (CmdY2) Name() string   { return "y2" }
func (c CmdY2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdY2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdY3 struct {
	Value int32 `json:"v"`
}

func (CmdY3) opcode() opCode { return opY3 }
func (CmdY3) Name() string   { return "y3" }
func (c CmdY3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdY3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdY4 struct {
	Value int32 `json:"v"`
}

func (CmdY4) opcode() opCode { return opY4 }
func (CmdY4) Name() string   { return "y4" }
func (c CmdY4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdY4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdZ0 struct{}

func (CmdZ0) opcode() opCode { return opZ0 }
func (CmdZ0) Name() string   { return "z0" }
func (c CmdZ0) write(w *Writer) {
	w.op(c.opcode())
}
func (c *CmdZ0) read(r *Reader) {
	_ = r.op()
}

type CmdZ1 struct {
	Value int32 `json:"v"`
}

func (CmdZ1) opcode() opCode { return opZ1 }
func (CmdZ1) Name() string   { return "z1" }
func (c CmdZ1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdZ1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdZ2 struct {
	Value int32 `json:"v"`
}

func (CmdZ2) opcode() opCode { return opZ2 }
func (CmdZ2) Name() string   { return "z2" }
func (c CmdZ2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdZ2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdZ3 struct {
	Value int32 `json:"v"`
}

func (CmdZ3) opcode() opCode { return opZ3 }
func (CmdZ3) Name() string   { return "z3" }
func (c CmdZ3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdZ3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdZ4 struct {
	Value int32 `json:"v"`
}

func (CmdZ4) opcode() opCode { return opZ4 }
func (CmdZ4) Name() string   { return "z4" }
func (c CmdZ4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdZ4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdDown1 struct {
	Value int32 `json:"v"`
}

func (CmdDown1) opcode() opCode { return opDown1 }
func (CmdDown1) Name() string   { return "down1" }
func (c CmdDown1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.Value))
}
func (c *CmdDown1) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU8())
}

type CmdDown2 struct {
	Value int32 `json:"v"`
}

func (CmdDown2) opcode() opCode { return opDown2 }
func (CmdDown2) Name() string   { return "down2" }
func (c CmdDown2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.Value))
}
func (c *CmdDown2) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU16())
}

type CmdDown3 struct {
	Value int32 `json:"v"`
}

func (CmdDown3) opcode() opCode { return opDown3 }
func (CmdDown3) Name() string   { return "down3" }
func (c CmdDown3) write(w *Writer) {
	w.op(c.opcode())
	w.writeI24(c.Value)
}
func (c *CmdDown3) read(r *Reader) {
	_ = r.op()
	c.Value = r.readI24()
}

type CmdDown4 struct {
	Value int32 `json:"v"`
}

func (CmdDown4) opcode() opCode { return opDown4 }
func (CmdDown4) Name() string   { return "down4" }
func (c CmdDown4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.Value))
}
func (c *CmdDown4) read(r *Reader) {
	_ = r.op()
	c.Value = int32(r.readU32())
}

type CmdFntNum struct {
	ID uint8 `json:"-"`
}

func (cmd CmdFntNum) opcode() opCode { return (opFntNum00 + opCode(cmd.ID)) }
func (cmd CmdFntNum) Name() string   { return fmt.Sprintf("fnt_num_%d", cmd.ID) }
func (c CmdFntNum) write(w *Writer) {
	w.op(c.opcode())
}
func (c *CmdFntNum) read(r *Reader) {
	op := r.op()
	c.ID = op - uint8(opFntNum00)
}

type CmdFnt1 struct {
	ID uint32
}

func (CmdFnt1) opcode() opCode { return opFnt1 }
func (CmdFnt1) Name() string   { return "fnt_1" }
func (c CmdFnt1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(c.ID))
}
func (c *CmdFnt1) read(r *Reader) {
	_ = r.op()
	c.ID = uint32(r.readU8())
}

type CmdFnt2 struct {
	ID uint32
}

func (CmdFnt2) opcode() opCode { return opFnt2 }
func (CmdFnt2) Name() string   { return "fnt_2" }
func (c CmdFnt2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(c.ID))
}
func (c *CmdFnt2) read(r *Reader) {
	_ = r.op()
	c.ID = uint32(r.readU16())
}

type CmdFnt3 struct {
	ID uint32
}

func (CmdFnt3) opcode() opCode { return opFnt3 }
func (CmdFnt3) Name() string   { return "fnt_3" }
func (c CmdFnt3) write(w *Writer) {
	w.op(c.opcode())
	w.writeU24(uint32(c.ID))
}
func (c *CmdFnt3) read(r *Reader) {
	_ = r.op()
	c.ID = uint32(r.readU24())
}

type CmdFnt4 struct {
	ID int32
}

func (CmdFnt4) opcode() opCode { return opFnt4 }
func (CmdFnt4) Name() string   { return "fnt_4" }
func (c CmdFnt4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.ID))
}
func (c *CmdFnt4) read(r *Reader) {
	_ = r.op()
	c.ID = int32(r.readU32())
}

type CmdXXX1 struct {
	Value []byte
}

func (CmdXXX1) opcode() opCode { return opXXX1 }
func (CmdXXX1) Name() string   { return "xxx1" }
func (c CmdXXX1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(uint8(len(c.Value)))
	w.writeBuf(c.Value)
}
func (c *CmdXXX1) read(r *Reader) {
	_ = r.op()
	n := int(r.readU8())
	c.Value = r.readBuf(n)
}

type CmdXXX2 struct {
	Value []byte
}

func (CmdXXX2) opcode() opCode { return opXXX2 }
func (CmdXXX2) Name() string   { return "xxx2" }
func (c CmdXXX2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(uint16(len(c.Value)))
	w.writeBuf(c.Value)
}
func (c *CmdXXX2) read(r *Reader) {
	_ = r.op()
	n := int(r.readU16())
	c.Value = r.readBuf(n)
}

type CmdXXX3 struct {
	Value []byte
}

func (CmdXXX3) opcode() opCode { return opXXX3 }
func (CmdXXX3) Name() string   { return "xxx3" }
func (c CmdXXX3) write(w *Writer) {
	w.op(c.opcode())
	w.writeU24(uint32(len(c.Value)))
	w.writeBuf(c.Value)
}
func (c *CmdXXX3) read(r *Reader) {
	_ = r.op()
	n := int(r.readU24())
	c.Value = r.readBuf(n)
}

type CmdXXX4 struct {
	Value []byte
}

func (CmdXXX4) opcode() opCode { return opXXX4 }
func (CmdXXX4) Name() string   { return "xxx4" }
func (c CmdXXX4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(len(c.Value)))
	w.writeBuf(c.Value)
}
func (c *CmdXXX4) read(r *Reader) {
	_ = r.op()
	n := int(r.readU32())
	c.Value = r.readBuf(n)
}

type CmdFntDef1 struct {
	ID       uint8  `json:"id"`
	Checksum uint32 `json:"chksum"`
	Size     int32  `json:"size"`
	Design   int32  `json:"dsize"`
	Area     string `json:"area"`
	Font     string `json:"name"`
}

func (CmdFntDef1) opcode() opCode { return opFntDef1 }
func (CmdFntDef1) Name() string   { return "fnt_def1" }
func (c CmdFntDef1) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(c.ID)
	w.writeU32(c.Checksum)
	w.writeU32(uint32(c.Size))
	w.writeU32(uint32(c.Design))
	w.writeU8(uint8(len(c.Area)))
	w.writeU8(uint8(len(c.Font)))
	w.writeBuf([]byte(c.Area))
	w.writeBuf([]byte(c.Font))
}
func (c *CmdFntDef1) read(r *Reader) {
	_ = r.op()
	c.ID = r.readU8()
	c.Checksum = r.readU32()
	c.Size = int32(r.readU32())
	c.Design = int32(r.readU32())
	area := r.readU8()
	font := r.readU8()
	buf := r.readBuf(int(area + font))
	c.Area = string(buf[:area])
	c.Font = string(buf[area:])
}

type CmdFntDef2 struct {
	ID       uint16 `json:"id"`
	Checksum uint32 `json:"chksum"`
	Size     int32  `json:"size"`
	Design   int32  `json:"dsize"`
	Area     string `json:"area"`
	Font     string `json:"name"`
}

func (CmdFntDef2) opcode() opCode { return opFntDef2 }
func (CmdFntDef2) Name() string   { return "fnt_def2" }
func (c CmdFntDef2) write(w *Writer) {
	w.op(c.opcode())
	w.writeU16(c.ID)
	w.writeU32(c.Checksum)
	w.writeU32(uint32(c.Size))
	w.writeU32(uint32(c.Design))
	w.writeU8(uint8(len(c.Area)))
	w.writeU8(uint8(len(c.Font)))
	w.writeBuf([]byte(c.Area))
	w.writeBuf([]byte(c.Font))
}
func (c *CmdFntDef2) read(r *Reader) {
	_ = r.op()
	c.ID = r.readU16()
	c.Checksum = r.readU32()
	c.Size = int32(r.readU32())
	c.Design = int32(r.readU32())
	area := r.readU8()
	font := r.readU8()
	buf := r.readBuf(int(area + font))
	c.Area = string(buf[:area])
	c.Font = string(buf[area:])
}

type CmdFntDef3 struct {
	ID       uint32 `json:"id"`
	Checksum uint32 `json:"chksum"`
	Size     int32  `json:"size"`
	Design   int32  `json:"dsize"`
	Area     string `json:"area"`
	Font     string `json:"name"`
}

func (CmdFntDef3) opcode() opCode { return opFntDef3 }
func (CmdFntDef3) Name() string   { return "fnt_def3" }
func (c CmdFntDef3) write(w *Writer) {
	w.op(c.opcode())
	w.writeU24(c.ID)
	w.writeU32(c.Checksum)
	w.writeU32(uint32(c.Size))
	w.writeU32(uint32(c.Design))
	w.writeU8(uint8(len(c.Area)))
	w.writeU8(uint8(len(c.Font)))
	w.writeBuf([]byte(c.Area))
	w.writeBuf([]byte(c.Font))
}
func (c *CmdFntDef3) read(r *Reader) {
	_ = r.op()
	c.ID = r.readU24()
	c.Checksum = r.readU32()
	c.Size = int32(r.readU32())
	c.Design = int32(r.readU32())
	area := r.readU8()
	font := r.readU8()
	buf := r.readBuf(int(area + font))
	c.Area = string(buf[:area])
	c.Font = string(buf[area:])
}

type CmdFntDef4 struct {
	ID       int32  `json:"id"`
	Checksum uint32 `json:"chksum"`
	Size     int32  `json:"size"`
	Design   int32  `json:"dsize"`
	Area     string `json:"area"`
	Font     string `json:"name"`
}

func (CmdFntDef4) opcode() opCode { return opFntDef4 }
func (CmdFntDef4) Name() string   { return "fnt_def4" }
func (c CmdFntDef4) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(uint32(c.ID))
	w.writeU32(c.Checksum)
	w.writeU32(uint32(c.Size))
	w.writeU32(uint32(c.Design))
	w.writeU8(uint8(len(c.Area)))
	w.writeU8(uint8(len(c.Font)))
	w.writeBuf([]byte(c.Area))
	w.writeBuf([]byte(c.Font))
}
func (c *CmdFntDef4) read(r *Reader) {
	_ = r.op()
	c.ID = int32(r.readU32())
	c.Checksum = r.readU32()
	c.Size = int32(r.readU32())
	c.Design = int32(r.readU32())
	area := r.readU8()
	font := r.readU8()
	buf := r.readBuf(int(area + font))
	c.Area = string(buf[:area])
	c.Font = string(buf[area:])
}

type CmdPre struct {
	Version uint8  `json:"version"`
	Num     int32  `json:"num"`
	Den     int32  `json:"den"`
	Mag     int32  `json:"mag"`
	Msg     string `json:"msg"`
}

func (cmd CmdPre) opcode() opCode { return opPre }
func (cmd CmdPre) Name() string   { return "pre" }
func (c CmdPre) write(w *Writer) {
	w.op(c.opcode())
	w.writeU8(c.Version)
	w.writeU32(uint32(c.Num))
	w.writeU32(uint32(c.Den))
	w.writeU32(uint32(c.Mag))
	w.writeU8(uint8(len(c.Msg)))
	w.writeBuf([]byte(c.Msg))
}
func (c *CmdPre) read(r *Reader) {
	_ = r.op()
	c.Version = r.readU8()
	c.Num = int32(r.readU32())
	c.Den = int32(r.readU32())
	c.Mag = int32(r.readU32())
	n := r.readU8()
	c.Msg = string(r.readBuf(int(n)))
}

type CmdPost struct {
	BOP      uint32 `json:"bop"`
	Num      int32  `json:"num"`
	Den      int32  `json:"den"`
	Mag      int32  `json:"mag"`
	Height   uint32 `json:"height`
	Width    uint32 `json:"width"`
	MaxStack uint16 `json:"max_stack"`
	Pages    uint16 `json:"pages"`
}

func (CmdPost) opcode() opCode { return opPost }
func (CmdPost) Name() string   { return "post" }
func (c CmdPost) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(c.BOP)
	w.writeU32(uint32(c.Num))
	w.writeU32(uint32(c.Den))
	w.writeU32(uint32(c.Mag))
	w.writeU32(c.Height)
	w.writeU32(c.Width)
	w.writeU16(c.MaxStack)
	w.writeU16(c.Pages)
}
func (c *CmdPost) read(r *Reader) {
	_ = r.op()
	c.BOP = r.readU32()
	c.Num = int32(r.readU32())
	c.Den = int32(r.readU32())
	c.Mag = int32(r.readU32())
	c.Height = r.readU32()
	c.Width = r.readU32()
	c.MaxStack = r.readU16()
	c.Pages = r.readU16()
}

type CmdPostPost struct {
	BOP     uint32 `json:"bop"`
	Version uint8  `json:"version"`
	Trailer uint8  `json:"trailer"`
}

func (CmdPostPost) opcode() opCode { return opPostPost }
func (CmdPostPost) Name() string   { return "post_post" }
func (c CmdPostPost) write(w *Writer) {
	w.op(c.opcode())
	w.writeU32(c.BOP)
	w.writeU8(c.Version)
	for i := 0; i < int(c.Trailer); i++ {
		w.writeU8(223)
	}
}

func (c *CmdPostPost) read(r *Reader) {
	_ = r.op()
	c.BOP = r.readU32()
	c.Version = r.readU8()
	c.Trailer = 0
	for {
		v, err := r.tailU8()
		if err != nil {
			break
		}
		if v != 223 {
			_ = r.r.UnreadByte()
			break
		}
		c.Trailer++
	}
}
