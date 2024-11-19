#ifndef CLOX_VALUE_H
#define CLOX_VALUE_H

#include "common.h"
#include "memory.h"
#include "value.h"

typedef double Value;

typedef struct {
    int capacity;
    int count;
    Value* values;
} ValueArray;

void initValueArray(ValueArray* array);
void writeValueArray(ValueArray* array, Value value);
void freeValueArray(ValueArray* array);
void printValue(Value value);

#endif
