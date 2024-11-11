#ifndef EVAVM_H
#define EVAVM_H

#include "includes.h"

#include "OpCode.h"
#include "Logger.h"
#include "EvaValues.h"

#define STACK_LIMIT 512

class EvaVM {
private:
    uint8_t* instruction_pointer;
    EvaValue* stack_pointer;
    std::vector<uint8_t> code;
    std::vector<EvaValue> constants;
    std::array<EvaValue, STACK_LIMIT> stack;
public:
    EvaVM();
    ~EvaVM();
    uint8_t read_byte();
    
    EvaValue exec(const std::string &program);
    EvaValue eval();

    void push(const EvaValue&);
    EvaValue pop();
    
    EvaValue get_const();
};

#endif