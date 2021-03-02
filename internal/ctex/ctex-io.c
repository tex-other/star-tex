// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "ctex-io.h"

void loadU8(ctex_file r, int *mode, uint8_t *v) {
  if (*mode == 1) {
    *v = ctex_io_readU8(r);
  } else {
    *mode = 1;
  }
}

uint8_t *readU8(ctex_file r, int *mode, uint8_t *v) {
  if (*mode == 1) {
    *mode = 2;
    *v = ctex_io_readU8(r);
  }
  return v;
}

void loadU32(ctex_file r, int *mode, memory_word *v) {
  if (*mode == 1) {
    v->int_ = ctex_io_readU32(r);
  } else {
    *mode = 1;
  }
}

memory_word *readU32(ctex_file r, int *mode, memory_word *v) {
  if (*mode == 1) {
    *mode = 2;
    v->int_ = ctex_io_readU32(r);
  }
  return v;
}

void writeU32(ctex_file w, int *mode, memory_word *v) {
  ctex_io_writeU32(w, v->int_);
  *mode = 0;
}

int erstat(ctex_file f) { return (f == 0) || (ctex_io_ferror(f) != 0); }

int fpeek(ctex_file f) { return ctex_io_fpeek(f); }

void break_in(ctex_file ios, bool_t v) {}

bool_t eoln(ctex_file f) {
  int c = ctex_io_fpeek(f);
  return (c == EOF) || (c == '\n');
}

const char *trim_name(char *filename,
                      size_t length) // never called on a string literal;
                                     // note the lack of a const
{
  for (char *p = filename + length - 1; *p == ' '; --p)
    *p = '\0';

  return filename;
}

void io_error(int error, const char *name) {}
