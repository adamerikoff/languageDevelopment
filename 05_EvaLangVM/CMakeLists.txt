cmake_minimum_required(VERSION 3.31.0)
project(EvaLangVM VERSION 1.0 LANGUAGES CXX)

# Set the C++ standard to C++23
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Ensure the correct macOS SDK is used
set(CMAKE_OSX_SYSROOT /Library/Developer/CommandLineTools/SDKs/MacOSX.sdk)

# Define source directories
set(SRC_DIR ${CMAKE_SOURCE_DIR}/src)
set(VM_DIR ${SRC_DIR}/vm)
set(BC_DIR ${SRC_DIR}/bytecode)
set(PR_DIR ${SRC_DIR}/parser)
set(CM_DIR ${SRC_DIR}/compiler)

# Include directories
include_directories(${VM_DIR})
include_directories(${BC_DIR})
include_directories(${PR_DIR})

# Add executable target
add_executable(EvaLangVM
    ${SRC_DIR}/main.cpp
    ${VM_DIR}/EvaVM.cpp
    ${CM_DIR}/EvaCompiler.cpp
)

# Set up headers for IDEs and proper dependency tracking
target_include_directories(EvaLangVM PRIVATE ${VM_DIR} ${BC_DIR} ${PR_DIR})

# Set compiler options
target_compile_options(EvaLangVM PRIVATE -Wall)
