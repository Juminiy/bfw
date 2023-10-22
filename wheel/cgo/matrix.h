#ifndef MATRIX_H
#define MATRIX_H 
   
#pragma once
#include "buffer_pool.h"

typedef struct {
    int rows;
    int cols;
    double **data;
} matrix;

matrix* create_matrix(int rows, int cols);
matrix* random_matrix(int rows, int cols);
void destroy_matrix(matrix *m);
void print_matrix(matrix *m);

#endif 
// MATRIX_H
