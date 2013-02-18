package dw

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
type HICN unsafe.Pointer
type COLOR C.ulong

const (
   FALSE C.int = iota
   TRUE
)

var DESKTOP HWND = nil

func RESOURCE(id uintptr) unsafe.Pointer {
   return unsafe.Pointer(id);
}

func Init(newthread C.int) C.int {
   return C.go_init(newthread);
}

func Messagebox(title string, flags C.int, message string) C.int {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cmessage := C.CString(message);
   defer C.free(unsafe.Pointer(cmessage));
   
   return C.go_messagebox(ctitle, flags, cmessage);
}

func Window_new(owner HWND, title string, flags C.ulong) HWND {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   
   return HWND(C.go_window_new(unsafe.Pointer(owner), ctitle, flags));
}

func Window_show(handle HWND) C.int {
   return C.go_window_show(unsafe.Pointer(handle));
}

func Window_hide(handle HWND) C.int {
   return C.go_window_hide(unsafe.Pointer(handle));
}

func Window_lower(handle HWND) C.int {
   return C.go_window_lower(unsafe.Pointer(handle));
}

func Window_raise(handle HWND) C.int {
   return C.go_window_raise(unsafe.Pointer(handle));
}

func Window_minimize(handle HWND) C.int {
   return C.go_window_minimize(unsafe.Pointer(handle));
}

func Window_set_pos(handle HWND, x C.long, y C.long) {
   C.go_window_set_pos(unsafe.Pointer(handle), x, y);
}

func Window_set_pos_size(handle HWND, x C.long, y C.long, width C.ulong, height C.ulong) {
   C.go_window_set_pos_size(unsafe.Pointer(handle), x, y, width, height);
}

func Window_set_size(handle HWND, width C.ulong, height C.ulong) {
   C.go_window_set_size(unsafe.Pointer(handle), width, height);
}

func Window_set_color(handle HWND, fore COLOR, back COLOR) C.int {
   return C.go_window_set_color(unsafe.Pointer(handle), C.ulong(fore), C.ulong(back));
}

func Window_set_style(handle HWND, style C.ulong, mask C.ulong) {
   C.go_window_set_style(unsafe.Pointer(handle), style, mask);
}

func Window_click_default(window HWND, next HWND) {
   C.go_window_click_default(unsafe.Pointer(window), unsafe.Pointer(next));
}

func Window_default(window HWND, defaultitem HWND) {
   C.go_window_default(unsafe.Pointer(window), unsafe.Pointer(defaultitem));
}

func Window_destroy(handle HWND) C.int {
   return C.go_window_destroy(unsafe.Pointer(handle));
}

func Window_disable(handle HWND) {
   C.go_window_disable(unsafe.Pointer(handle));
}

func Window_enable(handle HWND) {
   C.go_window_enable(unsafe.Pointer(handle));
}

func Window_from_id(handle HWND, cid C.int) HWND {
   return HWND(C.go_window_from_id(unsafe.Pointer(handle), cid));
}

func Window_get_data(window HWND, dataname string) unsafe.Pointer {
   cdataname := C.CString(dataname);
   defer C.free(unsafe.Pointer(cdataname));
   
   return C.go_window_get_data(unsafe.Pointer(window), cdataname);
}

func Window_get_font(handle HWND) string {
   cfontname := C.go_window_get_font(unsafe.Pointer(handle));
   fontname := C.GoString(cfontname);
   C.dw_free(unsafe.Pointer(cfontname));
   return fontname;
}

func Window_set_font(handle HWND, fontname string) C.int {
   cfontname := C.CString(fontname);
   defer C.free(unsafe.Pointer(cfontname));
   
   return C.go_window_set_font(unsafe.Pointer(handle), cfontname);
}

func Window_get_pos_size(handle HWND) (C.long, C.long, C.ulong, C.ulong) {
   var x, y C.long;
   var width, height C.ulong;
   C.go_window_get_pos_size(unsafe.Pointer(handle), &x, &y, &width, &height);
   return x, y, width, height;
}

func Window_get_preferred_size(handle HWND) (C.int, C.int) {
   var width, height C.int;
   C.go_window_get_preferred_size(unsafe.Pointer(handle), &width, &height);
   return width, height;
}

func Window_get_text(handle HWND) string {
   ctext := C.go_window_get_text(unsafe.Pointer(handle));
   text := C.GoString(ctext);
   C.dw_free(unsafe.Pointer(ctext));
   return text;
}

func Window_set_text(handle HWND, text string) {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   C.go_window_set_text(unsafe.Pointer(handle), ctext);
}

func Window_set_tooltip(handle HWND, bubbletext string) {
   cbubbletext := C.CString(bubbletext);
   defer C.free(unsafe.Pointer(cbubbletext));
   
   C.go_window_set_tooltip(unsafe.Pointer(handle), cbubbletext);
}

func Window_redraw(handle HWND) {
   C.go_window_redraw(unsafe.Pointer(handle));
}

func Window_capture(handle HWND) {
   C.go_window_capture(unsafe.Pointer(handle));
}

func Window_release() {
   C.dw_window_release();
}

func Window_set_bitmap(window HWND, id C.ulong, filename string) {
   cfilename := C.CString(filename);
   defer C.free(unsafe.Pointer(cfilename));
   
   C.go_window_set_bitmap(unsafe.Pointer(window), id, cfilename);
}

func Window_set_border(handle HWND, border C.int) {
   C.go_window_set_border(unsafe.Pointer(handle), border);
}

func Window_set_focus(handle HWND) {
   C.go_window_set_focus(unsafe.Pointer(handle));
}

func Window_set_gravity(handle HWND, horz C.int, vert C.int) {
   C.go_window_set_gravity(unsafe.Pointer(handle), horz, vert);
}

func Window_set_icon(handle HWND, icon HICN) {
   C.go_window_set_icon(unsafe.Pointer(handle), unsafe.Pointer(icon));
}

func Window_set_pointer(handle HWND, cursortype C.int) {
   C.go_window_set_pointer(unsafe.Pointer(handle), cursortype);
}

func Main() {
   C.dw_main();
}

func Main_iteration() {
   C.dw_main_iteration();
}

func Main_quit() {
   C.dw_main_quit();
}

func Main_sleep(milliseconds C.int) {
   C.dw_main_sleep(milliseconds);
}

func Box_new(btype C.int, pad C.int) HWND {
   return HWND(C.go_box_new(btype, pad));
}

func Box_pack_at_index(box HWND, item HWND, index C.int, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_at_index(unsafe.Pointer(box), unsafe.Pointer(item), index, width, height, hsize, vsize, pad);
}

func Box_pack_end(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_end(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func Box_pack_start(box HWND, item HWND, width C.int, height C.int, hsize C.int, vsize C.int, pad C.int) {
   C.go_box_pack_start(unsafe.Pointer(box), unsafe.Pointer(item), width, height, hsize, vsize, pad);
}

func Box_unpack(handle HWND) C.int {
   return C.go_box_unpack(unsafe.Pointer(handle));
}

func Box_unpack_at_index(handle HWND, index C.int) HWND {
   return HWND(C.go_box_unpack_at_index(unsafe.Pointer(handle), index));
}

func Text_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_text_new(ctext, id));
}

func Entryfield_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_new(ctext, id));
}

func Entryfield_password_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_password_new(ctext, id));
}

func Entryfield_set_limit(handle HWND, limit C.int) {
   C.go_entryfield_set_limit(unsafe.Pointer(handle), limit);
}

func Button_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_button_new(ctext, id));
}

func Clipboard_get_text() string {
   ctext := C.dw_clipboard_get_text();
   text := C.GoString(ctext);
   C.dw_free(unsafe.Pointer(ctext));
   return text;
}

func Clipboard_set_text(text string) {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   C.dw_clipboard_set_text(ctext, C.int(C.strlen(ctext)));
}

func File_browse(title string, defpath string, ext string, flags C.int) string {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cdefpath := C.CString(defpath);
   defer C.free(unsafe.Pointer(cdefpath));
   cext := C.CString(ext);
   defer C.free(unsafe.Pointer(cext));
   
   result := C.dw_file_browse(ctitle, cdefpath, cext, flags);
   defer C.dw_free(unsafe.Pointer(result));
   return C.GoString(result);
}

func Color_choose(value COLOR) COLOR {
   return COLOR(C.dw_color_choose(C.ulong(value)));
}

func Signal_connect(window HWND, signame string, sigfunc unsafe.Pointer, data unsafe.Pointer) {
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

