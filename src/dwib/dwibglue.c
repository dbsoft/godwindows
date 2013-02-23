#include <dwib.h>
#include <stdlib.h>
#include <string.h>

static void *goib_load(void *handle, char *name)
{
   return (void *)dwib_load((DWIB)handle, name);
}

static int goib_load_at_index(void *handle, char *name, char *dataname, void *window, void *box, int index)
{
   return dwib_load_at_index((DWIB)handle, name, dataname, (HWND)window, (HWND)box, index);
}

static void goib_show(void *window)
{
   dwib_show((HWND)window);
}

static void *goib_open(char *filename)
{
   return (void *)dwib_open(filename);
}

static void goib_close(void *handle)
{
   dwib_close((DWIB)handle);
}

static int goib_image_root_set(char *path)
{
   return dwib_image_root_set(path);
}

static int goib_locale_set(char *loc)
{
   return dwib_locale_set(loc);
}

static void *goib_window_get_handle(void *handle, char *dataname)
{
   return (void *)dwib_window_get_handle((HWND)handle, dataname);
}
