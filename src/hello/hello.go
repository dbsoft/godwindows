package main

import (
   "dw"
)

// Global variables
var APP_NAME = "DW Hello World Example"

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
    window := dw.WindowNew(dw.DESKTOP, APP_NAME, dw.FCF_SYSMENU | dw.FCF_TITLEBAR | dw.FCF_TASKLIST | dw.FCF_DLGBORDER | dw.FCF_SIZEBORDER | dw.FCF_MINMAX);

    label := dw.TextNew("Hello, 世界", 0);
    window.PackStart(label, 0, 0, dw.TRUE, dw.TRUE, 0);
    
    /* Connect the signal handlers */
    window.ConnectDelete(func(window dw.HWND) int { return exit_handler(window, message); });

    /* Set the size and show the window */
    window.SetSize(640, 550);
    window.Show();
    
    dw.Main();

    /* Destroy the main window */
    window.Destroy();

    /* Call dw.Shutdown() to shutdown the Dynamic Windows engine */
    dw.Shutdown();
    return;
}

