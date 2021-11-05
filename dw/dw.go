package dw

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/DBSoft/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/DBSoft/dwindows -ldw
#include "dwglue.c"
*/
import "C"
import "unsafe"
import "runtime"
import "reflect"
import "os"

type HANDLE interface {
	GetHandle() unsafe.Pointer
	GetType() C.uint
}
type DRAWABLE interface {
	DrawPoint(x int, y int)
	DrawLine(x1 int, y1 int, x2 int, y2 int)
	DrawPolygon(flags int, x []int, y []int)
	DrawRect(fill int, x int, y int, width int, height int)
	DrawArc(flags int, xorigin int, yorigin int, x1 int, y1 int, x2 int, y2 int)
	DrawText(x int, y int, text string)
	BitBltStretchWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int, srcwidth int, srcheight int) int
	BitBltStretchPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int, srcwidth int, srcheight int) int
	BitBltWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int)
	BitBltPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int)
}
type HGENERIC struct {
	hwnd unsafe.Pointer
}
type HWND struct {
	hwnd unsafe.Pointer
}
type HENTRYFIELD struct {
	hwnd unsafe.Pointer
}
type HTEXT struct {
	hwnd unsafe.Pointer
}
type HTREE struct {
	hwnd unsafe.Pointer
}
type HCONTAINER struct {
	hwnd       unsafe.Pointer
	filesystem bool
}
type HMLE struct {
	hwnd unsafe.Pointer
}
type HBUTTON struct {
	hwnd unsafe.Pointer
}
type HSPINBUTTON struct {
	hwnd unsafe.Pointer
}
type HNOTEBOOK struct {
	hwnd unsafe.Pointer
}
type HBOX struct {
	hwnd unsafe.Pointer
}
type HSCROLLBOX struct {
	hwnd unsafe.Pointer
}
type HMENUITEM struct {
	hwnd unsafe.Pointer
}
type HLISTBOX struct {
	hwnd unsafe.Pointer
}
type HPERCENT struct {
	hwnd unsafe.Pointer
}
type HSLIDER struct {
	hwnd unsafe.Pointer
}
type HSCROLLBAR struct {
	hwnd unsafe.Pointer
}
type HRENDER struct {
	hwnd unsafe.Pointer
}
type HHTML struct {
	hwnd unsafe.Pointer
}
type HCALENDAR struct {
	hwnd unsafe.Pointer
}
type HBITMAP struct {
	hwnd unsafe.Pointer
}
type HSPLITBAR struct {
	hwnd unsafe.Pointer
}
type HTREEITEM struct {
	htreeitem unsafe.Pointer
	htree     HANDLE
}
type HCONTINS struct {
	ptr        unsafe.Pointer
	rowcount   int
	hcont      HANDLE
	filesystem bool
}
type HDIALOG struct {
	hdialog unsafe.Pointer
}
type HEV struct {
	hev unsafe.Pointer
}
type HMTX struct {
	hmtx unsafe.Pointer
}
type HICN unsafe.Pointer
type HTIMER struct {
	tid C.int
}
type HMENUI struct {
	hmenui unsafe.Pointer
}
type HPIXMAP struct {
	hpixmap unsafe.Pointer
}
type HPRINT struct {
	hprint  unsafe.Pointer
	jobname string
}
type HNOTEPAGE struct {
	pageid    C.ulong
	hnotebook HANDLE
}
type COLOR C.ulong
type POINTER unsafe.Pointer
type SIGNAL_FUNC unsafe.Pointer

type Env struct {
	OSName, BuildDate, BuildTime                       string
	MajorVersion, MinorVersion, MajorBuild, MinorBuild C.short
	DWMajorVersion, DWMinorVersion, DWSubVersion       C.short
}

// Define our exported constants
const (
	FALSE int = iota
	TRUE
)

var DESKTOP HWND

// Varaibles to pass if "none/nil" is intended
var NOHWND HWND
var NOHTIMER HTIMER
var NOHPRINT HPRINT
var NOHPIXMAP HPIXMAP
var NOHMENUI HMENUI
var NOMENU HMENUI
var NOHTREEITEM HTREEITEM
var NOHICN HICN = nil

// Import as much as we can from C
var HORZ = C.DW_HORZ
var VERT = C.DW_VERT

// Message box return values
var MB_RETURN_OK = C.DW_MB_RETURN_OK
var MB_RETURN_YES = C.DW_MB_RETURN_YES
var MB_RETURN_NO = C.DW_MB_RETURN_NO
var MB_RETURN_CANCEL = C.DW_MB_RETURN_CANCEL

// Message box button options
var MB_OK = C.DW_MB_OK
var MB_OKCANCEL = C.DW_MB_OKCANCEL
var MB_YESNO = C.DW_MB_YESNO
var MB_YESNOCANCEL = C.DW_MB_YESNOCANCEL

// Message box icons
var MB_WARNING = C.DW_MB_WARNING
var MB_ERROR = C.DW_MB_ERROR
var MB_INFORMATION = C.DW_MB_INFORMATION
var MB_QUESTION = C.DW_MB_QUESTION

/* Preset Pointers */
var POINTER_DEFAULT = C.DW_POINTER_DEFAULT
var POINTER_ARROW = C.DW_POINTER_ARROW
var POINTER_CLOCK = C.DW_POINTER_CLOCK
var POINTER_QUESTION = C.DW_POINTER_QUESTION

/* Draw Text Flags */
var DT_LEFT uint = C.DW_DT_LEFT
var DT_CENTER uint = C.DW_DT_CENTER
var DT_RIGHT uint = C.DW_DT_RIGHT
var DT_VCENTER uint = C.DW_DT_VCENTER
var DT_WORDBREAK uint = C.DW_DT_WORDBREAK

/* Window Frame Creation Flags */
var FCF_CLOSEBUTTON uint = C.DW_FCF_CLOSEBUTTON
var FCF_TITLEBAR uint = C.DW_FCF_TITLEBAR
var FCF_SYSMENU uint = C.DW_FCF_SYSMENU
var FCF_SIZEBORDER uint = C.DW_FCF_SIZEBORDER
var FCF_MINBUTTON uint = C.DW_FCF_MINBUTTON
var FCF_MAXBUTTON uint = C.DW_FCF_MAXBUTTON
var FCF_MINMAX uint = C.DW_FCF_MINMAX
var FCF_DLGBORDER uint = C.DW_FCF_DLGBORDER
var FCF_BORDER uint = C.DW_FCF_BORDER
var FCF_TASKLIST uint = C.DW_FCF_TASKLIST
var FCF_HIDEBUTTON uint = C.DW_FCF_HIDEBUTTON
var FCF_HIDEMAX uint = C.DW_FCF_HIDEMAX
var FCF_MAXIMIZE uint = C.DW_FCF_MAXIMIZE
var FCF_MINIMIZE uint = C.DW_FCF_MINIMIZE
var FCF_COMPOSITED uint = C.DW_FCF_COMPOSITED
var FCF_TEXTURED uint = C.DW_FCF_TEXTURED

var LIT_NONE = C.DW_LIT_NONE

var MLE_CASESENSITIVE = C.DW_MLE_CASESENSITIVE

/* Button Styles */
var BS_NOBORDER uint = C.DW_BS_NOBORDER

/* Key Code Modifiers */
var KC_CTRL = C.KC_CTRL
var KC_SHIFT = C.KC_SHIFT
var KC_ALT = C.KC_ALT

/* Menu Presets */
var MENU_SEPARATOR = C.DW_MENU_SEPARATOR
var MENU_AUTO uint = C.DW_MENU_AUTO
var MENU_POPUP uint = ^uint(0)

var PERCENT_INDETERMINATE uint = ^uint(0)

/* Return value error codes */
var ERROR_NONE = C.DW_ERROR_NONE
var ERROR_GENERAL = C.DW_ERROR_GENERAL
var ERROR_TIMEOUT = C.DW_ERROR_TIMEOUT
var ERROR_NON_INIT = C.DW_ERROR_NON_INIT
var ERROR_NO_MEM = C.DW_ERROR_NO_MEM
var ERROR_INTERRUPT = C.DW_ERROR_INTERRUPT
var ERROR_UNKNOWN = C.DW_ERROR_UNKNOWN

/* Embedded HTML actions */
var HTML_GOBACK = C.DW_HTML_GOBACK
var HTML_GOFORWARD = C.DW_HTML_GOFORWARD
var HTML_GOHOME = C.DW_HTML_GOHOME
var HTML_SEARCH = C.DW_HTML_SEARCH
var HTML_RELOAD = C.DW_HTML_RELOAD
var HTML_STOP = C.DW_HTML_STOP
var HTML_PRINT = C.DW_HTML_PRINT

/* Drawing flags  */
var DRAW_DEFAULT = C.DW_DRAW_DEFAULT
var DRAW_FILL = C.DW_DRAW_FILL
var DRAW_FULL = C.DW_DRAW_FULL
var DRAW_NOAA = C.DW_DRAW_NOAA

/* Preset Drawing Colors */
var CLR_BLACK = COLOR(C.DW_CLR_BLACK)
var CLR_DARKRED = COLOR(C.DW_CLR_DARKRED)
var CLR_DARKGREEN = COLOR(C.DW_CLR_DARKGREEN)
var CLR_BROWN = COLOR(C.DW_CLR_BROWN)
var CLR_DARKBLUE = COLOR(C.DW_CLR_DARKBLUE)
var CLR_DARKPINK = COLOR(C.DW_CLR_DARKPINK)
var CLR_DARKCYAN = COLOR(C.DW_CLR_DARKCYAN)
var CLR_PALEGRAY = COLOR(C.DW_CLR_PALEGRAY)
var CLR_DARKGRAY = COLOR(C.DW_CLR_DARKGRAY)
var CLR_RED = COLOR(C.DW_CLR_RED)
var CLR_GREEN = COLOR(C.DW_CLR_GREEN)
var CLR_YELLOW = COLOR(C.DW_CLR_YELLOW)
var CLR_BLUE = COLOR(C.DW_CLR_BLUE)
var CLR_PINK = COLOR(C.DW_CLR_PINK)
var CLR_CYAN = COLOR(C.DW_CLR_CYAN)
var CLR_WHITE = COLOR(C.DW_CLR_WHITE)
var CLR_DEFAULT = COLOR(C.DW_CLR_DEFAULT)

/* Signal handler defines */
var SIGNAL_CONFIGURE = C.DW_SIGNAL_CONFIGURE
var SIGNAL_KEY_PRESS = C.DW_SIGNAL_KEY_PRESS
var SIGNAL_BUTTON_PRESS = C.DW_SIGNAL_BUTTON_PRESS
var SIGNAL_BUTTON_RELEASE = C.DW_SIGNAL_BUTTON_RELEASE
var SIGNAL_MOTION_NOTIFY = C.DW_SIGNAL_MOTION_NOTIFY
var SIGNAL_DELETE = C.DW_SIGNAL_DELETE
var SIGNAL_EXPOSE = C.DW_SIGNAL_EXPOSE
var SIGNAL_CLICKED = C.DW_SIGNAL_CLICKED
var SIGNAL_ITEM_ENTER = C.DW_SIGNAL_ITEM_ENTER
var SIGNAL_ITEM_CONTEXT = C.DW_SIGNAL_ITEM_CONTEXT
var SIGNAL_ITEM_SELECT = C.DW_SIGNAL_ITEM_SELECT
var SIGNAL_LIST_SELECT = C.DW_SIGNAL_LIST_SELECT
var SIGNAL_SET_FOCUS = C.DW_SIGNAL_SET_FOCUS
var SIGNAL_VALUE_CHANGED = C.DW_SIGNAL_VALUE_CHANGED
var SIGNAL_SWITCH_PAGE = C.DW_SIGNAL_SWITCH_PAGE
var SIGNAL_COLUMN_CLICK = C.DW_SIGNAL_COLUMN_CLICK
var SIGNAL_TREE_EXPAND = C.DW_SIGNAL_TREE_EXPAND

/* status of menu items */
var MIS_ENABLED uint = C.DW_MIS_ENABLED
var MIS_DISABLED uint = C.DW_MIS_DISABLED
var MIS_CHECKED uint = C.DW_MIS_CHECKED
var MIS_UNCHECKED uint = C.DW_MIS_UNCHECKED

/* Gravity */
var GRAV_TOP = C.DW_GRAV_TOP
var GRAV_LEFT = C.DW_GRAV_LEFT
var GRAV_CENTER = C.DW_GRAV_CENTER
var GRAV_RIGHT = C.DW_GRAV_RIGHT
var GRAV_BOTTOM = C.DW_GRAV_BOTTOM
var GRAV_OBSTACLES = C.DW_GRAV_OBSTACLES

/* Container Flags */
var CFA_BITMAPORICON uint = C.DW_CFA_BITMAPORICON
var CFA_STRING uint = C.DW_CFA_STRING
var CFA_ULONG uint = C.DW_CFA_ULONG
var CFA_TIME uint = C.DW_CFA_TIME
var CFA_DATE uint = C.DW_CFA_DATE
var CFA_CENTER uint = C.DW_CFA_CENTER
var CFA_LEFT uint = C.DW_CFA_LEFT
var CFA_RIGHT uint = C.DW_CFA_RIGHT

var CFA_STRINGANDICON uint = C.DW_CFA_STRINGANDICON
var CFA_HORZSEPARATOR uint = C.DW_CFA_HORZSEPARATOR
var CFA_SEPARATOR uint = C.DW_CFA_SEPARATOR

var CRA_SELECTED uint = C.DW_CRA_SELECTED
var CRA_CUROSRED uint = C.DW_CRA_CURSORED

/* Mouse buttons */
var BUTTON1_MASK = C.DW_BUTTON1_MASK
var BUTTON2_MASK = C.DW_BUTTON2_MASK
var BUTTON3_MASK = C.DW_BUTTON3_MASK

/* File dialog */
var FILE_OPEN = C.DW_FILE_OPEN
var FILE_SAVE = C.DW_FILE_SAVE
var DIRECTORY_OPEN = C.DW_DIRECTORY_OPEN

/* Key codes */
var VK_LBUTTON = int(C.VK_LBUTTON)
var VK_RBUTTON = int(C.VK_RBUTTON)
var VK_CANCEL = int(C.VK_CANCEL)
var VK_MBUTTON = int(C.VK_MBUTTON)
var VK_TAB = int(C.VK_TAB)
var VK_CLEAR = int(C.VK_CLEAR)
var VK_RETURN = int(C.VK_RETURN)
var VK_PAUSE = int(C.VK_PAUSE)
var VK_CAPITAL = int(C.VK_CAPITAL)
var VK_ESCAPE = int(C.VK_ESCAPE)
var VK_SPACE = int(C.VK_SPACE)
var VK_PRIOR = int(C.VK_PRIOR)
var VK_NEXT = int(C.VK_NEXT)
var VK_END = int(C.VK_END)
var VK_HOME = int(C.VK_HOME)
var VK_LEFT = int(C.VK_LEFT)
var VK_UP = int(C.VK_UP)
var VK_RIGHT = int(C.VK_RIGHT)
var VK_DOWN = int(C.VK_DOWN)
var VK_SELECT = int(C.VK_SELECT)
var VK_PRINT = int(C.VK_PRINT)
var VK_EXECUTE = int(C.VK_EXECUTE)
var VK_SNAPSHOT = int(C.VK_SNAPSHOT)
var VK_INSERT = int(C.VK_INSERT)
var VK_DELETE = int(C.VK_DELETE)
var VK_HELP = int(C.VK_HELP)
var VK_LWIN = int(C.VK_LWIN)
var VK_RWIN = int(C.VK_RWIN)
var VK_NUMPAD0 = int(C.VK_NUMPAD0)
var VK_NUMPAD1 = int(C.VK_NUMPAD1)
var VK_NUMPAD2 = int(C.VK_NUMPAD2)
var VK_NUMPAD3 = int(C.VK_NUMPAD3)
var VK_NUMPAD4 = int(C.VK_NUMPAD4)
var VK_NUMPAD5 = int(C.VK_NUMPAD5)
var VK_NUMPAD6 = int(C.VK_NUMPAD6)
var VK_NUMPAD7 = int(C.VK_NUMPAD7)
var VK_NUMPAD8 = int(C.VK_NUMPAD8)
var VK_NUMPAD9 = int(C.VK_NUMPAD9)
var VK_MULTIPLY = int(C.VK_MULTIPLY)
var VK_ADD = int(C.VK_ADD)
var VK_SEPARATOR = int(C.VK_SEPARATOR)
var VK_SUBTRACT = int(C.VK_SUBTRACT)
var VK_DECIMAL = int(C.VK_DECIMAL)
var VK_DIVIDE = int(C.VK_DIVIDE)
var VK_F1 = int(C.VK_F1)
var VK_F2 = int(C.VK_F2)
var VK_F3 = int(C.VK_F3)
var VK_F4 = int(C.VK_F4)
var VK_F5 = int(C.VK_F5)
var VK_F6 = int(C.VK_F6)
var VK_F7 = int(C.VK_F7)
var VK_F8 = int(C.VK_F8)
var VK_F9 = int(C.VK_F9)
var VK_F10 = int(C.VK_F10)
var VK_F11 = int(C.VK_F11)
var VK_F12 = int(C.VK_F12)
var VK_F13 = int(C.VK_F13)
var VK_F14 = int(C.VK_F14)
var VK_F15 = int(C.VK_F15)
var VK_F16 = int(C.VK_F16)
var VK_F17 = int(C.VK_F17)
var VK_F18 = int(C.VK_F18)
var VK_F19 = int(C.VK_F19)
var VK_F20 = int(C.VK_F20)
var VK_F21 = int(C.VK_F21)
var VK_F22 = int(C.VK_F22)
var VK_F23 = int(C.VK_F23)
var VK_F24 = int(C.VK_F24)
var VK_NUMLOCK = int(C.VK_NUMLOCK)
var VK_SCROLL = int(C.VK_SCROLL)
var VK_LSHIFT = int(C.VK_LSHIFT)
var VK_RSHIFT = int(C.VK_RSHIFT)
var VK_LCONTROL = int(C.VK_LCONTROL)
var VK_RCONTROL = int(C.VK_RCONTROL)

// Cache the function pointers so they don't get garbage collected
var backs []unsafe.Pointer

// Convert a resource ID into a pointer
func RESOURCE(id uintptr) unsafe.Pointer {
	return unsafe.Pointer(id)
}

// Convert component colors into a COLOR type
func RGB(red uint8, green uint8, blue uint8) COLOR {
	lred := C.ulong(red)
	lgreen := C.ulong(green)
	lblue := C.ulong(blue)
	return COLOR((0xF0000000 | (lred) | (lgreen << 8) | (lblue << 16)))
}

// Convert a POINTER to a HANDLE (use with care)
func POINTER_TO_HANDLE(ptr POINTER) HANDLE {
	return HANDLE(HGENERIC{unsafe.Pointer(ptr)})
}

// Convert a HANDLE to a UINTPTR, mostly used for display purposes
func HANDLE_TO_UINTPTR(handle HANDLE) uintptr {
	return uintptr(handle.GetHandle())
}

// Convert a HANDLE to a POINTER (use with care)
func HANDLE_TO_POINTER(handle HANDLE) POINTER {
	return POINTER(handle.GetHandle())
}

// Convert a HNOTEPAGE to a UINT, mostly used for display purposes
func HNOTEPAGE_TO_UINT(handle HNOTEPAGE) uint {
	return uint(handle.pageid)
}

/* Functions to convert from HANDLE to specific types..
 * These will only work if the hanldle was of the
 * correct type, or were HGENERIC. Use with care.
 */

// Convert HANDLE to HWND (use with care)
func HANDLE_TO_HWND(handle HANDLE) HWND {
	if handle.GetType() == 1 || handle.GetType() == 0 {
		return HWND{handle.GetHandle()}
	}
	return HWND{nil}
}

// Convert HANDLE to HENTRYFIELD (use with care)
func HANDLE_TO_HENTRYFIELD(handle HANDLE) HENTRYFIELD {
	if handle.GetType() == 2 || handle.GetType() == 0 {
		return HENTRYFIELD{handle.GetHandle()}
	}
	return HENTRYFIELD{nil}
}

// Convert HANDLE to HTEXT (use with care)
func HANDLE_TO_HTEXT(handle HANDLE) HTEXT {
	if handle.GetType() == 3 || handle.GetType() == 0 {
		return HTEXT{handle.GetHandle()}
	}
	return HTEXT{nil}
}

// Convert HANDLE to HTREE (use with care)
func HANDLE_TO_HTREE(handle HANDLE) HTREE {
	if handle.GetType() == 4 || handle.GetType() == 0 {
		return HTREE{handle.GetHandle()}
	}
	return HTREE{nil}
}

// Convert HANDLE to HCONTAINER (use with care)
func HANDLE_TO_HCONTAINER(handle HANDLE) HCONTAINER {
	if handle.GetType() == 5 || handle.GetType() == 0 {
		filesystem := false
		if Window_get_data(HCONTAINER{handle.GetHandle(), false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return HCONTAINER{handle.GetHandle(), filesystem}
	}
	return HCONTAINER{nil, false}
}

// Convert HANDLE to HMLE (use with care)
func HANDLE_TO_HMLE(handle HANDLE) HMLE {
	if handle.GetType() == 6 || handle.GetType() == 0 {
		return HMLE{handle.GetHandle()}
	}
	return HMLE{nil}
}

// Convert HANDLE to HBUTTON (use with care)
func HANDLE_TO_HBUTTON(handle HANDLE) HBUTTON {
	if handle.GetType() == 7 || handle.GetType() == 0 {
		return HBUTTON{handle.GetHandle()}
	}
	return HBUTTON{nil}
}

// Convert HANDLE to HSPINBUTTON (use with care)
func HANDLE_TO_HSPINBUTTON(handle HANDLE) HSPINBUTTON {
	if handle.GetType() == 8 || handle.GetType() == 0 {
		return HSPINBUTTON{handle.GetHandle()}
	}
	return HSPINBUTTON{nil}
}

// Convert HANDLE to HNOTEBOOK (use with care)
func HANDLE_TO_HNOTEBOOK(handle HANDLE) HNOTEBOOK {
	if handle.GetType() == 9 || handle.GetType() == 0 {
		return HNOTEBOOK{handle.GetHandle()}
	}
	return HNOTEBOOK{nil}
}

// Convert HANDLE to HBOX (use with care)
func HANDLE_TO_HBOX(handle HANDLE) HBOX {
	if handle.GetType() == 10 || handle.GetType() == 0 {
		return HBOX{handle.GetHandle()}
	}
	return HBOX{nil}
}

// Convert HANDLE to HSCROLLBOX (use with care)
func HANDLE_TO_HSCROLLBOX(handle HANDLE) HSCROLLBOX {
	if handle.GetType() == 11 || handle.GetType() == 0 {
		return HSCROLLBOX{handle.GetHandle()}
	}
	return HSCROLLBOX{nil}
}

// Convert HANDLE to HMENUITEM (use with care)
func HANDLE_TO_HMENUITEM(handle HANDLE) HMENUITEM {
	if handle.GetType() == 12 || handle.GetType() == 0 {
		return HMENUITEM{handle.GetHandle()}
	}
	return HMENUITEM{nil}
}

// Convert HANDLE to HLISTBOX (use with care)
func HANDLE_TO_HLISTBOX(handle HANDLE) HLISTBOX {
	if handle.GetType() == 13 || handle.GetType() == 0 {
		return HLISTBOX{handle.GetHandle()}
	}
	return HLISTBOX{nil}
}

// Convert HANDLE to HPERCENT (use with care)
func HANDLE_TO_HPERCENT(handle HANDLE) HPERCENT {
	if handle.GetType() == 14 || handle.GetType() == 0 {
		return HPERCENT{handle.GetHandle()}
	}
	return HPERCENT{nil}
}

// Convert HANDLE to HSLIDER (use with care)
func HANDLE_TO_HSLIDER(handle HANDLE) HSLIDER {
	if handle.GetType() == 15 || handle.GetType() == 0 {
		return HSLIDER{handle.GetHandle()}
	}
	return HSLIDER{nil}
}

// Convert HANDLE to HSCROLLBAR (use with care)
func HANDLE_TO_HSCROLLBAR(handle HANDLE) HSCROLLBAR {
	if handle.GetType() == 16 || handle.GetType() == 0 {
		return HSCROLLBAR{handle.GetHandle()}
	}
	return HSCROLLBAR{nil}
}

// Convert HANDLE to HRENDER (use with care)
func HANDLE_TO_HRENDER(handle HANDLE) HRENDER {
	if handle.GetType() == 17 || handle.GetType() == 0 {
		return HRENDER{handle.GetHandle()}
	}
	return HRENDER{nil}
}

// Convert HANDLE to HHTML (use with care)
func HANDLE_TO_HHTML(handle HANDLE) HHTML {
	if handle.GetType() == 18 || handle.GetType() == 0 {
		return HHTML{handle.GetHandle()}
	}
	return HHTML{nil}
}

// Convert HANDLE to HCALENDAR (use with care)
func HANDLE_TO_HCALENDAR(handle HANDLE) HCALENDAR {
	if handle.GetType() == 19 || handle.GetType() == 0 {
		return HCALENDAR{handle.GetHandle()}
	}
	return HCALENDAR{nil}
}

// Convert HANDLE to HBITMAP (use with care)
func HANDLE_TO_HBITMAP(handle HANDLE) HBITMAP {
	if handle.GetType() == 20 || handle.GetType() == 0 {
		return HBITMAP{handle.GetHandle()}
	}
	return HBITMAP{nil}
}

// Convert HANDLE to HSPLITBAR (use with care)
func HANDLE_TO_HSPLITBAR(handle HANDLE) HSPLITBAR {
	if handle.GetType() == 21 || handle.GetType() == 0 {
		return HSPLITBAR{handle.GetHandle()}
	}
	return HSPLITBAR{nil}
}

func (window HGENERIC) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HGENERIC) GetType() C.uint {
	return 0
}

func (window HWND) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HWND) GetType() C.uint {
	return 1
}

func (window HENTRYFIELD) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HENTRYFIELD) GetType() C.uint {
	return 2
}

func (window HTEXT) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HTEXT) GetType() C.uint {
	return 3
}

func (window HTREE) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HTREE) GetType() C.uint {
	return 4
}

func (window HCONTAINER) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HCONTAINER) GetType() C.uint {
	return 5
}

func (window HMLE) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HMLE) GetType() C.uint {
	return 6
}

func (window HBUTTON) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HBUTTON) GetType() C.uint {
	return 7
}

func (window HSPINBUTTON) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HSPINBUTTON) GetType() C.uint {
	return 8
}

func (window HNOTEBOOK) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HNOTEBOOK) GetType() C.uint {
	return 9
}

func (window HBOX) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HBOX) GetType() C.uint {
	return 10
}

func (window HSCROLLBOX) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HSCROLLBOX) GetType() C.uint {
	return 11
}

func (window HMENUITEM) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HMENUITEM) GetType() C.uint {
	return 12
}

func (window HLISTBOX) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HLISTBOX) GetType() C.uint {
	return 13
}

func (window HPERCENT) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HPERCENT) GetType() C.uint {
	return 14
}

func (window HSLIDER) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HSLIDER) GetType() C.uint {
	return 15
}

func (window HSCROLLBAR) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HSCROLLBAR) GetType() C.uint {
	return 16
}

func (window HRENDER) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HRENDER) GetType() C.uint {
	return 17
}

func (window HHTML) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HHTML) GetType() C.uint {
	return 18
}

func (window HCALENDAR) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HCALENDAR) GetType() C.uint {
	return 19
}

func (window HBITMAP) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HBITMAP) GetType() C.uint {
	return 20
}

func (window HSPLITBAR) GetHandle() unsafe.Pointer {
	return window.hwnd
}

func (window HSPLITBAR) GetType() C.uint {
	return 21
}

func (treeitem HTREEITEM) GetHandle() unsafe.Pointer {
	return treeitem.htreeitem
}

func (window HTREEITEM) GetType() C.uint {
	return 22
}

func (contins HCONTINS) GetHandle() unsafe.Pointer {
	return contins.ptr
}

func (window HCONTINS) GetType() C.uint {
	return 0
}

func (event HEV) GetHandle() unsafe.Pointer {
	return event.hev
}

func (window HEV) GetType() C.uint {
	return 0
}

func (mutex HMTX) GetHandle() unsafe.Pointer {
	return mutex.hmtx
}

func (window HMTX) GetType() C.uint {
	return 0
}

// Initializes the Dynamic Windows engine.
func Init(newthread int) int {
	if len(os.Args) > 0 {
		var argc C.int = C.int(len(os.Args))
		argv := C.go_string_array_make(argc)
		defer C.go_string_array_free(argv, argc)
		for i, s := range os.Args {
			C.go_string_array_set(argv, C.CString(s), C.int(i))
		}
		return int(C.go_init(C.int(newthread), argc, argv))
	}
	return int(C.go_init(C.int(newthread), 0, nil))
}

// Initializes a Go Routine for Dynamic Windows calls
func InitThread() {
	runtime.LockOSThread()

	C._dw_init_thread()
}

// Deinitializes a Go Routine when Dyamic Windows calls are complete
func DeinitThread() {
	C._dw_deinit_thread()

	runtime.UnlockOSThread()
}

// Cleanly terminates a DW session, should be signal handler safe but does not exit.
func Shutdown() {
	C.dw_shutdown()
}

// Returns some information about the current operating environment.
func Environment_query(env *Env) {
	var cenv C.DWEnv
	C.dw_environment_query(&cenv)
	env.OSName = C.GoString((*C.char)(unsafe.Pointer(&cenv.osName[0])))
	env.BuildDate = C.GoString((*C.char)(unsafe.Pointer(&cenv.buildDate[0])))
	env.BuildTime = C.GoString((*C.char)(unsafe.Pointer(&cenv.buildTime[0])))
	env.MajorVersion = cenv.MajorVersion
	env.MinorVersion = cenv.MajorVersion
	env.MajorBuild = cenv.MajorBuild
	env.MinorBuild = cenv.MinorBuild
	env.DWMajorVersion = cenv.DWMajorVersion
	env.DWMinorVersion = cenv.DWMinorVersion
	env.DWSubVersion = cenv.DWSubVersion
}

// Returns some information about the current operating environment.
func EnvironmentGet() Env {
	var env Env
	Environment_query(&env)
	return env
}

// Displays a Message Box with given text and title.
func Messagebox(title string, flags int, message string) int {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))

	return int(C.go_messagebox(ctitle, C.int(flags), cmessage))
}

// Displays a Message Box with given text and title.
func MessageBox(title string, flags int, message string) int {
	return Messagebox(title, flags, message)
}

// Create a new Window Frame.
func Window_new(owner HWND, title string, flags uint) HWND {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return HWND{C.go_window_new(unsafe.Pointer(owner.hwnd), ctitle, C.ulong(flags))}
}

// Create a new Window Frame.
func WindowNew(owner HWND, title string, flags uint) HWND {
	return Window_new(owner, title, flags)
}

// Makes the window visible.
func Window_show(handle HANDLE) int {
	return int(C.go_window_show(handle.GetHandle()))
}

// Makes the window visible.
func (window HWND) Show() int {
	return Window_show(window)
}

// Makes the window invisible.
func Window_hide(handle HANDLE) int {
	return int(C.go_window_hide(handle.GetHandle()))
}

// Makes the window invisible.
func (window HWND) Hide() int {
	return Window_hide(window)
}

// Makes the window bottommost.
func Window_lower(handle HANDLE) int {
	return int(C.go_window_lower(handle.GetHandle()))
}

// Makes the window bottommost.
func (window HWND) Lower() int {
	return Window_lower(window)
}

// Makes the window topmost.
func Window_raise(handle HANDLE) int {
	return int(C.go_window_raise(handle.GetHandle()))
}

// Makes the window topmost.
func (window HWND) Raise() int {
	return Window_raise(window)
}

// Minimizes or Iconifies a top-level window.
func Window_minimize(handle HANDLE) int {
	return int(C.go_window_minimize(handle.GetHandle()))
}

// Minimizes or Iconifies a top-level window.
func (window HWND) Minimize() int {
	return Window_minimize(window)
}

// Sets the position of a given window.
func Window_set_pos(handle HANDLE, x int, y int) {
	C.go_window_set_pos(handle.GetHandle(), C.long(x), C.long(y))
}

// Sets the position of a given window.
func (window HWND) SetPos(x int, y int) {
	Window_set_pos(window, x, y)
}

// Sets the position and size of a given window.
func Window_set_pos_size(handle HANDLE, x int, y int, width uint, height uint) {
	C.go_window_set_pos_size(handle.GetHandle(), C.long(x), C.long(y), C.ulong(width), C.ulong(height))
}

// Sets the position and size of a given window.
func (window HWND) SetPosSize(x int, y int, width uint, height uint) {
	Window_set_pos_size(window, x, y, width, height)
}

// Sets the size of a given window.
func Window_set_size(handle HANDLE, width uint, height uint) {
	C.go_window_set_size(handle.GetHandle(), C.ulong(width), C.ulong(height))
}

// Sets the size of a given window.
func (window HWND) SetSize(width uint, height uint) {
	Window_set_size(window, width, height)
}

// Sets the colors used by a specified widget handle.
func Window_set_color(handle HANDLE, fore COLOR, back COLOR) int {
	return int(C.go_window_set_color(handle.GetHandle(), C.ulong(fore), C.ulong(back)))
}

// Sets the style of a given widget.
func Window_set_style(handle HANDLE, style uint, mask uint) {
	C.go_window_set_style(handle.GetHandle(), C.ulong(style), C.ulong(mask))
}

// Sets widget to click the default dialog item when an ENTER is pressed.
func Window_click_default(window HANDLE, next HANDLE) {
	C.go_window_click_default(window.GetHandle(), next.GetHandle())
}

// Sets widget to click the default dialog item when an ENTER is pressed.
func (window HWND) ClickDefault(next HANDLE) {
	Window_click_default(window, next)
}

// Sets the default focus item for a window/dialog.
func Window_default(window HWND, defaultitem HANDLE) {
	C.go_window_default(unsafe.Pointer(window.hwnd), defaultitem.GetHandle())
}

// Sets the default focus item for a window/dialog.
func (window HWND) Default(defaultitem HANDLE) {
	Window_default(window, defaultitem)
}

// Destroys a window and all of it's children.
func Window_destroy(handle HANDLE) int {
	return int(C.go_window_destroy(handle.GetHandle()))
}

// Disables given widget.
func Window_disable(handle HANDLE) {
	C.go_window_disable(handle.GetHandle())
}

// Enables given widget.
func Window_enable(handle HANDLE) {
	C.go_window_enable(handle.GetHandle())
}

// Gets the child widget handle with specified ID.
func Window_from_id(handle HANDLE, cid int) HGENERIC {
	return HGENERIC{C.go_window_from_id(handle.GetHandle(), C.int(cid))}
}

// Gets the child widget handle with specified ID.
func (window HWND) FromID(cid int) HGENERIC {
	return Window_from_id(window, cid)
}

// Gets a named user data item from a widget handle.
func Window_get_data(window HANDLE, dataname string) POINTER {
	cdataname := C.CString(dataname)
	defer C.free(unsafe.Pointer(cdataname))

	return POINTER(C.go_window_get_data(window.GetHandle(), cdataname))
}

// Add a named user data item to a widget handle.
func Window_set_data(window HANDLE, dataname string, data POINTER) {
	cdataname := C.CString(dataname)
	defer C.free(unsafe.Pointer(cdataname))

	C.go_window_set_data(window.GetHandle(), cdataname, unsafe.Pointer(data))
}

// Returns the current font for the specified widget
func Window_get_font(handle HANDLE) string {
	cfontname := C.go_window_get_font(handle.GetHandle())
	fontname := C.GoString(cfontname)
	C.dw_free(unsafe.Pointer(cfontname))
	return fontname
}

// Sets the font used by a specified widget handle.
func Window_set_font(handle HANDLE, fontname string) int {
	cfontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(cfontname))

	return int(C.go_window_set_font(handle.GetHandle(), cfontname))
}

// Gets the position and size of a given window.
func Window_get_pos_size(handle HANDLE) (int, int, uint, uint) {
	var x, y C.long
	var width, height C.ulong
	C.go_window_get_pos_size(handle.GetHandle(), &x, &y, &width, &height)
	return int(x), int(y), uint(width), uint(height)
}

// Gets the position and size of a given window.
func (window HWND) GetPosSize() (int, int, uint, uint) {
	return Window_get_pos_size(window)
}

// Gets the size the system thinks the widget should be.
func Window_get_preferred_size(handle HANDLE) (int, int) {
	var width, height C.int
	C.go_window_get_preferred_size(handle.GetHandle(), &width, &height)
	return int(width), int(height)
}

// Gets the text used for a given widget.
func Window_get_text(handle HANDLE) string {
	ctext := C.go_window_get_text(handle.GetHandle())
	text := C.GoString(ctext)
	C.dw_free(unsafe.Pointer(ctext))
	return text
}

// Sets the text used for a given widget.
func Window_set_text(handle HANDLE, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_window_set_text(handle.GetHandle(), ctext)
}

// Sets the text used for a given widget's floating bubble help.
func Window_set_tooltip(handle HANDLE, bubbletext string) {
	cbubbletext := C.CString(bubbletext)
	defer C.free(unsafe.Pointer(cbubbletext))

	C.go_window_set_tooltip(handle.GetHandle(), cbubbletext)
}

// Causes entire window to be invalidated and redrawn.
func Window_redraw(handle HANDLE) {
	C.go_window_redraw(handle.GetHandle())
}

// Causes entire window to be invalidated and redrawn.
func (window HWND) Redraw() {
	Window_redraw(window)
}

// Captures the mouse input to this window even if it is outside the bounds.
func Window_capture(handle HANDLE) {
	C.go_window_capture(handle.GetHandle())
}

// Captures the mouse input to this window even if it is outside the bounds.
func (window HWND) Capture() {
	Window_capture(window)
}

// Releases previous mouse capture.
func Window_release() {
	C.dw_window_release()
}

// Releases previous mouse capture.
func (window HWND) Release() {
	Window_release()
}

// Sets the bitmap used for a given widget.
func Window_set_bitmap(window HANDLE, id uint, filename string) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.go_window_set_bitmap(window.GetHandle(), C.ulong(id), cfilename)
}

// Sets the bitmap used for a given widget.
func (window HBUTTON) SetBitmap(id uint, filename string) {
	Window_set_bitmap(window, id, filename)
}

// Sets the bitmap used for a given widget.
func (window HBITMAP) SetBitmap(id uint, filename string) {
	Window_set_bitmap(window, id, filename)
}

// Sets the border size of a specified window handle.
// This function may only work on OS/2.
func Window_set_border(handle HANDLE, border int) {
	C.go_window_set_border(handle.GetHandle(), C.int(border))
}

// Sets the border size of a specified window handle.
// This function may only work on OS/2.
func (window HWND) SetBorder(border int) {
	Window_set_border(window, border)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func Window_set_focus(handle HANDLE) {
	C.go_window_set_focus(handle.GetHandle())
}

// Sets the gravity of a given window.
// Gravity controls which corner of the screen and window the position is relative to.
func Window_set_gravity(handle HANDLE, horz int, vert int) {
	C.go_window_set_gravity(handle.GetHandle(), C.int(horz), C.int(vert))
}

// Sets the gravity of a given window.
// Gravity controls which corner of the screen and window the position is relative to.
func (window HWND) SetGravity(horz int, vert int) {
	Window_set_gravity(window, horz, vert)
}

// Sets the icon used for a given window.
func Window_set_icon(handle HANDLE, icon HICN) {
	C.go_window_set_icon(handle.GetHandle(), unsafe.Pointer(icon))
}

// Sets the icon used for a given window.
func (window HWND) SetIcon(icon HICN) {
	Window_set_icon(window, icon)
}

// Changes the appearance of the mouse pointer.
func Window_set_pointer(handle HANDLE, cursortype int) {
	C.go_window_set_pointer(handle.GetHandle(), C.int(cursortype))
}

/* Start Generic Section ---
 * These need to be implemented by basically every class/type
 */

// Destroys a window and all of it's children.
func (window HWND) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HWND) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HWND) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HWND) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HWND) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HWND) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HWND) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HWND) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HWND) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HWND) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HWND) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HWND) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HWND) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HWND) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Destroys a widget and all of it's children.
func (window HENTRYFIELD) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HENTRYFIELD) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HENTRYFIELD) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HENTRYFIELD) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HENTRYFIELD) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HENTRYFIELD) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HENTRYFIELD) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HENTRYFIELD) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HENTRYFIELD) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HENTRYFIELD) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HENTRYFIELD) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HENTRYFIELD) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HENTRYFIELD) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HENTRYFIELD) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HENTRYFIELD) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HTEXT) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HTEXT) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HTEXT) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HTEXT) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HTEXT) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HTEXT) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HTEXT) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HTEXT) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HTEXT) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HTEXT) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HTEXT) SetText(text string) {
	Window_set_text(window, text)
}

// Changes the appearance of the mouse pointer.
func (window HTEXT) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HTEXT) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HTEXT) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HTREE) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HTREE) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HTREE) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HTREE) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HTREE) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HTREE) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HTREE) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HTREE) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HTREE) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HTREE) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HTREE) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HTREE) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HTREE) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HTREE) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HTREE) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HCONTAINER) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HCONTAINER) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HCONTAINER) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HCONTAINER) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HCONTAINER) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HCONTAINER) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HCONTAINER) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HCONTAINER) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HCONTAINER) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HCONTAINER) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HCONTAINER) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HCONTAINER) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HCONTAINER) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HCONTAINER) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HCONTAINER) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HMLE) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HMLE) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HMLE) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HMLE) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HMLE) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HMLE) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HMLE) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HMLE) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HMLE) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HMLE) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HMLE) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HMLE) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HMLE) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HMLE) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HMLE) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HBUTTON) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HBUTTON) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HBUTTON) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HBUTTON) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HBUTTON) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HBUTTON) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HBUTTON) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HBUTTON) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HBUTTON) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HBUTTON) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HBUTTON) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HBUTTON) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HBUTTON) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HBUTTON) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HBUTTON) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HSPINBUTTON) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HSPINBUTTON) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HSPINBUTTON) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HSPINBUTTON) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HSPINBUTTON) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HSPINBUTTON) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HSPINBUTTON) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HSPINBUTTON) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HSPINBUTTON) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HSPINBUTTON) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HSPINBUTTON) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HSPINBUTTON) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HSPINBUTTON) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HSPINBUTTON) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HSPINBUTTON) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HNOTEBOOK) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HNOTEBOOK) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HNOTEBOOK) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HNOTEBOOK) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HNOTEBOOK) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HNOTEBOOK) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HNOTEBOOK) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HNOTEBOOK) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HNOTEBOOK) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HNOTEBOOK) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HNOTEBOOK) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HNOTEBOOK) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HNOTEBOOK) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HNOTEBOOK) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HNOTEBOOK) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HLISTBOX) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HLISTBOX) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HLISTBOX) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HLISTBOX) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HLISTBOX) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HLISTBOX) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Sets the colors used by a specified widget handle.
func (window HLISTBOX) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HLISTBOX) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HLISTBOX) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget's floating bubble help.
func (window HLISTBOX) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HLISTBOX) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HLISTBOX) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HLISTBOX) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HPERCENT) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HPERCENT) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HPERCENT) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HPERCENT) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HPERCENT) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HPERCENT) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HPERCENT) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HPERCENT) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HPERCENT) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HPERCENT) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HPERCENT) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HPERCENT) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HPERCENT) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HPERCENT) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HPERCENT) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HSLIDER) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HSLIDER) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HSLIDER) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HSLIDER) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HSLIDER) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HSLIDER) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HSLIDER) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HSLIDER) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HSLIDER) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HSLIDER) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HSLIDER) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HSLIDER) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HSLIDER) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HSLIDER) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HSLIDER) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HSCROLLBAR) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HSCROLLBAR) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HSCROLLBAR) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HSCROLLBAR) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HSCROLLBAR) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HSCROLLBAR) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Gets the text used for a given widget.
func (window HSCROLLBAR) GetText() string {
	return Window_get_text(window)
}

// Sets the colors used by a specified widget handle.
func (window HSCROLLBAR) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HSCROLLBAR) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HSCROLLBAR) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget.
func (window HSCROLLBAR) SetText(text string) {
	Window_set_text(window, text)
}

// Sets the text used for a given widget's floating bubble help.
func (window HSCROLLBAR) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HSCROLLBAR) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HSCROLLBAR) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HSCROLLBAR) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HCALENDAR) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HCALENDAR) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HCALENDAR) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HCALENDAR) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HCALENDAR) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HCALENDAR) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Sets the colors used by a specified widget handle.
func (window HCALENDAR) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HCALENDAR) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HCALENDAR) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget's floating bubble help.
func (window HCALENDAR) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HCALENDAR) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HCALENDAR) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HCALENDAR) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HBITMAP) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HBITMAP) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HBITMAP) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HBITMAP) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HBITMAP) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HBITMAP) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Sets the colors used by a specified widget handle.
func (window HBITMAP) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HBITMAP) SetFocus() {
	Window_set_focus(window)
}

// Sets the font used by a specified widget handle.
func (window HBITMAP) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Sets the text used for a given widget's floating bubble help.
func (window HBITMAP) SetTooltip(bubbletext string) {
	Window_set_tooltip(window, bubbletext)
}

// Changes the appearance of the mouse pointer.
func (window HBITMAP) SetPointer(cursortype int) {
	Window_set_pointer(window, cursortype)
}

// Sets the style of a given widget.
func (window HBITMAP) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HBITMAP) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HHTML) Destroy() int {
	return Window_destroy(window)
}

// Disables given widget.
func (window HHTML) Disable() {
	Window_disable(window)
}

// Enables given widget.
func (window HHTML) Enable() {
	Window_enable(window)
}

// Gets a named user data item from a widget handle.
func (window HHTML) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HHTML) GetFont() string {
	return Window_get_font(window)
}

// Gets the size the system thinks the widget should be.
func (window HHTML) GetPreferredSize() (int, int) {
	return Window_get_preferred_size(window)
}

// Sets the current focus widget for a window/dialog.
// This is for use after showing the window/dialog.
func (window HHTML) SetFocus() {
	Window_set_focus(window)
}

// Sets the style of a given widget.
func (window HHTML) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

// Remove widget from the box it is packed into.
func (window HHTML) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HSPLITBAR) Destroy() int {
	return Window_destroy(window)
}

// Gets a named user data item from a widget handle.
func (window HSPLITBAR) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Remove widget from the box it is packed into.
func (window HSPLITBAR) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HBOX) Destroy() int {
	return Window_destroy(window)
}

// Gets a named user data item from a widget handle.
func (window HBOX) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Sets the colors used by a specified widget handle.
func (window HBOX) SetColor(fore COLOR, back COLOR) int {
	return Window_set_color(window, fore, back)
}

// Remove widget from the box it is packed into.
func (window HBOX) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HSCROLLBOX) Destroy() int {
	return Window_destroy(window)
}

// Gets a named user data item from a widget handle.
func (window HSCROLLBOX) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Remove widget from the box it is packed into.
func (window HSCROLLBOX) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HRENDER) Destroy() int {
	return Window_destroy(window)
}

// Gets a named user data item from a widget handle.
func (window HRENDER) GetData(dataname string) POINTER {
	return Window_get_data(window, dataname)
}

// Returns the current font for the specified widget
func (window HRENDER) GetFont() string {
	return Window_get_font(window)
}

// Get the width and height of a text string.
func (window HRENDER) GetTextExtents(text string) (int, int) {
	return Font_text_extents_get(window, NOHPIXMAP, text)
}

// Sets the font used by a specified widget handle.
func (window HRENDER) SetFont(fontname string) int {
	return Window_set_font(window, fontname)
}

// Remove widget from the box it is packed into.
func (window HRENDER) Unpack() int {
	return Box_unpack(window)
}

// Destroys a widget and all of it's children.
func (window HMENUITEM) Destroy() int {
	return Window_destroy(window)
}

// Sets the style of a given widget.
func (window HMENUITEM) SetStyle(style uint, mask uint) {
	Window_set_style(window, style, mask)
}

/* End Generic Section */

// Runs a message loop for Dynamic Windows.
func Main() {
	C.dw_main()
}

// Processes a single message iteration and returns.
func Main_iteration() {
	C.dw_main_iteration()
}

// Processes a single message iteration and returns.
func MainIteration() {
	Main_iteration()
}

// Causes running dw.Main() to return.
func Main_quit() {
	C.dw_main_quit()
}

// Causes running dw.Main() to return.
func MainQuit() {
	Main_quit()
}

// Runs a message loop for Dynamic Windows, for a period of milliseconds.
func Main_sleep(milliseconds int) {
	C.dw_main_sleep(C.int(milliseconds))
}

// Runs a message loop for Dynamic Windows, for a period of milliseconds.
func MainSleep(milliseconds int) {
	Main_sleep(milliseconds)
}

// Create a new Box to be packed.
func Box_new(btype int, pad int) HBOX {
	return HBOX{C.go_box_new(C.int(btype), C.int(pad))}
}

// Create a new Box to be packed.
func BoxNew(btype int, pad int) HBOX {
	return Box_new(btype, pad)
}

// Pack widgets into a box at an arbitrary location.
func Box_pack_at_index(box HANDLE, item HANDLE, index int, width int, height int, hsize int, vsize int, pad int) {
	C.go_box_pack_at_index(box.GetHandle(), item.GetHandle(), C.int(index), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad))
}

// Pack widgets into a box at an arbitrary location.
func (window HWND) PackAtIndex(item HANDLE, index int, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_at_index(window, item, index, width, height, hsize, vsize, pad)
}

// Pack widgets into a box at an arbitrary location.
func (window HBOX) PackAtIndex(item HANDLE, index int, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_at_index(window, item, index, width, height, hsize, vsize, pad)
}

// Pack widgets into a box at an arbitrary location.
func (window HSCROLLBOX) PackAtIndex(item HANDLE, index int, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_at_index(window, item, index, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the end (or bottom).
func Box_pack_end(box HANDLE, item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	C.go_box_pack_end(box.GetHandle(), item.GetHandle(), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad))
}

// Pack widgets into a box from the end (or bottom).
func (window HWND) PackEnd(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_end(window, item, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the end (or bottom).
func (window HBOX) PackEnd(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_end(window, item, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the end (or bottom).
func (window HSCROLLBOX) PackEnd(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_end(window, item, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the start (or top).
func Box_pack_start(box HANDLE, item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	C.go_box_pack_start(box.GetHandle(), item.GetHandle(), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad))
}

// Pack widgets into a box from the start (or top).
func (window HWND) PackStart(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_start(window, item, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the start (or top).
func (window HBOX) PackStart(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_start(window, item, width, height, hsize, vsize, pad)
}

// Pack widgets into a box from the start (or top).
func (window HSCROLLBOX) PackStart(item HANDLE, width int, height int, hsize int, vsize int, pad int) {
	Box_pack_start(window, item, width, height, hsize, vsize, pad)
}

// Remove widget from the box it is packed into.
func Box_unpack(handle HANDLE) int {
	return int(C.go_box_unpack(handle.GetHandle()))
}

// Remove widgets from a box at an arbitrary location.
func Box_unpack_at_index(handle HANDLE, index int) HANDLE {
	return HANDLE(HWND{C.go_box_unpack_at_index(handle.GetHandle(), C.int(index))})
}

// Remove widgets from a box at an arbitrary location.
func (window HWND) UnpackAtIndex(index int) HANDLE {
	return Box_unpack_at_index(window, index)
}

// Remove widgets from a box at an arbitrary location.
func (window HBOX) UnpackAtIndex(index int) HANDLE {
	return Box_unpack_at_index(window, index)
}

// Remove widgets from a box at an arbitrary location.
func (window HSCROLLBOX) UnpackAtIndex(index int) HANDLE {
	return Box_unpack_at_index(window, index)
}

// Create a new static text widget to be packed.
func Text_new(text string, id uint) HTEXT {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HTEXT{C.go_text_new(ctext, C.ulong(id))}
}

// Create a new static text widget to be packed.
func TextNew(text string, id uint) HTEXT {
	return Text_new(text, id)
}

// Create a new status text widget to be packed.
func Status_text_new(text string, id uint) HTEXT {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HTEXT{C.go_status_text_new(ctext, C.ulong(id))}
}

// Create a new status text widget to be packed.
func StatusTextNew(text string, id uint) HTEXT {
	return Status_text_new(text, id)
}

// Create a new entryfield widget to be packed.
func Entryfield_new(text string, id uint) HENTRYFIELD {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HENTRYFIELD{C.go_entryfield_new(ctext, C.ulong(id))}
}

// Create a new entryfield widget to be packed.
func EntryfieldNew(text string, id uint) HENTRYFIELD {
	return Entryfield_new(text, id)
}

// Create a new entryfield password widget to be packed.
func Entryfield_password_new(text string, id uint) HENTRYFIELD {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HENTRYFIELD{C.go_entryfield_password_new(ctext, C.ulong(id))}
}

// Create a new entryfield password widget to be packed.
func EntryfieldPasswordNew(text string, id uint) HENTRYFIELD {
	return Entryfield_password_new(text, id)
}

// Sets the entryfield character limit.
func Entryfield_set_limit(handle HANDLE, limit int) {
	C.go_entryfield_set_limit(handle.GetHandle(), C.int(limit))
}

// Sets the entryfield character limit.
func (handle HENTRYFIELD) SetLimit(limit int) {
	Entryfield_set_limit(handle, limit)
}

// Create a new button widget to be packed.
func Button_new(text string, id uint) HBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HBUTTON{C.go_button_new(ctext, C.ulong(id))}
}

// Create a new button widget to be packed.
func ButtonNew(text string, id uint) HBUTTON {
	return Button_new(text, id)
}

// Gets the contents of the default clipboard as text.
func Clipboard_get_text() string {
	ctext := C.dw_clipboard_get_text()
	text := C.GoString(ctext)
	C.dw_free(unsafe.Pointer(ctext))
	return text
}

// Gets the contents of the default clipboard as text.
func ClipboardGetText() string {
	return Clipboard_get_text()
}

// Sets the contents of the default clipboard to the supplied text.
func Clipboard_set_text(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.dw_clipboard_set_text(ctext, C.int(C.strlen(ctext)))
}

// Sets the contents of the default clipboard to the supplied text.
func ClipboardSetText(text string) {
	Clipboard_set_text(text)
}

// Opens a file dialog and queries user selection.
func File_browse(title string, defpath string, ext string, flags int) string {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cdefpath := C.CString(defpath)
	defer C.free(unsafe.Pointer(cdefpath))
	cext := C.CString(ext)
	defer C.free(unsafe.Pointer(cext))

	result := C.dw_file_browse(ctitle, cdefpath, cext, C.int(flags))
	defer C.dw_free(unsafe.Pointer(result))
	return C.GoString(result)
}

// Opens a file dialog and queries user selection.
func FileBrowse(title string, defpath string, ext string, flags int) string {
	return File_browse(title, defpath, ext, flags)
}

// Allows the user to choose a color using the system's color chooser dialog.
func Color_choose(value COLOR) COLOR {
	return COLOR(C.dw_color_choose(C.ulong(value)))
}

// Allows the user to choose a color using the system's color chooser dialog.
func ColorChoose(value COLOR) COLOR {
	return Color_choose(value)
}

// Add a callback to a timer event.
func Timer_connect(interval int, sigfunc SIGNAL_FUNC, data POINTER) HTIMER {
	backs = append(backs, unsafe.Pointer(sigfunc))
	return HTIMER{C.go_timer_connect(C.int(interval), unsafe.Pointer(sigfunc), unsafe.Pointer(data), 0)}
}

// Create a new timer.
func TimerNew() HTIMER {
	return HTIMER{0}
}

//Removes timer callback.
func Timer_disconnect(id HTIMER) {
	if id.tid > 0 {
		C.dw_timer_disconnect(C.int(id.tid))
	}
}

// Add a callback to a widget event.
func Signal_connect(window HANDLE, signame string, sigfunc SIGNAL_FUNC, data POINTER) {
	csigname := C.CString(signame)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(sigfunc))
	thissigfunc := unsafe.Pointer(sigfunc)
	thisdata := unsafe.Pointer(data)
	C.go_signal_connect(window.GetHandle(), csigname, thissigfunc, thisdata, window.GetType()<<8)
}

// Emits a beep.
func Beep(freq int, dur int) {
	C.dw_beep(C.int(freq), C.int(dur))
}

// Create a menu object to be popped up.
func Menu_new(id uint) HMENUI {
	return HMENUI{C.go_menu_new(C.ulong(id))}
}

// Create a menu object to be popped up.
func MenuNew(id uint) HMENUI {
	return Menu_new(id)
}

// Create a menubar on a window.
func Menubar_new(location HWND) HMENUI {
	return HMENUI{C.go_menubar_new(unsafe.Pointer(location.hwnd))}
}

// Create a menubar on a window.
func (window HWND) MenubarNew() HMENUI {
	return Menubar_new(window)
}

// Adds a menuitem or submenu to an existing menu.
func Menu_append_item(menu HMENUI, title string, id uint, flags uint, end int, check int, submenu HMENUI) HMENUITEM {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return HMENUITEM{C.go_menu_append_item(menu.hmenui, ctitle, C.ulong(id), C.ulong(flags), C.int(end), C.int(check), submenu.hmenui)}
}

// Adds a menuitem or submenu to an existing menu.
func (menui HMENUI) AppendItem(title string, id uint, flags uint, end int, check int, submenu HMENUI) HMENUITEM {
	return Menu_append_item(menui, title, id, flags, end, check, submenu)
}

// Deletes the menu item specified.
func Menu_delete_item(menu HMENUI, id uint) {
	C.go_menu_delete_item(menu.hmenui, C.ulong(id))
}

// Deletes the menu item specified.
func (menui HMENUI) DeleteItem(id uint) {
	Menu_delete_item(menui, id)
}

// Destroys a menu created with dw.Menubar_new or dw.Menu_new.
func Menu_destroy(menu HMENUI) {
	C.go_menu_destroy(menu.hmenui)
}

// Destroys a menu created with dw.MenubarNew or dw.MenuNew.
func (menui HMENUI) Destroy() {
	Menu_destroy(menui)
}

// Sets the state of a menu item.
func Menu_item_set_state(menu HMENUI, id uint, flags uint) {
	C.go_menu_item_set_state(menu.hmenui, C.ulong(id), C.ulong(flags))
}

// Sets the state of a menu item.
func (menui HMENUI) SetState(id uint, flags uint) {
	Menu_item_set_state(menui, id, flags)
}

// Pops up a context menu at given x and y coordinates.
func Menu_popup(menu HMENUI, parent HANDLE, x int, y int) {
	C.go_menu_popup(menu.hmenui, parent.GetHandle(), C.int(x), C.int(y))
}

// Pops up a context menu at given x and y coordinates.
func (menui HMENUI) Popup(parent HANDLE, x int, y int) {
	Menu_popup(menui, parent, x, y)
}

// Create a notebook widget to be packed.
func Notebook_new(id uint, top int) HNOTEBOOK {
	return HNOTEBOOK{C.go_notebook_new(C.ulong(id), C.int(top))}
}

// Create a notebook widget to be packed.
func NotebookNew(id uint, top int) HNOTEBOOK {
	return Notebook_new(id, top)
}

// Packs the specified box into the notebook page.
func Notebook_pack(handle HANDLE, pageid HNOTEPAGE, page HANDLE) {
	C.go_notebook_pack(handle.GetHandle(), pageid.pageid, page.GetHandle())
}

// Packs the specified box into the notebook page.
func (handle HNOTEPAGE) Pack(page HANDLE) {
	Notebook_pack(handle.hnotebook, handle, page)
}

// The contents of the notebook page will be destroyed as well.
func Notebook_page_destroy(handle HANDLE, pageid HNOTEPAGE) {
	C.go_notebook_page_destroy(handle.GetHandle(), pageid.pageid)
}

// The contents of the notebook page will be destroyed as well.
func (handle HNOTEPAGE) Destroy() {
	Notebook_page_destroy(handle.hnotebook, handle)
}

// Get the currently visible page.
func Notebook_page_get(handle HANDLE) HNOTEPAGE {
	return HNOTEPAGE{C.go_notebook_page_get(handle.GetHandle()), handle}
}

// Get the currently visible page.
func (handle HNOTEBOOK) PageGet() HNOTEPAGE {
	return Notebook_page_get(handle)
}

// Adds a new page to specified notebook.
func Notebook_page_new(handle HANDLE, flags uint, front int) HNOTEPAGE {
	return HNOTEPAGE{C.go_notebook_page_new(handle.GetHandle(), C.ulong(flags), C.int(front)), handle}
}

// Adds a new page to specified notebook.
func (handle HNOTEBOOK) PageNew(flags uint, front int) HNOTEPAGE {
	return Notebook_page_new(handle, flags, front)
}

// Sets the currently visible page.
func Notebook_page_set(handle HANDLE, pageid HNOTEPAGE) {
	C.go_notebook_page_set(handle.GetHandle(), pageid.pageid)
}

// Sets the text on the specified notebook tab.
func Notebook_page_set_text(handle HANDLE, pageid HNOTEPAGE, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_notebook_page_set_text(handle.GetHandle(), pageid.pageid, ctext)
}

// Sets the text on the specified notebook tab.
func (handle HNOTEPAGE) SetText(text string) {
	Notebook_page_set_text(handle.hnotebook, handle, text)
}

// Obtains an icon from a file.
func Icon_load_from_file(filename string) HICN {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	return HICN(C.go_icon_load_from_file(cfilename))
}

// Obtains an icon from a file.
func IconLoadFromFile(filename string) HICN {
	return Icon_load_from_file(filename)
}

// Obtains an icon from a module (or header in GTK).
func Icon_load(id uint) HICN {
	return HICN(C.go_icon_load(0, C.ulong(id)))
}

// Obtains an icon from a module (or header in GTK).
func IconLoad(id uint) HICN {
	return Icon_load(id)
}

// Deletes an icon from the taskbar.
func Taskbar_delete(handle HANDLE, icon HICN) {
	C.go_taskbar_delete(handle.GetHandle(), unsafe.Pointer(icon))
}

// Deletes an icon from the taskbar.
func TaskbarDelete(handle HANDLE, icon HICN) {
	Taskbar_delete(handle, icon)
}

// Inserts an icon into the taskbar.
func Taskbar_insert(handle HANDLE, icon HICN, bubbletext string) {
	cbubbletext := C.CString(bubbletext)
	defer C.free(unsafe.Pointer(cbubbletext))

	C.go_taskbar_insert(handle.GetHandle(), unsafe.Pointer(icon), cbubbletext)
}

// Inserts an icon into the taskbar.
func TaskbarInsert(handle HANDLE, icon HICN, bubbletext string) {
	Taskbar_insert(handle, icon, bubbletext)
}

// Create a new Combobox widget to be packed.
func Combobox_new(text string, id uint) HLISTBOX {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HLISTBOX{C.go_combobox_new(ctext, C.ulong(id))}
}

// Create a new Combobox widget to be packed.
func ComboboxNew(text string, id uint) HLISTBOX {
	return Combobox_new(text, id)
}

// Create a new listbox widget to be packed.
func Listbox_new(id uint, multi int) HLISTBOX {
	return HLISTBOX{C.go_listbox_new(C.ulong(id), C.int(multi))}
}

// Create a new listbox widget to be packed.
func ListboxNew(id uint, multi int) HLISTBOX {
	return Listbox_new(id, multi)
}

// Appends the specified text to the listbox's (or combobox) entry list.
func Listbox_append(handle HANDLE, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_listbox_append(handle.GetHandle(), ctext)
}

// Appends the specified text to the listbox's (or combobox) entry list.
func (handle HLISTBOX) Append(text string) {
	Listbox_append(handle, text)
}

// Appends the specified text items to the listbox's (or combobox) entry list.
func Listbox_list_append(handle HANDLE, text []string) {
	count := len(text)
	ctext := C.go_string_array_make(C.int(count))
	defer C.go_string_array_free(ctext, C.int(count))

	for i, s := range text {
		C.go_string_array_set(ctext, C.CString(s), C.int(i))
	}

	C.go_listbox_list_append(handle.GetHandle(), ctext, C.int(count))
}

// Appends the specified text items to the listbox's (or combobox) entry list.
func (handle HLISTBOX) AppendList(text []string) {
	Listbox_list_append(handle, text)
}

func Listbox_insert(handle HANDLE, text string, pos int) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_listbox_insert(handle.GetHandle(), ctext, C.int(pos))
}

// Inserts the specified text into the listbox's (or combobox) entry list.
func (handle HLISTBOX) Insert(text string, pos int) {
	Listbox_insert(handle, text, pos)
}

// Clears the listbox's (or combobox) list of all entries.
func Listbox_clear(handle HANDLE) {
	C.go_listbox_clear(handle.GetHandle())
}

// Clears the listbox's (or combobox) list of all entries.
func (handle HLISTBOX) Clear() {
	Listbox_clear(handle)
}

// Returns the listbox's item count.
func Listbox_count(handle HANDLE) int {
	return int(C.go_listbox_count(handle.GetHandle()))
}

// Returns the listbox's item count.
func (handle HLISTBOX) Count() int {
	return Listbox_count(handle)
}

// Sets the topmost item in the viewport.
func Listbox_set_top(handle HANDLE, top int) {
	C.go_listbox_set_top(handle.GetHandle(), C.int(top))
}

// Sets the topmost item in the viewport.
func (handle HLISTBOX) SetTop(top int) {
	Listbox_set_top(handle, top)
}

// Sets the selection state of a given index.
func Listbox_select(handle HANDLE, index int, state int) {
	C.go_listbox_select(handle.GetHandle(), C.int(index), C.int(state))
}

// Sets the selection state of a given index.
func (handle HLISTBOX) Select(index int, state int) {
	Listbox_select(handle, index, state)
}

// Deletes the item with given index from the list.
func Listbox_delete(handle HANDLE, index int) {
	C.go_listbox_delete(handle.GetHandle(), C.int(index))
}

// Deletes the item with given index from the list.
func (handle HLISTBOX) Delete(index int) {
	Listbox_delete(handle, index)
}

// Get the given index item's text.
func Listbox_get_text(handle HANDLE, index int) string {
	var buf [201]C.char

	C.go_listbox_get_text(handle.GetHandle(), C.int(index), &buf[0], 200)
	return C.GoString((*C.char)(unsafe.Pointer(&buf[0])))
}

// Get the given index item's text.
func (handle HLISTBOX) GetText(index int) string {
	return Listbox_get_text(handle, index)
}

// Sets the text of a given listbox entry.
func Listbox_set_text(handle HANDLE, index int, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_listbox_set_text(handle.GetHandle(), C.int(index), ctext)
}

// Sets the text of a given listbox entry.
func (handle HLISTBOX) SetText(index int, text string) {
	Listbox_set_text(handle, index, text)
}

// Returns the index to the item in the list currently selected.
func Listbox_selected(handle HANDLE) int {
	return int(C.go_listbox_selected(handle.GetHandle()))
}

// Returns the index to the item in the list currently selected.
func (handle HLISTBOX) Selected() int {
	return Listbox_selected(handle)
}

// Returns the index to the current selected item or -1 when done.
func Listbox_selected_multi(handle HANDLE, where int) int {
	return int(C.go_listbox_selected_multi(handle.GetHandle(), C.int(where)))
}

// Returns the index to the current selected item or -1 when done.
func (handle HLISTBOX) SelectedMulti(where int) int {
	return Listbox_selected_multi(handle, where)
}

// Returns the width of the screen.
func Screen_width() int {
	return int(C.dw_screen_width())
}

// Returns the width of the screen.
func ScreenWidth() int {
	return Screen_width()
}

// Returns the height of the screen.
func Screen_height() int {
	return int(C.dw_screen_height())
}

// Returns the height of the screen.
func ScreenHeight() int {
	return Screen_height()
}

// This should return the current color depth.
func Color_depth_get() uint {
	return uint(C.dw_color_depth_get())
}

// This should return the current color depth.
func ColorDepthGet() uint {
	return Color_depth_get()
}

// Sets the current foreground drawing color. Either a color constant such as dw.CLR_BLACK or an RGB color using the dw.RGB().
func Color_foreground_set(color COLOR) {
	C.dw_color_foreground_set(C.ulong(color))
}

// Sets the current foreground drawing color. Either a color constant such as dw.CLR_BLACK or an RGB color using the dw.RGB().
func ColorForegroundSet(color COLOR) {
	Color_foreground_set(color)
}

// Sets the current background drawing color. Either a color constant such as dw.CLR_BLACK or an RGB color using the dw.RGB().
func Color_background_set(color COLOR) {
	C.dw_color_background_set(C.ulong(color))
}

// Sets the current background drawing color. Either a color constant such as dw.CLR_BLACK or an RGB color using the dw.RGB().
func ColorBackgroundSet(color COLOR) {
	Color_background_set(color)
}

// Create a new spinbutton widget to be packed.
func Spinbutton_new(text string, id C.ulong) HSPINBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HSPINBUTTON{C.go_spinbutton_new(ctext, id)}
}

// Create a new spinbutton widget to be packed.
func SpinButtonNew(text string, id C.ulong) HSPINBUTTON {
	return Spinbutton_new(text, id)
}

// Sets the spinbutton value.
func Spinbutton_set_pos(handle HANDLE, position int) {
	C.go_spinbutton_set_pos(handle.GetHandle(), C.long(position))
}

// Sets the spinbutton value.
func (handle HSPINBUTTON) SetPos(position int) {
	Spinbutton_set_pos(handle, position)
}

// Sets the spinbutton limits.
func Spinbutton_set_limits(handle HANDLE, upper int, lower int) {
	C.go_spinbutton_set_limits(handle.GetHandle(), C.long(upper), C.long(lower))
}

// Sets the spinbutton limits.
func (handle HSPINBUTTON) SetLimits(upper int, lower int) {
	Spinbutton_set_limits(handle, upper, lower)
}

// Returns the current value of the spinbutton.
func Spinbutton_get_pos(handle HANDLE) int {
	return int(C.go_spinbutton_get_pos(handle.GetHandle()))
}

// Returns the current value of the spinbutton.
func (handle HSPINBUTTON) GetPos() int {
	return Spinbutton_get_pos(handle)
}

// Create a new radiobutton widget to be packed.
func Radiobutton_new(text string, id uint) HBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HBUTTON{C.go_radiobutton_new(ctext, C.ulong(id))}
}

// Create a new radiobutton widget to be packed.
func RadioButtonNew(text string, id uint) HBUTTON {
	return Radiobutton_new(text, id)
}

// Create a new checkbox widget to be packed.
func Checkbox_new(text string, id uint) HBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HBUTTON{C.go_checkbox_new(ctext, C.ulong(id))}
}

// Create a new checkbox widget to be packed.
func CheckButtonNew(text string, id uint) HBUTTON {
	return Checkbox_new(text, id)
}

// Returns the state of the checkbox.
func Checkbox_get(handle HANDLE) int {
	return int(C.go_checkbox_get(handle.GetHandle()))
}

// Returns the state of the checkbox.
func (handle HBUTTON) Get() int {
	return Checkbox_get(handle)
}

// Sets the state of the checkbox.
func Checkbox_set(handle HANDLE, value int) {
	C.go_checkbox_set(handle.GetHandle(), C.int(value))
}

// Sets the state of the checkbox.
func (handle HBUTTON) Set(value int) {
	Checkbox_set(handle, value)
}

// Create a new percent bar widget to be packed.
func Percent_new(id C.ulong) HPERCENT {
	return HPERCENT{C.go_percent_new(id)}
}

// Create a new percent bar widget to be packed.
func PercentNew(id C.ulong) HPERCENT {
	return Percent_new(id)
}

// Sets the percent bar position.
func Percent_set_pos(handle HANDLE, position uint) {
	C.go_percent_set_pos(handle.GetHandle(), C.uint(position))
}

// Sets the percent bar position.
func (handle HPERCENT) SetPos(position uint) {
	Percent_set_pos(handle, position)
}

// Create a new slider widget to be packed.
func Slider_new(vertical int, increments int, id uint) HSLIDER {
	return HSLIDER{C.go_slider_new(C.int(vertical), C.int(increments), C.ulong(id))}
}

// Create a new slider widget to be packed.
func SliderNew(vertical int, increments int, id uint) HSLIDER {
	return Slider_new(vertical, increments, id)
}

// Create a new scrollbar widget to be packed.
func Scrollbar_new(vertical int, id uint) HSCROLLBAR {
	return HSCROLLBAR{C.go_scrollbar_new(C.int(vertical), C.ulong(id))}
}

// Create a new scrollbar widget to be packed.
func ScrollbarNew(vertical int, id uint) HSCROLLBAR {
	return Scrollbar_new(vertical, id)
}

// Returns the position of the slider.
func Slider_get_pos(handle HANDLE) uint {
	return uint(C.go_slider_get_pos(handle.GetHandle()))
}

// Returns the position of the slider.
func (handle HSLIDER) GetPos() uint {
	return Slider_get_pos(handle)
}

// Sets the slider position.
func Slider_set_pos(handle HANDLE, position uint) {
	C.go_slider_set_pos(handle.GetHandle(), C.uint(position))
}

// Sets the slider position.
func (handle HSLIDER) SetPos(position uint) {
	Slider_set_pos(handle, position)
}

// Returns the position of the scrollbar.
func Scrollbar_get_pos(handle HANDLE) uint {
	return uint(C.go_scrollbar_get_pos(handle.GetHandle()))
}

// Returns the position of the scrollbar.
func (handle HSCROLLBAR) GetPos() uint {
	return Scrollbar_get_pos(handle)
}

// Sets the scrollbar position.
func Scrollbar_set_pos(handle HANDLE, position uint) {
	C.go_scrollbar_set_pos(handle.GetHandle(), C.uint(position))
}

// Sets the scrollbar position.
func (handle HSCROLLBAR) SetPos(position uint) {
	Scrollbar_set_pos(handle, position)
}

// Sets the scrollbar range.
func Scrollbar_set_range(handle HANDLE, srange uint, visible uint) {
	C.go_scrollbar_set_range(handle.GetHandle(), C.uint(srange), C.uint(visible))
}

// Sets the scrollbar range.
func (handle HSCROLLBAR) SetRange(srange uint, visible uint) {
	Scrollbar_set_range(handle, srange, visible)
}

// Create a new scrollable Box to be packed.
func Scrollbox_new(btype int, pad int) HSCROLLBOX {
	return HSCROLLBOX{C.go_scrollbox_new(C.int(btype), C.int(pad))}
}

// Create a new scrollable Box to be packed.
func ScrollBoxNew(btype int, pad int) HSCROLLBOX {
	return Scrollbox_new(btype, pad)
}

// Returns the position of the scrollbar in the scrollbox.
func Scrollbox_get_pos(handle HANDLE) (int, int) {
	return int(C.go_scrollbox_get_pos(handle.GetHandle(), C.int(C.DW_HORZ))), int(C.go_scrollbox_get_pos(handle.GetHandle(), C.int(C.DW_VERT)))
}

// Returns the position of the scrollbar in the scrollbox.
func (handle HSCROLLBOX) GetPos() (int, int) {
	return Scrollbox_get_pos(handle)
}

// Gets the range for the scrollbar in the scrollbox.
func Scrollbox_get_range(handle HANDLE) (int, int) {
	return int(C.go_scrollbox_get_range(handle.GetHandle(), C.int(C.DW_HORZ))), int(C.go_scrollbox_get_range(handle.GetHandle(), C.int(C.DW_VERT)))
}

// Gets the range for the scrollbar in the scrollbox.
func (handle HSCROLLBOX) GetRange() (int, int) {
	return Scrollbox_get_range(handle)
}

// Create a new Group Box to be packed.
func Groupbox_new(btype C.int, pad int, title string) HBOX {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return HBOX{C.go_groupbox_new(btype, C.int(pad), ctitle)}
}

// Create a new Group Box to be packed.
func GroupboxNew(btype C.int, pad int, title string) HBOX {
	return Groupbox_new(btype, pad, title)
}

// Creates a rendering context widget to be packed.
func Render_new(id uint) HRENDER {
	return HRENDER{C.go_render_new(C.ulong(id))}
}

// Creates a rendering context widget to be packed.
func RenderNew(id uint) HRENDER {
	return Render_new(id)
}

// Allows the user to choose a font using the system's font chooser dialog.
func Font_choose(currfont string) string {
	ccurrfont := C.CString(currfont)
	defer C.free(unsafe.Pointer(ccurrfont))
	newfont := C.dw_font_choose(ccurrfont)
	defer C.dw_free(unsafe.Pointer(newfont))
	return C.GoString(newfont)
}

// Allows the user to choose a font using the system's font chooser dialog.
func FontChoose(currfont string) string {
	return Font_choose(currfont)
}

// Sets the default font used on text based widgets.
func Font_set_default(fontname string) {
	cfontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(cfontname))
	C.dw_font_set_default(cfontname)
}

// Sets the default font used on text based widgets.
func FontSetFefault(fontname string) {
	Font_set_default(fontname)
}

// Get the width and height of a text string.
func Font_text_extents_get(handle HANDLE, pixmap HPIXMAP, text string) (int, int) {
	var width, height C.int

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_font_text_extents_get(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), ctext, &width, &height)
	return int(width), int(height)
}

// Get the width and height of a text string.
func (pixmap HPIXMAP) GetTextExtents(text string) (int, int) {
	return Font_text_extents_get(NOHWND, pixmap, text)
}

// Creates a pixmap with given parameters.
func Pixmap_new(handle HANDLE, width uint, height uint, depth uint) HPIXMAP {
	return HPIXMAP{C.go_pixmap_new(handle.GetHandle(), C.ulong(width), C.ulong(height), C.ulong(depth))}
}

// Creates a pixmap with given parameters.
func PixmapNew(handle HANDLE, width uint, height uint, depth uint) HPIXMAP {
	return Pixmap_new(handle, width, height, depth)
}

// Creates a pixmap from a file.
func Pixmap_new_from_file(handle HANDLE, filename string) HPIXMAP {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	return HPIXMAP{C.go_pixmap_new_from_file(handle.GetHandle(), cfilename)}
}

// Creates a pixmap from a file.
func PixmapNewFromFile(handle HANDLE, filename string) HPIXMAP {
	return Pixmap_new_from_file(handle, filename)
}

// Creates a pixmap from internal resource graphic specified by id.
func Pixmap_grab(handle HANDLE, id uint) HPIXMAP {
	return HPIXMAP{C.go_pixmap_grab(handle.GetHandle(), C.ulong(id))}
}

// Creates a pixmap from internal resource graphic specified by id.
func PixmapGrab(handle HANDLE, id uint) HPIXMAP {
	return Pixmap_grab(handle, id)
}

// Creates a pixmap from internal resource graphic specified by id.
func (window HRENDER) PixmapGrab(id uint) HPIXMAP {
	return Pixmap_grab(window, id)
}

// Copies from one surface to another.
func Pixmap_bitblt(dest HANDLE, destp HPIXMAP, xdest int, ydest int, width int, height int, src HANDLE, srcp HPIXMAP, xsrc int, ysrc int) {
	C.go_pixmap_bitblt(dest.GetHandle(), unsafe.Pointer(destp.hpixmap), C.int(xdest), C.int(ydest), C.int(width), C.int(height), src.GetHandle(), unsafe.Pointer(srcp.hpixmap), C.int(xsrc), C.int(ysrc))
}

// Copies from one surface to another allowing for stretching.
func Pixmap_stretch_bitblt(dest HANDLE, destp HPIXMAP, xdest int, ydest int, width int, height int, src HANDLE, srcp HPIXMAP, xsrc int, ysrc int, srcwidth int, srcheight int) int {
	return int(C.go_pixmap_stretch_bitblt(dest.GetHandle(), unsafe.Pointer(destp.hpixmap), C.int(xdest), C.int(ydest), C.int(width), C.int(height), src.GetHandle(), unsafe.Pointer(srcp.hpixmap), C.int(xsrc), C.int(ysrc), C.int(srcwidth), C.int(srcheight)))
}

// Copies from one surface to another allowing for stretching.
func (window HRENDER) BitBltStretchPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int, srcwidth int, srcheight int) int {
	return Pixmap_stretch_bitblt(window, NOHPIXMAP, xdest, ydest, width, height, NOHWND, srcp, xsrc, ysrc, srcwidth, srcheight)
}

// Copies from one surface to another allowing for stretching.
func (window HRENDER) BitBltStretchWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int, srcwidth int, srcheight int) int {
	return Pixmap_stretch_bitblt(window, NOHPIXMAP, xdest, ydest, width, height, src, NOHPIXMAP, xsrc, ysrc, srcwidth, srcheight)
}

// Copies from one surface to another allowing for stretching.
func (pixmap HPIXMAP) BitBltStretchPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int, srcwidth int, srcheight int) int {
	return Pixmap_stretch_bitblt(NOHWND, pixmap, xdest, ydest, width, height, NOHWND, srcp, xsrc, ysrc, srcwidth, srcheight)
}

// Copies from one surface to another allowing for stretching.
func (pixmap HPIXMAP) BitBltStretchWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int, srcwidth int, srcheight int) int {
	return Pixmap_stretch_bitblt(NOHWND, pixmap, xdest, ydest, width, height, src, NOHPIXMAP, xsrc, ysrc, srcwidth, srcheight)
}

// Copies from one surface to another.
func (window HRENDER) BitBltPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int) {
	Pixmap_bitblt(window, NOHPIXMAP, xdest, ydest, width, height, NOHWND, srcp, xsrc, ysrc)
}

// Copies from one surface to another.
func (window HRENDER) BitBltWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int) {
	Pixmap_bitblt(window, NOHPIXMAP, xdest, ydest, width, height, src, NOHPIXMAP, xsrc, ysrc)
}

// Copies from one surface to another.
func (pixmap HPIXMAP) BitBltPixmap(xdest int, ydest int, width int, height int, srcp HPIXMAP, xsrc int, ysrc int) {
	Pixmap_bitblt(NOHWND, pixmap, xdest, ydest, width, height, NOHWND, srcp, xsrc, ysrc)
}

// Copies from one surface to another.
func (pixmap HPIXMAP) BitBltWindow(xdest int, ydest int, width int, height int, src HANDLE, xsrc int, ysrc int) {
	Pixmap_bitblt(NOHWND, pixmap, xdest, ydest, width, height, src, NOHPIXMAP, xsrc, ysrc)
}

// Sets the transparent color for a pixmap.
func Pixmap_set_transparent_color(pixmap HPIXMAP, color COLOR) {
	C.go_pixmap_set_transparent_color(unsafe.Pointer(pixmap.hpixmap), C.ulong(color))
}

// Sets the transparent color for a pixmap.
func (pixmap HPIXMAP) SetTransparentColor(color COLOR) {
	Pixmap_set_transparent_color(pixmap, color)
}

// Sets the font used by a specified pixmap.
func Pixmap_set_font(pixmap HPIXMAP, fontname string) int {
	cfontname := C.CString(fontname)
	defer C.free(unsafe.Pointer(cfontname))

	return int(C.go_pixmap_set_font(unsafe.Pointer(pixmap.hpixmap), cfontname))
}

// Sets the font used by a specified pixmap.
func (pixmap HPIXMAP) SetFont(fontname string) int {
	return Pixmap_set_font(pixmap, fontname)
}

// Destroys an allocated pixmap.
func Pixmap_destroy(pixmap HPIXMAP) {
	C.go_pixmap_destroy(unsafe.Pointer(pixmap.hpixmap))
}

// Destroys an allocated pixmap.
func (pixmap HPIXMAP) Destroy() {
	Pixmap_destroy(pixmap)
}

// Returns the width in pixels of the specified pixmap.
func Pixmap_width(pixmap HPIXMAP) int {
	return int(C.go_pixmap_width(unsafe.Pointer(pixmap.hpixmap)))
}

// Returns the width in pixels of the specified pixmap.
func (pixmap HPIXMAP) GetWidth() int {
	return Pixmap_width(pixmap)
}

// Returns the height in pixels of the specified pixmap.
func Pixmap_height(pixmap HPIXMAP) int {
	return int(C.go_pixmap_height(unsafe.Pointer(pixmap.hpixmap)))
}

// Returns the height in pixels of the specified pixmap.
func (pixmap HPIXMAP) GetHeight() int {
	return Pixmap_height(pixmap)
}

// Draw a point.
func Draw_point(handle HANDLE, pixmap HPIXMAP, x int, y int) {
	C.go_draw_point(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(x), C.int(y))
}

// Draw a point on a widget.
func (window HRENDER) DrawPoint(x int, y int) {
	Draw_point(window, NOHPIXMAP, x, y)
}

// Draw a point on a pixmap.
func (pixmap HPIXMAP) DrawPoint(x int, y int) {
	Draw_point(NOHWND, pixmap, x, y)
}

// Draw a line.
func Draw_line(handle HANDLE, pixmap HPIXMAP, x1 int, y1 int, x2 int, y2 int) {
	C.go_draw_line(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// Draw a line on a widget.
func (window HRENDER) DrawLine(x1 int, y1 int, x2 int, y2 int) {
	Draw_line(window, NOHPIXMAP, x1, y1, x2, y2)
}

// Draw a line on a pixmap.
func (pixmap HPIXMAP) DrawLine(x1 int, y1 int, x2 int, y2 int) {
	Draw_line(NOHWND, pixmap, x1, y1, x2, y2)
}

// Draw a polygon.
func Draw_polygon(handle HANDLE, pixmap HPIXMAP, flags int, x []int, y []int) {
	count := len(x)
	if len(y) < count {
		count = len(y)
	}
	cx := make([]C.int, count)
	cy := make([]C.int, count)
	for n := 0; n < count; n++ {
		cx[n] = C.int(x[n])
		cy[n] = C.int(y[n])
	}
	xHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cx)))
	yHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cy)))

	C.go_draw_polygon(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(flags), C.int(count), (*C.int)(unsafe.Pointer(xHeader.Data)), (*C.int)(unsafe.Pointer(yHeader.Data)))
}

// Draw a polygon on a widget.
func (window HRENDER) DrawPolygon(flags int, x []int, y []int) {
	Draw_polygon(window, NOHPIXMAP, flags, x, y)
}

// Draw a polygon on a pixmap.
func (pixmap HPIXMAP) DrawPolygon(flags int, x []int, y []int) {
	Draw_polygon(NOHWND, pixmap, flags, x, y)
}

// Draw a rectangle.
func Draw_rect(handle HANDLE, pixmap HPIXMAP, fill int, x int, y int, width int, height int) {
	C.go_draw_rect(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(fill), C.int(x), C.int(y), C.int(width), C.int(height))
}

// Draw a rectangle on a widget.
func (window HRENDER) DrawRect(fill int, x int, y int, width int, height int) {
	Draw_rect(window, NOHPIXMAP, fill, x, y, width, height)
}

// Draw a rectangle on a pixmap.
func (pixmap HPIXMAP) DrawRect(fill int, x int, y int, width int, height int) {
	Draw_rect(NOHWND, pixmap, fill, x, y, width, height)
}

// Draw an arc.
func Draw_arc(handle HANDLE, pixmap HPIXMAP, flags int, xorigin int, yorigin int, x1 int, y1 int, x2 int, y2 int) {
	C.go_draw_arc(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(flags), C.int(xorigin), C.int(yorigin), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// Draw an arc on a widget.
func (window HRENDER) DrawArc(flags int, xorigin int, yorigin int, x1 int, y1 int, x2 int, y2 int) {
	Draw_arc(window, NOHPIXMAP, flags, xorigin, yorigin, x1, y1, x2, y2)
}

// Draw an arc on a pixmap.
func (pixmap HPIXMAP) DrawArc(flags int, xorigin int, yorigin int, x1 int, y1 int, x2 int, y2 int) {
	Draw_arc(NOHWND, pixmap, flags, xorigin, yorigin, x1, y1, x2, y2)
}

// Draw text.
func Draw_text(handle HANDLE, pixmap HPIXMAP, x int, y int, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_draw_text(handle.GetHandle(), unsafe.Pointer(pixmap.hpixmap), C.int(x), C.int(y), ctext)
}

// Draw text on a widget.
func (window HRENDER) DrawText(x int, y int, text string) {
	Draw_text(window, NOHPIXMAP, x, y, text)
}

// Draw text on a pixmap.
func (pixmap HPIXMAP) DrawText(x int, y int, text string) {
	Draw_text(NOHWND, pixmap, x, y, text)
}

// Returns the current X and Y coordinates of the mouse pointer.
func Pointer_query_pos() (int, int) {
	var x, y C.long
	C.dw_pointer_query_pos(&x, &y)
	return int(x), int(y)
}

// Returns the current X and Y coordinates of the mouse pointer.
func PointerGetPos() (int, int) {
	return Pointer_query_pos()
}

// Sets the X and Y coordinates of the mouse pointer.
func Pointer_set_pos(x int, y int) {
	C.dw_pointer_set_pos(C.long(x), C.long(y))
}

// Sets the X and Y coordinates of the mouse pointer.
func PointerSetPos(x int, y int) {
	Pointer_set_pos(x, y)
}

// Call this after drawing to the screen to make sure anything you have drawn is visible.
func Flush() {
	C.dw_flush()
}

// Create a tree widget to be packed.
func Tree_new(id uint) HTREE {
	return HTREE{C.go_tree_new(C.ulong(id))}
}

// Create a tree widget to be packed.
func TreeNew(id uint) HTREE {
	return Tree_new(id)
}

// Inserts an item into a tree widget.
func Tree_insert(handle HANDLE, title string, icon HICN, parent HTREEITEM, itemdata POINTER) HTREEITEM {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return HTREEITEM{C.go_tree_insert(handle.GetHandle(), ctitle, unsafe.Pointer(icon), parent.htreeitem, unsafe.Pointer(itemdata)), handle}
}

// Inserts an item into a tree widget.
func (handle HTREE) Insert(title string, icon HICN, parent HTREEITEM, itemdata POINTER) HTREEITEM {
	return Tree_insert(handle, title, icon, parent, itemdata)
}

// Inserts an item into a tree widget after another item.
func Tree_insert_after(handle HANDLE, item HTREEITEM, title string, icon HICN, parent HTREEITEM, itemdata POINTER) HTREEITEM {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	return HTREEITEM{C.go_tree_insert_after(handle.GetHandle(), item.htreeitem, ctitle, unsafe.Pointer(icon), parent.htreeitem, unsafe.Pointer(itemdata)), handle}
}

// Inserts an item into a tree widget after another item.
func (handle HTREE) InsertAfter(item HTREEITEM, title string, icon HICN, parent HTREEITEM, itemdata POINTER) HTREEITEM {
	return Tree_insert_after(handle, item, title, icon, parent, itemdata)
}

// Removes all nodes from a tree.
func Tree_clear(handle HANDLE) {
	C.go_tree_clear(handle.GetHandle())
}

// Removes all nodes from a tree.
func (handle HTREE) Clear() {
	Tree_clear(handle)
}

// Removes a node from a tree.
func Tree_item_delete(handle HANDLE, item HTREEITEM) {
	C.go_tree_item_delete(handle.GetHandle(), item.htreeitem)
}

// Removes a node from a tree.
func (handle HTREEITEM) Delete() {
	Tree_item_delete(handle.htree, handle)
}

// Sets the text and icon of an item in a tree widget.
func Tree_item_change(handle HANDLE, item HTREEITEM, title string, icon HICN) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.go_tree_item_change(handle.GetHandle(), item.htreeitem, ctitle, unsafe.Pointer(icon))
}

// Sets the text and icon of an item in a tree widget.
func (handle HTREEITEM) Change(title string, icon HICN) {
	Tree_item_change(handle.htree, handle, title, icon)
}

// Expands a node on a tree.
func Tree_item_expand(handle HANDLE, item HTREEITEM) {
	C.go_tree_item_expand(handle.GetHandle(), item.htreeitem)
}

// Expands a node on a tree.
func (handle HTREEITEM) Expand() {
	Tree_item_expand(handle.htree, handle)
}

// Collapses a node on a tree.
func Tree_item_collapse(handle HANDLE, item HTREEITEM) {
	C.go_tree_item_collapse(handle.GetHandle(), item.htreeitem)
}

// Collapses a node on a tree.
func (handle HTREEITEM) Collapse() {
	Tree_item_collapse(handle.htree, handle)
}

// Sets this item as the active selection.
func Tree_item_select(handle HANDLE, item HTREEITEM) {
	C.go_tree_item_select(handle.GetHandle(), item.htreeitem)
}

// Sets this item as the active selection.
func (handle HTREEITEM) Select() {
	Tree_item_select(handle.htree, handle)
}

// Sets the item data of a tree item.
func Tree_item_set_data(handle HANDLE, item HTREEITEM, itemdata POINTER) {
	C.go_tree_item_set_data(handle.GetHandle(), item.htreeitem, unsafe.Pointer(itemdata))
}

// Sets the item data of a tree item.
func (handle HTREEITEM) SetData(itemdata POINTER) {
	Tree_item_set_data(handle.htree, handle, itemdata)
}

// Gets the item data of a tree item.
func Tree_item_get_data(handle HANDLE, item HTREEITEM) POINTER {
	return POINTER(C.go_tree_item_get_data(handle.GetHandle(), item.htreeitem))
}

// Gets the item data of a tree item.
func (handle HTREEITEM) GetData() POINTER {
	return Tree_item_get_data(handle.htree, handle)
}

// Gets the text an item in a tree widget.
func Tree_get_title(handle HANDLE, item HTREEITEM) string {
	ctitle := C.go_tree_get_title(handle.GetHandle(), item.htreeitem)
	title := C.GoString(ctitle)
	C.dw_free(unsafe.Pointer(ctitle))
	return title
}

// Gets the text an item in a tree widget.
func (handle HTREEITEM) GetTitle() string {
	return Tree_get_title(handle.htree, handle)
}

// Create a new HTML widget to be packed.
func Html_new(id uint) HHTML {
	return HHTML{C.go_html_new(C.ulong(id))}
}

// Create a new HTML widget to be packed.
func HtmlNew(id uint) HHTML {
	return Html_new(id)
}

// Causes the embedded HTML widget to take action.
func Html_action(handle HANDLE, action int) {
	C.go_html_action(handle.GetHandle(), C.int(action))
}

// Causes the embedded HTML widget to take action.
func (handle HHTML) Action(action int) {
	Html_action(handle, action)
}

// Render raw HTML code in the embedded HTML widget.
func Html_raw(handle HANDLE, code string) int {
	ccode := C.CString(code)
	defer C.free(unsafe.Pointer(ccode))

	return int(C.go_html_raw(handle.GetHandle(), ccode))
}

// Render raw HTML code in the embedded HTML widget.
func (handle HHTML) Raw(code string) {
	Html_raw(handle, code)
}

// Render file or web page in the embedded HTML widget.
func Html_url(handle HANDLE, url string) int {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))

	return int(C.go_html_url(handle.GetHandle(), curl))
}

// Render file or web page in the embedded HTML widget.
func (handle HHTML) URL(url string) int {
	return Html_url(handle, url)
}

// Create a new Multiline Editbox widget to be packed.
func Mle_new(id uint) HMLE {
	return HMLE{C.go_mle_new(C.ulong(id))}
}

// Create a new Multiline Editbox widget to be packed.
func MLENew(id uint) HMLE {
	return Mle_new(id)
}

// Adds text to an MLE box and returns the current point.
func Mle_import(handle HANDLE, buffer string, startpoint int) int {
	cbuffer := C.CString(buffer)
	defer C.free(unsafe.Pointer(cbuffer))

	return int(C.go_mle_import(handle.GetHandle(), cbuffer, C.int(startpoint)))
}

// Adds text to an MLE box and returns the current point.
func (handle HMLE) Import(buffer string, startpoint int) int {
	return Mle_import(handle, buffer, startpoint)
}

// Grabs text from an MLE box.
func Mle_export(handle HANDLE, startpoint int, length int) string {
	cbuf := C.calloc(1, C.size_t(length+1))
	C.go_mle_export(handle.GetHandle(), (*C.char)(cbuf), C.int(startpoint), C.int(length))
	buf := C.GoString((*C.char)(cbuf))
	C.free(cbuf)
	return buf
}

// Grabs text from an MLE box.
func (handle HMLE) Export(startpoint int, length int) string {
	return Mle_export(handle, startpoint, length)
}

// Obtains information about an MLE box.
func Mle_get_size(handle HANDLE) (int, int) {
	var bytes, lines C.ulong
	C.go_mle_get_size(handle.GetHandle(), &bytes, &lines)
	return int(bytes), int(lines)
}

// Obtains information about an MLE box.
func (handle HMLE) GetSize() (int, int) {
	return Mle_get_size(handle)
}

// Deletes text from an MLE box.
func Mle_delete(handle HANDLE, startpoint int, length int) {
	C.go_mle_delete(handle.GetHandle(), C.int(startpoint), C.int(length))
}

// Deletes text from an MLE box.
func (handle HMLE) Delete(startpoint int, length int) {
	Mle_delete(handle, startpoint, length)
}

// Clears all text from an MLE box.
func Mle_clear(handle HANDLE) {
	C.go_mle_clear(handle.GetHandle())
}

// Clears all text from an MLE box.
func (handle HMLE) Clear() {
	Mle_clear(handle)
}

// Stops redrawing of an MLE box.
func Mle_freeze(handle HANDLE) {
	C.go_mle_freeze(handle.GetHandle())
}

// Stops redrawing of an MLE box.
func (handle HMLE) Freeze() {
	Mle_freeze(handle)
}

// Resumes redrawing of an MLE box.
func Mle_thaw(handle HANDLE) {
	C.go_mle_thaw(handle.GetHandle())
}

// Resumes redrawing of an MLE box.
func (handle HMLE) Thaw() {
	Mle_thaw(handle)
}

// Sets the current cursor position of an MLE box.
func Mle_set_cursor(handle HANDLE, point int) {
	C.go_mle_set_cursor(handle.GetHandle(), C.int(point))
}

// Sets the current cursor position of an MLE box.
func (handle HMLE) SetCursor(point int) {
	Mle_set_cursor(handle, point)
}

// Sets the visible line of an MLE box.
func Mle_set_visible(handle HANDLE, line int) {
	C.go_mle_set_visible(handle.GetHandle(), C.int(line))
}

// Sets the visible line of an MLE box.
func (handle HMLE) SetVisible(line int) {
	Mle_set_visible(handle, line)
}

// Sets the editablity of an MLE box.
func Mle_set_editable(handle HANDLE, state int) {
	C.go_mle_set_editable(handle.GetHandle(), C.int(state))
}

// Sets the editablity of an MLE box.
func (handle HMLE) SetEditable(state int) {
	Mle_set_editable(handle, state)
}

// Sets the word wrap state of an MLE box.
func Mle_set_word_wrap(handle HANDLE, state int) {
	C.go_mle_set_word_wrap(handle.GetHandle(), C.int(state))
}

// Sets the word wrap state of an MLE box.
func (handle HMLE) SetWordWrap(state int) {
	Mle_set_word_wrap(handle, state)
}

// Finds text in an MLE box.
func Mle_search(handle HANDLE, text string, point int, flags uint) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return int(C.go_mle_search(handle.GetHandle(), ctext, C.int(point), C.ulong(flags)))
}

// Finds text in an MLE box.
func (handle HMLE) Search(text string, point int, flags uint) int {
	return Mle_search(handle, text, point, flags)
}

// Create a container widget to be packed.
func Container_new(id uint, multi int) HCONTAINER {
	return HCONTAINER{C.go_container_new(C.ulong(id), C.int(multi)), false}
}

// Create a container widget to be packed.
func ContainerNew(id uint, multi int) HCONTAINER {
	return Container_new(id, multi)
}

// Sets up the container columns.
func Container_setup(handle HANDLE, flags []uint, titles []string, separator int) int {
	count := len(flags)
	if len(titles) < count {
		count = len(titles)
	}

	ctitles := C.go_string_array_make(C.int(len(titles)))
	defer C.go_string_array_free(ctitles, C.int(len(titles)))
	for i, s := range titles {
		C.go_string_array_set(ctitles, C.CString(s), C.int(i))
	}

	cflags := make([]C.ulong, count)
	for n := 0; n < count; n++ {
		cflags[n] = C.ulong(flags[n])
	}
	flagsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cflags)))
	return int(C.go_container_setup(handle.GetHandle(), (*C.ulong)(unsafe.Pointer(flagsHeader.Data)), ctitles, C.int(count), C.int(separator)))
}

// Sets up the container columns.
func (handle HCONTAINER) Setup(flags []uint, titles []string, separator int) int {
	return Container_setup(handle, flags, titles, separator)
}

// Sets up the filesystem columns, note: filesystem always has an icon/filename field.
func Filesystem_setup(handle HANDLE, flags []uint, titles []string) int {
	count := len(flags)
	if len(titles) < count {
		count = len(titles)
	}

	ctitles := C.go_string_array_make(C.int(len(titles)))
	defer C.go_string_array_free(ctitles, C.int(len(titles)))
	for i, s := range titles {
		C.go_string_array_set(ctitles, C.CString(s), C.int(i))
	}

	cflags := make([]C.ulong, count)
	for n := 0; n < count; n++ {
		cflags[n] = C.ulong(flags[n])
	}
	flagsHeader := (*reflect.SliceHeader)((unsafe.Pointer(&cflags)))
	Window_set_data(handle, "_go_filesystem", POINTER(uintptr(1)))
	return int(C.go_filesystem_setup(handle.GetHandle(), (*C.ulong)(unsafe.Pointer(flagsHeader.Data)), ctitles, C.int(count)))
}

// Sets up the filesystem columns, note: filesystem always has an icon/filename field.
func (handle *HCONTAINER) FileSystemSetup(flags []uint, titles []string) int {
	handle.filesystem = true
	return Filesystem_setup(handle, flags, titles)
}

// Allocates memory used to populate a container.
func Container_alloc(handle HANDLE, rowcount int) HCONTINS {
	return HCONTINS{C.go_container_alloc(handle.GetHandle(), C.int(rowcount)), rowcount, handle, false}
}

// Allocates memory used to populate a container.
func (handle *HCONTAINER) Alloc(rowcount int) HCONTINS {
	contins := Container_alloc(handle, rowcount)
	contins.filesystem = handle.filesystem
	return contins
}

// Sets an item in specified row and column to the given data.
func Container_set_item(handle HANDLE, contins HCONTINS, column int, row int, data POINTER) {
	C.go_container_set_item(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), unsafe.Pointer(data))
}

// Sets an item in specified row and column to the given data.
func (handle HCONTINS) SetItem(column int, row int, data POINTER) {
	if handle.filesystem == true {
		Filesystem_set_item(handle.hcont, handle, column, row, data)
	}
	Container_set_item(handle.hcont, handle, column, row, data)
}

// Sets an item in specified row and column to the given unsigned integer.
func Container_set_item_ulong(handle HANDLE, contins HCONTINS, column int, row int, val uint) {
	C.go_container_set_item_ulong(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.ulong(val))
}

// Sets an item in specified row and column to the given unsigned integer.
func (handle HCONTINS) SetItemULong(column int, row int, val uint) {
	if handle.filesystem == true {
		Filesystem_set_item_ulong(handle.hcont, handle, column, row, val)
	}
	Container_set_item_ulong(handle.hcont, handle, column, row, val)
}

// Sets an item in specified row and column to the given icon.
func Container_set_item_icon(handle HANDLE, contins HCONTINS, column int, row int, icon HICN) {
	C.go_container_set_item_icon(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), unsafe.Pointer(icon))
}

// Sets an item in specified row and column to the given icon.
func (handle HCONTINS) SetItemIcon(column int, row int, icon HICN) {
	if handle.filesystem == true {
		Filesystem_set_item_icon(handle.hcont, handle, column, row, icon)
	}
	Container_set_item_icon(handle.hcont, handle, column, row, icon)
}

// Sets an item in specified row and column to the given time.
func Container_set_item_time(handle HANDLE, contins HCONTINS, column int, row int, seconds int, minutes int, hours int) {
	C.go_container_set_item_time(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.int(seconds), C.int(minutes), C.int(hours))
}

// Sets an item in specified row and column to the given time.
func (handle HCONTINS) SetItemTime(column int, row int, seconds int, minutes int, hours int) {
	if handle.filesystem == true {
		Filesystem_set_item_time(handle.hcont, handle, column, row, seconds, minutes, hours)
	}
	Container_set_item_time(handle.hcont, handle, column, row, seconds, minutes, hours)
}

// Sets an item in specified row and column to the given date.
func Container_set_item_date(handle HANDLE, contins HCONTINS, column int, row int, day int, month int, year int) {
	C.go_container_set_item_date(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.int(day), C.int(month), C.int(year))
}

// Sets an item in specified row and column to the given date.
func (handle HCONTINS) SetItemDate(column int, row int, day int, month int, year int) {
	if handle.filesystem == true {
		Filesystem_set_item_date(handle.hcont, handle, column, row, day, month, year)
	}
	Container_set_item_date(handle.hcont, handle, column, row, day, month, year)
}

// Changes an existing item in specified row and column to the given data.
func Container_change_item(handle HANDLE, column int, row int, data POINTER) {
	C.go_container_change_item(handle.GetHandle(), C.int(column), C.int(row), unsafe.Pointer(data))
}

// Changes an existing item in specified row and column to the given data.
func (handle HCONTAINER) ChangeItem(column int, row int, data POINTER) {
	if handle.filesystem == true {
		Filesystem_change_item(handle, column, row, data)
	}
	Container_change_item(handle, column, row, data)
}

// Changes an existing item in specified row and column to the given unsigned integer.
func Container_change_item_ulong(handle HANDLE, column int, row int, val uint) {
	C.go_container_change_item_ulong(handle.GetHandle(), C.int(column), C.int(row), C.ulong(val))
}

// Changes an existing item in specified row and column to the given unsigned integer.
func (handle HCONTAINER) ChangeItemULong(column int, row int, val uint) {
	if handle.filesystem == true {
		Filesystem_change_item_ulong(handle, column, row, val)
	}
	Container_change_item_ulong(handle, column, row, val)
}

// Changes an existing item in specified row and column to the given icon.
func Container_change_item_icon(handle HANDLE, column int, row int, icon HICN) {
	C.go_container_change_item_icon(handle.GetHandle(), C.int(column), C.int(row), unsafe.Pointer(icon))
}

// Changes an existing item in specified row and column to the given icon.
func (handle HCONTAINER) ChangeItemIcon(column int, row int, icon HICN) {
	if handle.filesystem == true {
		Filesystem_change_item_icon(handle, column, row, icon)
	}
	Container_change_item_icon(handle, column, row, icon)
}

// Changes an existing item in specified row and column to the given time.
func Container_change_item_time(handle HANDLE, column int, row int, seconds int, minutes int, hours int) {
	C.go_container_change_item_time(handle.GetHandle(), C.int(column), C.int(row), C.int(seconds), C.int(minutes), C.int(hours))
}

// Changes an existing item in specified row and column to the given time.
func (handle HCONTAINER) ChangeItemTime(column int, row int, seconds int, minutes int, hours int) {
	if handle.filesystem == true {
		Filesystem_change_item_time(handle, column, row, seconds, minutes, hours)
	}
	Container_change_item_time(handle, column, row, seconds, minutes, hours)
}

// Changes an existing item in specified row and column to the given date.
func Container_change_item_date(handle HANDLE, column int, row int, day int, month int, year int) {
	C.go_container_change_item_date(handle.GetHandle(), C.int(column), C.int(row), C.int(day), C.int(month), C.int(year))
}

// Changes an existing item in specified row and column to the given date.
func (handle HCONTAINER) ChangeItemDate(column int, row int, day int, month int, year int) {
	if handle.filesystem == true {
		Filesystem_change_item_date(handle, column, row, day, month, year)
	}
	Container_change_item_date(handle, column, row, day, month, year)
}

// Sets the width of a column in the container.
func Container_set_column_width(handle HANDLE, column int, width int) {
	C.go_container_set_column_width(handle.GetHandle(), C.int(column), C.int(width))
}

// Sets the width of a column in the container.
func (handle HCONTAINER) SetColumnWidth(column int, width int) {
	Container_set_column_width(handle, column, width)
}

// Sets the title of a row in the container.
func Container_set_row_title(contins HCONTINS, row int, title string) {
	ctitle := C.CString(title)
	C.dw_container_set_row_title(contins.ptr, C.int(row), ctitle)
	defer C.free(unsafe.Pointer(ctitle))
}

// Sets the title of a row in the container.
func (handle HCONTINS) SetRowTitle(row int, title string) {
	Container_set_row_title(handle, row, title)
}

// Sets the pointer of a row in the container.
func Container_set_row_data(contins HCONTINS, row int, data POINTER) {
	C.dw_container_set_row_data(contins.ptr, C.int(row), unsafe.Pointer(data))
}

// Sets the pointer of a row in the container.
func (handle HCONTINS) SetRowData(row int, data POINTER) {
	Container_set_row_data(handle, row, data)
}

// Sets the title of a row in the container.
func Container_change_row_title(handle HANDLE, row int, title string) {
	ctitle := C.CString(title)
	C.go_container_change_row_title(handle.GetHandle(), C.int(row), ctitle)
}

// Sets the title of a row in the container.
func (handle HCONTAINER) ChangeRowTitle(row int, title string) {
	Container_change_row_title(handle, row, title)
}

// Sets the pointer of a row in the container.
func Container_change_row_data(handle HANDLE, row int, data unsafe.Pointer) {
	C.go_container_change_row_data(handle.GetHandle(), C.int(row), data)
}

// Sets the pointer of a row in the container.
func (handle HCONTAINER) ChangeRowData(row int, data POINTER) {
	Container_change_row_data(handle, row, unsafe.Pointer(data))
}

// Inserts allocated rows into the container widget.
func Container_insert(handle HANDLE, contins HCONTINS, rowcount int) {
	C.go_container_insert(handle.GetHandle(), contins.ptr, C.int(rowcount))
	contins.ptr = nil
	contins.rowcount = 0
}

// Inserts allocated rows into the container widget.
func (handle HCONTINS) Insert() {
	Container_insert(handle.hcont, handle, handle.rowcount)
}

// Removes all rows from a container.
func Container_clear(handle HANDLE, redraw int) {
	C.go_container_clear(handle.GetHandle(), C.int(redraw))
}

// Removes all rows from a container.
func (handle HCONTAINER) Clear(redraw int) {
	Container_clear(handle, redraw)
}

// Removes the first x rows from a container.
func Container_delete(handle HANDLE, rowcount int) {
	C.go_container_delete(handle.GetHandle(), C.int(rowcount))
}

// Removes the first x rows from a container.
func (handle HCONTAINER) Delete(rowcount int) {
	Container_delete(handle, rowcount)
}

// Starts a new query of a container.
func Container_query_start(handle HANDLE, flags uint) string {
	cresult := C.go_container_query_start(handle.GetHandle(), C.ulong(flags)&^C.DW_CR_RETDATA)
	result := C.GoString(cresult)
	defer C.dw_free(unsafe.Pointer(cresult))
	return result
}

func Container_query_start_data(handle HANDLE, flags uint) POINTER {
	cresult := C.go_container_query_start(handle.GetHandle(), C.ulong(flags)|C.DW_CR_RETDATA)
	return POINTER(cresult)
}

// Starts a new query of a container.
func (handle HCONTAINER) QueryStart(flags uint) string {
	return Container_query_start(handle, flags)
}

// Continues an existing query of a container.
func Container_query_next(handle HANDLE, flags uint) string {
	cresult := C.go_container_query_next(handle.GetHandle(), C.ulong(flags)&^C.DW_CR_RETDATA)
	result := C.GoString(cresult)
	defer C.dw_free(unsafe.Pointer(cresult))
	return result
}

func Container_query_next_data(handle HANDLE, flags uint) POINTER {
	cresult := C.go_container_query_next(handle.GetHandle(), C.ulong(flags)|C.DW_CR_RETDATA)
	return POINTER(cresult)
}

// Continues an existing query of a container.
func (handle HCONTAINER) QueryNext(flags uint) string {
	return Container_query_next(handle, flags)
}

// Scrolls container up or down.
func Container_scroll(handle HANDLE, direction int, rows int) {
	C.go_container_scroll(handle.GetHandle(), C.int(direction), C.long(rows))
}

// Scrolls container up or down.
func (handle HCONTAINER) Scroll(direction int, rows int) {
	Container_scroll(handle, direction, rows)
}

// Cursors the item with the title speficied, and scrolls to that item.
func Container_cursor(handle HANDLE, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_container_cursor(handle.GetHandle(), ctext)
}

// Cursors the item with the title speficied, and scrolls to that item.
func (handle HCONTAINER) Cursor(text string) {
	Container_cursor(handle, text)
}

// Cursors the item with the data speficied, and scrolls to that item.
func Container_cursor_by_data(handle HANDLE, data POINTER) {
	C.go_container_cursor_by_data(handle.GetHandle(), unsafe.Pointer(data))
}

// Cursors the item with the data speficied, and scrolls to that item.
func (handle HCONTAINER) CursorByData(data POINTER) {
	Container_cursor_by_data(handle, data)
}

// Deletes the item with the title specified.
func Container_delete_row(handle HANDLE, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	C.go_container_delete_row(handle.GetHandle(), ctext)
}

// Deletes the item with the title specified.
func (handle HCONTAINER) DeleteRow(text string) {
	Container_delete_row(handle, text)
}

// Deletes the item with the data specified.
func Container_delete_row_by_data(handle HANDLE, data POINTER) {
	C.go_container_delete_row_by_data(handle.GetHandle(), unsafe.Pointer(data))
}

// Deletes the item with the data specified.
func (handle HCONTAINER) DeleteRowByData(data POINTER) {
	Container_delete_row_by_data(handle, data)
}

// Optimizes the column widths so that all data is visible.
func Container_optimize(handle HANDLE) {
	C.go_container_optimize(handle.GetHandle())
}

// Optimizes the column widths so that all data is visible.
func (handle HCONTAINER) Optimize() {
	Container_optimize(handle)
}

// Sets the alternating row colors for container widget handle.
func Container_set_stripe(handle HANDLE, oddcolor COLOR, evencolor COLOR) {
	C.go_container_set_stripe(handle.GetHandle(), C.ulong(oddcolor), C.ulong(evencolor))
}

// Sets the alternating row colors for container widget handle.
func (handle HCONTAINER) SetStripe(oddcolor COLOR, evencolor COLOR) {
	Container_set_stripe(handle, oddcolor, evencolor)
}

// Gets column type for a container column.
func Container_get_column_type(handle HANDLE, column int) uint {
	return uint(C.go_container_get_column_type(handle.GetHandle(), C.int(column)))
}

// Gets column type for a container column.
func (handle HCONTAINER) GetColumnType(column int) uint {
	if handle.filesystem == true {
		return Filesystem_get_column_type(handle, column)
	}
	return Container_get_column_type(handle, column)
}

// Gets column type for a filesystem container column.
func Filesystem_get_column_type(handle HANDLE, column int) uint {
	return uint(C.go_filesystem_get_column_type(handle.GetHandle(), C.int(column)))
}

// Configures the main filesystem column title for localization.
func Filesystem_set_column_title(handle HANDLE, title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.go_filesystem_set_column_title(handle.GetHandle(), ctitle)
}

// Configures the main filesystem column title for localization.
func (handle HCONTAINER) SetColumnTitle(title string) {
	Filesystem_set_column_title(handle, title)
}

// Sets an item in specified row and column to the given data.
func Filesystem_set_item(handle HANDLE, contins HCONTINS, column int, row int, data POINTER) {
	C.go_filesystem_set_item(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), unsafe.Pointer(data))
}

// Sets an item in specified row and column to the given unsigned integer.
func Filesystem_set_item_ulong(handle HANDLE, contins HCONTINS, column int, row int, val uint) {
	C.go_filesystem_set_item_ulong(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.ulong(val))
}

// Sets an item in specified row and column to the given icon.
func Filesystem_set_item_icon(handle HANDLE, contins HCONTINS, column int, row int, icon HICN) {
	C.go_filesystem_set_item_icon(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), unsafe.Pointer(icon))
}

// Sets an item in specified row and column to the given time.
func Filesystem_set_item_time(handle HANDLE, contins HCONTINS, column int, row int, seconds int, minutes int, hours int) {
	C.go_filesystem_set_item_time(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.int(seconds), C.int(minutes), C.int(hours))
}

// Sets an item in specified row and column to the given date.
func Filesystem_set_item_date(handle HANDLE, contins HCONTINS, column int, row int, day int, month int, year int) {
	C.go_filesystem_set_item_date(handle.GetHandle(), contins.ptr, C.int(column), C.int(row), C.int(day), C.int(month), C.int(year))
}

// Sets the filename and icon of the row in a filesystem style container.
func Filesystem_set_file(handle HANDLE, contins HCONTINS, row int, filename string, icon HICN) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.go_filesystem_set_file(handle.GetHandle(), contins.ptr, C.int(row), cfilename, unsafe.Pointer(icon))
}

// Sets the filename and icon of the row in a filesystem style container.
func (handle HCONTINS) SetFile(row int, filename string, icon HICN) {
	if handle.filesystem == true {
		Filesystem_set_file(handle.hcont, handle, row, filename, icon)
	}
}

// Changes an existing item in specified row and column to the given data.
func Filesystem_change_item(handle HANDLE, column int, row int, data POINTER) {
	C.go_filesystem_change_item(handle.GetHandle(), C.int(column), C.int(row), unsafe.Pointer(data))
}

// Changes an existing item in specified row and column to the given unsigned integer.
func Filesystem_change_item_ulong(handle HANDLE, column int, row int, val uint) {
	C.go_filesystem_change_item_ulong(handle.GetHandle(), C.int(column), C.int(row), C.ulong(val))
}

// Changes an existing item in specified row and column to the given icon.
func Filesystem_change_item_icon(handle HANDLE, column int, row int, icon HICN) {
	C.go_filesystem_change_item_icon(handle.GetHandle(), C.int(column), C.int(row), unsafe.Pointer(icon))
}

// Changes an existing item in specified row and column to the given time.
func Filesystem_change_item_time(handle HANDLE, column int, row int, seconds int, minutes int, hours int) {
	C.go_filesystem_change_item_time(handle.GetHandle(), C.int(column), C.int(row), C.int(seconds), C.int(minutes), C.int(hours))
}

// Changes an existing item in specified row and column to the given date.
func Filesystem_change_item_date(handle HANDLE, column int, row int, day int, month int, year int) {
	C.go_filesystem_change_item_date(handle.GetHandle(), C.int(column), C.int(row), C.int(day), C.int(month), C.int(year))
}

// Changes the filename and icon of the row in a filesystem style container.
func Filesystem_change_file(handle HANDLE, row int, filename string, icon HICN) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	C.go_filesystem_change_file(handle.GetHandle(), C.int(row), cfilename, unsafe.Pointer(icon))
}

// Changes the filename and icon of the row in a filesystem style container.
func (handle HCONTAINER) ChangeFile(row int, filename string, icon HICN) {
	if handle.filesystem == true {
		Filesystem_change_file(handle, row, filename, icon)
	}
}

// Create a new calendar widget to be packed.
func Calendar_new(id uint) HCALENDAR {
	return HCALENDAR{C.go_calendar_new(C.ulong(id))}
}

// Create a new calendar widget to be packed.
func CalendarNew(id uint) HCALENDAR {
	return Calendar_new(id)
}

// Sets the current date of a calendar.
func Calendar_set_date(handle HANDLE, year uint, month uint, day uint) {
	C.go_calendar_set_date(handle.GetHandle(), C.uint(year), C.uint(month), C.uint(day))
}

// Sets the current date of a calendar.
func (handle HCALENDAR) SetDate(year uint, month uint, day uint) {
	Calendar_set_date(handle, year, month, day)
}

// Gets the year, month and day set in the calendar widget.
func Calendar_get_date(handle HANDLE) (uint, uint, uint) {
	var year, month, day C.uint

	C.go_calendar_get_date(handle.GetHandle(), &year, &month, &day)
	return uint(year), uint(month), uint(day)
}

// Gets the year, month and day set in the calendar widget.
func (handle HCALENDAR) GetDate() (uint, uint, uint) {
	return Calendar_get_date(handle)
}

// Create a bitmap widget to be packed.
func Bitmap_new(id uint) HBITMAP {
	return HBITMAP{C.go_bitmap_new(C.ulong(id))}
}

// Create a bitmap widget to be packed.
func BitmapNew(id uint) HBITMAP {
	return Bitmap_new(id)
}

// Create a new bitmap button widget to be packed.
func Bitmapbutton_new(text string, id uint) HBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	return HBUTTON{C.go_bitmapbutton_new(ctext, C.ulong(id))}
}

// Create a new bitmap button widget to be packed.
func BitmapButtonNew(text string, id uint) HBUTTON {
	return Bitmapbutton_new(text, id)
}

// Create a new bitmap button widget to be packed from a file.
func Bitmapbutton_new_from_file(text string, id uint, filename string) HBUTTON {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	return HBUTTON{C.go_bitmapbutton_new_from_file(ctext, C.ulong(id), cfilename)}
}

// Create a new bitmap button widget to be packed from a file.
func BitmapButtonNewFromFile(text string, id uint, filename string) HBUTTON {
	return Bitmapbutton_new_from_file(text, id, filename)
}

// Creates a splitbar widget with given widgets on either side.
func Splitbar_new(btype int, topleft HWND, bottomright HWND, id uint) HSPLITBAR {
	return HSPLITBAR{C.go_splitbar_new(C.int(btype), unsafe.Pointer(topleft.hwnd), unsafe.Pointer(bottomright.hwnd), C.ulong(id))}
}

// Creates a splitbar widget with given widgets on either side.
func SplitbarNew(btype int, topleft HWND, bottomright HWND, id uint) HSPLITBAR {
	return Splitbar_new(btype, topleft, bottomright, id)
}

// Sets the position of a splitbar (pecentage).
func Splitbar_set(handle HANDLE, position float32) {
	C.go_splitbar_set(handle.GetHandle(), C.float(position))
}

// Sets the position of a splitbar (pecentage).
func (handle HSPLITBAR) Set(position float32) {
	Splitbar_set(handle, position)
}

// Gets the position of a splitbar (pecentage).
func Splitbar_get(handle HANDLE) float32 {
	return float32(C.go_splitbar_get(handle.GetHandle()))
}

// Gets the position of a splitbar (pecentage).
func (handle HSPLITBAR) Get() float32 {
	return Splitbar_get(handle)
}

// Creates a new print object.
func PrintNew(jobname string) HPRINT {
	return HPRINT{nil, jobname}
}

// Creates a new print object.
func Print_new(jobname string, flags uint, pages uint, drawfunc SIGNAL_FUNC, drawdata POINTER) HPRINT {
	backs = append(backs, unsafe.Pointer(drawfunc))
	cjobname := C.CString(jobname)
	defer C.free(unsafe.Pointer(cjobname))

	return HPRINT{C.go_print_new(cjobname, C.ulong(flags), C.uint(pages), unsafe.Pointer(drawfunc), unsafe.Pointer(drawdata), 0), jobname}
}

// Runs the print job, causing the draw page callbacks to fire.
func Print_run(print HPRINT, flags uint) int {
	if print.hprint != nil {
		return int(C.go_print_run(unsafe.Pointer(print.hprint), C.ulong(flags)))
	}
	return C.DW_ERROR_UNKNOWN
}

// Runs the print job, causing the draw page callbacks to fire.
func (print HPRINT) Run(flags uint) {
	Print_run(print, flags)
}

// Cancels the print job, typically called from a draw page callback.
func Print_cancel(print HPRINT) {
	if print.hprint != nil {
		C.go_print_cancel(unsafe.Pointer(print.hprint))
	}
}

// Cancels the print job, typically called from a draw page callback.
func (print HPRINT) Cancel() {
	Print_cancel(print)
}

func init() {
	runtime.LockOSThread()
}

// Creates a new mutex semaphore.
func Mutex_new() HMTX {
	return HMTX{C.go_mutex_new()}
}

// Creates a new mutex semaphore.
func MutexNew() HMTX {
	return Mutex_new()
}

// Closes a semaphore created by Mutex_new().
func Mutex_close(handle HMTX) {
	C.go_mutex_close(unsafe.Pointer(handle.hmtx))
}

// Closes a semaphore created by MutexNew().
func (handle HMTX) Close() {
	Mutex_close(handle)
}

// Tries to gain access to the semaphore, if it can't it blocks.
func Mutex_lock(handle HMTX) {
	C.go_mutex_lock(unsafe.Pointer(handle.hmtx))
}

// Tries to gain access to the semaphore, if it can't it blocks.
func (handle HMTX) Lock() {
	Mutex_lock(handle)
}

// Reliquishes the access to the semaphore.
func Mutex_unlock(handle HMTX) {
	C.go_mutex_unlock(unsafe.Pointer(handle.hmtx))
}

// Reliquishes the access to the semaphore.
func (handle HMTX) Unlock() {
	Mutex_unlock(handle)
}

// Tries to gain access to the semaphore.
func Mutex_trylock(handle HMTX) int {
	return int(C.go_mutex_trylock(unsafe.Pointer(handle.hmtx)))
}

// Tries to gain access to the semaphore.
func (handle HMTX) TryLock() int {
	return Mutex_trylock(handle)
}

// Allocates and initializes a dialog.
func Dialog_new() HDIALOG {
	return HDIALOG{C.go_dialog_new()}
}

// Allocates and initializes a dialog.
func DialogNew() HDIALOG {
	return Dialog_new()
}

// Accepts a dialog and returns the given data to the initial call of Dialog_wait().
func Dialog_dismiss(handle HDIALOG, result POINTER) int {
	return int(C.go_dialog_dismiss(unsafe.Pointer(handle.hdialog), unsafe.Pointer(result)))
}

// Returns the given data to the initial call of Wait().
func (handle HDIALOG) Dismiss(result uintptr) int {
	return Dialog_dismiss(handle, POINTER(result))
}

// Accepts a dialog, waits for Dialog_dismiss() to be called by a signal handler with the given dialog.
func Dialog_wait(handle HDIALOG) POINTER {
	return POINTER(C.go_dialog_wait(unsafe.Pointer(handle.hdialog)))
}

// Waits for Dismiss() to be called by a signal handler.
func (handle HDIALOG) Wait() uintptr {
	return uintptr(Dialog_wait(handle))
}

// Creates an unnamed event semaphore.
func Event_new() HEV {
	return HEV{C.go_event_new()}
}

// Creates an unnamed event semaphore.
func EventNew() HEV {
	return Event_new()
}

// Closes a semaphore created by Event_new().
func Event_close(handle *HEV) int {
	retval := int(C.go_event_close(unsafe.Pointer(handle.hev)))
	handle.hev = nil
	return retval
}

// Closes a semaphore created by EventNew().
func (handle *HEV) Close() int {
	return Event_close(handle)
}

// Posts a semaphore created by Event_new(). Causing all threads waiting on this event in Event_wait() to continue.
func Event_post(handle HEV) int {
	return int(C.go_event_post(unsafe.Pointer(handle.hev)))
}

// Posts a semaphore created by EventNew(). Causing all threads waiting on this event in Wait() to continue.
func (handle HEV) Post() int {
	return Event_post(handle)
}

// Resets a semaphore created by Event_new().
func Event_reset(handle HEV) int {
	return int(C.go_event_reset(unsafe.Pointer(handle.hev)))
}

// Resets a semaphore created by EventNew().
func (handle HEV) Reset() int {
	return Event_reset(handle)
}

// Waits on a semaphore created by Event_new(), until the event gets posted or until the timeout expires.
func Event_wait(handle HEV, timeout int) int {
	return int(C.go_event_wait(unsafe.Pointer(handle.hev), C.ulong(timeout)))
}

// Waits on a semaphore created by EventNew(), until the event gets posted or until the timeout expires.
func (handle HEV) Wait(timeout int) int {
	return Event_wait(handle, timeout)
}

var go_flags_no_data C.uint = 1

// Connect a function or closure to a window close (delete) event.
func (window HWND) ConnectDelete(sigfunc func(window HWND) int) {
	csigname := C.CString(C.DW_SIGNAL_DELETE)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a widget clicked event.
func (window HBUTTON) ConnectClicked(sigfunc func(window HBUTTON) int) {
	csigname := C.CString(C.DW_SIGNAL_CLICKED)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a focus clicked event.
func (window HWND) ConnectSetFocus(sigfunc func(window HWND) int) {
	csigname := C.CString(C.DW_SIGNAL_SET_FOCUS)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a key press event.
func (window HWND) ConnectKeyPress(sigfunc func(window HWND, ch uint8, vk int, state int, utf8 string) int) {
	csigname := C.CString(C.DW_SIGNAL_KEY_PRESS)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a key press event.
func (window HRENDER) ConnectKeyPress(sigfunc func(window HRENDER, ch uint8, vk int, state int, utf8 string) int) {
	csigname := C.CString(C.DW_SIGNAL_KEY_PRESS)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a mouse motion event.
func (window HRENDER) ConnectMotion(sigfunc func(window HRENDER, x int, y int, mask int) int) {
	csigname := C.CString(C.DW_SIGNAL_MOTION_NOTIFY)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a mouse button press event.
func (window HRENDER) ConnectButtonPress(sigfunc func(window HRENDER, x int, y int, mask int) int) {
	csigname := C.CString(C.DW_SIGNAL_BUTTON_PRESS)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a mouse button release event.
func (window HRENDER) ConnectButtonRelease(sigfunc func(window HRENDER, x int, y int, mask int) int) {
	csigname := C.CString(C.DW_SIGNAL_BUTTON_RELEASE)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a render expose event.
func (window HRENDER) ConnectExpose(sigfunc func(window HRENDER, x int, y int, width int, height int) int) {
	csigname := C.CString(C.DW_SIGNAL_EXPOSE)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a configure (size change) event.
func (window HRENDER) ConnectConfigure(sigfunc func(window HRENDER, width int, height int) int) {
	csigname := C.CString(C.DW_SIGNAL_CONFIGURE)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a tree ENTER/RETURN press event.
func (window HTREE) ConnectItemEnter(sigfunc func(window HTREE, str string, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_ENTER)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a container ENTER/RETURN press event.
func (window HCONTAINER) ConnectItemEnter(sigfunc func(window HCONTAINER, str string, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_ENTER)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a tree context event.
func (window HTREE) ConnectItemContext(sigfunc func(window HTREE, text string, x int, y int, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_CONTEXT)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a container context event.
func (window HCONTAINER) ConnectItemContext(sigfunc func(window HCONTAINER, text string, x int, y int, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_CONTEXT)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a tree select event.
func (window HTREE) ConnectItemSelect(sigfunc func(window HTREE, item HTREEITEM, text string, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_SELECT)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a container select event.
func (window HCONTAINER) ConnectItemSelect(sigfunc func(window HCONTAINER, item HTREEITEM, text string, itemdata POINTER) int) {
	csigname := C.CString(C.DW_SIGNAL_ITEM_SELECT)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a listbox select event.
func (window HLISTBOX) ConnectListSelect(sigfunc func(window HLISTBOX, index int) int) {
	csigname := C.CString(C.DW_SIGNAL_LIST_SELECT)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a value changed event.
func (window HSCROLLBAR) ConnectValueChanged(sigfunc func(window HSCROLLBAR, index int) int) {
	csigname := C.CString(C.DW_SIGNAL_VALUE_CHANGED)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a value changed event.
func (window HSLIDER) ConnectValueChanged(sigfunc func(window HSLIDER, index int) int) {
	csigname := C.CString(C.DW_SIGNAL_VALUE_CHANGED)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a value changed event.
func (window HSPINBUTTON) ConnectValueChanged(sigfunc func(window HSPINBUTTON, index int) int) {
	csigname := C.CString(C.DW_SIGNAL_VALUE_CHANGED)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a column title click event.
func (window HCONTAINER) ConnectColumnClick(sigfunc func(window HCONTAINER, index int) int) {
	csigname := C.CString(C.DW_SIGNAL_COLUMN_CLICK)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a notebook switch page event.
func (window HNOTEBOOK) ConnectSwitchPage(sigfunc func(window HNOTEBOOK, pageid HNOTEPAGE) int) {
	csigname := C.CString(C.DW_SIGNAL_SWITCH_PAGE)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a tree item (node) expand event.
func (window HTREE) ConnectTreeExpand(sigfunc func(window HTREE, item HTREEITEM) int) {
	csigname := C.CString(C.DW_SIGNAL_TREE_EXPAND)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a menu item clicked event.
func (window HMENUITEM) ConnectClicked(sigfunc func(window HMENUITEM) int) {
	csigname := C.CString(C.DW_SIGNAL_CLICKED)
	defer C.free(unsafe.Pointer(csigname))

	backs = append(backs, unsafe.Pointer(&sigfunc))
	C.go_signal_connect(unsafe.Pointer(window.hwnd), csigname, unsafe.Pointer(&sigfunc), nil, (window.GetType()<<8)|go_flags_no_data)
}

// Connect a function or closure to a timer event.
func (id *HTIMER) Connect(sigfunc func() int, interval int) {
	if id.tid == 0 {
		backs = append(backs, unsafe.Pointer(&sigfunc))
		id.tid = C.go_timer_connect(C.int(interval), unsafe.Pointer(&sigfunc), nil, go_flags_no_data)
	}
}

// Disconnect an active timer event.
func (id HTIMER) Disconnect() {
	if id.tid > 0 {
		C.dw_timer_disconnect(C.int(id.tid))
	}
}

// Connect a function or closure to a print object draw page event.
func (print HPRINT) Connect(drawfunc func(HPRINT, HPIXMAP, int) int, flags uint, pages int) {
	if print.hprint == nil {
		backs = append(backs, unsafe.Pointer(&drawfunc))
		cjobname := C.CString(print.jobname)
		defer C.free(unsafe.Pointer(cjobname))

		print.hprint = C.go_print_new(cjobname, C.ulong(flags), C.uint(pages), unsafe.Pointer(&drawfunc), nil, go_flags_no_data)
	}
}

//export go_callback_remove
func go_callback_remove(pfunc unsafe.Pointer) {
	// Scan through the callback function pointer list...
	for i, p := range backs {
		// When we find the pointer of the function
		// we are removing...
		if p == pfunc {
			// Remove it from the callback list...
			// So it can be garbage collected if not used
			backs = append(backs[:i], backs[i+1:]...)
			//delete(backs, i);
			return
		}
	}
}

//export go_int_callback_basic
func go_int_callback_basic(pfunc unsafe.Pointer, window unsafe.Pointer, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, POINTER) int)(pfunc)
		return C.int(thisfunc(HWND{window}, POINTER(data)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, POINTER) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, POINTER(data)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, POINTER) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, POINTER(data)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, POINTER(data)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, POINTER(data)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, POINTER) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, POINTER(data)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, POINTER) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, POINTER(data)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, POINTER(data)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, POINTER) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, POINTER(data)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, POINTER) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, POINTER(data)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, POINTER(data)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, POINTER) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, POINTER(data)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, POINTER) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, POINTER(data)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, POINTER) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, POINTER(data)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, POINTER) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, POINTER(data)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, POINTER(data)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, POINTER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, POINTER(data)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, POINTER) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, POINTER(data)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, POINTER) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, POINTER(data)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, POINTER) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, POINTER(data)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND) int)(pfunc)
		return C.int(thisfunc(HWND{window}))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE) int)(pfunc)
		return C.int(thisfunc(HTREE{window}))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE) int)(pfunc)
		return C.int(thisfunc(HMLE{window}))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX) int)(pfunc)
		return C.int(thisfunc(HBOX{window}))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML) int)(pfunc)
		return C.int(thisfunc(HHTML{window}))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}))
	}
	thisfunc := *(*func(HANDLE, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, POINTER(data)))
}

//export go_int_callback_configure
func go_int_callback_configure(pfunc unsafe.Pointer, window unsafe.Pointer, width C.int, height C.int, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(width), int(height), POINTER(data)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(width), int(height), POINTER(data)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(width), int(height), POINTER(data)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(width), int(height), POINTER(data)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int, int, POINTER) C.int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(width), int(height), POINTER(data)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(width), int(height), POINTER(data)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(width), int(height), POINTER(data)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(width), int(height), POINTER(data)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(width), int(height), POINTER(data)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(width), int(height), POINTER(data)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(width), int(height), POINTER(data)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(width), int(height), POINTER(data)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(width), int(height), POINTER(data)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(width), int(height), POINTER(data)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(width), int(height), POINTER(data)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(width), int(height), POINTER(data)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(width), int(height), POINTER(data)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(width), int(height), POINTER(data)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(width), int(height), POINTER(data)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(width), int(height), POINTER(data)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int, int, POINTER) C.int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(width), int(height), POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, int, int) C.int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, int(width), int(height)))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND, int, int) C.int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(width), int(height)))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int, int) C.int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(width), int(height)))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT, int, int) C.int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(width), int(height)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, int, int) C.int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(width), int(height)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int, int) C.int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(width), int(height)))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE, int, int) C.int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(width), int(height)))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON, int, int) C.int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(width), int(height)))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int, int) C.int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(width), int(height)))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int, int) C.int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(width), int(height)))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX, int, int) C.int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(width), int(height)))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int, int) C.int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(width), int(height)))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int, int) C.int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(width), int(height)))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int, int) C.int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(width), int(height)))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT, int, int) C.int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(width), int(height)))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER, int, int) C.int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(width), int(height)))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int, int) C.int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(width), int(height)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, int, int) C.int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(width), int(height)))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML, int, int) C.int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(width), int(height)))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int, int) C.int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(width), int(height)))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP, int, int) C.int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(width), int(height)))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int, int) C.int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(width), int(height)))
	}
	thisfunc := *(*func(HANDLE, int, int, POINTER) C.int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, int(width), int(height), POINTER(data)))
}

//export go_int_callback_keypress
func go_int_callback_keypress(pfunc unsafe.Pointer, window unsafe.Pointer, ch C.char, vk C.int, state C.int, data unsafe.Pointer, utf8 *C.char, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HWND{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, uint8, int, int, POINTER, string) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, uint8, int, int, POINTER, string) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HWND{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, uint8, int, int, string) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, uint8, int, int, string) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, uint8(ch), int(vk), int(state), C.GoString(utf8)))
	}
	thisfunc := *(*func(HANDLE, uint8, int, int, POINTER, string) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, uint8(ch), int(vk), int(state), POINTER(data), C.GoString(utf8)))
}

//export go_int_callback_mouse
func go_int_callback_mouse(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, mask C.int, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(x), int(y), int(mask), POINTER(data)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(x), int(y), int(mask), POINTER(data)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(x), int(y), int(mask), POINTER(data)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(x), int(y), int(mask), POINTER(data)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int, int, int, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(x), int(y), int(mask), POINTER(data)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(x), int(y), int(mask), POINTER(data)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(x), int(y), int(mask), POINTER(data)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(x), int(y), int(mask), POINTER(data)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(x), int(y), int(mask), POINTER(data)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(x), int(y), int(mask), POINTER(data)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(x), int(y), int(mask), POINTER(data)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(x), int(y), int(mask), POINTER(data)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(x), int(y), int(mask), POINTER(data)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(x), int(y), int(mask), POINTER(data)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(x), int(y), int(mask), POINTER(data)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(x), int(y), int(mask), POINTER(data)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(x), int(y), int(mask), POINTER(data)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(x), int(y), int(mask), POINTER(data)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(x), int(y), int(mask), POINTER(data)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(x), int(y), int(mask), POINTER(data)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(x), int(y), int(mask), POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, int, int, int) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, int(x), int(y), int(mask)))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND, int, int, int) int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(x), int(y), int(mask)))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int, int, int) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(x), int(y), int(mask)))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT, int, int, int) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(x), int(y), int(mask)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, int, int, int) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(x), int(y), int(mask)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int, int, int) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(x), int(y), int(mask)))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE, int, int, int) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(x), int(y), int(mask)))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON, int, int, int) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(x), int(y), int(mask)))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int, int, int) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(x), int(y), int(mask)))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int, int, int) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(x), int(y), int(mask)))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX, int, int, int) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(x), int(y), int(mask)))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int, int, int) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(x), int(y), int(mask)))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int, int, int) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(x), int(y), int(mask)))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int, int, int) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(x), int(y), int(mask)))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT, int, int, int) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(x), int(y), int(mask)))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER, int, int, int) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(x), int(y), int(mask)))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int, int, int) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(x), int(y), int(mask)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, int, int, int) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(x), int(y), int(mask)))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML, int, int, int) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(x), int(y), int(mask)))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int, int, int) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(x), int(y), int(mask)))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP, int, int, int) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(x), int(y), int(mask)))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int, int, int) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(x), int(y), int(mask)))
	}
	thisfunc := *(*func(HANDLE, int, int, int, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, int(x), int(y), int(mask), POINTER(data)))
}

//export go_int_callback_expose
func go_int_callback_expose(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, width C.int, height C.int, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, int, int, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(x), int(y), int(width), int(height), POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, int, int, int, int) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, int(x), int(y), int(width), int(height)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, int, int, int, int) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(x), int(y), int(width), int(height)))
	}
	thisfunc := *(*func(HANDLE, int, int, int, int, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, int(x), int(y), int(width), int(height), POINTER(data)))
}

//export go_int_callback_item_enter
func go_int_callback_item_enter(pfunc unsafe.Pointer, window unsafe.Pointer, text *C.char, data unsafe.Pointer, itemdata unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, string, POINTER, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, C.GoString(text), POINTER(data), POINTER(itemdata)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, string, POINTER, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, C.GoString(text), POINTER(data), POINTER(itemdata)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, string, POINTER) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, C.GoString(text), POINTER(itemdata)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, string, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, C.GoString(text), POINTER(itemdata)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, string, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, C.GoString(text), POINTER(itemdata)))
	}
	thisfunc := *(*func(HANDLE, string, POINTER, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, C.GoString(text), POINTER(data), POINTER(itemdata)))
}

//export go_int_callback_item_context
func go_int_callback_item_context(pfunc unsafe.Pointer, window unsafe.Pointer, text *C.char, x C.int, y C.int, data unsafe.Pointer, itemdata unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, string, int, int, POINTER, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, C.GoString(text), int(x), int(y), POINTER(data), POINTER(itemdata)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, string, int, int, POINTER, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, C.GoString(text), int(x), int(y), POINTER(data), POINTER(itemdata)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, string, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, C.GoString(text), int(x), int(y), POINTER(itemdata)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, string, int, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, C.GoString(text), int(x), int(y), POINTER(itemdata)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, string, int, int, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, C.GoString(text), int(x), int(y), POINTER(itemdata)))
	}
	thisfunc := *(*func(HANDLE, string, int, int, POINTER, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, C.GoString(text), int(x), int(y), POINTER(data), POINTER(itemdata)))
}

//export go_int_callback_item_select
func go_int_callback_item_select(pfunc unsafe.Pointer, window unsafe.Pointer, item unsafe.Pointer, text *C.char, data unsafe.Pointer, itemdata unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, HTREEITEM, string, POINTER, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(data), POINTER(itemdata)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, HTREEITEM, string, POINTER, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(data), POINTER(itemdata)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, HTREEITEM, string, POINTER) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(itemdata)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, HTREEITEM, string, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(itemdata)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, HTREEITEM, string, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(itemdata)))
	}
	thisfunc := *(*func(HANDLE, HTREEITEM, string, POINTER, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, HTREEITEM{item, HWND{window}}, C.GoString(text), POINTER(data), POINTER(itemdata)))
}

//export go_int_callback_numeric
func go_int_callback_numeric(pfunc unsafe.Pointer, window unsafe.Pointer, val C.int, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(val), POINTER(data)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(val), POINTER(data)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(val), POINTER(data)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(val), POINTER(data)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(val), POINTER(data)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(val), POINTER(data)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(val), POINTER(data)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(val), POINTER(data)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(val), POINTER(data)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(val), POINTER(data)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(val), POINTER(data)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(val), POINTER(data)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(val), POINTER(data)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(val), POINTER(data)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(val), POINTER(data)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(val), POINTER(data)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(val), POINTER(data)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(val), POINTER(data)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(val), POINTER(data)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(val), POINTER(data)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(val), POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, int) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, int(val)))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND, int) int)(pfunc)
		return C.int(thisfunc(HWND{window}, int(val)))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, int) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, int(val)))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT, int) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, int(val)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, int) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, int(val)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, int) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, int(val)))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE, int) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, int(val)))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON, int) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, int(val)))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, int) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, int(val)))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, int) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, int(val)))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX, int) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, int(val)))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, int) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, int(val)))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM, int) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, int(val)))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX, int) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, int(val)))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT, int) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, int(val)))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER, int) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, int(val)))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, int) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, int(val)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, int) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, int(val)))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML, int) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, int(val)))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR, int) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, int(val)))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP, int) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, int(val)))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, int) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, int(val)))
	}
	thisfunc := *(*func(HANDLE, int, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, int(val), POINTER(data)))
}

//export go_int_callback_ulong
func go_int_callback_ulong(pfunc unsafe.Pointer, window unsafe.Pointer, val C.ulong, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (1 << 8): // HWND
		thisfunc := *(*func(HWND, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HWND{window}, uint(val), POINTER(data)))
	case (2 << 8): // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, uint(val), POINTER(data)))
	case (3 << 8): // HTEXT
		thisfunc := *(*func(HTEXT, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, uint(val), POINTER(data)))
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, uint(val), POINTER(data)))
	case (5 << 8): // HCONTAINER
		thisfunc := *(*func(HCONTAINER, uint, POINTER) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, uint(val), POINTER(data)))
	case (6 << 8): // HMLE
		thisfunc := *(*func(HMLE, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, uint(val), POINTER(data)))
	case (7 << 8): // HBUTTON
		thisfunc := *(*func(HBUTTON, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, uint(val), POINTER(data)))
	case (8 << 8): // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, uint(val), POINTER(data)))
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, uint(val), POINTER(data)))
	case (10 << 8): // HBOX
		thisfunc := *(*func(HBOX, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, uint(val), POINTER(data)))
	case (11 << 8): // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, uint(val), POINTER(data)))
	case (12 << 8): // HMENUITEM
		thisfunc := *(*func(HMENUITEM, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, uint(val), POINTER(data)))
	case (13 << 8): // HLISTBOX
		thisfunc := *(*func(HLISTBOX, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, uint(val), POINTER(data)))
	case (14 << 8): // HPERCENT
		thisfunc := *(*func(HPERCENT, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, uint(val), POINTER(data)))
	case (15 << 8): // HSLIDER
		thisfunc := *(*func(HSLIDER, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, uint(val), POINTER(data)))
	case (16 << 8): // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, uint(val), POINTER(data)))
	case (17 << 8): // HRENDER
		thisfunc := *(*func(HRENDER, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, uint(val), POINTER(data)))
	case (18 << 8): // HHTML
		thisfunc := *(*func(HHTML, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, uint(val), POINTER(data)))
	case (19 << 8): // HCALENDAR
		thisfunc := *(*func(HCALENDAR, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, uint(val), POINTER(data)))
	case (20 << 8): // HBITMAP
		thisfunc := *(*func(HBITMAP, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, uint(val), POINTER(data)))
	case (21 << 8): // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, uint, POINTER) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, uint(val), POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, uint) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, uint(val)))
	case (1 << 8) | go_flags_no_data: // HWND
		thisfunc := *(*func(HWND, uint) int)(pfunc)
		return C.int(thisfunc(HWND{window}, uint(val)))
	case (2 << 8) | go_flags_no_data: // HENTRYFIELD
		thisfunc := *(*func(HENTRYFIELD, uint) int)(pfunc)
		return C.int(thisfunc(HENTRYFIELD{window}, uint(val)))
	case (3 << 8) | go_flags_no_data: // HTEXT
		thisfunc := *(*func(HTEXT, uint) int)(pfunc)
		return C.int(thisfunc(HTEXT{window}, uint(val)))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, uint) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, uint(val)))
	case (5 << 8) | go_flags_no_data: // HCONTAINER
		thisfunc := *(*func(HCONTAINER, uint) int)(pfunc)
		filesystem := false
		if Window_get_data(HCONTAINER{window, false}, "_go_filesystem") != nil {
			filesystem = true
		}
		return C.int(thisfunc(HCONTAINER{window, filesystem}, uint(val)))
	case (6 << 8) | go_flags_no_data: // HMLE
		thisfunc := *(*func(HMLE, uint) int)(pfunc)
		return C.int(thisfunc(HMLE{window}, uint(val)))
	case (7 << 8) | go_flags_no_data: // HBUTTON
		thisfunc := *(*func(HBUTTON, uint) int)(pfunc)
		return C.int(thisfunc(HBUTTON{window}, uint(val)))
	case (8 << 8) | go_flags_no_data: // HSPINBUTTON
		thisfunc := *(*func(HSPINBUTTON, uint) int)(pfunc)
		return C.int(thisfunc(HSPINBUTTON{window}, uint(val)))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, uint) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, uint(val)))
	case (10 << 8) | go_flags_no_data: // HBOX
		thisfunc := *(*func(HBOX, uint) int)(pfunc)
		return C.int(thisfunc(HBOX{window}, uint(val)))
	case (11 << 8) | go_flags_no_data: // HSCROLLBOX
		thisfunc := *(*func(HSCROLLBOX, uint) int)(pfunc)
		return C.int(thisfunc(HSCROLLBOX{window}, uint(val)))
	case (12 << 8) | go_flags_no_data: // HMENUITEM
		thisfunc := *(*func(HMENUITEM, uint) int)(pfunc)
		return C.int(thisfunc(HMENUITEM{window}, uint(val)))
	case (13 << 8) | go_flags_no_data: // HLISTBOX
		thisfunc := *(*func(HLISTBOX, uint) int)(pfunc)
		return C.int(thisfunc(HLISTBOX{window}, uint(val)))
	case (14 << 8) | go_flags_no_data: // HPERCENT
		thisfunc := *(*func(HPERCENT, uint) int)(pfunc)
		return C.int(thisfunc(HPERCENT{window}, uint(val)))
	case (15 << 8) | go_flags_no_data: // HSLIDER
		thisfunc := *(*func(HSLIDER, uint) int)(pfunc)
		return C.int(thisfunc(HSLIDER{window}, uint(val)))
	case (16 << 8) | go_flags_no_data: // HSCROLLBAR
		thisfunc := *(*func(HSCROLLBAR, uint) int)(pfunc)
		return C.int(thisfunc(HSCROLLBAR{window}, uint(val)))
	case (17 << 8) | go_flags_no_data: // HRENDER
		thisfunc := *(*func(HRENDER, uint) int)(pfunc)
		return C.int(thisfunc(HRENDER{window}, uint(val)))
	case (18 << 8) | go_flags_no_data: // HHTML
		thisfunc := *(*func(HHTML, uint) int)(pfunc)
		return C.int(thisfunc(HHTML{window}, uint(val)))
	case (19 << 8) | go_flags_no_data: // HCALENDAR
		thisfunc := *(*func(HCALENDAR, uint) int)(pfunc)
		return C.int(thisfunc(HCALENDAR{window}, uint(val)))
	case (20 << 8) | go_flags_no_data: // HBITMAP
		thisfunc := *(*func(HBITMAP, uint) int)(pfunc)
		return C.int(thisfunc(HBITMAP{window}, uint(val)))
	case (21 << 8) | go_flags_no_data: // HSPLITBAR
		thisfunc := *(*func(HSPLITBAR, uint) int)(pfunc)
		return C.int(thisfunc(HSPLITBAR{window}, uint(val)))
	}
	thisfunc := *(*func(HANDLE, uint, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, uint(val), POINTER(data)))
}

//export go_int_callback_notepage
func go_int_callback_notepage(pfunc unsafe.Pointer, window unsafe.Pointer, val C.ulong, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (9 << 8): // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, HNOTEPAGE, POINTER) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, HNOTEPAGE{val, HNOTEBOOK{window}}, POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, HNOTEPAGE) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, HNOTEPAGE{val, HNOTEBOOK{window}}))
	case (9 << 8) | go_flags_no_data: // HNOTEBOOK
		thisfunc := *(*func(HNOTEBOOK, HNOTEPAGE) int)(pfunc)
		return C.int(thisfunc(HNOTEBOOK{window}, HNOTEPAGE{val, HNOTEBOOK{window}}))
	}
	thisfunc := *(*func(HANDLE, HNOTEPAGE, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, HNOTEPAGE{val, HNOTEBOOK{window}}, POINTER(data)))
}

//export go_int_callback_tree
func go_int_callback_tree(pfunc unsafe.Pointer, window unsafe.Pointer, tree unsafe.Pointer, data unsafe.Pointer, flags C.uint) C.int {
	switch flags {
	case (4 << 8): // HTREE
		thisfunc := *(*func(HTREE, HTREEITEM, POINTER) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, HTREEITEM{tree, HWND{window}}, POINTER(data)))
	case go_flags_no_data:
		thisfunc := *(*func(HANDLE, HTREEITEM) int)(pfunc)
		return C.int(thisfunc(HGENERIC{window}, HTREEITEM{tree, HWND{window}}))
	case (4 << 8) | go_flags_no_data: // HTREE
		thisfunc := *(*func(HTREE, HTREEITEM) int)(pfunc)
		return C.int(thisfunc(HTREE{window}, HTREEITEM{tree, HWND{window}}))
	}
	thisfunc := *(*func(HANDLE, HTREEITEM, POINTER) int)(pfunc)
	return C.int(thisfunc(HGENERIC{window}, HTREEITEM{tree, HWND{window}}, POINTER(data)))
}

//export go_int_callback_timer
func go_int_callback_timer(pfunc unsafe.Pointer, data unsafe.Pointer, flags C.uint) C.int {
	if (flags & go_flags_no_data) == go_flags_no_data {
		thisfunc := *(*func() int)(pfunc)
		return C.int(thisfunc())
	}
	thisfunc := *(*func(POINTER) int)(pfunc)
	return C.int(thisfunc(POINTER(data)))
}

//export go_int_callback_print
func go_int_callback_print(pfunc unsafe.Pointer, print unsafe.Pointer, pixmap unsafe.Pointer, page_num C.int, data unsafe.Pointer, flags C.uint) C.int {
	if (flags & go_flags_no_data) == go_flags_no_data {
		thisfunc := *(*func(HPRINT, HPIXMAP, int) int)(pfunc)
		return C.int(thisfunc(HPRINT{print, ""}, HPIXMAP{pixmap}, int(page_num)))
	}
	thisfunc := *(*func(HPRINT, HPIXMAP, int, POINTER) int)(pfunc)
	return C.int(thisfunc(HPRINT{print, ""}, HPIXMAP{pixmap}, int(page_num), POINTER(data)))
}
