#include "EvaVM.h"

EvaVM::EvaVM() {
    stack_pointer = stack.begin();
}

EvaVM::~EvaVM() {
}

EvaValue EvaVM::exec(const std::string &program) {
    (void)program;

    constants.push_back(EvaValue(11));
    constants.push_back(EvaValue(22));

    code = {
        OP_CONST,
        0,
        OP_CONST,
        1,
        OP_ADD,
        OP_HALT,
    };

    instruction_pointer = &code[0];
    stack_pointer = &stack[0];

    return eval();
}

EvaValue EvaVM::eval() {
    for (;;) {
        uint8_t opcode = read_byte();
        LOG_OPCODE(opcode);
        switch (opcode) {
            case OP_HALT: {
                std::cout << "OP_HALT CODE" << std::endl;
                return pop();
                break;
            }
            case OP_CONST: {
                std::cout << "OP_CONST CODE" << std::endl;
                push(get_const());
                break;
            }
            case OP_ADD: {
                std::cout << "OP_ADD CODE" << std::endl;
                double op1 = pop().asNumber();
                double op2 = pop().asNumber();
                double result = op1 + op2;
                push(EvaValue(result));
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
    return *instruction_pointer++;
}

void EvaVM::push(const EvaValue& value) {
    if ((size_t)(stack_pointer - stack.begin()) == STACK_LIMIT) {
        DIE << "push(): Stack overflow!" << std::endl;
    }
    *stack_pointer = value;
    stack_pointer++;
}

EvaValue EvaVM::pop() {
    if (stack_pointer == stack.begin()) {
        DIE << "pop(): Stack empty!" << std::endl;
    }
    --stack_pointer;
    return *stack_pointer;
}

EvaValue EvaVM::get_const() {
    uint8_t constIndex = read_byte();
    
    if (constIndex >= constants.size()) {
        DIE << "get_const(): Invalid constant index!" << std::endl;
    }

    EvaValue constant = constants[constIndex];
    return constant;
}