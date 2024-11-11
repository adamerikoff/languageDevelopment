#include "EvaVM.h"

EvaVM::EvaVM(/* args */) {
}

EvaVM::~EvaVM() {
}

void EvaVM::exec(const std::string &program) {

    code = {
        OP_HALT
    };

    ip = &code[0];

    return eval();
}

void EvaVM::eval() {
    for (;;) {
        uint8_t opcode = read_byte();
        LOG(opcode);
        switch (opcode) {
        case OP_HALT:
            return ;
            break;
        default:
            DIE << "UNKNOWN OPCODE: " << std::hex << opcode << std::dec << std::endl;
            break;
        }
    }
}

uint8_t EvaVM::read_byte() {
    return *ip++;
}
