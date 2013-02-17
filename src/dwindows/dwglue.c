#include <dw.h>
#include <stdlib.h>
#include <string.h>

static int go_init(int newthread)
{
   int argc = 0;
   char **argv = NULL;
   
   return dw_init(newthread, argc, argv);
}

static int go_messagebox(char *title, int flags, char *message)
{
   return dw_messagebox(title, flags, message);
}

static void *go_window_new(void *owner, char *title, unsigned long flags)
{
   return (void *)dw_window_new((HWND)owner, title, flags);
}

static int go_window_show(void *handle)
{
   return dw_window_show((HWND)handle);
}

static void go_window_set_pos(void *handle, long x, long y)
{
   dw_window_set_pos((HWND)handle, x, y);
}

static void go_window_set_pos_size(void *handle, long x, long y, unsigned long width, unsigned long height)
{
   dw_window_set_pos_size((HWND)handle, x, y, width, height);
}

static void go_window_set_size(void *handle, unsigned long width, unsigned long height)
{
   dw_window_set_size((HWND)handle, width, height);
}

static int go_window_set_color(void *handle, unsigned long fore, unsigned long back)
{
   return dw_window_set_color((HWND)handle, fore, back);
}

static void go_window_set_style(void *handle, unsigned long style, unsigned long mask)
{
   dw_window_set_style((HWND)handle, style, mask);
}

static void go_window_click_default(void *window, void *next)
{
   dw_window_click_default((HWND)window, (HWND)next);
}

static void go_window_default(void *window, void *defaultitem)
{
   dw_window_default((HWND)window, (HWND)defaultitem);
}

static void *go_box_new(int type, int pad)
{
   return (void *)dw_box_new(type, pad);
}

static void go_box_pack_at_index(void *box, void *item, int index, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_at_index((HWND)box, (HWND)item, index, width, height, hsize, vsize, pad);
}

static void go_box_pack_end(void *box, void *item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_end((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

static void go_box_pack_start(void *box, void *item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_start((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

static int go_box_unpack(void *handle)
{
   return dw_box_unpack((HWND)handle);
}

static void *go_box_unpack_at_index(void *handle, int index)
{
   return (void *)dw_box_unpack_at_index((HWND)handle, index);
}

static void *go_text_new(char *text, unsigned long id)
{
   return (void *)dw_text_new(text, id);
}

static void *go_entryfield_new(char *text, unsigned long id)
{
   return (void *)dw_entryfield_new(text, id);
}

static void *go_entryfield_password_new(char *text, unsigned long id)
{
   return (void *)dw_entryfield_password_new(text, id);
}

static void go_entryfield_set_limit(void *handle, int limit)
{
   dw_entryfield_set_limit((HWND)handle, limit);
}

static void *go_button_new(char *text, unsigned long id)
{
   return (void *)dw_button_new(text, id);
}

extern int go_int_callback_basic(void *pfunc, void* window, void *data);
extern int go_int_callback_configure(void *pfunc, void* window, int width, int height, void *data);
extern int go_int_callback_keypress(void *pfunc, void *window, char ch, int vk, int state, void *data, char *utf8);
extern int go_int_callback_mouse(void *pfunc, void* window, int x, int y, int mask, void *data);
extern int go_int_callback_expose(void *pfunc, void* window, int x, int y, int width, int height, void *data);
extern int go_int_callback_string(void *pfunc, void* window, char *str, void *data);
extern int go_int_callback_item_context(void *pfunc, void *window, char *text, int x, int y, void *data, void *itemdata);
extern int go_int_callback_item_select(void *pfunc, void *window, void *item, char *text, void *data, void *itemdata);
extern int go_int_callback_numeric(void *pfunc, void* window, int val, void *data);
extern int go_int_callback_ulong(void *pfunc, void* window, unsigned long val, void *data);
extern int go_int_callback_tree(void *pfunc, void* window, void *item, void *data);

static int DWSIGNAL go_callback_basic(HWND window, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_basic(param[0], (void *)window, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_configure(HWND window, int width, int height, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_configure(param[0], (void *)window, width, height, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_keypress(HWND window, char ch, int vk, int state, void *data, char *utf8)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_keypress(param[0], (void *)window, ch, vk, state, param[1], utf8);
   }
   return 0;
}

static int DWSIGNAL go_callback_mouse(HWND window, int x, int y, int mask, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_mouse(param[0], (void *)window, x, y, mask, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_expose(HWND window,  DWExpose *exp, void *data)
{
   if(data && exp)
   {
      void **param = (void **)data;
      return go_int_callback_expose(param[0], (void *)window, exp->x, exp->y, exp->width, exp->height, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_string(HWND window, char *str, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_string(param[0], (void *)window, str, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_item_context(HWND window, char *text, int x, int y, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_context(param[0], (void *)window, text, x, y, param[1], itemdata);
   }
   return 0;
}

static int DWSIGNAL go_callback_item_select(HWND window, HTREEITEM item, char *text, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_select(param[0], (void *)window, (void *)item, text, param[1], itemdata);
   }
   return 0;
}

static int DWSIGNAL go_callback_numeric(HWND window, int val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_numeric(param[0], (void *)window, val, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_ulong(HWND window, unsigned long val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_ulong(param[0], (void *)window, val, param[1]);
   }
   return 0;
}

static int DWSIGNAL go_callback_tree(HWND window, HTREEITEM tree, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_tree(param[0], (void *)window, (void *)tree, param[1]);
   }
   return 0;
}

static void go_signal_connect(void *window, char *signame, void *sigfunc, void *data)
{
   void **param = malloc(sizeof(void *) * 2);
   void *func = (void *)go_callback_basic;
   
   if(param)
   {
      param[0] = sigfunc;
      param[1] = data;
      
      if(strcmp(signame, DW_SIGNAL_CONFIGURE) == 0)
      {
         func = (void *)go_callback_configure;
      }
      else if(strcmp(signame, DW_SIGNAL_KEY_PRESS) == 0)
      {
         func = (void *)go_callback_keypress;
      }
      else if(strcmp(signame, DW_SIGNAL_BUTTON_PRESS) == 0 ||
              strcmp(signame, DW_SIGNAL_BUTTON_RELEASE) == 0 ||
              strcmp(signame, DW_SIGNAL_MOTION_NOTIFY) == 0)
      {
         func = (void *)go_callback_mouse;
      }
      else if(strcmp(signame, DW_SIGNAL_EXPOSE) == 0)
      {
         func = (void *)go_callback_expose;
      }
      else if(strcmp(signame, DW_SIGNAL_ITEM_ENTER) == 0)
      {
         func = (void *)go_callback_string;
      }
      else if(strcmp(signame, DW_SIGNAL_ITEM_CONTEXT) == 0)
      {
         func = (void *)go_callback_item_context;
      }
      else if(strcmp(signame, DW_SIGNAL_ITEM_SELECT) == 0)
      {
         func = (void *)go_callback_item_select;
      }
      else if(strcmp(signame, DW_SIGNAL_LIST_SELECT) == 0 ||
              strcmp(signame, DW_SIGNAL_VALUE_CHANGED) == 0 ||
              strcmp(signame, DW_SIGNAL_COLUMN_CLICK) == 0)
      {
         func = (void *)go_callback_numeric;
      }
      else if(strcmp(signame, DW_SIGNAL_SWITCH_PAGE) == 0)
      {
         func = (void *)go_callback_ulong;
      }
      else if(strcmp(signame, DW_SIGNAL_TREE_EXPAND) == 0)
      {
         func = (void *)go_callback_tree;
      }
      
      dw_signal_connect((HWND)window, signame, func, param);
   }
}
