#ifndef EVAVM_H
#define EVAVM_H

#include "includes.h"

#include "OpCode.h"
#include "Logger.h"

class EvaVM {
private:
    uint8_t* ip;
    std::vector<uint8_t> code;
public:
    EvaVM(/* args */);
    ~EvaVM();
    void exec(const std::string &program);
    void eval();
    uint8_t read_byte();
};

#endif