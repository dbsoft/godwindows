package main

import (
   "dw"
   "dwib"
   "fmt"
   "go/build"
)

// Global variables
const (
   FALSE int = iota
   TRUE
)

var APP_NAME = "DWIB Example"
var SRCROOT string

/* Handle exiting the application */
func exit_handler() int {
    if dw.MessageBox(APP_NAME, dw.MB_YESNO | dw.MB_QUESTION, "Are you sure you want to exit?") == dw.MB_RETURN_YES {
        /* Exit the application cleanly */
        dw.MainQuit();
    }
    return TRUE;
}

var exit_handler_func = exit_handler;

func main() {
   /* Locate the source root of the package */
   pkg, err := build.Import("dwibtest", "", build.FindOnly);
   if err == nil && len(pkg.SrcRoot) > 0 {
      SRCROOT = fmt.Sprintf("%s/dwibtest", pkg.SrcRoot);
   }
   
    /* Initialize Dynamic Windows */
    dw.Init(TRUE);

    /* Load the interface XML file */
    handle := dwib.Open("example.xml");
    if handle == nil && len(SRCROOT) > 0 {
       handle = dwib.Open(fmt.Sprintf("%s/example.xml", SRCROOT));
    }

    /* Show an error if it fails to load */
    if handle == nil {
        dw.MessageBox(APP_NAME, dw.MB_OK | dw.MB_ERROR, "Unable to load the interface XML.");
        return;
    }

    /* Create the loading window... */
    window := dwib.Load(handle, "Test");
    dwib.Show(window);

    /* Connect the signal handlers */
    window.ConnectDelete(func(window dw.HWND) int { return exit_handler(); });
    /* Handler for Mac application menu Quit */
    dw.DESKTOP.ConnectDelete(func(window dw.HWND) int { return exit_handler(); });
    quitmenu := dw.HANDLE_TO_HMENUITEM(dwib.GetHandle(window, "quitmenu"));
    quitmenu.ConnectClicked(func(window dw.HMENUITEM) int { return exit_handler(); });

    dw.Main();

    /* Destroy the main window */
    window.Destroy();
    /* Close the Interface Builder XML */
    dwib.Close(handle);

    /* Call dw.Shutdown() to shutdown the Dynamic Windows engine */
    dw.Shutdown();
    return;
}

