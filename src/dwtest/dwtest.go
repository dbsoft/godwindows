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
import "fmt"

const (
   FALSE C.int = iota
   TRUE
)

func exit_callback(window dw.HWND, data unsafe.Pointer) C.int {
   if dw.Messagebox("dwtest", C.DW_MB_YESNO | C.DW_MB_QUESTION, "Are you sure you want to exit?") != 0 {
      dw.Main_quit();
   }
   return C.TRUE;
}

var copypastefield, entryfield, cursortogglebutton dw.HWND
var current_file string
var current_color dw.COLOR
var cursor_arrow bool

func copy_clicked_callback(button dw.HWND, data unsafe.Pointer) C.int {
   test := dw.Window_get_text(copypastefield);

   if len(test) > 0 {
     dw.Clipboard_set_text(test);
   }
   dw.Window_set_focus(entryfield);
   return TRUE;
}

func paste_clicked_callback(button dw.HWND, data unsafe.Pointer) C.int {
    test := dw.Clipboard_get_text();
    
    if len(test) > 0 {
        dw.Window_set_text(copypastefield, test);
    }
    return TRUE;
}

func browse_file_callback(window dw.HWND, data unsafe.Pointer) C.int {
    tmp := dw.File_browse("Pick a file", "dwtest.c", "c", C.DW_FILE_OPEN);
    if len(tmp) > 0 {
        current_file = tmp;
        dw.Window_set_text(entryfield, current_file);
        /*read_file();
        current_col = current_row = 0;
        update_render();*/
    }
    dw.Window_set_focus(copypastefield);
    return FALSE;
}

func browse_folder_callback(window dw.HWND, data unsafe.Pointer) C.int {
    tmp := dw.File_browse("Pick a folder", ".", "c", C.DW_DIRECTORY_OPEN);
    fmt.Printf("Folder picked: %s\n", tmp);
    return FALSE;
}

func colorchoose_callback(window dw.HWND, data unsafe.Pointer) C.int {
    current_color = dw.Color_choose(current_color);
    return FALSE;
}

func cursortoggle_callback(window dw.HWND, data unsafe.Pointer) C.int {
    if cursor_arrow {
        dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - ARROW");
        dw.Window_set_pointer(dw.HWND(data), C.DW_POINTER_CLOCK);
        cursor_arrow = false;
    } else {
        dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - CLOCK");
        dw.Window_set_pointer(dw.HWND(data), C.DW_POINTER_DEFAULT);
        cursor_arrow = true;
    }
    return FALSE;
}

func beep_callback(window dw.HWND, data unsafe.Pointer) C.int {
    //dw.Timer_disconnect(timerid);
    return TRUE;
}

var exit_callback_func = exit_callback;
var copy_clicked_callback_func = copy_clicked_callback;
var paste_clicked_callback_func = paste_clicked_callback;
var browse_file_callback_func = browse_file_callback;
var browse_folder_callback_func = browse_folder_callback;
var colorchoose_callback_func = colorchoose_callback;
var cursortoggle_callback_func = cursortoggle_callback;
var beep_callback_func = beep_callback;

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

   copypastefield = dw.Entryfield_new("", 0);

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

   entryfield = dw.Entryfield_new("", 100);

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

   cursortogglebutton = dw.Button_new("Set Cursor pointer - CLOCK", 1003);
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

   dw.Signal_connect(browsefilebutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&browse_file_callback_func), nil);
   dw.Signal_connect(browsefolderbutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&browse_folder_callback_func), nil);
   dw.Signal_connect(copybutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&copy_clicked_callback_func), unsafe.Pointer(copypastefield));
   dw.Signal_connect(pastebutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&paste_clicked_callback_func), unsafe.Pointer(copypastefield));
   dw.Signal_connect(okbutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&beep_callback_func), nil);
   dw.Signal_connect(cancelbutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
   dw.Signal_connect(cursortogglebutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&cursortoggle_callback_func), unsafe.Pointer(mainwindow));
   dw.Signal_connect(colorchoosebutton, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&colorchoose_callback_func), unsafe.Pointer(mainwindow));
   
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
   dw.Signal_connect(dw.DESKTOP, C.DW_SIGNAL_DELETE, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
   //timerid = dw.timer_connect(2000, DW_SIGNAL_FUNC(timer_callback), 0);
   dw.Window_set_size(mainwindow, 640, 550);
   dw.Window_show(mainwindow);
   
   dw.Main();
}

