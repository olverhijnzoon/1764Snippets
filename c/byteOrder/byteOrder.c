#include <stdio.h>
#include <arpa/inet.h>

int main() {

    printf("# 1764Snippets\n");
    printf("## C byteOrder\n");

    /*
    This shows the usage of the htons function in C to show a conversion from the host byte order (typically Little Endian) to network byte order which is usually Big Endian (configurable header)

    lldb) breakpoint set --file byteOrder.c --line 42
    Breakpoint 1: where = byteOrderX`main + 108 at byteOrder.c:42:35, address = 0x0000000100003efc
    (lldb) r
    ...
    hostshort = 0x06e4
    ...
    -> 42       printf("netshort = 0x%04x\n", netshort);
    ...
    (lldb) frame variable
    (uint16_t) hostshort = 1764
    (uint16_t) netshort = 58374
    (lldb) memory read &hostshort
    0x16fdfedaa: e4 06 00 00 00 00 e0 ef df 6f 01 00 00 00 e0 10  .........o......
    0x16fdfedba: 25 81 01 00 00 00 00 00 00 00 00 00 00 00 00 00  %...............
    (lldb) memory read &netshort
    0x16fdfeda8: 06 e4 e4 06 00 00 00 00 e0 ef df 6f 01 00 00 00  ...........o....
    0x16fdfedb8: e0 10 25 81 01 00 00 00 00 00 00 00 00 00 00 00  ..%.............

    */

    // Declare a 16-bit number in host byte order
    uint16_t hostshort = 1764;

    // Print the value in hexadecimal format
    printf("hostshort = 0x%04x\n", hostshort);

    // Convert the number to network byte order using htons
    uint16_t netshort = htons(hostshort);

    // Print the value in hexadecimal format
    printf("netshort = 0x%04x\n", netshort);
    // Output:
    // hostshort = 0x1389
    // netshort = 0x8913

    return 0;

    }
