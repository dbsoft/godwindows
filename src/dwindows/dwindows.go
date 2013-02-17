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
type HTREEITEM unsafe.Pointer
type DW struct { }

const (
   FALSE C.int = iota
   TRUE
)

var HWND_DESKTOP HWND = nil
var DW_DESKTOP HWND = nil

func (dw DW) Init(newthread C.int) C.int {
   return C.go_init(newthread);
}

func (dw DW) Messagebox(title string, flags C.int, message string) C.int {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cmessage := C.CString(message);
   defer C.free(unsafe.Pointer(cmessage));
   
   return C.go_messagebox(ctitle, flags, cmessage);
}

func (dw DW) Window_new(owner HWND, title string, flags C.ulong) HWND {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   
   return HWND(C.go_window_new(unsafe.Pointer(owner), ctitle, flags));
}

func (dw DW) Window_show(handle HWND) C.int {
   return C.go_window_show(unsafe.Pointer(handle));
}

func (dw DW) Window_set_pos(handle HWND, x C.long, y C.long) {
   C.go_window_set_pos(unsafe.Pointer(handle), x, y);
}

func (dw DW) Window_set_pos_size(handle HWND, x C.long, y C.long, width C.ulong, height C.ulong) {
   C.go_window_set_pos_size(unsafe.Pointer(handle), x, y, width, height);
}

func (dw DW) Window_set_size(handle HWND, width C.ulong, height C.ulong) {
   C.go_window_set_size(unsafe.Pointer(handle), width, height);
}

func (dw DW) Window_set_color(handle HWND, fore C.ulong, back C.ulong) C.int {
   return C.go_window_set_color(unsafe.Pointer(handle), fore, back);
}

func (dw DW) Window_set_style(handle HWND, style C.ulong, mask C.ulong) {
   C.go_window_set_style(unsafe.Pointer(handle), style, mask);
}

func (dw DW) Window_click_default(window HWND, next HWND) {
   C.go_window_click_default(unsafe.Pointer(window), unsafe.Pointer(next));
}

func (dw DW) Window_default(window HWND, defaultitem HWND) {
   C.go_window_default(unsafe.Pointer(window), unsafe.Pointer(defaultitem));
}

func (dw DW) Main() {
   C.dw_main();
}

func (dw DW) Main_iteration() {
   C.dw_main_iteration();
}

func (dw DW) Main_quit() {
   C.dw_main_quit();
}

func (dw DW) Main_sleep(milliseconds C.int) {
   C.dw_main_sleep(milliseconds);
}

func (dw DW) Box_new(btype C.int, pad C.int) HWND {
   return HWND(C.go_box_new(btype, pad));
}

func (dw DW) Box_pack_at_index(box HWND, item HWND, index C.int, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_at_index(unsafe.Pointer(box), unsafe.Pointer(item), index, width, height, hsize, vsize, pad);
}

func (dw DW) Box_pack_end(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_end(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func (dw DW) Box_pack_start(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_start(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func (dw DW) Box_unpack(handle HWND) C.int {
   return C.go_box_unpack(unsafe.Pointer(handle));
}

func (dw DW) Box_unpack_at_index(handle HWND, index C.int) HWND {
   return HWND(C.go_box_unpack_at_index(unsafe.Pointer(handle), index));
}

func (dw DW) Text_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_text_new(ctext, id));
}

func (dw DW) Entryfield_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_new(ctext, id));
}

func (dw DW) Entryfield_password_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_password_new(ctext, id));
}

func (dw DW) Entryfield_set_limit(handle HWND, limit C.int) {
   C.go_entryfield_set_limit(unsafe.Pointer(handle), limit);
}

func (dw DW) Button_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_button_new(ctext, id));
}

func (dw DW) Signal_connect(window HWND, signame string, sigfunc unsafe.Pointer, data unsafe.Pointer) {
   csigname := C.CString(signame);
   defer C.free(unsafe.Pointer(csigname));
   
   C.go_signal_connect(unsafe.Pointer(window), csigname, sigfunc, data);
}

func init() {
   runtime.LockOSThread();
}

//export go_int_callback_basic
func go_int_callback_basic(pfunc unsafe.Pointer, window unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), data);
}

//export go_int_callback_configure
func go_int_callback_configure(pfunc unsafe.Pointer, window unsafe.Pointer, width C.int, height C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, C.int, C.int, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), width, height, data);
}

//export go_int_callback_keypress
func go_int_callback_keypress(pfunc unsafe.Pointer, window unsafe.Pointer, ch C.char, vk C.int, state C.int, data unsafe.Pointer, utf8 *C.char) C.int {
   thisfunc := *(*func(HWND, C.char, C.int, C.int, unsafe.Pointer, string) C.int)(pfunc);
   return thisfunc(HWND(window), ch, vk, state, data, C.GoString(utf8));
}

//export go_int_callback_mouse
func go_int_callback_mouse(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, mask C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, C.int, C.int, C.int, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), x, y, mask, data);
}

//export go_int_callback_expose
func go_int_callback_expose(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, width C.int, height C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, C.int, C.int, C.int, C.int, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), x, y, width, height, data);
}

//export go_int_callback_string
func go_int_callback_string(pfunc unsafe.Pointer, window unsafe.Pointer, str *C.char, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, string, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), C.GoString(str), data);
}

//export go_int_callback_item_context
func go_int_callback_item_context(pfunc unsafe.Pointer, window unsafe.Pointer, text *C.char, x C.int, y C.int, data unsafe.Pointer, itemdata unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, string, C.int, C.int, unsafe.Pointer, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), C.GoString(text), x, y, data, itemdata);
}

//export go_int_callback_item_select
func go_int_callback_item_select(pfunc unsafe.Pointer, window unsafe.Pointer, item unsafe.Pointer, text *C.char, data unsafe.Pointer, itemdata unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, HTREEITEM, string, unsafe.Pointer, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), HTREEITEM(item), C.GoString(text), data, itemdata);
}

//export go_int_callback_numeric
func go_int_callback_numeric(pfunc unsafe.Pointer, window unsafe.Pointer, val C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, C.int, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), val, data);
}

//export go_int_callback_ulong
func go_int_callback_ulong(pfunc unsafe.Pointer, window unsafe.Pointer, val C.ulong, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, C.ulong, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), val, data);
}

//export go_int_callback_tree
func go_int_callback_tree(pfunc unsafe.Pointer, window unsafe.Pointer, tree unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, HTREEITEM, unsafe.Pointer) C.int)(pfunc);
   return thisfunc(HWND(window), HTREEITEM(tree), data);
}

