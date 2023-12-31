#include <stdio.h>
#include <math.h>

double absoluteValue(double x) {
    return fabs(x);
}

int absoluteValueBitwise(int n) {
    // shifts the bits of n to the right
    // negative n results in a mask of all ones; positive, the mask will be all zeros
    int const mask = n >> (sizeof(int) * 8 - 1);
    // printf("Mask %d", mask);
    // bitwise XOR operation; will flip all the bits if negative
    return ((n + mask) ^ mask);
}

double sumAbsoluteValues(double tensor[4][4]) {
    double sum = 0.0;
    for(int i = 0; i < 4; i++) {
        for(int j = 0; j < 4; j++) {
            sum += absoluteValue(tensor[i][j]);
        }
    }
    return sum;
}

int main() {
    
    printf("1764Snippets\n");
    printf("Absolute Values\n\n");

    int num = -42;
    printf("The absolute value of %d is %d.\n\n", num, absoluteValueBitwise(num));

    double tensor[4][4] = {
        {1.0, -2.0, 13.0, 42.0},
        {-2.0, 25.0, 6.0, 700.0},
        {13.0, 6.0, 8.0, -97.0},
        {42.0, 700.0, -97.0, 10.0}
    };

    for(int i = 0; i < 4; i++) {
        for(int j = 0; j < 4; j++) {
            printf("%.2f ", absoluteValue(tensor[i][j]));
        }
        printf("\n");
    }

    double sum = sumAbsoluteValues(tensor);
    printf("\nSum of Absolute Values of Elements: %.2f\n", sum);

    return 0;
}
