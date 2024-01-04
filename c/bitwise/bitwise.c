#include <stdio.h>

void print_binary(unsigned int num) {
    for(int i = 31; i >= 0; i--) {
        if(num & (1 << i))
            printf("1");
        else
            printf("0");
    }
    printf("\n");
}

int main() {
    unsigned int x = 1764;  
    unsigned int y = 42;    
    int z = 0;           

    printf("Binary representation of x:            ");
    print_binary(x);

    printf("Binary representation of y:            ");
    print_binary(y);

    z = x & y;             
    printf("Bitwise AND - result is is %d          ", z);
    print_binary(z);

    z = x | y;             
    printf("Bitwise OR - result is %d            ", z );
    print_binary(z);

    z = x ^ y;             
    printf("Bitwise XOR - result is %d           ", z );
    print_binary(z);

    z = ~x;                
    printf("Bitwise complement - result is %d   ", z );
    print_binary(z);

    z = x << 2;            
    printf("Left shift - result is %d            ", z );
    print_binary(z);

    z = x >> 2;            
    printf("Right shift - result is %d            ", z );
    print_binary(z);

    return 0;
}
