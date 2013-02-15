package main

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/Netlabs/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/Netlabs/dwindows -ldw
#include <dw.h>
#include <stdlib.h>

int go_init(int newthread)
{
   int argc = 0;
   char **argv = NULL;
   
   return dw_init(newthread, argc, argv);
}

int go_messagebox(char *title, int flags, char *message)
{
   return dw_messagebox(title, flags, message);
}

void *go_window_new(void *owner, char *title, unsigned long flags)
{
   return (void *)dw_window_new((HWND)owner, title, flags);
}

int go_window_show(void *handle)
{
   return dw_window_show((HWND)handle);
}

void go_window_set_pos(void *handle, long x, long y)
{
   dw_window_set_pos((HWND)handle, x, y);
}

void go_window_set_pos_size(void *handle, long x, long y, unsigned long width, unsigned long height)
{
   dw_window_set_pos_size((HWND)handle, x, y, width, height);
}

void go_window_set_size(void *handle, unsigned long width, unsigned long height)
{
   dw_window_set_size((HWND)handle, width, height);
}

int go_window_set_color(void *handle, unsigned long fore, unsigned long back)
{
   return dw_window_set_color((HWND)handle, fore, back);
}

void go_window_set_style(void *handle, unsigned long style, unsigned long mask)
{
   dw_window_set_style((HWND)handle, style, mask);
}

void go_window_click_default(void *window, void *next)
{
   dw_window_click_default((HWND)window, (HWND)next);
}

void go_window_default(void *window, void *defaultitem)
{
   dw_window_default((HWND)window, (HWND)defaultitem);
}

void *go_box_new(int type, int pad)
{
   return (void *)dw_box_new(type, pad);
}

void go_box_pack_at_index(void *box, void *item, int index, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_at_index((HWND)box, (HWND)item, index, width, height, hsize, vsize, pad);
}

void go_box_pack_end(void *box, void *item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_end((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

void go_box_pack_start(void *box, void *item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_start((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

int go_box_unpack(void *handle)
{
   return dw_box_unpack((HWND)handle);
}

void *go_box_unpack_at_index(void *handle, int index)
{
   return (void *)dw_box_unpack_at_index((HWND)handle, index);
}

void *go_text_new(char *text, unsigned long id)
{
   return (void *)dw_text_new(text, id);
}

void *go_entryfield_new(char *text, unsigned long id)
{
   return (void *)dw_entryfield_new(text, id);
}

void *go_entryfield_password_new(char *text, unsigned long id)
{
   return (void *)dw_entryfield_password_new(text, id);
}

void go_entryfield_set_limit(void *handle, int limit)
{
   dw_entryfield_set_limit((HWND)handle, limit);
}

void *go_button_new(char *text, unsigned long id)
{
   return (void *)dw_button_new(text, id);
}
*/
import "C"
import "unsafe"

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

func main() {
   dw := new(DW);
   
   /* Initialize the Dynamic Windows engine */
   dw.init(TRUE);

   /* Create our window */
   mainwindow := dw.window_new( HWND_DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", C.DW_FCF_SYSMENU | C.DW_FCF_TITLEBAR | C.DW_FCF_TASKLIST | C.DW_FCF_DLGBORDER | C.DW_FCF_SIZEBORDER | C.DW_FCF_MINMAX);
   
   lbbox := dw.box_new(C.DW_VERT, 10);

   dw.box_pack_start(mainwindow, lbbox, 150, 70, TRUE, TRUE, 0);

   /* Copy and Paste */
   browsebox := dw.box_new(C.DW_HORZ, 0);

   dw.box_pack_start(lbbox, browsebox, 0, 0, FALSE, FALSE, 0);

   copypastefield := dw.entryfield_new("", 0);

   dw.entryfield_set_limit(copypastefield, 260);

   dw.box_pack_start(browsebox, copypastefield, -1, -1, TRUE, FALSE, 4);

   copybutton := dw.button_new("Copy", 0);

   dw.box_pack_start(browsebox, copybutton, -1, -1, FALSE, FALSE, 0);

   pastebutton := dw.button_new("Paste", 0);

   dw.box_pack_start(browsebox, pastebutton, -1, -1, FALSE, FALSE, 0);

   /* Archive Name */
   stext := dw.text_new("File to browse", 0);

   dw.window_set_style(stext, C.DW_DT_VCENTER, C.DW_DT_VCENTER);

   dw.box_pack_start(lbbox, stext, 130, 15, TRUE, TRUE, 2);

   browsebox = dw.box_new(C.DW_HORZ, 0);

   dw.box_pack_start(lbbox, browsebox, 0, 0, TRUE, TRUE, 0);

   entryfield := dw.entryfield_new("", 100);

   dw.entryfield_set_limit(entryfield, 260);

   dw.box_pack_start(browsebox, entryfield, 100, 15, TRUE, TRUE, 4);

   browsefilebutton := dw.button_new("Browse File", 1001);

   dw.box_pack_start(browsebox, browsefilebutton, 40, 15, TRUE, TRUE, 0);

   browsefolderbutton := dw.button_new("Browse Folder", 1001);

   dw.box_pack_start(browsebox, browsefolderbutton, 40, 15, TRUE, TRUE, 0);

   dw.window_set_color(browsebox, C.DW_CLR_PALEGRAY, C.DW_CLR_PALEGRAY);
   dw.window_set_color(stext, C.DW_CLR_BLACK, C.DW_CLR_PALEGRAY);

   /* Buttons */
   buttonbox := dw.box_new(C.DW_HORZ, 10);

   dw.box_pack_start(lbbox, buttonbox, 0, 0, TRUE, TRUE, 0);

   cancelbutton := dw.button_new("Exit", 1002);
   dw.box_pack_start(buttonbox, cancelbutton, 130, 30, TRUE, TRUE, 2);

   cursortogglebutton := dw.button_new("Set Cursor pointer - CLOCK", 1003);
   dw.box_pack_start(buttonbox, cursortogglebutton, 130, 30, TRUE, TRUE, 2);

   okbutton := dw.button_new("Turn Off Annoying Beep!", 1001);
   dw.box_pack_start(buttonbox, okbutton, 130, 30, TRUE, TRUE, 2);

   dw.box_unpack(cancelbutton);
   dw.box_pack_start(buttonbox, cancelbutton, 130, 30, TRUE, TRUE, 2);
   dw.window_click_default(mainwindow, cancelbutton);

   colorchoosebutton := dw.button_new("Color Chooser Dialog", 1004);
   dw.box_pack_at_index(buttonbox, colorchoosebutton, 1, 130, 30, TRUE, TRUE, 2);

   /* Set some nice fonts and colors */
   dw.window_set_color(lbbox, C.DW_CLR_DARKCYAN, C.DW_CLR_PALEGRAY);
   dw.window_set_color(buttonbox, C.DW_CLR_DARKCYAN, C.DW_CLR_PALEGRAY);
   dw.window_set_color(okbutton, C.DW_CLR_PALEGRAY, C.DW_CLR_DARKCYAN);

   /*dw_signal_connect(browsefilebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(browse_file_callback), DW_POINTER(notebookbox1));
   dw_signal_connect(browsefolderbutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(browse_folder_callback), DW_POINTER(notebookbox1));
   dw_signal_connect(copybutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(copy_clicked_callback), DW_POINTER(copypastefield));
   dw_signal_connect(pastebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(paste_clicked_callback), DW_POINTER(copypastefield));
   dw_signal_connect(okbutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(beep_callback), DW_POINTER(notebookbox1));
   dw_signal_connect(cancelbutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(exit_callback), DW_POINTER(mainwindow));
   dw_signal_connect(cursortogglebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(cursortoggle_callback), DW_POINTER(mainwindow));
   dw_signal_connect(colorchoosebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(colorchoose_callback), DW_POINTER(mainwindow));*/
   
   /* Set the default field */
   dw.window_default(mainwindow, copypastefield);

   //dw.signal_connect(mainwindow, DW_SIGNAL_DELETE, DW_SIGNAL_FUNC(exit_callback), DW_POINTER(mainwindow));
   /*
   * The following is a special case handler for the Mac and other platforms which contain
   * an application object which can be closed.  It function identically to a window delete/close
   * request except it applies to the entire application not an individual window. If it is not
   * handled or you allow the default handler to take place the entire application will close.
   * On platforms which do not have an application object this line will be ignored.
   */
   //dw.signal_connect(DW_DESKTOP, DW_SIGNAL_DELETE, DW_SIGNAL_FUNC(exit_callback), DW_POINTER(mainwindow));
   //timerid = dw.timer_connect(2000, DW_SIGNAL_FUNC(timer_callback), 0);
   dw.window_set_size(mainwindow, 640, 550);
   dw.window_show(mainwindow);
   
   dw.main();
}

