#include <stdio.h>

int main() {
    
    printf("1764Snippets\n");
    printf("LLDBDebugger\n");

    // points to memory location 0, which is not accessible
	int *ptr = NULL;
    printf("This will cause a segmentation fault\n");
    *ptr = 0;
    return 0;
    
}
