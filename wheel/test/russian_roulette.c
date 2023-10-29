#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(){
    // [0,6)+1 = [1,6]
    srand(getpid());
    int x = rand() % 6 + 1;
    printf("you got %d\n", x);
    if (x & 1) {
        printf("You are dead!\n");
        int status = rmdir("fake_os");
        printf("rmdir status = %s\n", ((status == 0) ? "Ok":"Error"));
    } else {
        printf("You are alive!\n");
    }
    return 0;
}

