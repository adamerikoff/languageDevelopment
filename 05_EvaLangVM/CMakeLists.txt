cmake_minimum_required(VERSION 3.31.0)
project(EvaLangVM VERSION 1.0 LANGUAGES CXX)

# Set the C++ standard to C++23
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Define source directories
set(SRC_DIR ${CMAKE_SOURCE_DIR}/src)
set(VM_DIR ${SRC_DIR}/vm)

# Include directories
include_directories(${VM_DIR})

# Add executable target
add_executable(EvaLangVM
    ${SRC_DIR}/main.cpp
    ${VM_DIR}/EvaVM.cpp
    ${VM_DIR}/Logger.cpp
)

# Set up headers for IDEs and proper dependency tracking
target_include_directories(EvaLangVM PRIVATE ${VM_DIR})

# Specify additional files to be used as header dependencies
target_sources(EvaLangVM PRIVATE
    ${VM_DIR}/EvaVM.h
    ${VM_DIR}/Logger.h
    ${VM_DIR}/OpCode.h
    ${VM_DIR}/includes.h
)

target_compile_options(EvaLangVM PRIVATE -Wall -Wextra -pedantic)
