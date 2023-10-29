#ifndef MATRIX_H
#define MATRIX_H 
   
#pragma once

typedef struct real_matrix {
    int rows,cols;
    double *data;
} real_matrix;

real_matrix* create_real_matrix(int , int );
void destroy_real_matrix(real_matrix *);
real_matrix* create_rand_matrix(int , int , double , double );

real_matrix* add(real_matrix *, real_matrix *);
real_matrix* sub(real_matrix *, real_matrix *);
real_matrix* mul(real_matrix *, real_matrix *);
real_matrix* mulV2(real_matrix *, real_matrix *);
real_matrix* transpose(real_matrix *); 
real_matrix* dot(real_matrix *, real_matrix *);
real_matrix* hadamard(real_matrix *, real_matrix *);
real_matrix* scalar(real_matrix *, double );
int equal_real_matrix(real_matrix *, real_matrix *);
real_matrix* apply(real_matrix *, double (*)(double));
real_matrix* apply2(real_matrix *, double (*)(double,double), double);
real_matrix* apply3(real_matrix *, double (*)(double,double,double), double, double);
real_matrix* apply4(real_matrix *, double (*)(double,double,double,double), double, double, double);
real_matrix* apply5(real_matrix *, double (*)(double,double,double,double,double), double, double, double, double);


void print_real_matrix(real_matrix *);

#endif 
// MATRIX_H
