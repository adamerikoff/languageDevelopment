#include "EvaCompiler.h"

#define ALLOC_CONST(tester, converter, allocator, value)                 \
    do {                                                                 \
        for (auto i = 0; i < this->codeObject->constants.size(); i++) {  \
            if (!tester(this->codeObject->constants[i])) {               \
                continue;                                                \
            }                                                            \
            if (converter(this->codeObject->constants[i]) == value) {    \
                return i;                                                \
            }                                                            \
        }                                                                \
        this->codeObject->constants.push_back(allocator(value));         \
        return this->codeObject->constants.size() - 1;                   \
    } while (false)

#define GEN_BINARY_OP(op)      \
    do {                       \
        generate(exp.list[1]); \
        generate(exp.list[2]); \
        emit(op);              \
    } while (false)

std::map<std::string, uint8_t> EvaCompiler::compareOps = {
    {"<", 0},
    {">", 1},
    {"==", 2},
    {">=", 3},
    {"<=", 4},
    {"!=", 5},
};

EvaCompiler::EvaCompiler() {
}

CodeObject* EvaCompiler::compile(const Exp& exp) {
    this->codeObject = AS_CODE(ALLOC_CODE("main"));
    this->generate(exp);
    this->emit(OP_HALT);
    return this->codeObject;
}

void EvaCompiler::generate(const Exp& exp) {
    switch (exp.type) {
        case ExpType::NUMBER: {
            this->emit(OP_CONST);
            this->emit(this->numericConstID(exp.number));
            break;
        }
        case ExpType::STRING: {
            this->emit(OP_CONST);
            this->emit(this->stringConstID(exp.string));
            break;
        }
        case ExpType::SYMBOL: {
            if (exp.string == "true" || exp.string == "false") {
                this->emit(OP_CONST);
                this->emit(booleanConstID(exp.string == "true" ? true : false));
            }
            
            break;
        }
        case ExpType::LIST: {
            Value tag = exp.list[0];
            if (tag.type == ExpType::SYMBOL) {
                std::string op = tag.string;

                if (op == "+") {
                    GEN_BINARY_OP(OP_ADD);
                }
                if (op == "-") {
                    GEN_BINARY_OP(OP_SUB);
                }
                if (op == "*") {
                    GEN_BINARY_OP(OP_MUL);
                }
                if (op == "/") {
                    GEN_BINARY_OP(OP_DIV);
                }
                if (compareOps.count(op) != 0) {
                    generate(exp.list[1]);
                    generate(exp.list[2]);
                    emit(OP_COMPARE);
                    emit(compareOps[op]);
                }
            }
            break;
        }
    }
}

void EvaCompiler::emit(uint8_t code) {
    this->codeObject->code.push_back(code);
}

size_t EvaCompiler::numericConstID(double value) {
    ALLOC_CONST(IS_NUMBER, AS_NUMBER, NUMBER, value);
    return this->codeObject->constants.size() - 1;
}

size_t EvaCompiler::stringConstID(const std::string value) {
    ALLOC_CONST(IS_STRING, AS_CPPSTRING, ALLOC_STRING, value);
    return this->codeObject->constants.size() - 1;
}

size_t EvaCompiler::booleanConstID(const bool value) {
    ALLOC_CONST(IS_BOOLEAN, AS_BOOLEAN, BOOLEAN, value);
    return this->codeObject->constants.size() - 1;
}