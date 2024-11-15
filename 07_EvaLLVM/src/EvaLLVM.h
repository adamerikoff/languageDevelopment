#ifndef EVALLVM_h
#define EVALLVM_h

#include "includes.h"
// #include "./Environment.h"
#include "./Logger.h"
#include "./parser/EvaParser.h"

using syntax::EvaParser;

class EvaLLVM {
private:
    std::unique_ptr<EvaParser> parser;
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

    void compile(const Exp& ast) {
        fn = createFunction("main", llvm::FunctionType::get(builder->getInt32Ty(), false));
        gen(ast);
        builder->CreateRet(builder->getInt32(0));
    }

    llvm::Value* gen(const Exp& exp) {
        std::regex re;
        std::string str;
        switch (exp.type) {
            case ExpType::NUMBER:
                return builder->getInt32(exp.number);
            case ExpType::STRING:
                re = std::regex("\\\\n");
                str = std::regex_replace(exp.string, re, "\\n");
                return builder->CreateGlobalStringPtr(str);
            case ExpType::SYMBOL:
                return builder->getInt32(0);
            case ExpType::LIST:
                auto tag = exp.list[0];
                if (tag.type == ExpType::SYMBOL) {
                    auto op = tag.string;
                    if (op == "printf") {
                        auto printfFn = module->getFunction("printf");
                        std::vector<llvm::Value*> args{};
                        for (auto i = 1; i < exp.list.size(); i++) {
                            args.push_back(gen(exp.list[i]));
                        }
                        return builder->CreateCall(printfFn, args);
                    }
                }
        }        
        return builder->getInt32(0);
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

    void setupExternFunctions() {
        auto bytePtrTy = builder->getInt8Ty()->getPointerTo();
        module->getOrInsertFunction("printf", llvm::FunctionType::get(builder->getInt32Ty(), bytePtrTy, true));
    }
public: 
    EvaLLVM() : parser(std::make_unique<EvaParser>()){
        moduleInit();
        setupExternFunctions();
    }

    void exec(const std::string& program) {
        auto ast = parser->parse(program);
        compile(ast);
        module->print(llvm::outs(), nullptr);
        std::cout << "\n";
        saveModuleToFile("./out.ll");
    }

};

#endif