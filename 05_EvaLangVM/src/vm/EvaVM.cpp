#include "EvaVM.h"

#define COMPARE_VALUES(op, v1, v2)  \
do {                                \
    bool res;                       \
    switch (op) {                   \
    case 0:                         \
        res = v1 < v2;              \
        break;                      \
    case 1:                         \
        res = v1 > v2;              \
        break;                      \
    case 2:                         \
        res = v1 == v2;             \
        break;                      \
    case 3:                         \
        res = v1 >= v2;             \
        break;                      \
    case 4:                         \
        res = v1 <= v2;             \
        break;                      \
    case 5:                         \
        res = v1 != v2;             \
        break;                      \
    }                               \
    push(BOOLEAN(res));             \
} while (false)

EvaVM::EvaVM() {
    this->parser = std::make_unique<EvaParser>();
    this->compiler = std::make_unique<EvaCompiler>();
    this->instruction_index = 0;
}

EvaVM::~EvaVM() {}

EvaValue EvaVM::exec(const std::string &program) {
    auto ast = parser->parse(program);
    this->codeObject = compiler->compile(ast);
    this->instruction_index = 0;
    return this->eval();
}

EvaValue EvaVM::eval() {
    for (;;) {
        uint8_t opcode = this->read_byte();
        switch (opcode) {
            case OP_HALT: {
                std::cout << "code: OP_HALT" << std::endl;
                return this->pop();
            }
            case OP_CONST: {
                std::cout << "code: OP_CONST" << std::endl;
                this->push(this->get_const());
                break;
            }
            case OP_ADD: {
                std::cout << "code: OP_ADD" << std::endl;
                this->binaryOperation("+");
                break;
            }
            case OP_SUB: {
                std::cout << "code: OP_SUB" << std::endl;
                this->binaryOperation("-");
                break;
            }
            case OP_MUL: {
                std::cout << "code: OP_MUL" << std::endl;
                this->binaryOperation("*");
                break;
            }
            case OP_DIV: {
                std::cout << "code: OP_DIV" << std::endl;
                this->binaryOperation("/");
                break;
            }
            case OP_COMPARE: {
                std::cout << "code: OP_COMPARE" << std::endl;
                uint8_t op = this->read_byte();
                EvaValue op2 = this->pop();
                EvaValue op1 = this->pop();
                if (IS_NUMBER(op1) && IS_NUMBER(op2)) {
                    double v1 = AS_NUMBER(op1);
                    double v2 = AS_NUMBER(op2);
                    COMPARE_VALUES(op, v1, v2);
                } else if (IS_STRING(op1) && IS_STRING(op2)) {
                    std::string s1 = AS_CPPSTRING(op1);
                    std::string s2 = AS_CPPSTRING(op2);
                    COMPARE_VALUES(op, s1, s2);
                }
                break;
            }
            default: {
                DIE << "UNKNOWN code: " << std::hex << opcode << std::dec << std::endl;
                break;
            }
        }
    }
}

uint8_t EvaVM::read_byte() {
    if (instruction_index >= this->codeObject->code.size()) {
        DIE << "read_byte(): Out of code bounds!" << std::endl;
    }
    return this->codeObject->code[instruction_index++];
}

void EvaVM::push(const EvaValue& value) {
    this->stack.push_back(value);
}

EvaValue EvaVM::pop() {
    if (this->stack.empty()) {
        DIE << "pop(): Stack empty!" << std::endl;
    }
    EvaValue value = this->stack.back();
    this->stack.pop_back();
    return value;
}

EvaValue EvaVM::get_const() {
    uint8_t constIndex = this->read_byte();
    if (constIndex >= this->codeObject->constants.size()) {
        DIE << "get_const(): Invalid constant index!" << std::endl;
    }
    return this->codeObject->constants[constIndex];
}

void EvaVM::binaryOperation(const char* op) {
    EvaValue op2 = this->pop();
    EvaValue op1 = this->pop();
    
    EvaValue result;

    if (strcmp(op, "+") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            result = NUMBER(AS_NUMBER(op1) + AS_NUMBER(op2));
        } else if (IS_STRING(op1) && IS_STRING(op2)) {
            result = ALLOC_STRING(AS_CPPSTRING(op1) + AS_CPPSTRING(op2));
        }
    }
    else if (strcmp(op, "-") == 0 && IS_NUMBER(op1) && IS_NUMBER(op2)) {
        result = NUMBER(AS_NUMBER(op1) - AS_NUMBER(op2));
    }
    else if (strcmp(op, "*") == 0 && IS_NUMBER(op1) && IS_NUMBER(op2)) {
        result = NUMBER(AS_NUMBER(op1) * AS_NUMBER(op2));
    }
    else if (strcmp(op, "/") == 0 && IS_NUMBER(op1) && IS_NUMBER(op2)) {
        if (AS_NUMBER(op2) == 0.0) {
            DIE << "Error: Division by zero!" << std::endl;
            return;
        }
        result = NUMBER(AS_NUMBER(op1) / AS_NUMBER(op2));
    }
    else {
        DIE << "Error: Unsupported operator " << op << std::endl;
        return;
    }

    this->push(result);
}

