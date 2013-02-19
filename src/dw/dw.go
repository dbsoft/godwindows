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
type COLORI C.uchar

type Env struct 
{
    OSName, BuildDate, BuildTime string
    MajorVersion, MinorVersion, MajorBuild, MinorBuild C.short
    DWMajorVersion, DWMinorVersion, DWSubVersion C.short
}

const (
   FALSE C.int = iota
   TRUE
)

var DESKTOP HWND = nil
var NOMENU HMENUI = nil

func RESOURCE(id uintptr) unsafe.Pointer {
   return unsafe.Pointer(id);
}

func RGB(red COLORI, green COLORI, blue COLORI) COLOR {
    lred := C.ulong(red);
    lgreen := C.ulong(green);
    lblue := C.ulong(blue);
    return COLOR((0xF0000000 | (lred) | (lgreen << 8) | (lblue << 16)));
}

func Init(newthread C.int) C.int {
   return C.go_init(newthread);
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

func Status_text_new(text string, id C.ulong) HWND {
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   return HWND(C.go_status_text_new(ctext, id));
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

func Timer_connect(interval C.int, sigfunc unsafe.Pointer, data unsafe.Pointer) HTIMER {
   return HTIMER(C.go_timer_connect(interval, sigfunc, data));
}

func Timer_disconnect(id HTIMER) {
   C.dw_timer_disconnect(C.int(id));
}

func Signal_connect(window HWND, signame string, sigfunc unsafe.Pointer, data unsafe.Pointer) {
   csigname := C.CString(signame);
   defer C.free(unsafe.Pointer(csigname));
   
   C.go_signal_connect(unsafe.Pointer(window), csigname, sigfunc, data);
}

func Beep(freq C.int, dur C.int) {
    C.dw_beep(freq, dur);
}

func Menu_new(id C.ulong) HMENUI {
    return HMENUI(C.go_menu_new(id));
}

func Menubar_new(location HWND) HMENUI {
    return HMENUI(C.go_menubar_new(unsafe.Pointer(location)));
}

func Menu_append_item(menu HMENUI, title string, id C.ulong, flags C.ulong, end C.int, check C.int, submenu HMENUI) HWND {
    ctitle := C.CString(title);
    defer C.free(unsafe.Pointer(ctitle));

    return HWND(C.go_menu_append_item(unsafe.Pointer(menu), ctitle, id, flags, end, check, unsafe.Pointer(submenu)));
}

func Menu_delete_item(menu HMENUI, id C.ulong) {
    C.go_menu_delete_item(unsafe.Pointer(menu), id);
}

func Menu_destroy(menu HMENUI) {
    C.go_menu_destroy(unsafe.Pointer(menu));
}

func Menu_item_set_state(menu HMENUI, id C.ulong, flags C.ulong) {
    C.go_menu_item_set_state(unsafe.Pointer(menu), id, flags);
}

func Menu_poup(menu HMENUI, parent HWND, x C.int, y C.int) {
    C.go_menu_popup(unsafe.Pointer(menu), unsafe.Pointer(parent), x, y);
}

func Notebook_new(id C.ulong, top C.int) HWND {
    return HWND(C.go_notebook_new(id, top));
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

func Notebook_page_new(handle HWND, flags C.ulong, front C.int) HNOTEPAGE {
    return HNOTEPAGE(C.go_notebook_page_new(unsafe.Pointer(handle), flags, front));
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

func Icon_load(id C.ulong) HICN {
    return HICN(C.go_icon_load(0, id));
}

func Taskbar_delete(handle HWND, icon HICN) {
    C.go_taskbar_delete(unsafe.Pointer(handle), unsafe.Pointer(icon));
}

func Taskbar_insert(handle HWND, icon HICN, bubbletext string) {
    cbubbletext := C.CString(bubbletext);
    defer C.free(unsafe.Pointer(cbubbletext));
    
    C.go_taskbar_insert(unsafe.Pointer(handle), unsafe.Pointer(icon), cbubbletext);
}

func Combobox_new(text string, id C.ulong) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_combobox_new(ctext, id));
}

func Listbox_new(id C.ulong, multi C.int) HWND {
    return HWND(C.go_listbox_new(id, multi));
}

func Listbox_append(handle HWND, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_append(unsafe.Pointer(handle), ctext);
}

func Listbox_insert(handle HWND, text string, pos C.int) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_insert(unsafe.Pointer(handle), ctext, pos);
}

func Listbox_clear(handle HWND) {
    C.go_listbox_clear(unsafe.Pointer(handle));
}

func Listbox_count(handle HWND) C.int {
    return C.go_listbox_count(unsafe.Pointer(handle));
}

func Listbox_set_top(handle HWND, top C.int) {
    C.go_listbox_set_top(unsafe.Pointer(handle), top);
}

func Listbox_select(handle HWND, index C.int, state C.int) {
    C.go_listbox_select(unsafe.Pointer(handle), index, state);
}

func Listbox_delete(handle HWND, index C.int) {
    C.go_listbox_delete(unsafe.Pointer(handle), index);
}

func Listbox_get_text(handle HWND, index C.int) string {
    var buf [201]C.char;
    
    C.go_listbox_get_text(unsafe.Pointer(handle), index, &buf[0], 200);
    return C.GoString((*C.char)(unsafe.Pointer(&buf[0])));
}

func Listbox_set_text(handle HWND, index C.int, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_listbox_set_text(unsafe.Pointer(handle), index, ctext);
}

func Listbox_selected(handle HWND) C.int {
    return C.go_listbox_selected(unsafe.Pointer(handle));
}

func Listbox_selected_multi(handle HWND, where C.int) C.int {
    return C.go_listbox_selected_multi(unsafe.Pointer(handle), where);
}

func Screen_width() C.int {
    return C.dw_screen_width();
}

func Screen_height() C.int {
    return C.dw_screen_height();
}

func Color_depth_get() C.ulong {
    return C.dw_color_depth_get();
}

func Spinbutton_new(text string, id C.ulong) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_spinbutton_new(ctext, id));
}

func Spinbutton_set_pos(handle HWND, position C.long) {
    C.go_spinbutton_set_pos(unsafe.Pointer(handle), position);
}

func Spinbutton_set_limits(handle HWND, upper C.long, lower C.long) {
    C.go_spinbutton_set_limits(unsafe.Pointer(handle), upper, lower);
}

func Spinbutton_get_pos(handle HWND) C.long {
    return C.go_spinbutton_get_pos(unsafe.Pointer(handle));
}

func Radiobutton_new(text string, id C.ulong) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_radiobutton_new(ctext, id));
}

func Checkbox_new(text string, id C.ulong) HWND {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    return HWND(C.go_checkbox_new(ctext, id));
}

func Checkbox_get(handle HWND) C.int {
    return C.go_checkbox_get(unsafe.Pointer(handle));
}

func Checkbox_set(handle HWND, value C.int) {
    C.go_checkbox_set(unsafe.Pointer(handle), value);
}

func Percent_new(id C.ulong) HWND {
    return HWND(C.go_percent_new(id));
}

func Slider_new(vertical C.int, increments C.int, id C.ulong) HWND {
    return HWND(C.go_slider_new(vertical, increments, id));
}

func Scrollbar_new(vertical C.int, id C.ulong) HWND {
    return HWND(C.go_scrollbar_new(vertical, id));
}

func Slider_get_pos(handle HWND) C.uint {
    return C.go_slider_get_pos(unsafe.Pointer(handle));
}

func Slider_set_pos(handle HWND, position C.uint) {
    C.go_slider_set_pos(unsafe.Pointer(handle), position);
}

func Scrollbar_get_pos(handle HWND) C.uint {
    return C.go_scrollbar_get_pos(unsafe.Pointer(handle));
}

func Scrollbar_set_pos(handle HWND, position C.uint) {
    C.go_scrollbar_set_pos(unsafe.Pointer(handle), position);
}

func Scrollbar_set_range(handle HWND, srange C.uint, visible C.uint) {
    C.go_scrollbar_set_range(unsafe.Pointer(handle), srange, visible);
}

func Scrollbox_new(btype C.int, pad C.int) HWND {
    return HWND(C.go_scrollbox_new(btype, pad));
}

func Scrollbox_get_pos(handle HWND, orient C.int) C.int {
    return C.go_scrollbox_get_pos(unsafe.Pointer(handle), orient);
}

func Scrollbox_get_range(handle HWND, orient C.int) C.int {
    return C.go_scrollbox_get_range(unsafe.Pointer(handle), orient);
}

func Groupbox_new(btype C.int, pad C.int, title string) HWND {
    ctitle := C.CString(title);
    defer C.free(unsafe.Pointer(ctitle));
    
    return HWND(C.go_groupbox_new(btype, pad, ctitle));
}

func Render_new(id C.ulong) HWND {
    return HWND(C.go_render_new(id));
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

func Font_text_extents_get(handle HWND, pixmap HPIXMAP, text string) (C.int, C.int) {
   var width, height C.int;
   
   ctext := C.CString(text);
   defer C.free(unsafe.Pointer(ctext));
   
   C.go_font_text_extents_get(unsafe.Pointer(handle), unsafe.Pointer(pixmap), ctext, &width, &height);
   return width, height;
}

func Pixmap_new(handle HWND, width C.ulong, height C.ulong, depth C.ulong) HPIXMAP {
    return HPIXMAP(C.go_pixmap_new(unsafe.Pointer(handle), width, height, depth));
}

func Pixmap_new_from_file(handle HWND, filename string) HPIXMAP {
    cfilename := C.CString(filename);
    defer C.free(unsafe.Pointer(cfilename));
    
    return HPIXMAP(C.go_pixmap_new_from_file(unsafe.Pointer(handle), cfilename));
}

func Pixmap_grab(handle HWND, id C.ulong) HPIXMAP {
    return HPIXMAP(C.go_pixmap_grab(unsafe.Pointer(handle), id));
}

func Pixmap_bitblt(dest HWND, destp HPIXMAP, xdest C.int, ydest C.int, width C.int, height C.int, src HWND, srcp HPIXMAP, xsrc C.int, ysrc C.int) {
    C.go_pixmap_bitblt(unsafe.Pointer(dest), unsafe.Pointer(srcp), xdest, ydest, width, height, unsafe.Pointer(src), unsafe.Pointer(srcp), xsrc, ysrc); 
}

func Pixmap_stretch_bitblt(dest HWND, destp HPIXMAP, xdest C.int, ydest C.int, width C.int, height C.int, src HWND, srcp HPIXMAP, xsrc C.int, ysrc C.int, srcwidth C.int, srcheight C.int) C.int {
    return C.go_pixmap_stretch_bitblt(unsafe.Pointer(dest), unsafe.Pointer(srcp), xdest, ydest, width, height, unsafe.Pointer(src), unsafe.Pointer(srcp), xsrc, ysrc, srcwidth, srcheight); 
}

func Pixmap_set_transparent_color(pixmap HPIXMAP, color COLOR) {
    C.go_pixmap_set_transparent_color(unsafe.Pointer(pixmap), C.ulong(color));
}

func Pixmap_set_font(handle HWND, fontname string) C.int {
    cfontname := C.CString(fontname);
    defer C.free(unsafe.Pointer(cfontname));
    
    return C.go_pixmap_set_font(unsafe.Pointer(handle), cfontname);
}

func Pixmap_destroy(pixmap HPIXMAP) {
    C.go_pixmap_destroy(unsafe.Pointer(pixmap));
}

func Draw_point(handle HWND, pixmap HPIXMAP, x C.int, y C.int) {
    C.go_draw_point(unsafe.Pointer(handle), unsafe.Pointer(pixmap), x, y);
}

func Draw_line(handle HWND, pixmap HPIXMAP, x1 C.int, y1 C.int, x2 C.int, y2 C.int) {
    C.go_draw_line(unsafe.Pointer(handle), unsafe.Pointer(pixmap), x1, y1, x2, y2);
}

func Draw_rect(handle HWND, pixmap HPIXMAP, fill C.int, x C.int, y C.int, width C.int, height C.int) {
    C.go_draw_rect(unsafe.Pointer(handle), unsafe.Pointer(pixmap), fill, x, y, width, height);
}

func Draw_arc(handle HWND, pixmap HPIXMAP, flags C.int, xorigin C.int, yorigin C.int, x1 C.int, y1 C.int, x2 C.int, y2 C.int) {
    C.go_draw_arc(unsafe.Pointer(handle), unsafe.Pointer(pixmap), flags, xorigin, yorigin, x1, y1, x2, y2);
}

func Draw_text(handle HWND, pixmap HPIXMAP, x C.int, y C.int, text string) {
    ctext := C.CString(text);
    defer C.free(unsafe.Pointer(ctext));
    
    C.go_draw_text(unsafe.Pointer(handle), unsafe.Pointer(pixmap), x, y, ctext);
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

//export go_int_callback_timer
func go_int_callback_timer(pfunc unsafe.Pointer, data unsafe.Pointer) C.int {
   thisfunc := *(*func(unsafe.Pointer) C.int)(pfunc);
   return thisfunc(data);
}

