#include <stdio.h>

// Define enums for the days of the week and cardinal directions
enum week {Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday};
enum direction {North, East, South, West};

// Function prototypes
void printDirection(enum direction dir);

int main() {

    printf("\033[H\033[2J\033[32m");
    printf("1764Snippets\n");
    printf("Enums\n");

    // Declare and initialize a variable of type week
    enum week today = Wednesday;
    printf("Wednesday is day %d\n", today);
    printf("%d\n", today + 1764 + Sunday);
    printf("%d\n", today + 1764 + Monday);

    // Declare and initialize a variable of type direction
    enum direction chosenDirection = South + 1;

    // Check if chosenDirection is a valid enum value
    if (chosenDirection >= North && chosenDirection <= West) {
        printDirection(chosenDirection);
    } else {
        printf("Invalid direction!\n");
    }

    return 0;
}

// Function to print the chosen direction
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
