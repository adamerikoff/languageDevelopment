#ifndef EVAVALUES_H
#define EVAVALUES_H

#include "includes.h"

enum class EvaValueType {
    NUMBER,
};


struct EvaValue {
    EvaValueType type;
    union {
        double number;
    };

    EvaValue() : type(EvaValueType::NUMBER), number(0.0) {}
    
    EvaValue(double value){
        type = EvaValueType::NUMBER;
        number = value;
    }
};


#endif