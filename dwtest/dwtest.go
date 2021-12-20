package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/build"
	"hg.code.sf.net/p/godwindows/code.hg/dw"
	"io"
	"os"
	"runtime"
)

// Global variables
const (
	FALSE int = iota
	TRUE
)

var FIXEDFONT = "10.monospace"

var mainwindow dw.HWND

// Page 1
var notebookbox1 dw.HBOX
var cursortogglebutton dw.HBUTTON
var noncheckable_menuitem, checkable_menuitem dw.HMENUITEM
var copypastefield, entryfield dw.HENTRYFIELD
var current_color dw.COLOR = dw.RGB(100, 100, 100)
var cursor_arrow bool = true
var timerid dw.HTIMER

// Page 2
var notebookbox2 dw.HBOX
var textbox1, textbox2 dw.HRENDER
var status1, status2 dw.HTEXT
var vscrollbar, hscrollbar dw.HSCROLLBAR
var rendcombo dw.HLISTBOX
var imagexspin, imageyspin dw.HSPINBUTTON
var imagestretchcheck dw.HBUTTON
var text1pm, text2pm, image dw.HPIXMAP
var image_x = 20
var image_y = 20
var image_stretch int = FALSE
var font_width = 8
var font_height = 12
var rows = 10
var width1 = 6
var cols = 80
var current_row = 0
var current_col = 0
var max_linewidth = 0
var SHAPES_DOUBLE_BUFFERED = 0
var SHAPES_DIRECT = 1
var DRAW_FILE = 2
var render_type = SHAPES_DOUBLE_BUFFERED

// Page 3
var notebookbox3 dw.HBOX
var tree dw.HTREE

// Page 4
var notebookbox4 dw.HBOX
var container_mle dw.HMLE
var container dw.HCONTAINER
var mle_point = 0

// Page 5
var notebookbox5, buttonboxperm, buttonsbox dw.HBOX
var combobox1, combobox2 dw.HLISTBOX
var cal dw.HCALENDAR
var spinbutton dw.HSPINBUTTON
var slider dw.HSLIDER
var percent dw.HPERCENT

// Page 7
var notebookbox7 dw.HBOX
var html dw.HHTML

// Page 8
var notebookbox8 dw.HBOX
var scrollbox dw.HSCROLLBOX
var MAX_WIDGETS = 20

var iteration = 0

// Page 9
var notebookbox9 dw.HBOX
var threadmle dw.HMLE
var startbutton dw.HBUTTON
var mutex dw.HMTX
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
		file   *os.File
		part   []byte
		prefix bool
		length int
		err    error
	)

	lines = nil
	max_linewidth = 0

	if file, err = os.Open(current_file); err != nil {
		return
	}
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	buffer.Reset()
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			length = len(buffer.String())
			if length > max_linewidth {
				max_linewidth = length
			}
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols))
	dw.Scrollbar_set_pos(hscrollbar, 0)
	dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows))
	dw.Scrollbar_set_pos(vscrollbar, 0)
}

// Call back section
func exit_callback(window dw.HWND, data dw.POINTER) int {
	if dw.Messagebox("dwtest", dw.MB_YESNO|dw.MB_QUESTION, "Are you sure you want to exit?") != 0 {
		dw.Main_quit()
	}
	return TRUE
}

func exit_menuitem_callback(window dw.HMENUITEM, data dw.POINTER) int {
	return exit_callback(dw.HWND(window), data)
}

func exit_button_callback(window dw.HBUTTON, data dw.POINTER) int {
	return exit_callback(dw.HWND(window), data)
}

func switch_page_callback(window dw.HNOTEBOOK, page_num dw.HNOTEPAGE, itemdata dw.POINTER) int {
	fmt.Printf("DW_SIGNAL_SWITCH_PAGE: PageNum: %d\n", dw.HNOTEPAGE_TO_UINT(page_num))
	return FALSE
}

func menu_callback(window dw.HMENUITEM, data dw.POINTER) int {
	info := dw.POINTER_TO_STRING(data)
	buf := fmt.Sprintf("%s menu item selected", info)
	dw.Messagebox("Menu Item Callback", dw.MB_OK|dw.MB_INFORMATION, buf)
	return FALSE
}

func menutoggle_callback(window dw.HMENUITEM, data dw.POINTER) int {
	if menu_enabled {
		dw.Window_set_style(checkable_menuitem, dw.MIS_DISABLED, dw.MIS_DISABLED)
		dw.Window_set_style(noncheckable_menuitem, dw.MIS_DISABLED, dw.MIS_DISABLED)
		menu_enabled = false
	} else {
		dw.Window_set_style(checkable_menuitem, dw.MIS_DISABLED, dw.MIS_ENABLED)
		dw.Window_set_style(noncheckable_menuitem, dw.MIS_DISABLED, dw.MIS_ENABLED)
		menu_enabled = true
	}
	return FALSE
}

func helpabout_callback(window dw.HMENUITEM, data dw.POINTER) int {
	var env dw.Env

	dw.Environment_query(&env)
	message := fmt.Sprintf("dwindows test\n\nOS: %s %s %s Version: %d.%d.%d.%d\n\nHTML: %s\n\ndwindows Version: %d.%d.%d\n\nScreen: %dx%d %dbpp",
		env.OSName, env.BuildDate, env.BuildTime,
		env.MajorVersion, env.MinorVersion, env.MajorBuild, env.MinorBuild,
		env.HTMLEngine, env.DWMajorVersion, env.DWMinorVersion, env.DWSubVersion,
		dw.Screen_width(), dw.Screen_height(), dw.Color_depth_get())
	dw.Messagebox("About dwindows", dw.MB_OK|dw.MB_INFORMATION, message)
	return FALSE
}

// Page 1 Callbacks
func paste_clicked_callback(button dw.HBUTTON, data dw.POINTER) int {
	test := dw.Clipboard_get_text()

	if len(test) > 0 {
		dw.Window_set_text(copypastefield, test)
	}
	return TRUE
}

func copy_clicked_callback(button dw.HBUTTON, data dw.POINTER) int {
	test := dw.Window_get_text(copypastefield)

	if len(test) > 0 {
		dw.Clipboard_set_text(test)
	}
	dw.Window_set_focus(entryfield)
	return TRUE
}

func notification_clicked_callback(button dw.HNOTIFICATION, data dw.POINTER) int {
	fmt.Printf("Notification clicked\n")
	return TRUE
}

func browse_file_callback(window dw.HBUTTON, data dw.POINTER) int {
	tmp := dw.File_browse("Pick a file", "dwtest.c", "go", dw.FILE_OPEN)
	if len(tmp) > 0 {
		notification := dw.Notification_new("New file loaded", "image/test.png", "dwtest loaded \""+tmp+"\" into the file browser on the Render tab, with \"File Display\" selected from the drop down list.")

		current_file = tmp
		dw.Window_set_text(entryfield, current_file)
		read_file()
		current_col = 0
		current_row = 0
		render_draw()
		dw.Signal_connect(notification, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(notification_clicked_callback), nil)
		dw.Notification_send(notification)
	}
	dw.Window_set_focus(copypastefield)
	return FALSE
}

func browse_folder_callback(window dw.HBUTTON, data dw.POINTER) int {
	tmp := dw.File_browse("Pick a folder", ".", "c", dw.DIRECTORY_OPEN)
	fmt.Printf("Folder picked: %s\n", tmp)
	return FALSE
}

func colorchoose_callback(window dw.HBUTTON, data dw.POINTER) int {
	current_color = dw.Color_choose(current_color)
	return FALSE
}

func cursortoggle_callback(window dw.HBUTTON, data dw.POINTER) int {
	if cursor_arrow {
		dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - ARROW")
		dw.Window_set_pointer(dw.POINTER_TO_HANDLE(data), dw.POINTER_CLOCK)
		cursor_arrow = false
	} else {
		dw.Window_set_text(cursortogglebutton, "Set Cursor pointer - CLOCK")
		dw.Window_set_pointer(dw.POINTER_TO_HANDLE(data), dw.POINTER_DEFAULT)
		cursor_arrow = true
	}
	return FALSE
}

func beep_callback(window dw.HBUTTON, data dw.POINTER) int {
	dw.Timer_disconnect(timerid)
	return TRUE
}

/* Beep every second */
func timer_callback(data dw.POINTER) int {
	dw.Beep(200, 200)

	/* Return TRUE so we get called again */
	return TRUE
}

// Page 2 Callbacks
func motion_notify_event(window dw.HRENDER, x int, y int, buttonmask int, data dw.POINTER) int {
	var which = "button_press"

	if uintptr(data) > 0 {
		which = "motion_notify"
	}
	dw.Window_set_text(status2, fmt.Sprintf("%s: %dx%d buttons %d", which, x, y, buttonmask))
	return FALSE
}

func show_window_callback(window dw.HMENUITEM, data dw.POINTER) int {
	thiswindow := dw.POINTER_TO_HANDLE(data)

	if thiswindow != dw.NOHWND {
		dw.Window_show(thiswindow)
		dw.Window_raise(thiswindow)
	}
	return TRUE
}

func context_menu_event(window dw.HANDLE, x int, y int, buttonmask int, data dw.POINTER) int {
	hwndMenu := dw.Menu_new(0)
	menuitem := dw.Menu_append_item(hwndMenu, "~Quit", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)

	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(exit_menuitem_callback), dw.HANDLE_TO_POINTER(mainwindow))
	dw.Menu_append_item(hwndMenu, dw.MENU_SEPARATOR, dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	menuitem = dw.Menu_append_item(hwndMenu, "~Show Window", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(show_window_callback), dw.HANDLE_TO_POINTER(mainwindow))
	px, py := dw.Pointer_query_pos()
	/* Use the toplevel window handle here.... because on the Mac..
	 * using the control itself, when a different tab is active
	 * the control is removed from the window and can no longer
	 * handle the messages.
	 */
	dw.Menu_popup(hwndMenu, mainwindow, px, py)
	return TRUE
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_file(row int, col int, nrows int, fheight int, hpma dw.HPIXMAP) {
	var hpm dw.HPIXMAP

	if hpma == dw.NOHPIXMAP {
		hpm = text2pm
	} else {
		hpm = hpma
	}

	if len(current_file) > 0 {
		var i int

		dw.Color_foreground_set(dw.CLR_WHITE)
		if hpma == dw.NOHPIXMAP {
			dw.Draw_rect(dw.NOHWND, text1pm, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(text1pm), dw.Pixmap_height(text1pm))
		}
		dw.Draw_rect(dw.NOHWND, hpm, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(hpm), dw.Pixmap_height(hpm))

		for i = 0; (i < nrows) && (i+row < len(lines)); i++ {
			fileline := i + row - 1
			y := i * fheight
			dw.Color_background_set(dw.COLOR(1 + (fileline % 15)))
			dw.Color_foreground_set(dw.COLOR(fileline % 16))
			if hpma == dw.NOHPIXMAP {
				dw.Draw_text(dw.NOHWND, text1pm, 0, y, fmt.Sprintf("%6.6d", i+row))
			}
			thisline := lines[i+row]
			if len(thisline) > col {
				dw.Draw_text(dw.NOHWND, hpm, 0, y, thisline[col:])
			}
		}
	}
}

/* When hpma is not NULL we are printing.. so handle things differently */
func draw_shapes(direct int, hpma dw.HPIXMAP) {
	var hpm, pixmap dw.HPIXMAP = dw.NOHPIXMAP, dw.NOHPIXMAP
	var window dw.HRENDER
	if hpma != dw.NOHPIXMAP {
		hpm = hpma
	} else {
		hpm = text2pm
	}
	if direct == TRUE {
		window = textbox2
	} else {
		pixmap = hpm
	}

	width := dw.Pixmap_width(hpm)
	height := dw.Pixmap_height(hpm)

	x := []int{20, 180, 180, 230, 180, 180, 20}
	y := []int{50, 50, 20, 70, 120, 90, 90}

	image_x = dw.Spinbutton_get_pos(imagexspin)
	image_y = dw.Spinbutton_get_pos(imageyspin)
	image_stretch = dw.Checkbox_get(imagestretchcheck)

	dw.Color_foreground_set(dw.CLR_WHITE)
	dw.Draw_rect(window, pixmap, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, width, height)
	dw.Color_foreground_set(dw.CLR_DARKPINK)
	dw.Draw_rect(window, pixmap, dw.DRAW_FILL|dw.DRAW_NOAA, 10, 10, width-20, height-20)
	dw.Color_foreground_set(dw.CLR_GREEN)
	dw.Color_background_set(dw.CLR_DARKRED)
	dw.Draw_text(window, pixmap, 10, 10, "This should be aligned with the edges.")
	dw.Color_foreground_set(dw.CLR_YELLOW)
	dw.Draw_line(window, pixmap, width-10, 10, 10, height-10)
	dw.Color_foreground_set(dw.CLR_BLUE)
	dw.Draw_polygon(window, pixmap, dw.DRAW_FILL, x, y)
	dw.Color_foreground_set(dw.CLR_BLACK)
	dw.Draw_rect(window, pixmap, dw.DRAW_FILL|dw.DRAW_NOAA, 80, 80, 80, 40)
	dw.Color_foreground_set(dw.CLR_CYAN)
	/* Bottom right corner */
	dw.Draw_arc(window, pixmap, 0, width-30, height-30, width-10, height-30, width-30, height-10)
	/* Top right corner */
	dw.Draw_arc(window, pixmap, 0, width-30, 30, width-30, 10, width-10, 30)
	/* Bottom left corner */
	dw.Draw_arc(window, pixmap, 0, 30, height-30, 30, height-10, 10, height-30)
	/* Full circle in the left top area */
	dw.Draw_arc(window, pixmap, dw.DRAW_FULL, 120, 100, 80, 80, 160, 120)
	if image != dw.NOHPIXMAP {
		if image_stretch == TRUE {
			dw.Pixmap_stretch_bitblt(window, pixmap, 10, 10, width-20, height-20, dw.NOHWND, image, 0, 0, dw.Pixmap_width(image), dw.Pixmap_height(image))
		} else {
			dw.Pixmap_bitblt(window, pixmap, image_x, image_y, dw.Pixmap_width(image), dw.Pixmap_height(image), dw.NOHWND, image, 0, 0)
		}
	}
}

func update_render() {
	switch render_type {
	case SHAPES_DOUBLE_BUFFERED:
		draw_shapes(FALSE, dw.NOHPIXMAP)
	case SHAPES_DIRECT:
		draw_shapes(TRUE, dw.NOHPIXMAP)
	case DRAW_FILE:
		draw_file(current_row, current_col, rows, font_height, dw.NOHPIXMAP)
	}
}

/* Request that the render widgets redraw...
 * If not using direct rendering, call update_render() to
 * redraw the in memory pixmaps. Then trigger the expose events.
 * Expose will call update_render() to draw directly or bitblt the pixmaps.
 */
func render_draw() {
	/* If we are double buffered, draw to the pixmaps */
	if render_type != SHAPES_DIRECT {
		update_render()
	}
	/* Trigger expose event */
	dw.Render_redraw(textbox1)
	dw.Render_redraw(textbox2)
}

func draw_page(print dw.HPRINT, pixmap dw.HPIXMAP, page_num int, data dw.POINTER) int {
	dw.Pixmap_set_font(pixmap, FIXEDFONT)
	if page_num == 0 {
		draw_shapes(FALSE, pixmap)
	} else if page_num == 1 {
		/* If we have a file to display... */
		if len(current_file) > 0 {
			/* Calculate new dimensions */
			_, fheight := dw.Font_text_extents_get(dw.NOHWND, pixmap, "(g")
			nrows := int(dw.Pixmap_height(pixmap) / fheight)

			/* Do the actual drawing */
			draw_file(0, 0, nrows, fheight, pixmap)
		} else {
			/* We don't have a file so center an error message on the page */
			var text = "No file currently selected!"

			/* Get the font size for this printer context... */
			fwidth, fheight := dw.Font_text_extents_get(dw.NOHWND, pixmap, text)

			posx := int(dw.Pixmap_width(pixmap)-fwidth) / 2
			posy := int(dw.Pixmap_height(pixmap)-fheight) / 2

			dw.Color_foreground_set(dw.CLR_BLACK)
			dw.Color_background_set(dw.CLR_WHITE)
			dw.Draw_text(dw.NOHWND, pixmap, posx, posy, text)
		}
	}
	return TRUE
}

func print_callback(window dw.HANDLE, data dw.POINTER) int {
	print := dw.Print_new("DWTest Job", 0, 2, dw.SIGNAL_FUNC(draw_page), nil)
	dw.Print_run(print, 0)
	return FALSE
}

/* This gets called when a part of the graph needs to be repainted. */
func text_expose(hwnd dw.HRENDER, x int, y int, width int, height int, data dw.POINTER) int {
	if render_type != SHAPES_DIRECT {
		var hpm dw.HPIXMAP

		if hwnd.GetHandle() == textbox1.GetHandle() {
			hpm = text1pm
		} else if hwnd.GetHandle() == textbox2.GetHandle() {
			hpm = text2pm
		} else {
			return TRUE
		}

		width = dw.Pixmap_width(hpm)
		height = dw.Pixmap_height(hpm)

		dw.Pixmap_bitblt(hwnd, dw.NOHPIXMAP, 0, 0, width, height, dw.NOHWND, hpm, 0, 0)
		dw.Flush()
	} else {
		update_render()
	}
	return TRUE
}

/* Handle size change of the main render window */
func configure_event(hwnd dw.HRENDER, width int, height int, data dw.POINTER) int {
	old1 := text1pm
	old2 := text2pm
	depth := dw.Color_depth_get()

	rows = height / font_height
	cols = width / font_width

	/* Create new pixmaps with the current sizes */
	text1pm = dw.Pixmap_new(textbox1, uint(font_width*(width1)), uint(height), depth)
	text2pm = dw.Pixmap_new(textbox2, uint(width), uint(height), depth)

	/* Make sure the side area is cleared */
	dw.Color_foreground_set(dw.CLR_WHITE)
	dw.Draw_rect(dw.NOHWND, text1pm, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, dw.Pixmap_width(text1pm), dw.Pixmap_height(text1pm))

	/* Destroy the old pixmaps */
	dw.Pixmap_destroy(old1)
	dw.Pixmap_destroy(old2)

	/* Update scrollbar ranges with new values */
	dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols))
	dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows))

	/* Redraw the render widgets */
	render_draw()
	return TRUE
}

func refresh_callback(window dw.HBUTTON, data dw.POINTER) int {
	render_draw()
	return FALSE
}

func render_select_event_callback(window dw.HLISTBOX, index int, data dw.POINTER) int {
	if index != render_type {
		if index == DRAW_FILE {
			dw.Scrollbar_set_range(hscrollbar, uint(max_linewidth), uint(cols))
			dw.Scrollbar_set_pos(hscrollbar, 0)
			dw.Scrollbar_set_range(vscrollbar, uint(len(lines)), uint(rows))
			dw.Scrollbar_set_pos(vscrollbar, 0)
			current_col = 0
			current_row = 0
		} else {
			dw.Scrollbar_set_range(hscrollbar, 0, 0)
			dw.Scrollbar_set_pos(hscrollbar, 0)
			dw.Scrollbar_set_range(vscrollbar, 0, 0)
			dw.Scrollbar_set_pos(vscrollbar, 0)
		}
		render_type = index
		render_draw()
	}
	return FALSE
}

/* Callback to handle user selection of the scrollbar position */
func scrollbar_valuechanged_callback(hwnd dw.HSCROLLBAR, value int, data dw.POINTER) int {
	if data != nil {
		stext := dw.POINTER_TO_HANDLE(data)

		if hwnd == vscrollbar {
			current_row = value
		} else {
			current_col = value
		}
		dw.Window_set_text(stext, fmt.Sprintf("Row:%d Col:%d Lines:%d Cols:%d", current_row, current_col, len(lines), max_linewidth))
		render_draw()
	}
	return FALSE
}

func resolve_keyname(vk int) string {
	var keyname string = "<unknown>"

	switch vk {
	case dw.VK_LBUTTON:
		keyname = "VK_LBUTTON"
	case dw.VK_RBUTTON:
		keyname = "VK_RBUTTON"
	case dw.VK_CANCEL:
		keyname = "VK_CANCEL"
	case dw.VK_MBUTTON:
		keyname = "VK_MBUTTON"
	case dw.VK_TAB:
		keyname = "VK_TAB"
	case dw.VK_CLEAR:
		keyname = "VK_CLEAR"
	case dw.VK_RETURN:
		keyname = "VK_RETURN"
	case dw.VK_PAUSE:
		keyname = "VK_PAUSE"
	case dw.VK_CAPITAL:
		keyname = "VK_CAPITAL"
	case dw.VK_ESCAPE:
		keyname = "VK_ESCAPE"
	case dw.VK_SPACE:
		keyname = "VK_SPACE"
	case dw.VK_PRIOR:
		keyname = "VK_PRIOR"
	case dw.VK_NEXT:
		keyname = "VK_NEXT"
	case dw.VK_END:
		keyname = "VK_END"
	case dw.VK_HOME:
		keyname = "VK_HOME"
	case dw.VK_LEFT:
		keyname = "VK_LEFT"
	case dw.VK_UP:
		keyname = "VK_UP"
	case dw.VK_RIGHT:
		keyname = "VK_RIGHT"
	case dw.VK_DOWN:
		keyname = "VK_DOWN"
	case dw.VK_SELECT:
		keyname = "VK_SELECT"
	case dw.VK_PRINT:
		keyname = "VK_PRINT"
	case dw.VK_EXECUTE:
		keyname = "VK_EXECUTE"
	case dw.VK_SNAPSHOT:
		keyname = "VK_SNAPSHOT"
	case dw.VK_INSERT:
		keyname = "VK_INSERT"
	case dw.VK_DELETE:
		keyname = "VK_DELETE"
	case dw.VK_HELP:
		keyname = "VK_HELP"
	case dw.VK_LWIN:
		keyname = "VK_LWIN"
	case dw.VK_RWIN:
		keyname = "VK_RWIN"
	case dw.VK_NUMPAD0:
		keyname = "VK_NUMPAD0"
	case dw.VK_NUMPAD1:
		keyname = "VK_NUMPAD1"
	case dw.VK_NUMPAD2:
		keyname = "VK_NUMPAD2"
	case dw.VK_NUMPAD3:
		keyname = "VK_NUMPAD3"
	case dw.VK_NUMPAD4:
		keyname = "VK_NUMPAD4"
	case dw.VK_NUMPAD5:
		keyname = "VK_NUMPAD5"
	case dw.VK_NUMPAD6:
		keyname = "VK_NUMPAD6"
	case dw.VK_NUMPAD7:
		keyname = "VK_NUMPAD7"
	case dw.VK_NUMPAD8:
		keyname = "VK_NUMPAD8"
	case dw.VK_NUMPAD9:
		keyname = "VK_NUMPAD9"
	case dw.VK_MULTIPLY:
		keyname = "VK_MULTIPLY"
	case dw.VK_ADD:
		keyname = "VK_ADD"
	case dw.VK_SEPARATOR:
		keyname = "VK_SEPARATOR"
	case dw.VK_SUBTRACT:
		keyname = "VK_SUBTRACT"
	case dw.VK_DECIMAL:
		keyname = "VK_DECIMAL"
	case dw.VK_DIVIDE:
		keyname = "VK_DIVIDE"
	case dw.VK_F1:
		keyname = "VK_F1"
	case dw.VK_F2:
		keyname = "VK_F2"
	case dw.VK_F3:
		keyname = "VK_F3"
	case dw.VK_F4:
		keyname = "VK_F4"
	case dw.VK_F5:
		keyname = "VK_F5"
	case dw.VK_F6:
		keyname = "VK_F6"
	case dw.VK_F7:
		keyname = "VK_F7"
	case dw.VK_F8:
		keyname = "VK_F8"
	case dw.VK_F9:
		keyname = "VK_F9"
	case dw.VK_F10:
		keyname = "VK_F10"
	case dw.VK_F11:
		keyname = "VK_F11"
	case dw.VK_F12:
		keyname = "VK_F12"
	case dw.VK_F13:
		keyname = "VK_F13"
	case dw.VK_F14:
		keyname = "VK_F14"
	case dw.VK_F15:
		keyname = "VK_F15"
	case dw.VK_F16:
		keyname = "VK_F16"
	case dw.VK_F17:
		keyname = "VK_F17"
	case dw.VK_F18:
		keyname = "VK_F18"
	case dw.VK_F19:
		keyname = "VK_F19"
	case dw.VK_F20:
		keyname = "VK_F20"
	case dw.VK_F21:
		keyname = "VK_F21"
	case dw.VK_F22:
		keyname = "VK_F22"
	case dw.VK_F23:
		keyname = "VK_F23"
	case dw.VK_F24:
		keyname = "VK_F24"
	case dw.VK_NUMLOCK:
		keyname = "VK_NUMLOCK"
	case dw.VK_SCROLL:
		keyname = "VK_SCROLL"
	case dw.VK_LSHIFT:
		keyname = "VK_LSHIFT"
	case dw.VK_RSHIFT:
		keyname = "VK_RSHIFT"
	case dw.VK_LCONTROL:
		keyname = "VK_LCONTROL"
	case dw.VK_RCONTROL:
		keyname = "VK_RCONTROL"
	}
	return keyname
}

func resolve_keymodifiers(mask int) string {
	if (mask&dw.KC_CTRL) == dw.KC_CTRL && (mask&dw.KC_SHIFT) == dw.KC_SHIFT && (mask&dw.KC_ALT) == dw.KC_ALT {
		return "KC_CTRL KC_SHIFT KC_ALT"
	} else if (mask&dw.KC_CTRL) == dw.KC_CTRL && (mask&dw.KC_SHIFT) == dw.KC_SHIFT {
		return "KC_CTRL KC_SHIFT"
	} else if (mask&dw.KC_CTRL) == dw.KC_CTRL && (mask&dw.KC_ALT) == dw.KC_ALT {
		return "KC_CTRL KC_ALT"
	} else if (mask&dw.KC_SHIFT) == dw.KC_SHIFT && (mask&dw.KC_ALT) == dw.KC_ALT {
		return "KC_SHIFT KC_ALT"
	} else if (mask & dw.KC_SHIFT) == dw.KC_SHIFT {
		return "KC_SHIFT"
	} else if (mask & dw.KC_CTRL) == dw.KC_CTRL {
		return "KC_CTRL"
	} else if (mask & dw.KC_ALT) == dw.KC_ALT {
		return "KC_ALT"
	}
	return "none"
}

func keypress_callback(window dw.HWND, ch uint8, vk int, state int, data dw.POINTER, utf8 string) int {
	var message string

	if ch != 0 {
		message = fmt.Sprintf("Key: %c(%d) Modifiers: %s(%d) utf8 %s", ch, ch, resolve_keymodifiers(state), state, utf8)
	} else {
		message = fmt.Sprintf("Key: %s(%d) Modifiers: %s(%d) utf8 %s", resolve_keyname(vk), vk, resolve_keymodifiers(state), state, utf8)
	}
	dw.Window_set_text(status1, message)
	return FALSE
}

// Page 3 and 4 Callbacks
func word_wrap_click_cb(wordwrap dw.HBUTTON, data dw.POINTER) int {
	container_mle := dw.POINTER_TO_HANDLE(data)

	dw.Mle_set_word_wrap(container_mle, dw.Checkbox_get(wordwrap))
	return TRUE
}

func color_combobox() dw.HLISTBOX {
	combobox := dw.Combobox_new("DW_CLR_DEFAULT", 0)

	dw.Listbox_append(combobox, "DW_CLR_DEFAULT")
	dw.Listbox_append(combobox, "DW_CLR_BLACK")
	dw.Listbox_append(combobox, "DW_CLR_DARKRED")
	dw.Listbox_append(combobox, "DW_CLR_DARKGREEN")
	dw.Listbox_append(combobox, "DW_CLR_BROWN")
	dw.Listbox_append(combobox, "DW_CLR_DARKBLUE")
	dw.Listbox_append(combobox, "DW_CLR_DARKPINK")
	dw.Listbox_append(combobox, "DW_CLR_DARKCYAN")
	dw.Listbox_append(combobox, "DW_CLR_PALEGRAY")
	dw.Listbox_append(combobox, "DW_CLR_DARKGRAY")
	dw.Listbox_append(combobox, "DW_CLR_RED")
	dw.Listbox_append(combobox, "DW_CLR_GREEN")
	dw.Listbox_append(combobox, "DW_CLR_YELLOW")
	dw.Listbox_append(combobox, "DW_CLR_BLUE")
	dw.Listbox_append(combobox, "DW_CLR_PINK")
	dw.Listbox_append(combobox, "DW_CLR_CYAN")
	dw.Listbox_append(combobox, "DW_CLR_WHITE")
	return combobox
}

func combobox_color(colortext string) dw.COLOR {
	color := dw.CLR_DEFAULT

	if colortext == "DW_CLR_BLACK" {
		color = dw.CLR_BLACK
	} else if colortext == "DW_CLR_DARKRED" {
		color = dw.CLR_DARKRED
	} else if colortext == "DW_CLR_DARKGREEN" {
		color = dw.CLR_DARKGREEN
	} else if colortext == "DW_CLR_BROWN" {
		color = dw.CLR_BROWN
	} else if colortext == "DW_CLR_DARKBLUE" {
		color = dw.CLR_DARKBLUE
	} else if colortext == "DW_CLR_DARKPINK" {
		color = dw.CLR_DARKPINK
	} else if colortext == "DW_CLR_DARKCYAN" {
		color = dw.CLR_DARKCYAN
	} else if colortext == "DW_CLR_PALEGRAY" {
		color = dw.CLR_PALEGRAY
	} else if colortext == "DW_CLR_DARKGRAY" {
		color = dw.CLR_DARKGRAY
	} else if colortext == "DW_CLR_RED" {
		color = dw.CLR_RED
	} else if colortext == "DW_CLR_GREEN" {
		color = dw.CLR_GREEN
	} else if colortext == "DW_CLR_YELLOW" {
		color = dw.CLR_YELLOW
	} else if colortext == "DW_CLR_BLUE" {
		color = dw.CLR_BLUE
	} else if colortext == "DW_CLR_PINK" {
		color = dw.CLR_PINK
	} else if colortext == "DW_CLR_CYAN" {
		color = dw.CLR_CYAN
	} else if colortext == "DW_CLR_WHITE" {
		color = dw.CLR_WHITE
	}

	return color
}

func mle_color_cb(hwnd dw.HLISTBOX, pos int, data dw.POINTER) int {
	hbox := dw.POINTER_TO_HANDLE(data)
	mlefore := dw.POINTER_TO_HANDLE(dw.Window_get_data(hbox, "mlefore"))
	mleback := dw.POINTER_TO_HANDLE(dw.Window_get_data(hbox, "mleback"))
	fore := dw.CLR_DEFAULT
	back := dw.CLR_DEFAULT

	if dw.Window_compare(hwnd, mlefore) {
		colortext := dw.Listbox_get_text(mlefore, pos)
		fore = combobox_color(colortext)
	} else {
		text := dw.Window_get_text(mlefore)

		if text != "" {
			fore = combobox_color(text)
		}
	}
	if dw.Window_compare(hwnd, mleback) {
		colortext := dw.Listbox_get_text(mleback, pos)
		back = combobox_color(colortext)
	} else {
		text := dw.Window_get_text(mleback)

		if text != "" {
			back = combobox_color(text)
		}
	}

	dw.Window_set_color(container_mle, fore, back)
	return FALSE
}

func mle_font_set(mle dw.HMLE, fontsize int, fontname string) {
	if fontname != "" {
		dw.Window_set_font(mle, fmt.Sprintf("%d.%s", fontsize, fontname))
	} else {
		dw.Window_set_font(mle, fontname)
	}
}

func mle_fontname_cb(hwnd dw.HLISTBOX, pos int, data dw.POINTER) int {
	hbox := dw.POINTER_TO_HANDLE(data)
	fontsize := dw.POINTER_TO_HANDLE(dw.Window_get_data(hbox, "fontsize"))
	fontname := dw.POINTER_TO_HANDLE(dw.Window_get_data(hbox, "fontname"))
	font := dw.Listbox_get_text(fontname, pos)

	if font == "Default" {
		font = ""
	}

	mle_font_set(container_mle, dw.Spinbutton_get_pos(fontsize), font)
	return FALSE
}

func mle_fontsize_cb(hwnd dw.HSPINBUTTON, size int, data dw.POINTER) int {
	hbox := dw.POINTER_TO_HANDLE(data)
	fontname := dw.POINTER_TO_HANDLE(dw.Window_get_data(hbox, "fontname"))
	font := dw.Window_get_text(fontname)

	if font != "" {
		if font == "Default" {
			font = ""
		}
		mle_font_set(container_mle, size, font)
	} else {
		mle_font_set(container_mle, size, "")
	}
	return FALSE
}

func item_enter_cb(window dw.HCONTAINER, text string, data dw.POINTER, itemdata dw.POINTER) int {
	message := fmt.Sprintf("DW_SIGNAL_ITEM_ENTER: Window: %x Text: %s Itemdata: %x", dw.HANDLE_TO_UINTPTR(window), text, uintptr(itemdata))
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	return FALSE
}

/* Context menus */
func context_menu_cb(hwnd dw.HMENUITEM, data dw.POINTER) int {
	statline := dw.POINTER_TO_HANDLE(data)

	dw.Window_set_text(statline, fmt.Sprintf("DW_SIGNAL_CLICKED: Menu: %x Container context menu clicked", dw.HANDLE_TO_UINTPTR(hwnd)))
	return FALSE
}

func item_context_menu_new(text string, data dw.POINTER) dw.HMENUI {
	hwndMenu := dw.Menu_new(0)
	hwndSubMenu := dw.Menu_new(0)
	menuitem := dw.Menu_append_item(hwndSubMenu, "File", dw.MENU_POPUP, 0, TRUE, TRUE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)
	menuitem = dw.Menu_append_item(hwndSubMenu, "Date", dw.MENU_POPUP, 0, TRUE, TRUE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)
	menuitem = dw.Menu_append_item(hwndSubMenu, "Size", dw.MENU_POPUP, 0, TRUE, TRUE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)
	menuitem = dw.Menu_append_item(hwndSubMenu, "None", dw.MENU_POPUP, 0, TRUE, TRUE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)

	menuitem = dw.Menu_append_item(hwndMenu, "Sort", dw.MENU_POPUP, 0, TRUE, FALSE, hwndSubMenu)

	menuitem = dw.Menu_append_item(hwndMenu, "Make Directory", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)

	dw.Menu_append_item(hwndMenu, "", 0, 0, TRUE, FALSE, dw.NOMENU)
	menuitem = dw.Menu_append_item(hwndMenu, "Rename Entry", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)

	menuitem = dw.Menu_append_item(hwndMenu, "Delete Entry", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)

	dw.Menu_append_item(hwndMenu, "", 0, 0, TRUE, FALSE, dw.NOMENU)
	menuitem = dw.Menu_append_item(hwndMenu, "View File", dw.MENU_POPUP, 0, TRUE, FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(context_menu_cb), data)

	return hwndMenu
}

func item_context_cb(window dw.HCONTAINER, text string, x int, y int, data dw.POINTER, itemdata dw.POINTER) int {
	message := fmt.Sprintf("DW_SIGNAL_ITEM_CONTEXT: Window: %x Text: %s x: %d y: %d Itemdata: %x", dw.HANDLE_TO_UINTPTR(window), text, x, y, uintptr(itemdata))
	popupmenu := item_context_menu_new(text, data)
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	dw.Menu_popup(popupmenu, mainwindow, x, y)
	return FALSE
}

func list_select_cb(window dw.HLISTBOX, item int, data dw.POINTER) int {
	message := fmt.Sprintf("DW_SIGNAL_LIST_SELECT: Window: %x Item: %d", dw.HANDLE_TO_UINTPTR(window), item)
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	return FALSE
}

func item_select_cb(window dw.HTREE, item dw.HTREEITEM, text string, data dw.POINTER, itemdata dw.POINTER) int {
	message := fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", dw.HANDLE_TO_UINTPTR(window),
		dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata))
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	return FALSE
}

func container_select_cb(window dw.HCONTAINER, item dw.HTREEITEM, text string, data dw.POINTER, itemdata dw.POINTER) int {
	message := fmt.Sprintf("DW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x", dw.HANDLE_TO_UINTPTR(window),
		dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata))
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	message = fmt.Sprintf("\r\nDW_SIGNAL_ITEM_SELECT: Window: %x Item: %x Text: %s Itemdata: %x\r\n", dw.HANDLE_TO_UINTPTR(window),
		dw.HANDLE_TO_UINTPTR(item), text, uintptr(itemdata))
	mle_point = dw.Mle_import(container_mle, message, mle_point)
	str := dw.Container_query_start(container, dw.CRA_SELECTED)
	for len(str) > 0 {
		mle_point = dw.Mle_import(container_mle, fmt.Sprintf("Selected: %s\r\n", str), mle_point)
		str = dw.Container_query_next(container, dw.CRA_SELECTED)
	}
	/* Make the last inserted point the cursor location */
	dw.Mle_set_cursor(container_mle, mle_point)
	/* set the details of item 0 to new data */
	dw.Filesystem_change_file(container, 0, "new data", fileicon)
	dw.Filesystem_change_item_ulong(container, 1, 0, 999)
	return FALSE
}

func combobox_select_event_callback(window dw.HLISTBOX, index int, data dw.POINTER) int {
	fmt.Printf("got combobox_select_event for index: %d, iteration: %d\n", index, iteration)
	iteration++
	return FALSE
}

func column_click_cb(window dw.HCONTAINER, column_num int, data dw.POINTER) int {
	var stype = "Unknown"

	if column_num == 0 {
		stype = "Filename"
	} else {
		column_type := dw.Filesystem_get_column_type(window, column_num-1)
		if column_type == dw.CFA_STRING {
			stype = "String"
		} else if column_type == dw.CFA_ULONG {
			stype = "ULong"
		} else if column_type == dw.CFA_DATE {
			stype = "Date"
		} else if column_type == dw.CFA_TIME {
			stype = "Time"
		} else if column_type == dw.CFA_BITMAPORICON {
			stype = "BitmapOrIcon"
		}
	}
	message := fmt.Sprintf("DW_SIGNAL_COLUMN_CLICK: Window: %x Column: %d Type: %s Itemdata: %x", dw.HANDLE_TO_UINTPTR(window),
		column_num, stype)
	dw.Window_set_text(dw.POINTER_TO_HANDLE(data), message)
	return FALSE
}

// Page 5 Callbacks
func button_callback(window dw.HBUTTON, data dw.POINTER) int {
	idx := dw.Listbox_selected(combobox1)
	buf1 := dw.Listbox_get_text(combobox1, idx)
	idx = dw.Listbox_selected(combobox2)
	buf2 := dw.Listbox_get_text(combobox2, idx)
	y, m, d := dw.Calendar_get_date(cal)
	spvalue := dw.Spinbutton_get_pos(spinbutton)
	message := fmt.Sprintf("spinbutton: %d\ncombobox1: \"%s\"\ncombobox2: \"%s\"\ncalendar: %d-%d-%d",
		spvalue,
		buf1, buf2,
		y, m, d)
	dw.Messagebox("Values", dw.MB_OK|dw.MB_INFORMATION, message)
	return FALSE
}

var isfoldericon bool = true

func bitmap_toggle_callback(window dw.HBUTTON, data dw.POINTER) int {
	if isfoldericon == true {
		isfoldericon = false
		dw.Window_set_bitmap(window, 0, FILE_ICON_NAME)
		dw.Window_set_tooltip(window, "File Icon")
	} else {
		isfoldericon = true
		//dw.Window_set_bitmap_from_data(window, 0, folder_ico, sizeof(folder_ico));
		dw.Window_set_tooltip(window, "Folder Icon")
	}
	return FALSE
}

func percent_button_box_callback(window dw.HBUTTON, data dw.POINTER) int {
	dw.Percent_set_pos(percent, dw.PERCENT_INDETERMINATE)
	return FALSE
}

func change_color_red_callback(window dw.HBUTTON, data dw.POINTER) int {
	dw.Window_set_color(buttonsbox, dw.CLR_RED, dw.CLR_RED)
	return FALSE
}

func change_color_yellow_callback(window dw.HBUTTON, data dw.POINTER) int {
	dw.Window_set_color(buttonsbox, dw.CLR_YELLOW, dw.CLR_YELLOW)
	return FALSE
}

/* Callback to handle user selection of the spinbutton position */
func spinbutton_valuechanged_callback(hwnd dw.HSPINBUTTON, value int, data dw.POINTER) int {
	dw.Messagebox("DWTest", dw.MB_OK, fmt.Sprintf("New value from spinbutton: %d\n", value))
	return FALSE
}

/* Callback to handle user selection of the slider position */
func slider_valuechanged_callback(hwnd dw.HSLIDER, value int, data dw.POINTER) int {
	dw.Percent_set_pos(percent, uint(value*10))
	return FALSE
}

// Page 8 Callbacks
func scrollbox_button_callback(window dw.HBUTTON, data dw.POINTER) int {
	_, pos := dw.Scrollbox_get_pos(scrollbox)
	_, rng := dw.Scrollbox_get_range(scrollbox)
	fmt.Printf("Pos %d Range %d\n", pos, rng)
	return FALSE
}

// Page 9 Callbacks
func run_thread(threadnum int) {
	dw.InitThread()
	update_mle(fmt.Sprintf("Thread %d started.\r\n", threadnum), TRUE)

	/* Increment the ready count while protected by mutex */
	dw.Mutex_lock(mutex)
	ready++
	/* If all 4 threads have incrememted the ready count...
	 * Post the control event semaphore so things will get started.
	 */
	if ready == 4 {
		dw.Event_post(controlevent)
	}
	dw.Mutex_unlock(mutex)

	for finished == 0 {
		result := dw.Event_wait(workevent, 2000)

		if result == dw.ERROR_TIMEOUT {
			update_mle(fmt.Sprintf("Thread %d timeout waiting for event.\r\n", threadnum), dw.TRUE)
		} else if result == dw.ERROR_NONE {
			update_mle(fmt.Sprintf("Thread %d doing some work.\r\n", threadnum), dw.TRUE)
			/* Pretend to do some work */
			dw.Main_sleep(1000 * threadnum)

			/* Increment the ready count while protected by mutex */
			dw.Mutex_lock(mutex)
			ready++
			buf := fmt.Sprintf("Thread %d work done. ready=%d", threadnum, ready)
			/* If all 4 threads have incrememted the ready count...
			 * Post the control event semaphore so things will get started.
			 */
			if ready == 4 {
				dw.Event_post(controlevent)
				buf = fmt.Sprintf("%s%s", buf, " Control posted.")
			}
			dw.Mutex_unlock(mutex)
			update_mle(fmt.Sprintf("%s\r\n", buf), dw.TRUE)
		} else {
			update_mle(fmt.Sprintf("Thread %d error %d.\r\n", threadnum), dw.TRUE)
			dw.Main_sleep(10000)
		}
	}
	update_mle(fmt.Sprintf("Thread %d finished.\r\n", threadnum), dw.TRUE)
	dw.DeinitThread()
}

func control_thread() {
	dw.InitThread()

	inprogress := 5

	for inprogress != 0 {
		result := dw.Event_wait(controlevent, 2000)

		if result == dw.ERROR_TIMEOUT {
			update_mle("Control thread timeout waiting for event.\r\n", dw.TRUE)
		} else if result == dw.ERROR_NONE {
			/* Reset the control event */
			dw.Event_reset(controlevent)
			ready = 0
			update_mle(fmt.Sprintf("Control thread starting worker threads. Inprogress=%d\r\n", inprogress), dw.TRUE)
			/* Start the work threads */
			dw.Event_post(workevent)
			dw.Main_sleep(100)
			/* Reset the work event */
			dw.Event_reset(workevent)
			inprogress--
		} else {
			update_mle(fmt.Sprintf("Control thread error %d.\r\n", result), dw.TRUE)
			dw.Main_sleep(10000)
		}
	}
	/* Tell the other threads we are done */
	finished = dw.TRUE
	dw.Event_post(workevent)
	/* Close the control event */
	dw.Event_close(&controlevent)
	update_mle("Control thread finished.\r\n", dw.TRUE)
	dw.Window_enable(startbutton)
	dw.DeinitThread()
}

func start_threads_button_callback(window dw.HWND, data dw.POINTER) int {
	dw.Window_disable(startbutton)
	dw.Mutex_lock(mutex)
	controlevent = dw.Event_new()
	dw.Event_reset(workevent)
	finished = FALSE
	ready = 0
	update_mle("Starting thread 1\r\n", FALSE)
	go run_thread(1)
	update_mle("Starting thread 2\r\n", FALSE)
	go run_thread(2)
	update_mle("Starting thread 3\r\n", FALSE)
	go run_thread(3)
	update_mle("Starting thread 4\r\n", FALSE)
	go run_thread(4)
	update_mle("Starting control thread\r\n", FALSE)
	go control_thread()
	dw.Mutex_unlock(mutex)
	return FALSE
}

/* Handle web back navigation */
func web_back_clicked(button dw.HBUTTON, data dw.POINTER) int {
	dw.Html_action(dw.POINTER_TO_HANDLE(data), dw.HTML_GOBACK)
	return FALSE
}

/* Handle web forward navigation */
func web_forward_clicked(button dw.HBUTTON, data dw.POINTER) int {
	dw.Html_action(dw.POINTER_TO_HANDLE(data), dw.HTML_GOFORWARD)
	return FALSE
}

/* Handle web reload */
func web_reload_clicked(button dw.HBUTTON, data dw.POINTER) int {
	dw.Html_action(dw.POINTER_TO_HANDLE(data), dw.HTML_RELOAD)
	return FALSE
}

/* Handle web run */
func web_run_clicked(button dw.HBUTTON, data dw.POINTER) int {
	html := dw.POINTER_TO_HANDLE(data)
	javascript := dw.POINTER_TO_HANDLE(dw.Window_get_data(button, "javascript"))
	script := dw.Window_get_text(javascript)

	dw.Html_javascript_run(html, script, nil)
	return FALSE
}

/* Handle web javascript result */
func web_html_result(html dw.HHTML, status int, result string, script_data dw.POINTER, user_data dw.POINTER) int {
	var style = dw.MB_INFORMATION
	var message = result
	if status != dw.ERROR_NONE {
		style = dw.MB_ERROR
	}
	if result == "" {
		message = "Javascript result is not a string value"
	}
	dw.Messagebox("Javascript Result", style, message)
	return TRUE
}

/* Handle web html changed */
func web_html_changed(html dw.HHTML, status int, url string, data dw.POINTER) int {
	statusnames := []string{"none", "started", "redirect", "loading", "complete"}

	if status < 5 {
		dw.Window_set_text(dw.POINTER_TO_HANDLE(data), "Status "+statusnames[status]+": "+url)
	}
	return FALSE
}

var checkable_string = "checkable"
var noncheckable_string = "non-checkable"

// Create the menu
func menu_add() {
	mainmenubar := dw.Menubar_new(mainwindow)
	/* add menus to the menubar */
	menu := dw.Menu_new(0)
	menuitem := dw.Menu_append_item(menu, "~Quit", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(exit_menuitem_callback), dw.HANDLE_TO_POINTER(mainwindow))
	/*
	 * Add the "File" menu to the menubar...
	 */
	dw.Menu_append_item(mainmenubar, "~File", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu)

	changeable_menu := dw.Menu_new(0)
	checkable_menuitem = dw.Menu_append_item(changeable_menu, "~Checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU)
	dw.Signal_connect(checkable_menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(menu_callback), dw.OBJECT_TO_POINTER(checkable_string))
	noncheckable_menuitem = dw.Menu_append_item(changeable_menu, "~Non-checkable Menu Item", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU)
	dw.Signal_connect(noncheckable_menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(menu_callback), dw.OBJECT_TO_POINTER(noncheckable_string))
	dw.Menu_append_item(changeable_menu, "~Disabled menu Item", dw.MENU_AUTO, dw.MIS_DISABLED|dw.MIS_CHECKED, dw.TRUE, dw.TRUE, dw.NOMENU)
	/* seperator */
	dw.Menu_append_item(changeable_menu, dw.MENU_SEPARATOR, dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU)
	menuitem = dw.Menu_append_item(changeable_menu, "~Menu Items Disabled", dw.MENU_AUTO, 0, dw.TRUE, dw.TRUE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(menutoggle_callback), nil)
	/*
	 * Add the "Menu" menu to the menubar...
	 */
	dw.Menu_append_item(mainmenubar, "~Menu", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, changeable_menu)

	menu = dw.Menu_new(0)
	menuitem = dw.Menu_append_item(menu, "~About", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, dw.NOMENU)
	dw.Signal_connect(menuitem, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(helpabout_callback), dw.HANDLE_TO_POINTER(mainwindow))
	/*
	 * Add the "Help" menu to the menubar...
	 */
	dw.Menu_append_item(mainmenubar, "~Help", dw.MENU_AUTO, 0, dw.TRUE, dw.FALSE, menu)
}

// Create Page 1
func archive_add() {
	lbbox := dw.Box_new(dw.VERT, 10)

	dw.Box_pack_start(notebookbox1, lbbox, 150, 70, dw.TRUE, dw.TRUE, 0)

	/* Copy and Paste */
	browsebox := dw.Box_new(dw.HORZ, 0)

	dw.Box_pack_start(lbbox, browsebox, 0, 0, dw.FALSE, dw.FALSE, 0)

	copypastefield = dw.Entryfield_new("", 0)

	dw.Entryfield_set_limit(copypastefield, 260)

	dw.Box_pack_start(browsebox, copypastefield, dw.SIZE_AUTO, dw.SIZE_AUTO, dw.TRUE, dw.FALSE, 4)

	copybutton := dw.Button_new("Copy", 0)

	dw.Box_pack_start(browsebox, copybutton, dw.SIZE_AUTO, dw.SIZE_AUTO, dw.FALSE, dw.FALSE, 0)

	pastebutton := dw.Button_new("Paste", 0)

	dw.Box_pack_start(browsebox, pastebutton, dw.SIZE_AUTO, dw.SIZE_AUTO, dw.FALSE, dw.FALSE, 0)

	/* Archive Name */
	stext := dw.Text_new("File to browse", 0)

	dw.Window_set_style(stext, dw.DT_VCENTER, dw.DT_VCENTER)

	dw.Box_pack_start(lbbox, stext, 130, 15, dw.TRUE, dw.TRUE, 2)

	browsebox = dw.Box_new(dw.HORZ, 0)

	dw.Box_pack_start(lbbox, browsebox, 0, 0, dw.TRUE, dw.TRUE, 0)

	entryfield = dw.Entryfield_new("", 100)

	dw.Entryfield_set_limit(entryfield, 260)

	dw.Box_pack_start(browsebox, entryfield, 100, 15, dw.TRUE, dw.TRUE, 4)

	browsefilebutton := dw.Button_new("Browse File", 1001)

	dw.Box_pack_start(browsebox, browsefilebutton, 40, 15, dw.TRUE, dw.TRUE, 0)

	browsefolderbutton := dw.Button_new("Browse Folder", 1001)

	dw.Box_pack_start(browsebox, browsefolderbutton, 40, 15, dw.TRUE, dw.TRUE, 0)

	dw.Window_set_color(browsebox, dw.CLR_PALEGRAY, dw.CLR_PALEGRAY)
	dw.Window_set_color(stext, dw.CLR_BLACK, dw.CLR_PALEGRAY)

	/* Buttons */
	buttonbox := dw.Box_new(dw.HORZ, 10)

	dw.Box_pack_start(lbbox, buttonbox, 0, 0, dw.TRUE, dw.TRUE, 0)

	cancelbutton := dw.Button_new("Exit", 1002)
	dw.Box_pack_start(buttonbox, cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2)

	cursortogglebutton = dw.Button_new("Set Cursor pointer - CLOCK", 1003)
	dw.Box_pack_start(buttonbox, cursortogglebutton, 130, 30, dw.TRUE, dw.TRUE, 2)

	okbutton := dw.Button_new("Turn Off Annoying Beep!", 1001)
	dw.Box_pack_start(buttonbox, okbutton, 130, 30, dw.TRUE, dw.TRUE, 2)

	dw.Box_unpack(cancelbutton)
	dw.Box_pack_start(buttonbox, cancelbutton, 130, 30, dw.TRUE, dw.TRUE, 2)
	dw.Window_click_default(mainwindow, cancelbutton)

	colorchoosebutton := dw.Button_new("Color Chooser Dialog", 1004)
	dw.Box_pack_at_index(buttonbox, colorchoosebutton, 1, 130, 30, dw.TRUE, dw.TRUE, 2)

	/* Set some nice fonts and colors */
	dw.Window_set_color(lbbox, dw.CLR_DARKCYAN, dw.CLR_PALEGRAY)
	dw.Window_set_color(buttonbox, dw.CLR_DARKCYAN, dw.CLR_PALEGRAY)
	dw.Window_set_color(okbutton, dw.CLR_PALEGRAY, dw.CLR_DARKCYAN)

	dw.Signal_connect(browsefilebutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(browse_file_callback), nil)
	dw.Signal_connect(browsefolderbutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(browse_folder_callback), nil)
	dw.Signal_connect(copybutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(copy_clicked_callback), dw.HANDLE_TO_POINTER(copypastefield))
	dw.Signal_connect(pastebutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(paste_clicked_callback), dw.HANDLE_TO_POINTER(copypastefield))
	dw.Signal_connect(okbutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(beep_callback), nil)
	dw.Signal_connect(cancelbutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(exit_button_callback), dw.HANDLE_TO_POINTER(mainwindow))
	dw.Signal_connect(cursortogglebutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(cursortoggle_callback), dw.HANDLE_TO_POINTER(mainwindow))
	dw.Signal_connect(colorchoosebutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(colorchoose_callback), dw.HANDLE_TO_POINTER(mainwindow))
}

// Create Page 2
func text_add() {
	depth := dw.Color_depth_get()

	/* create a box to pack into the notebook page */
	pagebox := dw.Box_new(dw.HORZ, 2)
	dw.Box_pack_start(notebookbox2, pagebox, 0, 0, dw.TRUE, dw.TRUE, 0)
	/* now a status area under this box */
	hbox := dw.Box_new(dw.HORZ, 1)
	dw.Box_pack_start(notebookbox2, hbox, 100, 20, dw.TRUE, dw.FALSE, 1)
	status1 = dw.Status_text_new("", 0)
	dw.Box_pack_start(hbox, status1, 100, dw.SIZE_AUTO, dw.TRUE, dw.FALSE, 1)
	status2 = dw.Status_text_new("", 0)
	dw.Box_pack_start(hbox, status2, 100, dw.SIZE_AUTO, dw.TRUE, dw.FALSE, 1)
	/* a box with combobox and button */
	hbox = dw.Box_new(dw.HORZ, 1)
	dw.Box_pack_start(notebookbox2, hbox, 100, 25, dw.TRUE, dw.FALSE, 1)
	rendcombo := dw.Combobox_new("Shapes Double Buffered", 0)
	dw.Box_pack_start(hbox, rendcombo, 80, 25, dw.TRUE, dw.TRUE, 0)
	dw.Listbox_append(rendcombo, "Shapes Double Buffered")
	dw.Listbox_append(rendcombo, "Shapes Direct")
	dw.Listbox_append(rendcombo, "File Display")
	label := dw.Text_new("Image X:", 100)
	dw.Window_set_style(label, dw.DT_VCENTER|dw.DT_CENTER, dw.DT_VCENTER|dw.DT_CENTER)
	dw.Box_pack_start(hbox, label, dw.SIZE_AUTO, 25, dw.FALSE, dw.TRUE, 0)
	imagexspin = dw.Spinbutton_new("20", 1021)
	dw.Box_pack_start(hbox, imagexspin, 25, 25, dw.TRUE, dw.TRUE, 0)
	label = dw.Text_new("Y:", 100)
	dw.Window_set_style(label, dw.DT_VCENTER|dw.DT_CENTER, dw.DT_VCENTER|dw.DT_CENTER)
	dw.Box_pack_start(hbox, label, dw.SIZE_AUTO, 25, dw.FALSE, dw.TRUE, 0)
	imageyspin = dw.Spinbutton_new("20", 1021)
	dw.Box_pack_start(hbox, imageyspin, 25, 25, dw.TRUE, dw.TRUE, 0)
	dw.Spinbutton_set_limits(imagexspin, 2000, 0)
	dw.Spinbutton_set_limits(imageyspin, 2000, 0)
	dw.Spinbutton_set_pos(imagexspin, 20)
	dw.Spinbutton_set_pos(imageyspin, 20)
	imagestretchcheck = dw.Checkbox_new("Stretch", 1021)
	dw.Box_pack_start(hbox, imagestretchcheck, dw.SIZE_AUTO, 25, dw.FALSE, dw.TRUE, 0)

	button1 := dw.Button_new("Refresh", 1223)
	dw.Box_pack_start(hbox, button1, dw.SIZE_AUTO, 25, dw.FALSE, dw.TRUE, 0)
	button2 := dw.Button_new("Print", 1224)
	dw.Box_pack_start(hbox, button2, dw.SIZE_AUTO, 25, dw.FALSE, dw.TRUE, 0)

	/* Pre-create the scrollbars so we can query their sizes */
	vscrollbar = dw.Scrollbar_new(dw.VERT, 50)
	hscrollbar = dw.Scrollbar_new(dw.HORZ, 50)
	vscrollbarwidth, _ := dw.Window_get_preferred_size(vscrollbar)
	_, hscrollbarheight := dw.Window_get_preferred_size(hscrollbar)

	/* On GTK with overlay scrollbars enabled this returns us 0...
	 * so in that case we need to give it some real values.
	 */
	if vscrollbarwidth == 0 {
		vscrollbarwidth = 8
	}
	if hscrollbarheight == 0 {
		hscrollbarheight = 8
	}

	/* create render box for number pixmap */
	textbox1 = dw.Render_new(100)
	dw.Window_set_font(textbox1, FIXEDFONT)
	font_width, font_height = dw.Font_text_extents_get(textbox1, dw.NOHPIXMAP, "(g")
	font_width = font_width / 2
	vscrollbox := dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(vscrollbox, textbox1, font_width*width1, font_height*rows, dw.FALSE, dw.TRUE, 0)
	dw.Box_pack_start(vscrollbox, dw.NOHWND, font_width*(width1+1), hscrollbarheight, dw.FALSE, dw.FALSE, 0)
	dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0)

	/* pack empty space 1 character wide */
	dw.Box_pack_start(pagebox, dw.NOHWND, font_width, 0, dw.FALSE, dw.TRUE, 0)

	/* create box for filecontents and horz scrollbar */
	textboxA := dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(pagebox, textboxA, 0, 0, dw.TRUE, dw.TRUE, 0)

	/* create render box for filecontents pixmap */
	textbox2 = dw.Render_new(101)
	dw.Box_pack_start(textboxA, textbox2, 10, 10, dw.TRUE, dw.TRUE, 0)
	dw.Window_set_font(textbox2, FIXEDFONT)
	/* create horizonal scrollbar */
	dw.Box_pack_start(textboxA, hscrollbar, dw.SIZE_AUTO, dw.SIZE_AUTO, dw.TRUE, dw.FALSE, 0)

	/* create vertical scrollbar */
	vscrollbox = dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(vscrollbox, vscrollbar, dw.SIZE_AUTO, dw.SIZE_AUTO, dw.FALSE, dw.TRUE, 0)
	/* Pack an area of empty space 14x14 pixels */
	dw.Box_pack_start(vscrollbox, dw.NOHWND, vscrollbarwidth, hscrollbarheight, dw.FALSE, dw.FALSE, 0)
	dw.Box_pack_start(pagebox, vscrollbox, 0, 0, dw.FALSE, dw.TRUE, 0)

	text1pm = dw.Pixmap_new(textbox1, uint(font_width*width1), uint(font_height*rows), depth)
	text2pm = dw.Pixmap_new(textbox2, uint(font_width*cols), uint(font_height*rows), depth)
	image = dw.Pixmap_new_from_file(textbox2, "test")
	if image == dw.NOHPIXMAP && len(SRCROOT) > 0 {
		image = dw.Pixmap_new_from_file(textbox2, fmt.Sprintf("%s/test", SRCROOT))
	}
	if image != dw.NOHPIXMAP {
		dw.Pixmap_set_transparent_color(image, dw.CLR_WHITE)
	}

	dw.Messagebox("DWTest", dw.MB_OK|dw.MB_INFORMATION, fmt.Sprintf("Width: %d Height: %d\n", font_width, font_height))
	dw.Draw_rect(dw.NOHWND, text1pm, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, font_width*width1, font_height*rows)
	dw.Draw_rect(dw.NOHWND, text2pm, dw.DRAW_FILL|dw.DRAW_NOAA, 0, 0, font_width*cols, font_height*rows)
	dw.Signal_connect(textbox1, dw.SIGNAL_BUTTON_PRESS, dw.SIGNAL_FUNC(context_menu_event), nil)
	dw.Signal_connect(textbox1, dw.SIGNAL_EXPOSE, dw.SIGNAL_FUNC(text_expose), nil)
	dw.Signal_connect(textbox2, dw.SIGNAL_EXPOSE, dw.SIGNAL_FUNC(text_expose), nil)
	dw.Signal_connect(textbox2, dw.SIGNAL_CONFIGURE, dw.SIGNAL_FUNC(configure_event), nil)
	dw.Signal_connect(textbox2, dw.SIGNAL_MOTION_NOTIFY, dw.SIGNAL_FUNC(motion_notify_event), dw.POINTER(uintptr(1)))
	dw.Signal_connect(textbox2, dw.SIGNAL_BUTTON_PRESS, dw.SIGNAL_FUNC(motion_notify_event), nil)
	dw.Signal_connect(hscrollbar, dw.SIGNAL_VALUE_CHANGED, dw.SIGNAL_FUNC(scrollbar_valuechanged_callback), dw.HANDLE_TO_POINTER(status1))
	dw.Signal_connect(vscrollbar, dw.SIGNAL_VALUE_CHANGED, dw.SIGNAL_FUNC(scrollbar_valuechanged_callback), dw.HANDLE_TO_POINTER(status1))
	dw.Signal_connect(imagestretchcheck, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(refresh_callback), nil)
	dw.Signal_connect(button1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(refresh_callback), nil)
	dw.Signal_connect(button2, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(print_callback), nil)
	dw.Signal_connect(rendcombo, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(render_select_event_callback), nil)
	dw.Signal_connect(mainwindow, dw.SIGNAL_KEY_PRESS, dw.SIGNAL_FUNC(keypress_callback), nil)

	dw.Taskbar_insert(textbox1, fileicon, "DWTest")
}

// Page 3
func tree_add() {
	/* create a box to pack into the notebook page */
	listbox := dw.Listbox_new(1024, TRUE)
	dw.Box_pack_start(notebookbox3, listbox, 500, 200, TRUE, TRUE, 0)
	dw.Listbox_append(listbox, "Test 1")
	dw.Listbox_append(listbox, "Test 2")
	dw.Listbox_append(listbox, "Test 3")
	dw.Listbox_append(listbox, "Test 4")
	dw.Listbox_append(listbox, "Test 5")

	/* now a tree area under this box */
	tree = dw.Tree_new(101)
	dw.Box_pack_start(notebookbox3, tree, 500, 200, TRUE, TRUE, 1)

	/* and a status area to see whats going on */
	tree_status := dw.Status_text_new("", 0)
	dw.Box_pack_start(notebookbox3, tree_status, 100, dw.SIZE_AUTO, TRUE, FALSE, 1)

	/* set up our signal trappers... */
	dw.Signal_connect(tree, dw.SIGNAL_ITEM_CONTEXT, dw.SIGNAL_FUNC(item_context_cb), dw.HANDLE_TO_POINTER(tree_status))
	dw.Signal_connect(tree, dw.SIGNAL_ITEM_SELECT, dw.SIGNAL_FUNC(item_select_cb), dw.HANDLE_TO_POINTER(tree_status))

	t1 := dw.Tree_insert(tree, "tree folder 1", foldericon, dw.NOHTREEITEM, dw.POINTER(uintptr(1)))
	t2 := dw.Tree_insert(tree, "tree folder 2", foldericon, dw.NOHTREEITEM, dw.POINTER(uintptr(2)))
	dw.Tree_insert(tree, "tree file 1", fileicon, t1, dw.POINTER(uintptr(3)))
	dw.Tree_insert(tree, "tree file 2", fileicon, t1, dw.POINTER(uintptr(4)))
	dw.Tree_insert(tree, "tree file 3", fileicon, t2, dw.POINTER(uintptr(5)))
	dw.Tree_insert(tree, "tree file 4", fileicon, t2, dw.POINTER(uintptr(6)))
	dw.Tree_item_change(tree, t1, "tree folder 1", foldericon)
	dw.Tree_item_change(tree, t2, "tree folder 2", foldericon)
	dw.Tree_item_set_data(tree, t2, dw.POINTER(uintptr(100)))
	fmt.Printf("t1 title \"%s\" data %d t2 data %d\n", dw.Tree_get_title(tree, t1), uintptr(dw.Tree_item_get_data(tree, t1)), uintptr(dw.Tree_item_get_data(tree, t2)))
}

// Page 4
func container_add() {
	var z int
	titles := []string{"Type", "Size", "Time", "Date"}
	flags := []uint{dw.CFA_BITMAPORICON | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
		dw.CFA_ULONG | dw.CFA_RIGHT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
		dw.CFA_TIME | dw.CFA_CENTER | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR,
		dw.CFA_DATE | dw.CFA_LEFT | dw.CFA_HORZSEPARATOR | dw.CFA_SEPARATOR}

	/* create a box to pack into the notebook page */
	containerbox := dw.Box_new(dw.HORZ, 2)
	dw.Box_pack_start(notebookbox4, containerbox, 500, 200, TRUE, TRUE, 0)

	/* Add a word wrap/font style box */
	hbox := dw.Box_new(dw.HORZ, 0)

	checkbox := dw.Checkbox_new("Word wrap", 0)
	dw.Box_pack_start(hbox, checkbox, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, TRUE, 1)
	text := dw.Text_new("Foreground:", 0)
	dw.Window_set_style(text, dw.DT_VCENTER, dw.DT_VCENTER)
	dw.Box_pack_start(hbox, text, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, TRUE, 1)
	mlefore := color_combobox()
	dw.Box_pack_start(hbox, mlefore, 150, dw.SIZE_AUTO, TRUE, FALSE, 1)
	text = dw.Text_new("Background:", 0)
	dw.Window_set_style(text, dw.DT_VCENTER, dw.DT_VCENTER)
	dw.Box_pack_start(hbox, text, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, TRUE, 1)
	mleback := color_combobox()
	dw.Box_pack_start(hbox, mleback, 150, dw.SIZE_AUTO, TRUE, FALSE, 1)
	dw.Checkbox_set(checkbox, TRUE)
	text = dw.Text_new("Font:", 0)
	dw.Window_set_style(text, dw.DT_VCENTER, dw.DT_VCENTER)
	dw.Box_pack_start(hbox, text, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, TRUE, 1)
	fontsize := dw.Spinbutton_new("9", 0)
	dw.Box_pack_start(hbox, fontsize, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, FALSE, 1)
	dw.Spinbutton_set_limits(fontsize, 100, 5)
	dw.Spinbutton_set_pos(fontsize, 9)
	fontname := dw.Combobox_new("Default", 0)
	dw.Listbox_append(fontname, "Default")
	dw.Listbox_append(fontname, "Arial")
	dw.Listbox_append(fontname, "Geneva")
	dw.Listbox_append(fontname, "Verdana")
	dw.Listbox_append(fontname, "Helvetica")
	dw.Listbox_append(fontname, "DejaVu Sans")
	dw.Listbox_append(fontname, "Times New Roman")
	dw.Listbox_append(fontname, "Times New Roman Bold")
	dw.Listbox_append(fontname, "Times New Roman Italic")
	dw.Listbox_append(fontname, "Times New Roman Bold Italic")
	dw.Box_pack_start(hbox, fontname, 150, dw.SIZE_AUTO, TRUE, FALSE, 1)
	dw.Box_pack_start(notebookbox4, hbox, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 1)

	dw.Window_set_data(hbox, "mlefore", dw.HANDLE_TO_POINTER(mlefore))
	dw.Window_set_data(hbox, "mleback", dw.HANDLE_TO_POINTER(mleback))
	dw.Window_set_data(hbox, "fontsize", dw.HANDLE_TO_POINTER(fontsize))
	dw.Window_set_data(hbox, "fontname", dw.HANDLE_TO_POINTER(fontname))

	/* now a container area under this box */
	container = dw.Container_new(100, TRUE)
	dw.Box_pack_start(notebookbox4, container, 500, 200, TRUE, FALSE, 1)

	/* and a status area to see whats going on */
	container_status := dw.Status_text_new("", 0)
	dw.Box_pack_start(notebookbox4, container_status, 100, dw.SIZE_AUTO, TRUE, FALSE, 1)

	dw.Filesystem_set_column_title(container, "Test")
	dw.Filesystem_setup(container, flags, titles)
	dw.Container_set_stripe(container, dw.CLR_DEFAULT, dw.CLR_DEFAULT)
	containerinfo := dw.Container_alloc(container, 3)

	for z = 0; z < 3; z++ {
		var thisicon dw.HICN = fileicon

		if z == 0 {
			thisicon = foldericon
		}
		fmt.Printf("Initial: container: %x containerinfo: %x icon: %x\n", uintptr(dw.HANDLE_TO_POINTER(container)),
			dw.HANDLE_TO_UINTPTR(containerinfo), uintptr(thisicon))
		dw.Filesystem_set_file(container, containerinfo, z, fmt.Sprintf("Filename %d", z+1), thisicon)
		dw.Filesystem_set_item_icon(container, containerinfo, 0, z, thisicon)
		dw.Filesystem_set_item_ulong(container, containerinfo, 1, z, uint(z*100))
		dw.Filesystem_set_item_time(container, containerinfo, 2, z, z+10, z+10, z+10)
		dw.Filesystem_set_item_date(container, containerinfo, 3, z, z+10, z+10, z+2000)
		dw.Container_set_row_title(containerinfo, z, fmt.Sprintf("We can now allocate from the stack: Item: %d", z))
		dw.Container_set_row_data(containerinfo, z, dw.POINTER(uintptr(z)))
	}
	dw.Container_insert(container, containerinfo, 3)

	containerinfo = dw.Container_alloc(container, 1)
	dw.Filesystem_set_file(container, containerinfo, 0, "Yikes", foldericon)
	dw.Filesystem_set_item_icon(container, containerinfo, 0, 0, foldericon)
	dw.Filesystem_set_item_ulong(container, containerinfo, 1, 0, 324)
	dw.Filesystem_set_item_time(container, containerinfo, 2, 0, z+10, z+10, z+10)
	dw.Filesystem_set_item_date(container, containerinfo, 3, 0, z+10, z+10, z+2000)
	dw.Container_set_row_title(containerinfo, 0, "Extra")

	dw.Container_insert(container, containerinfo, 1)
	dw.Container_optimize(container)

	container_mle = dw.Mle_new(111)
	dw.Box_pack_start(containerbox, container_mle, 500, 200, TRUE, TRUE, 0)

	mle_point = dw.Mle_import(container_mle, "", -1)
	mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]", mle_point), mle_point)
	mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]abczxydefijkl", mle_point), mle_point)
	dw.Mle_delete(container_mle, 9, 3)
	mle_point = dw.Mle_import(container_mle, "gh", 12)
	newpoint, _ := dw.Mle_get_size(container_mle)
	mle_point = newpoint
	mle_point = dw.Mle_import(container_mle, fmt.Sprintf("[%d]\r\n\r\n", mle_point), mle_point)
	dw.Mle_set_cursor(container_mle, mle_point)
	/* connect our event trappers... */
	dw.Signal_connect(container, dw.SIGNAL_ITEM_ENTER, dw.SIGNAL_FUNC(item_enter_cb), dw.HANDLE_TO_POINTER(container_status))
	dw.Signal_connect(container, dw.SIGNAL_ITEM_CONTEXT, dw.SIGNAL_FUNC(item_context_cb), dw.HANDLE_TO_POINTER(container_status))
	dw.Signal_connect(container, dw.SIGNAL_ITEM_SELECT, dw.SIGNAL_FUNC(container_select_cb), dw.HANDLE_TO_POINTER(container_status))
	dw.Signal_connect(container, dw.SIGNAL_COLUMN_CLICK, dw.SIGNAL_FUNC(column_click_cb), dw.HANDLE_TO_POINTER(container_status))
	dw.Signal_connect(checkbox, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(word_wrap_click_cb), dw.HANDLE_TO_POINTER(container_mle))
	dw.Signal_connect(mlefore, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(mle_color_cb), dw.HANDLE_TO_POINTER(hbox))
	dw.Signal_connect(mleback, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(mle_color_cb), dw.HANDLE_TO_POINTER(hbox))
	dw.Signal_connect(fontname, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(mle_fontname_cb), dw.HANDLE_TO_POINTER(hbox))
	dw.Signal_connect(fontsize, dw.SIGNAL_VALUE_CHANGED, dw.SIGNAL_FUNC(mle_fontsize_cb), dw.HANDLE_TO_POINTER(hbox))
}

// Page 5
func buttons_add() {
	var i int

	/* create a box to pack into the notebook page */
	buttonsbox = dw.Box_new(dw.VERT, 2)
	dw.Box_pack_start(notebookbox5, buttonsbox, 25, 200, TRUE, TRUE, 0)
	dw.Window_set_color(buttonsbox, dw.CLR_RED, dw.CLR_RED)

	calbox := dw.Box_new(dw.HORZ, 0)
	dw.Box_pack_start(notebookbox5, calbox, 500, 200, TRUE, TRUE, 1)
	cal = dw.Calendar_new(100)
	dw.Box_pack_start(calbox, cal, 180, 120, TRUE, TRUE, 0)
	/*
	   dw.Calendar_set_date(cal, 2001, 1, 1);
	*/
	/*
	 * Create our file toolbar boxes...
	 */
	buttonboxperm = dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(buttonsbox, buttonboxperm, 25, 0, FALSE, TRUE, 2)
	dw.Window_set_color(buttonboxperm, dw.CLR_WHITE, dw.CLR_WHITE)
	abutton1 := dw.Bitmapbutton_new_from_file("Top Button", 0, fmt.Sprintf("%s/%s", SRCROOT, FILE_ICON_NAME))
	dw.Box_pack_start(buttonboxperm, abutton1, 100, 30, FALSE, FALSE, 0)
	dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(button_callback), nil)
	dw.Box_pack_start(buttonboxperm, dw.NOHWND, 25, 5, FALSE, FALSE, 0)
	abutton2 := dw.Bitmapbutton_new_from_file("Bottom", 0, fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME))
	dw.Box_pack_start(buttonsbox, abutton2, 25, 25, FALSE, FALSE, 0)
	dw.Signal_connect(abutton2, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(button_callback), nil)
	dw.Window_set_bitmap(abutton2, 0, FILE_ICON_NAME)

	create_button(false)
	/* make a combobox */
	combox := dw.Box_new(dw.VERT, 2)
	dw.Box_pack_start(notebookbox5, combox, 25, 200, TRUE, FALSE, 0)
	combobox1 = dw.Combobox_new("fred", 0) /* no point in specifying an initial value */
	dw.Listbox_append(combobox1, "fred")
	dw.Box_pack_start(combox, combobox1, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)
	/*
	   dw_window_set_text( combobox, "initial value");
	*/
	dw.Signal_connect(combobox1, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(combobox_select_event_callback), nil)

	combobox2 = dw.Combobox_new("joe", 0) /* no point in specifying an initial value */
	dw.Box_pack_start(combox, combobox2, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)
	/*
	   dw_window_set_text( combobox, "initial value");
	*/
	dw.Signal_connect(combobox2, dw.SIGNAL_LIST_SELECT, dw.SIGNAL_FUNC(combobox_select_event_callback), nil)
	/* add LOTS of items */
	fmt.Printf("before appending 500 items to combobox using dw_listbox_list_append()\n")
	text := make([]string, 500)
	for i = 0; i < 500; i++ {
		text[i] = fmt.Sprintf("item %d", i)
	}
	dw.Listbox_list_append(combobox2, text)
	fmt.Printf("after appending 500 items to combobox\n")
	/* now insert a couple of items */
	dw.Listbox_insert(combobox2, "inserted item 2", 2)
	dw.Listbox_insert(combobox2, "inserted item 5", 5)
	/* make a spinbutton */
	spinbutton = dw.Spinbutton_new("", 0) /* no point in specifying text */
	dw.Box_pack_start(combox, spinbutton, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)
	dw.Spinbutton_set_limits(spinbutton, 100, 1)
	dw.Spinbutton_set_pos(spinbutton, 30)
	dw.Signal_connect(spinbutton, dw.SIGNAL_VALUE_CHANGED, dw.SIGNAL_FUNC(spinbutton_valuechanged_callback), nil)
	/* make a slider */
	slider = dw.Slider_new(FALSE, 11, 0) /* no point in specifying text */
	dw.Box_pack_start(combox, slider, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)
	dw.Signal_connect(slider, dw.SIGNAL_VALUE_CHANGED, dw.SIGNAL_FUNC(slider_valuechanged_callback), nil)
	/* make a percent */
	percent = dw.Percent_new(0)
	dw.Box_pack_start(combox, percent, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)
}

func create_button(redraw bool) {
	filetoolbarbox := dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(buttonboxperm, filetoolbarbox, 0, 0, TRUE, TRUE, 0)

	abutton1 := dw.Bitmapbutton_new_from_file("Empty image. Should be under Top button", 0, "junk")
	dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0)
	dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(change_color_red_callback), nil)
	dw.Box_pack_start(filetoolbarbox, dw.NOHWND, 25, 5, FALSE, FALSE, 0)

	abutton1 = dw.Bitmapbutton_new_from_file("A borderless bitmapbitton", 0, fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME))
	dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0)
	dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(change_color_yellow_callback), nil)
	dw.Box_pack_start(filetoolbarbox, dw.NOHWND, 25, 5, FALSE, FALSE, 0)
	dw.Window_set_style(abutton1, dw.BS_NOBORDER, dw.BS_NOBORDER)

	//abutton1 = dw.Bitmapbutton_new_from_data("A button from data", 0, folder_ico, 1718 );
	abutton1 = dw.Bitmapbutton_new_from_file("A button from data", 0, "junk")
	dw.Box_pack_start(filetoolbarbox, abutton1, 25, 25, FALSE, FALSE, 0)
	dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(percent_button_box_callback), nil)
	dw.Box_pack_start(filetoolbarbox, dw.NOHWND, 25, 5, FALSE, FALSE, 0)
	if redraw == true {
		dw.Window_redraw(filetoolbarbox)
		dw.Window_redraw(mainwindow)
	}
}

// Page 7
func html_add() {
	rawhtml := dw.Html_new(1001)
	if rawhtml.GetHandle() != 0 {
		hbox := dw.Box_new(dw.HORZ, 0)
		javascript := dw.Combobox_new("", 0)

		dw.Listbox_append(javascript, "window.scrollTo(0,500);")
		dw.Listbox_append(javascript, "window.document.title;")
		dw.Listbox_append(javascript, "window.navigator.userAgent;")

		dw.Box_pack_start(notebookbox7, rawhtml, 0, 100, TRUE, FALSE, 0)
		dw.Html_raw(rawhtml, "<html><body><center><h1>dwtest</h1></center></body></html>")
		html := dw.Html_new(1002)

		dw.Box_pack_start(notebookbox7, hbox, 0, 0, TRUE, FALSE, 0)

		/* Add navigation buttons */
		item := dw.Button_new("Back", 0)
		dw.Box_pack_start(hbox, item, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, FALSE, 0)
		dw.Signal_connect(item, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(web_back_clicked), dw.HANDLE_TO_POINTER(html))

		item = dw.Button_new("Forward", 0)
		dw.Box_pack_start(hbox, item, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, FALSE, 0)
		dw.Signal_connect(item, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(web_forward_clicked), dw.HANDLE_TO_POINTER(html))

		/* Put in some extra space */
		dw.Box_pack_start(hbox, dw.NOHWND, 5, 1, FALSE, FALSE, 0)

		item = dw.Button_new("Reload", 0)
		dw.Box_pack_start(hbox, item, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, FALSE, 0)
		dw.Signal_connect(item, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(web_reload_clicked), dw.HANDLE_TO_POINTER(html))

		/* Put in some extra space */
		dw.Box_pack_start(hbox, dw.NOHWND, 5, 1, FALSE, FALSE, 0)
		dw.Box_pack_start(hbox, javascript, dw.SIZE_AUTO, dw.SIZE_AUTO, TRUE, FALSE, 0)

		item = dw.Button_new("Run", 0)
		dw.Window_set_data(item, "javascript", dw.HANDLE_TO_POINTER(javascript))
		dw.Box_pack_start(hbox, item, dw.SIZE_AUTO, dw.SIZE_AUTO, FALSE, FALSE, 0)
		dw.Signal_connect(item, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(web_run_clicked), dw.HANDLE_TO_POINTER(html))
		dw.Window_click_default(javascript, item)

		dw.Box_pack_start(notebookbox7, html, 0, 100, TRUE, TRUE, 0)
		dw.Html_url(html, "https://dbsoft.org/dw_help.php")
		htmlstatus := dw.Status_text_new("HTML status loading...", 0)
		dw.Box_pack_start(notebookbox7, htmlstatus, 100, dw.SIZE_AUTO, TRUE, FALSE, 1)
		dw.Signal_connect(html, dw.SIGNAL_HTML_CHANGED, dw.SIGNAL_FUNC(web_html_changed), dw.HANDLE_TO_POINTER(htmlstatus))
		dw.Signal_connect(html, dw.SIGNAL_HTML_RESULT, dw.SIGNAL_FUNC(web_html_result), dw.HANDLE_TO_POINTER(javascript))
	} else {
		label := dw.Text_new("HTML widget not available.", 0)
		dw.Box_pack_start(notebookbox7, label, 0, 100, TRUE, TRUE, 0)
	}
}

// Page 8
func scrollbox_add() {
	var i int

	/* create a box to pack into the notebook page */
	scrollbox = dw.Scrollbox_new(dw.VERT, 0)
	dw.Box_pack_start(notebookbox8, scrollbox, 0, 0, TRUE, TRUE, 1)

	abutton1 := dw.Button_new("Show Adjustments", 0)
	dw.Box_pack_start(scrollbox, abutton1, dw.SIZE_AUTO, 30, FALSE, FALSE, 0)
	dw.Signal_connect(abutton1, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(scrollbox_button_callback), nil)

	for i = 0; i < MAX_WIDGETS; i++ {
		tmpbox := dw.Box_new(dw.HORZ, 0)
		dw.Box_pack_start(scrollbox, tmpbox, 0, 24, TRUE, FALSE, 2)
		label := dw.Text_new(fmt.Sprintf("Label %d", i), 0)
		dw.Box_pack_start(tmpbox, label, 0, 20, TRUE, FALSE, 0)
		item := dw.Entryfield_new(fmt.Sprintf("Entry %d", i), uint(i))
		dw.Box_pack_start(tmpbox, item, 0, 20, TRUE, FALSE, 0)
	}
}

// Page 9
func update_mle(text string, lock int) {
	/* Protect pos from being changed by different threads */
	if lock != 0 {
		dw.Mutex_lock(mutex)
	}
	mlepos = dw.Mle_import(threadmle, text, mlepos)
	dw.Mle_set_cursor(threadmle, mlepos)
	if lock != 0 {
		dw.Mutex_unlock(mutex)
	}
}

func thread_add() {
	/* create a box to pack into the notebook page */
	tmpbox := dw.Box_new(dw.VERT, 0)
	dw.Box_pack_start(notebookbox9, tmpbox, 0, 0, dw.TRUE, dw.TRUE, 1)

	startbutton = dw.Button_new("Start Threads", 0)
	dw.Box_pack_start(tmpbox, startbutton, dw.SIZE_AUTO, 30, dw.FALSE, dw.FALSE, 0)
	dw.Signal_connect(startbutton, dw.SIGNAL_CLICKED, dw.SIGNAL_FUNC(start_threads_button_callback), nil)

	/* Create the base threading components */
	threadmle = dw.Mle_new(0)
	dw.Box_pack_start(tmpbox, threadmle, 1, 1, TRUE, TRUE, 0)
	mutex = dw.Mutex_new()
	workevent = dw.Event_new()
}

func main() {
	DWFeatureList := []string{
		"Supports the HTML Widget",
		"Supports the DW_SIGNAL_HTML_RESULT callback",
		"Supports custom window border sizes",
		"Supports window frame transparency",
		"Supports Dark Mode user interface",
		"Supports auto completion in Multi-line Edit boxes",
		"Supports word wrapping in Multi-line Edit boxes",
		"Supports striped line display in container widgets",
		"Supports Multiple Document Interface window frame",
		"Supports status text area on notebook/tabbed controls",
		"Supports sending system notifications",
		"Supports UTF8 encoded Unicode text",
		"Supports Rich Edit based MLE control (Windows)",
		"Supports icons in the taskbar or similar system widget",
		"Supports the Tree Widget",
		"Supports arbitrary window placement"}

	/* Pick an approriate font for our platform */
	if runtime.GOOS == "windows" {
		FIXEDFONT = "10.Lucida Console"
	} else if runtime.GOOS == "darwin" {
		FIXEDFONT = "9.Monaco"
	}

	/* Locate the source root of the package */
	pkg, err := build.Import("dwtest", "", build.FindOnly)
	if err == nil && len(pkg.SrcRoot) > 0 {
		SRCROOT = fmt.Sprintf("%s/dwtest", pkg.SrcRoot)
	}

	/* Setup the Application ID for sending notifications */
	dw.App_id_set("org.dbsoft.dwindows.dwtest", "Dynamic Windows Test")

	/* Enable full dark mode on platforms that support it */
	if os.Getenv("DW_DARK_MODE") != "" {
		dw.Feature_set(dw.FEATURE_DARK_MODE, dw.DARK_MODE_FULL)
	}

	/* Initialize the Dynamic Windows engine */
	dw.Init(dw.TRUE)

	/* Test all the features and display the results */
	for feat := 0; feat < dw.FEATURE_MAX && feat < len(DWFeatureList); feat++ {
		result := dw.Feature_get(feat)
		status := "Unsupported"

		if result == 0 {
			status = "Disabled"
		} else if result > 0 {
			status = "Enabled"
		}

		fmt.Printf("%s: %s (%d)\n", DWFeatureList[feat], status, result)
	}

	/* Create our window */
	mainwindow = dw.Window_new(dw.DESKTOP, "dwindows test UTF8 中国語 (繁体) cañón", dw.FCF_SYSMENU|dw.FCF_TITLEBAR|dw.FCF_TASKLIST|dw.FCF_DLGBORDER|dw.FCF_SIZEBORDER|dw.FCF_MINMAX)

	menu_add()

	notebookbox := dw.Box_new(dw.VERT, 5)
	dw.Box_pack_start(mainwindow, notebookbox, 0, 0, dw.TRUE, dw.TRUE, 0)

	foldericon = dw.Icon_load_from_file(FOLDER_ICON_NAME)
	if foldericon == dw.NOHICN && len(SRCROOT) > 0 {
		foldericon = dw.Icon_load_from_file(fmt.Sprintf("%s/%s", SRCROOT, FOLDER_ICON_NAME))
	}
	fileicon = dw.Icon_load_from_file(FILE_ICON_NAME)
	if fileicon == dw.NOHICN && len(SRCROOT) > 0 {
		fileicon = dw.Icon_load_from_file(fmt.Sprintf("%s/%s", SRCROOT, FILE_ICON_NAME))
	}
	notebook := dw.Notebook_new(1, dw.TRUE)
	dw.Box_pack_start(notebookbox, notebook, 100, 100, dw.TRUE, dw.TRUE, 0)
	dw.Signal_connect(notebook, dw.SIGNAL_SWITCH_PAGE, dw.SIGNAL_FUNC(switch_page_callback), nil)

	notebookbox1 = dw.Box_new(dw.VERT, 5)
	notebookpage1 := dw.Notebook_page_new(notebook, 0, dw.TRUE)
	dw.Notebook_pack(notebook, notebookpage1, notebookbox1)
	dw.Notebook_page_set_text(notebook, notebookpage1, "buttons and entry")
	archive_add()

	notebookbox2 = dw.Box_new(dw.VERT, 5)
	notebookpage2 := dw.Notebook_page_new(notebook, 1, dw.FALSE)
	dw.Notebook_pack(notebook, notebookpage2, notebookbox2)
	dw.Notebook_page_set_text(notebook, notebookpage2, "render")
	text_add()

	notebookbox3 = dw.Box_new(dw.VERT, 5)
	notebookpage3 := dw.Notebook_page_new(notebook, 1, dw.FALSE)
	dw.Notebook_pack(notebook, notebookpage3, notebookbox3)
	dw.Notebook_page_set_text(notebook, notebookpage3, "tree")
	tree_add()

	notebookbox4 = dw.Box_new(dw.VERT, 5)
	notebookpage4 := dw.Notebook_page_new(notebook, 1, FALSE)
	dw.Notebook_pack(notebook, notebookpage4, notebookbox4)
	dw.Notebook_page_set_text(notebook, notebookpage4, "container")
	container_add()

	notebookbox5 = dw.Box_new(dw.VERT, 5)
	notebookpage5 := dw.Notebook_page_new(notebook, 1, FALSE)
	dw.Notebook_pack(notebook, notebookpage5, notebookbox5)
	dw.Notebook_page_set_text(notebook, notebookpage5, "buttons")
	buttons_add()

	/* DEPRECATED
	   notebookbox6 = dw.Box_new(dw.VERT, 5);
	   notebookpage6 := dw.Notebook_page_new( notebook, 1, FALSE );
	   dw.Notebook_pack(notebook, notebookpage6, notebookbox6);
	   dw.Notebook_page_set_text(notebook, notebookpage6, "mdi");
	   mdi_add();
	*/

	notebookbox7 = dw.Box_new(dw.VERT, 6)
	notebookpage7 := dw.Notebook_page_new(notebook, 1, FALSE)
	dw.Notebook_pack(notebook, notebookpage7, notebookbox7)
	dw.Notebook_page_set_text(notebook, notebookpage7, "html")
	html_add()

	notebookbox8 = dw.Box_new(dw.VERT, 7)
	notebookpage8 := dw.Notebook_page_new(notebook, 1, FALSE)
	dw.Notebook_pack(notebook, notebookpage8, notebookbox8)
	dw.Notebook_page_set_text(notebook, notebookpage8, "scrollbox")
	scrollbox_add()

	notebookbox9 = dw.Box_new(dw.VERT, 8)
	notebookpage9 := dw.Notebook_page_new(notebook, 1, FALSE)
	dw.Notebook_pack(notebook, notebookpage9, notebookbox9)
	dw.Notebook_page_set_text(notebook, notebookpage9, "thread/event")
	thread_add()

	/* Set the default field */
	dw.Window_default(mainwindow, copypastefield)

	dw.Signal_connect(mainwindow, dw.SIGNAL_DELETE, dw.SIGNAL_FUNC(exit_callback), dw.HANDLE_TO_POINTER(mainwindow))
	/*
	 * The following is a special case handler for the Mac and other platforms which contain
	 * an application object which can be closed.  It function identically to a window delete/close
	 * request except it applies to the entire application not an individual window. If it is not
	 * handled or you allow the default handler to take place the entire application will close.
	 * On platforms which do not have an application object this line will be ignored.
	 */
	dw.Signal_connect(dw.DESKTOP, dw.SIGNAL_DELETE, dw.SIGNAL_FUNC(exit_callback), dw.HANDLE_TO_POINTER(mainwindow))
	timerid = dw.Timer_connect(2000, dw.SIGNAL_FUNC(timer_callback), nil)
	dw.Window_set_size(mainwindow, 640, 550)
	dw.Window_show(mainwindow)

	/* Now that the window is created and shown...
	 * run the main loop until we get dw_main_quit()
	 */
	dw.Main()

	/* Now that the loop is done we can cleanup */
	dw.Taskbar_delete(textbox1, fileicon)
	dw.Window_destroy(mainwindow)

	fmt.Printf("dwtest exiting...\n")
	/* Call dw.Shutdown() to shutdown the Dynamic Windows engine */
	dw.Shutdown()
}
