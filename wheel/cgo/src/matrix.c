#include "matrix.h"
#include <stdio.h>
#include <stdlib.h>

real_matrix* create_real_matrix(int rows, int cols){
    real_matrix* matrix = (real_matrix*)malloc(sizeof(real_matrix));
    matrix->rows = rows;
    matrix->cols = cols;
    matrix->data = (double*)malloc(rows*cols*sizeof(double));
    return matrix;
}

void destroy_real_matrix(real_matrix* matrix){
    matrix->rows = 0;
    matrix->cols = 0;
    free(matrix->data);
    free(matrix);
}

real_matrix* create_rand_matrix(int rows, int cols, double range_start, double range_end){
    real_matrix* matrix = create_real_matrix(rows,cols);
    for(int i=0;i<rows;i++){
        for(int j=0;j<cols;j++){
            matrix->data[i*cols+j] = rand() % (int)(range_end - range_start + 1) + range_start;
        }
    }
    return matrix;
}

int is_same_shape(real_matrix* matrix1, real_matrix* matrix2){
    return matrix1->rows == matrix2->rows && matrix1->cols == matrix2->cols;
}

int is_phalanx(real_matrix* matrix){
    return matrix->rows == matrix->cols;
}

int can_multiply(real_matrix* matrix1, real_matrix* matrix2){
    return matrix1->cols == matrix2->rows;
}

void elem_swap(real_matrix* matrix, int i1, int j1, int i2, int j2){
    double temp = matrix->data[i1*matrix->cols+j1];
    matrix->data[i1*matrix->cols+j1] = matrix->data[i2*matrix->cols+j2];
    matrix->data[i2*matrix->cols+j2] = temp;
}

void real_matrix_swap(real_matrix* matrix1, real_matrix* matrix2){
    int temp = matrix1->rows;
    matrix1->rows = matrix2->rows;
    matrix2->rows = temp;
    temp = matrix1->cols;
    matrix1->cols = matrix2->cols;
    matrix2->cols = temp;
    double* temp_data = matrix1->data;
    matrix1->data = matrix2->data;
    matrix2->data = temp_data;
}

void real_matrix_copy(real_matrix *dest, real_matrix *src){
    dest->rows = src->rows;
    dest->cols = src->cols;
    dest->data = (double*)malloc(dest->rows*dest->cols*sizeof(double));
    for (int i=0;i<dest->rows;i++){
        for (int j=0;j<dest->cols;j++){
            dest->data[i*dest->cols+j] = src->data[i*src->cols+j];
        }
    }
}

real_matrix* add(real_matrix *A, real_matrix *B){
    if (!A || !B || !is_same_shape(A, B)){
        return NULL;
    }
    real_matrix* matrix = create_real_matrix(A->rows,A->cols);
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<A->cols;j++){
            int index = i*A->cols+j;
            matrix->data[index] = A->data[index] + B->data[index];
        }
    }
    return matrix;
}
real_matrix* sub(real_matrix *A, real_matrix *B){
    if (!A || !B || !is_same_shape(A, B)){
        return NULL;
    }
    real_matrix* matrix = create_real_matrix(A->rows,A->cols);
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<A->cols;j++){
            int index = i*A->cols+j;
            matrix->data[index] = A->data[index] - B->data[index];
        }
    }
    return matrix;
}
real_matrix* mul(real_matrix *A, real_matrix *B){
    if (!A || !B || !can_multiply(A, B)){
        return NULL;
    }
    real_matrix* matrix = create_real_matrix(A->rows,B->cols);
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<B->cols;j++){
            double sum = 0;
            for (int k=0;k<A->cols;k++){
                sum += A->data[i*A->cols+k] * B->data[k*B->cols+j];
            }
            matrix->data[i*A->cols+j] = sum;
        }
    }
    return matrix;
}

real_matrix* mulV2(real_matrix *A, real_matrix *B){
    if (!A || !B || !can_multiply(A, B)){
        return NULL;
    }
    B = transpose(B);
    real_matrix* matrix = create_real_matrix(A->rows,B->rows);
    for (int i=0;i<A->rows;i++){
        for (int j=0; j<B->rows;j++){
            double sum = 0;
            for (int k=0;k<A->cols;k++){
                sum += A->data[i*A->cols+k] * B->data[j*B->cols+k];
            }
            matrix->data[i*A->cols+j] = sum;
        }
    }
    return matrix;
}

// test pass 
void phalanx_transpose(real_matrix* A){
    if (!A || !is_phalanx(A)){
        return ;
    }
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<i;j++){
            elem_swap(A,i,j,j,i);
        }
    }
}

real_matrix* transpose(real_matrix *A){
    if (!A){
        return NULL;
    }
    if (is_phalanx(A)){
        phalanx_transpose(A);
        return A;
    }
    real_matrix* matrix = create_real_matrix(A->cols,A->rows);
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<A->cols;j++){
            matrix->data[j*A->rows+i] = A->data[i*A->cols+j];
        }
    }
    return matrix;
} 
real_matrix* dot(real_matrix *, real_matrix *);
real_matrix* hadamard(real_matrix *, real_matrix *);
real_matrix* scalar(real_matrix *, double );

int equal_real_matrix(real_matrix *A, real_matrix *B){
    if (!A || !B || !is_same_shape(A, B)){
        return 0;
    }
    for (int i=0;i<A->rows;i++){
        for (int j=0;j<A->cols;j++){
            if (A->data[i*A->cols+j] != B->data[i*A->cols+j]){
                return 0;
            }
        }
    }
    return 1;
}


void print_real_matrix(real_matrix* matrix){
    if (!matrix){
        puts("null");
        return;
    }
    for(int i=0;i<matrix->rows;i++){
        for(int j=0;j<matrix->cols;j++){
            printf(" %.1f",matrix->data[i*matrix->cols+j]);
        }
        puts("");
    }
}
