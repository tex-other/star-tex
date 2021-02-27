// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dvi

import (
	"fmt"
	"io"
	"strings"
)

type Interpreter struct {
	w     io.Writer
	state state

	page int
	pre  CmdPre
	conv float32
}

func NewInterpreter(w io.Writer) *Interpreter {
	return &Interpreter{
		w:     w,
		state: newState(),
	}
}

func (in *Interpreter) Run(cmd Cmd) error {
	switch op := cmd.opcode(); op {
	case opSetChar000, opSetChar001, opSetChar002, opSetChar003, opSetChar004,
		opSetChar005, opSetChar006, opSetChar007, opSetChar008, opSetChar009,
		opSetChar010, opSetChar011, opSetChar012, opSetChar013, opSetChar014,
		opSetChar015, opSetChar016, opSetChar017, opSetChar018, opSetChar019,
		opSetChar020, opSetChar021, opSetChar022, opSetChar023, opSetChar024,
		opSetChar025, opSetChar026, opSetChar027, opSetChar028, opSetChar029,
		opSetChar030, opSetChar031, opSetChar032, opSetChar033, opSetChar034,
		opSetChar035, opSetChar036, opSetChar037, opSetChar038, opSetChar039,
		opSetChar040, opSetChar041, opSetChar042, opSetChar043, opSetChar044,
		opSetChar045, opSetChar046, opSetChar047, opSetChar048, opSetChar049,
		opSetChar050, opSetChar051, opSetChar052, opSetChar053, opSetChar054,
		opSetChar055, opSetChar056, opSetChar057, opSetChar058, opSetChar059,
		opSetChar060, opSetChar061, opSetChar062, opSetChar063, opSetChar064,
		opSetChar065, opSetChar066, opSetChar067, opSetChar068, opSetChar069,
		opSetChar070, opSetChar071, opSetChar072, opSetChar073, opSetChar074,
		opSetChar075, opSetChar076, opSetChar077, opSetChar078, opSetChar079,
		opSetChar080, opSetChar081, opSetChar082, opSetChar083, opSetChar084,
		opSetChar085, opSetChar086, opSetChar087, opSetChar088, opSetChar089,
		opSetChar090, opSetChar091, opSetChar092, opSetChar093, opSetChar094,
		opSetChar095, opSetChar096, opSetChar097, opSetChar098, opSetChar099,
		opSetChar100, opSetChar101, opSetChar102, opSetChar103, opSetChar104,
		opSetChar105, opSetChar106, opSetChar107, opSetChar108, opSetChar109,
		opSetChar110, opSetChar111, opSetChar112, opSetChar113, opSetChar114,
		opSetChar115, opSetChar116, opSetChar117, opSetChar118, opSetChar119,
		opSetChar120, opSetChar121, opSetChar122, opSetChar123, opSetChar124,
		opSetChar125, opSetChar126, opSetChar127:
		cmd := cmd.(*CmdSetChar)
		cur := in.state.cur()
		old := cur.h
		cur.h += int32(cmd.Value)
		fmt.Fprintf(in.w,
			"%s h:=%d%+d=%d, hh:=%d\n",
			strings.Replace(cmd.Name(), "_", "", -1),
			old, cmd.Value, cur.h, in.pixels(cur.h),
		)

	case opBOP:
		in.page++
		fmt.Fprintf(in.w, "beginning of page %d\n", in.page)

	case opPush:
		fmt.Fprintf(in.w, "push\n")
		lvl := len(in.state.stack) - 1
		in.state.push()
		st := in.state.cur()
		fmt.Fprintf(in.w,
			"level %d:(h=%d,v=%d,w=%d,x=%d,y=%d,z=%d,hh=%d,vv=%d)\n",
			lvl,
			st.h, st.v, st.w, st.x, st.y, st.z,
			in.pixels(st.h), in.pixels(st.v),
		)

	case opPop:
		fmt.Fprintf(in.w, "pop\n")
		in.state.pop()
		lvl := len(in.state.stack) - 1
		st := in.state.cur()
		fmt.Fprintf(in.w,
			"level %d:(h=%d,v=%d,w=%d,x=%d,y=%d,z=%d,hh=%d,vv=%d)\n",
			lvl,
			st.h, st.v, st.w, st.x, st.y, st.z,
			in.pixels(st.h), in.pixels(st.v),
		)

	case opRight1:
		cmd := cmd.(*CmdRight1)
		cur := in.state.cur()
		old := cur.h
		cur.h += cmd.Value
		fmt.Fprintf(in.w,
			"right1 %d h:=%d%+d=%+d, hh:=%d\n",
			cmd.Value, old, cmd.Value, cur.h, in.pixels(cur.h),
		)

	case opRight2:
		cmd := cmd.(*CmdRight2)
		cur := in.state.cur()
		old := cur.h
		cur.h += cmd.Value
		fmt.Fprintf(in.w,
			"right2 %d h:=%d%+d=%+d, hh:=%d\n",
			cmd.Value, old, cmd.Value, cur.h, in.pixels(cur.h),
		)

	case opRight3:
		cmd := cmd.(*CmdRight3)
		cur := in.state.cur()
		old := cur.h
		cur.h += cmd.Value
		fmt.Fprintf(in.w,
			"right3 %d h:=%d%+d=%+d, hh:=%d\n",
			cmd.Value, old, cmd.Value, cur.h, in.pixels(cur.h),
		)

	case opRight4:
		cmd := cmd.(*CmdRight4)
		cur := in.state.cur()
		old := cur.h
		cur.h += cmd.Value
		fmt.Fprintf(in.w,
			"right4 %d h:=%d%+d=%+d, hh:=%d\n",
			cmd.Value, old, cmd.Value, cur.h, in.pixels(cur.h),
		)

	case opDown1:
		cmd := cmd.(*CmdDown1)
		cur := in.state.cur()
		old := cur.v
		cur.v += cmd.Value
		fmt.Fprintf(in.w,
			"down1 %d v:=%d%+d=%+d, vv:=%d\n",
			cmd.Value, old, cmd.Value, cur.v, in.pixels(cur.v),
		)

	case opDown2:
		cmd := cmd.(*CmdDown2)
		cur := in.state.cur()
		old := cur.v
		cur.v += cmd.Value
		fmt.Fprintf(in.w,
			"down2 %d v:=%d%+d=%+d, vv:=%d\n",
			cmd.Value, old, cmd.Value, cur.v, in.pixels(cur.v),
		)

	case opDown3:
		cmd := cmd.(*CmdDown3)
		cur := in.state.cur()
		old := cur.v
		cur.v += cmd.Value
		fmt.Fprintf(in.w,
			"down3 %d v:=%d%+d=%+d, vv:=%d\n",
			cmd.Value, old, cmd.Value, cur.v, in.pixels(cur.v),
		)

	case opDown4:
		cmd := cmd.(*CmdDown4)
		cur := in.state.cur()
		old := cur.v
		cur.v += cmd.Value
		fmt.Fprintf(in.w,
			"down4 %d v:=%d%+d=%+d, vv:=%d\n",
			cmd.Value, old, cmd.Value, cur.v, in.pixels(cur.v),
		)

	case opFntNum00, opFntNum01, opFntNum02, opFntNum03, opFntNum04,
		opFntNum05, opFntNum06, opFntNum07, opFntNum08, opFntNum09,
		opFntNum10, opFntNum11, opFntNum12, opFntNum13, opFntNum14,
		opFntNum15, opFntNum16, opFntNum17, opFntNum18, opFntNum19,
		opFntNum20, opFntNum21, opFntNum22, opFntNum23, opFntNum24,
		opFntNum25, opFntNum26, opFntNum27, opFntNum28, opFntNum29,
		opFntNum30, opFntNum31, opFntNum32, opFntNum33, opFntNum34,
		opFntNum35, opFntNum36, opFntNum37, opFntNum38, opFntNum39,
		opFntNum40, opFntNum41, opFntNum42, opFntNum43, opFntNum44,
		opFntNum45, opFntNum46, opFntNum47, opFntNum48, opFntNum49,
		opFntNum50, opFntNum51, opFntNum52, opFntNum53, opFntNum54,
		opFntNum55, opFntNum56, opFntNum57, opFntNum58, opFntNum59,
		opFntNum60, opFntNum61, opFntNum62, opFntNum63:
		cmd := cmd.(*CmdFntNum)
		in.state.f = int(cmd.ID)
		fmt.Fprintf(in.w, "%v current font is %s\n",
			strings.Replace(cmd.Name(), "_", "", -1),
			in.state.fonts[in.state.f].Font,
		)

	case opFntDef1:
		cmd := cmd.(*CmdFntDef1)
		fmt.Fprintf(in.w, "fntdef1 %d: %s\n", cmd.ID, cmd.Font)
		in.state.fonts[int(cmd.ID)] = fntdef{
			ID:       int(cmd.ID),
			Checksum: cmd.Checksum,
			Size:     cmd.Size,
			Design:   cmd.Design,
			Area:     cmd.Area,
			Font:     cmd.Font,
		}

	case opFntDef2:
		cmd := cmd.(*CmdFntDef2)
		fmt.Fprintf(in.w, "fntdef2 %d: %s\n", cmd.ID, cmd.Font)
		in.state.fonts[int(cmd.ID)] = fntdef{
			ID:       int(cmd.ID),
			Checksum: cmd.Checksum,
			Size:     cmd.Size,
			Design:   cmd.Design,
			Area:     cmd.Area,
			Font:     cmd.Font,
		}

	case opFntDef3:
		cmd := cmd.(*CmdFntDef3)
		fmt.Fprintf(in.w, "fntdef3 %d: %s\n", cmd.ID, cmd.Font)
		in.state.fonts[int(cmd.ID)] = fntdef{
			ID:       int(cmd.ID),
			Checksum: cmd.Checksum,
			Size:     cmd.Size,
			Design:   cmd.Design,
			Area:     cmd.Area,
			Font:     cmd.Font,
		}

	case opFntDef4:
		cmd := cmd.(*CmdFntDef4)
		fmt.Fprintf(in.w, "fntdef4 %d: %s\n", cmd.ID, cmd.Font)
		in.state.fonts[int(cmd.ID)] = fntdef{
			ID:       int(cmd.ID),
			Checksum: cmd.Checksum,
			Size:     cmd.Size,
			Design:   cmd.Design,
			Area:     cmd.Area,
			Font:     cmd.Font,
		}

	case opPre:
		in.pre = *cmd.(*CmdPre)
		fmt.Fprintf(in.w, "numerator/denominator=%d/%d\n", in.pre.Num, in.pre.Den)
		res := float32(300.0)
		conv := float32(in.pre.Num) / 254000.0 * (res / float32(in.pre.Den))
		in.conv = conv * float32(in.pre.Mag) / 1000.0
		fmt.Fprintf(in.w, "magnification=%d;       %10.8f pixels per DVI unit\n", in.pre.Mag, in.conv)
		fmt.Fprintf(in.w, "'%s'\n", in.pre.Msg)
	default:
		panic(fmt.Errorf("unknown cmd %T", cmd))
	}

	return nil
}

func (in *Interpreter) pixels(v int32) int32 {
	x := in.conv * float32(v)
	return roundF32(x)
}

func roundF32(v float32) int32 {
	if v > 0 {
		return int32(v + 0.5)
	}
	return int32(v - 0.5)
}
