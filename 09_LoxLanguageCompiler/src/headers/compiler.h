#ifndef CLOX_COMPILER_H
#define CLOX_COMPILER_H

#include "common.h"
#include "scanner.h"
#include "vm.h"

typedef struct {
    Token current;
    Token previous;
    bool hadError;
    bool panicMode;
} Parser;

typedef enum {
    PREC_NONE,        // No precedence
    PREC_ASSIGNMENT,  // =
    PREC_OR,          // or
    PREC_AND,         // and
    PREC_EQUALITY,    // ==, !=
    PREC_COMPARISON,  // <, >, <=, >=
    PREC_TERM,        // +, -
    PREC_FACTOR,      // *, /
    PREC_UNARY,       // !, -
    PREC_CALL,        // . ()
    PREC_PRIMARY      // Primary expressions
} Precedence;

typedef void (*ParseFn)();

typedef struct {
    ParseFn prefix;
    ParseFn infix;
    Precedence precedence;
} ParseRule;

bool compile(const char* source, Chunk* chunk);


#ifdef DEBUG_PRINT_CODE
    #include "debug.h"
#endif

#endif
