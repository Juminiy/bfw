#include "buffer_pool.h"

// Initialize a buffer pool with buffers of the given size
void buffer_pool_init(buffer_pool_t* pool, size_t buffer_size) {
    pool->num_buffers = 0;
    for (size_t i = 0; i < MAX_NUM_BUFFERS; i++) {
        pool->buffers[i].data = malloc(buffer_size);
        pool->buffers[i].size = buffer_size;
        pool->buffers[i].in_use = false;
        pool->num_buffers++;
    }
}

// Allocate a buffer from the pool
void* buffer_pool_alloc(buffer_pool_t* pool, size_t size) {
    for (size_t i = 0; i < pool->num_buffers; i++) {
        if (!pool->buffers[i].in_use && pool->buffers[i].size >= size) {
            pool->buffers[i].in_use = true;
            return pool->buffers[i].data;
        }
    }
    return NULL; // No available buffer found
}

// Free a buffer and return it to the pool
void buffer_pool_free(buffer_pool_t* pool, void* data) {
    for (size_t i = 0; i < pool->num_buffers; i++) {
        if (pool->buffers[i].data == data) {
            pool->buffers[i].in_use = false;
            return;
        }
    }
}

