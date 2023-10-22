#ifndef FUNC_CALLBACK_H
#define FUNC_CALLBACK_H

#include <stddef.h> // size_t

// Define the struct
typedef struct func_ptr{
    void **func; // Array of void* function pointers
    size_t size; // Size of the array
} func_ptr;

#endif 
// FUNC_CALLBACK_H
