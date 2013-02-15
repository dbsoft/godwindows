package main

/*
#cgo linux CFLAGS: -I/usr/local/include -g -O2 -pthread -I/usr/include/gtk-3.0 -I/usr/include/pango-1.0 -I/usr/include/gio-unix-2.0/ -I/usr/include/atk-1.0 -I/usr/include/cairo -I/usr/include/gdk-pixbuf-2.0 -I/usr/include/freetype2 -I/usr/include/glib-2.0 -I/usr/lib/i386-linux-gnu/glib-2.0/include -I/usr/include/pixman-1 -I/usr/include/libpng12 -pthread -I/usr/include/glib-2.0 -I/usr/lib/i386-linux-gnu/glib-2.0/include -I/usr/include/gtk-3.0 -I/usr/include/libsoup-2.4 -I/usr/include/pango-1.0 -I/usr/include/gio-unix-2.0/ -I/usr/include/atk-1.0 -I/usr/include/cairo -I/usr/include/gdk-pixbuf-2.0 -I/usr/include/freetype2 -I/usr/include/pixman-1 -I/usr/include/libpng12 -I/usr/include/libxml2 -I/usr/include/webkitgtk-3.0 -D__UNIX__
#cgo linux LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -lgtk-3 -lgdk-3 -latk-1.0 -lgio-2.0 -lpangocairo-1.0 -lgdk_pixbuf-2.0 -lcairo-gobject -lpango-1.0 -lcairo -lgobject-2.0 -lglib-2.0 -lpthread
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
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

