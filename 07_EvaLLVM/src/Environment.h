#ifndef ENVIRONMENT_H
#define ENVIRONMENT_H

#include "includes.h"
#include "./Logger.h"


class Environment : public std::enable_shared_from_this<Environment> {
public:
    Environment(std::map<std::string, llvm::Value*> record, std::shared_ptr<Environment> parent) : record_(record), parent_(parent) {

    }

    llvm::Value* define(const std::string& name, llvm::Value* value) {

    }

    llvm::Value* lookup(const std::string& name) {
        return resolve(name)->record_[name];
    }

private:
    std::map<std::string, llvm::Value*> record_;
    std::shared_ptr<Environment> parent_;
    std::shared_ptr<Environment> resolve(const std::string& name) {
        
    }
};

#endif