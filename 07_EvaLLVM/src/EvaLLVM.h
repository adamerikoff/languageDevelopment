#ifndef EVALLVM_h
#define EVALLVM_h

#include "includes.h"
// #include "./Environment.h"
#include "./Logger.h"
#include "./parser/EvaParser.h"

using syntax::EvaParser;

class EvaLLVM {
private:
    std::unique_ptr<llvm::LLVMContext> ctx;
    std::unique_ptr<llvm::Module> module;
    std::unique_ptr<llvm::IRBuilder<>> builder;
    void moduleInit() {
        ctx = std::make_unique<llvm::LLVMContext>();
        module = std::make_unique<llvm::Module>("EvaLLVM", *ctx);
        builder = std::make_unique<llvm::IRBuilder<>>(*ctx);
    }
    void saveModuleToFile(const std::string& fileName) {
        std::error_code errorCode;
        llvm::raw_fd_ostream ouLL(fileName, errorCode);
        module->print(ouLL, nullptr);
    }
public: 
    EvaLLVM() {
        moduleInit();
    }

    void exec(const std::string& program) {
        // auto ast = parser->parser(program)
        // compile(ast);
        module->print(llvm::outs(), nullptr);
        saveModuleToFile("./out.ll");
    }

};

#endif