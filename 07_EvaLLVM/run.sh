#!/bin/bash

LLVM_CONFIG="/opt/homebrew/Cellar/llvm/19.1.3/bin/llvm-config"

# Create the build directory if it doesn't exist
if [ ! -d "build" ]; then
    mkdir build
fi

clang++ -o ./build/eva-llvm -target arm64-apple-darwin `$LLVM_CONFIG --cxxflags --ldflags --system-libs --libs all` main.cpp

cd build

./eva-llvm
