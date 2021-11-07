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

static uintptr_t go_window_new(uintptr_t owner, char *title, unsigned long flags)
{
   return (uintptr_t)dw_window_new((HWND)owner, title, flags);
}

static int go_window_show(uintptr_t handle)
{
   return dw_window_show((HWND)handle);
}

static int go_window_hide(uintptr_t handle)
{
   return dw_window_hide((HWND)handle);
}

static int go_window_lower(uintptr_t handle)
{
   return dw_window_lower((HWND)handle);
}

static int go_window_raise(uintptr_t handle)
{
   return dw_window_raise((HWND)handle);
}

static int go_window_minimize(uintptr_t handle)
{
   return dw_window_minimize((HWND)handle);
}

static void go_window_set_pos(uintptr_t handle, long x, long y)
{
   dw_window_set_pos((HWND)handle, x, y);
}

static void go_window_set_pos_size(uintptr_t handle, long x, long y, unsigned long width, unsigned long height)
{
   dw_window_set_pos_size((HWND)handle, x, y, width, height);
}

static void go_window_set_size(uintptr_t handle, unsigned long width, unsigned long height)
{
   dw_window_set_size((HWND)handle, width, height);
}

static int go_window_set_color(uintptr_t handle, unsigned long fore, unsigned long back)
{
   return dw_window_set_color((HWND)handle, fore, back);
}

static void go_window_set_style(uintptr_t handle, unsigned long style, unsigned long mask)
{
   dw_window_set_style((HWND)handle, style, mask);
}

static void go_window_click_default(uintptr_t window, uintptr_t next)
{
   dw_window_click_default((HWND)window, (HWND)next);
}

static void go_window_default(uintptr_t window, uintptr_t defaultitem)
{
   dw_window_default((HWND)window, (HWND)defaultitem);
}

static int go_window_destroy(uintptr_t handle)
{
   return dw_window_destroy((HWND)handle);
}

static void go_window_disable(uintptr_t handle)
{
   dw_window_disable((HWND)handle);
}

static void go_window_enable(uintptr_t handle)
{
   dw_window_enable((HWND)handle);
}

static uintptr_t go_window_from_id(uintptr_t handle, int id)
{
   return (uintptr_t)dw_window_from_id((HWND)handle, id);
}

static void *go_window_get_data(uintptr_t handle, char *dataname)
{
   return dw_window_get_data((HWND)handle, dataname);
}

static void go_window_set_data(uintptr_t handle, char *dataname, void *data)
{
   dw_window_set_data((HWND)handle, dataname, data);
}

static char *go_window_get_font(uintptr_t handle)
{
   return dw_window_get_font((HWND)handle);
}

static int go_window_set_font(uintptr_t handle, char *fontname)
{
   return dw_window_set_font((HWND)handle, fontname);
}

static void go_window_get_pos_size(uintptr_t handle, long *x, long *y, unsigned long *width, unsigned long *height)
{
   dw_window_get_pos_size((HWND)handle, x, y, width, height);
}

static void go_window_get_preferred_size(uintptr_t handle, int *width, int *height)
{
   dw_window_get_preferred_size((HWND)handle, width, height);
}

static char *go_window_get_text(uintptr_t handle)
{
   return dw_window_get_text((HWND)handle);
}

static void go_window_set_text(uintptr_t handle, char *text)
{
   dw_window_set_text((HWND)handle, text);
}

static void go_window_set_tooltip(uintptr_t handle, char *bubbletext)
{
   dw_window_set_tooltip((HWND)handle, bubbletext);
}

static void go_window_redraw(uintptr_t handle)
{
   dw_window_redraw((HWND)handle);
}

static void go_window_capture(uintptr_t handle)
{
   dw_window_capture((HWND)handle);
}

static void go_window_set_bitmap(uintptr_t handle, unsigned long cid, char *filename)
{
   dw_window_set_bitmap((HWND)handle, cid, filename);
}

static int go_window_set_border(uintptr_t handle, int border)
{
   return dw_window_set_border((HWND)handle, border);
}

static void go_window_set_focus(uintptr_t handle)
{
   dw_window_set_focus((HWND)handle);
}

static void go_window_set_gravity(uintptr_t handle, int horz, int vert)
{
   dw_window_set_gravity((HWND)handle, horz, vert);
}

static void go_window_set_icon(uintptr_t handle, uintptr_t icon)
{
   dw_window_set_icon((HWND)handle, (HICN)icon);
}

static void go_window_set_pointer(uintptr_t handle, int cursortype)
{
   dw_window_set_pointer((HWND)handle, cursortype);
}

static uintptr_t go_box_new(int type, int pad)
{
   return (uintptr_t)dw_box_new(type, pad);
}

static void go_box_pack_at_index(uintptr_t box, uintptr_t item, int index, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_at_index((HWND)box, (HWND)item, index, width, height, hsize, vsize, pad);
}

static void go_box_pack_end(uintptr_t box, uintptr_t item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_end((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

static void go_box_pack_start(uintptr_t box, uintptr_t item, int width, int height, int hsize, int vsize, int pad)
{
   dw_box_pack_start((HWND)box, (HWND)item, width, height, hsize, vsize, pad);
}

static int go_box_unpack(uintptr_t handle)
{
   return dw_box_unpack((HWND)handle);
}

static uintptr_t go_box_unpack_at_index(uintptr_t handle, int index)
{
   return (uintptr_t)dw_box_unpack_at_index((HWND)handle, index);
}

static uintptr_t go_text_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_text_new(text, id);
}

static uintptr_t go_status_text_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_status_text_new(text, id);
}

static uintptr_t go_entryfield_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_entryfield_new(text, id);
}

static uintptr_t go_entryfield_password_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_entryfield_password_new(text, id);
}

static void go_entryfield_set_limit(uintptr_t handle, int limit)
{
   dw_entryfield_set_limit((HWND)handle, limit);
}

static uintptr_t go_button_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_button_new(text, id);
}

static uintptr_t go_menu_new(unsigned long cid)
{
    return (uintptr_t)dw_menu_new(cid);
}

static uintptr_t go_menubar_new(uintptr_t location)
{
    return (uintptr_t)dw_menubar_new((HWND)location);
}

static uintptr_t go_menu_append_item(uintptr_t menu, char *title, unsigned long id, unsigned long flags, int end, int check, uintptr_t submenu)
{
    return (uintptr_t)dw_menu_append_item((HMENUI)menu, title, id, flags, end, check, (HMENUI)submenu);
}

static int go_menu_delete_item(uintptr_t menu, unsigned long cid)
{
    return dw_menu_delete_item((HMENUI)menu, cid);
}

static void go_menu_destroy(uintptr_t menu)
{
    HMENUI thismenu = (HMENUI)menu;
    dw_menu_destroy(&thismenu);
}

static void go_menu_item_set_state(uintptr_t menu, unsigned long cid, unsigned long flags)
{
    dw_menu_item_set_state((HMENUI)menu, cid, flags);
}

static void go_menu_popup(uintptr_t menu, uintptr_t parent, int x, int y)
{
    HMENUI thismenu = (HMENUI)menu;
    dw_menu_popup(&thismenu, (HWND)parent, x, y);
}

static uintptr_t go_notebook_new(unsigned long cid, int top)
{
    return (uintptr_t)dw_notebook_new(cid, top);
}

static void go_notebook_pack(uintptr_t handle, unsigned long pageid, uintptr_t page)
{
    dw_notebook_pack((HWND)handle, pageid, (HWND)page);
}

static void go_notebook_page_destroy(uintptr_t handle, unsigned long pageid)
{
    dw_notebook_page_destroy((HWND)handle, (unsigned int)pageid);
}

static unsigned long go_notebook_page_get(uintptr_t handle)
{
    return dw_notebook_page_get((HWND)handle);
}

static unsigned long go_notebook_page_new(uintptr_t handle, unsigned long flags, int front)
{
    return dw_notebook_page_new((HWND)handle, flags, front);
}

static void go_notebook_page_set(uintptr_t handle, unsigned long pageid)
{
    dw_notebook_page_set((HWND)handle, (unsigned int)pageid);
}

static void go_notebook_page_set_text(uintptr_t handle, unsigned long pageid, char *text)
{
    dw_notebook_page_set_text((HWND)handle, pageid, text);
}

static uintptr_t go_icon_load_from_file(char *filename)
{
    return (uintptr_t)dw_icon_load_from_file(filename);
}

static uintptr_t go_icon_load(unsigned long module, unsigned long cid)
{
    return (uintptr_t)dw_icon_load(module, cid);
}

static void go_taskbar_delete(uintptr_t handle, uintptr_t icon)
{
    dw_taskbar_delete((HWND)handle, (HICN)icon);
}

static void go_taskbar_insert(uintptr_t handle, uintptr_t icon, char *bubbletext)
{
    dw_taskbar_insert((HWND)handle, (HICN)icon, bubbletext);
}

static uintptr_t go_combobox_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_combobox_new(text, id);
}

static uintptr_t go_listbox_new(unsigned long id, int multi)
{
   return (uintptr_t)dw_listbox_new(id, multi);
}

static void go_listbox_append(uintptr_t handle, char *text)
{
    dw_listbox_append((HWND)handle, text);
}

static void go_listbox_list_append(uintptr_t handle, char **text, int count)
{
    dw_listbox_list_append((HWND)handle, text, count);
}

static void go_listbox_insert(uintptr_t handle, char *text, int pos)
{
    dw_listbox_insert((HWND)handle, text, pos);
}

static void go_listbox_clear(uintptr_t handle)
{
    dw_listbox_clear((HWND)handle);
}

static int go_listbox_count(uintptr_t handle)
{
    return dw_listbox_count((HWND)handle);
}

static void go_listbox_set_top(uintptr_t handle, int top)
{
    dw_listbox_set_top((HWND)handle, top);
}

static void go_listbox_select(uintptr_t handle, int index, int state)
{
    dw_listbox_select((HWND)handle, index, state);
}

static void go_listbox_delete(uintptr_t handle, int index)
{
    dw_listbox_delete((HWND)handle, index);
}

static void go_listbox_get_text(uintptr_t handle, int index, char *text, int length)
{
    dw_listbox_get_text((HWND)handle, index, text, length);
}

static void go_listbox_set_text(uintptr_t handle, int index, char *text)
{
    dw_listbox_set_text((HWND)handle, index, text);
}

static int go_listbox_selected(uintptr_t handle)
{
    return dw_listbox_selected((HWND)handle);
}

static int go_listbox_selected_multi(uintptr_t handle, int where)
{
    return dw_listbox_selected_multi((HWND)handle, where);
}

static uintptr_t go_spinbutton_new(char *text, unsigned long id)
{
    return (uintptr_t)dw_spinbutton_new(text, id);
}

static void go_spinbutton_set_pos(uintptr_t handle, long position)
{
    dw_spinbutton_set_pos((HWND)handle, position);
}
static void go_spinbutton_set_limits(uintptr_t handle, long upper, long lower)
{
    dw_spinbutton_set_limits((HWND)handle, upper, lower);
}

static long go_spinbutton_get_pos(uintptr_t handle)
{
    return dw_spinbutton_get_pos((HWND)handle);
}

static uintptr_t go_radiobutton_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_radiobutton_new(text, id);
}

static uintptr_t go_checkbox_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_checkbox_new(text, id);
}

static int go_checkbox_get(uintptr_t handle)
{
    return dw_checkbox_get((HWND)handle);
}

static void go_checkbox_set(uintptr_t handle, int value)
{
    return dw_checkbox_set((HWND)handle, value);
}

static uintptr_t go_percent_new(unsigned long id)
{
   return (uintptr_t)dw_percent_new(id);
}

static void go_percent_set_pos(uintptr_t handle, unsigned int position)
{
   dw_percent_set_pos((HWND)handle, position);
}

static uintptr_t go_slider_new(int vertical, int increments, unsigned long id)
{
   return (uintptr_t)dw_slider_new(vertical, increments, id);
}

static unsigned int go_slider_get_pos(uintptr_t handle)
{
   return dw_slider_get_pos((HWND)handle);
}

static void go_slider_set_pos(uintptr_t handle, unsigned int pos)
{
    dw_slider_set_pos((HWND)handle, pos);
}

static uintptr_t go_scrollbar_new(int vertical, unsigned long id)
{
   return (uintptr_t)dw_scrollbar_new(vertical, id);
}

static unsigned int go_scrollbar_get_pos(uintptr_t handle)
{
   return dw_scrollbar_get_pos((HWND)handle);
}

static void go_scrollbar_set_pos(uintptr_t handle, unsigned int pos)
{
    dw_scrollbar_set_pos((HWND)handle, pos);
}

static void go_scrollbar_set_range(uintptr_t handle, unsigned int range, unsigned int visible)
{
    dw_scrollbar_set_range((HWND)handle, range, visible);
}

static uintptr_t go_scrollbox_new(int type, int pad)
{
   return (uintptr_t)dw_scrollbox_new(type, pad);
}

static int go_scrollbox_get_pos(uintptr_t handle, int orient)
{
    return dw_scrollbox_get_pos((HWND)handle, orient);
}

static int go_scrollbox_get_range(uintptr_t handle, int orient)
{
    return dw_scrollbox_get_range((HWND)handle, orient);
}

static uintptr_t go_groupbox_new(int type, int pad, char *title)
{
   return (uintptr_t)dw_groupbox_new(type, pad, title);
}

static uintptr_t go_render_new(unsigned long id)
{
   return (uintptr_t)dw_render_new(id);
}

static void go_font_text_extents_get(uintptr_t handle, uintptr_t pixmap, char *text, int *width, int *height)
{
   dw_font_text_extents_get((HWND)handle, (HPIXMAP)pixmap, text, width, height);
}

static uintptr_t go_pixmap_new(uintptr_t handle, unsigned long width, unsigned long height, unsigned long depth) 
{
    return (uintptr_t)dw_pixmap_new((HWND)handle, width, height, (int)depth);
}

static uintptr_t go_pixmap_new_from_file(uintptr_t handle, char *filename) 
{
    return (uintptr_t)dw_pixmap_new_from_file((HWND)handle, filename);
}

static uintptr_t go_pixmap_grab(uintptr_t handle, unsigned long cid) 
{
    return (uintptr_t)dw_pixmap_grab((HWND)handle, cid);
}

static void go_pixmap_bitblt(uintptr_t dest, uintptr_t destp, int xdest, int ydest, int width, int height, uintptr_t src, uintptr_t srcp, int xsrc, int ysrc)
{
    dw_pixmap_bitblt((HWND)dest, (HPIXMAP)destp, xdest, ydest, width, height, (HWND)src, (HPIXMAP)srcp, xsrc, ysrc);
}

static int go_pixmap_stretch_bitblt(uintptr_t dest, uintptr_t destp, int xdest, int ydest, int width, int height, uintptr_t src, uintptr_t srcp, int xsrc, int ysrc, int srcwidth, int srcheight)
{
    return dw_pixmap_stretch_bitblt((HWND)dest, (HPIXMAP)destp, xdest, ydest, width, height, (HWND)src, (HPIXMAP)srcp, xsrc, ysrc, srcwidth, srcheight);
}

static void go_pixmap_set_transparent_color(uintptr_t pixmap, unsigned long color)
{
    dw_pixmap_set_transparent_color((HPIXMAP)pixmap, color);
}

static int go_pixmap_set_font(uintptr_t pixmap, char *fontname)
{
    return dw_pixmap_set_font((HPIXMAP)pixmap, fontname);
}

static void go_pixmap_destroy(uintptr_t pixmap)
{
    dw_pixmap_destroy((HPIXMAP)pixmap);
}

static int go_pixmap_width(uintptr_t pixmap)
{
    return (int)DW_PIXMAP_WIDTH(((HPIXMAP)pixmap));
}

static int go_pixmap_height(uintptr_t pixmap)
{
    return (int)DW_PIXMAP_HEIGHT(((HPIXMAP)pixmap));
}

static void go_draw_point(uintptr_t handle, uintptr_t pixmap, int x, int y)
{
    dw_draw_point((HWND)handle, (HPIXMAP)pixmap, x, y);
}

static void go_draw_line(uintptr_t handle, uintptr_t pixmap, int x1, int y1, int x2, int y2)
{
    dw_draw_line((HWND)handle, (HPIXMAP)pixmap, x1, y1, x2, y2);
}

static void go_draw_polygon(uintptr_t handle, uintptr_t pixmap, int fill, int count, int x[], int y[])
{
    dw_draw_polygon((HWND)handle, (HPIXMAP)pixmap, fill, count, x, y);
}

static void go_draw_rect(uintptr_t handle, uintptr_t pixmap, int fill, int x, int y, int width, int height)
{
    dw_draw_rect((HWND)handle, (HPIXMAP)pixmap, fill, x, y, width, height);
}

static void go_draw_arc(uintptr_t handle, uintptr_t pixmap, int flags, int xorigin, int yorigin, int x1, int y1, int x2, int y2)
{
    dw_draw_arc((HWND)handle, (HPIXMAP)pixmap, flags, xorigin, yorigin, x1, y1, x2, y2);
}

static void go_draw_text(uintptr_t handle, uintptr_t pixmap, int x, int y, char *text)
{
    dw_draw_text((HWND)handle, (HPIXMAP)pixmap, x, y, text);
}

static uintptr_t go_tree_new(unsigned long id)
{
   return (uintptr_t)dw_tree_new(id);
}

static uintptr_t go_tree_insert(uintptr_t handle, char *title, uintptr_t icon, uintptr_t parent, void *itemdata)
{
    return (uintptr_t)dw_tree_insert((HWND)handle, title, (HICN)icon, (HTREEITEM)parent, itemdata);
}

static uintptr_t go_tree_insert_after(uintptr_t handle, uintptr_t item, char *title, uintptr_t icon, uintptr_t parent, void *itemdata)
{
    return (uintptr_t)dw_tree_insert_after((HWND)handle, (HTREEITEM)item, title, (HICN)icon, (HTREEITEM)parent, itemdata);
}

static void go_tree_clear(uintptr_t handle)
{
    dw_tree_clear((HWND)handle);
}

static void go_tree_item_delete(uintptr_t handle, uintptr_t item)
{
    dw_tree_item_delete((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_change(uintptr_t handle, uintptr_t item, char *title, uintptr_t icon)
{
    dw_tree_item_change((HWND)handle, (HTREEITEM)item, title, (HICN)icon);
}

static void go_tree_item_expand(uintptr_t handle, uintptr_t item)
{
    dw_tree_item_expand((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_collapse(uintptr_t handle, uintptr_t item)
{
    dw_tree_item_collapse((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_select(uintptr_t handle, uintptr_t item)
{
    dw_tree_item_select((HWND)handle, (HTREEITEM)item);
}

static void go_tree_item_set_data(uintptr_t handle, uintptr_t item, void *itemdata)
{
    dw_tree_item_set_data((HWND)handle, (HTREEITEM)item, itemdata);
}

static void *go_tree_item_get_data(uintptr_t handle, uintptr_t item)
{
    return dw_tree_item_get_data((HWND)handle, (HTREEITEM)item);
}

static char *go_tree_get_title(uintptr_t handle, uintptr_t item)
{
    return dw_tree_get_title((HWND)handle, (HTREEITEM)item);
}

static uintptr_t go_html_new(unsigned long id)
{
   return (uintptr_t)dw_html_new(id);
}


static void go_html_action(uintptr_t hwnd, int action)
{
    dw_html_action((HWND)hwnd, action);
}

static int go_html_raw(uintptr_t hwnd, char *string)
{
    return dw_html_raw((HWND)hwnd, string);
}

static int go_html_url(uintptr_t hwnd, char *url)
{
    return dw_html_url((HWND)hwnd, url);
}

static uintptr_t go_mle_new(unsigned long id)
{
   return (uintptr_t)dw_mle_new(id);
}

static unsigned int go_mle_import(uintptr_t handle, char *buffer, int startpoint)
{
    return dw_mle_import((HWND)handle, buffer, startpoint);
}

static void go_mle_export(uintptr_t handle, char *buffer, int startpoint, int length)
{
    dw_mle_export((HWND)handle, buffer, startpoint, length);
}

static void go_mle_get_size(uintptr_t handle, unsigned long *bytes, unsigned long *lines)
{
    dw_mle_get_size((HWND)handle, bytes, lines);
}

static void go_mle_delete(uintptr_t handle, int startpoint, int length)
{
    dw_mle_delete((HWND)handle, startpoint, length);
}

static void go_mle_clear(uintptr_t handle)
{
    dw_mle_clear((HWND)handle);
}

static void go_mle_freeze(uintptr_t handle)
{
    dw_mle_freeze((HWND)handle);
}

static void go_mle_thaw(uintptr_t handle)
{
    dw_mle_thaw((HWND)handle);
}

static void go_mle_set_cursor(uintptr_t handle, int point)
{
    dw_mle_set_cursor((HWND)handle, point);
}

static void go_mle_set_visible(uintptr_t handle, int line)
{
    dw_mle_set_visible((HWND)handle, line);
}

static void go_mle_set_editable(uintptr_t handle, int state)
{
    dw_mle_set_editable((HWND)handle, state);
}

static void go_mle_set_word_wrap(uintptr_t handle, int state)
{
    dw_mle_set_word_wrap((HWND)handle, state);
}

static int go_mle_search(uintptr_t handle, char *text, int point, unsigned long flags)
{
    return dw_mle_search((HWND)handle, text, point, flags);
}

static uintptr_t go_container_new(unsigned long id, int multi)
{
    return (uintptr_t)dw_container_new(id, multi);
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

static int go_container_setup(uintptr_t handle, unsigned long *flags, char **titles, int count, int separator)
{
    return dw_container_setup((HWND)handle, flags, titles, count, separator);
}

static uintptr_t go_container_alloc(uintptr_t handle, int rowcount)
{
    return (uintptr_t)dw_container_alloc((HWND)handle, rowcount);
}

static void go_container_set_item(uintptr_t handle, uintptr_t pointer, int column, int row, void *data)
{
    dw_container_set_item((HWND)handle, (void *)pointer, column, row, data);
}

static void go_container_set_item_ulong(uintptr_t handle, uintptr_t pointer, int column, int row, unsigned long val)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_container_set_item((HWND)handle, (void *)pointer, column, row, &val);
}

static void go_container_set_item_icon(uintptr_t handle, uintptr_t pointer, int column, int row, uintptr_t icon)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_container_set_item((HWND)handle, (void *)pointer, column, row, &icon);
}

static void go_container_set_item_time(uintptr_t handle, uintptr_t pointer, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_container_set_item((HWND)handle, (void *)pointer, column, row, &time);
}

static void go_container_set_item_date(uintptr_t handle, uintptr_t pointer, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_container_set_item((HWND)handle, (void *)pointer, column, row, &date);
}

static void go_container_change_item(uintptr_t handle, int column, int row, void *data)
{
    dw_container_change_item((HWND)handle, column, row, data);
}

static void go_container_change_item_ulong(uintptr_t handle, int column, int row, unsigned long val)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_container_change_item((HWND)handle, column, row, &val);
}

static void go_container_change_item_icon(uintptr_t handle, int column, int row, uintptr_t icon)
{
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_container_change_item((HWND)handle, column, row, &icon);
}

static void go_container_change_item_time(uintptr_t handle, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_container_change_item((HWND)handle, column, row, &time);
}

static void go_container_change_item_date(uintptr_t handle, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_container_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_container_change_item((HWND)handle, column, row, &date);
}

static void go_container_set_column_width(uintptr_t handle, int column, int width)
{
    dw_container_set_column_width((HWND)handle, column, width);
}

static void go_container_change_row_title(uintptr_t handle, int row, char *title)
{
    dw_container_change_row_title((HWND)handle, row, title);
}

static void go_container_change_row_data(uintptr_t handle, int row, void *data)
{
    dw_container_change_row_data((HWND)handle, row, data);
}

static void go_container_set_row_title(uintptr_t pointer, int row, char *title)
{
    dw_container_set_row_title((void *)pointer, row, title);
}

static void go_container_set_row_data(uintptr_t pointer, int row, void *data)
{
    dw_container_set_row_data((void *)pointer, row, data);
}

static void go_container_insert(uintptr_t handle, uintptr_t pointer, int rowcount)
{
    dw_container_insert((HWND)handle, (void *)pointer, rowcount);
}

static void go_container_clear(uintptr_t handle, int redraw)
{
    dw_container_clear((HWND)handle, redraw);
}

static void go_container_delete(uintptr_t handle, int rowcount)
{
    dw_container_delete((HWND)handle, rowcount);
}

static char *go_container_query_start(uintptr_t handle, unsigned long flags)
{
    return dw_container_query_start((HWND)handle, flags);
}

static char *go_container_query_next(uintptr_t handle, unsigned long flags)
{
    return dw_container_query_next((HWND)handle, flags);
}

static void go_container_scroll(uintptr_t handle, int direction, long rows)
{
    dw_container_scroll((HWND)handle, direction, rows);
}

static void go_container_cursor(uintptr_t handle, char *text)
{
    dw_container_cursor((HWND)handle, text);
}

static void go_container_cursor_by_data(uintptr_t handle, void *data)
{
    dw_container_cursor_by_data((HWND)handle, data);
}

static void go_container_delete_row(uintptr_t handle, char *text)
{
    dw_container_delete_row((HWND)handle, text);
}

static void go_container_delete_row_by_data(uintptr_t handle, void *data)
{
    dw_container_delete_row_by_data((HWND)handle, data);
}

static void go_container_optimize(uintptr_t handle)
{
    dw_container_optimize((HWND)handle);
}

static void go_container_set_stripe(uintptr_t handle, unsigned long oddcolor, unsigned long evencolor)
{
    dw_container_set_stripe((HWND)handle, oddcolor, evencolor);
}

static void go_filesystem_set_column_title(uintptr_t handle, char *title)
{
    dw_filesystem_set_column_title((HWND)handle, title);
}

static int go_filesystem_setup(uintptr_t handle, unsigned long *flags, char **titles, int count)
{
    return dw_filesystem_setup((HWND)handle, flags, titles, count);
}

static void go_filesystem_set_item(uintptr_t handle, uintptr_t pointer, int column, int row, void *data)
{
    dw_filesystem_set_item((HWND)handle, (void *)pointer, column, row, data);
}

static void go_filesystem_set_item_ulong(uintptr_t handle, uintptr_t pointer, int column, int row, unsigned long val)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_filesystem_set_item((HWND)handle, (void *)pointer, column, row, &val);
}

static void go_filesystem_set_item_icon(uintptr_t handle, uintptr_t pointer, int column, int row, uintptr_t icon)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_filesystem_set_item((HWND)handle, (void *)pointer, column, row, &icon);
}

static void go_filesystem_set_item_time(uintptr_t handle, uintptr_t pointer, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_filesystem_set_item((HWND)handle, (void *)pointer, column, row, &time);
}

static void go_filesystem_set_item_date(uintptr_t handle, uintptr_t pointer, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_filesystem_set_item((HWND)handle, (void *)pointer, column, row, &date);
}

static void go_filesystem_set_file(uintptr_t handle, uintptr_t pointer, int row, char *filename, uintptr_t icon)
{
    dw_filesystem_set_file((HWND)handle, (void *)pointer, row, filename, (HICN)icon);
}

static void go_filesystem_change_item(uintptr_t handle, int column, int row, void *data)
{
    dw_filesystem_change_item((HWND)handle, column, row, data);
}

static void go_filesystem_change_item_ulong(uintptr_t handle, int column, int row, unsigned long val)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_ULONG)
        dw_filesystem_change_item((HWND)handle, column, row, &val);
}

static void go_filesystem_change_item_icon(uintptr_t handle, int column, int row, uintptr_t icon)
{
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_BITMAPORICON)
        dw_filesystem_change_item((HWND)handle, column, row, &icon);
}

static void go_filesystem_change_item_time(uintptr_t handle, int column, int row, int seconds, int minutes, int hours)
{
    CTIME time;
    
    time.seconds = seconds;
    time.minutes = minutes;
    time.hours = hours;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_TIME)
        dw_filesystem_change_item((HWND)handle, column, row, &time);
}

static void go_filesystem_change_item_date(uintptr_t handle, int column, int row, int day, int month, int year)
{
    CDATE date;
    
    date.day = day;
    date.month = month;
    date.year = year;
    
    if(dw_filesystem_get_column_type((HWND)handle, column) == DW_CFA_DATE)
        dw_filesystem_change_item((HWND)handle, column, row, &date);
}

static void go_filesystem_change_file(uintptr_t handle, int row, char *filename, uintptr_t icon)
{
    dw_filesystem_change_file((HWND)handle, row, filename, (HICN)icon);
}

static int go_container_get_column_type(uintptr_t handle, int column)
{
    return dw_container_get_column_type((HWND)handle, column);
}
static int go_filesystem_get_column_type(uintptr_t handle, int column)
{
    return dw_filesystem_get_column_type((HWND)handle, column);
}

static uintptr_t go_calendar_new(unsigned long id)
{
   return (uintptr_t)dw_calendar_new(id);
}

static void go_calendar_set_date(uintptr_t handle, unsigned int year, unsigned int month, unsigned int day)
{
    dw_calendar_set_date((HWND)handle, year, month, day);
}

static void go_calendar_get_date(uintptr_t handle, unsigned int *year, unsigned int *month, unsigned int *day)
{
    dw_calendar_get_date((HWND)handle, year, month, day);
}

static uintptr_t go_bitmap_new(unsigned long id)
{
   return (uintptr_t)dw_bitmap_new(id);
}

static uintptr_t go_bitmapbutton_new(char *text, unsigned long id)
{
   return (uintptr_t)dw_bitmapbutton_new(text, id);
}

static uintptr_t go_bitmapbutton_new_from_file(char *text, unsigned long id, char *filename)
{
   return (uintptr_t)dw_bitmapbutton_new_from_file(text, id, filename);
}

static uintptr_t go_splitbar_new(int type, uintptr_t topleft, uintptr_t bottomright, unsigned long cid)
{
   return (uintptr_t)dw_splitbar_new(type, (HWND)topleft, (HWND)bottomright, cid);
}

static void go_splitbar_set(uintptr_t handle, float position)
{
    dw_splitbar_set((HWND)handle, position);
}

static float go_splitbar_get(uintptr_t handle)
{
    return dw_splitbar_get((HWND)handle);
}

static int go_print_run(uintptr_t print, unsigned long flags)
{
    return dw_print_run((HPRINT)print, flags);
}

static void go_print_cancel(uintptr_t print)
{
    return dw_print_cancel((HPRINT)print);
}

static uintptr_t go_mutex_new(void)
{
    return (uintptr_t)dw_mutex_new();
}

static void go_mutex_close(uintptr_t mutex)
{
    dw_mutex_close((HMTX)mutex);
}

static void go_mutex_lock(uintptr_t mutex)
{
    dw_mutex_lock((HMTX)mutex);
}

static void go_mutex_unlock(uintptr_t mutex)
{
    dw_mutex_unlock((HMTX)mutex);
}

static int go_mutex_trylock(uintptr_t mutex)
{
    return dw_mutex_trylock((HMTX)mutex);
}

static uintptr_t go_dialog_new(void)
{
    return (uintptr_t)dw_dialog_new(NULL);
}

static int go_dialog_dismiss(uintptr_t dialog, void *result)
{
	return dw_dialog_dismiss((DWDialog *)dialog, result);
}

static void *go_dialog_wait(uintptr_t dialog)
{
	return dw_dialog_wait((DWDialog *)dialog);
}

static uintptr_t go_event_new(void)
{
    return (uintptr_t)dw_event_new();
}

static int go_event_close(uintptr_t event)
{
    HEV thisevent = (HEV)event;
    return dw_event_close(&thisevent);
}

static int go_event_post(uintptr_t event)
{
    return dw_event_post((HEV)event);
}

static int go_event_reset(uintptr_t event)
{
    return dw_event_reset((HEV)event);
}

static int go_event_wait(uintptr_t event, unsigned long timeout)
{
    return dw_event_wait((HEV)event, timeout);
}

extern int go_int_callback_basic(void *pfunc, uintptr_t window, void *data, unsigned int flags);
extern int go_int_callback_configure(void *pfunc, uintptr_t window, int width, int height, void *data, unsigned int flags);
extern int go_int_callback_keypress(void *pfunc, uintptr_t window, char ch, int vk, int state, void *data, char *utf8, unsigned int flags);
extern int go_int_callback_mouse(void *pfunc, uintptr_t window, int x, int y, int mask, void *data, unsigned int flags);
extern int go_int_callback_expose(void *pfunc, uintptr_t window, int x, int y, int width, int height, void *data, unsigned int flags);
extern int go_int_callback_item_enter(void *pfunc, uintptr_t window, char *text, void *data, void *itemdata, unsigned int flags);
extern int go_int_callback_item_context(void *pfunc, uintptr_t window, char *text, int x, int y, void *data, void *itemdata, unsigned int flags);
extern int go_int_callback_item_select(void *pfunc, uintptr_t window, uintptr_t item, char *text, void *data, void *itemdata, unsigned int flags);
extern int go_int_callback_numeric(void *pfunc, uintptr_t window, int val, void *data, unsigned int flags);
extern int go_int_callback_ulong(void *pfunc, uintptr_t window, unsigned long val, void *data, unsigned int flags);
extern int go_int_callback_notepage(void *pfunc, uintptr_t window, unsigned long val, void *data, unsigned int flags);
extern int go_int_callback_tree(void *pfunc, uintptr_t window, uintptr_t item, void *data, unsigned int flags);
extern int go_int_callback_timer(void *pfunc, void *data, unsigned int flags);
extern int go_int_callback_print(void *pfunc, uintptr_t print, uintptr_t pixmap, int page_num, void *data, unsigned int flags);
extern void go_callback_remove(void *pfunc);

static int DWSIGNAL go_callback_basic(HWND window, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_basic(param[0], (uintptr_t)window, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_configure(HWND window, int width, int height, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_configure(param[0], (uintptr_t)window, width, height, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_keypress(HWND window, char ch, int vk, int state, void *data, char *utf8)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_keypress(param[0], (uintptr_t)window, ch, vk, state, param[1], utf8, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_mouse(HWND window, int x, int y, int mask, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_mouse(param[0], (uintptr_t)window, x, y, mask, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_expose(HWND window,  DWExpose *exp, void *data)
{
   if(data && exp)
   {
      void **param = (void **)data;
      return go_int_callback_expose(param[0], (uintptr_t)window, exp->x, exp->y, exp->width, exp->height, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_item_enter(HWND window, char *text, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_enter(param[0], (uintptr_t)window, text, param[1], itemdata, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_item_context(HWND window, char *text, int x, int y, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_context(param[0], (uintptr_t)window, text, x, y, param[1], itemdata, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_item_select(HWND window, HTREEITEM item, char *text, void *data, void *itemdata)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_item_select(param[0], (uintptr_t)window, (uintptr_t)item, text, param[1], itemdata, DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_numeric(HWND window, int val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_numeric(param[0], (uintptr_t)window, val, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

/*static int DWSIGNAL go_callback_ulong(HWND window, unsigned long val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_ulong(param[0], (uintptr_t)window, val, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}*/

static int DWSIGNAL go_callback_notepage(HWND window, unsigned long val, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_notepage(param[0], (uintptr_t)window, val, param[1], DW_POINTER_TO_INT(param[2]));
   }
   return 0;
}

static int DWSIGNAL go_callback_tree(HWND window, HTREEITEM tree, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      return go_int_callback_tree(param[0], (uintptr_t)window, (uintptr_t)tree, param[1], DW_POINTER_TO_INT(param[2]));
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
       return go_int_callback_print(param[0], (uintptr_t)print, (uintptr_t)pixmap, page_num, param[1], DW_POINTER_TO_INT(param[2]));
    }
    return 0;
}

static uintptr_t go_print_new(char *jobname, unsigned long flags, unsigned int pages, void *drawfunc, void *drawdata, unsigned int sflags)
{
    void **param = malloc(sizeof(void *) * 3);
   
    if(param && drawfunc)
    {
       param[0] = drawfunc;
       param[1] = drawdata;
       param[2] = DW_UINT_TO_POINTER(sflags);
       return (uintptr_t)dw_print_new(jobname, flags, pages, DW_SIGNAL_FUNC(go_callback_print), param);
    }
    return 0;
}

static int go_timer_connect(int interval, void *sigfunc, void *data, unsigned int flags)
{
   void **param = malloc(sizeof(void *) * 3);
   
   if(param && sigfunc)
   {
      param[0] = sigfunc;
      param[1] = data;
      param[2] = DW_UINT_TO_POINTER(flags);
      return dw_timer_connect(interval, DW_SIGNAL_FUNC(go_callback_timer), param);
   }
   return 0;
}

static void DWSIGNAL go_signal_free(HWND window, void *data)
{
   if(data)
   {
      void **param = (void **)data;
      go_callback_remove(param[0]);
      free(data);
   }
}

static void go_signal_connect(uintptr_t window, char *signame, void *sigfunc, void *data, unsigned int flags)
{
   void **param = malloc(sizeof(void *) * 3);
   void *func = (void *)go_callback_basic;
   
   if(param && sigfunc)
   {
      param[0] = sigfunc;
      param[1] = data;
      param[2] = DW_UINT_TO_POINTER(flags);
      
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
         func = (void *)go_callback_item_enter;
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
         func = (void *)go_callback_notepage;
      }
      else if(strcmp(signame, DW_SIGNAL_TREE_EXPAND) == 0)
      {
         func = (void *)go_callback_tree;
      }
      
      dw_signal_connect_data((HWND)window, signame, func, DW_SIGNAL_FUNC(go_signal_free), param);
   }
}
