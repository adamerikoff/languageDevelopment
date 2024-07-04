#include <stdio.h>
#include <stdlib.h>

#include <editline/readline.h>

#include "mpc.h"

int main(int argc, char** argv) {
    /* Create Some Parsers */
    mpc_parser_t* Number     = mpc_new("number");
    mpc_parser_t* Operator   = mpc_new("operator");
    mpc_parser_t* Expression = mpc_new("expression");
    mpc_parser_t* Lisp       = mpc_new("lisp");

    /* Define them with the following Language */
    mpca_lang(MPCA_LANG_DEFAULT,
        "                                                      \
        number : /-?[0-9]+/;                                   \
        operator : '+' | '-' | '*' | '/' ;                     \
        expression : <number> | '(' <operator> <expr>+ ')' ;   \
        lisp : /^/ <operator> <expr>+ /$/ ;                    \
        ",
        Number, Operator, Expression, Lisp);

    puts("LispInterpreter Version 0.0.0.0.1");
    puts("Press Ctrl+c to Exit\n\n");

    while (1) {
        char* input = readline("lisp> ");
        add_history(input);

        /* Attempt to parse the user input */
        mpc_result_t r;
        if (mpc_parse("<stdin>", input, Lisp, &r)) {
        /* On Success Print the AST */
        mpc_ast_print(r.output);
        mpc_ast_delete(r.output);
        } else {
        /* Otherwise Print the Error */
        mpc_err_print(r.error);
        mpc_err_delete(r.error); }
        free(input);
    }
    mpc_cleanup(4, Number, Operator, Expression, Lisp);
    return 0;
}
