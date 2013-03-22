package dwib

/*
#cgo linux pkg-config: dwib dwindows libxml-2.0
#cgo freebsd pkg-config: dwib dwindows libxml-2.0
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -ldwib -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/BitBucket/dwindows -IC:/Work/BitBucket/dwib -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/BitBucket/dwindows -ldw -LC:/Work/BitBucket/dwib -ldwib
#include "dwibglue.c"
*/
import "C"
import "unsafe"
import "dw"

type DWIB unsafe.Pointer

func Load(handle DWIB, name string) dw.HWND {
   cname := C.CString(name);
   defer C.free(unsafe.Pointer(cname));
   
   return dw.HANDLE_TO_HWND(dw.POINTER_TO_HANDLE(dw.POINTER(C.goib_load(unsafe.Pointer(handle), cname))));
}

func Load_at_index(handle DWIB, name string, dataname string, window dw.HANDLE, box dw.HANDLE, index int) int {
   cname := C.CString(name);
   defer C.free(unsafe.Pointer(cname));
   cdataname := C.CString(dataname);
   defer C.free(unsafe.Pointer(cdataname));
   
   return int(C.goib_load_at_index(unsafe.Pointer(handle), cname, cdataname, unsafe.Pointer(dw.HANDLE_TO_POINTER(window)), unsafe.Pointer(dw.HANDLE_TO_POINTER(box)), C.int(index)));
}

func LoadAtIndex(handle DWIB, name string, dataname string, window dw.HANDLE, box dw.HANDLE, index int) int {
    return Load_at_index(handle, name, dataname, window, box, index);
}

func Show(handle dw.HANDLE) {
   C.goib_show(unsafe.Pointer(dw.HANDLE_TO_POINTER(handle)));
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

func ImageRootSet(path string) int {
    return Image_root_set(path);
}

func Locale_set(loc string) int {
   cloc := C.CString(loc);
   defer C.free(unsafe.Pointer(cloc));
   
   return int(C.goib_locale_set(cloc));
}

func LocaleSet(loc string) int {
    return Locale_set(loc);
}

func Window_get_handle(handle dw.HANDLE, dataname string) dw.HANDLE {
   cdataname := C.CString(dataname);
   defer C.free(unsafe.Pointer(cdataname));

   return dw.POINTER_TO_HANDLE(dw.POINTER(C.goib_window_get_handle(unsafe.Pointer(dw.HANDLE_TO_POINTER(handle)), cdataname)));
}

func GetHandle(handle dw.HANDLE, dataname string) dw.HANDLE {
    return Window_get_handle(handle, dataname);
}


