#ifndef GLIBNESS_H_
#define GLIBNESS_H_

#include <stddef.h>
#include <stdint.h>

enum glib_response_code {
    GLIB_RESPONSE_NONE = 0,
    GLIB_RESPONSE_ERROR,
    GLIB_RESPONSE_SAY,
    GLIB_RESPONSE_CHOICE,
    GLIB_RESPONSE_INTERNAL_CHANGE,
    GLIB_RESPONSE_FINISHED
};

extern uintptr_t glib_new_engine();
extern void      glib_free_engine(uintptr_t engine);

extern char glib_load_string(uintptr_t engine, const char* string);
extern char glib_load_file(uintptr_t engine, const char* path);

extern char glib_start(uintptr_t engine, const char* dialogue);
extern int  glib_next(uintptr_t engine);
extern char glib_choose(uintptr_t engine, int index);

extern size_t glib_get_sentence(uintptr_t engine, char* buffer);
extern size_t glib_get_speaker(uintptr_t engine, char* buffer);
extern size_t glib_get_error(uintptr_t engine, char* buffer);
extern char   glib_is_active(uintptr_t engine);

extern size_t glib_get_str(uintptr_t engine, const char* key, char* buffer);
extern char   glib_get_int(uintptr_t engine, const char* key, int* value);
extern char   glib_get_bool(uintptr_t engine, const char* key, char* value);

extern char glib_set_str(uintptr_t engine, const char* key, const char* value);
extern char glib_set_int(uintptr_t engine, const char* key, int value);
extern char glib_set_bool(uintptr_t engine, const char* key, char value);

#endif  // GLIBNESS_H_