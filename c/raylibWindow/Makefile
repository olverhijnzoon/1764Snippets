RAYLIB_PATH=/opt/homebrew/opt/raylib
OUTPUT_DIR = tmp
CC = clang
CFLAGS = -I$(RAYLIB_PATH)/include
LDFLAGS = -L$(RAYLIB_PATH)/lib -lraylib -framework OpenGL -framework Cocoa -lm
OUTPUT = $(OUTPUT_DIR)/raylibWindow
SOURCES = raylibWindow.c

all: $(SOURCES) prep
	$(CC) $(CFLAGS) $(SOURCES) $(LDFLAGS) -o $(OUTPUT)

run: all
	./$(OUTPUT)

clean:
	rm -rf $(OUTPUT_DIR)

prep:
	mkdir -p $(OUTPUT_DIR)
