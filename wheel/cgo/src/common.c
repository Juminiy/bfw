#include <stdio.h>
#include "common.h"


void scanf_int(int *x){
    scanf("%d", x);
}
void scanf_double(double *x){
    scanf("%lf", x);
}
void scanf_string(char *x){
    scanf("%s", x);
}
void scanf_char(char *x){
    scanf("%c", x);
}
void scanf_long(long *x){
    scanf("%ld", x);
}
void scanf_float(float *x){
    scanf("%f", x);
}
void scanf_unsigned_int(unsigned int *x){
    scanf("%u", x);
}
void scanf_unsigned_long(unsigned long *x){
    scanf("%lu", x);
}
void scanf_unsigned_long_long(unsigned long long *x){
    scanf("%llu", x);
}
void scanf_unsigned_char(unsigned char *x){
    scanf("%hhu", x);
}
void scanf_unsigned_short(unsigned short *x){
    scanf("%hu", x);
}
void scanf_long_double(long double *x){
    scanf("%Lf", x);
}
void scanf_long_long(long long *x){
    scanf("%lld", x);
}


void printf_int(int x){
    printf("%d", x);
}
void printf_double(double x){
    printf("%lf", x);
}
void printf_string(char *x){
    printf("%s", x);
}
void printf_char(char x){
    printf("%c", x);
}
void printf_long(long x){
    printf("%ld", x);
}
void printf_float(float x){
    printf("%f", x);
}
void printf_unsigned_int(unsigned int x){
    printf("%u", x);
}
void printf_unsigned_long(unsigned long x){
    printf("%lu", x);
}
void printf_unsigned_long_long(unsigned long long x){
    printf("%llu", x);
}
void printf_unsigned_char(unsigned char x){
    printf("%hhu", x);
}
void printf_unsigned_short(unsigned short x){
    printf("%hu", x);
}
void printf_long_double(long double x){
    printf("%Lf", x);
}
void printf_long_long(long long x){
    printf("%lld", x);
}
