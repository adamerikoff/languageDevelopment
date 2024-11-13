#ifndef EVACOMPILER_H
#define EVACOMPILER_H

#include "../includes.h"
#include "../parser/EvaParser.h"
#include "../vm/EvaValue.h"
#include "../bytecode/OpCode.h"

class EvaCompiler {
private:
    CodeObject* codeObject;
public:
    EvaCompiler();

    CodeObject* compile(const Exp& exp);
    void generate(const Exp& exp);
    void emit(uint8_t code);
    size_t numericConstID(double value);
    size_t stringConstID(const std::string value);
    size_t booleanConstID(const bool value);

    static std::map<std::string, uint8_t> compareOps;
};

#endif