#include <dwib.h>
#include <stdlib.h>
#include <string.h>

static uintptr_t goib_load(uintptr_t handle, char *name)
{
   return (uintptr_t)dwib_load((DWIB)handle, name);
}

static int goib_load_at_index(uintptr_t handle, char *name, char *dataname, uintptr_t window, uintptr_t box, int index)
{
   return dwib_load_at_index((DWIB)handle, name, dataname, (HWND)window, (HWND)box, index);
}

static void goib_show(uintptr_t window)
{
   dwib_show((HWND)window);
}

static uintptr_t goib_open(char *filename)
{
   return (uintptr_t)dwib_open(filename);
}

static void goib_close(uintptr_t handle)
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

static uintptr_t goib_window_get_handle(uintptr_t handle, char *dataname)
{
   return (uintptr_t)dwib_window_get_handle((HWND)handle, dataname);
}
