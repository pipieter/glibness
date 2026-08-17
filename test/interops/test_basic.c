#include <assert.h>
#include <stdint.h>
#include <string.h>

#include <glibness/glibness.h>

void test_basic_execution() {
    uintptr_t engine = glib_new_engine();
    assert(engine);

    char speaker[128];
    char sentence[128];
    char error[128];
    int  my_int;
    char my_bool;

    const char* script =
        "dialogue test {\n"
        "   set speaker \"Glibness\"\n"
        "   say \"This is a message.\"\n"
        "   set my_int 333\n"
        "   set my_bool true\n"
        "}";

    assert(glib_load_string(engine, script));
    assert(glib_start(engine, "test"));
    assert(glib_next(engine) == GLIB_RESPONSE_INTERNAL_CHANGE);
    assert(glib_next(engine) == GLIB_RESPONSE_SAY);
    assert(glib_next(engine) == GLIB_RESPONSE_INTERNAL_CHANGE);
    assert(glib_next(engine) == GLIB_RESPONSE_INTERNAL_CHANGE);

    assert(glib_is_active(engine));
    assert(glib_get_speaker(engine, speaker) == 8);
    assert(glib_get_sentence(engine, sentence) == 18);
    assert(glib_get_error(engine, error) == 0);
    assert(glib_get_int(engine, "my_int", &my_int));
    assert(glib_get_bool(engine, "my_bool", &my_bool));

    assert(strcmp(speaker, "Glibness") == 0);
    assert(strcmp(sentence, "This is a message.") == 0);
    assert(my_int == 333);
    assert(my_bool == 1);

    assert(glib_next(engine) == GLIB_RESPONSE_FINISHED);
    assert(!glib_is_active(engine));

    glib_free_engine(engine);
}

void test_basic_set_get() {
    uintptr_t engine = glib_new_engine();
    assert(engine);

    int   my_set_int  = 1234;
    char* my_set_str  = "abcdef";
    char  my_set_bool = 1;

    int  my_set_int_result;
    char my_set_str_result[10];
    char my_set_bool_result;

    assert(glib_set_int(engine, "my_set_int", my_set_int));
    assert(glib_set_str(engine, "my_set_str", my_set_str));
    assert(glib_set_bool(engine, "my_set_bool", my_set_bool));

    assert(glib_get_int(engine, "my_set_int", &my_set_int_result));
    assert(glib_get_str(engine, "my_set_str", my_set_str_result));
    assert(glib_get_bool(engine, "my_set_bool", &my_set_bool_result));

    assert(my_set_int == my_set_int_result);
    assert(strcmp(my_set_str, my_set_str_result) == 0);
    assert(my_set_bool == my_set_bool_result);

    glib_free_engine(engine);
}

void test_basic_choice() {
    uintptr_t engine = glib_new_engine();
    assert(engine);

    const char* script =
        "dialogue test {\n"
        "  choose {\n"
        "     choice \"A\" {\n"
        "        say \"A.1\"\n"
        "     }\n"
        "     choice \"BB\" {\n"
        "        say \"B.1\"\n"
        "        say \"B.2\"\n"
        "     }\n"
        "     choice \"CCC\" {\n"
        "        say \"C.1\"\n"
        "     }\n"
        "  }\n"
        "}";

    char choiceA[8];
    char choiceB[8];
    char choiceC[8];
    char choiceX[8];

    assert(glib_load_string(engine, script));
    assert(glib_start(engine, "test"));
    assert(glib_next(engine) == GLIB_RESPONSE_CHOICE);

    assert(glib_get_choice_count(engine) == 3);
    assert(glib_get_choice(engine, 0, choiceA));
    assert(glib_get_choice(engine, 1, choiceB));
    assert(glib_get_choice(engine, 2, choiceC));

    assert(!glib_get_choice(engine, 37, choiceX));
    assert(!glib_get_choice(engine, -2, choiceX));

    assert(strcmp(choiceA, "A") == 0);
    assert(strcmp(choiceB, "BB") == 0);
    assert(strcmp(choiceC, "CCC") == 0);

    assert(glib_choose(engine, 0));
    assert(glib_next(engine) == GLIB_RESPONSE_SAY);
    assert(glib_next(engine) == GLIB_RESPONSE_FINISHED);

    glib_free_engine(engine);
}
