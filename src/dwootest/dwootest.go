package main

import (
   "dw"
   "fmt"
   "os"
   "bufio"
   "bytes"
   "io"
   "runtime"
   "go/build"
)

var FIXEDFONT = "10.monospace"

var mainwindow dw.HWND

// Page 1
var current_color dw.COLOR = dw.RGB(100, 100, 100)
var cursor_arrow bool = true
var timerid dw.HTIMER

// Page 2
var textbox1, textbox2 dw.HRENDER
var status1, status2 dw.HTEXT
var vscrollbar, hscrollbar dw.HSCROLLBAR
var imagexspin, imageyspin dw.HSPINBUTTON
var imagestretchcheck dw.HBUTTON
var text1pm, text2pm, image dw.HPIXMAP
var font_width = 8
var font_height = 12
var rows = 10
var width1 = 6
var cols = 80
var render_type = 0
var current_row = 0
var current_col = 0
var max_linewidth = 0

// Page 4
var mle_point = 0

// Page 5
var iteration = 0;

// Page 8
var MAX_WIDGETS = 20

// Page 9
var threadmle dw.HMLE
var startbutton dw.HBUTTON
var mutex dw.HMTX;
var workevent, controlevent dw.HEV
var finished = 0
var ready = 0
var mlepos = 0

// Miscellaneous
var fileicon, foldericon dw.HICN
var current_file string
var lines []string
var menu_enabled bool = true

var FOLDER_ICON_NAME string = "folder"
var FILE_ICON_NAME string = "file"
var SRCROOT string

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
    hscrollbar.SetRange(uint(max_linewidth), uint(cols));
    hscrollbar.SetPos(0);
    vscrollbar.SetRange(uint(len(lines)), uint(rows));
    vscrollbar.SetPos(0);
}

// Call back section
func exit_handler() int {
   if dw.MessageBox("dwtest", dw.MB_YESNO | dw.MB_QUESTION, "Are you sure you want to exit?") != 0 {
      dw.MainQuit();
   }
   return dw.TRUE;
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_file(row int, col int, nrows int, fheight int, hpma dw.HPIXMAP) {
    var hpm dw.HPIXMAP

    if hpma == dw.NOHPIXMAP {
        hpm = text2pm;
    } else {
        hpm = hpma;
    }

    if len(current_file) > 0 {
        var i int

        dw.ColorForegroundSet(dw.CLR_WHITE);
        if hpma == dw.NOHPIXMAP {
            text1pm.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, text1pm.GetWidth(), text1pm.GetHeight());
        }
        hpm.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, hpm.GetWidth(), hpm.GetHeight());

        for i = 0; (i < nrows) && (i+row < len(lines)); i++ {
            fileline := i + row - 1;
            y := i*fheight;
            dw.ColorBackgroundSet(dw.COLOR(1 + (fileline % 15)));
            dw.ColorForegroundSet(dw.COLOR(fileline % 16));
            if hpma == dw.NOHPIXMAP {
                text1pm.DrawText(0, y, fmt.Sprintf("%6.6d", i+row));
            }
            thisline := lines[i+row];
            if len(thisline) > col {
               hpm.DrawText(0, y, thisline[col:]);
            }
        }
        if hpma == dw.NOHPIXMAP {
            text_expose(textbox1, text1pm);
            text_expose(textbox2, text2pm);
        }
    }
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_shapes(direct int, hpma dw.HPIXMAP) {
    var hpm dw.HPIXMAP = dw.NOHPIXMAP
    var drawable dw.DRAWABLE = textbox2
    if hpma != dw.NOHPIXMAP {
        hpm = hpma;
    } else {
        hpm = text2pm;
    }
    if direct != dw.TRUE {
        drawable = hpm;
    }

    width := hpm.GetWidth();
    height := hpm.GetHeight();

    x := []int{ 20, 180, 180, 230, 180, 180, 20 };
    y := []int{ 50, 50, 20, 70, 120, 90, 90 };

    image_x := imagexspin.GetPos();
    image_y := imageyspin.GetPos();
    image_stretch := imagestretchcheck.Get();

    dw.ColorForegroundSet(dw.CLR_WHITE);
    drawable.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, width, height);
    dw.ColorForegroundSet(dw.CLR_DARKPINK);
    drawable.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 10, 10, width - 20, height - 20);
    dw.ColorForegroundSet(dw.CLR_GREEN);
    dw.ColorBackgroundSet(dw.CLR_DARKRED);
    drawable.DrawText(10, 10, "This should be aligned with the edges.");
    dw.ColorForegroundSet(dw.CLR_YELLOW);
    drawable.DrawLine(width - 10, 10, 10, height - 10);
    dw.ColorForegroundSet(dw.CLR_BLUE);
    drawable.DrawPolygon(dw.DRAW_FILL, x, y);
    dw.ColorForegroundSet(dw.CLR_BLACK);
    drawable.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 80, 80, 80, 40);
    dw.ColorForegroundSet(dw.CLR_CYAN);
    /* Bottom right corner */
    drawable.DrawArc(0, width - 30, height - 30, width - 10, height - 30, width - 30, height - 10);
    /* Top right corner */
    drawable.DrawArc(0, width - 30, 30, width - 30, 10, width - 10, 30);
    /* Bottom left corner */
    drawable.DrawArc(0, 30, height - 30, 30, height - 10, 10, height - 30);
    /* Full circle in the left top area */
    drawable.DrawArc(dw.DRAW_FULL, 120, 100, 80, 80, 160, 120);
    if image != dw.NOHPIXMAP {
        if image_stretch == dw.TRUE {
            drawable.BitBltStretchPixmap(10, 10, width - 20, height - 20, image, 0, 0, image.GetWidth(), image.GetHeight());
        } else {
            drawable.BitBltPixmap(image_x, image_y, image.GetWidth(), image.GetHeight(), image, 0, 0);
        }
    }

    /* If we aren't drawing direct do a bitblt */
    if direct == dw.FALSE && hpma == dw.NOHPIXMAP {
        text_expose(textbox2, text2pm);
    }
}

func update_render() {
    switch render_type {
        case 0:
            draw_shapes(dw.FALSE, dw.NOHPIXMAP);
        case 1:
            draw_shapes(dw.TRUE, dw.NOHPIXMAP);
        case 2:
            draw_file(current_row, current_col, rows, font_height, dw.NOHPIXMAP);
    }
}

func print_callback() {
   print := dw.PrintNew("DWTest Job");
   print.Connect(func(print dw.HPRINT, pixmap dw.HPIXMAP, page_num int) int {
                   pixmap.SetFont(FIXEDFONT);
                   if page_num == 0 {
                       draw_shapes(dw.FALSE, pixmap);
                   } else if page_num == 1 {
                       /* If we have a file to display... */
                       if len(current_file) > 0 {
                           /* Calculate new dimensions */
                           _, fheight := pixmap.GetTextExtents("(g");
                           nrows := int(pixmap.GetHeight() / fheight);

                           /* Do the actual drawing */
                           draw_file(0, 0, nrows, fheight, pixmap);
                       } else {
                           /* We don't have a file so center an error message on the page */
                           var text = "No file currently selected!";

                           /* Get the font size for this printer context... */
                           fwidth, fheight := pixmap.GetTextExtents(text);

                           posx := int(pixmap.GetWidth() - fwidth)/2;
                           posy := int(pixmap.GetHeight() - fheight)/2;

                           dw.ColorForegroundSet(dw.CLR_BLACK);
                           dw.ColorBackgroundSet(dw.CLR_WHITE);
                           pixmap.DrawText(posx, posy, text);
                       }
                   }
                   return dw.TRUE;
                }, 0, 2);
   print.Run(0);
}

func context_menu() {
    hwndMenu := dw.MenuNew(0);
    menuitem := hwndMenu.AppendItem("~Quit", dw.MENU_POPUP, 0, dw.TRUE, dw.FALSE, dw.NOMENU);

    menuitem.ConnectClicked(func(window dw.HMENUITEM) int { return exit_handler(); });
    hwndMenu.AppendItem(dw.MENU_SEPARATOR, dw.MENU_POPUP, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem = hwndMenu.AppendItem("~Show Window", dw.MENU_POPUP, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem.ConnectClicked(func(window dw.HMENUITEM) int {
                                mainwindow.Show();
                                mainwindow.Raise();
                                return dw.TRUE;
                            });
    px, py := dw.PointerGetPos();
    /* Use the toplevel window handle here.... because on the Mac..
     * using the control itself, when a different tab is active
     * the control is removed from the window and can no longer
     * handle the messages.
     */
    hwndMenu.Popup(mainwindow, px, py);
}

/* This gets called when a part of the graph needs to be repainted. */
func text_expose(hwnd dw.HRENDER, hpm dw.HPIXMAP) int {
    if render_type != 1 {
        width := hpm.GetWidth();
        height := hpm.GetHeight();

        hwnd.BitBltPixmap(0, 0, width, height, hpm, 0, 0);
        dw.Flush();
    } else {
        update_render();
    }
    return dw.TRUE;
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


func button_callback(combobox1 dw.HLISTBOX, combobox2 dw.HLISTBOX, spinbutton dw.HSPINBUTTON, cal dw.HCALENDAR) {
    idx := combobox1.Selected();
    buf1 := combobox1.GetText(idx);
    idx = combobox2.Selected();
    buf2 := combobox2.GetText(idx);
    y, m, d := cal.GetDate();
    spvalue := spinbutton.GetPos();
    message := fmt.Sprintf("spinbutton: %d\ncombobox1: \"%s\"\ncombobox2: \"%s\"\ncalendar: %d-%d-%d",
                  spvalue,
                  buf1, buf2,
                  y, m, d);
    dw.MessageBox( "Values", dw.MB_OK | dw.MB_INFORMATION, message);
}

// Create the menu
func menu_add(mainwindow dw.HWND) {
    mainmenubar := mainwindow.MenubarNew();
    /* add menus to the menubar */
    menu := dw.MenuNew(0);
    menuitem := menu.AppendItem("~Quit", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem.ConnectClicked(func(window dw.HMENUITEM) int { return exit_handler(); });
    /*
     * Add the "File" menu to the menubar...
     */
    mainmenubar.AppendItem("~File", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);

    changeable_menu := dw.MenuNew(0);
    checkable_menuitem := changeable_menu.AppendItem("~Checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    checkable_menuitem.ConnectClicked(func(window dw.HMENUITEM) int { dw.MessageBox("Menu Item Callback", dw.MB_OK | dw.MB_INFORMATION, "Checkable menu item selected"); return dw.FALSE });


    noncheckable_menuitem := changeable_menu.AppendItem("~Non-checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    noncheckable_menuitem.ConnectClicked(func(window dw.HMENUITEM) int { dw.MessageBox("Menu Item Callback", dw.MB_OK | dw.MB_INFORMATION, "Non-checkable menu item selected"); return dw.FALSE });
    changeable_menu.AppendItem("~Disabled menu Item", dw.MENU_AUTO, dw.MIS_DISABLED | dw.MIS_CHECKED, dw.TRUE, dw.TRUE, dw.NOMENU);
    /* seperator */
    changeable_menu.AppendItem(dw.MENU_SEPARATOR, dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem = changeable_menu.AppendItem("~Menu Items Disabled", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU);
    menuitem.ConnectClicked(func(window dw.HMENUITEM) int {
                                if menu_enabled {
                                    checkable_menuitem.SetStyle(dw.MIS_DISABLED, dw.MIS_DISABLED);
                                    noncheckable_menuitem.SetStyle(dw.MIS_DISABLED, dw.MIS_DISABLED);
                                    menu_enabled = false;
                                } else {
                                    checkable_menuitem.SetStyle(dw.MIS_DISABLED, dw.MIS_ENABLED);
                                    noncheckable_menuitem.SetStyle(dw.MIS_DISABLED, dw.MIS_ENABLED);
                                    menu_enabled = true;
                                }
                                return dw.FALSE;
                            });


    /*
     * Add the "Menu" menu to the menubar...
     */
    mainmenubar.AppendItem("~Menu", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, changeable_menu);

    menu = dw.MenuNew(0);
    menuitem = menu.AppendItem("~About", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU);
    menuitem.ConnectClicked(func(window dw.HMENUITEM) int {
                                env := dw.EnvironmentGet();
                                message := fmt.Sprintf("dwindows test\n\nOS: %s %s %s Version: %d.%d.%d.%d\n\ndwindows Version: %d.%d.%d",
                                                          env.OSName, env.BuildDate, env.BuildTime,
                                                          env.MajorVersion, env.MinorVersion, env.MajorBuild, env.MinorBuild,
                                                          env.DWMajorVersion, env.DWMinorVersion, env.DWSubVersion);
                                dw.MessageBox("About dwindows", dw.MB_OK | dw.MB_INFORMATION, message);
                                return dw.FALSE;
                            });

    /*
     * Add the "Help" menu to the menubar...
     */
    mainmenubar.AppendItem("~Help", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu);
}

// Create Page 1
func archive_add(notebookbox1 dw.HBOX) {
    lbbox := dw.BoxNew(dw.VERT, 10);

    notebookbox1.PackStart(lbbox, 150, 70, dw.TRUE, dw.TRUE, 0);

    /* Copy and Paste */
    browsebox := dw.BoxNew(dw.HORZ, 0);

    lbbox.PackStart(browsebox, 0, 0, dw.FALSE, dw.FALSE, 0);

    copypastefield := dw.EntryfieldNew("", 0);

    copypastefield.SetLimit(260);

    browsebox.PackStart(copypastefield, -1, -1, dw.TRUE, dw.FALSE, 4);

    copybutton := dw.ButtonNew("Copy", 0);

    browsebox.PackStart(copybutton, -1, -1, dw.FALSE, dw.FALSE, 0);

    pastebutton := dw.ButtonNew("Paste", 0);

    browsebox.PackStart(pastebutton, -1, -1, dw.FALSE, dw.FALSE, 0);

    /* Archive Name */
    stext := dw.TextNew("File to browse", 0);

    stext.SetStyle(dw.DT_VCENTER, dw.DT_VCENTER);

    lbbox.PackStart(stext, 130, 15, dw.TRUE, dw.TRUE, 2);

    browsebox = dw.BoxNew(dw.HORZ, 0);

    lbbox.PackStart(browsebox, 0, 0, dw.TRUE, dw.TRUE, 0);

    entryfield := dw.EntryfieldNew("", 100);

    entryfield.SetLimit(260);

    browsebox.PackStart(entryfield, 100, 15, dw.TRUE, dw.TRUE, 4);

    browsefilebutton := dw.ButtonNew("Browse File", 1001);

    browsebox.PackStart(browsefilebutton, 40, 15, dw.TRUE, dw.TRUE, 0);

    browsefolderbutton := dw.ButtonNew("Browse Folder", 1001);

    browsebox.PackStart(browsefolderbutton, 40, 15, dw.TRUE, dw.TRUE, 0);

    browsebox.SetColor(dw.CLR_PALEGRAY, dw.CLR_PALEGRAY);
    stext.SetColor(dw.CLR_BLACK, dw.CLR_PALEGRAY);

    /* Buttons */
    buttonbox := dw.BoxNew(dw.HORZ, 10);

    lbbox.PackStart(buttonbox, 0, 0, dw.TRUE, dw.TRUE, 0);

    cancelbutton := dw.ButtonNew("Exit", 1002);
    buttonbox.PackStart(cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2);

    cursortogglebutton := dw.ButtonNew("Set Cursor pointer - CLOCK", 1003);
    buttonbox.PackStart(cursortogglebutton, 130, 30, dw.TRUE, dw.TRUE, 2);

    okbutton := dw.ButtonNew("Turn Off Annoying Beep!", 1001);
    buttonbox.PackStart(okbutton, 130, 30, dw.TRUE, dw.TRUE, 2);

    cancelbutton.Unpack();
    buttonbox.PackStart(cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2);
    mainwindow.ClickDefault(cancelbutton);

    colorchoosebutton := dw.ButtonNew("Color Chooser Dialog", 1004);
    buttonbox.PackAtIndex(colorchoosebutton, 1, 130, 30, dw.TRUE, dw.TRUE, 2);

    /* Set some nice fonts and colors */
    lbbox.SetColor(dw.CLR_DARKCYAN, dw.CLR_PALEGRAY);
    buttonbox.SetColor(dw.CLR_DARKCYAN, dw.CLR_PALEGRAY);
    okbutton.SetColor(dw.CLR_PALEGRAY, dw.CLR_DARKCYAN);

    browsefilebutton.ConnectClicked(func(window dw.HBUTTON) int {
                                    tmp := dw.FileBrowse("Pick a file", "dwootest.go", "go", dw.FILE_OPEN);
                                    if len(tmp) > 0 {
                                        current_file = tmp;
                                        entryfield.SetText(current_file);
                                        read_file();
                                        current_col = 0;
                                        current_row = 0;
                                        update_render();
                                    }
                                    copypastefield.SetFocus();
                                    return dw.FALSE;
                                });

    browsefolderbutton.ConnectClicked(func(window dw.HBUTTON) int {
                                        tmp := dw.FileBrowse("Pick a folder", ".", "c", dw.DIRECTORY_OPEN);
                                        fmt.Printf("Folder picked: %s\n", tmp);
                                        return dw.FALSE;
                                    });
    copybutton.ConnectClicked(func(button dw.HBUTTON) int {
                               test := copypastefield.GetText();

                               if len(test) > 0 {
                                 dw.ClipboardSetText(test);
                               }
                               entryfield.SetFocus();
                               return dw.TRUE;
                            });
    pastebutton.ConnectClicked(func(button dw.HBUTTON) int {
                                test := dw.ClipboardGetText();

                                if len(test) > 0 {
                                    copypastefield.SetText(test);
                                }
                                return dw.TRUE;
                            });
    okbutton.ConnectClicked(func(window dw.HBUTTON) int { timerid.Disconnect(); return dw.TRUE; });
    cancelbutton.ConnectClicked(func(window dw.HBUTTON) int { return exit_handler(); });
    cursortogglebutton.ConnectClicked(func(window dw.HBUTTON) int {
                                        if cursor_arrow {
                                            cursortogglebutton.SetText("Set Cursor pointer - ARROW");
                                            mainwindow.SetPointer(dw.POINTER_CLOCK);
                                            cursor_arrow = false;
                                        } else {
                                            cursortogglebutton.SetText("Set Cursor pointer - CLOCK");
                                            mainwindow.SetPointer(dw.POINTER_DEFAULT);
                                            cursor_arrow = true;
                                        }
                                        return dw.FALSE;
                                    });
    colorchoosebutton.ConnectClicked(func(window dw.HBUTTON) int { current_color = dw.Color_choose(current_color); return dw.FALSE;});
    
    /* Set the default field */
    mainwindow.Default(copypastefield);
}

// Create Page 2
func text_add(notebookbox2 dw.HBOX) {
    depth := dw.ColorDepthGet();

    /* create a box to pack into the notebook page */
    pagebox := dw.BoxNew(dw.HORZ, 2);
    notebookbox2.PackStart(pagebox, 0, 0, dw.TRUE, dw.TRUE, 0);
    /* now a status area under this box */
    hbox := dw.BoxNew(dw.HORZ, 1);
    notebookbox2.PackStart(hbox, 100, 20, dw.TRUE, dw.FALSE, 1);
    status1 = dw.StatusTextNew("", 0);
    hbox.PackStart(status1, 100, -1, dw.TRUE, dw.FALSE, 1);
    status2 = dw.StatusTextNew("", 0);
    hbox.PackStart(status2, 100, -1, dw.TRUE, dw.FALSE, 1);
    /* a box with combobox and button */
    hbox = dw.BoxNew(dw.HORZ, 1);
    notebookbox2.PackStart(hbox, 100, 25, dw.TRUE, dw.FALSE, 1);
    rendcombo := dw.ComboboxNew("Shapes Double Buffered", 0);
    hbox.PackStart(rendcombo, 80, 25, dw.TRUE, dw.FALSE, 0);
    rendcombo.Append("Shapes Double Buffered");
    rendcombo.Append("Shapes Direct");
    rendcombo.Append("File Display");
    label := dw.TextNew("Image X:", 100);
    label.SetStyle(dw.DT_VCENTER | dw.DT_CENTER, dw.DT_VCENTER | dw.DT_CENTER);
    hbox.PackStart(label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imagexspin = dw.SpinButtonNew("20", 1021);
    hbox.PackStart(imagexspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    label = dw.TextNew("Y:", 100);
    label.SetStyle(dw.DT_VCENTER | dw.DT_CENTER, dw.DT_VCENTER | dw.DT_CENTER);
    hbox.PackStart(label, -1, 25, dw.FALSE, dw.FALSE, 0);
    imageyspin = dw.SpinButtonNew("20", 1021);
    hbox.PackStart(imageyspin, 25, 25, dw.TRUE, dw.FALSE, 0);
    imagexspin.SetLimits(2000, 0);
    imageyspin.SetLimits(2000, 0);
    imagexspin.SetPos(20);
    imageyspin.SetPos(20);
    imagestretchcheck = dw.CheckButtonNew("Stretch", 1021);
    hbox.PackStart(imagestretchcheck, -1, 25, dw.FALSE, dw.FALSE, 0);

    button1 := dw.ButtonNew("Refresh", 1223);
    hbox.PackStart(button1, 100, 25, dw.FALSE, dw.FALSE, 0);
    button2 := dw.ButtonNew("Print", 1224);
    hbox.PackStart(button2, 100, 25, dw.FALSE, dw.FALSE, 0);

    /* Pre-create the scrollbars so we can query their sizes */
    vscrollbar = dw.ScrollbarNew(dw.VERT, 50);
    hscrollbar = dw.ScrollbarNew(dw.HORZ, 50);
    vscrollbarwidth, _ := vscrollbar.GetPreferredSize();
    _, hscrollbarheight := hscrollbar.GetPreferredSize();

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
    textbox1 = dw.RenderNew(100);
    textbox1.SetFont(FIXEDFONT);
    font_width, font_height = textbox1.GetTextExtents("(g");
    font_width = font_width / 2;
    vscrollbox := dw.BoxNew(dw.VERT, 0);
    vscrollbox.PackStart(textbox1, font_width * width1, font_height * rows, dw.FALSE, dw.TRUE, 0);
    vscrollbox.PackStart(dw.NOHWND, font_width * (width1 + 1), hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    pagebox.PackStart(vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    /* pack empty space 1 character wide */
    pagebox.PackStart(dw.NOHWND, font_width, 0, dw.FALSE, dw.TRUE, 0);

    /* create box for filecontents and horz scrollbar */
    textboxA := dw.BoxNew(dw.VERT, 0);
    pagebox.PackStart(textboxA, 0, 0, dw.TRUE, dw.TRUE, 0);

    /* create render box for filecontents pixmap */
    textbox2 = dw.RenderNew(101);
    textboxA.PackStart(textbox2, 10, 10, dw.TRUE, dw.TRUE, 0);
    textbox2.SetFont(FIXEDFONT);
    /* create horizonal scrollbar */
    textboxA.PackStart(hscrollbar, -1, -1, dw.TRUE, dw.FALSE, 0);

    /* create vertical scrollbar */
    vscrollbox = dw.BoxNew(dw.VERT, 0);
    vscrollbox.PackStart(vscrollbar, -1, -1, dw.FALSE, dw.TRUE, 0);
    /* Pack an area of empty space 14x14 pixels */
    vscrollbox.PackStart(dw.NOHWND, vscrollbarwidth, hscrollbarheight, dw.FALSE, dw.FALSE, 0);
    pagebox.PackStart(vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0);

    text1pm = dw.PixmapNew(textbox1, uint(font_width * width1), uint(font_height * rows), depth);
    text2pm = dw.PixmapNew(textbox2, uint(font_width * cols), uint(font_height * rows), depth);
    image = dw.PixmapNewFromFile(textbox2, "test");
    if image == dw.NOHPIXMAP && len(SRCROOT) > 0 {
        image = dw.PixmapNewFromFile(textbox2, fmt.Sprintf("%s/test", SRCROOT));
    }
    if image != dw.NOHPIXMAP {
        image.SetTransparentColor(dw.CLR_WHITE);
    }

    dw.MessageBox("DWTest", dw.MB_OK | dw.MB_INFORMATION, fmt.Sprintf("Width: %d Height: %d\n", font_width, font_height));
    text1pm.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, font_width * width1, font_height * rows);
    text2pm.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, font_width * cols, font_height * rows);
    textbox1.ConnectButtonPress(func(window dw.HRENDER, x int, y int, buttonmask int) int { context_menu(); return dw.TRUE; });
    textbox1.ConnectExpose(func(hwnd dw.HRENDER, x int, y int, width int, height int) int { return text_expose(hwnd, text1pm); });
    textbox2.ConnectExpose(func(hwnd dw.HRENDER, x int, y int, width int, height int) int { return text_expose(hwnd, text2pm); });
    textbox2.ConnectConfigure(func(hwnd dw.HRENDER, width int, height int) int {
                                old1 := text1pm;
                                old2 := text2pm;
                                depth := dw.ColorDepthGet();

                                rows = height / font_height;
                                cols = width / font_width;

                                /* Create new pixmaps with the current sizes */
                                text1pm = dw.PixmapNew(textbox1, uint(font_width*(width1)), uint(height), depth);
                                text2pm = dw.PixmapNew(textbox2, uint(width), uint(height), depth);

                                /* Make sure the side area is cleared */
                                dw.ColorForegroundSet(dw.CLR_WHITE);
                                text1pm.DrawRect(dw.DRAW_FILL | dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(text1pm), dw.Pixmap_height(text1pm));

                               /* Destroy the old pixmaps */
                                old1.Destroy();
                                old2.Destroy();

                                /* Update scrollbar ranges with new values */
                                hscrollbar.SetRange(uint(max_linewidth), uint(cols));
                                vscrollbar.SetRange(uint(len(lines)), uint(rows));

                                /* Redraw the window */
                                update_render();
                                return dw.TRUE;
                            });
    textbox2.ConnectMotion(func(window dw.HRENDER, x int, y int, buttonmask int) int { status2.SetText(fmt.Sprintf("motion_notify: %dx%d", x, y)); return dw.FALSE; });
    textbox2.ConnectButtonPress(func(window dw.HRENDER, x int, y int, buttonmask int) int { status2.SetText(fmt.Sprintf("button_press: %dx%d", x, y)); return dw.FALSE; });
    hscrollbar.ConnectValueChanged(func(hwnd dw.HSCROLLBAR, value int) int {
                                        current_col = value;
                                        status1.SetText(fmt.Sprintf("Row:%d Col:%d Lines:%d Cols:%d", current_row, current_col, len(lines), max_linewidth));
                                        update_render();
                                        return dw.FALSE;
                                    });
    vscrollbar.ConnectValueChanged(func(hwnd dw.HSCROLLBAR, value int) int {
                                        current_row = value;
                                        status1.SetText(fmt.Sprintf("Row:%d Col:%d Lines:%d Cols:%d", current_row, current_col, len(lines), max_linewidth));
                                        update_render();
                                        return dw.FALSE;
                                    });    
    imagestretchcheck.ConnectClicked(func(window dw.HBUTTON) int { update_render(); return dw.FALSE; });
    button1.ConnectClicked(func(window dw.HBUTTON) int { update_render(); return dw.FALSE; });
    button2.ConnectClicked(func(window dw.HBUTTON) int { print_callback(); return dw.FALSE; });
    rendcombo.ConnectListSelect(func(window dw.HLISTBOX, index int) int {
                                    if index != render_type {
                                        if index == 2 {
                                            hscrollbar.SetRange(uint(max_linewidth), uint(cols));
                                            hscrollbar.SetPos(0);
                                            vscrollbar.SetRange(uint(len(lines)), uint(rows));
                                            vscrollbar.SetPos(0);
                                            current_col = 0;
                                            current_row = 0;
                                        } else {
                                            hscrollbar.SetRange(0, 0);
                                            hscrollbar.SetPos(0);
                                            vscrollbar.SetRange(0, 0);
                                            vscrollbar.SetPos(0);
                                        }
                                        render_type = index;
                                        update_render();
                                    }
                                    return dw.FALSE;
                                });
    mainwindow.ConnectKeyPress(func(window dw.HWND, ch uint8, vk int, state int, utf8 string) int {
                                var message string

                                if ch != 0 {
                                    message = fmt.Sprintf("Key: %c(%d) Modifiers: %s(%d) utf8 %s", ch, ch, resolve_keymodifiers(state), state,  utf8);
                                } else {
                                    message = fmt.Sprintf("Key: %s(%d) Modifiers: %s(%d) utf8 %s", resolve_keyname(vk), vk, resolve_keymodifiers(state), state, utf8);
                                }
                                status1.SetText(message);
                                return dw.FALSE;
                            });


    dw.TaskbarInsert(textbox1, fileicon, "DWTest");
}

// Page 3
func tree_add(notebookbox3 dw.HBOX) {
    /* create a box to pack into the notebook page */
    listbox := dw.ListboxNew(1024, dw.TRUE);
    notebookbox3.PackStart(listbox, 500, 200, dw.TRUE, dw.TRUE, 0);
    listbox.Append("Test 1");
    listbox.Append("Test 2");
    listbox.Append("Test 3");
    listbox.Append("Test 4");
    listbox.Append("Test 5");

    /* now a tree area under this box */
    tree := dw.TreeNew(101);
    notebookbox3.PackStart(tree, 500, 200, dw.TRUE, dw.TRUE, 1);

    /* and a status area to see whats going on */
    tree_status := dw.StatusTextNew("", 0);
    notebookbox3.PackStart(tree_status, 100, -1, dw.TRUE, dw.FALSE, 1);

    /* set up our signal trappers... */
    tree.ConnectItemContext(func(window dw.HTREE, text string, x int, y int, itemdata dw.POINTER) int {
                                tree_status.SetText(fmt.Sprintf("DW_SIGNAL_ITEM_CONTEXT: Window: %x Text: %s x: %d y: %d Itemdata: %x", 
                                        dw.HANDLE_TO_UINTPTR(window), text, x, y, uintptr(itemdata)));
                                return dw.FALSE;
                            });
    tree.ConnectItemSelect(func(window dw.HTREE, item dw.HTREEITEM, text string, itemdata dw.POINTER) int {
                                tree_status.SetText(fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", 
                                        dw.HANDLE_TO_UINTPTR(window), dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata)));
                                return dw.FALSE;
                            });

    t1 := tree.Insert("tree folder 1", foldericon, dw.NOHTREEITEM, dw.POINTER(uintptr(1)));
    t2 := tree.Insert("tree folder 2", foldericon, dw.NOHTREEITEM, dw.POINTER(uintptr(2)));
    tree.Insert("tree file 1", fileicon, t1, dw.POINTER(uintptr(3)));
    tree.Insert("tree file 2", fileicon, t1, dw.POINTER(uintptr(4)));
    tree.Insert("tree file 3", fileicon, t2, dw.POINTER(uintptr(5)));
    tree.Insert("tree file 4", fileicon, t2, dw.POINTER(uintptr(6)));
    t1.Change("tree folder 1", foldericon);
    t2.Change("tree folder 2", foldericon);
    t2.SetData(dw.POINTER(uintptr(100)));
    fmt.Printf("t1 title \"%s\" data %d t2 data %d\n", t1.GetTitle(), uintptr(t1.GetData()), uintptr(t2.GetData())); 
}

// Page 4
func container_add(notebookbox4 dw.HBOX) {
    var z int
    titles := []string{ "Type", "Size", "Time", "Date" };
    flags := []uint{   dw.CFA_BITMAPORICON | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_ULONG | dw.CFA_RIGHT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_TIME | dw.CFA_CENTER | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
                         dw.CFA_DATE | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR };


    /* create a box to pack into the notebook page */
    containerbox := dw.BoxNew(dw.HORZ, 2);
    notebookbox4.PackStart(containerbox, 500, 200, dw.TRUE, dw.TRUE, 0);

    /* now a container area under this box */
    container := dw.ContainerNew(100, dw.TRUE);
    notebookbox4.PackStart(container, 500, 200, dw.TRUE, dw.FALSE, 1);

    /* and a status area to see whats going on */
    container_status := dw.StatusTextNew("", 0);
    notebookbox4.PackStart(container_status, 100, -1, dw.TRUE, dw.FALSE, 1);

    container.SetColumnTitle("Test");
    container.FileSystemSetup(flags, titles);
    container.SetStripe(dw.CLR_DEFAULT, dw.CLR_DEFAULT);
    containerinfo := container.Alloc(3);

    for z=0; z<3; z++ {
        var thisicon dw.HICN = fileicon; 
        
        if z == 0 {
             thisicon = foldericon;
        } 
        fmt.Printf("Initial: container: %x containerinfo: %x icon: %x\n", uintptr(dw.HANDLE_TO_POINTER(container)),
                  dw.HANDLE_TO_UINTPTR(containerinfo), uintptr(dw.POINTER(thisicon)));
        containerinfo.SetFile(z, fmt.Sprintf("Filename %d", z+1), thisicon);
        containerinfo.SetItemIcon(0, z, thisicon);
        containerinfo.SetItemULong(1, z, uint(z*100));
        containerinfo.SetItemTime(2, z, z+10, z+10, z+10);
        containerinfo.SetItemDate(3, z, z+10, z+10, z+2000);
        containerinfo.SetRowTitle(z, fmt.Sprintf("We can now allocate from the stack: Item: %d", z));
        containerinfo.SetRowData(z, dw.POINTER(uintptr(z)));
    }
    containerinfo.Insert();

    containerinfo = container.Alloc(1);
    containerinfo.SetFile(0, "Yikes", foldericon);
    containerinfo.SetItemIcon(0, 0, foldericon);
    containerinfo.SetItemULong(1, 0, 324);
    containerinfo.SetItemTime(2, 0, z+10, z+10, z+10);
    containerinfo.SetItemDate(3, 0, z+10, z+10, z+2000);
    containerinfo.SetRowTitle(0, "Extra");

    containerinfo.Insert();
    container.Optimize();

    container_mle := dw.MLENew(111);
    containerbox.PackStart(container_mle, 500, 200, dw.TRUE, dw.TRUE, 0);

    mle_point = container_mle.Import("", -1);
    mle_point = container_mle.Import(fmt.Sprintf("[%d]", mle_point), mle_point);
    mle_point = container_mle.Import(fmt.Sprintf("[%d]abczxydefijkl", mle_point), mle_point);
    dw.Mle_delete(container_mle, 9, 3);
    mle_point = container_mle.Import("gh", 12);
    newpoint, _ := container_mle.GetSize();
    mle_point = newpoint;
    mle_point = container_mle.Import(fmt.Sprintf("[%d]\r\n\r\n", mle_point), mle_point);
    container_mle.SetCursor(mle_point);
    /* connect our event trappers... */
    container.ConnectItemEnter(func(window dw.HCONTAINER, text string, itemdata dw.POINTER) int {
        container_status.SetText(fmt.Sprintf("DW_SIGNAL_ITEM_ENTER: Window: %x Text: %s Itemdata: %x", dw.HANDLE_TO_UINTPTR(window), text, uintptr(itemdata)));
        return dw.FALSE;
    });
    container.ConnectItemContext(func(window dw.HCONTAINER, text string, x int, y int, itemdata dw.POINTER) int {
                                    container_status.SetText(fmt.Sprintf("DW_SIGNAL_ITEM_CONTEXT: Window: %x Text: %s x: %d y: %d Itemdata: %x", 
                                            dw.HANDLE_TO_UINTPTR(window), text, x, y, uintptr(itemdata)));
                                    return dw.FALSE;
                                });
    
    container.ConnectItemSelect(func(window dw.HCONTAINER, item dw.HTREEITEM, text string, itemdata dw.POINTER)  int {
                                    message := fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", 
                                            dw.HANDLE_TO_UINTPTR(window), dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata));
                                    container_status.SetText(message);
                                    message = fmt.Sprintf("\r\nDW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x\r\n", 
                                             dw.HANDLE_TO_UINTPTR(window), dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata));
                                    mle_point = container_mle.Import(message, mle_point);
                                    str := container.QueryStart(dw.CRA_SELECTED);
                                    for len(str) > 0 {
                                        mle_point = container_mle.Import(fmt.Sprintf("Selected: %s\r\n", str), mle_point);
                                        str = container.QueryNext(dw.CRA_SELECTED);
                                    }
                                    /* Make the last inserted point the cursor location */
                                    container_mle.SetCursor(mle_point);
                                    /* set the details of item 0 to new data */
                                    container.ChangeFile(0, "new data", fileicon);
                                    container.ChangeItemULong(1, 0, 999);
                                    return dw.FALSE;
                                });

    container.ConnectColumnClick(func(window dw.HCONTAINER, column_num int) int {
                                    var stype = "Unknown";

                                    if column_num == 0 {
                                        stype = "Filename";
                                    } else {
                                        column_type := window.GetColumnType(column_num-1);
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
                                    container_status.SetText(fmt.Sprintf("DW_SIGNAL_COLUMN_CLICK: Window: %x Column: %d Type: %s Itemdata: %x", 
                                            dw.HANDLE_TO_UINTPTR(window), column_num, stype));
                                    return dw.FALSE;
                                });
}

// Page 5
func buttons_add(notebookbox5 dw.HBOX) {
    var i int;
    
    /* create a box to pack into the notebook page */
    buttonsbox := dw.BoxNew(dw.VERT, 2);
    notebookbox5.PackStart(buttonsbox, 25, 200, dw.TRUE, dw.TRUE, 0);
    buttonsbox.SetColor(dw.CLR_RED, dw.CLR_RED);

    calbox := dw.BoxNew(dw.HORZ, 0);
    notebookbox5.PackStart(calbox, 500, 200, dw.TRUE, dw.TRUE, 1);
    cal := dw.CalendarNew(100);
    calbox.PackStart(cal, 180, 120, dw.TRUE, dw.TRUE, 0);
    /*
     cal.SetDate(2001, 1, 1);
     */
    /*
     * Create our file toolbar boxes...
     */
    buttonboxperm := dw.BoxNew(dw.VERT, 0);
    buttonsbox.PackStart(buttonboxperm, 25, 0, dw.FALSE, dw.TRUE, 2);
    buttonboxperm.SetColor(dw.CLR_WHITE, dw.CLR_WHITE);
    abutton1 := dw.BitmapButtonNewFromFile("Top Button", 0, fmt.Sprintf("%s/%s", SRCROOT, FILE_ICON_NAME));
    buttonboxperm.PackStart(abutton1, 100, 30, dw.FALSE, dw.FALSE, 0);
    buttonboxperm.PackStart(dw.NOHWND, 25, 5, dw.FALSE, dw.FALSE, 0);
    abutton2 := dw.BitmapButtonNewFromFile("Bottom", 0, fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME));
    buttonsbox.PackStart(abutton2, 25, 25, dw.FALSE, dw.FALSE, 0);
    abutton2.SetBitmap(0, FILE_ICON_NAME);

    /* Pre-create the percent */
    percent := dw.PercentNew(0);
    
    create_button(buttonboxperm, buttonsbox, percent);
    
    /* make a combobox */
    combox := dw.BoxNew(dw.VERT, 2);
    notebookbox5.PackStart(combox, 25, 200, dw.TRUE, dw.FALSE, 0);
    combobox1 := dw.ComboboxNew("fred", 0 ); /* no point in specifying an initial value */
    combobox1.Append("fred" );
    combox.PackStart(combobox1, -1, -1, dw.TRUE, dw.FALSE, 0);


    combobox2 := dw.ComboboxNew("joe", 0); /* no point in specifying an initial value */
    combox.PackStart(combobox2, -1, -1, dw.TRUE, dw.FALSE, 0);
    /* add LOTS of items */
    fmt.Printf("before appending 500 items to combobox using dw_listbox_list_append()\n");
    text := make([]string, 500);
    for  i = 0; i < 500; i++ {
        text[i] = fmt.Sprintf("item %d", i);
    }
    combobox2.AppendList(text);
    fmt.Printf("after appending 500 items to combobox\n");
    /* now insert a couple of items */
    combobox2.Insert("inserted item 2", 2 );
    combobox2.Insert("inserted item 5", 5 );
    /* make a spinbutton */
    spinbutton := dw.SpinButtonNew("", 0); /* no point in specifying text */
    combox.PackStart(spinbutton, -1, -1, dw.TRUE, dw.FALSE, 0);
    spinbutton.SetLimits(100, 1);
    spinbutton.SetPos(30);

    /* make a slider */
    slider := dw.SliderNew(dw.FALSE, 11, 0); /* no point in specifying text */
    combox.PackStart(slider, -1, -1, dw.TRUE, dw.FALSE, 0);

    /* Pack the percent */
    combox.PackStart(percent, -1, -1, dw.TRUE, dw.FALSE, 0);
    
    /* Connect the handlers */
    abutton1.ConnectClicked(func(window dw.HBUTTON) int { button_callback(combobox1, combobox2, spinbutton, cal); return dw.TRUE; });
    abutton2.ConnectClicked(func(window dw.HBUTTON) int { button_callback(combobox1, combobox2, spinbutton, cal); return dw.TRUE; });
    combobox1.ConnectListSelect(func(window dw.HLISTBOX, index int) int {
                                    fmt.Printf("got combobox_select_event for index: %d, iteration: %d\n", index, iteration);
                                    iteration++;
                                    return dw.FALSE;
                                });
    combobox2.ConnectListSelect(func(window dw.HLISTBOX, index int) int {
                                    fmt.Printf("got combobox_select_event for index: %d, iteration: %d\n", index, iteration);
                                    iteration++;
                                    return dw.FALSE;
                                });
    spinbutton.ConnectValueChanged(func(hwnd dw.HSPINBUTTON, value int) int { dw.MessageBox("DWTest", dw.MB_OK, fmt.Sprintf("New value from spinbutton: %d\n", value)); return dw.FALSE; });
    slider.ConnectValueChanged(func(hwnd dw.HSLIDER, value int) int { percent.SetPos(uint(value * 10)); return dw.FALSE; });
}

func create_button(buttonboxperm, buttonsbox dw.HBOX, percent dw.HPERCENT) {
    filetoolbarbox := dw.BoxNew(dw.VERT, 0);
    buttonboxperm.PackStart(filetoolbarbox, 0, 0, dw.TRUE, dw.TRUE, 0);

    abutton1 := dw.BitmapButtonNewFromFile("Empty image. Should be under Top button", 0, "junk");
    filetoolbarbox.PackStart(abutton1, 25, 25, dw.FALSE, dw.FALSE, 0);
    abutton1.ConnectClicked(func(window dw.HBUTTON) int { buttonsbox.SetColor(dw.CLR_RED, dw.CLR_RED); return dw.FALSE; });

    filetoolbarbox.PackStart(dw.NOHWND, 25, 5, dw.FALSE, dw.FALSE, 0);

    abutton1 = dw.BitmapButtonNewFromFile("A borderless bitmapbitton", 0, fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME));
    filetoolbarbox.PackStart(abutton1, 25, 25, dw.FALSE, dw.FALSE, 0);
    abutton1.ConnectClicked(func(window dw.HBUTTON) int { buttonsbox.SetColor(dw.CLR_YELLOW, dw.CLR_YELLOW); return dw.FALSE; });

    filetoolbarbox.PackStart(dw.NOHWND, 25, 5, dw.FALSE, dw.FALSE, 0);
    abutton1.SetStyle(dw.BS_NOBORDER, dw.BS_NOBORDER);

    //abutton1 = dw.Bitmapbutton_new_from_data("A button from data", 0, folder_ico, 1718 );
    abutton1 = dw.BitmapButtonNewFromFile("A button from data", 0, "junk");
    filetoolbarbox.PackStart(abutton1, 25, 25, dw.FALSE, dw.FALSE, 0);
    abutton1.ConnectClicked(func(window dw.HBUTTON) int { percent.SetPos(dw.PERCENT_INDETERMINATE); return dw.FALSE; });

    filetoolbarbox.PackStart(dw.NOHWND, 25, 5, dw.FALSE, dw.FALSE, 0);
}

// Page 8
func scrollbox_add(notebookbox8 dw.HBOX) {
   var i int;

    /* create a box to pack into the notebook page */
    scrollbox := dw.ScrollBoxNew(dw.VERT, 0);
    notebookbox8.PackStart(scrollbox, 0, 0, dw.TRUE, dw.TRUE, 1);

    abutton1 := dw.ButtonNew("Show Adjustments", 0);
    scrollbox.PackStart(abutton1, -1, 30, dw.FALSE, dw.FALSE, 0 );
    abutton1.ConnectClicked(func(window dw.HBUTTON) int { _, pos := scrollbox.GetPos(); _, rng := scrollbox.GetRange(); fmt.Printf("Pos %d Range %d\n", pos, rng); return dw.FALSE; });

    for i = 0; i < MAX_WIDGETS; i++ {
        tmpbox := dw.BoxNew(dw.HORZ, 0);
        scrollbox.PackStart(tmpbox, 0, 24, dw.TRUE, dw.FALSE, 2);
        label := dw.TextNew(fmt.Sprintf("Label %d", i), 0 );
        tmpbox.PackStart(label, 0, 20, dw.TRUE, dw.FALSE, 0);
        item := dw.EntryfieldNew(fmt.Sprintf("Entry %d", i), uint(i));
        tmpbox.PackStart(item, 0, 20, dw.TRUE, dw.FALSE, 0);
    }
}

// Page 9
func update_mle(text string, lock int) {
    /* Protect pos from being changed by different threads */
    if(lock != 0) {
        mutex.Lock();
    }
    mlepos = threadmle.Import(text, mlepos);
    threadmle.SetCursor(mlepos);
    if(lock != 0) {
        mutex.Unlock();
    }
}

func thread_add(notebookbox9 dw.HBOX) {
    /* create a box to pack into the notebook page */
    tmpbox := dw.BoxNew(dw.VERT, 0);
    notebookbox9.PackStart(tmpbox, 0, 0, dw.TRUE, dw.TRUE, 1);

    startbutton = dw.ButtonNew("Start Threads", 0);
    tmpbox.PackStart(startbutton, -1, 30, dw.FALSE, dw.FALSE, 0);
    /* Create the base threading components */
    threadmle = dw.MLENew(0);
    tmpbox.PackStart(threadmle, 1, 1, dw.TRUE, dw.TRUE, 0);
    mutex = dw.MutexNew();
    workevent = dw.EventNew();
    /* Connect signal handlers */
    startbutton.ConnectClicked(func(window dw.HBUTTON) int {
        startbutton.Disable();
        mutex.Lock();
        controlevent = dw.EventNew();
        workevent.Reset();
        finished = dw.FALSE;
        ready = 0;
        update_mle("Starting thread 1\r\n", dw.FALSE);
        go run_thread(1);
        update_mle("Starting thread 2\r\n", dw.FALSE);
        go run_thread(2);
        update_mle("Starting thread 3\r\n", dw.FALSE);
        go run_thread(3);
        update_mle("Starting thread 4\r\n", dw.FALSE);
        go run_thread(4);
        update_mle("Starting control thread\r\n", dw.FALSE);
        go control_thread();
        mutex.Unlock();
        return dw.FALSE;
    });

}

func run_thread(threadnum int) {
    dw.InitThread();
    update_mle(fmt.Sprintf("Thread %d started.\r\n", threadnum), dw.TRUE);

    /* Increment the ready count while protected by mutex */
    mutex.Lock();
    ready++;
    /* If all 4 threads have incrememted the ready count...
     * Post the control event semaphore so things will get started.
     */
    if(ready == 4) {
        controlevent.Post();
    }
    mutex.Unlock();

    for finished == 0 {
        result := workevent.Wait(2000);

        if(result == dw.ERROR_TIMEOUT) {
            update_mle(fmt.Sprintf("Thread %d timeout waiting for event.\r\n", threadnum), dw.TRUE);
        } else if(result == dw.ERROR_NONE) {
            update_mle(fmt.Sprintf("Thread %d doing some work.\r\n", threadnum), dw.TRUE);
            /* Pretend to do some work */
            dw.MainSleep(1000 * threadnum);

            /* Increment the ready count while protected by mutex */
            mutex.Lock();
            ready++;
            buf := fmt.Sprintf("Thread %d work done. ready=%d", threadnum, ready);
            /* If all 4 threads have incrememted the ready count...
            * Post the control event semaphore so things will get started.
            */
            if(ready == 4) {
                controlevent.Post();
                buf = fmt.Sprintf("%s%s", buf, " Control posted.");
            }
            mutex.Unlock();
            update_mle(fmt.Sprintf("%s\r\n", buf), dw.TRUE);
        } else {
            update_mle(fmt.Sprintf("Thread %d error %d.\r\n", threadnum), dw.TRUE);
            dw.MainSleep(10000);
        }
    }
    update_mle(fmt.Sprintf("Thread %d finished.\r\n", threadnum), dw.TRUE);
    dw.DeinitThread();
}

func control_thread() {
    dw.InitThread();
    
    inprogress := 5;

    for inprogress != 0 {
        result := controlevent.Wait(2000);

        if(result == dw.ERROR_TIMEOUT) {
            update_mle("Control thread timeout waiting for event.\r\n", dw.TRUE);
        } else if(result == dw.ERROR_NONE) {
            /* Reset the control event */
            controlevent.Reset();
            ready = 0;
            update_mle(fmt.Sprintf("Control thread starting worker threads. Inprogress=%d\r\n", inprogress), dw.TRUE);
            /* Start the work threads */
            workevent.Post();
            dw.MainSleep(100);
            /* Reset the work event */
            workevent.Reset();
            inprogress--;
        } else {
            update_mle(fmt.Sprintf("Control thread error %d.\r\n", result), dw.TRUE);
            dw.MainSleep(10000);
        }
    }
    /* Tell the other threads we are done */
    finished = dw.TRUE;
    workevent.Post();
    /* Close the control event */
    controlevent.Close();
    update_mle("Control thread finished.\r\n", dw.TRUE);
    startbutton.Enable();
    dw.DeinitThread();
}

func main() {
   /* Pick an approriate font for our platform */
   if runtime.GOOS == "windows" {
      FIXEDFONT = "10.Lucida Console";
   } else if runtime.GOOS == "darwin" {
      FIXEDFONT = "9.Monaco";
   }
   
   /* Locate the source root of the package */
   pkg, err := build.Import("dwtest", "", build.FindOnly);
   if err == nil && len(pkg.SrcRoot) > 0 {
      SRCROOT = fmt.Sprintf("%s/dwtest", pkg.SrcRoot);
   }
   
   /* Initialize the Dynamic Windows engine */
   dw.Init(dw.TRUE);

   /* Create our window */
   mainwindow := dw.WindowNew(dw.DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", dw.FCF_SYSMENU | dw.FCF_TITLEBAR | dw.FCF_TASKLIST | dw.FCF_DLGBORDER | dw.FCF_SIZEBORDER | dw.FCF_MINMAX);

   menu_add(mainwindow);

   notebookbox := dw.BoxNew(dw.VERT, 5);
   mainwindow.PackStart(notebookbox, 0, 0, dw.TRUE, dw.TRUE, 0);

   foldericon = dw.IconLoadFromFile(FOLDER_ICON_NAME);
   if foldericon == dw.NOHICN && len(SRCROOT) > 0 {
      foldericon = dw.IconLoadFromFile(fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME));
   }
   fileicon = dw.IconLoadFromFile(FILE_ICON_NAME);
   if fileicon == dw.NOHICN && len(SRCROOT) > 0 {
      fileicon = dw.IconLoadFromFile(fmt.Sprintf("%s/%s", SRCROOT, FILE_ICON_NAME));
   }
   notebook := dw.NotebookNew(1, dw.TRUE);
   notebookbox.PackStart(notebook, 100, 100, dw.TRUE, dw.TRUE, 0);
   notebook.ConnectSwitchPage(func(window dw.HNOTEBOOK, page_num dw.HNOTEPAGE) int {
                                    fmt.Printf("DW_SIGNAL_SWITCH_PAGE: PageNum: %d\n", dw.HNOTEPAGE_TO_UINT(page_num));
                                    return dw.FALSE;
                                });

   notebookbox1 := dw.BoxNew(dw.VERT, 5);
   notebookpage1 := notebook.PageNew(0, dw.TRUE);
   notebookpage1.Pack(notebookbox1);
   notebookpage1.SetText("buttons and entry");
   archive_add(notebookbox1);

   notebookbox2 := dw.BoxNew(dw.VERT, 5);
   notebookpage2 := notebook.PageNew(1, dw.FALSE);
   notebookpage2.Pack(notebookbox2);
   notebookpage2.SetText("render");
   text_add(notebookbox2);

   notebookbox3 := dw.BoxNew(dw.VERT, 5);
   notebookpage3 := notebook.PageNew(1, dw.FALSE);
   notebookpage3.Pack(notebookbox3);
   notebookpage3.SetText("tree");
   tree_add(notebookbox3);
   
   notebookbox4 := dw.BoxNew(dw.VERT, 5);
   notebookpage4 := notebook.PageNew(1, dw.FALSE);
   notebookpage4.Pack(notebookbox4);
   notebookpage4.SetText("container");
   container_add(notebookbox4);

   notebookbox5 := dw.BoxNew(dw.VERT, 5);
   notebookpage5 := notebook.PageNew(1, dw.FALSE);
   notebookpage5.Pack(notebookbox5);
   notebookpage5.SetText("buttons");
   buttons_add(notebookbox5);

/* DEPRECATED
   notebookbox6 := dw.BoxNew(dw.VERT, 5);
   notebookpage6 := notebook.PageNew(1, dw.FALSE );
   notebookpage6.Pack(notebookbox6);
   notebookpage6.SetText("mdi");
   mdi_add(notebookbox6);
*/

   notebookbox7 := dw.BoxNew(dw.VERT, 6);
   notebookpage7 := notebook.PageNew(1, dw.FALSE);
   notebookpage7.Pack(notebookbox7);
   notebookpage7.SetText("html");
   
   rawhtml := dw.HtmlNew(1001);
   if rawhtml.GetHandle() != nil {
       notebookbox7.PackStart(rawhtml, 0, 100, dw.TRUE, dw.FALSE, 0);
       rawhtml.Raw("<html><body><center><h1>dwtest</h1></center></body></html>");
       html := dw.HtmlNew(1002);
       notebookbox7.PackStart(html, 0, 100, dw.TRUE, dw.TRUE, 0);
       html.URL("http://dwindows.netlabs.org");
   } else {
       label := dw.Text_new("HTML widget not available.", 0);
       notebookbox7.PackStart(label, 0, 100, dw.TRUE, dw.TRUE, 0);
   }

   notebookbox8 := dw.BoxNew(dw.VERT, 7);
   notebookpage8 := notebook.PageNew(1, dw.FALSE);
   notebookpage8.Pack(notebookbox8);
   notebookpage8.SetText("scrollbox");
   scrollbox_add(notebookbox8);

   notebookbox9 := dw.BoxNew(dw.VERT, 8);
   notebookpage9 := notebook.PageNew(1, dw.FALSE);
   notebookpage9.Pack(notebookbox9);
   notebookpage9.SetText("thread/event");
   thread_add(notebookbox9);

   mainwindow.ConnectDelete(func(window dw.HWND) int { return exit_handler(); });
   /*
   * The following is a special case handler for the Mac and other platforms which contain
   * an application object which can be closed.  It function identically to a window delete/close
   * request except it applies to the entire application not an individual window. If it is not
   * handled or you allow the default handler to take place the entire application will close.
   * On platforms which do not have an application object this line will be ignored.
   */
   dw.DESKTOP.ConnectDelete(func(window dw.HWND) int { return exit_handler(); });
   timerid = dw.TimerNew();
   /* Return dw.TRUE so we get called again */
   timerid.Connect(func() int { dw.Beep(200, 200); return dw.TRUE; }, 2000);

   mainwindow.SetSize(640, 550);
   mainwindow.Show();

  /* Now that the window is created and shown...
   * run the main loop until we get dw_main_quit()
   */
   dw.Main();

   /* Now that the loop is done we can cleanup */
   dw.TaskbarDelete(textbox1, fileicon);
   mainwindow.Destroy();
   
   fmt.Printf("dwtest exiting...\n");
   /* Call dw.Shutdown() to shutdown the Dynamic Windows engine */
   dw.Shutdown();
}
