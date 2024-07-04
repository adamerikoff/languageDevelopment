#include <stdio.h>
#include <stdlib.h>
#include <editline/readline.h>

int main(int argc, char** argv) {
    puts("LispInterpreter Version 0.0.0.0.1");
    puts("Press Ctrl+c to Exit\n\n");

    while (1) {
        char* input = readline("|===> ");
        add_history(input);

        /* Echo input back to user */
        printf("No you're a %s\n\n", input);
        free(input);
    }
    return 0;
}
