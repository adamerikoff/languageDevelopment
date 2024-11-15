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
    llvm::Function* fn;
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
    void compile() {
        fn = createFunction("main", llvm::FunctionType::get(builder->getInt32Ty(), false));
        auto result = gen();
        auto i32Result = builder->CreateIntCast(result, builder->getInt32Ty(), true);
        builder->CreateRet(i32Result);
    }
    llvm::Value* gen() {
        return builder->getInt32(42);
    }
    llvm::Function* createFunction(const std::string& fnName, llvm::FunctionType* fnType) {
        auto fn = module->getFunction(fnName);
        if (fn == nullptr) {
            fn = createFunctionProto(fnName, fnType);
        }
        createFunctionBlock(fn);
        return fn;
    }
    llvm::Function* createFunctionProto(const std::string& fnName, llvm::FunctionType* fnType) {
        auto fn = llvm::Function::Create(fnType, llvm::Function::ExternalLinkage, fnName, *module);
        verifyFunction(*fn);
        return fn;
    }
    void createFunctionBlock(llvm::Function* fn) {
        auto entry = createBB("entry", fn);
        builder->SetInsertPoint(entry);
    }
    llvm::BasicBlock* createBB(std::string name, llvm::Function* fn = nullptr) {
        return llvm::BasicBlock::Create(*ctx, name, fn);
    }
public: 
    EvaLLVM() {
        moduleInit();
    }

    void exec(const std::string& program) {
        // auto ast = parser->parser(program)
        // compile(ast);
        compile();
        module->print(llvm::outs(), nullptr);
        saveModuleToFile("./out.ll");
    }

};

#endif