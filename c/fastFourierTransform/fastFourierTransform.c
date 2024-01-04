#include <complex.h>
#include <math.h>
#include <stdio.h>

#ifndef M_PI
#define M_PI 3.141592653589793238
#endif

void fastFourierTransform(complex double* complexNumbers, int arraySize) {
    if (arraySize == 1)
        return;

    complex double evenIndexedComplexNumbers[arraySize/2];
    complex double oddIndexedComplexNumbers[arraySize/2];
    for (int index = 0; index < arraySize/2; index++) {
        evenIndexedComplexNumbers[index] = complexNumbers[index*2];
        oddIndexedComplexNumbers[index] = complexNumbers[index*2 + 1];
    }

    fastFourierTransform(evenIndexedComplexNumbers, arraySize/2);
    fastFourierTransform(oddIndexedComplexNumbers, arraySize/2);

    for (int index = 0; index < arraySize/2; index++) {
        complex double t = cexp(-2.0 * I * M_PI * index / arraySize) * oddIndexedComplexNumbers[index];
        complexNumbers[index] = evenIndexedComplexNumbers[index] + t;
        complexNumbers[index + arraySize/2] = evenIndexedComplexNumbers[index] - t;
    }
}

void inverseFastFourierTransform(complex double* complexNumbers, int arraySize) {
    for (int i = 0; i < arraySize; i++) {
        complexNumbers[i] = conj(complexNumbers[i]);
    }

    fastFourierTransform(complexNumbers, arraySize);

    for (int i = 0; i < arraySize; i++) {
        complexNumbers[i] = conj(complexNumbers[i]) / arraySize;
    }
}

void applyLowPassFilter(complex double* complexNumbers, int arraySize, double cutoffFrequency, double samplingRate) {
    double frequencyStep = samplingRate / arraySize;

    for (int index = 0; index < arraySize; index++) {
        double frequency = index * frequencyStep;

        if (frequency > cutoffFrequency) {
            complexNumbers[index] = 0;
        }
    }
}

int main() {
    complex double complexNumbersArray[] = {0, 1, 3, 4, 4, 3, 2, 1};
    int arraySize = sizeof(complexNumbersArray) / sizeof(complexNumbersArray[0]);
    double samplingRate = 1000.0;
    double cutoffFrequency = 200.0;

    printf("Original array:\n");
    for (int index = 0; index < arraySize; index++) {
        printf("(%f, %f)\n", creal(complexNumbersArray[index]), cimag(complexNumbersArray[index]));
    }

    fastFourierTransform(complexNumbersArray, arraySize);
    applyLowPassFilter(complexNumbersArray, arraySize, cutoffFrequency, samplingRate);
    inverseFastFourierTransform(complexNumbersArray, arraySize);

    printf("\nAfter FFT, filtering, and IFFT:\n");
    for (int index = 0; index < arraySize; index++) {
        printf("(%f, %f)\n", creal(complexNumbersArray[index]), cimag(complexNumbersArray[index]));
    }

    return 0;
}
