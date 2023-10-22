#ifndef BUFFER_POOL_H
#define BUFFER_POOL_H

#pragma once
#include <stdlib.h>
#include <stdbool.h>

// 16 KB = 2^14 B
#define MAX_BUFFER_SIZE (1 << 14) // Maximum size of a buffer in bytes
#define MAX_NUM_BUFFERS (1 << 3) // Maximum number of buffers in the pool

typedef struct {
    void* data; // Pointer to the buffer data
    size_t size; // Size of the buffer in bytes
    bool in_use; // Whether the buffer is currently in use
} buffer_t;

typedef struct {
    buffer_t buffers[MAX_NUM_BUFFERS]; // Array of buffers in the pool
    size_t num_buffers; // Number of buffers in the pool
} buffer_pool_t;

// Initialize a buffer pool with buffers of the given size
void buffer_pool_init(buffer_pool_t* pool, size_t buffer_size);

// Allocate a buffer from the pool
void* buffer_pool_alloc(buffer_pool_t* pool, size_t size);

// Free a buffer and return it to the pool
void buffer_pool_free(buffer_pool_t* pool, void* data);

// get_buffer from buffer pool
void* get_buffer(size_t size);

// release_buffer to buffer pool
void release_buffer(void* data);

#endif  
// BUFFER_POOL_H
