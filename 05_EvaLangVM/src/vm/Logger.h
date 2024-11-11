#ifndef LOGGER_H
#define LOGGER_H

#include "includes.h"

class ErrorLogMessage : public std::basic_ostringstream<char> {
public:
    ~ErrorLogMessage() {
        fprintf(stderr, "Fatal error: %s\n", str().c_str());
        exit(EXIT_FAILURE);
    }
};

#define DIE ErrorLogMessage()
#define LOG(value) std::cout << #value << " = 0x" << std::setw(2) << std::setfill('0') << std::hex << (int)(value) << std::endl;

#endif
