CC = clang
CFLAGS = -Wall -Wextra -pedantic -std=c99
all: matrix_test lib

buffer_pool.o: buffer_pool.c buffer_pool.h
	$(CC) $(CFLAGS) -o $@ -c $<

matrix.o: matrix.c matrix.h
	$(CC) $(CFLAGS) -o $@ -c $<

matrix_test: matrix_test.c matrix.h buffer_pool.h lib
	$(CC) $(CFLAGS) -o $@ $< -L. -lmatrix -lpthread 

lib: matrix.o buffer_pool.o
	ar rcs libmatrix.a matrix.o buffer_pool.o

run_test: matrix_test
	chmod 777 $< && ./matrix_test
