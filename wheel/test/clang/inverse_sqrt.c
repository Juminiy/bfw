#include <stdio.h>

float Q_rsqrt(float x ) {
    float xhalf = 0.5f * x;
    int i = *(int*)&x; // get bits for floating VALUE
    i = 0x5f3759df - (i>>1); // gives initial guess y0
    x = *(float*)&i; // convert bits BACK to float
    x = x*(1.5f-xhalf*x*x); // Newton step, repeating increases accuracy
    return x;
}


int main(int argc, char ** args) {
    float x = 0.15625;
    long i ;
    i = *(long*)&x;
    printf("x = %f, i = %x\n", x, i);
    return 0;
}