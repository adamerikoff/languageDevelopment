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

    double asNumber() const {
        if (type == EvaValueType::NUMBER) {
            return number;
        }
        std::cerr << "Error: Attempted to access non-number value as a number." << std::endl;
        return 0.0;  // Return a default value (could throw an exception instead)
    }
};


#endif