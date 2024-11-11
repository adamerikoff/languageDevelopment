#ifndef EVAVM_H
#define EVAVM_H

#include "includes.h"

#include "OpCode.h"
#include "Logger.h"
#include "EvaValue.h"

#define STACK_LIMIT 512

class EvaVM {
private:
    uint8_t instruction_index;
    std::vector<uint8_t> code;
    std::vector<EvaValue> constants;
    std::vector<EvaValue> stack;
public:
    EvaVM();
    ~EvaVM();
    uint8_t read_byte();
    
    EvaValue exec(const std::string &program);
    EvaValue eval();

    void push(const EvaValue&);
    EvaValue pop();
    
    EvaValue get_const();

    void binaryOperation(const char* op);
};

#endif