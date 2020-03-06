#include <stdio.h>
#include <math.h>


bool judgeSquareSum(int c){
    long a = 0;
    long b = (long)sqrt(c);

    while(a <= b) {
        long sum = a*a + b*b;
        if(sum == c) {
            return true;
        }

        if(sum < c) {
            a++;
        } else {
            b--;
        }
    }

    return false;
}

int main() {
	printf("%d\n", judgeSquareSum(3));
	printf("%d\n", judgeSquareSum(5));
	printf("%d\n", judgeSquareSum(8));
	printf("%d\n", judgeSquareSum(7));
}