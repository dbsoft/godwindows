package main

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/Netlabs/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/Netlabs/dwindows -ldw
#include "dwglue.h"
*/
import "C"
import "unsafe"
import "runtime"
import "dwindows"

var dw DW;

func exit_handler(window HWND, data unsafe.Pointer) C.int {
   dw.main_quit();
   return FALSE;
}

func main() {
   dw = new(DW);
   
   /* Initialize the Dynamic Windows engine */
   dw.init(TRUE);

   /* Create our window */
   mainwindow := dw.window_new( HWND_DESKTOP, "dwindows test UTF8 ??? (??) cañón", C.DW_FCF_SYSMENU | C.DW_FCF_TITLEBAR | C.DW_FCF_TASKLIST | C.DW_FCF_DLGBORDER | C.DW_FCF_SIZEBORDER | C.DW_FCF_MINMAX);
   
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

   dw.signal_connect(mainwindow, C.DW_SIGNAL_DELETE, unsafe.Pointer(&exit_callback), unsafe.Pointer(mainwindow));
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

