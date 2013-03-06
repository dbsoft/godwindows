#include <dw.h>
#include <stdlib.h>
#include <string.h>

static int go_init(int newthread, int argc, char *argv[])
{
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

static int go_window_hide(void *handle)
{
   return dw_window_hide((HWND)handle);
}

static int go_window_lower(void *handle)
{
   return dw_window_lower((HWND)handle);
}

static int go_window_raise(void *handle)
{
   return dw_window_raise((HWND)handle);
}

static int go_window_minimize(void *handle)
{
   return dw_window_minimize((HWND)handle);
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

static int go_window_destroy(void *handle)
{
   return dw_window_destroy((HWND)handle);
}

static void go_window_disable(void *handle)
{
   dw_window_disable((HWND)handle);
}

static void go_window_enable(void *handle)
{
   dw_window_enable((HWND)handle);
}

static void *go_window_from_id(void *handle, int id)
{
   return dw_window_from_id((HWND)handle, id);
}

static void *go_window_get_data(void *handle, char *dataname)
{
   return dw_window_get_data((HWND)handle, dataname);
}

static char *go_window_get_font(void *handle)
{
   return dw_window_get_font((HWND)handle);
}

static int go_window_set_font(void *handle, char *fontname)
{
   return dw_window_set_font((HWND)handle, fontname);
}

static void go_window_get_pos_size(void *handle, long *x, long *y, unsigned long *width, unsigned long *height)
{
   dw_window_get_pos_size((HWND)handle, x, y, width, height);
}

static void go_window_get_preferred_size(void *handle, int *width, int *height)
{
   dw_window_get_preferred_size((HWND)handle, width, height);
}

static char *go_window_get_text(void *handle)
{
   return dw_window_get_text((HWND)handle);
}

static void go_window_set_text(void *handle, char *text)
{
   dw_window_set_text((HWND)handle, text);
}

static void go_window_set_tooltip(void *handle, char *bubbletext)
{
   dw_window_set_tooltip((HWND)handle, bubbletext);
}

static void go_window_redraw(void *handle)
{
   dw_window_redraw((HWND)handle);
}

static void go_window_capture(void *handle)
{
   dw_window_capture((HWND)handle);
}

static void go_window_set_bitmap(void *handle, unsigned long cid, char *filename)
{
   dw_window_set_bitmap((HWND)handle, cid, filename);
}

static int go_window_set_border(void *handle, int border)
{
   return dw_window_set_border((HWND)handle, border);
}

static void go_window_set_focus(void *handle)
{
   dw_window_set_focus((HWND)handle);
}

static void go_window_set_gravity(void *handle, int horz, int vert)
{
   dw_window_set_gravity((HWND)handle, horz, vert);
}

static void go_window_set_icon(void *handle, void *icon)
{
   dw_window_set_icon((HWND)handle, (HICN)icon);
}

static void go_window_set_pointer(void *handle, int cursortype)
{
   dw_window_set_pointer((HWND)handle, cursortype);
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

static void *go_status_text_new(char *text, unsigned long id)
{
   return (void *)dw_status_text_new(text, id);
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

static void *go_menu_new(unsigned long cid)
{
    return (void *)dw_menu_new(cid);
}

static void *go_menubar_new(void *location)
{
    return (void *)dw_menubar_new((HWND)location);
}

static void *go_menu_append_item(void *menu, char *title, unsigned long id, unsigned long flags, int end, int check, void *submenu)
{
    return dw_menu_append_item((HMENUI)menu, title, id, flags, end, check, submenu);
}

static int go_menu_delete_item(void *menu, unsigned long cid)
{
    return dw_menu_delete_item((HMENUI)menu, cid);
}

static void go_menu_destroy(void *menu)
{
    HMENUI thismenu = (HMENUI)menu;
    dw_menu_destroy(&thismenu);
}

static void go_menu_item_set_state(void *menu, unsigned long cid, unsigned long flags)
{
    dw_menu_item_set_state((HMENUI)menu, cid, flags);
}

static void go_menu_popup(void *menu, void *parent, int x, int y)
{
    HMENUI thismenu = (HMENUI)menu;
    dw_menu_popup(&thismenu, (HWND)parent, x, y);
}

static void *go_notebook_new(unsigned long cid, int top)
{
    return (void *)dw_notebook_new(cid, top);
}

static void go_notebook_pack(void *handle, unsigned long pageid, void *page)
{
    dw_notebook_pack((HWND)handle, pageid, (HWND)page);
}

static void go_notebook_page_destroy(void *handle, unsigned long pageid)
{
    dw_notebook_page_destroy((HWND)handle, (unsigned int)pageid);
}

static unsigned long go_notebook_page_get(void *handle)
{
    return dw_notebook_page_get((HWND)handle);
}

static unsigned long go_notebook_page_new(void *handle, unsigned long flags, int front)
{
    return dw_notebook_page_new((HWND)handle, flags, front);
}

static void go_notebook_page_set(void *handle, unsigned long pageid)
{
    dw_notebook_page_set((HWND)handle, (unsigned int)pageid);
}

static void go_notebook_page_set_text(void *handle, unsigned long pageid, char *text)
{
    dw_notebook_page_set_text((HWND)handle, pageid, text);
}

static void *go_icon_load_from_file(char *filename)
{
    return (void *)dw_icon_load_from_file(filename);
}

static void *go_icon_load(unsigned long module, unsigned long cid)
{
    return (void *)dw_icon_load(module, cid);
}

static void go_taskbar_delete(void *handle, void *icon)
{
    dw_taskbar_delete((HWND)handle, (HICN)icon);
}

static void go_taskbar_insert(void *handle, void *icon, char *bubbletext)
{
    dw_taskbar_insert((HWND)handle, (HICN)icon, bubbletext);
}

static void *go_combobox_new(char *text, unsigned long id)
{
   return (void *)dw_combobox_new(text, id);
}

static void *go_listbox_new(unsigned long id, int multi)
{
   return (void *)dw_listbox_new(id, multi);
}

static void go_listbox_append(void *handle, char *text)
{
    dw_listbox_append((HWND)handle, text);
}

static void go_listbox_list_append(void *handle, char **text, int count)
{
    dw_listbox_list_append((HWND)handle, text, count);
}

static void go_listbox_insert(void *handle, char *text, int pos)
{
    dw_listbox_insert((HWND)handle, text, pos);
}

static void go_listbox_clear(void *handle)
{
    dw_listbox_clear((HWND)handle);
}

static int go_listbox_count(void *handle)
{
    return dw_listbox_count((HWND)handle);
}

static void go_listbox_set_top(void *handle, int top)
{
    dw_listbox_set_top((HWND)handle, top);
}

static void go_listbox_select(void *handle, int index, int state)
{
    dw_listbox_select((HWND)handle, index, state);
}

static void go_listbox_delete(void *handle, int index)
{
    dw_listbox_delete((HWND)handle, index);
}

static void go_listbox_get_text(void *handle, int index, char *text, int length)
{
    dw_listbox_get_text((HWND)handle, index, text, length);
}

static void go_listbox_set_text(void *handle, int index, char *text)
{
    dw_listbox_set_text((HWND)handle, index, text);
}

static int go_listbox_selected(void *handle)
{
    return dw_listbox_selected((HWND)handle);
}

static int go_listbox_selected_multi(void *handle, int where)
{
    return dw_listbox_selected_multi((HWND)handle, where);
}

static void *go_spinbutton_new(char *text, unsigned long id)
{
    return (void *)dw_spinbutton_new(text, id);
}

static void go_spinbutton_set_pos(void *handle, long position)
{
    dw_spinbutton_set_pos((HWND)handle, position);
}
static void go_spinbutton_set_limits(void *handle, long upper, long lower)
{
    dw_spinbutton_set_limits((HWND)handle, upper, lower);
}

static long go_spinbutton_get_pos(void *handle)
{
    return dw_spinbutton_get_pos((HWND)handle);
}

static void *go_radiobutton_new(char *text, unsigned long id)
{
   return (void *)dw_radiobutton_new(text, id);
}

static void *go_checkbox_new(char *text, unsigned long id)
{
   return (void *)dw_checkbox_new(text, id);
}

static int go_checkbox_get(void *handle)
{
    return dw_checkbox_get((HWND)handle);
}

static void go_checkbox_set(void *handle, int value)
{
    return dw_checkbox_set((HWND)handle, value);
}

static void *go_percent_new(unsigned long id)
{
   return (void *)dw_percent_new(id);
}

static void go_percent_set_pos(void *handle, unsigned int position)
{
   dw_percent_set_pos((HWND)handle, position);
}

static void *go_slider_new(int vertical, int increments, unsigned long id)
{
   return (void *)dw_slider_new(vertical, increments, id);
}

static unsigned int go_slider_get_pos(void *handle)
{
   return dw_slider_get_pos((HWND)handle);
}

static void go_slider_set_pos(void *handle, unsigned int pos)
{
    dw_slider_set_pos((HWND)handle, pos);
}

static void *go_scrollbar_new(int vertical, unsigned long id)
{
   return (void *)dw_scrollbar_new(vertical, id);
}

static unsigned int go_scrollbar_get_pos(void *handle)
{
   return dw_scrollbar_get_pos((HWND)handle);
}

static void go_scrollbar_set_pos(void *handle, unsigned int pos)
{
    dw_scrollbar_set_pos((HWND)handle, pos);
}

static void go_scrollbar_set_range(void *handle, unsigned int range, unsigned int visible)
{
    dw_scrollbar_set_range((HWND)handle, range, visible);
}

static void *go_scrollbox_new(int type, int pad)
{
   return (void *)dw_scrollbox_new(type, pad);
}

static int go_scrollbox_get_pos(void *handle, int orient)
{
    return dw_scrollbox_get_pos((HWND)handle, orient);
}

static int go_scrollbox_get_range(void *handle, int orient)
{
    return dw_scrollbox_get_range((HWND)handle, orient);
}

static void *go_groupbox_new(int type, int pad, char *title)
{
   return (void *)dw_groupbox_new(type, pad, title);
}

static void *go_render_new(unsigned long id)
{
   return (void *)dw_render_new(id);
}

static void go_font_text_extents_get(void *handle, void *pixmap, char *text, int *width, int *height)
{
   dw_font_text_extents_get((HWND)handle, (HPIXMAP)pixmap, text, width, height);
}

static void *go_pixmap_new(void *handle, unsigned long width, unsigned long height, unsigned long depth) 
{
    return (void *)dw_pixmap_new((HWND)handle, width, height, (int)depth);
}

static void *go_pixmap_new_from_file(void *handle, char *filename) 
{
    return (void *)dw_pixmap_new_from_file((HWND)handle, filename);
}

static void *go_pixmap_grab(void *handle, unsigned long cid) 
{
    return (void *)dw_pixmap_grab((HWND)handle, cid);
}

static void go_pixmap_bitblt(void *dest, void *destp, int xdest, int ydest, int width, int height, void *src, void *srcp, int xsrc, int ysrc)
{
    dw_pixmap_bitblt((HWND)dest, (HPIXMAP)destp, xdest, ydest, width, height, (HWND)src, (HPIXMAP)srcp, xsrc, ysrc);
}

static int go_pixmap_stretch_bitblt(void *dest, void *destp, int xdest, int ydest, int width, int height, void *src, void *srcp, int xsrc, int ysrc, int srcwidth, int srcheight)
{
    return dw_pixmap_stretch_bitblt((HWND)dest, (HPIXMAP)destp, xdest, ydest, width, height, (HWND)src, (HPIXMAP)srcp, xsrc, ysrc, srcwidth, srcheight);
}

static void go_pixmap_set_transparent_color(void *pixmap, unsigned long color)
{
    dw_pixmap_set_transparent_color((HPIXMAP)pixmap, color);
}

static int go_pixmap_set_font(void *pixmap, char *fontname)
{
    return dw_pixmap_set_font((HPIXMAP)pixmap, fontname);
}

static void go_pixmap_destroy(void *pixmap)
{
    dw_pixmap_destroy((HPIXMAP)pixmap);
}

static int go_pixmap_width(void *pixmap)
{
    return (int)DW_PIXMAP_WIDTH(((HPIXMAP)pixmap));
}

static int go_pixmap_height(void *pixmap)
{
    return (int)DW_PIXMAP_HEIGHT(((HPIXMAP)pixmap));
}

static void go_draw_point(void *handle, void *pixmap, int x, int y)
{
    dw_draw_point((HWND)handle, (HPIXMAP)pixmap, x, y);
}

static void go_draw_line(void *handle, void *pixmap, int x1, int y1, int x2, int y2)
{
    dw_draw_line((HWND)handle, (HPIXMAP)pixmap, x1, y1, x2, y2);
}

static void go_draw_polygon(void *handle, void *pixmap, int fill, int count, int x[], int y[])
{
    dw_draw_polygon((HWND)handle, (HPIXMAP)pixmap, fill, count, x, y);
}

static void go_draw_rect(void *handle, void *pixmap, int fill, int x, int y, int width, int height)
{
    dw_draw_rect((HWND)handle, (HPIXMAP)pixmap, fill, x, y, width, height);
}

static void go_draw_arc(void *handle, void *pixmap, int flags, int xorigin, int yorigin, int x1, int y1, int x2, int y2)
{
    dw_draw_arc((HWND)handle, (HPIXMAP)pixmap, flags, xorigin, yorigin, x1, y1, x2, y2);
}

static void go_draw_text(void *handle, void *pixmap, int x, int y, char *text)
{
    dw_draw_text((HWND)handle, (HPIXMAP)pixmap, x, y, text);
}

static void *go_tree_new(unsigned long id)
{
   return (void *)dw_tree_new(id);
}

static void *go_tree_insert(void *handle, char *title, void *icon, void *parent, void *itemdata)
{
    return (void *)dw_tree_insert((HWND)handle, title, (HICN)icon, (HTREEITEM)parent, itemdata);
}

static void *go_tree_insert_after(void *handle, void *item, char *title, void *icon, void *parent, void *itemdata)
{
    return (void *)dw_tree_insert_after((HWND)handle, (HTREEITEM)item, title, (HICN)icon, (HTREEITEM)parent, itemdata);
}

static void go_tree_clear(void *handle)
{
    dw_tree_clear((HWND)handle);
}

static void go_tree_item_delete(void *handle, void *item)
{
    dw_tree_item_delete((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_change(void *handle, void *item, char *title, void *icon)
{
    dw_tree_item_change((HWND)handle, (HTREEITEM)item, title, (HICN)icon);
}

static void go_tree_item_expand(void *handle, void *item)
{
    dw_tree_item_expand((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_collapse(void *handle, void *item)
{
    dw_tree_item_collapse((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_select(void *handle, void *item)
{
    dw_tree_item_select((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_set_data(void *handle, void *item, void *itemdata)
{
    dw_tree_item_set_data((HWND)handle, (HTREEITEM)item, itemdata);
}

static void *go_tree_item_get_data(void *handle, void *item)
{
    return dw_tree_item_get_data((HWND)handle, (HTREEITEM)item);
}

static char *go_tree_get_title(void *handle, void *item)
{
    return dw_tree_get_title((HWND)handle, (HTREEITEM)item);
}

static void *go_html_new(unsigned long id)
{
   return (void *)dw_html_new(id);
}


static void go_html_action(void *hwnd, int action)
{
    dw_html_action((HWND)hwnd, action);
}

static int go_html_raw(void *hwnd, char *string)
{
    return dw_html_raw((HWND)hwnd, string);
}

static int go_html_url(void *hwnd, char *url)
{
    return dw_html_url((HWND)hwnd, url);
}

static void *go_mle_new(unsigned long id)
{
   return (void *)dw_mle_new(id);
}

static unsigned int go_mle_import(void *handle, char *buffer, int startpoint)
{
    return dw_mle_import((HWND)handle, buffer, startpoint);
}

static void go_mle_export(void *handle, char *buffer, int startpoint, int length)
{
    dw_mle_export((HWND)handle, buffer, startpoint, length);
}

static void go_mle_get_size(void *handle, unsigned long *bytes, unsigned long *lines)
{
    dw_mle_get_size((HWND)handle, bytes, lines);
}

static void go_mle_delete(void *handle, int startpoint, int length)
{
    dw_mle_delete((HWND)handle, startpoint, length);
}

static void go_mle_clear(void *handle)
{
    dw_mle_clear((HWND)handle);
}

static void go_mle_freeze(void *handle)
{
    dw_mle_freeze((HWND)handle);
}

static void go_mle_thaw(void *handle)
{
    dw_mle_thaw((HWND)handle);
}

static void go_mle_set_cursor(void *handle, int point)
{
    dw_mle_set_cursor((HWND)handle, point);
}

static void go_mle_set_visible(void *handle, int line)
{
    dw_mle_set_visible((HWND)handle, line);
}

static void go_mle_set_editable(void *handle, int state)
{
    dw_mle_set_editable((HWND)handle, state);
}

static void go_mle_set_word_wrap(void *handle, int state)
{
    dw_mle_set_word_wrap((HWND)handle, state);
}

static int go_mle_search(void *handle, char *text, int point, unsigned long flags)
{
    return dw_mle_search((HWND)handle, text, point, flags);
}

static void *go_container_new(unsigned long id, int multi)
{
    return (void *)dw_container_new(id, multi);
}

static char **go_string_array_make(int size) 
{
    return calloc(sizeof(char*), size);
}

static void go_string_array_set(char **a, char *s, int n) 
{
    a[n] = s;
}

static void go_string_array_free(char **a, int size) 
{
    int x;
    
    for(x = 0; x < size; x++)
        free(a[x]);
    free(a);
}

static int go_container_setup(void *handle, unsigned long *flags, char **titles, int count, int separator)
{
    return dw_container_setup((HWND)handle, flags, titles, count, separator);
}

static void * go_container_alloc(void *handle, int rowcount)
{
    return dw_container_alloc((HWND)handle, rowcount);
}

static void go_container_set_item(void *handle, void *pointer, int column, int row, void *data)
{
    dw_container_set_item((HWND)handle, pointer, column, row, data);
}

static void go_container_set_item_ulong(void *handle, void *pointer, int column, int row, unsigned long val)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_container_set_item((HWND)handle, pointer, column, row, &val);
}

static void go_container_set_item_icon(void *handle, void *pointer, int column, int row, void *icon)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_container_set_item((HWND)handle, pointer, column, row, &icon);
}

static void go_container_set_item_time(void *handle, void *pointer, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_container_set_item((HWND)handle, pointer, column, row, &time);
}

static void go_container_set_item_date(void *handle, void *pointer, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_container_set_item((HWND)handle, pointer, column, row, &date);
}

static void go_container_change_item(void *handle, int column, int row, void *data)
{
    dw_container_change_item((HWND)handle, column, row, data);
}

static void go_container_change_item_ulong(void *handle, int column, int row, unsigned long val)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_container_change_item((HWND)handle, column, row, &val);
}

static void go_container_change_item_icon(void *handle, int column, int row, void *icon)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_container_change_item((HWND)handle, column, row, &icon);
}

static void go_container_change_item_time(void *handle, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_container_change_item((HWND)handle, column, row, &time);
}

static void go_container_change_item_date(void *handle, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_container_change_item((HWND)handle, column, row, &date);
}

static void go_container_set_column_width(void *handle, int column, int width)
{
    dw_container_set_column_width((HWND)handle, column, width);
}

static void go_container_change_row_title(void *handle, int row, char *title)
{
    dw_container_change_row_title((HWND)handle, row, title);
}

static void go_container_change_row_data(void *handle, int row, void *data)
{
    dw_container_change_row_title((HWND)handle, row, (char *)data);
}

static void go_container_insert(void *handle, void *pointer, int rowcount)
{
    dw_container_insert((HWND)handle, pointer, rowcount);
}

static void go_container_clear(void *handle, int redraw)
{
    dw_container_clear((HWND)handle, redraw);
}

static void go_container_delete(void *handle, int rowcount)
{
    dw_container_delete((HWND)handle, rowcount);
}

static char *go_container_query_start(void *handle, unsigned long flags)
{
    return dw_container_query_start((HWND)handle, flags);
}

static char *go_container_query_next(void *handle, unsigned long flags)
{
    return dw_container_query_next((HWND)handle, flags);
}

static void go_container_scroll(void *handle, int direction, long rows)
{
    dw_container_scroll((HWND)handle, direction, rows);
}

static void go_container_cursor(void *handle, char *text)
{
    dw_container_cursor((HWND)handle, text);
}

static void go_container_delete_row(void *handle, char *text)
{
    dw_container_delete_row((HWND)handle, text);
}

static void go_container_optimize(void *handle)
{
    dw_container_optimize((HWND)handle);
}

static void go_container_set_stripe(void *handle, unsigned long oddcolor, unsigned long evencolor)
{
    dw_container_set_stripe((HWND)handle, oddcolor, evencolor);
}

static void go_filesystem_set_column_title(void *handle, char *title)
{
    dw_filesystem_set_column_title((HWND)handle, title);
}

static int go_filesystem_setup(void *handle, unsigned long *flags, char **titles, int count)
{
    return dw_filesystem_setup((HWND)handle, flags, titles, count);
}

static void go_filesystem_set_item(void *handle, void *pointer, int column, int row, void *data)
{
    dw_filesystem_set_item((HWND)handle, pointer, column, row, data);
}

static void go_filesystem_set_item_ulong(void *handle, void *pointer, int column, int row, unsigned long val)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_filesystem_set_item((HWND)handle, pointer, column, row, &val);
}

static void go_filesystem_set_item_icon(void *handle, void *pointer, int column, int row, void *icon)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_filesystem_set_item((HWND)handle, pointer, column, row, &icon);
}

static void go_filesystem_set_item_time(void *handle, void *pointer, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_filesystem_set_item((HWND)handle, pointer, column, row, &time);
}

static void go_filesystem_set_item_date(void *handle, void *pointer, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_filesystem_set_item((HWND)handle, pointer, column, row, &date);
}

static void go_filesystem_set_file(void *handle, void *pointer, int row, char *filename, void *icon)
{
    dw_filesystem_set_file((HWND)handle, pointer, row, filename, (HICN)icon);
}

static void go_filesystem_change_item(void *handle, int column, int row, void *data)
{
    dw_filesystem_change_item((HWND)handle, column, row, data);
}

static void go_filesystem_change_item_ulong(void *handle, int column, int row, unsigned long val)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_filesystem_change_item((HWND)handle, column, row, &val);
}

static void go_filesystem_change_item_icon(void *handle, int column, int row, void *icon)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_filesystem_change_item((HWND)handle, column, row, &icon);
}

static void go_filesystem_change_item_time(void *handle, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_filesystem_change_item((HWND)handle, column, row, &time);
}

static void go_filesystem_change_item_date(void *handle, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_filesystem_change_item((HWND)handle, column, row, &date);
}

static void go_filesystem_change_file(void *handle, int row, char *filename, void *icon)
{
    dw_filesystem_change_file((HWND)handle, row, filename, (HICN)icon);
}

static int go_container_get_column_type(void *handle, int column)
{
    return dw_container_get_column_type((HWND)handle, column);
}
static int go_filesystem_get_column_type(void *handle, int column)
{
    return dw_filesystem_get_column_type((HWND)handle, column);
}

static void *go_calendar_new(unsigned long id)
{
   return (void *)dw_calendar_new(id);
}

static void go_calendar_set_date(void *handle, unsigned int year, unsigned int month, unsigned int day)
{
    dw_calendar_set_date((HWND)handle, year, month, day);
}

static void go_calendar_get_date(void *handle, unsigned int *year, unsigned int *month, unsigned int *day)
{
    dw_calendar_get_date((HWND)handle, year, month, day);
}

static void *go_bitmap_new(unsigned long id)
{
   return (void *)dw_bitmap_new(id);
}

static void *go_bitmapbutton_new(char *text, unsigned long id)
{
   return (void *)dw_bitmapbutton_new(text, id);
}

static void *go_bitmapbutton_new_from_file(char *text, unsigned long id, char *filename)
{
   return (void *)dw_bitmapbutton_new_from_file(text, id, filename);
}

static void *go_splitbar_new(int type, void *topleft, void *bottomright, unsigned long cid)
{
   return (void *)dw_splitbar_new(type, (HWND)topleft, (HWND)bottomright, cid);
}

static void go_splitbar_set(void *handle, float position)
{
    dw_splitbar_set((HWND)handle, position);
}

static float go_splitbar_get(void *handle)
{
    return dw_splitbar_get((HWND)handle);
}

static int go_print_run(void *print, unsigned long flags)
{
    return dw_print_run((HPRINT)print, flags);
}

static void go_print_cancel(void *print)
{
    return dw_print_cancel((HPRINT)print);
}

extern int go_int_callback_basic(void *pfunc, void* window, void *data, int flags);
extern int go_int_callback_configure(void *pfunc, void* window, int width, int height, void *data, int flags);
extern int go_int_callback_keypress(void *pfunc, void *window, char ch, int vk, int state, void *data, char *utf8, int flags);
extern int go_int_callback_mouse(void *pfunc, void* window, int x, int y, int mask, void *data, int flags);
extern int go_int_callback_expose(void *pfunc, void* window, int x, int y, int width, int height, void *data, int flags);
extern int go_int_callback_string(void *pfunc, void* window, char *str, void *data, int flags);
extern int go_int_callback_item_context(void *pfunc, void *window, char *text, int x, int y, void *data, void *itemdata, int flags);
extern int go_int_callback_item_select(void *pfunc, void *window, void *item, char *text, void *data, void *itemdata, int flags);
extern int go_int_callback_numeric(void *pfunc, void* window, int val, void *data, int flags);
extern int go_int_callback_ulong(void *pfunc, void* window, unsigned long val, void *data, int flags);
extern int go_int_callback_tree(void *pfunc, void* window, void *item, void *data, int flags);
extern int go_int_callback_timer(void *pfunc, void *data, int flags);
extern int go_int_callback_print(void *pfunc, void *print, void *pixmap, int page_num, void *data, int flags);

static int DWSIGNAL go_callback_basic(HWND window, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_basic(param[0], (void *)window, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_configure(HWND window, int width, int height, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_configure(param[0], (void *)window, width, height, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_keypress(HWND window, char ch, int vk, int state, void *data, char *utf8)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_keypress(param[0], (void *)window, ch, vk, state, param[1], utf8, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_mouse(HWND window, int x, int y, int mask, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_mouse(param[0], (void *)window, x, y, mask, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_expose(HWND window,  DWExpose *exp, void *data)
{
   if(data && exp)
   {
      void **param = (void **)data;
      return go_int_callback_expose(param[0], (void *)window, exp->x, exp->y, exp->width, exp->height, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_string(HWND window, char *str, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_string(param[0], (void *)window, str, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_item_context(HWND window, char *text, int x, int y, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_context(param[0], (void *)window, text, x, y, param[1], itemdata, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_item_select(HWND window, HTREEITEM item, char *text, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_select(param[0], (void *)window, (void *)item, text, param[1], itemdata, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_numeric(HWND window, int val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_numeric(param[0], (void *)window, val, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_ulong(HWND window, unsigned long val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_ulong(param[0], (void *)window, val, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_tree(HWND window, HTREEITEM tree, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_tree(param[0], (void *)window, (void *)tree, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_timer(void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_timer(param[0], param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_print(HPRINT print, HPIXMAP pixmap, int page_num, void *data)
{
    if(data)
    {
       void **param = (void **)data;
       return go_int_callback_print(param[0], (void *)print, (void *)pixmap, page_num, param[1], DW_POINTER_TO_INT(param[2]));
    }
    return 0;
}

static void *go_print_new(char *jobname, unsigned long flags, unsigned int pages, void *drawfunc, void *drawdata, int sflags)
{
    void **param = malloc(sizeof(void *) * 3);
   
    if(param && drawfunc)
    {
       param[0] = drawfunc;
       param[1] = drawdata;
       param[2] = DW_INT_TO_POINTER(sflags);
       return (void*)dw_print_new(jobname, flags, pages, DW_SIGNAL_FUNC(go_callback_print), param);
    }
    return NULL;
}

static int go_timer_connect(int interval, void *sigfunc, void *data, int flags)
{
   void **param = malloc(sizeof(void *) * 3);
   
   if(param && sigfunc)
   {
      param[0] = sigfunc;
      param[1] = data;
      param[2] = DW_INT_TO_POINTER(flags);
      return dw_timer_connect(interval, DW_SIGNAL_FUNC(go_callback_timer), param);
   }
   return 0;
}

static void DWSIGNAL go_signal_free(HWND window, void *data)
{
   if(data)
   {
      free(data);
   }
}

static void go_signal_connect(void *window, char *signame, void *sigfunc, void *data, int flags)
{
   void **param = malloc(sizeof(void *) * 3);
   void *func = (void *)go_callback_basic;
   
   if(param && sigfunc)
   {
      param[0] = sigfunc;
      param[1] = data;
      param[2] = DW_INT_TO_POINTER(flags);
      
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
      
      dw_signal_connect_data((HWND)window, signame, func, DW_SIGNAL_FUNC(go_signal_free), param);
   }
}
