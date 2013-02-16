package dwindows

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/Netlabs/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/Netlabs/dwindows -ldw
#include "dwglue.c"
*/
import "C"
import "unsafe"
import "runtime"

type HWND unsafe.Pointer
type DW struct { }

const (
   FALSE C.int = iota
   TRUE
)

var HWND_DESKTOP HWND = nil
var DW_DESKTOP HWND = nil

func (dw DW) init(newthread C.int) C.int {
   return C.go_init(newthread);
}

func (dw DW) messagebox(title string, flags C.int, message string) C.int {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cmessage := C.CString(message);
   defer C.free(unsafe.Pointer(cmessage));
   
   return C.go_messagebox(ctitle, flags, cmessage);
}

func (dw DW) window_new(owner HWND, title string, flags C.ulong) HWND {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   
   return HWND(C.go_window_new(unsafe.Pointer(owner), ctitle, flags));
}

func (dw DW) window_show(handle HWND) C.int {
   return C.go_window_show(unsafe.Pointer(handle));
}

func (dw DW) window_set_pos(handle HWND, x C.long, y C.long) {
   C.go_window_set_pos(unsafe.Pointer(handle), x, y);
}

func (dw DW) window_set_pos_size(handle HWND, x C.long, y C.long, width C.ulong, height C.ulong) {
   C.go_window_set_pos_size(unsafe.Pointer(handle), x, y, width, height);
}

func (dw DW) window_set_size(handle HWND, width C.ulong, height C.ulong) {
   C.go_window_set_size(unsafe.Pointer(handle), width, height);
}

func (dw DW) window_set_color(handle HWND, fore C.ulong, back C.ulong) C.int {
   return C.go_window_set_color(unsafe.Pointer(handle), fore, back);
}

func (dw DW) window_set_style(handle HWND, style C.ulong, mask C.ulong) {
   C.go_window_set_style(unsafe.Pointer(handle), style, mask);
}

func (dw DW) window_click_default(window HWND, next HWND) {
   C.go_window_click_default(unsafe.Pointer(window), unsafe.Pointer(next));
}

func (dw DW) window_default(window HWND, defaultitem HWND) {
   C.go_window_default(unsafe.Pointer(window), unsafe.Pointer(defaultitem));
}

func (dw DW) main() {
   C.dw_main();
}

func (dw DW) main_iteration() {
   C.dw_main_iteration();
}

func (dw DW) main_quit() {
   C.dw_main_quit();
}

func (dw DW) main_sleep(milliseconds C.int) {
   C.dw_main_sleep(milliseconds);
}

func (dw DW) box_new(btype C.int, pad C.int) HWND {
   return HWND(C.go_box_new(btype, pad));
}

func (dw DW) box_pack_at_index(box HWND, item HWND, index C.int, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_at_index(unsafe.Pointer(box), unsafe.Pointer(item), index, width, height, hsize, vsize, pad);
}

func (dw DW) box_pack_end(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_end(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func (dw DW) box_pack_start(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_start(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func (dw DW) box_unpack(handle HWND) C.int {
   return C.go_box_unpack(unsafe.Pointer(handle));
}

func (dw DW) box_unpack_at_index(handle HWND, index C.int) HWND {
   return HWND(C.go_box_unpack_at_index(unsafe.Pointer(handle), index));
}

func (dw DW) text_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_text_new(ctext, id));
}

func (dw DW) entryfield_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_new(ctext, id));
}

func (dw DW) entryfield_password_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_password_new(ctext, id));
}

func (dw DW) entryfield_set_limit(handle HWND, limit C.int) {
   C.go_entryfield_set_limit(unsafe.Pointer(handle), limit);
}

func (dw DW) button_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_button_new(ctext, id));
}

func (dw DW) signal_connect(window HWND, signame string, sigfunc unsafe.Pointer, data unsafe.Pointer) {
   csigname := C.CString(signame);
   defer C.free(unsafe.Pointer(csigname));
   
   C.go_signal_connect(unsafe.Pointer(window), csigname, sigfunc, data);
}

func init() {
   runtime.LockOSThread();
}

// export go_int_callback_basic
func go_int_callback_basic(pfunc unsafe.Pointer, window unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), data);
}

