#include "vm/EvaVM.h"

int main(int argc, char const *argv[]) {
    EvaVM vm;

    vm.exec(R"(
        42
    )");

    std::cout << "ALL DONE!" << std::endl;

    return 0;
}
