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
            break;
        }
        case ExpType::LIST: {
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