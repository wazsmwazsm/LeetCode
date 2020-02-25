// C99
#include <stdio.h>
#include <string.h>
#include <malloc.h>

char ** generatePossibleNextMoves(char * s, int* returnSize);

char *tests[5] = {
    (char *)"+++--++-",
	(char *)"++++-++-+--+-++-",
	(char *)"++++",
	(char *)"--",
	(char *)"-+-++"};

int main() {

    char **res;
    int *returnSize;
 
    for (int i = 0; i < (sizeof(tests)/sizeof(char *) - 1); i++) {
        res = generatePossibleNextMoves(tests[i], returnSize);
        
        printf("input:\n %s\n", tests[i]);
        printf("output:\n");
        for (int j = 0; j < *returnSize; j++) {
            printf("[ %s ]\n", res[j]);
        }
        printf("\n");
        
    }
    

    return 0;
}

char ** generatePossibleNextMoves(char * s, int* returnSize){
    if (s == NULL) {
        return NULL;
    }

    if (0 == strcmp(s,"")) {
        *returnSize = 0;
        return NULL;
    }

    int cnt = 0;
    int len = strlen(s);

    char ** result = (char **)malloc((len - 1) * sizeof(char *));
    
    for (int i = 0; i < len - 1; i++) {
        if (s[i] == '+' && s[i+1] == '+') {
            result[cnt] = (char *)malloc(sizeof(char *) * (len+1));
            strcpy(result[cnt], s);

            result[cnt][i] = '-';
            result[cnt][i+1] = '-';
            cnt++;
        }
    }

    *returnSize = cnt;

    return result;
}


