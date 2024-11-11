#include "vm/EvaVM.h"

int main(int argc, char const *argv[]) {
    (void)argc;
    (void)argv;

    EvaVM vm;

    auto result = vm.exec(R"(
        42
    )");

    LOG(result.asCPPString());

    std::cout << "ALL DONE!" << std::endl;

    return 0;
}
