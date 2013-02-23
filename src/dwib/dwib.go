package dwib

/*
#cgo linux pkg-config: dwib dwindows libxml-2.0
#cgo freebsd pkg-config: dwib dwindows libxml-2.0
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -ldwib -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/BitBucket/dwindows -IC:/Work/BitBucket/dwib -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/BitBucket/dwindows -ldwindows -LC:/Work/BitBucket/dwib -ldwib
#include "dwibglue.c"
*/
import "C"
import "unsafe"
import "dw"

type DWIB unsafe.Pointer

func Load(handle DWIB, name string) dw.HWND {
   cname := C.CString(name);
   defer C.free(unsafe.Pointer(cname));
   
   return dw.HWND(C.goib_load(unsafe.Pointer(handle), cname));
}

func Load_at_index(handle DWIB, name string, dataname string, window dw.HWND, box dw.HWND, index int) int {
   cname := C.CString(name);
   defer C.free(unsafe.Pointer(cname));
   cdataname := C.CString(dataname);
   defer C.free(unsafe.Pointer(cdataname));
   
   return int(C.goib_load_at_index(unsafe.Pointer(handle), cname, cdataname, unsafe.Pointer(window), unsafe.Pointer(box), C.int(index)));
}

func Show(handle dw.HWND) {
   C.goib_show(unsafe.Pointer(handle));
}

func Open(filename string) DWIB {
   cfilename := C.CString(filename);
   defer C.free(unsafe.Pointer(cfilename));
   
   return DWIB(C.goib_open(cfilename));
}

func Close(handle DWIB) {
   C.goib_close(unsafe.Pointer(handle));
}

func Image_root_set(path string) int {
   cpath := C.CString(path);
   defer C.free(unsafe.Pointer(cpath));
   
   return int(C.goib_image_root_set(cpath));
}

func Locale_set(loc string) int {
   cloc := C.CString(loc);
   defer C.free(unsafe.Pointer(cloc));
   
   return int(C.goib_locale_set(cloc));
}

func Window_get_handle(handle dw.HWND, dataname string) dw.HWND {
   cdataname := C.CString(dataname);
   defer C.free(unsafe.Pointer(cdataname));

   return dw.HWND(C.goib_window_get_handle(unsafe.Pointer(handle), cdataname));
}

