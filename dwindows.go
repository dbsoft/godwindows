package main

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/Netlabs/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/Netlabs/dwindows -ldwindows
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
   return dw_messagebox(title, DW_MB_OK, message);
}

*/
import "C"
import "unsafe"

const (
   FALSE C.int = iota
   TRUE
)

type DW struct { }

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

func main() {
   dw := new(DW);
   
   /* Initialize the Dynamic Windows engine */
   dw.init(TRUE);

   /* Create our window */
   dw.messagebox("Test", 0, "This is a test");
}

