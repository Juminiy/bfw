#include <stdio.h>
#include "matrix.h"

int main(void) {
    
    int size = 100;
    matrix* m = random_matrix(size, size);

    printf("Random matrix:\n");
    print_matrix(m);
    
    // Free memory
    destroy_matrix(m);
    
    return 0;
}
