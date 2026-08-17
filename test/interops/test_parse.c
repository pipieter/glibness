#include <assert.h>
#include <stdint.h>

#include <glibness/glibness.h>

void test_basic_parsing() {
    uintptr_t engine = glib_new_engine();
    assert(engine);

    const char* script =
        "dialogue test {\n"
        "   set speaker \"Glibness\"\n"
        "   say \"This is a test message.\"\n"
        "   say \"This is another test message.\"\n"
        "}";

    assert(glib_load_string(engine, script));
    assert(glib_start(engine, "test"));
    glib_free_engine(engine);
}

void test_missing_closing_bracket() {
    uintptr_t engine = glib_new_engine();
    assert(engine);

    const char* script =
        "dialogue test {\n"
        "   set speaker \"Glibness\"\n"
        "   say \"This is a test message.\"\n"
        "   say \"This is another test message.\"\n";

    assert(!glib_load_string(engine, script));
}
