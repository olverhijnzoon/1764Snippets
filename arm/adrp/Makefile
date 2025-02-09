AS = as
LD = ld

ASM_FLAGS =
LD_FLAGS =
TARGET = adrpX
SOURCE = adrp.s
OBJECT = adrp.o

all: $(TARGET)

$(OBJECT): $(SOURCE)
	$(AS) $(ASM_FLAGS) $(SOURCE) -o $(OBJECT)

$(TARGET): $(OBJECT)
	$(LD) $(LD_FLAGS) $(OBJECT) -o $(TARGET)

clean:
	rm -f $(TARGET) $(OBJECT)

debug: $(TARGET)
	lldb $(TARGET)
