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

var copypastefield, entryfield, cursortogglebutton, mainwindow, notebookbox1, noncheckable_menuitem, checkable_menuitem dw.HWND
var textbox1, textbox2 dw.HWND
var fileicon, foldericon dw.HICN
var current_file string
var current_color dw.COLOR = dw.RGB(100, 100, 100)
var cursor_arrow bool = true
var menu_enabled bool = true
var timerid dw.HTIMER

var FOLDER_ICON_NAME string = "mac/folder"
var FILE_ICON_NAME string = "mac/file"

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
    dw.Timer_disconnect(timerid);
    return TRUE;
}

/* Beep every second */
func timer_callback(data unsafe.Pointer) C.int {
    dw.Beep(200, 200);

    /* Return TRUE so we get called again */
    return TRUE;
}

func switch_page_callback(window dw.HWND, page_num dw.HNOTEPAGE, itemdata unsafe.Pointer) C.int {
    fmt.Printf("DW_SIGNAL_SWITCH_PAGE: PageNum: %u\n", uint(page_num));
    return FALSE;
}

func menu_callback(window dw.HWND, data unsafe.Pointer) C.int {
    info:= *(*string)(data);
    buf := fmt.Sprintf("%s menu item selected", info);
    dw.Messagebox("Menu Item Callback", C.DW_MB_OK | C.DW_MB_INFORMATION, buf);
    return 0;
}

func menutoggle_callback(window dw.HWND, data unsafe.Pointer) C.int {
    if menu_enabled {
        dw.Window_set_style(checkable_menuitem, C.DW_MIS_DISABLED, C.DW_MIS_DISABLED);
        dw.Window_set_style(noncheckable_menuitem, C.DW_MIS_DISABLED, C.DW_MIS_DISABLED);
        menu_enabled = false;
    } else {
        dw.Window_set_style(checkable_menuitem, C.DW_MIS_DISABLED, C.DW_MIS_ENABLED);
        dw.Window_set_style(noncheckable_menuitem, C.DW_MIS_DISABLED, C.DW_MIS_ENABLED);
        menu_enabled = true;
    }
    return FALSE;
}

func helpabout_callback(window dw.HWND, data unsafe.Pointer) C.int {
    var env dw.Env;
    
    dw.Environment_query(&env);
    message := fmt.Sprintf("dwindows test\n\nOS: %s %s %s Version: %d.%d.%d.%d\n\ndwindows Version: %d.%d.%d",
                              env.OSName, env.BuildDate, env.BuildTime,
                              env.MajorVersion, env.MinorVersion, env.MajorBuild, env.MinorBuild,
                              env.DWMajorVersion, env.DWMinorVersion, env.DWSubVersion);
    dw.Messagebox("About dwindows", C.DW_MB_OK | C.DW_MB_INFORMATION, message);
    return FALSE;
}

var exit_callback_func = exit_callback;
var copy_clicked_callback_func = copy_clicked_callback;
var paste_clicked_callback_func = paste_clicked_callback;
var browse_file_callback_func = browse_file_callback;
var browse_folder_callback_func = browse_folder_callback;
var colorchoose_callback_func = colorchoose_callback;
var cursortoggle_callback_func = cursortoggle_callback;
var beep_callback_func = beep_callback;
var timer_callback_func = timer_callback;
var switch_page_callback_func = switch_page_callback;
var helpabout_callback_func = helpabout_callback;
var menu_callback_func = menu_callback;
var menutoggle_callback_func = menutoggle_callback;

var checkable_string = "checkable";
var noncheckable_string = "non-checkable";

func menu_add() {
    mainmenubar := dw.Menubar_new(mainwindow);
    /* add menus to the menubar */
    menu := dw.Menu_new(0);
    menuitem := dw.Menu_append_item(menu, "~Quit", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(menuitem, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
    /*
     * Add the "File" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~File", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);

    changeable_menu := dw.Menu_new(0);
    checkable_menuitem = dw.Menu_append_item(changeable_menu, "~Checkable Menu Item", C.DW_MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    dw.Signal_connect(checkable_menuitem, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&menu_callback_func), unsafe.Pointer(&checkable_string));
    noncheckable_menuitem = dw.Menu_append_item(changeable_menu, "~Non-checkable Menu Item", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(noncheckable_menuitem, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&menu_callback_func), unsafe.Pointer(&noncheckable_string));
    dw.Menu_append_item(changeable_menu, "~Disabled menu Item", C.DW_MENU_AUTO, C.DW_MIS_DISABLED | C.DW_MIS_CHECKED, dw.TRUE, dw.TRUE, dw.NOMENU);
    /* seperator */
    dw.Menu_append_item(changeable_menu, C.DW_MENU_SEPARATOR, C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem = dw.Menu_append_item(changeable_menu, "~Menu Items Disabled", C.DW_MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    dw.Signal_connect(menuitem, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&menutoggle_callback_func), nil);
    /*
     * Add the "Menu" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~Menu", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, changeable_menu);

    menu = dw.Menu_new(0);
    menuitem = dw.Menu_append_item(menu, "~About", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(menuitem, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&helpabout_callback_func), unsafe.Pointer(mainwindow));
    /*
     * Add the "Help" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~Help", C.DW_MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);
}

func archive_add() {
    lbbox := dw.Box_new(C.DW_VERT, 10);

    dw.Box_pack_start(notebookbox1, lbbox, 150, 70, dw.TRUE, dw.TRUE, 0);

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
}

func text_add() {
    depth := dw_color_depth_get();

    /* create a box to pack into the notebook page */
    pagebox := dw.Box_new(C.DW_HORZ, 2);
    dw.Box_pack_start(notebookbox2, pagebox, 0, 0, dw.TRUE, dw.TRUE, 0);
    /* now a status area under this box */
    hbox := dw.Box_new(C.DW_HORZ, 1);
    dw.Box_pack_start(notebookbox2, hbox, 100, 20, dw.TRUE, dw.FALSE, 1);
    status1 = dw.Status_text_new("", 0);
    dw.Box_pack_start(hbox, status1, 100, -1, dw.TRUE, dw.FALSE, 1);
    status2 = dw.Status_text_new("", 0);
    dw.Box_pack_start(hbox, status2, 100, -1, dw.TRUE, dw.FALSE, 1);
    /* a box with combobox and button */
    hbox = dw.Box_new(C.DW_HORZ, 1);
    dw.Box_pack_start(notebookbox2, hbox, 100, 25, dw.TRUE, dw.FALSE, 1);
    rendcombo := dw.Combobox_new("Shapes Double Buffered", 0);
    dw.Box_pack_start(hbox, rendcombo, 80, 25, dw.TRUE, dw.FALSE, 0);
    dw.Listbox_append(rendcombo, "Shapes Double Buffered");
    dw.Listbox_append(rendcombo, "Shapes Direct");
    dw.Listbox_append(rendcombo, "File Display");
    label := dw.Text_new("Image X:", 100);
    dw.Window_set_style(label, C.DW_DT_VCENTER | C.DW_DT_CENTER, C.DW_DT_VCENTER | C.DW_DT_CENTER);
    dw.Box_pack_start( hbox, label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imagexspin := dw.Spinbutton_new("20", 1021);
    dw.Box_pack_start(hbox, imagexspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    label = dw.Text_new("Y:", 100);
    dw.Window_set_style(label, C.DW_DT_VCENTER | C.DW_DT_CENTER, C.DW_DT_VCENTER | C.DW_DT_CENTER);
    dw.Box_pack_start(hbox, label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imageyspin := dw.Spinbutton_new("20", 1021);
    dw.Box_pack_start(hbox, imageyspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    dw.Spinbutton_set_limits(imagexspin, 2000, 0);
    dw.Spinbutton_set_limits(imageyspin, 2000, 0);
    dw.Spinbutton_set_pos(imagexspin, 20);
    dw.Spinbutton_set_pos(imageyspin, 20);
    imagestretchcheck = dw.Checkbox_new("Stretch", 1021);
    dw.Box_pack_start(hbox, imagestretchcheck, -1, 25, dw.FALSE, dw.FALSE, 0);

    button1 = dw.Button_new("Refresh", 1223);
    dw.Box_pack_start(hbox, button1, 100, 25, dw.FALSE, dw.FALSE, 0);
    button2 = dw.Button_new("Print", 1224);
    dw.Box_pack_start(hbox, button2, 100, 25, dw.FALSE, dw.FALSE, 0);

    /* Pre-create the scrollbars so we can query their sizes */
    vscrollbar = dw.Scrollbar_new(C.DW_VERT, 50);
    hscrollbar = dw.Scrollbar_new(C.DW_HORZ, 50);
    dw.Window_get_preferred_size(vscrollbar, &vscrollbarwidth, NULL);
    dw.Window_get_preferred_size(hscrollbar, NULL, &hscrollbarheight);

    /* On GTK with overlay scrollbars enabled this returns us 0...
     * so in that case we need to give it some real values.
     */
    if vscrollbarwidth == 0 {
        vscrollbarwidth = 8;
    }
    if hscrollbarheight == 0 {
        hscrollbarheight = 8;
    }

    /* create render box for number pixmap */
    textbox1 = dw.Render_new(100);
    //dw.Window_set_font(textbox1, FIXEDFONT);
    dw.Font_text_extents_get(textbox1, NULL, "(g", &font_width, &font_height);
    font_width = font_width / 2;
    vscrollbox = dw.Box_new(C.DW_VERT, 0);
    dw.Box_pack_start(vscrollbox, textbox1, font_width*width1, font_height*rows, dw.FALSE, dw.TRUE, 0);
    dw.Box_pack_start(vscrollbox, 0, (font_width*(width1+1)), hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    /* pack empty space 1 character wide */
    dw.Box_pack_start(pagebox, 0, font_width, 0, dw.FALSE, dw.TRUE, 0);

    /* create box for filecontents and horz scrollbar */
    textboxA := dw.Box_new(C.DW_VERT, 0);
    dw.Box_pack_start(pagebox, textboxA, 0, 0, dw.TRUE, dw.TRUE, 0);

    /* create render box for filecontents pixmap */
    textbox2 = dw.Render_new(101);
    dw.Box_pack_start(textboxA, textbox2, 10, 10, dw.TRUE, dw.TRUE, 0);
    dw.Window_set_font(textbox2, FIXEDFONT);
    /* create horizonal scrollbar */
    dw.Box_pack_start(textboxA, hscrollbar, -1, -1, dw.TRUE, dw.FALSE, 0);

    /* create vertical scrollbar */
    vscrollbox = dw.Box_new(C.DW_VERT, 0);
    dw.Box_pack_start(vscrollbox, vscrollbar, -1, -1, dw.FALSE, dw.TRUE, 0);
    /* Pack an area of empty space 14x14 pixels */
    dw.Box_pack_start(vscrollbox, 0, vscrollbarwidth, hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    text1pm = dw.Pixmap_new(textbox1, font_width*width1, font_height*rows, depth);
    text2pm = dw.Pixmap_new(textbox2, font_width*cols, font_height*rows, depth);
    image = dw.Pixmap_new_from_file(textbox2, "image/test");
    if image == nil {
        image = dw.Pixmap_new_from_file(textbox2, "~/test");
    }
    if image != nil {
        dw.Pixmap_set_transparent_color(image, C.DW_CLR_WHITE);
    }

    dw.Messagebox("DWTest", C.DW_MB_OK | C.DW_MB_INFORMATION, fmt.Sprintf("Width: %d Height: %d\n", font_width, font_height));
    dw.Draw_rect(nil, text1pm, C.DW_DRAW_FILL | C.DW_DRAW_NOAA, 0, 0, font_width*width1, font_height*rows);
    dw.Draw_rect(nil, text2pm, C.DW_DRAW_FILL | C.DW_DRAW_NOAA, 0, 0, font_width*cols, font_height*rows);
    /*dw.Signal_connect(textbox1, C.DW_SIGNAL_BUTTON_PRESS, unsafe.Pointer(&context_menu_event_func), nil);
    dw.Signal_connect(textbox1, C.DW_SIGNAL_EXPOSE, unsafe.Pointer(&text_expose_func), nil);
    dw.Signal_connect(textbox2, C.DW_SIGNAL_EXPOSE, unsafe.Pointer(&text_expose_func), nil);
    dw.Signal_connect(textbox2, C.DW_SIGNAL_CONFIGURE, unsafe.Pointer(&configure_event_func), text2pm);
    dw.Signal_connect(textbox2, C.DW_SIGNAL_MOTION_NOTIFY, unsafe.Pointer(&motion_notify_event_func), DW_INT_TO_POINTER(1));
    dw.Signal_connect(textbox2, C.DW_SIGNAL_BUTTON_PRESS, unsafe.Pointer(&motion_notify_event_func), DW_INT_TO_POINTER(0));
    dw.Signal_connect(hscrollbar, C.DW_SIGNAL_VALUE_CHANGED, unsafe.Pointer(&scrollbar_valuechanged_callback_func), DW_POINTER(status1));
    dw.Signal_connect(vscrollbar, C.DW_SIGNAL_VALUE_CHANGED, unsafe.Pointer(&scrollbar_valuechanged_callback_func), DW_POINTER(status1));
    dw.Signal_connect(imagestretchcheck, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&refresh_callback_func), nil);
    dw.Signal_connect(button1, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&refresh_callback_func), nil);
    dw.Signal_connect(button2, C.DW_SIGNAL_CLICKED, unsafe.Pointer(&print_callback_func), nil);
    dw.Signal_connect(rendcombo, C.DW_SIGNAL_LIST_SELECT, unsafe.Pointer(&render_select_event_callback_func), nil);
    dw.Signal_connect(mainwindow, C.DW_SIGNAL_KEY_PRESS, unsafe.Pointer(&keypress_callback_func), nil);*/

    dw.Taskbar_insert(textbox1, fileicon, "DWTest");
}

func main() {
   /* Initialize the Dynamic Windows engine */
   dw.Init(dw.TRUE);

   /* Create our window */
   mainwindow = dw.Window_new(dw.DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", C.DW_FCF_SYSMENU | C.DW_FCF_TITLEBAR | C.DW_FCF_TASKLIST | C.DW_FCF_DLGBORDER | C.DW_FCF_SIZEBORDER | C.DW_FCF_MINMAX);
   
   menu_add();

   notebookbox := dw.Box_new(C.DW_VERT, 5);
   dw.Box_pack_start(mainwindow, notebookbox, 0, 0, dw.TRUE, dw.TRUE, 0);

   foldericon = dw.Icon_load_from_file(FOLDER_ICON_NAME);
   fileicon = dw.Icon_load_from_file(FILE_ICON_NAME);

   notebook := dw.Notebook_new(1, dw.TRUE);
   dw.Box_pack_start(notebookbox, notebook, 100, 100, dw.TRUE, dw.TRUE, 0);
   dw.Signal_connect(notebook, C.DW_SIGNAL_SWITCH_PAGE, unsafe.Pointer(&switch_page_callback_func), nil);

   notebookbox1 = dw.Box_new(C.DW_VERT, 5);
   notebookpage1 := dw.Notebook_page_new(notebook, 0, dw.TRUE);
   dw.Notebook_pack(notebook, notebookpage1, notebookbox1);
   dw.Notebook_page_set_text(notebook, notebookpage1, "buttons and entry");
   archive_add();
   
   notebookbox2 = dw.Box_new(C.DW_VERT, 5);
   notebookpage2 := dw.Notebook_page_new(notebook, 1, dw.FALSE);
   dw.Notebook_pack(notebook, notebookpage2, notebookbox2);
   dw.Notebook_page_set_text(notebook, notebookpage2, "render");
   text_add();

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
   timerid = dw.Timer_connect(2000, unsafe.Pointer(&timer_callback_func), nil);
   dw.Window_set_size(mainwindow, 640, 550);
   dw.Window_show(mainwindow);
   
  /* Now that the window is created and shown...
   * run the main loop until we get dw_main_quit()
   */
   dw.Main();
   
   /* Now that the loop is done we can cleanup */
   dw.Taskbar_delete(textbox1, fileicon);
   dw.Window_destroy(mainwindow);

   fmt.Printf("dwtest exiting...\n");
}

