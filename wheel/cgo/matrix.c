#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "matrix.h"

buffer_pool_t* pool;


// get_buffer from buffer pool  
void* get_buffer(size_t size) {
    return buffer_pool_alloc(pool, size);
}

// release_buffer to buffer pool
void release_buffer(void* data) {
    buffer_pool_free(pool, data);
}

matrix* create_matrix(int rows, int cols) {
    if (pool == NULL) {
        buffer_pool_init(pool, MAX_BUFFER_SIZE);
    }
    matrix* m = (matrix*)get_buffer(sizeof(matrix));
    m->rows = rows;
    m->cols = cols;
    m->data = (double **)get_buffer(rows * sizeof(double *));
    for (int i = 0; i < rows; i++) {
        m->data[i] = (double *)get_buffer(cols * sizeof(double));
    }
    return m;
}

void destroy_matrix(matrix* m) {
    for (int i = 0; i < m->rows; i++) {
        release_buffer(m->data[i]);
    }
    release_buffer(m->data);
    release_buffer(m);
}

matrix* random_matrix(int rows, int cols) {
    matrix* m = create_matrix(rows, cols);
    srand(time(NULL));
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            m->data[i][j] = (double)rand() / RAND_MAX;
        }
    }
    return m;
}

void print_matrix(matrix* m) {
    printf("Matrix %dx%d:\n", m->rows, m->cols);
    for (int i = 0; i < m->rows; i++) {
        printf("[");
        for (int j = 0; j < m->cols; j++) {
            printf("%f ", m->data[i][j]);
        }
        printf("]\n");
    }
}

