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
type HTIMER C.int
type HMENUI unsafe.Pointer
type HPIXMAP unsafe.Pointer
type HNOTEPAGE C.ulong
type COLOR C.ulong

type Env struct {
    OSName, BuildDate, BuildTime string
    MajorVersion, MinorVersion, MajorBuild, MinorBuild C.short
    DWMajorVersion, DWMinorVersion, DWSubVersion C.short
}

// Define our exported constants
const (
   FALSE int = iota
   TRUE
)

var DESKTOP HWND = nil
var NOMENU HMENUI = nil

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

var POINTER_DEFAULT = C.DW_POINTER_DEFAULT
var POINTER_ARROW = C.DW_POINTER_ARROW
var POINTER_CLOCK = C.DW_POINTER_CLOCK
var POINTER_QUESTION = C.DW_POINTER_QUESTION

var DT_LEFT uint = C.DW_DT_LEFT
var DT_CENTER uint = C.DW_DT_CENTER
var DT_RIGHT uint = C.DW_DT_RIGHT
var DT_VCENTER uint = C.DW_DT_VCENTER
var DT_WORDBREAK uint = C.DW_DT_WORDBREAK

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

var BS_NOBORDER uint = C.DW_BS_NOBORDER

var KC_CTRL = C.KC_CTRL
var KC_SHIFT = C.KC_SHIFT
var KC_ALT = C.KC_ALT

var MENU_SEPARATOR = C.DW_MENU_SEPARATOR
var MENU_AUTO uint = C.DW_MENU_AUTO
var MENU_POPUP uint = ^uint(0)

var PERCENT_INDETERMINATE = -1

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

var BUTTON1_MASK = C.DW_BUTTON1_MASK
var BUTTON2_MASK = C.DW_BUTTON2_MASK
var BUTTON3_MASK = C.DW_BUTTON3_MASK

var FILE_OPEN = C.DW_FILE_OPEN
var FILE_SAVE = C.DW_FILE_SAVE
var DIRECTORY_OPEN = C.DW_DIRECTORY_OPEN

var VK_LBUTTON  = int(C.VK_LBUTTON)
var VK_RBUTTON  = int(C.VK_RBUTTON)
var VK_CANCEL   = int(C.VK_CANCEL)
var VK_MBUTTON  = int(C.VK_MBUTTON)
var VK_TAB      = int(C.VK_TAB)
var VK_CLEAR    = int(C.VK_CLEAR)
var VK_RETURN   = int(C.VK_RETURN)
var VK_PAUSE    = int(C.VK_PAUSE)
var VK_CAPITAL  = int(C.VK_CAPITAL)
var VK_ESCAPE   = int(C.VK_ESCAPE)
var VK_SPACE    = int(C.VK_SPACE)
var VK_PRIOR    = int(C.VK_PRIOR)
var VK_NEXT     = int(C.VK_NEXT)
var VK_END      = int(C.VK_END)
var VK_HOME     = int(C.VK_HOME)
var VK_LEFT     = int(C.VK_LEFT)
var VK_UP       = int(C.VK_UP)
var VK_RIGHT    = int(C.VK_RIGHT)
var VK_DOWN     = int(C.VK_DOWN)
var VK_SELECT   = int(C.VK_SELECT)
var VK_PRINT    = int(C.VK_PRINT)
var VK_EXECUTE  = int(C.VK_EXECUTE)
var VK_SNAPSHOT = int(C.VK_SNAPSHOT)
var VK_INSERT   = int(C.VK_INSERT)
var VK_DELETE   = int(C.VK_DELETE)
var VK_HELP     = int(C.VK_HELP)
var VK_LWIN     = int(C.VK_LWIN)
var VK_RWIN     = int(C.VK_RWIN)
var VK_NUMPAD0  = int(C.VK_NUMPAD0)
var VK_NUMPAD1  = int(C.VK_NUMPAD1)
var VK_NUMPAD2  = int(C.VK_NUMPAD2)
var VK_NUMPAD3  = int(C.VK_NUMPAD3)
var VK_NUMPAD4  = int(C.VK_NUMPAD4)
var VK_NUMPAD5  = int(C.VK_NUMPAD5)
var VK_NUMPAD6  = int(C.VK_NUMPAD6)
var VK_NUMPAD7  = int(C.VK_NUMPAD7)
var VK_NUMPAD8  = int(C.VK_NUMPAD8)
var VK_NUMPAD9  = int(C.VK_NUMPAD9)
var VK_MULTIPLY = int(C.VK_MULTIPLY)
var VK_ADD      = int(C.VK_ADD)
var VK_SEPARATOR= int(C.VK_SEPARATOR)
var VK_SUBTRACT = int(C.VK_SUBTRACT)
var VK_DECIMAL  = int(C.VK_DECIMAL)
var VK_DIVIDE   = int(C.VK_DIVIDE)
var VK_F1       = int(C.VK_F1)
var VK_F2       = int(C.VK_F2)
var VK_F3       = int(C.VK_F3)
var VK_F4       = int(C.VK_F4)
var VK_F5       = int(C.VK_F5)
var VK_F6       = int(C.VK_F6)
var VK_F7       = int(C.VK_F7)
var VK_F8       = int(C.VK_F8)
var VK_F9       = int(C.VK_F9)
var VK_F10      = int(C.VK_F10)
var VK_F11      = int(C.VK_F11)
var VK_F12      = int(C.VK_F12)
var VK_F13      = int(C.VK_F13)
var VK_F14      = int(C.VK_F14)
var VK_F15      = int(C.VK_F15)
var VK_F16      = int(C.VK_F16)
var VK_F17      = int(C.VK_F17)
var VK_F18      = int(C.VK_F18)
var VK_F19      = int(C.VK_F19)
var VK_F20      = int(C.VK_F20)
var VK_F21      = int(C.VK_F21)
var VK_F22      = int(C.VK_F22)
var VK_F23      = int(C.VK_F23)
var VK_F24      = int(C.VK_F24)
var VK_NUMLOCK  = int(C.VK_NUMLOCK)
var VK_SCROLL   = int(C.VK_SCROLL)
var VK_LSHIFT   = int(C.VK_LSHIFT)
var VK_RSHIFT   = int(C.VK_RSHIFT)
var VK_LCONTROL = int(C.VK_LCONTROL)
var VK_RCONTROL = int(C.VK_RCONTROL)

func RESOURCE(id uintptr) unsafe.Pointer {
   return unsafe.Pointer(id);
}

func RGB(red uint8, green uint8, blue uint8) COLOR {
    lred := C.ulong(red);
    lgreen := C.ulong(green);
    lblue := C.ulong(blue);
    return COLOR((0xF0000000 | (lred) | (lgreen << 8) | (lblue << 16)));
}

func Init(newthread int) int {
   return int(C.go_init(C.int(newthread)));
}

func Environment_query(env *Env) {
    var cenv C.DWEnv;
    C.dw_environment_query(&cenv);
    env.OSName = C.GoString((*C.char)(unsafe.Pointer(&cenv.osName[0])));
    env.BuildDate = C.GoString((*C.char)(unsafe.Pointer(&cenv.buildDate[0])));
    env.BuildTime = C.GoString((*C.char)(unsafe.Pointer(&cenv.buildTime[0])));
    env.MajorVersion = cenv.MajorVersion;
    env.MinorVersion = cenv.MajorVersion;
    env.MajorBuild = cenv.MajorBuild;
    env.MinorBuild = cenv.MinorBuild;
    env.DWMajorVersion = cenv.DWMajorVersion;
    env.DWMinorVersion = cenv.DWMinorVersion;
    env.DWSubVersion = cenv.DWSubVersion;
}

func Messagebox(title string, flags int, message string) int {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cmessage := C.CString(message);
   defer C.free(unsafe.Pointer(cmessage));
   
   return int(C.go_messagebox(ctitle, C.int(flags), cmessage));
}

func Window_new(owner HWND, title string, flags uint) HWND {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   
   return HWND(C.go_window_new(unsafe.Pointer(owner), ctitle, C.ulong(flags)));
}

func Window_show(handle HWND) int {
   return int(C.go_window_show(unsafe.Pointer(handle)));
}

func Window_hide(handle HWND) int {
   return int(C.go_window_hide(unsafe.Pointer(handle)));
}

func Window_lower(handle HWND) int {
   return int(C.go_window_lower(unsafe.Pointer(handle)));
}

func Window_raise(handle HWND) int {
   return int(C.go_window_raise(unsafe.Pointer(handle)));
}

func Window_minimize(handle HWND) int {
   return int(C.go_window_minimize(unsafe.Pointer(handle)));
}

func Window_set_pos(handle HWND, x int, y int) {
   C.go_window_set_pos(unsafe.Pointer(handle), C.long(x), C.long(y));
}

func Window_set_pos_size(handle HWND, x int, y int, width uint, height uint) {
   C.go_window_set_pos_size(unsafe.Pointer(handle), C.long(x), C.long(y), C.ulong(width), C.ulong(height));
}

func Window_set_size(handle HWND, width uint, height uint) {
   C.go_window_set_size(unsafe.Pointer(handle), C.ulong(width), C.ulong(height));
}

func Window_set_color(handle HWND, fore COLOR, back COLOR) int {
   return int(C.go_window_set_color(unsafe.Pointer(handle), C.ulong(fore), C.ulong(back)));
}

func Window_set_style(handle HWND, style uint, mask uint) {
   C.go_window_set_style(unsafe.Pointer(handle), C.ulong(style), C.ulong(mask));
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

func Window_from_id(handle HWND, cid int) HWND {
   return HWND(C.go_window_from_id(unsafe.Pointer(handle), C.int(cid)));
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

func Window_set_font(handle HWND, fontname string) int {
   cfontname := C.CString(fontname);
   defer C.free(unsafe.Pointer(cfontname));
   
   return int(C.go_window_set_font(unsafe.Pointer(handle), cfontname));
}

func Window_get_pos_size(handle HWND) (int, int, uint, uint) {
   var x, y C.long;
   var width, height C.ulong;
   C.go_window_get_pos_size(unsafe.Pointer(handle), &x, &y, &width, &height);
   return int(x), int(y), uint(width), uint(height);
}

func Window_get_preferred_size(handle HWND) (int, int) {
   var width, height C.int;
   C.go_window_get_preferred_size(unsafe.Pointer(handle), &width, &height);
   return int(width), int(height);
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

func Window_set_bitmap(window HWND, id uint, filename string) {
   cfilename := C.CString(filename);
   defer C.free(unsafe.Pointer(cfilename));
   
   C.go_window_set_bitmap(unsafe.Pointer(window), C.ulong(id), cfilename);
}

func Window_set_border(handle HWND, border int) {
   C.go_window_set_border(unsafe.Pointer(handle), C.int(border));
}

func Window_set_focus(handle HWND) {
   C.go_window_set_focus(unsafe.Pointer(handle));
}

func Window_set_gravity(handle HWND, horz int, vert int) {
   C.go_window_set_gravity(unsafe.Pointer(handle), C.int(horz), C.int(vert));
}

func Window_set_icon(handle HWND, icon HICN) {
   C.go_window_set_icon(unsafe.Pointer(handle), unsafe.Pointer(icon));
}

func Window_set_pointer(handle HWND, cursortype int) {
   C.go_window_set_pointer(unsafe.Pointer(handle), C.int(cursortype));
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

func Main_sleep(milliseconds int) {
   C.dw_main_sleep(C.int(milliseconds));
}

func Box_new(btype int, pad int) HWND {
   return HWND(C.go_box_new(C.int(btype), C.int(pad)));
}

func Box_pack_at_index(box HWND, item HWND, index int, width int, height int, hsize int, vsize int, pad int) {
   C.go_box_pack_at_index(unsafe.Pointer(box), unsafe.Pointer(item), C.int(index), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad));
}

func Box_pack_end(box HWND, item HWND, width int, height int, hsize int, vsize int, pad int) {
   C.go_box_pack_end(unsafe.Pointer(box), unsafe.Pointer(item), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad));
}

func Box_pack_start(box HWND, item HWND, width int, height int, hsize int, vsize int, pad int) {
   C.go_box_pack_start(unsafe.Pointer(box), unsafe.Pointer(item), C.int(width), C.int(height), C.int(hsize), C.int(vsize), C.int(pad));
}

func Box_unpack(handle HWND) int {
   return int(C.go_box_unpack(unsafe.Pointer(handle)));
}

func Box_unpack_at_index(handle HWND, index int) HWND {
   return HWND(C.go_box_unpack_at_index(unsafe.Pointer(handle), C.int(index)));
}

func Text_new(text string, id uint) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_text_new(ctext, C.ulong(id)));
}

func Status_text_new(text string, id uint) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_status_text_new(ctext, C.ulong(id)));
}

func Entryfield_new(text string, id uint) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_new(ctext, C.ulong(id)));
}

func Entryfield_password_new(text string, id uint) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_entryfield_password_new(ctext, C.ulong(id)));
}

func Entryfield_set_limit(handle HWND, limit int) {
   C.go_entryfield_set_limit(unsafe.Pointer(handle), C.int(limit));
}

func Button_new(text string, id uint) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_button_new(ctext, C.ulong(id)));
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

func File_browse(title string, defpath string, ext string, flags int) string {
   ctitle := C.CString(title);
   defer C.free(unsafe.Pointer(ctitle));
   cdefpath := C.CString(defpath);
   defer C.free(unsafe.Pointer(cdefpath));
   cext := C.CString(ext);
   defer C.free(unsafe.Pointer(cext));
   
   result := C.dw_file_browse(ctitle, cdefpath, cext, C.int(flags));
   defer C.dw_free(unsafe.Pointer(result));
   return C.GoString(result);
}

func Color_choose(value COLOR) COLOR {
   return COLOR(C.dw_color_choose(C.ulong(value)));
}

func Timer_connect(interval int, sigfunc unsafe.Pointer, data unsafe.Pointer) HTIMER {
   return HTIMER(C.go_timer_connect(C.int(interval), sigfunc, data));
}

func Timer_disconnect(id HTIMER) {
   C.dw_timer_disconnect(C.int(id));
}

func Signal_connect(window HWND, signame string, sigfunc unsafe.Pointer, data unsafe.Pointer) {
   csigname := C.CString(signame);
   defer C.free(unsafe.Pointer(csigname));
   
   C.go_signal_connect(unsafe.Pointer(window), csigname, sigfunc, data);
}

func Beep(freq int, dur int) {
    C.dw_beep(C.int(freq), C.int(dur));
}

func Menu_new(id uint) HMENUI {
    return HMENUI(C.go_menu_new(C.ulong(id)));
}

func Menubar_new(location HWND) HMENUI {
    return HMENUI(C.go_menubar_new(unsafe.Pointer(location)));
}

func Menu_append_item(menu HMENUI, title string, id uint, flags uint, end int, check int, submenu HMENUI) HWND {
    ctitle := C.CString(title);
    defer C.free(unsafe.Pointer(ctitle));

    return HWND(C.go_menu_append_item(unsafe.Pointer(menu), ctitle, C.ulong(id), C.ulong(flags), C.int(end), C.int(check), unsafe.Pointer(submenu)));
}

func Menu_delete_item(menu HMENUI, id uint) {
    C.go_menu_delete_item(unsafe.Pointer(menu), C.ulong(id));
}

func Menu_destroy(menu HMENUI) {
    C.go_menu_destroy(unsafe.Pointer(menu));
}

func Menu_item_set_state(menu HMENUI, id uint, flags uint) {
    C.go_menu_item_set_state(unsafe.Pointer(menu), C.ulong(id), C.ulong(flags));
}

func Menu_popup(menu HMENUI, parent HWND, x int, y int) {
    C.go_menu_popup(unsafe.Pointer(menu), unsafe.Pointer(parent), C.int(x), C.int(y));
}

func Notebook_new(id uint, top int) HWND {
    return HWND(C.go_notebook_new(C.ulong(id), C.int(top)));
}

func Notebook_pack(handle HWND, pageid HNOTEPAGE, page HWND) {
    C.go_notebook_pack(unsafe.Pointer(handle), C.ulong(pageid), unsafe.Pointer(page));
}

func Notebook_page_destroy(handle HWND, pageid HNOTEPAGE) {
    C.go_notebook_page_destroy(unsafe.Pointer(handle), C.ulong(pageid));
}

func Notebook_page_get(handle HWND) HNOTEPAGE {
    return HNOTEPAGE(C.go_notebook_page_get(unsafe.Pointer(handle)));
}

func Notebook_page_new(handle HWND, flags uint, front int) HNOTEPAGE {
    return HNOTEPAGE(C.go_notebook_page_new(unsafe.Pointer(handle), C.ulong(flags), C.int(front)));
}

func Notebook_page_set(handle HWND, pageid HNOTEPAGE) {
    C.go_notebook_page_set(unsafe.Pointer(handle), C.ulong(pageid));
}

func Notebook_page_set_text(handle HWND, pageid HNOTEPAGE, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_notebook_page_set_text(unsafe.Pointer(handle), C.ulong(pageid), ctext);
}

func Icon_load_from_file(filename string) HICN {
    cfilename := C.CString(filename);
    defer C.free(unsafe.Pointer(cfilename));
    
    return HICN(C.go_icon_load_from_file(cfilename));
}

func Icon_load(id uint) HICN {
    return HICN(C.go_icon_load(0, C.ulong(id)));
}

func Taskbar_delete(handle HWND, icon HICN) {
    C.go_taskbar_delete(unsafe.Pointer(handle), unsafe.Pointer(icon));
}

func Taskbar_insert(handle HWND, icon HICN, bubbletext string) {
    cbubbletext := C.CString(bubbletext);
    defer C.free(unsafe.Pointer(cbubbletext));
    
    C.go_taskbar_insert(unsafe.Pointer(handle), unsafe.Pointer(icon), cbubbletext);
}

func Combobox_new(text string, id uint) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_combobox_new(ctext, C.ulong(id)));
}

func Listbox_new(id uint, multi int) HWND {
    return HWND(C.go_listbox_new(C.ulong(id), C.int(multi)));
}

func Listbox_append(handle HWND, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_append(unsafe.Pointer(handle), ctext);
}

func Listbox_insert(handle HWND, text string, pos int) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_insert(unsafe.Pointer(handle), ctext, C.int(pos));
}

func Listbox_clear(handle HWND) {
    C.go_listbox_clear(unsafe.Pointer(handle));
}

func Listbox_count(handle HWND) int {
    return int(C.go_listbox_count(unsafe.Pointer(handle)));
}

func Listbox_set_top(handle HWND, top int) {
    C.go_listbox_set_top(unsafe.Pointer(handle), C.int(top));
}

func Listbox_select(handle HWND, index C.int, state C.int) {
    C.go_listbox_select(unsafe.Pointer(handle), index, state);
}

func Listbox_delete(handle HWND, index int) {
    C.go_listbox_delete(unsafe.Pointer(handle), C.int(index));
}

func Listbox_get_text(handle HWND, index int) string {
    var buf [201]C.char;
    
    C.go_listbox_get_text(unsafe.Pointer(handle), C.int(index), &buf[0], 200);
    return C.GoString((*C.char)(unsafe.Pointer(&buf[0])));
}

func Listbox_set_text(handle HWND, index int, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_set_text(unsafe.Pointer(handle), C.int(index), ctext);
}

func Listbox_selected(handle HWND) int {
    return int(C.go_listbox_selected(unsafe.Pointer(handle)));
}

func Listbox_selected_multi(handle HWND, where int) int {
    return int(C.go_listbox_selected_multi(unsafe.Pointer(handle), C.int(where)));
}

func Screen_width() int {
    return int(C.dw_screen_width());
}

func Screen_height() int {
    return int(C.dw_screen_height());
}

func Color_depth_get() uint {
    return uint(C.dw_color_depth_get());
}

func Color_foreground_set(color COLOR) {
    C.dw_color_foreground_set(C.ulong(color));
}

func Color_background_set(color COLOR) {
    C.dw_color_background_set(C.ulong(color));
}

func Spinbutton_new(text string, id C.ulong) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_spinbutton_new(ctext, id));
}

func Spinbutton_set_pos(handle HWND, position int) {
    C.go_spinbutton_set_pos(unsafe.Pointer(handle), C.long(position));
}

func Spinbutton_set_limits(handle HWND, upper int, lower int) {
    C.go_spinbutton_set_limits(unsafe.Pointer(handle), C.long(upper), C.long(lower));
}

func Spinbutton_get_pos(handle HWND) int {
    return int(C.go_spinbutton_get_pos(unsafe.Pointer(handle)));
}

func Radiobutton_new(text string, id uint) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_radiobutton_new(ctext, C.ulong(id)));
}

func Checkbox_new(text string, id uint) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_checkbox_new(ctext, C.ulong(id)));
}

func Checkbox_get(handle HWND) int {
    return int(C.go_checkbox_get(unsafe.Pointer(handle)));
}

func Checkbox_set(handle HWND, value int) {
    C.go_checkbox_set(unsafe.Pointer(handle), C.int(value));
}

func Percent_new(id C.ulong) HWND {
    return HWND(C.go_percent_new(id));
}

func Slider_new(vertical int, increments int, id uint) HWND {
    return HWND(C.go_slider_new(C.int(vertical), C.int(increments), C.ulong(id)));
}

func Scrollbar_new(vertical int, id uint) HWND {
    return HWND(C.go_scrollbar_new(C.int(vertical), C.ulong(id)));
}

func Slider_get_pos(handle HWND) uint {
    return uint(C.go_slider_get_pos(unsafe.Pointer(handle)));
}

func Slider_set_pos(handle HWND, position uint) {
    C.go_slider_set_pos(unsafe.Pointer(handle), C.uint(position));
}

func Scrollbar_get_pos(handle HWND) uint {
    return uint(C.go_scrollbar_get_pos(unsafe.Pointer(handle)));
}

func Scrollbar_set_pos(handle HWND, position uint) {
    C.go_scrollbar_set_pos(unsafe.Pointer(handle), C.uint(position));
}

func Scrollbar_set_range(handle HWND, srange uint, visible uint) {
    C.go_scrollbar_set_range(unsafe.Pointer(handle), C.uint(srange), C.uint(visible));
}

func Scrollbox_new(btype int, pad int) HWND {
    return HWND(C.go_scrollbox_new(C.int(btype), C.int(pad)));
}

func Scrollbox_get_pos(handle HWND, orient int) int {
    return int(C.go_scrollbox_get_pos(unsafe.Pointer(handle), C.int(orient)));
}

func Scrollbox_get_range(handle HWND, orient int) int {
    return int(C.go_scrollbox_get_range(unsafe.Pointer(handle), C.int(orient)));
}

func Groupbox_new(btype C.int, pad int, title string) HWND {
    ctitle := C.CString(title);
    defer C.free(unsafe.Pointer(ctitle));
    
    return HWND(C.go_groupbox_new(btype, C.int(pad), ctitle));
}

func Render_new(id uint) HWND {
    return HWND(C.go_render_new(C.ulong(id)));
}

func Font_choose(currfont string) string {
    ccurrfont := C.CString(currfont);
    defer C.free(unsafe.Pointer(ccurrfont));
    newfont := C.dw_font_choose(ccurrfont);
    defer C.dw_free(unsafe.Pointer(newfont));
    return C.GoString(newfont);
}

func Font_set_default(fontname string) {
    cfontname := C.CString(fontname);
    defer C.free(unsafe.Pointer(cfontname));
    C.dw_font_set_default(cfontname);
}

func Font_text_extents_get(handle HWND, pixmap HPIXMAP, text string) (int, int) {
   var width, height C.int;
   
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   C.go_font_text_extents_get(unsafe.Pointer(handle), unsafe.Pointer(pixmap), ctext, &width, &height);
   return int(width), int(height);
}

func Pixmap_new(handle HWND, width uint, height uint, depth uint) HPIXMAP {
    return HPIXMAP(C.go_pixmap_new(unsafe.Pointer(handle), C.ulong(width), C.ulong(height), C.ulong(depth)));
}

func Pixmap_new_from_file(handle HWND, filename string) HPIXMAP {
    cfilename := C.CString(filename);
    defer C.free(unsafe.Pointer(cfilename));
    
    return HPIXMAP(C.go_pixmap_new_from_file(unsafe.Pointer(handle), cfilename));
}

func Pixmap_grab(handle HWND, id uint) HPIXMAP {
    return HPIXMAP(C.go_pixmap_grab(unsafe.Pointer(handle), C.ulong(id)));
}

func Pixmap_bitblt(dest HWND, destp HPIXMAP, xdest int, ydest int, width int, height int, src HWND, srcp HPIXMAP, xsrc int, ysrc int) {
    C.go_pixmap_bitblt(unsafe.Pointer(dest), unsafe.Pointer(srcp), C.int(xdest), C.int(ydest), C.int(width), C.int(height), unsafe.Pointer(src), unsafe.Pointer(srcp), C.int(xsrc), C.int(ysrc)); 
}

func Pixmap_stretch_bitblt(dest HWND, destp HPIXMAP, xdest int, ydest int, width int, height int, src HWND, srcp HPIXMAP, xsrc int, ysrc int, srcwidth int, srcheight int) C.int {
    return C.go_pixmap_stretch_bitblt(unsafe.Pointer(dest), unsafe.Pointer(srcp), C.int(xdest), C.int(ydest), C.int(width), C.int(height), unsafe.Pointer(src), unsafe.Pointer(srcp), C.int(xsrc), C.int(ysrc), C.int(srcwidth), C.int(srcheight)); 
}

func Pixmap_set_transparent_color(pixmap HPIXMAP, color COLOR) {
    C.go_pixmap_set_transparent_color(unsafe.Pointer(pixmap), C.ulong(color));
}

func Pixmap_set_font(handle HWND, fontname string) int {
    cfontname := C.CString(fontname);
    defer C.free(unsafe.Pointer(cfontname));
    
    return int(C.go_pixmap_set_font(unsafe.Pointer(handle), cfontname));
}

func Pixmap_destroy(pixmap HPIXMAP) {
    C.go_pixmap_destroy(unsafe.Pointer(pixmap));
}

func Pixmap_width(pixmap HPIXMAP) int {
    return int(C.go_pixmap_width(unsafe.Pointer(pixmap)));
}

func Pixmap_height(pixmap HPIXMAP) int {
    return int(C.go_pixmap_height(unsafe.Pointer(pixmap)));
}

func Draw_point(handle HWND, pixmap HPIXMAP, x int, y int) {
    C.go_draw_point(unsafe.Pointer(handle), unsafe.Pointer(pixmap), C.int(x), C.int(y));
}

func Draw_line(handle HWND, pixmap HPIXMAP, x1 int, y1 int, x2 int, y2 int) {
    C.go_draw_line(unsafe.Pointer(handle), unsafe.Pointer(pixmap), C.int(x1), C.int(y1), C.int(x2), C.int(y2));
}

func Draw_rect(handle HWND, pixmap HPIXMAP, fill int, x int, y int, width int, height int) {
    C.go_draw_rect(unsafe.Pointer(handle), unsafe.Pointer(pixmap), C.int(fill), C.int(x), C.int(y), C.int(width), C.int(height));
}

func Draw_arc(handle HWND, pixmap HPIXMAP, flags int, xorigin int, yorigin int, x1 int, y1 int, x2 int, y2 int) {
    C.go_draw_arc(unsafe.Pointer(handle), unsafe.Pointer(pixmap), C.int(flags), C.int(xorigin), C.int(yorigin), C.int(x1), C.int(y1), C.int(x2), C.int(y2));
}

func Draw_text(handle HWND, pixmap HPIXMAP, x int, y int, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_draw_text(unsafe.Pointer(handle), unsafe.Pointer(pixmap), C.int(x), C.int(y), ctext);
}

func Pointer_query_pos() (int, int) {
   var x, y C.long;
   C.dw_pointer_query_pos(&x, &y);
   return int(x), int(y);
}

func Pointer_set_pos(x int, y int) {
   C.dw_pointer_set_pos(C.long(x), C.long(y));
}

func Flush() {
    C.dw_flush();
}

func init() {
   runtime.LockOSThread();
}

//export go_int_callback_basic
func go_int_callback_basic(pfunc unsafe.Pointer, window unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), data));
}

//export go_int_callback_configure
func go_int_callback_configure(pfunc unsafe.Pointer, window unsafe.Pointer, width C.int, height C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, int, int, unsafe.Pointer) C.int)(pfunc);
   return C.int(thisfunc(HWND(window), int(width), int(height), data));
}

//export go_int_callback_keypress
func go_int_callback_keypress(pfunc unsafe.Pointer, window unsafe.Pointer, ch C.char, vk C.int, state C.int, data unsafe.Pointer, utf8 *C.char) C.int {
   thisfunc := *(*func(HWND, uint8, int, int, unsafe.Pointer, string) int)(pfunc);
   return C.int(thisfunc(HWND(window), uint8(ch), int(vk), int(state), data, C.GoString(utf8)));
}

//export go_int_callback_mouse
func go_int_callback_mouse(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, mask C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, int, int, int, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), int(x), int(y), int(mask), data));
}

//export go_int_callback_expose
func go_int_callback_expose(pfunc unsafe.Pointer, window unsafe.Pointer, x C.int, y C.int, width C.int, height C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, int, int, int, int, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), int(x), int(y), int(width), int(height), data));
}

//export go_int_callback_string
func go_int_callback_string(pfunc unsafe.Pointer, window unsafe.Pointer, str *C.char, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, string, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), C.GoString(str), data));
}

//export go_int_callback_item_context
func go_int_callback_item_context(pfunc unsafe.Pointer, window unsafe.Pointer, text *C.char, x C.int, y C.int, data unsafe.Pointer, itemdata unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, string, int, int, unsafe.Pointer, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), C.GoString(text), int(x), int(y), data, itemdata));
}

//export go_int_callback_item_select
func go_int_callback_item_select(pfunc unsafe.Pointer, window unsafe.Pointer, item unsafe.Pointer, text *C.char, data unsafe.Pointer, itemdata unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, HTREEITEM, string, unsafe.Pointer, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), HTREEITEM(item), C.GoString(text), data, itemdata));
}

//export go_int_callback_numeric
func go_int_callback_numeric(pfunc unsafe.Pointer, window unsafe.Pointer, val C.int, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, int, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), int(val), data));
}

//export go_int_callback_ulong
func go_int_callback_ulong(pfunc unsafe.Pointer, window unsafe.Pointer, val C.ulong, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, uint, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), uint(val), data));
}

//export go_int_callback_tree
func go_int_callback_tree(pfunc unsafe.Pointer, window unsafe.Pointer, tree unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(HWND, HTREEITEM, unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(HWND(window), HTREEITEM(tree), data));
}

//export go_int_callback_timer
func go_int_callback_timer(pfunc unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(unsafe.Pointer) int)(pfunc);
   return C.int(thisfunc(data));
}

