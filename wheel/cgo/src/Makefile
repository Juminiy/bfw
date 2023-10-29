CC = clang
CFLAGS = -Wall -Wextra -pedantic -std=c99

home_dir = ..
binary_dir = $(home_dir)/bin
include_dir = $(home_dir)/include
library_dir = $(home_dir)/lib

binary_out = *.o *.a *.so *.dylib *.exe *.exe~
header_file = *.h
matrix_lib = libmatrix.a
common_lib = libcommon.a
matrix_dylib = matrix.so
common_dylib = common.so

all: matrix.lib dup_header

test: matrix_test

dup_header: $(header_file)
	cp $^ $(include_dir)

buffer_pool.o: buffer_pool.c buffer_pool.h
	$(CC) $(CFLAGS) -o $@ -c $<

matrix.o: matrix.c matrix.h
	$(CC) $(CFLAGS) -o $@ -c $< 

matrix_dy.o: matrix.c matrix.h
	$(CC) $(CFLAGS) -o $@ -c -fPIC $<

common.o: common.c common.h
	$(CC) $(CFLAGS) -o $@ -c $<

common_dy.o: common.c common.h
	$(CC) $(CFLAGS) -o $@ -c -fPIC $<

matrix.lib: matrix.o
	ar rcs $(matrix_lib) $^
	cp $(matrix_lib) $(library_dir)/$(matrix_lib)
	rm -f $^ $(matrix_lib)

common.lib: common.o
	ar rcs $(common_lib) $^
	cp $(common_lib) $(library_dir)/$(common_lib)
	rm -f $^ $(common_lib)

matrix.dylib: matrix_dy.o
	$(CC) -shared -o $(matrix_dylib) $<
	cp $(matrix_dylib) $(library_dir)/$(matrix_dylib)

common.dylib: common_dy.o
	$(CC) -shared -o $(common_dylib) $<
	cp $(common_dylib) $(library_dir)/$(common_dylib)

stlib: common.lib matrix.lib
	make dup_header
	make clean_src

dylib: common.dylib matrix.dylib
	make dup_header
	make clean_src

matrix_test: matrix_test.c matrix.h buffer_pool.h matrix.lib
	$(CC) $(CFLAGS) -o $@ $< -L$(library_dir) -lmatrix -lpthread
	mv $@ $(binary_dir)/$@

run_test: matrix_test
	cp $(library_dir)/$(matrix_lib) $(binary_dir)/$(matrix_lib)
	cd $(binary_dir) && chmod 777 $< && ./$<

clean_src:
	rm -f $(binary_out)

clean_bin:
	cd $(binary_dir) && rm -f *

clean_lib:
	cd $(library_dir) && rm -f $(binary_out)

clean_include:
	cd $(include_dir) && rm -f $(header_file)


clean:
	make clean_src
	make clean_bin
	make clean_lib
	make clean_include
	clear