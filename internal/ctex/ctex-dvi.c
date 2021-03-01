// Copyright ©2021 The star-tex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "ctex-dvi.h"

#include <stdio.h>
#include <stdlib.h>

void ctex_dvi_swap(ctex_dvi_t *self);
void ctex_dvi_write_dvi(ctex_dvi_t *self, dvi_index a, dvi_index b);

void ctex_dvi_init(ctex_dvi_t *self) {
  self->total_pages = 0;
  self->dvi_offset = 0;
  self->dvi_gone = 0;

  self->dvi_h = 0;
  self->dvi_v = 0;

  self->file = NULL;

  for (int k = 0; k < dvi_buf_size + 1; k++) {
    self->buf[k] = 0;
  }
  self->half_buf = dvi_buf_size / 2;
  self->dvi_limit = dvi_buf_size;
  self->dvi_ptr = 0;
  self->dvi_f = 0;
}

uint8_t ctex_dvi_at(ctex_dvi_t *self, int i) { return self->buf[i]; }

void ctex_dvi_set(ctex_dvi_t *self, int i, uint8_t v) { self->buf[i] = v; }

FILE *ctex_dvi_file(ctex_dvi_t *self) { return self->file; }

void ctex_dvi_set_file(ctex_dvi_t *self, FILE *f) { self->file = f; }

int ctex_dvi_fclose(ctex_dvi_t *self) {
  int rc = 0;
  if (self->file) {
    rc = fclose(self->file);
    self->file = NULL;
  }
  return rc;
}

void ctex_dvi_add_page(ctex_dvi_t *self) { self->total_pages++; }

integer ctex_dvi_pages(ctex_dvi_t *self) { return self->total_pages; }

integer ctex_dvi_offset(ctex_dvi_t *self) { return self->dvi_offset; }

integer ctex_dvi_gone(ctex_dvi_t *self) { return self->dvi_gone; }

void ctex_dvi_flush(ctex_dvi_t *self) {
  if (self->dvi_limit == self->half_buf) {
    ctex_dvi_write_dvi(self, self->half_buf, dvi_buf_size - 1);
  }
  if (self->dvi_ptr > 0) {
    ctex_dvi_write_dvi(self, 0, self->dvi_ptr - 1);
  }
}

void ctex_dvi_wU8(ctex_dvi_t *self, uint8_t v) {
  self->buf[self->dvi_ptr] = v;
  ++self->dvi_ptr;
  if (self->dvi_ptr == self->dvi_limit) {
    ctex_dvi_swap(self);
  }
}

void ctex_dvi_swap(ctex_dvi_t *self) {
  if (self->dvi_limit == dvi_buf_size) {
    ctex_dvi_write_dvi(self, 0, self->half_buf - 1);
    self->dvi_limit = self->half_buf;
    self->dvi_offset += dvi_buf_size;
    self->dvi_ptr = 0;
  } else {
    ctex_dvi_write_dvi(self, self->half_buf, dvi_buf_size - 1);
    self->dvi_limit = dvi_buf_size;
  }
  self->dvi_gone += self->half_buf;
}

void ctex_dvi_write_dvi(ctex_dvi_t *self, dvi_index a, dvi_index b) {
  dvi_index k;
  for (k = a; k < b + 1; ++k) {
    fprintf(self->file, "%c", self->buf[k]);
  }
}

void ctex_dvi_four(ctex_dvi_t *self, integer x) {
  ctex_dvi_wU8(self, x >> 24);
  ctex_dvi_wU8(self, x >> 16);
  ctex_dvi_wU8(self, x >> 8);
  ctex_dvi_wU8(self, x);
}

void ctex_dvi_pop(ctex_dvi_t *self, integer l) {
  if ((l == (self->dvi_offset + self->dvi_ptr)) && (self->dvi_ptr > 0)) {
    --self->dvi_ptr;
    return;
  }
  ctex_dvi_wU8(self, dvi_cmd_pop);
}

integer ctex_dvi_pos(ctex_dvi_t *self) {
  return self->dvi_offset + self->dvi_ptr;
}

integer ctex_dvi_cap(ctex_dvi_t *self) { return dvi_buf_size - self->dvi_ptr; }

void ctex_dvi_set_font(ctex_dvi_t *self, internal_font_number f) {
  self->dvi_f = f;
}

internal_font_number ctex_dvi_get_font(ctex_dvi_t *self) { return self->dvi_f; }

void ctex_dvi_set_h(ctex_dvi_t *self, scaled h) { self->dvi_h = h; }

void ctex_dvi_set_v(ctex_dvi_t *self, scaled v) { self->dvi_v = v; }

scaled ctex_dvi_get_h(ctex_dvi_t *self) { return self->dvi_h; }

scaled ctex_dvi_get_v(ctex_dvi_t *self) { return self->dvi_v; }

void ctex_dvi_font_def(ctex_dvi_t *self, int fid, uint32_t chksum, int32_t size,
                       int32_t dsize, size_t areasz, const char *area,
                       size_t namesz, const char *name) {

  ctex_dvi_wU8(self, 243);
  ctex_dvi_wU8(self, fid);
  ctex_dvi_wU8(self, chksum >> 24);
  ctex_dvi_wU8(self, chksum >> 16);
  ctex_dvi_wU8(self, chksum >> 8);
  ctex_dvi_wU8(self, chksum);
  ctex_dvi_four(self, size);
  ctex_dvi_four(self, dsize);
  ctex_dvi_wU8(self, areasz);
  ctex_dvi_wU8(self, namesz);
  for (int i = 0; i < areasz; i++) {
    ctex_dvi_wU8(self, area[i]);
  }
  for (int i = 0; i < namesz; i++) {
    ctex_dvi_wU8(self, name[i]);
  }
}

void ctex_dvi_wcmd(ctex_dvi_t *self, uint8_t cmd, int32_t v) {
  uint32_t u = abs(v);

  if (u >= (1 << 23)) {
    ctex_dvi_wU8(self, cmd + 3);
    ctex_dvi_four(self, v);
    return;
  }

  if (u >= (1 << 15)) {
    ctex_dvi_wU8(self, cmd + 2);
    ctex_dvi_wU8(self, v >> 16);
    ctex_dvi_wU8(self, v >> 8);
    ctex_dvi_wU8(self, v);
    return;
  }

  if (u >= (1 << 7)) {
    ctex_dvi_wU8(self, cmd + 1);
    ctex_dvi_wU8(self, v >> 8);
    ctex_dvi_wU8(self, v);
    return;
  }

  ctex_dvi_wU8(self, cmd);
  ctex_dvi_wU8(self, v);
}
