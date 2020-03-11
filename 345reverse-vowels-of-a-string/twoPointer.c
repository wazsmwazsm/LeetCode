#include <stdio.h>
#include <string.h>
bool isVowels(char c);
char * reverseVowels(char * s);


int main() {
    char s[] ="leetcode";
    printf("%s\n", reverseVowels(s));
    char s2[] = "hello";
    printf("%s\n", reverseVowels(s2));
}


char * reverseVowels(char * s){
    int start = 0;
    int end = strlen(s) - 1;

    while(start < end) {
        while(!isVowels(s[start]) && start < end) start++;
        while(!isVowels(s[end]) && start < end) end--;
        char tmp = s[start];
        s[start] = s[end];
        s[end] = tmp;

        start++;
        end--;
    }

    return s;
}

bool isVowels(char c) {
    if (c == 'a' || c =='e'|| c == 'i'|| c == 'o'|| c == 'u'|| 
        c == 'A'|| c == 'E'|| c == 'I'|| c == 'O'|| c == 'U' 
    ) {
        return true;
    }
    return false;
}