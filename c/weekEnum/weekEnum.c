#include <stdio.h>

enum week {Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday};

enum direction {North, East, South, West};

// cardinal switch function
void printDirection(enum direction dir) {
    switch (dir)
    {
    case North:
        printf("You're heading North!\n");
        break;
    case East:
        printf("You're heading East!\n");
        break;
    case South:
        printf("You're heading South!\n");
        break;
    case West:
        printf("You're heading West!\n");
        break;
    default:
        printf("Invalid direction!\n");
    }
}

int main() {

    printf("\033[H\033[2J\033[32m");
    printf("1764Snippets\n");
    printf("Enums\n");

    // Declare a variable of type week
    enum week today;
    today = Wednesday;
    printf("Wednesday is day %d\n",today);
    printf("%d\n",today+1764+Sunday);
    printf("%d\n",today+1764+Monday);

    enum direction chosenDirrection = South+1;
    printDirection(chosenDirrection);

    return 0;
}
