package main

/*
#cgo linux pkg-config: dwindows
#cgo freebsd pkg-config: dwindows
#cgo darwin CFLAGS: -I/usr/local/include -g -O2 -D__MAC__
#cgo darwin LDFLAGS: -L/usr/local/lib -ldwindows -lresolv -framework Cocoa -framework WebKit -lpthread
#cgo windows CFLAGS: -IC:/Work/Netlabs/dwindows -g -O2 -D__WIN32__ -mthreads
#cgo windows LDFLAGS: -LC:/Work/Netlabs/dwindows -ldw
#include "../dw/dwglue.h"
*/
import "C"
import "unsafe"
import "dw"

func exit_callback(window dw.HWND, data unsafe.Pointer) C.int {
   if dw.Messagebox("dwtest", C.DW_MB_YESNO | C.DW_MB_QUESTION, "Are you sure you want to exit?") != 0 {
      dw.Main_quit();
   }
   return C.TRUE;
}

var exit_callback_func = exit_callback;

func main() {
   /* Initialize the Dynamic Windows engine */
   dw.Init(dw.TRUE);

   /* Create our window */
   mainwindow := dw.Window_new(dw.DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", C.DW_FCF_SYSMENU | C.DW_FCF_TITLEBAR | C.DW_FCF_TASKLIST | C.DW_FCF_DLGBORDER | C.DW_FCF_SIZEBORDER | C.DW_FCF_MINMAX);
   
   lbbox := dw.Box_new(C.DW_VERT, 10);

   dw.Box_pack_start(mainwindow, lbbox, 150, 70, dw.TRUE, dw.TRUE, 0);

   /* Copy and Paste */
   browsebox := dw.Box_new(C.DW_HORZ, 0);

   dw.Box_pack_start(lbbox, browsebox, 0, 0, dw.FALSE, dw.FALSE, 0);

   copypastefield := dw.Entryfield_new("", 0);

   dw.Entryfield_set_limit(copypastefield, 260);

   dw.Box_pack_start(browsebox, copypastefield, -1, -1, dw.TRUE, dw.FALSE, 4);

   copybutton := dw.Button_new("Copy", 0);

   dw.Box_pack_start(browsebox, copybutton, -1, -1, dw.FALSE, dw.FALSE, 0);

   pastebutton := dw.Button_new("Paste", 0);

   dw.Box_pack_start(browsebox, pastebutton, -1, -1, dw.FALSE, dw.FALSE, 0);

   /* Archive Name */
   stext := dw.Text_new("File to browse", 0);

   dw.Window_set_style(stext, C.DW_DT_VCENTER, C.DW_DT_VCENTER);

   dw.Box_pack_start(lbbox, stext, 130, 15, dw.TRUE, dw.TRUE, 2);

   browsebox = dw.Box_new(C.DW_HORZ, 0);

   dw.Box_pack_start(lbbox, browsebox, 0, 0, dw.TRUE, dw.TRUE, 0);

   entryfield := dw.Entryfield_new("", 100);

   dw.Entryfield_set_limit(entryfield, 260);

   dw.Box_pack_start(browsebox, entryfield, 100, 15, dw.TRUE, dw.TRUE, 4);

   browsefilebutton := dw.Button_new("Browse File", 1001);

   dw.Box_pack_start(browsebox, browsefilebutton, 40, 15, dw.TRUE, dw.TRUE, 0);

   browsefolderbutton := dw.Button_new("Browse Folder", 1001);

   dw.Box_pack_start(browsebox, browsefolderbutton, 40, 15, dw.TRUE, dw.TRUE, 0);

   dw.Window_set_color(browsebox, C.DW_CLR_PALEGRAY, C.DW_CLR_PALEGRAY);
   dw.Window_set_color(stext, C.DW_CLR_BLACK, C.DW_CLR_PALEGRAY);

   /* Buttons */
   buttonbox := dw.Box_new(C.DW_HORZ, 10);

   dw.Box_pack_start(lbbox, buttonbox, 0, 0, dw.TRUE, dw.TRUE, 0);

   cancelbutton := dw.Button_new("Exit", 1002);
   dw.Box_pack_start(buttonbox, cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2);

   cursortogglebutton := dw.Button_new("Set Cursor pointer - CLOCK", 1003);
   dw.Box_pack_start(buttonbox, cursortogglebutton, 130, 30, dw.TRUE, dw.TRUE, 2);

   okbutton := dw.Button_new("Turn Off Annoying Beep!", 1001);
   dw.Box_pack_start(buttonbox, okbutton, 130, 30, dw.TRUE, dw.TRUE, 2);

   dw.Box_unpack(cancelbutton);
   dw.Box_pack_start(buttonbox, cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2);
   dw.Window_click_default(mainwindow, cancelbutton);

   colorchoosebutton := dw.Button_new("Color Chooser Dialog", 1004);
   dw.Box_pack_at_index(buttonbox, colorchoosebutton, 1, 130, 30, dw.TRUE, dw.TRUE, 2);

   /* Set some nice fonts and colors */
   dw.Window_set_color(lbbox, C.DW_CLR_DARKCYAN, C.DW_CLR_PALEGRAY);
   dw.Window_set_color(buttonbox, C.DW_CLR_DARKCYAN, C.DW_CLR_PALEGRAY);
   dw.Window_set_color(okbutton, C.DW_CLR_PALEGRAY, C.DW_CLR_DARKCYAN);

   /*dw_signal_connect(browsefilebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(browse_file_callback), DW_POINTER(notebookbox1));
   dw_signal_connect(browsefolderbutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(browse_folder_callback), DW_POINTER(notebookbox1));
   dw_signal_connect(copybutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(copy_clicked_callback), DW_POINTER(copypastefield));
   dw_signal_connect(pastebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(paste_clicked_callback), DW_POINTER(copypastefield));
   dw_signal_connect(okbutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(beep_callback), DW_POINTER(notebookbox1));*/
   dw.Signal_connect(cancelbutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
   /*dw_signal_connect(cursortogglebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(cursortoggle_callback), DW_POINTER(mainwindow));
   dw_signal_connect(colorchoosebutton, DW_SIGNAL_CLICKED, DW_SIGNAL_FUNC(colorchoose_callback), DW_POINTER(mainwindow));*/
   
   /* Set the default field */
   dw.Window_default(mainwindow, copypastefield);

   dw.Signal_connect(mainwindow, C.DW_SIGNAL_DELETE, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
   /*
   * The following is a special case handler for the Mac and other platforms which contain
   * an application object which can be closed.  It function identically to a window delete/close
   * request except it applies to the entire application not an individual window. If it is not
   * handled or you allow the default handler to take place the entire application will close.
   * On platforms which do not have an application object this line will be ignored.
   */
   //dw.signal_connect(DW_DESKTOP, DW_SIGNAL_DELETE, DW_SIGNAL_FUNC(exit_callback), DW_POINTER(mainwindow));
   //timerid = dw.timer_connect(2000, DW_SIGNAL_FUNC(timer_callback), 0);
   dw.Window_set_size(mainwindow, 640, 550);
   dw.Window_show(mainwindow);
   
   dw.Main();
}

