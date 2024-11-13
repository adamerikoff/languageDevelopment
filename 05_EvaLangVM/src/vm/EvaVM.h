#ifndef EVAVM_H
#define EVAVM_H

#include "../includes.h"

#include "../bytecode/OpCode.h"
#include "../parser/EvaParser.h"
#include "../logger/Logger.h"
#include "../compiler/EvaCompiler.h"
#include "./EvaValue.h"

using syntax::EvaParser;

class EvaVM {    
private:
    std::unique_ptr<EvaParser> parser;
    std::unique_ptr<EvaCompiler> compiler;
    uint8_t instruction_index;
    std::vector<EvaValue> stack;
    CodeObject* codeObject;
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