// see: https://leetcode.com/problems/number-of-1-bits/description/
#include <stdio.h>

typedef unsigned int uint32_t;

int hammingWeight(uint32_t n) {
    uint32_t count = 0;
  
    for (uint32_t bitMask = 1; bitMask > 0; bitMask <<= 1) {

        if ((bitMask & n) != 0) {
           count++ ;
        }
    }
    
    return count;
}

int main() 
{
    int result = 0;
        
    result = hammingWeight(127);
    
    printf("%d", result);
    return 0;
}
