#include <dw.h>
#include <stdlib.h>
#include <string.h>

/* Go to C Glue function prototypes */
extern int go_init(int newthread);
extern int go_messagebox(char *title, int flags, char *message);
extern void *go_window_new(void *owner, char *title, unsigned long flags);
extern int go_window_show(void *handle);
extern void go_window_set_pos(void *handle, long x, long y);
extern void go_window_set_pos_size(void *handle, long x, long y, unsigned long width, unsigned long height);
extern void go_window_set_size(void *handle, unsigned long width, unsigned long height);
extern int go_window_set_color(void *handle, unsigned long fore, unsigned long back);
extern void go_window_set_style(void *handle, unsigned long style, unsigned long mask);
extern void go_window_click_default(void *window, void *next);
extern void go_window_default(void *window, void *defaultitem);
extern void *go_box_new(int type, int pad);
extern void go_box_pack_at_index(void *box, void *item, int index, int width, int height, int hsize, int vsize, int pad);
extern void go_box_pack_end(void *box, void *item, int width, int height, int hsize, int vsize, int pad);
extern void go_box_pack_start(void *box, void *item, int width, int height, int hsize, int vsize, int pad);
extern int go_box_unpack(void *handle);
extern void *go_box_unpack_at_index(void *handle, int index);
extern void *go_text_new(char *text, unsigned long id);
extern void *go_entryfield_new(char *text, unsigned long id);
extern void *go_entryfield_password_new(char *text, unsigned long id);
extern void go_entryfield_set_limit(void *handle, int limit);
extern void *go_button_new(char *text, unsigned long id);
extern void go_signal_connect(void *window, char *signame, void *sigfunc, void *data);

/* C to Go Glue function prototypes */
extern int DWSIGNAL go_callback_basic(HWND window, void *data);
extern int DWSIGNAL go_callback_configure(HWND window, int width, int height, void *data);
extern int DWSIGNAL go_callback_keypress(HWND window, char ch, int vk, int state, void *data, char *utf8);
extern int DWSIGNAL go_callback_mouse(HWND window, int x, int y, int mask, void *data);
extern int DWSIGNAL go_callback_expose(HWND window,  DWExpose *exp, void *data);
extern int DWSIGNAL go_callback_string(HWND window, char *str, void *data);
extern int DWSIGNAL go_callback_item_context(HWND window, char *text, int x, int y, void *data, void *itemdata);
extern int DWSIGNAL go_callback_item_select(HWND window, HTREEITEM item, char *text, void *data, void *itemdata);
extern int DWSIGNAL go_callback_numeric(HWND window,  int val, void *data);
extern int DWSIGNAL go_callback_ulong(HWND window, unsigned long val, void *data);
extern int DWSIGNAL go_callback_tree(HWND window, HTREEITEM *tree, void *data);

/* Exported Go Glue function prototypes */
extern int go_int_callback_basic(void *pfunc, void* window, void *data);;
