#include <stdio.h>
#include <string.h>

#define BLOCK_SIZE 8 // Define the block size for block encryption

// Function to perform simple XOR encryption
void simpleXorEncrypt(char *plaintext, char *encryptionKey) {
    int plaintextLength = strlen(plaintext);
    int keyLength = strlen(encryptionKey);

    // XOR each character of the plaintext with the encryption key
    for (int i = 0; i < plaintextLength; i++) {
        plaintext[i] = plaintext[i] ^ encryptionKey[i % keyLength];
    }
}

// Function to perform XOR operation on a block of text with a key
void xorBlockOperation(char *block, const char *key) {
    for (int i = 0; i < BLOCK_SIZE; ++i) {
        block[i] = block[i] ^ key[i];
    }
}

// Function to perform XOR Chain Block Encryption
void xorChainBlockEncrypt(char *plaintext, const char *encryptionKey, int plaintextLength) {
    char previousBlock[BLOCK_SIZE];
    memset(previousBlock, 0, BLOCK_SIZE); // Initialize with zeros as the initial vector

    for (int i = 0; i < plaintextLength; i += BLOCK_SIZE) {
        for (int j = 0; j < BLOCK_SIZE; ++j) {
            if (i + j < plaintextLength) {
                plaintext[i + j] = plaintext[i + j] ^ previousBlock[j];
            }
        }
        xorBlockOperation(&plaintext[i], encryptionKey);
        memcpy(previousBlock, &plaintext[i], BLOCK_SIZE); // Copy current encrypted block for next iteration
    }
}

// Function to perform XOR Chain Block Decryption
void xorChainBlockDecrypt(char *ciphertext, const char *decryptionKey, int ciphertextLength) {
    char previousBlock[BLOCK_SIZE], tempBlock[BLOCK_SIZE];
    memset(previousBlock, 0, BLOCK_SIZE); // Initialize with zeros as the initial vector

    for (int i = 0; i < ciphertextLength; i += BLOCK_SIZE) {
        memcpy(tempBlock, &ciphertext[i], BLOCK_SIZE); // Copy current encrypted block before decrypting
        xorBlockOperation(&ciphertext[i], decryptionKey);

        for (int j = 0; j < BLOCK_SIZE; ++j) {
            if (i + j < ciphertextLength) {
                ciphertext[i + j] = ciphertext[i + j] ^ previousBlock[j];
            }
        }
        memcpy(previousBlock, tempBlock, BLOCK_SIZE); // Copy original encrypted block for next iteration
    }
}

int main() {
    printf("1764Snippets\n");
    printf("Crypto XOR Encryption Chain\n");

    // Demonstrating simple XOR encryption and decryption
    char simplePlaintext[] = "1764Snippets";
    char simpleKey[] = "key23";

    printf("Original text: %s\n", simplePlaintext);
    simpleXorEncrypt(simplePlaintext, simpleKey);
    printf("Encrypted text: ");
    for (int i = 0; i < strlen(simplePlaintext); i++) {
        printf("%02x", (unsigned char)simplePlaintext[i]);
    }
    printf("\n");

    // Decrypting the text (since XOR is reversible)
    simpleXorEncrypt(simplePlaintext, simpleKey);
    printf("Decrypted text: %s\n", simplePlaintext);

    // Demonstrating XOR Chain Block Encryption and Decryption
    char complexPlaintext[] = "1764Snippets, the successor to the 42Snippets series/repository, boldly multiplies the ordinary.";
    const char complexKey[BLOCK_SIZE] = "key12345"; // Ensure this key is exactly BLOCK_SIZE characters
    int complexTextLength = strlen(complexPlaintext);

    printf("Original text: %s\n", complexPlaintext);
    xorChainBlockEncrypt(complexPlaintext, complexKey, complexTextLength);
    printf("Encrypted text: ");
    for (int i = 0; i < complexTextLength; ++i) {
        printf("%02x", (unsigned char)complexPlaintext[i]);
    }
    printf("\n");

    xorChainBlockDecrypt(complexPlaintext, complexKey, complexTextLength);
    printf("Decrypted text: %s\n", complexPlaintext);

    return 0;
}
