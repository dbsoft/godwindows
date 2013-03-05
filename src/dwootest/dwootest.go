package main

import (
   "dw"
)

// Global variables
var APP_NAME = "DW Object Oriented Example"

/* Handle exiting the application */
func exit_handler(win dw.HWND, message string) int {
    if dw.Messagebox(APP_NAME, dw.MB_YESNO | dw.MB_QUESTION, message) == dw.MB_RETURN_YES {
        /* Exit the application cleanly */
        dw.Main_quit();
    }
    return dw.TRUE;
}

func main() {
    var message = "Are you sure you want to exit?";
    
    /* Initialize Dynamic Windows */
    dw.Init(dw.TRUE);

    /* Create our window */
    window := dw.Window_new(dw.DESKTOP, APP_NAME, dw.FCF_SYSMENU | dw.FCF_TITLEBAR | dw.FCF_TASKLIST | dw.FCF_DLGBORDER | dw.FCF_SIZEBORDER | dw.FCF_MINMAX);

    label := dw.Text_new("Hello, World", 0);
    dw.Box_pack_start(window, label, 0, 0, dw.TRUE, dw.TRUE, 0);
    
    /* Connect the signal handlers */
    window.Delete(func(window dw.HWND, data dw.POINTER) int { return exit_handler(window, message); });

    dw.Window_set_size(window, 640, 550);
    dw.Window_show(window);
    
    dw.Main();

    /* Destroy the main window */
    dw.Window_destroy(window);

    /* Call dw.Shutdown() to shutdown the Dynamic Windows engine */
    dw.Shutdown();
    return;
}

