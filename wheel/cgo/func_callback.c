#include "func_callback.h"

// Example usage
void func1(void) {
    printf("Function 1\n");
}

void func2(void) {
    printf("Function 2\n");
}

int main() {
    // Initialize the struct
    func_ptr arr = {
        .func = {func1, func2},
        .size = 2
    };

    // Call the functions in the array
    for (int i = 0; i < arr.size; i++) {
        // arr.func[i]();
    }

    return 0;
}
