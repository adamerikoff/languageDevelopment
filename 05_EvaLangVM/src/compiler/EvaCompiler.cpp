#include "EvaCompiler.h"

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
    for (auto i = 0; i < this->codeObject->constants.size(); i++) {
        if (!IS_NUMBER(this->codeObject->constants[i])) {
            continue;
        }
        if (AS_NUMBER(this->codeObject->constants[i]) == value) {
            return i;
        }
    }
    this->codeObject->constants.push_back(NUMBER(value));
    return this->codeObject->constants.size() - 1;
}

size_t EvaCompiler::stringConstID(const std::string value) {
    for (auto i = 0; i < this->codeObject->constants.size(); i++) {
        if (!IS_STRING(this->codeObject->constants[i])) {
            continue;
        }
        if (AS_CPPSTRING(this->codeObject->constants[i]) == value) {
            return i;
        }
    }
    this->codeObject->constants.push_back(ALLOC_STRING(value));
    return this->codeObject->constants.size() - 1;
}