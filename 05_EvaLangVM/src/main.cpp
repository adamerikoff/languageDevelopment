#include "vm/EvaVM.h"

int main(int argc, char const *argv[]) {
    (void)argc;
    (void)argv;

    EvaVM vm;

    auto result = vm.exec(R"(
        42
    )");

    LOG(AS_NUMBER(result));

    std::cout << "ALL DONE!" << std::endl;

    return 0;
}