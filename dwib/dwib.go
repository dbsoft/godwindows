package dwib

/*
#cgo linux pkg-config: dwib dwindows libxml-2.0
#cgo freebsd pkg-config: dwib dwindows libxml-2.0
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -ldwib -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/DBSoft/dwindows -IC:/Work/DBSoft/dwib -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/DBSoft/dwindows -ldw -LC:/Work/DBSoft/dwib -ldwib
#include "dwibglue.c"
*/
import "C"
import "unsafe"
import "hg.code.sf.net/p/godwindows/code.hg/dw"

type DWIB C.uintptr_t

// Loads a window with the specified name from an XML tree.
func Load(handle DWIB, name string) dw.HWND {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	return dw.UINTPTR_TO_HWND(uintptr(C.goib_load(C.uintptr_t(handle), cname)))
}

// Loads a part of a window layout specified by dataname with the specified window name from an XML tree and packs it into box at index.
func Load_at_index(handle DWIB, name string, dataname string, window dw.HANDLE, box dw.HANDLE, index int) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cdataname := C.CString(dataname)
	defer C.free(unsafe.Pointer(cdataname))

	return int(C.goib_load_at_index(C.uintptr_t(handle), cname, cdataname, C.uintptr_t(dw.HANDLE_TO_UINTPTR(window)), C.uintptr_t(dw.HANDLE_TO_UINTPTR(box)), C.int(index)))
}

// Loads a part of a window layout specified by dataname with the specified window name from an XML tree and packs it into box at index.
func LoadAtIndex(handle DWIB, name string, dataname string, window dw.HANDLE, box dw.HANDLE, index int) int {
	return Load_at_index(handle, name, dataname, window, box, index)
}

// Shows a window loaded with dwib.Load() using the stored settings.
func Show(handle dw.HANDLE) {
	C.goib_show(C.uintptr_t(dw.HANDLE_TO_UINTPTR(handle)))
}

// Loads an XML templates and returns a handle to the XML tree.
func Open(filename string) DWIB {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	return DWIB(C.goib_open(cfilename))
}

// Closes a handle to an XML tree returned by dwib.Open*() and frees the memory associated with the tree.
func Close(handle DWIB) {
	C.goib_close(C.uintptr_t(handle))
}

// Update the location of the image root for locating image files.
func Image_root_set(path string) int {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	return int(C.goib_image_root_set(cpath))
}

// Update the location of the image root for locating image files.
func ImageRootSet(path string) int {
	return Image_root_set(path)
}

// Update the locale used when identifying locating strings during creation.
func Locale_set(loc string) int {
	cloc := C.CString(loc)
	defer C.free(unsafe.Pointer(cloc))

	return int(C.goib_locale_set(cloc))
}

// Update the locale used when identifying locating strings during creation.
func LocaleSet(loc string) int {
	return Locale_set(loc)
}

// Gets the window handle for a named widget.
func Window_get_handle(handle dw.HANDLE, dataname string) dw.HANDLE {
	cdataname := C.CString(dataname)
	defer C.free(unsafe.Pointer(cdataname))

	return dw.UINTPTR_TO_HWND(uintptr(C.goib_window_get_handle(C.uintptr_t(dw.HANDLE_TO_UINTPTR(handle)), cdataname)))
}

// Gets the window handle for a named widget.
func GetHandle(handle dw.HANDLE, dataname string) dw.HANDLE {
	return Window_get_handle(handle, dataname)
}
