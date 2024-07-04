# Define the compiler
CC = cc

# Define compiler flags
CFLAGS = -std=c99 -Wall mpc.c -ledit

# Define the target executable
TARGET = a.out

# Define the source file
SOURCE = parser.c

# Default target to build the executable
all: $(TARGET)

# Rule to compile the source file into the target executable
$(TARGET): $(SOURCE)
	$(CC) $(CFLAGS) -o $@ $<

# Rule to clean up the executable
clean:
	rm -f $(TARGET)

# Rule to run the program
run: $(TARGET)
	./$(TARGET)

# Combined target to clean, compile, and run
start: clean all run

# Phony targets
.PHONY: all clean run build_and_run
