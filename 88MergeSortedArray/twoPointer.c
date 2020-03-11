#include <stdio.h>
#include <string.h>

void merge(int* nums1, int nums1Size, int m, int* nums2, int nums2Size, int n){
    int p1 = m - 1;
    int p2 = n - 1;
    int p = m + n -1;

    while(p1 >= 0 && p2 >= 0) {
        if (nums1[p1] < nums2[p2]) {
            nums1[p] = nums2[p2];
            p2--;
        } else {
            nums1[p] = nums1[p1];
            p1--;
        }
        p--;
    }

    memcpy(nums1, nums2, sizeof(int)*(p2+1));

}


int main() {
	int nums1[6] = {1,2,3,0,0,0};
	int nums2[3] = {2,3,5};
	int nums1Size = sizeof(nums1)/sizeof(int);
	int nums2Size = sizeof(nums2)/sizeof(int);
	merge(nums1, nums1Size, 3, nums2, nums2Size, 3);

	for (int i = 0; i < nums1Size; i++)
	{
		printf("%d ", nums1[i]);
	}
	
}