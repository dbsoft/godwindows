package main

import (
   "unsafe"
   "dw"
   "fmt"
   "os"
   "bufio"
   "bytes"
   "io"
   "runtime"
)

// Global variables
const (
   FALSE int = iota
   TRUE
)

var FIXEDFONT = "10.monospace"

// Page 1
var notebookbox1, copypastefield, entryfield, cursortogglebutton, mainwindow, noncheckable_menuitem, checkable_menuitem dw.HWND
var current_color dw.COLOR = dw.RGB(100, 100, 100)
var cursor_arrow bool = true
var timerid dw.HTIMER

// Page 2
var notebookbox2, textbox1, textbox2, status1, status2, vscrollbar, hscrollbar, rendcombo, imagexspin, imageyspin, imagestretchcheck dw.HWND
var text1pm, text2pm, image dw.HPIXMAP
var image_x = 20
var image_y = 20
var image_stretch int = FALSE
var font_width = 8
var font_height = 12
var rows = 10
var width1 = 6
var cols = 80
var render_type = 0
var current_row = 0
var current_col = 0
var max_linewidth = 0

// Page 3
var notebookbox3, tree dw.HWND

// Page 4
var notebookbox4, container_mle, container dw.HWND
var mle_point = 0

// Page 5
var notebookbox5, combobox1, combobox2, cal, spinbutton, slider, percent, buttonsbox, buttonboxperm dw.HWND

// Page 7
var notebookbox7, html dw.HWND

// Page 8
var notebookbox8, scrollbox dw.HWND
var MAX_WIDGETS = 20

var iteration = 0;


// Miscellaneous
var fileicon, foldericon dw.HICN
var current_file string
var lines []string
var menu_enabled bool = true

var FOLDER_ICON_NAME string = "mac/folder"
var FILE_ICON_NAME string = "mac/file"

func read_file() {
    var (
        file *os.File
        part []byte
        prefix bool
        length int
        err error
    )
    
    lines = nil;
    max_linewidth = 0;
    
    if file, err = os.Open(current_file); err != nil {
        return;
    }
    reader := bufio.NewReader(file);
    buffer := bytes.NewBuffer(make([]byte, 1024));
    buffer.Reset();
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break;
        }
        buffer.Write(part);
        if !prefix {
            lines = append(lines, buffer.String());
            length = len(buffer.String());
            if length > max_linewidth {
               max_linewidth = length;
            }
            buffer.Reset();
        }
    }
    if err == io.EOF {
        err = nil;
    }
    dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols));
    dw.Scrollbar_set_pos(hscrollbar, 0);
    dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows));
    dw.Scrollbar_set_pos(vscrollbar, 0);
}

// Call back section
func exit_callback(window dw.HWND, data unsafe.Pointer) int {
   if dw.Messagebox("dwtest", dw.MB_YESNO | dw.MB_QUESTION, "Are you sure you want to exit?") != 0 {
      dw.Main_quit();
   }
   return TRUE;
}

func switch_page_callback(window dw.HWND, page_num dw.HNOTEPAGE, itemdata unsafe.Pointer) int {
    fmt.Printf("DW_SIGNAL_SWITCH_PAGE: PageNum: %d\n", uint(page_num));
    return FALSE;
}

func menu_callback(window dw.HWND, data unsafe.Pointer) int {
    info:= *(*string)(data);
    buf := fmt.Sprintf("%s menu item selected", info);
    dw.Messagebox("Menu Item Callback", dw.MB_OK | dw.MB_INFORMATION, buf);
    return FALSE;
}

func menutoggle_callback(window dw.HWND, data unsafe.Pointer) int {
    if menu_enabled {
        dw.Window_set_style(checkable_menuitem, dw.MIS_DISABLED, dw.MIS_DISABLED);
        dw.Window_set_style(noncheckable_menuitem, dw.MIS_DISABLED, dw.MIS_DISABLED);
        menu_enabled = false;
    } else {
        dw.Window_set_style(checkable_menuitem, dw.MIS_DISABLED, dw.MIS_ENABLED);
        dw.Window_set_style(noncheckable_menuitem, dw.MIS_DISABLED, dw.MIS_ENABLED);
        menu_enabled = true;
    }
    return FALSE;
}

func helpabout_callback(window dw.HWND, data unsafe.Pointer) int {
    var env dw.Env;

    dw.Environment_query(&env);
    message := fmt.Sprintf("dwindows test\n\nOS: %s %s %s Version: %d.%d.%d.%d\n\ndwindows Version: %d.%d.%d",
                              env.OSName, env.BuildDate, env.BuildTime,
                              env.MajorVersion, env.MinorVersion, env.MajorBuild, env.MinorBuild,
                              env.DWMajorVersion, env.DWMinorVersion, env.DWSubVersion);
    dw.Messagebox("About dwindows", dw.MB_OK | dw.MB_INFORMATION, message);
    return FALSE;
}

// Page 1 Callbacks
func paste_clicked_callback(button dw.HWND, data unsafe.Pointer) int {
    test := dw.Clipboard_get_text();

    if len(test) > 0 {
        dw.Window_set_text(copypastefield, test);
    }
    return TRUE;
}

func copy_clicked_callback(button dw.HWND, data unsafe.Pointer) int {
   test := dw.Window_get_text(copypastefield);

   if len(test) > 0 {
     dw.Clipboard_set_text(test);
   }
   dw.Window_set_focus(entryfield);
   return TRUE;
}

func browse_file_callback(window dw.HWND, data unsafe.Pointer) int {
    tmp := dw.File_browse("Pick a file", "dwtest.c", "c", dw.FILE_OPEN);
    if len(tmp) > 0 {
        current_file = tmp;
        dw.Window_set_text(entryfield, current_file);
        read_file();
        current_col = 0;
        current_row = 0;
        update_render();
    }
    dw.Window_set_focus(copypastefield);
    return FALSE;
}

func browse_folder_callback(window dw.HWND, data unsafe.Pointer) int {
    tmp := dw.File_browse("Pick a folder", ".", "c", dw.DIRECTORY_OPEN);
    fmt.Printf("Folder picked: %s\n", tmp);
    return FALSE;
}

func colorchoose_callback(window dw.HWND, data unsafe.Pointer) int {
    current_color = dw.Color_choose(current_color);
    return FALSE;
}

func cursortoggle_callback(window dw.HWND, data unsafe.Pointer) int {
    if cursor_arrow {
        dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - ARROW");
        dw.Window_set_pointer(dw.HWND(data), dw.POINTER_CLOCK);
        cursor_arrow = false;
    } else {
        dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - CLOCK");
        dw.Window_set_pointer(dw.HWND(data), dw.POINTER_DEFAULT);
        cursor_arrow = true;
    }
    return FALSE;
}

func beep_callback(window dw.HWND, data unsafe.Pointer) int {
    dw.Timer_disconnect(timerid);
    return TRUE;
}

/* Beep every second */
func timer_callback(data unsafe.Pointer) int {
    dw.Beep(200, 200);

    /* Return TRUE so we get called again */
    return TRUE;
}

// Page 2 Callbacks
func motion_notify_event(window dw.HWND, x int, y int, buttonmask int, data unsafe.Pointer) int {
    var which = "button_press";

    if(uintptr(data) > 0) {
        which = "motion_notify";
    }
    dw.Window_set_text(status2, fmt.Sprintf("%s: %dx%d", which, x, y));
    return FALSE;
}

func show_window_callback(window dw.HWND, data unsafe.Pointer) int {
    thiswindow := dw.HWND(data);

    if thiswindow != nil {
        dw.Window_show(thiswindow);
        dw.Window_raise(thiswindow);
    }
    return TRUE;
}

func context_menu_event(window dw.HWND, x int, y int, buttonmask int, data unsafe.Pointer) int {
    hwndMenu := dw.Menu_new(0);
    menuitem := dw.Menu_append_item(hwndMenu, "~Quit", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU);

    dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
    dw.Menu_append_item(hwndMenu, dw.MENU_SEPARATOR, dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU);
    menuitem = dw.Menu_append_item(hwndMenu, "~Show Window", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU);
    dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&show_window_callback_func), unsafe.Pointer(mainwindow));
    px, py := dw.Pointer_query_pos();
    /* Use the toplevel window handle here.... because on the Mac..
     * using the control itself, when a different tab is active
     * the control is removed from the window and can no longer
     * handle the messages.
     */
    dw.Menu_popup(hwndMenu, mainwindow, px, py);
    return TRUE;
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_file(row int, col int, nrows int, fheight int, hpma dw.HPIXMAP) {
    var hpm dw.HPIXMAP

    if hpma == nil {
        hpm = text2pm;
    } else {
        hpm = hpma;
    }

    if len(current_file) > 0 {
        var i int

        dw.Color_foreground_set(dw.CLR_WHITE);
        if hpma == nil {
            dw.Draw_rect(nil, text1pm, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(text1pm), dw.Pixmap_height(text1pm));
        }
        dw.Draw_rect(nil, hpm, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(hpm), dw.Pixmap_height(hpm));

        for i = 0; (i < nrows) && (i+row < len(lines)); i++ {
            fileline := i + row - 1;
            y := i*fheight;
            dw.Color_background_set(dw.COLOR(1 + (fileline % 15)));
            dw.Color_foreground_set(dw.COLOR(fileline % 16));
            if hpma == nil {
                dw.Draw_text(nil, text1pm, 0, y, fmt.Sprintf("%6.6d", i+row));
            }
            thisline := lines[i+row];
            if len(thisline) > col {
               dw.Draw_text(nil, hpm, 0, y, thisline[col:]);
            }
        }
        if hpma == nil {
            text_expose(textbox1, 0, 0, 0, 0, nil);
            text_expose(textbox2, 0, 0, 0, 0, nil);
        }
    }
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_shapes(direct int, hpma dw.HPIXMAP) {
    var hpm, pixmap dw.HPIXMAP = nil, nil
    var window dw.HWND = nil
    if hpma != nil {
        hpm = hpma;
    } else {
        hpm = text2pm;
    }
    if direct == TRUE {
        window = textbox2;
    } else {
        pixmap = hpm;
    }

    width := dw.Pixmap_width(hpm);
    height := dw.Pixmap_height(hpm);

    x := []int{ 20, 180, 180, 230, 180, 180, 20 };
    y := []int{ 50, 50, 20, 70, 120, 90, 90 };

    image_x = dw.Spinbutton_get_pos(imagexspin);
    image_y = dw.Spinbutton_get_pos(imageyspin);
    image_stretch = dw.Checkbox_get(imagestretchcheck);

    dw.Color_foreground_set(dw.CLR_WHITE);
    dw.Draw_rect(window, pixmap, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, width, height);
    dw.Color_foreground_set(dw.CLR_DARKPINK);
    dw.Draw_rect(window, pixmap, dw.DRAW_FILL | dw.DRAW_NOAA, 10, 10, width - 20, height - 20);
    dw.Color_foreground_set(dw.CLR_GREEN);
    dw.Color_background_set(dw.CLR_DARKRED);
    dw.Draw_text(window, pixmap, 10, 10, "This should be aligned with the edges.");
    dw.Color_foreground_set(dw.CLR_YELLOW);
    dw.Draw_line(window, pixmap, width - 10, 10, 10, height - 10);
    dw.Color_foreground_set(dw.CLR_BLUE);
    dw.Draw_polygon(window, pixmap, dw.DRAW_FILL, x, y);
    dw.Color_foreground_set(dw.CLR_BLACK);
    dw.Draw_rect(window, pixmap, dw.DRAW_FILL | dw.DRAW_NOAA, 80, 80, 80, 40);
    dw.Color_foreground_set(dw.CLR_CYAN);
    /* Bottom right corner */
    dw.Draw_arc(window, pixmap, 0, width - 30, height - 30, width - 10, height - 30, width - 30, height - 10);
    /* Top right corner */
    dw.Draw_arc(window, pixmap, 0, width - 30, 30, width - 30, 10, width - 10, 30);
    /* Bottom left corner */
    dw.Draw_arc(window, pixmap, 0, 30, height - 30, 30, height - 10, 10, height - 30);
    /* Full circle in the left top area */
    dw.Draw_arc(window, pixmap, dw.DRAW_FULL, 120, 100, 80, 80, 160, 120);
    if image != nil {
        if image_stretch == TRUE {
            dw.Pixmap_stretch_bitblt(window, pixmap, 10, 10, width - 20, height - 20, nil, image, 0, 0, dw.Pixmap_width(image), dw.Pixmap_height(image));
        } else {
            dw.Pixmap_bitblt(window, pixmap, image_x, image_y, dw.Pixmap_width(image), dw.Pixmap_height(image), nil, image, 0, 0);
        }
    }

    /* If we aren't drawing direct do a bitblt */
    if direct == FALSE && hpma == nil {
        text_expose(textbox2, 0, 0, 0, 0, nil);
    }
}

func update_render() {
    switch render_type {
        case 0:
            draw_shapes(FALSE, nil);
        case 1:
            draw_shapes(TRUE, nil);
        case 2:
            draw_file(current_row, current_col, rows, font_height, nil);
    }
}

func draw_page(print dw.HPRINT, pixmap dw.HPIXMAP, page_num int, data unsafe.Pointer) int {
   dw.Pixmap_set_font(pixmap, FIXEDFONT);
   if page_num == 0 {
       draw_shapes(FALSE, pixmap);
   } else if page_num == 1 {
       /* If we have a file to display... */
       if len(current_file) > 0 {
           /* Calculate new dimensions */
           _, fheight := dw.Font_text_extents_get(nil, pixmap, "(g");
           nrows := int(dw.Pixmap_height(pixmap) / fheight);

           /* Do the actual drawing */
           draw_file(0, 0, nrows, fheight, pixmap);
       } else {
           /* We don't have a file so center an error message on the page */
           var text = "No file currently selected!";

           /* Get the font size for this printer context... */
           fwidth, fheight := dw.Font_text_extents_get(nil, pixmap, text);

           posx := int(dw.Pixmap_width(pixmap) - fwidth)/2;
           posy := int(dw.Pixmap_height(pixmap) - fheight)/2;

           dw.Color_foreground_set(dw.CLR_BLACK);
           dw.Color_background_set(dw.CLR_WHITE);
           dw.Draw_text(nil, pixmap, posx, posy, text);
       }
   }
   return TRUE;
}

func print_callback(window dw.HWND, data unsafe.Pointer) int {
   print := dw.Print_new("DWTest Job", 0, 2, unsafe.Pointer(&draw_page_func), nil);
   dw.Print_run(print, 0);
   return FALSE;
}


/* This gets called when a part of the graph needs to be repainted. */
func text_expose(hwnd dw.HWND, x int, y int, width int, height int, data unsafe.Pointer) int {
    if render_type != 1 {
        var hpm dw.HPIXMAP

        if hwnd == textbox1 {
            hpm = text1pm;
        } else if hwnd == textbox2 {
            hpm = text2pm;
        } else {
            return TRUE;
        }

        width = dw.Pixmap_width(hpm);
        height = dw.Pixmap_height(hpm);

        dw.Pixmap_bitblt(hwnd, nil, 0, 0, width, height, nil, hpm, 0, 0);
        dw.Flush();
    } else {
        update_render();
    }
    return TRUE;
}

/* Handle size change of the main render window */
func configure_event(hwnd dw.HWND, width int, height int, data unsafe.Pointer) int {
    old1 := text1pm;
    old2 := text2pm;
    depth := dw.Color_depth_get();

    rows = height / font_height;
    cols = width / font_width;

    /* Create new pixmaps with the current sizes */
    text1pm = dw.Pixmap_new(textbox1, uint(font_width*(width1)), uint(height), depth);
    text2pm = dw.Pixmap_new(textbox2, uint(width), uint(height), depth);

    /* Make sure the side area is cleared */
    dw.Color_foreground_set(dw.CLR_WHITE);
    dw.Draw_rect(nil, text1pm, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(text1pm), dw.Pixmap_height(text1pm));

   /* Destroy the old pixmaps */
    dw.Pixmap_destroy(old1);
    dw.Pixmap_destroy(old2);

    /* Update scrollbar ranges with new values */
    dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols));
    dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows));

    /* Redraw the window */
    update_render();
    return TRUE;
}

func refresh_callback(window dw.HWND, data unsafe.Pointer) int {
    update_render();
    return FALSE;
}

func render_select_event_callback(window dw.HWND, index int, data unsafe.Pointer) int {
    if index != render_type {
        if index == 2 {
            dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols));
            dw.Scrollbar_set_pos(hscrollbar, 0);
            dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows));
            dw.Scrollbar_set_pos(vscrollbar, 0);
            current_col = 0;
            current_row = 0;
        } else {
            dw.Scrollbar_set_range(hscrollbar, 0, 0);
            dw.Scrollbar_set_pos(hscrollbar, 0);
            dw.Scrollbar_set_range(vscrollbar, 0, 0);
            dw.Scrollbar_set_pos(vscrollbar, 0);
        }
        render_type = index;
        update_render();
    }
    return FALSE;
}

/* Callback to handle user selection of the scrollbar position */
func scrollbar_valuechanged_callback(hwnd dw.HWND, value int, data unsafe.Pointer) int {
    if data != nil {
        stext := dw.HWND(data);

        if hwnd == vscrollbar {
            current_row = value;
        } else {
            current_col = value;
        }
        dw.Window_set_text(stext, fmt.Sprintf("Row:%d Col:%d Lines:%d Cols:%d", current_row, current_col, len(lines), max_linewidth));
        update_render();
    }
    return FALSE;
}

func resolve_keyname(vk int) string {
    var keyname string = "<unknown>"

    switch vk {
        case  dw.VK_LBUTTON : keyname =  "VK_LBUTTON";
        case  dw.VK_RBUTTON : keyname =  "VK_RBUTTON";
        case  dw.VK_CANCEL  : keyname =  "VK_CANCEL";
        case  dw.VK_MBUTTON : keyname =  "VK_MBUTTON";
        case  dw.VK_TAB     : keyname =  "VK_TAB";
        case  dw.VK_CLEAR   : keyname =  "VK_CLEAR";
        case  dw.VK_RETURN  : keyname =  "VK_RETURN";
        case  dw.VK_PAUSE   : keyname =  "VK_PAUSE";
        case  dw.VK_CAPITAL : keyname =  "VK_CAPITAL";
        case  dw.VK_ESCAPE  : keyname =  "VK_ESCAPE";
        case  dw.VK_SPACE   : keyname =  "VK_SPACE";
        case  dw.VK_PRIOR   : keyname =  "VK_PRIOR";
        case  dw.VK_NEXT    : keyname =  "VK_NEXT";
        case  dw.VK_END     : keyname =  "VK_END";
        case  dw.VK_HOME    : keyname =  "VK_HOME";
        case  dw.VK_LEFT    : keyname =  "VK_LEFT";
        case  dw.VK_UP      : keyname =  "VK_UP";
        case  dw.VK_RIGHT   : keyname =  "VK_RIGHT";
        case  dw.VK_DOWN    : keyname =  "VK_DOWN";
        case  dw.VK_SELECT  : keyname =  "VK_SELECT";
        case  dw.VK_PRINT   : keyname =  "VK_PRINT";
        case  dw.VK_EXECUTE : keyname =  "VK_EXECUTE";
        case  dw.VK_SNAPSHOT: keyname =  "VK_SNAPSHOT";
        case  dw.VK_INSERT  : keyname =  "VK_INSERT";
        case  dw.VK_DELETE  : keyname =  "VK_DELETE";
        case  dw.VK_HELP    : keyname =  "VK_HELP";
        case  dw.VK_LWIN    : keyname =  "VK_LWIN";
        case  dw.VK_RWIN    : keyname =  "VK_RWIN";
        case  dw.VK_NUMPAD0 : keyname =  "VK_NUMPAD0";
        case  dw.VK_NUMPAD1 : keyname =  "VK_NUMPAD1";
        case  dw.VK_NUMPAD2 : keyname =  "VK_NUMPAD2";
        case  dw.VK_NUMPAD3 : keyname =  "VK_NUMPAD3";
        case  dw.VK_NUMPAD4 : keyname =  "VK_NUMPAD4";
        case  dw.VK_NUMPAD5 : keyname =  "VK_NUMPAD5";
        case  dw.VK_NUMPAD6 : keyname =  "VK_NUMPAD6";
        case  dw.VK_NUMPAD7 : keyname =  "VK_NUMPAD7";
        case  dw.VK_NUMPAD8 : keyname =  "VK_NUMPAD8";
        case  dw.VK_NUMPAD9 : keyname =  "VK_NUMPAD9";
        case  dw.VK_MULTIPLY: keyname =  "VK_MULTIPLY";
        case  dw.VK_ADD     : keyname =  "VK_ADD";
        case  dw.VK_SEPARATOR: keyname = "VK_SEPARATOR";
        case  dw.VK_SUBTRACT: keyname =  "VK_SUBTRACT";
        case  dw.VK_DECIMAL : keyname =  "VK_DECIMAL";
        case  dw.VK_DIVIDE  : keyname =  "VK_DIVIDE";
        case  dw.VK_F1      : keyname =  "VK_F1";
        case  dw.VK_F2      : keyname =  "VK_F2";
        case  dw.VK_F3      : keyname =  "VK_F3";
        case  dw.VK_F4      : keyname =  "VK_F4";
        case  dw.VK_F5      : keyname =  "VK_F5";
        case  dw.VK_F6      : keyname =  "VK_F6";
        case  dw.VK_F7      : keyname =  "VK_F7";
        case  dw.VK_F8      : keyname =  "VK_F8";
        case  dw.VK_F9      : keyname =  "VK_F9";
        case  dw.VK_F10     : keyname =  "VK_F10";
        case  dw.VK_F11     : keyname =  "VK_F11";
        case  dw.VK_F12     : keyname =  "VK_F12";
        case  dw.VK_F13     : keyname =  "VK_F13";
        case  dw.VK_F14     : keyname =  "VK_F14";
        case  dw.VK_F15     : keyname =  "VK_F15";
        case  dw.VK_F16     : keyname =  "VK_F16";
        case  dw.VK_F17     : keyname =  "VK_F17";
        case  dw.VK_F18     : keyname =  "VK_F18";
        case  dw.VK_F19     : keyname =  "VK_F19";
        case  dw.VK_F20     : keyname =  "VK_F20";
        case  dw.VK_F21     : keyname =  "VK_F21";
        case  dw.VK_F22     : keyname =  "VK_F22";
        case  dw.VK_F23     : keyname =  "VK_F23";
        case  dw.VK_F24     : keyname =  "VK_F24";
        case  dw.VK_NUMLOCK : keyname =  "VK_NUMLOCK";
        case  dw.VK_SCROLL  : keyname =  "VK_SCROLL";
        case  dw.VK_LSHIFT  : keyname =  "VK_LSHIFT";
        case  dw.VK_RSHIFT  : keyname =  "VK_RSHIFT";
        case  dw.VK_LCONTROL: keyname =  "VK_LCONTROL";
        case  dw.VK_RCONTROL: keyname =  "VK_RCONTROL";
    }
    return keyname;
}

func resolve_keymodifiers(mask int) string {
    if (mask & dw.KC_CTRL) == dw.KC_CTRL && (mask & dw.KC_SHIFT) == dw.KC_SHIFT && (mask & dw.KC_ALT) == dw.KC_ALT {
        return "KC_CTRL KC_SHIFT KC_ALT";
    } else if (mask & dw.KC_CTRL) == dw.KC_CTRL && (mask & dw.KC_SHIFT) == dw.KC_SHIFT {
        return "KC_CTRL KC_SHIFT";
    } else if (mask & dw.KC_CTRL) == dw.KC_CTRL && (mask & dw.KC_ALT) == dw.KC_ALT {
        return "KC_CTRL KC_ALT";
    } else if (mask & dw.KC_SHIFT) == dw.KC_SHIFT && (mask & dw.KC_ALT) == dw.KC_ALT {
        return "KC_SHIFT KC_ALT";
    } else if (mask & dw.KC_SHIFT) == dw.KC_SHIFT {
        return "KC_SHIFT";
    } else if (mask & dw.KC_CTRL) == dw.KC_CTRL {
        return "KC_CTRL";
    } else if (mask & dw.KC_ALT) == dw.KC_ALT {
        return "KC_ALT";
    }
    return "none";
}

func keypress_callback(window dw.HWND, ch uint8, vk int, state int, data unsafe.Pointer, utf8 string) int {
    var message string

    if ch != 0 {
        message = fmt.Sprintf("Key: %c(%d) Modifiers: %s(%d) utf8 %s", ch, ch, resolve_keymodifiers(state), state,  utf8);
    } else {
        message = fmt.Sprintf("Key: %s(%d) Modifiers: %s(%d) utf8 %s", resolve_keyname(vk), vk, resolve_keymodifiers(state), state, utf8);
    }
    dw.Window_set_text(status1, message);
    return FALSE;
}

// Page 3 and 4 Callbacks
func item_enter_cb(window dw.HWND, text string, data unsafe.Pointer) int {
    message := fmt.Sprintf("DW_SIGNAL_ITEM_ENTER: Window: %x Text: %s", uintptr(unsafe.Pointer(window)), text);
    dw.Window_set_text(dw.HWND(data), message);
    return FALSE;
}

func item_context_cb(window dw.HWND, text string, x int, y int, data unsafe.Pointer, itemdata unsafe.Pointer) int {
    message := fmt.Sprintf("DW_SIGNAL_ITEM_CONTEXT: Window: %x Text: %s x: %d y: %d Itemdata: %x", uintptr(unsafe.Pointer(window)), 
          text, x, y, uintptr(itemdata));
    dw.Window_set_text(dw.HWND(data), message);
    return FALSE;
}

func list_select_cb(window dw.HWND, item int, data unsafe.Pointer) int {
    message := fmt.Sprintf("DW_SIGNAL_LIST_SELECT: Window: %x Item: %d", uintptr(unsafe.Pointer(window)), item);
    dw.Window_set_text(dw.HWND(data), message);
    return FALSE;
}

func item_select_cb(window dw.HWND, item dw.HTREEITEM, text string, data unsafe.Pointer, itemdata unsafe.Pointer) int {
    message := fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", uintptr(unsafe.Pointer(window)),
            uintptr(unsafe.Pointer(item)), text, uintptr(itemdata));
    dw.Window_set_text(dw.HWND(data), message);
    return FALSE;
}

func container_select_cb(window dw.HWND, item dw.HTREEITEM, text string, data unsafe.Pointer, itemdata unsafe.Pointer)  int {
    message := fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", uintptr(unsafe.Pointer(window)),
            uintptr(unsafe.Pointer(item)), text, uintptr(itemdata));
    dw.Window_set_text(dw.HWND(data), message);
    mle_point = dw.Mle_import(container_mle, message, mle_point);
    str := dw.Container_query_start(container, dw.CRA_SELECTED);
    for len(str) > 0 {
        mle_point = dw.Mle_import( container_mle, fmt.Sprintf("Selected: %s\r\n", str), mle_point);
        str = dw.Container_query_next(container, dw.CRA_SELECTED);
    }
    /* Make the last inserted point the cursor location */
    dw.Mle_set_cursor(container_mle, mle_point);
    /* set the details of item 0 to new data */
    dw.Filesystem_change_file(container, 0, "new data", fileicon);
    dw.Filesystem_change_item_ulong(container, 1, 0, 999);
    return FALSE;
}

func combobox_select_event_callback(window dw.HWND, index int, data unsafe.Pointer) int {
    fmt.Printf("got combobox_select_event for index: %d, iteration: %d\n", index, iteration);
    iteration++;
    return FALSE;
}

func column_click_cb(window dw.HWND, column_num int, data unsafe.Pointer) int {
    var stype = "Unknown";

    if column_num == 0 {
        stype = "Filename";
    } else {
        column_type := dw.Filesystem_get_column_type(window, column_num-1);
        if column_type == dw.CFA_STRING {
            stype = "String";
        } else if column_type == dw.CFA_ULONG {
            stype = "ULong";
        } else if column_type == dw.CFA_DATE {
            stype = "Date";
        } else if  column_type == dw.CFA_TIME {
            stype = "Time";
        } else if column_type == dw.CFA_BITMAPORICON {
            stype = "BitmapOrIcon";
        }
    }
    message := fmt.Sprintf("DW_SIGNAL_COLUMN_CLICK: Window: %x Column: %d Type: %s Itemdata: %x", uintptr(unsafe.Pointer(window)),
            column_num, stype);
    dw.Window_set_text(dw.HWND(data), message);
    return FALSE;
}

// Page 5 Callbacks
func button_callback(window dw.HWND, data unsafe.Pointer) int {
    idx := dw.Listbox_selected(combobox1);
    buf1 := dw.Listbox_get_text(combobox1, idx);
    idx = dw.Listbox_selected( combobox2 );
    buf2 := dw.Listbox_get_text(combobox2, idx);
    y, m, d := dw.Calendar_get_date(cal);
    spvalue := dw.Spinbutton_get_pos(spinbutton);
    message := fmt.Sprintf("spinbutton: %d\ncombobox1: \"%s\"\ncombobox2: \"%s\"\ncalendar: %d-%d-%d",
                  spvalue,
                  buf1, buf2,
                  y, m, d );
    dw.Messagebox( "Values", dw.MB_OK | dw.MB_INFORMATION, message);
    return FALSE;
}

var isfoldericon bool = true

func bitmap_toggle_callback(window dw.HWND, data unsafe.Pointer) int {
    if isfoldericon == true {
       isfoldericon = false;
       dw.Window_set_bitmap(window, 0, FILE_ICON_NAME);
       dw.Window_set_tooltip(window, "File Icon" );
    } else {
       isfoldericon = true;
       //dw.Window_set_bitmap_from_data(window, 0, folder_ico, sizeof(folder_ico));
       dw.Window_set_tooltip(window, "Folder Icon");
    }
    return FALSE;
}

func percent_button_box_callback(window dw.HWND, data unsafe.Pointer) int {
    dw.Percent_set_pos(percent, dw.PERCENT_INDETERMINATE);
    return FALSE;
}

func change_color_red_callback(window dw.HWND, data unsafe.Pointer) int {
    dw.Window_set_color(buttonsbox, dw.CLR_RED, dw.CLR_RED);
    return FALSE;
}

func change_color_yellow_callback(window dw.HWND, data unsafe.Pointer) int {
    dw.Window_set_color(buttonsbox, dw.CLR_YELLOW, dw.CLR_YELLOW);
    return FALSE;
}

/* Callback to handle user selection of the spinbutton position */
func spinbutton_valuechanged_callback(hwnd dw.HWND, value int, data unsafe.Pointer) int {
    dw.Messagebox("DWTest", dw.MB_OK, fmt.Sprintf("New value from spinbutton: %d\n", value));
    return FALSE;
}

/* Callback to handle user selection of the slider position */
func slider_valuechanged_callback(hwnd dw.HWND, value int, data unsafe.Pointer) int {
    dw.Percent_set_pos(percent, uint(value * 10));
    return FALSE;
}

// Page 8 Callbacks
func scrollbox_button_callback(window dw.HWND, data unsafe.Pointer) int {
    pos := dw.Scrollbox_get_pos(scrollbox, dw.VERT);
    rng := dw.Scrollbox_get_range(scrollbox, dw.VERT);
    fmt.Printf("Pos %d Range %d\n", pos, rng);
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
var text_expose_func = text_expose;
var configure_event_func = configure_event;
var motion_notify_event_func = motion_notify_event;
var show_window_callback_func = show_window_callback;
var context_menu_event_func = context_menu_event;
var refresh_callback_func = refresh_callback;
var render_select_event_callback_func = render_select_event_callback;
var scrollbar_valuechanged_callback_func = scrollbar_valuechanged_callback;
var keypress_callback_func = keypress_callback;
var item_enter_cb_func = item_enter_cb;
var item_context_cb_func = item_context_cb;
var list_select_cb_func = list_select_cb; 
var item_select_cb_func = item_select_cb;
var column_click_cb_func = column_click_cb;
var container_select_cb_func = container_select_cb;
var combobox_select_event_callback_func = combobox_select_event_callback;
var scrollbox_button_callback_func = scrollbox_button_callback;
var button_callback_func = button_callback;
var bitmap_toggle_callback_func = bitmap_toggle_callback;
var percent_button_box_callback_func = percent_button_box_callback;
var change_color_red_callback_func = change_color_red_callback;
var change_color_yellow_callback_func = change_color_yellow_callback;
var spinbutton_valuechanged_callback_func = spinbutton_valuechanged_callback;
var slider_valuechanged_callback_func = slider_valuechanged_callback;
var print_callback_func = print_callback;
var draw_page_func = draw_page;

var checkable_string = "checkable";
var noncheckable_string = "non-checkable";

// Create the menu
func menu_add() {
    mainmenubar := dw.Menubar_new(mainwindow);
    /* add menus to the menubar */
    menu := dw.Menu_new(0);
    menuitem := dw.Menu_append_item(menu, "~Quit", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
    /*
     * Add the "File" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~File", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);

    changeable_menu := dw.Menu_new(0);
    checkable_menuitem = dw.Menu_append_item(changeable_menu, "~Checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    dw.Signal_connect(checkable_menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&menu_callback_func), unsafe.Pointer(&checkable_string));
    noncheckable_menuitem = dw.Menu_append_item(changeable_menu, "~Non-checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(noncheckable_menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&menu_callback_func), unsafe.Pointer(&noncheckable_string));
    dw.Menu_append_item(changeable_menu, "~Disabled menu Item", dw.MENU_AUTO, dw.MIS_DISABLED | dw.MIS_CHECKED, dw.TRUE, dw.TRUE, dw.NOMENU);
    /* seperator */
    dw.Menu_append_item(changeable_menu, dw.MENU_SEPARATOR, dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem = dw.Menu_append_item(changeable_menu, "~Menu Items Disabled", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&menutoggle_callback_func), nil);
    /*
     * Add the "Menu" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~Menu", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, changeable_menu);

    menu = dw.Menu_new(0);
    menuitem = dw.Menu_append_item(menu, "~About", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, unsafe.Pointer(&helpabout_callback_func), unsafe.Pointer(mainwindow));
    /*
     * Add the "Help" menu to the menubar...
     */
    dw.Menu_append_item(mainmenubar, "~Help", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);
}

// Create Page 1
func archive_add() {
    lbbox := dw.Box_new(dw.VERT, 10);

    dw.Box_pack_start(notebookbox1, lbbox, 150, 70, dw.TRUE, dw.TRUE, 0);

    /* Copy and Paste */
    browsebox := dw.Box_new(dw.HORZ, 0);

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

    dw.Window_set_style(stext, dw.DT_VCENTER, dw.DT_VCENTER);

    dw.Box_pack_start(lbbox, stext, 130, 15, dw.TRUE, dw.TRUE, 2);

    browsebox = dw.Box_new(dw.HORZ, 0);

    dw.Box_pack_start(lbbox, browsebox, 0, 0, dw.TRUE, dw.TRUE, 0);

    entryfield = dw.Entryfield_new("", 100);

    dw.Entryfield_set_limit(entryfield, 260);

    dw.Box_pack_start(browsebox, entryfield, 100, 15, dw.TRUE, dw.TRUE, 4);

    browsefilebutton := dw.Button_new("Browse File", 1001);

    dw.Box_pack_start(browsebox, browsefilebutton, 40, 15, dw.TRUE, dw.TRUE, 0);

    browsefolderbutton := dw.Button_new("Browse Folder", 1001);

    dw.Box_pack_start(browsebox, browsefolderbutton, 40, 15, dw.TRUE, dw.TRUE, 0);

    dw.Window_set_color(browsebox, dw.CLR_PALEGRAY, dw.CLR_PALEGRAY);
    dw.Window_set_color(stext, dw.CLR_BLACK, dw.CLR_PALEGRAY);

    /* Buttons */
    buttonbox := dw.Box_new(dw.HORZ, 10);

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
    dw.Window_set_color(lbbox, dw.CLR_DARKCYAN, dw.CLR_PALEGRAY);
    dw.Window_set_color(buttonbox, dw.CLR_DARKCYAN, dw.CLR_PALEGRAY);
    dw.Window_set_color(okbutton, dw.CLR_PALEGRAY, dw.CLR_DARKCYAN);

    dw.Signal_connect(browsefilebutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&browse_file_callback_func), nil);
    dw.Signal_connect(browsefolderbutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&browse_folder_callback_func), nil);
    dw.Signal_connect(copybutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&copy_clicked_callback_func), unsafe.Pointer(copypastefield));
    dw.Signal_connect(pastebutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&paste_clicked_callback_func), unsafe.Pointer(copypastefield));
    dw.Signal_connect(okbutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&beep_callback_func), nil);
    dw.Signal_connect(cancelbutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
    dw.Signal_connect(cursortogglebutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&cursortoggle_callback_func), unsafe.Pointer(mainwindow));
    dw.Signal_connect(colorchoosebutton, dw.SIGNAL_CLICKED, unsafe.Pointer(&colorchoose_callback_func), unsafe.Pointer(mainwindow));
}

// Create Page 2
func text_add() {
    depth := dw.Color_depth_get();

    /* create a box to pack into the notebook page */
    pagebox := dw.Box_new(dw.HORZ, 2);
    dw.Box_pack_start(notebookbox2, pagebox, 0, 0, dw.TRUE, dw.TRUE, 0);
    /* now a status area under this box */
    hbox := dw.Box_new(dw.HORZ, 1);
    dw.Box_pack_start(notebookbox2, hbox, 100, 20, dw.TRUE, dw.FALSE, 1);
    status1 = dw.Status_text_new("", 0);
    dw.Box_pack_start(hbox, status1, 100, -1, dw.TRUE, dw.FALSE, 1);
    status2 = dw.Status_text_new("", 0);
    dw.Box_pack_start(hbox, status2, 100, -1, dw.TRUE, dw.FALSE, 1);
    /* a box with combobox and button */
    hbox = dw.Box_new(dw.HORZ, 1);
    dw.Box_pack_start(notebookbox2, hbox, 100, 25, dw.TRUE, dw.FALSE, 1);
    rendcombo := dw.Combobox_new("Shapes Double Buffered", 0);
    dw.Box_pack_start(hbox, rendcombo, 80, 25, dw.TRUE, dw.FALSE, 0);
    dw.Listbox_append(rendcombo, "Shapes Double Buffered");
    dw.Listbox_append(rendcombo, "Shapes Direct");
    dw.Listbox_append(rendcombo, "File Display");
    label := dw.Text_new("Image X:", 100);
    dw.Window_set_style(label, dw.DT_VCENTER | dw.DT_CENTER, dw.DT_VCENTER | dw.DT_CENTER);
    dw.Box_pack_start( hbox, label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imagexspin = dw.Spinbutton_new("20", 1021);
    dw.Box_pack_start(hbox, imagexspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    label = dw.Text_new("Y:", 100);
    dw.Window_set_style(label, dw.DT_VCENTER | dw.DT_CENTER, dw.DT_VCENTER | dw.DT_CENTER);
    dw.Box_pack_start(hbox, label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imageyspin = dw.Spinbutton_new("20", 1021);
    dw.Box_pack_start(hbox, imageyspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    dw.Spinbutton_set_limits(imagexspin, 2000, 0);
    dw.Spinbutton_set_limits(imageyspin, 2000, 0);
    dw.Spinbutton_set_pos(imagexspin, 20);
    dw.Spinbutton_set_pos(imageyspin, 20);
    imagestretchcheck = dw.Checkbox_new("Stretch", 1021);
    dw.Box_pack_start(hbox, imagestretchcheck, -1, 25, dw.FALSE, dw.FALSE, 0);

    button1 := dw.Button_new("Refresh", 1223);
    dw.Box_pack_start(hbox, button1, 100, 25, dw.FALSE, dw.FALSE, 0);
    button2 := dw.Button_new("Print", 1224);
    dw.Box_pack_start(hbox, button2, 100, 25, dw.FALSE, dw.FALSE, 0);

    /* Pre-create the scrollbars so we can query their sizes */
    vscrollbar = dw.Scrollbar_new(dw.VERT, 50);
    hscrollbar = dw.Scrollbar_new(dw.HORZ, 50);
    vscrollbarwidth, _ := dw.Window_get_preferred_size(vscrollbar);
    _, hscrollbarheight := dw.Window_get_preferred_size(hscrollbar);

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
    dw.Window_set_font(textbox1, FIXEDFONT);
    font_width, font_height = dw.Font_text_extents_get(textbox1, nil, "(g");
    font_width = font_width / 2;
    vscrollbox := dw.Box_new(dw.VERT, 0);
    dw.Box_pack_start(vscrollbox, textbox1, font_width * width1, font_height * rows, dw.FALSE, dw.TRUE, 0);
    dw.Box_pack_start(vscrollbox, nil, font_width * (width1 + 1), hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    /* pack empty space 1 character wide */
    dw.Box_pack_start(pagebox, nil, font_width, 0, dw.FALSE, dw.TRUE, 0);

    /* create box for filecontents and horz scrollbar */
    textboxA := dw.Box_new(dw.VERT, 0);
    dw.Box_pack_start(pagebox, textboxA, 0, 0, dw.TRUE, dw.TRUE, 0);

    /* create render box for filecontents pixmap */
    textbox2 = dw.Render_new(101);
    dw.Box_pack_start(textboxA, textbox2, 10, 10, dw.TRUE, dw.TRUE, 0);
    dw.Window_set_font(textbox2, FIXEDFONT);
    /* create horizonal scrollbar */
    dw.Box_pack_start(textboxA, hscrollbar, -1, -1, dw.TRUE, dw.FALSE, 0);

    /* create vertical scrollbar */
    vscrollbox = dw.Box_new(dw.VERT, 0);
    dw.Box_pack_start(vscrollbox, vscrollbar, -1, -1, dw.FALSE, dw.TRUE, 0);
    /* Pack an area of empty space 14x14 pixels */
    dw.Box_pack_start(vscrollbox, nil, vscrollbarwidth, hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    text1pm = dw.Pixmap_new(textbox1, uint(font_width * width1), uint(font_height * rows), depth);
    text2pm = dw.Pixmap_new(textbox2, uint(font_width * cols), uint(font_height * rows), depth);
    image = dw.Pixmap_new_from_file(textbox2, "test");
    if image == nil {
        image = dw.Pixmap_new_from_file(textbox2, "~/test");
    }
    if image != nil {
        dw.Pixmap_set_transparent_color(image, dw.CLR_WHITE);
    }

    dw.Messagebox("DWTest", dw.MB_OK | dw.MB_INFORMATION, fmt.Sprintf("Width: %d Height: %d\n", font_width, font_height));
    dw.Draw_rect(nil, text1pm, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, font_width * width1, font_height * rows);
    dw.Draw_rect(nil, text2pm, dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, font_width * cols, font_height * rows);
    dw.Signal_connect(textbox1, dw.SIGNAL_BUTTON_PRESS, unsafe.Pointer(&context_menu_event_func), nil);
    dw.Signal_connect(textbox1, dw.SIGNAL_EXPOSE, unsafe.Pointer(&text_expose_func), nil);
    dw.Signal_connect(textbox2, dw.SIGNAL_EXPOSE, unsafe.Pointer(&text_expose_func), nil);
    dw.Signal_connect(textbox2, dw.SIGNAL_CONFIGURE, unsafe.Pointer(&configure_event_func), unsafe.Pointer(text2pm));
    dw.Signal_connect(textbox2, dw.SIGNAL_MOTION_NOTIFY, unsafe.Pointer(&motion_notify_event_func), unsafe.Pointer(uintptr(1)));
    dw.Signal_connect(textbox2, dw.SIGNAL_BUTTON_PRESS, unsafe.Pointer(&motion_notify_event_func), nil);
    dw.Signal_connect(hscrollbar, dw.SIGNAL_VALUE_CHANGED, unsafe.Pointer(&scrollbar_valuechanged_callback_func), unsafe.Pointer(status1));
    dw.Signal_connect(vscrollbar, dw.SIGNAL_VALUE_CHANGED, unsafe.Pointer(&scrollbar_valuechanged_callback_func), unsafe.Pointer(status1));
    dw.Signal_connect(imagestretchcheck, dw.SIGNAL_CLICKED, unsafe.Pointer(&refresh_callback_func), nil);
    dw.Signal_connect(button1, dw.SIGNAL_CLICKED, unsafe.Pointer(&refresh_callback_func), nil);
    dw.Signal_connect(button2, dw.SIGNAL_CLICKED, unsafe.Pointer(&print_callback_func), nil);
    dw.Signal_connect(rendcombo, dw.SIGNAL_LIST_SELECT, unsafe.Pointer(&render_select_event_callback_func), nil);
    dw.Signal_connect(mainwindow, dw.SIGNAL_KEY_PRESS, unsafe.Pointer(&keypress_callback_func), nil);

    dw.Taskbar_insert(textbox1, fileicon, "DWTest");
}

// Page 3
func tree_add() {
    /* create a box to pack into the notebook page */
    listbox := dw.Listbox_new(1024, TRUE);
    dw.Box_pack_start(notebookbox3, listbox, 500, 200, TRUE, TRUE, 0);
    dw.Listbox_append(listbox, "Test 1");
    dw.Listbox_append(listbox, "Test 2");
    dw.Listbox_append(listbox, "Test 3");
    dw.Listbox_append(listbox, "Test 4");
    dw.Listbox_append(listbox, "Test 5");

    /* now a tree area under this box */
    tree = dw.Tree_new(101);
    dw.Box_pack_start(notebookbox3, tree, 500, 200, TRUE, TRUE, 1);

    /* and a status area to see whats going on */
    tree_status := dw.Status_text_new("", 0);
    dw.Box_pack_start(notebookbox3, tree_status, 100, -1, TRUE, FALSE, 1);

    /* set up our signal trappers... */
    dw.Signal_connect(tree, dw.SIGNAL_ITEM_CONTEXT, unsafe.Pointer(&item_context_cb_func), unsafe.Pointer(tree_status));
    dw.Signal_connect(tree, dw.SIGNAL_ITEM_SELECT, unsafe.Pointer(&item_select_cb_func), unsafe.Pointer(tree_status));

    t1 := dw.Tree_insert(tree, "tree folder 1", foldericon, nil, unsafe.Pointer(uintptr(1)));
    t2 := dw.Tree_insert(tree, "tree folder 2", foldericon, nil, unsafe.Pointer(uintptr(2)));
    dw.Tree_insert(tree, "tree file 1", fileicon, t1, unsafe.Pointer(uintptr(3)));
    dw.Tree_insert(tree, "tree file 2", fileicon, t1, unsafe.Pointer(uintptr(4)));
    dw.Tree_insert(tree, "tree file 3", fileicon, t2, unsafe.Pointer(uintptr(5)));
    dw.Tree_insert(tree, "tree file 4", fileicon, t2, unsafe.Pointer(uintptr(6)));
    dw.Tree_item_change(tree, t1, "tree folder 1", foldericon);
    dw.Tree_item_change(tree, t2, "tree folder 2", foldericon);
}

// Page 4
func container_add() {
    var z int
    titles := []string{ "Type", "Size", "Time", "Date" };
    flags := []uint{   dw.CFA_BITMAPORICON | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_ULONG | dw.CFA_RIGHT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_TIME | dw.CFA_CENTER | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_DATE | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR };


    /* create a box to pack into the notebook page */
    containerbox := dw.Box_new(dw.HORZ, 2);
    dw.Box_pack_start(notebookbox4, containerbox, 500, 200, TRUE, TRUE, 0);

    /* now a container area under this box */
    container = dw.Container_new(100, TRUE);
    dw.Box_pack_start(notebookbox4, container, 500, 200, TRUE, FALSE, 1);

    /* and a status area to see whats going on */
    container_status := dw.Status_text_new("", 0);
    dw.Box_pack_start(notebookbox4, container_status, 100, -1, TRUE, FALSE, 1);

    dw.Filesystem_set_column_title(container, "Test");
    dw.Filesystem_setup(container, flags, titles);
    dw.Container_set_stripe(container, dw.CLR_DEFAULT, dw.CLR_DEFAULT);
    containerinfo := dw.Container_alloc(container, 3);

    for z=0; z<3; z++ {
        var thisicon dw.HICN = fileicon; 
        
        if z == 0 {
             thisicon = foldericon;
        } 
        fmt.Printf("Initial: container: %x containerinfo: %x icon: %x\n", uintptr(unsafe.Pointer(container)),
                  uintptr(containerinfo), uintptr(unsafe.Pointer(thisicon)));
        dw.Filesystem_set_file(container, containerinfo, z, fmt.Sprintf("Filename %d", z+1), thisicon);
        dw.Filesystem_set_item_icon(container, containerinfo, 0, z, thisicon);
        dw.Filesystem_set_item_ulong(container, containerinfo, 1, z, uint(z*100));
        dw.Filesystem_set_item_time(container, containerinfo, 2, z, z+10, z+10, z+10);
        dw.Filesystem_set_item_date(container, containerinfo, 3, z, z+10, z+10, z+2000);
        dw.Container_set_row_title(containerinfo, z, fmt.Sprintf("Don't allocate from stack: Item: %d", z));
    }
    dw.Container_insert(container, containerinfo, 3);

    containerinfo = dw.Container_alloc(container, 1);
    dw.Filesystem_set_file(container, containerinfo, 0, "Yikes", foldericon);
    dw.Filesystem_set_item_icon(container, containerinfo, 0, 0, foldericon);
    dw.Filesystem_set_item_ulong(container, containerinfo, 1, 0, 324);
    dw.Filesystem_set_item_time(container, containerinfo, 2, 0, z+10, z+10, z+10);
    dw.Filesystem_set_item_date(container, containerinfo, 3, 0, z+10, z+10, z+2000);
    dw.Container_set_row_title(containerinfo, 0, "Extra");

    dw.Container_insert(container, containerinfo, 1);
    dw.Container_optimize(container);

    container_mle = dw.Mle_new(111);
    dw.Box_pack_start( containerbox, container_mle, 500, 200, TRUE, TRUE, 0);

    mle_point = dw.Mle_import(container_mle, "", -1);
    mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]", mle_point), mle_point);
    mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]abczxydefijkl", mle_point), mle_point);
    dw.Mle_delete(container_mle, 9, 3);
    mle_point = dw.Mle_import(container_mle, "gh", 12);
    newpoint, _ := dw.Mle_get_size(container_mle);
    mle_point = newpoint;
    mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]\r\n\r\n", mle_point), mle_point);
    dw.Mle_set_cursor(container_mle, mle_point);
    /* connect our event trappers... */
    dw.Signal_connect(container, dw.SIGNAL_ITEM_ENTER, unsafe.Pointer(&item_enter_cb_func), unsafe.Pointer(container_status));
    dw.Signal_connect(container, dw.SIGNAL_ITEM_CONTEXT, unsafe.Pointer(&item_context_cb_func), unsafe.Pointer(container_status));
    dw.Signal_connect(container, dw.SIGNAL_ITEM_SELECT, unsafe.Pointer(&container_select_cb_func), unsafe.Pointer(container_status));
    dw.Signal_connect(container, dw.SIGNAL_COLUMN_CLICK, unsafe.Pointer(&column_click_cb_func), unsafe.Pointer(container_status));
}

// Page 5
func buttons_add() {
    var i int;
    
    /* create a box to pack into the notebook page */
    buttonsbox = dw.Box_new(dw.VERT, 2);
    dw.Box_pack_start(notebookbox5, buttonsbox, 25, 200, TRUE, TRUE, 0);
    dw.Window_set_color(buttonsbox, dw.CLR_RED, dw.CLR_RED);

    calbox := dw.Box_new(dw.HORZ, 0);
    dw.Box_pack_start(notebookbox5, calbox, 500, 200, TRUE, TRUE, 1);
    cal = dw.Calendar_new(100);
    dw.Box_pack_start(calbox, cal, 180, 120, TRUE, TRUE, 0);
    /*
     dw_calendar_set_date(cal, 2001, 1, 1);
     */
    /*
     * Create our file toolbar boxes...
     */
    buttonboxperm = dw.Box_new(dw.VERT, 0);
    dw.Box_pack_start(buttonsbox, buttonboxperm, 25, 0, FALSE, TRUE, 2);
    dw.Window_set_color(buttonboxperm, dw.CLR_WHITE, dw.CLR_WHITE);
    abutton1 := dw.Bitmapbutton_new_from_file("Top Button", 0, FILE_ICON_NAME);
    dw.Box_pack_start(buttonboxperm, abutton1, 100, 30, FALSE, FALSE, 0);
    dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, unsafe.Pointer(&button_callback_func), nil);
    dw.Box_pack_start(buttonboxperm, nil, 25, 5, FALSE, FALSE, 0);
    abutton2 := dw.Bitmapbutton_new_from_file("Bottom", 0, FOLDER_ICON_NAME);
    dw.Box_pack_start(buttonsbox, abutton2, 25, 25, FALSE, FALSE, 0);
    dw.Signal_connect(abutton2, dw.SIGNAL_CLICKED, unsafe.Pointer(&button_callback_func), nil);
    dw.Window_set_bitmap(abutton2, 0, FILE_ICON_NAME);

    create_button(false);
    /* make a combobox */
    combox := dw.Box_new(dw.VERT, 2);
    dw.Box_pack_start(notebookbox5, combox, 25, 200, TRUE, FALSE, 0);
    combobox1 = dw.Combobox_new("fred", 0 ); /* no point in specifying an initial value */
    dw.Listbox_append(combobox1, "fred" );
    dw.Box_pack_start(combox, combobox1, -1, -1, TRUE, FALSE, 0);
    /*
     dw_window_set_text( combobox, "initial value");
     */
    dw.Signal_connect(combobox1, dw.SIGNAL_LIST_SELECT, unsafe.Pointer(&combobox_select_event_callback_func), nil);

    combobox2 = dw.Combobox_new("joe", 0); /* no point in specifying an initial value */
    dw.Box_pack_start(combox, combobox2, -1, -1, TRUE, FALSE, 0);
    /*
     dw_window_set_text( combobox, "initial value");
     */
    dw.Signal_connect(combobox2, dw.SIGNAL_LIST_SELECT, unsafe.Pointer(&combobox_select_event_callback_func), nil);
    /* add LOTS of items */
    fmt.Printf("before appending 500 items to combobox using dw_listbox_list_append()\n");
    text := make([]string, 500);
    for  i = 0; i < 500; i++ {
        text[i] = fmt.Sprintf("item %d", i);
    }
    dw.Listbox_list_append(combobox2, text);
    fmt.Printf("after appending 500 items to combobox\n");
    /* now insert a couple of items */
    dw.Listbox_insert(combobox2, "inserted item 2", 2 );
    dw.Listbox_insert(combobox2, "inserted item 5", 5 );
    /* make a spinbutton */
    spinbutton = dw.Spinbutton_new("", 0); /* no point in specifying text */
    dw.Box_pack_start(combox, spinbutton, -1, -1, TRUE, FALSE, 0);
    dw.Spinbutton_set_limits(spinbutton, 100, 1);
    dw.Spinbutton_set_pos(spinbutton, 30);
    dw.Signal_connect(spinbutton, dw.SIGNAL_VALUE_CHANGED, unsafe.Pointer(&spinbutton_valuechanged_callback_func), nil);
    /* make a slider */
    slider = dw.Slider_new(FALSE, 11, 0); /* no point in specifying text */
    dw.Box_pack_start(combox, slider, -1, -1, TRUE, FALSE, 0);
    dw.Signal_connect(slider, dw.SIGNAL_VALUE_CHANGED, unsafe.Pointer(&slider_valuechanged_callback_func), nil);
    /* make a percent */
    percent = dw.Percent_new(0);
    dw.Box_pack_start(combox, percent, -1, -1, TRUE, FALSE, 0);
}

func create_button(redraw bool) {
    filetoolbarbox := dw.Box_new(dw.VERT, 0);
    dw.Box_pack_start(buttonboxperm, filetoolbarbox, 0, 0, TRUE, TRUE, 0);

    abutton1 := dw.Bitmapbutton_new_from_file("Empty image. Should be under Top button", 0, "junk");
    dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0);
    dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, unsafe.Pointer(&change_color_red_callback_func), nil);
    dw.Box_pack_start(filetoolbarbox, nil, 25, 5, FALSE, FALSE, 0);

    //abutton1 = dw_bitmapbutton_new_from_data("A borderless bitmapbitton", 0, folder_ico, 1718 );
    abutton1 = dw.Bitmapbutton_new_from_file("A borderless bitmapbitton", 0, "junk");
    dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0);
    dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, unsafe.Pointer(&change_color_yellow_callback_func), nil);
    dw.Box_pack_start(filetoolbarbox, nil, 25, 5, FALSE, FALSE, 0);
    dw.Window_set_style(abutton1, dw.BS_NOBORDER, dw.BS_NOBORDER);

    //abutton1 = dw.Bitmapbutton_new_from_data("A button from data", 0, folder_ico, 1718 );
    abutton1 = dw.Bitmapbutton_new_from_file("A button from data", 0, "junk");
    dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0);
    dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, unsafe.Pointer(&percent_button_box_callback_func), nil);
    dw.Box_pack_start(filetoolbarbox, nil, 25, 5, FALSE, FALSE, 0);
    if redraw == true {
        dw.Window_redraw(filetoolbarbox);
        dw.Window_redraw(mainwindow);
    }
}

// Page 8
func scrollbox_add() {
   var i int;

    /* create a box to pack into the notebook page */
    scrollbox = dw.Scrollbox_new(dw.VERT, 0);
    dw.Box_pack_start(notebookbox8, scrollbox, 0, 0, TRUE, TRUE, 1);

    abutton1 := dw.Button_new("Show Adjustments", 0);
    dw.Box_pack_start(scrollbox, abutton1, -1, 30, FALSE, FALSE, 0 );
    dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, unsafe.Pointer(&scrollbox_button_callback_func), nil);

    for i = 0; i < MAX_WIDGETS; i++ {
        tmpbox := dw.Box_new(dw.HORZ, 0);
        dw.Box_pack_start(scrollbox, tmpbox, 0, 24, TRUE, FALSE, 2);
        label := dw.Text_new(fmt.Sprintf("Label %d", i), 0 );
        dw.Box_pack_start(tmpbox, label, 0, 20, TRUE, FALSE, 0);
        item := dw.Entryfield_new(fmt.Sprintf("Entry %d", i), uint(i));
        dw.Box_pack_start(tmpbox, item, 0, 20, TRUE, FALSE, 0);
    }
}

func main() {
   /* Pick an approriate font for our platform */
   if runtime.GOOS == "windows" {
      FIXEDFONT = "10.Lucida Console";
   } else if runtime.GOOS == "darwin" {
      FIXEDFONT = "9.Monaco";
   }
    
   /* Initialize the Dynamic Windows engine */
   dw.Init(dw.TRUE);

   /* Create our window */
   mainwindow = dw.Window_new(dw.DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", dw.FCF_SYSMENU | dw.FCF_TITLEBAR | dw.FCF_TASKLIST | dw.FCF_DLGBORDER | dw.FCF_SIZEBORDER | dw.FCF_MINMAX);

   menu_add();

   notebookbox := dw.Box_new(dw.VERT, 5);
   dw.Box_pack_start(mainwindow, notebookbox, 0, 0, dw.TRUE, dw.TRUE, 0);

   foldericon = dw.Icon_load_from_file(FOLDER_ICON_NAME);
   fileicon = dw.Icon_load_from_file(FILE_ICON_NAME);

   notebook := dw.Notebook_new(1, dw.TRUE);
   dw.Box_pack_start(notebookbox, notebook, 100, 100, dw.TRUE, dw.TRUE, 0);
   dw.Signal_connect(notebook, dw.SIGNAL_SWITCH_PAGE, unsafe.Pointer(&switch_page_callback_func), nil);

   notebookbox1 = dw.Box_new(dw.VERT, 5);
   notebookpage1 := dw.Notebook_page_new(notebook, 0, dw.TRUE);
   dw.Notebook_pack(notebook, notebookpage1, notebookbox1);
   dw.Notebook_page_set_text(notebook, notebookpage1, "buttons and entry");
   archive_add();

   notebookbox2 = dw.Box_new(dw.VERT, 5);
   notebookpage2 := dw.Notebook_page_new(notebook, 1, dw.FALSE);
   dw.Notebook_pack(notebook, notebookpage2, notebookbox2);
   dw.Notebook_page_set_text(notebook, notebookpage2, "render");
   text_add();

   notebookbox3 = dw.Box_new(dw.VERT, 5);
   notebookpage3 := dw.Notebook_page_new(notebook, 1, dw.FALSE);
   dw.Notebook_pack(notebook, notebookpage3, notebookbox3);
   dw.Notebook_page_set_text(notebook, notebookpage3, "tree");
   tree_add();
   
   notebookbox4 = dw.Box_new(dw.VERT, 5);
   notebookpage4 := dw.Notebook_page_new(notebook, 1, FALSE);
   dw.Notebook_pack(notebook, notebookpage4, notebookbox4);
   dw.Notebook_page_set_text(notebook, notebookpage4, "container");
   container_add();

   notebookbox5 = dw.Box_new(dw.VERT, 5);
   notebookpage5 := dw.Notebook_page_new(notebook, 1, FALSE);
   dw.Notebook_pack(notebook, notebookpage5, notebookbox5);
   dw.Notebook_page_set_text(notebook, notebookpage5, "buttons");
   buttons_add();

/* DEPRECATED
   notebookbox6 = dw.Box_new(dw.VERT, 5);
   notebookpage6 := dw.Notebook_page_new( notebook, 1, FALSE );
   dw.Notebook_pack(notebook, notebookpage6, notebookbox6);
   dw.Notebook_page_set_text(notebook, notebookpage6, "mdi");
   mdi_add();
*/

   notebookbox7 = dw.Box_new(dw.VERT, 6);
   notebookpage7 := dw.Notebook_page_new(notebook, 1, FALSE);
   dw.Notebook_pack(notebook, notebookpage7, notebookbox7);
   dw.Notebook_page_set_text(notebook, notebookpage7, "html");
   
   rawhtml := dw.Html_new(1001);
   if rawhtml != nil {
       dw.Box_pack_start(notebookbox7, rawhtml, 0, 100, TRUE, FALSE, 0);
       dw.Html_raw(rawhtml, "<html><body><center><h1>dwtest</h1></center></body></html>");
       html = dw.Html_new(1002);
       dw.Box_pack_start(notebookbox7, html, 0, 100, TRUE, TRUE, 0);
       dw.Html_url(html, "http://dwindows.netlabs.org");
   } else {
       html = dw.Text_new("HTML widget not available.", 0);
       dw.Box_pack_start(notebookbox7, html, 0, 100, TRUE, TRUE, 0);
   }

   notebookbox8 = dw.Box_new(dw.VERT, 7);
   notebookpage8 := dw.Notebook_page_new(notebook, 1, FALSE);
   dw.Notebook_pack(notebook, notebookpage8, notebookbox8);
   dw.Notebook_page_set_text(notebook, notebookpage8, "scrollbox");
   scrollbox_add();

   /* Set the default field */
   dw.Window_default(mainwindow, copypastefield);

   dw.Signal_connect(mainwindow, dw.SIGNAL_DELETE, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
   /*
   * The following is a special case handler for the Mac and other platforms which contain
   * an application object which can be closed.  It function identically to a window delete/close
   * request except it applies to the entire application not an individual window. If it is not
   * handled or you allow the default handler to take place the entire application will close.
   * On platforms which do not have an application object this line will be ignored.
   */
   dw.Signal_connect(dw.DESKTOP, dw.SIGNAL_DELETE, unsafe.Pointer(&exit_callback_func), unsafe.Pointer(mainwindow));
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
