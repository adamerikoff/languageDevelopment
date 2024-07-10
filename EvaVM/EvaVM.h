//
// Created by Adam Erik on 11.07.2024.
//

#ifndef EVAVM_H
#define EVAVM_H

#include <string>

class EvaVM {
    public:
        EvaVM() {}

        void exec(const std::string &program) {
            ip = &code[0];
            return eval();
        }

        void eval() {

        }

        uint8_t* ip;
        std::vector<uint8_t> code;
};

#endif //EVAVM_H
