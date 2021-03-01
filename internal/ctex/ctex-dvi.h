// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef CTEX_CAPI_DVI_H
#define CTEX_CAPI_DVI_H 1

#ifdef __cplusplus
extern "C" {
#endif

#include "ctex-consts.h"
#include "ctex-types.h"

#include <stdio.h>

typedef int dvi_index;
typedef struct {
  integer total_pages;
  integer dvi_offset;
  integer dvi_gone; // the number of bytes already output to |dvi_file|

  scaled dvi_h; // current horizontal position
  scaled dvi_v; // current vertical position

  FILE *file;

  uint8_t buf[dvi_buf_size + 1];
  dvi_index half_buf;
  dvi_index dvi_limit;
  dvi_index dvi_ptr;

  internal_font_number dvi_f;
} ctex_dvi_t;

void ctex_dvi_init(ctex_dvi_t *self);

uint8_t ctex_dvi_at(ctex_dvi_t *self, int i);
void ctex_dvi_set(ctex_dvi_t *self, int i, uint8_t v);

FILE *ctex_dvi_file(ctex_dvi_t *self);
void ctex_dvi_set_file(ctex_dvi_t *self, FILE *f);
int ctex_dvi_fclose(ctex_dvi_t *self);

void ctex_dvi_add_page(ctex_dvi_t *self);

integer ctex_dvi_pages(ctex_dvi_t *self);

integer ctex_dvi_offset(ctex_dvi_t *self);

integer ctex_dvi_gone(ctex_dvi_t *self);

void ctex_dvi_flush(ctex_dvi_t *self);

void ctex_dvi_wU8(ctex_dvi_t *self, uint8_t v);

void ctex_dvi_four(ctex_dvi_t *self, integer x);

void ctex_dvi_pop(ctex_dvi_t *self, integer l);

integer ctex_dvi_pos(ctex_dvi_t *self);

integer ctex_dvi_cap(ctex_dvi_t *self);

void ctex_dvi_set_font(ctex_dvi_t *self, internal_font_number f);

internal_font_number ctex_dvi_get_font(ctex_dvi_t *self);

void ctex_dvi_set_h(ctex_dvi_t *self, scaled h);

void ctex_dvi_set_v(ctex_dvi_t *self, scaled v);

scaled ctex_dvi_get_h(ctex_dvi_t *self);

scaled ctex_dvi_get_v(ctex_dvi_t *self);

void ctex_dvi_font_def(ctex_dvi_t *self, int fid, uint32_t chksum, int32_t size,
                       int32_t dsize, size_t areasz, const char *area,
                       size_t namesz, const char *name);

void ctex_dvi_wcmd(ctex_dvi_t *self, uint8_t cmd, int32_t v);

// DVI commands
#define dvi_cmd_set1 128
#define dvi_cmd_set_rule 132
#define dvi_cmd_put_rule 137
#define dvi_cmd_bop 139
#define dvi_cmd_eop 140
#define dvi_cmd_push 141
#define dvi_cmd_pop 142
#define dvi_cmd_right1 143
#define dvi_cmd_down1 157
#define dvi_cmd_fnt_num(x) (170 + x)
#define dvi_cmd_xxx1 239
#define dvi_cmd_xxx4 242
#define dvi_cmd_fnt1 235
#define dvi_cmd_pre 247
#define dvi_cmd_post 248
#define dvi_cmd_post_post 249

#define dvi_version 2

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* CTEX_CAPI_DVI_H */
