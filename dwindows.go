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
*/
import "C"
import "unsafe"

type HWND unsafe.Pointer
type DW struct { }

const (
   FALSE C.int = iota
   TRUE
)

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

func main() {
   dw := new(DW);
   
   /* Initialize the Dynamic Windows engine */
   dw.init(TRUE);

   /* Create our window */
   mainwindow := dw.window_new( nil, "dwindows test UTF8 中国語 (繁体) cañón", C.DW_FCF_SYSMENU | C.DW_FCF_TITLEBAR | C.DW_FCF_TASKLIST | C.DW_FCF_DLGBORDER | C.DW_FCF_SIZEBORDER | C.DW_FCF_MINMAX);
   dw.window_show(mainwindow);
   
   dw.messagebox("Test", C.DW_MB_OK, "This is a test");
}

