# compiler flags
CC = gcc
CFLAGS = -Wall -g -Iinclude

C_SOURCES = $(wildcard src/*.c)
ZIG_SOURCES = $(wildcard src/*.zig)
C_OBJECTS = $(C_SOURCES:.c=.o)
ZIG_OBJECTS = $(ZIG_SOURCES:.zig=.o)
EXECUTABLE = calendar_event_creator
HEADER = $(wildcard include/**/*.h)

all: $(EXECUTABLE)

$(EXECUTABLE): $(C_OBJECTS) $(ZIG_OBJECTS)
	$(CC) $(CFLAGS) $(C_OBJECTS) $(ZIG_OBJECTS) -o $@

# create .o from corresponding .c
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

%.o: %.zig
	cd $(dir $<) && zig build-obj -I ../include -O Debug $(notdir $<)

run: $(EXECUTABLE)
	@ mkdir -p tmp/ics_files && \
	ICS_OUTPUT_DIR="tmp/ics_files" \
	./$(EXECUTABLE)

clean:
	rm -f $(C_OBJECTS) $(ZIG_OBJECTS) $(EXECUTABLE)

format:
	clang-format -i $(C_SOURCES) $(HEADER)
	zig fmt $(ZIG_SOURCES)
