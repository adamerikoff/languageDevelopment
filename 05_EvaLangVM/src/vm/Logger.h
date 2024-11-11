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

#define LOG_OPCODE(value) std::cout << #value << " = 0x" << std::setw(2) << std::setfill('0') << std::hex << (int)(value) << std::endl;

#define LOG(value) std::cout << #value << " = " << value << std::endl;

#endif
