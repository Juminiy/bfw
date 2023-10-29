#include <stdio.h>
#include <time.h>
#include "matrix.h"

int main(void) {
    
    int size = 0;
    printf("Please input the size of matrix: ");
    scanf("%d", &size);
    time_t start_time = clock();    

    real_matrix* m1 = create_rand_matrix(size,size,1e2,1e5);
    real_matrix* m2 = create_rand_matrix(size,size,1e2,1e5);
    real_matrix* m3 = mul(m1,m2);

    // print_real_matrix(m1);
    // transpose(m1);
    // print_real_matrix(m1);

    time_t end_time = clock();
    double elapsed_time = (double)(end_time - start_time) / CLOCKS_PER_SEC;
    printf("%d * %d Matrix Multiply Elapsed time: %f seconds\n",size , size ,elapsed_time);

    // print_real_matrix(m3);

    start_time = clock();
    real_matrix* m4 = mulV2(m1,m2);
    end_time = clock();
    elapsed_time = (double)(end_time - start_time) / CLOCKS_PER_SEC;
    printf("After speed up,%d * %d Elapsed time: %f seconds\n", size, size,elapsed_time);

    // print_real_matrix(m4);
    // printf("m3 equal m4 = %d\n", equal_real_matrix(m3,m4));
    return 0;
}
