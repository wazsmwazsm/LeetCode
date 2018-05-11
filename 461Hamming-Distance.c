#include <stdio.h>


int main() {

    int x = 3; 
    int y = 1;
    int h = 0;

    h = hammingDistance(x, y);
    printf("%d\n", h);

    return 0;
}

int hammingDistance(int x, int y) {
    int count = 0;
    int xor = x ^ y;
    for (int bitMask = 1; bitMask > 0; bitMask <<= 1) {
		// 使用 bitMask 移位来比较每一位是 1 还是 0
        if ((bitMask & xor) != 0) {
            count++; 
        }
    }
    
    return count;
}

