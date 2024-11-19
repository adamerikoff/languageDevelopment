#include "./headers/common.h"
#include "./headers/chunk.h"
#include "./headers/debug.h"

int main(int argc, const char* argv[]) {
    (void)argc;
    (void)argv;

    Chunk chunk;
    initChunk(&chunk);
    writeChunk(&chunk, OP_RETURN);
    disassembleChunk(&chunk, "test chunk");
    freeChunk(&chunk);
    return 0;
}
