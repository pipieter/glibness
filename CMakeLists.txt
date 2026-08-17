cmake_minimum_required(VERSION 3.25)
project(glibness)

set(GO_SOURCES interops/glibness.go)
set(GO_LIB glibness.a)

set(TEST_SOURCES ./test/interops/tests.c ./test/interops/test_basic.c ./test/interops/test_parse.c)
set(TEST_EXE glibness_tests)

# Create the Go Glibness library

# TODO: find a cleaner way of doing this
add_custom_target(alwaysbuild ALL)

add_custom_command(
  TARGET alwaysbuild
  DEPENDS ${GO_SOURCES}
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
  COMMAND go build -buildmode=c-archive -o ${CMAKE_CURRENT_BINARY_DIR} ${GO_SOURCES})

# Create the Glibness C tests
add_executable(${TEST_EXE} ${TEST_SOURCES})
add_dependencies(${TEST_EXE} alwaysbuild)

target_include_directories(${TEST_EXE} PUBLIC "${CMAKE_CURRENT_LIST_DIR}/include")
target_link_libraries(${TEST_EXE} "${CMAKE_CURRENT_BINARY_DIR}/${GO_LIB}")
