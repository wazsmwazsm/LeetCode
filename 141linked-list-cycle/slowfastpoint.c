#include <stdio.h>
#include <stdbool.h>

struct ListNode {
    int val;
    struct ListNode *next;
};


bool hasCycle(struct ListNode *head);


int main() {

	struct ListNode l4 = {4};
	struct ListNode l3 = {0, &l4};
	struct ListNode l2 = {2, &l3};
	struct ListNode l1 = {3, &l2};

	l4.next = &l2;

	printf("%d\n", hasCycle(&l1));

	return 0;
}

bool hasCycle(struct ListNode *head) {
	if (head == NULL) {
		return false;
	}
	struct ListNode *slow = head;
	struct ListNode *fast = head;

	while(slow != NULL && fast != NULL && fast->next != NULL) {
		slow = slow->next;
		fast = fast->next->next;

		if (slow == fast) {
			return true;
		}
	}

	return false;
}