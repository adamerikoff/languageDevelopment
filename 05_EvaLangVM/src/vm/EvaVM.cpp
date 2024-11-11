#include "EvaVM.h"

EvaVM::EvaVM() {
    this->instruction_index = 0;
}

EvaVM::~EvaVM() {}

EvaValue EvaVM::exec(const std::string &program) {
    (void)program;

    this->constants.push_back(EvaValue(33));
    this->constants.push_back(EvaValue(333));

    this->code = {
        OP_CONST,
        0,
        OP_CONST,
        1,
        OP_ADD,
        OP_HALT,
    };

    this->instruction_index = 0;

    return this->eval();
}

EvaValue EvaVM::eval() {
    for (;;) {
        uint8_t opcode = this->read_byte();
        LOG_OPCODE(opcode);
        
        switch (opcode) {
            case OP_HALT: {
                std::cout << "OP_HALT CODE" << std::endl;
                return this->pop();
            }
            case OP_CONST: {
                std::cout << "OP_CONST CODE" << std::endl;
                this->push(this->get_const());
                break;
            }
            case OP_ADD: {
                std::cout << "OP_ADD CODE" << std::endl;
                this->binaryOperation("+");
                break;
            }
            case OP_SUB: {
                std::cout << "OP_SUB CODE" << std::endl;
                this->binaryOperation("-");
                break;
            }
            case OP_MUL: {
                std::cout << "OP_MUL CODE" << std::endl;
                this->binaryOperation("*");
                break;
            }
            case OP_DIV: {
                std::cout << "OP_DIV CODE" << std::endl;
                this->binaryOperation("/");
                break;
            }
            case OP_SQR: {
                std::cout << "OP_SQR CODE" << std::endl;
                this->binaryOperation("**");
                break;
            }
            default: {
                DIE << "UNKNOWN OPCODE: " << std::hex << opcode << std::dec << std::endl;
                break;
            }
        }
    }
}

uint8_t EvaVM::read_byte() {
    if (instruction_index >= this->code.size()) {
        DIE << "read_byte(): Out of code bounds!" << std::endl;
    }

    return this->code[instruction_index++];
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

    if (constIndex >= this->constants.size()) {
        DIE << "get_const(): Invalid constant index!" << std::endl;
    }

    EvaValue constant = this->constants[constIndex];
    return constant;
}

void EvaVM::binaryOperation(const char* op) {
    EvaValue op2 = this->pop();
    EvaValue op1 = this->pop();
    
    EvaValue result;

    if (strcmp(op, "+") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            result = EvaValue(op1.number + op2.number);
        } 
        else if (op2.type == EvaValueType::OBJECT && op1.type == EvaValueType::OBJECT) {
            if (op2.object->type == ObjectType::STRING && op1.object->type == ObjectType::STRING) {
                result = EvaValue(op1.asCPPString() + op2.asCPPString());
            }
        }
    }
    else if (strcmp(op, "-") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            result = EvaValue(op1.number - op2.number);
        }
    }
    else if (strcmp(op, "*") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            result = EvaValue(op1.number * op2.number);
        }
    }
    else if (strcmp(op, "/") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            if (op2.number == 0.0) {
                DIE << "Error: Division by zero!" << std::endl;
                return;
            }
            result = EvaValue(op1.number / op2.number);
        }
    }
    else if (strcmp(op, "**") == 0) {
        if (op2.type == EvaValueType::NUMBER && op1.type == EvaValueType::NUMBER) {
            result = EvaValue(pow(op1.number, op2.number));
        }
    }
    else {
        DIE << "Error: Unsupported operator " << op << std::endl;
        return;
    }

    this->push(result);
}


