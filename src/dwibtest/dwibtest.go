package main

import (
   "unsafe"
   "dw"
   "dwib"
   "fmt"
)

// Global variables
const (
   FALSE int = iota
   TRUE
)

var APP_NAME = "DWIB Example"

/* Handle exiting the application */
func exit_handler(win dw.HWND, data unsafe.Pointer) int {
    if(dw.Messagebox(APP_NAME, dw.MB_YESNO | dw.MB_QUESTION, "Are you sure you want to exit"))
    {
        /* Exit the application cleanly */
        dw.Main_quit();
    }
    return TRUE;
}

var exit_handler_func = exit_handler;

func main() {
    /* Initialize Dynamic Windows */
    dw.Init(TRUE);

    /* Load the interface XML file */
    handle := dwib.Open("example.xml");

    /* Show an error if it fails to load */
    if handle == nil {
        dw.Messagebox(APP_NAME, dw.MB_OK | dw.MB_ERROR, "Unable to load the interface XML.");
        return;
    }

    /* Create the loading window... */
    window := dwib.Load(handle, "Test");
    dwib.Show(window);

    /* Connect the signal handlers */
    dw.Signal_connect(window, dw.SIGNAL_DELETE, unsafe.Pointer(&exit_handler_func), nil);
    /* Handler for Mac application menu Quit */
    dw.Signal_connect(dw.DESKTOP, dw.SIGNAL_DELETE, unsafe.Pointer(&exit_handler_func), nil);
    dw.Signal_connect(dwib.Window_get_handle(window, "quitmenu"), dw.SIGNAL_CLICKED, unsafe.Pointer(&exit_handler_func), nil);

    dw.Main();

    /* Destroy the main window */
    dw.Window_destroy(window);
    /* Close the Interface Builder XML */
    dwib.Close(handle);

    return;
}

