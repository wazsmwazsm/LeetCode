#include <stdio.h>
#include <malloc.h>


int* twoSum(int* numbers, int numbersSize, int target, int* returnSize){


    if(numbersSize <= 0) {
        return NULL;
    }

    int *res = (int*)malloc(2*sizeof(int));
    *returnSize = 2;

    int min = 0;
    int max = numbersSize - 1;
    while(min < max) {
        int sum = numbers[min] + numbers[max];
        if (sum == target){
            res[0] = min+1;
            res[1] = max+1;
            
            return res;
        }

        if (sum < target) {
            min++;
        } else {
            max--;
        }
    }

    return NULL;
}

int main() {

    int a[4] = {2,7,8,11};
    int target = 9;
    int * res;
    int resSize;

    res = twoSum(a, sizeof(a)/sizeof(int), target, &resSize);
    for(int i=0; i<resSize; i++) {
        printf("%d ", res[i]);
    }
    
    free(res);
}