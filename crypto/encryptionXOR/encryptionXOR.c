#include <stdio.h>
#include <string.h>

#define BLOCK_SIZE 8

void xorEncrypt(char *text, char *key) {
    int textLen = strlen(text);
    int keyLen = strlen(key);

    for (int i = 0; i < textLen; i++) {
        text[i] = text[i] ^ key[i % keyLen];
    }
}

void xorOperation(char *block, const char *key) {
    for (int i = 0; i < BLOCK_SIZE; ++i) {
        block[i] = block[i] ^ key[i];
    }
}

void xorChainBlockEncrypt(char *text, const char *key, int textLen) {
    char prevBlock[BLOCK_SIZE];
    memset(prevBlock, 0, BLOCK_SIZE); // Initialization Vector (IV)

    for (int i = 0; i < textLen; i += BLOCK_SIZE) {
        for (int j = 0; j < BLOCK_SIZE; ++j) {
            if (i + j < textLen) {
                text[i + j] = text[i + j] ^ prevBlock[j];
            }
        }
        xorOperation(&text[i], key);
        memcpy(prevBlock, &text[i], BLOCK_SIZE); // Copy current encrypted block
    }
}

void xorChainBlockDecrypt(char *text, const char *key, int textLen) {
    char prevBlock[BLOCK_SIZE], tempBlock[BLOCK_SIZE];
    memset(prevBlock, 0, BLOCK_SIZE); // IV

    for (int i = 0; i < textLen; i += BLOCK_SIZE) {
        memcpy(tempBlock, &text[i], BLOCK_SIZE); // Copy current encrypted block
        xorOperation(&text[i], key);

        for (int j = 0; j < BLOCK_SIZE; ++j) {
            if (i + j < textLen) {
                text[i + j] = text[i + j] ^ prevBlock[j];
            }
        }
        memcpy(prevBlock, tempBlock, BLOCK_SIZE); // Copy original encrypted block
    }
}

int main() {

    printf("1764Snippets\n");
    printf("Crypto XOR Encryption Chain\n");

    /*
    
        This snippet demonstrates a XOR encryption which is quite basic as it can be easily reversed by XORing with the same key which makes it symmetric. Even the chained version of it in the second part of the snippet should not be used in production but it is at least an implementation using a simplified Cipher Block Chaining (CBC) operation. 
    */

    char textSimple[] = "1764Snippets";
    char keySimple[] = "key23";

    printf("Original text: %s\n", textSimple);
    xorEncrypt(textSimple, keySimple);
    printf("Encrypted text: ");
    for (int i = 0; i < strlen(textSimple); i++) {
        printf("%02x", (unsigned char)textSimple[i]);
    }
    printf("\n");

    // Decrypting the text (since XOR is reversible)
    xorEncrypt(textSimple, keySimple);
    printf("Decrypted text: %s\n", textSimple);


    char text[] = "1764Snippets, the successor to the 42Snippets series/repository, boldly multiplies the ordinary.";
    const char key[BLOCK_SIZE] = "key12345"; // Ensure this is exactly BLOCK_SIZE characters
    int textLen = strlen(text);

    printf("Original text: %s\n", text);
    xorChainBlockEncrypt(text, key, textLen);
    printf("Encrypted text: ");
    for (int i = 0; i < textLen; ++i) {
        printf("%02x", (unsigned char)text[i]);
    }
    printf("\n");

    xorChainBlockDecrypt(text, key, textLen);
    printf("Decrypted text: %s\n", text);

    return 0;
}
