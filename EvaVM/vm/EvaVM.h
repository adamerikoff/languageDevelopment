//
// Created by Adam Erik on 11.07.2024.
//

#ifndef EVAVM_H
#define EVAVM_H

#include <string>
#include <vector>

#include "../bytecode/OpCode.h"

#define READ_BYTE() *ip++

class EvaVM {
    public:
        uint8_t* ip;
        std::vector<uint8_t> code;

        EvaVM() {}


        void exec(const std::string &program) {
            code = {OP_HALT};
            ip = &code[0];
            return eval();
        }

        void eval() {
            for(;;) {
                switch (READ_BYTE()) {
                    case OP_HALT:
                        return;
                }
            }
        }


};

#endif //EVAVM_H
