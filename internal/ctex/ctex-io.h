// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef CTEX_IO_H
#define CTEX_IO_H 1

#include <stdint.h>
#include <stdio.h>

#include "ctex-types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef int ctex_file;
ctex_file ctex_io_fopen(const char *name);
ctex_file ctex_io_fcreate(const char *name);
int ctex_io_fclose(ctex_file f);
int ctex_io_fflush(ctex_file f);
int ctex_io_feof(ctex_file f);
int ctex_io_ferror(ctex_file f);
int ctex_io_fgetc(ctex_file f);
int ctex_io_fpeek(ctex_file f);
uint8_t ctex_io_readU8(ctex_file f);
uint32_t ctex_io_readU32(ctex_file f);
void ctex_io_writeU32(ctex_file f, uint32_t v);

void ctex_io_fprintf(ctex_file f, const char *str);
void ctex_io_fprintfU8(ctex_file f, const char *fmt, uint8_t v);

void loadU8(ctex_file r, int *mode, uint8_t *v);
uint8_t *readU8(ctex_file r, int *mode, uint8_t *v);
void loadU32(ctex_file r, int *mode, memory_word *v);
memory_word *readU32(ctex_file r, int *mode, memory_word *v);
void writeU32(ctex_file w, int *mode, memory_word *v);

int erstat(ctex_file f);
int fpeek(ctex_file f);
void break_in(ctex_file ios, bool_t);
bool_t eoln(ctex_file f);

const char *trim_name(char *filename,
                      size_t length); // never called on a string literal;
                                      // note the lack of a const
void io_error(int error, const char *name);

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* CTEX_IO_H */
