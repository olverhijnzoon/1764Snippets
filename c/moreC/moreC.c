#include <stdio.h>

int main() {
    
    printf("1764Snippets\n");
    printf("SymRank2 MoreC\n");

    int matrix[4][4] = {
        {1, 2, 3, 4},
        {2, 5, 6, 7},
        {3, 6, 8, 9},
        {4, 7, 9, 10}
    };

    printf("4x4 Symmetric Matrix:\n");
    for(int i = 0; i < 4; i++) {
        for(int j = 0; j < 4; j++) {
            printf("%d ", matrix[i][j]);
        }
        printf("\n");
    }

    return 0;
}
