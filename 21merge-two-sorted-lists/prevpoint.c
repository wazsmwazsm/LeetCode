	#include <stdio.h>

	struct ListNode {
		int val;
		struct ListNode *next;
	};
	
	struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2);

	void printList(struct ListNode* l) {
		while (l != NULL) {
			printf(" %d ", l->val);
			l = l->next;
		}
	}

	int main() {

		struct ListNode l1 = {1};
		struct ListNode l12 = {2};
		struct ListNode l13 = {4};
		l12.next = &l13;
		l1.next = &l12;

		struct ListNode l2 = {1};
		struct ListNode l22 = {3};
		struct ListNode l23 = {4};
		l22.next = &l23;
		l2.next = &l22;

		printf("\n l1: ");
		printList(&l1);
		printf("\n l2: ");
		printList(&l2);

		struct ListNode *l = mergeTwoLists(&l1, &l2);

		printf("\n merged: ");
		printList(l);
		return 0;
	}

	struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){
		struct ListNode node = {};
		struct ListNode * head = &node;
		struct ListNode * prev = head;

		while (l1 != NULL && l2 != NULL) {
			if (l1->val <= l2->val) {
				prev->next = l1;
				l1 = l1->next;
			} else {
				prev->next = l2;
				l2 = l2->next;
			}
			prev = prev->next;
		}

		prev->next = l1 == NULL ? l2 : l1;

		return head->next;
	}