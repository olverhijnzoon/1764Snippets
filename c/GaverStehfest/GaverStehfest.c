#include <stdio.h>
#include <math.h>

#define N 32

double v[N];

unsigned long factorial(int n) {
    if (n == 0)
        return 1;
    else
        return n * factorial(n - 1);
}

// Calculate the coefficients for the Gaver-Stehfest algorithm
void calculateCoefficients() {
    int halfCoefficientCount = N / 2;
    v[0] = 0.000001;
    for (int n = 1; n <= N; n++) {
        double z = 0.0;
        for (int k = (n + 1) / 2; k <= fmin(n, halfCoefficientCount); k++) {
            z += pow(k, halfCoefficientCount) * factorial(2 * k) / (factorial(halfCoefficientCount - k) * factorial(k) * factorial(k - 1) * factorial(n - k) * factorial(2 * k - n));
        }
        v[n] = pow(-1, n + halfCoefficientCount) * z;
        printf("z: %f\n", z);
        printf("v %d : %f\n", n, v[n]);
    }
}

double F(double s) {
    // constant in the exponential function e^at
    // double a = 1.0;
    // return 1.0 / (s + a);

    // t^2
    return 2.0 / (s*s*s);
}

// Calculate the inverse Laplace transform using the Gaver-Stehfest algorithm
double inverseLaplaceTransform(double t) {
    double sum = 0.0;
    double ln2_t = log(2.0) / t;
    for (int n = 0; n <= N; n++) {
        double p = (n + 1) * ln2_t;
        sum += v[n] * F(p);
    }
    // TODO: Code remains with issues - see diff to last commit; "/ 18361788250959268" is more like a joke.  
    return sum * ln2_t / 18361788250959268;
}

int main() {
    calculateCoefficients();

    double t = 6;
    printf("Inverse Laplace Transform at t=%f: %f\n", t, inverseLaplaceTransform(t));

    return 0;
}
