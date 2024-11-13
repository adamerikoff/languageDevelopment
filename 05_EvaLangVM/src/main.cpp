#include "vm/EvaVM.h"

int main(int argc, char const *argv[]) {
    (void)argc;
    (void)argv;

    EvaVM vm;

    auto result = vm.exec(R"(
        "Hello"
    )");

    LOG(AS_CPPSTRING(result));

    std::cout << "ALL DONE!" << std::endl;

    return 0;
}
