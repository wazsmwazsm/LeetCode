#include <stdio.h>

/**

  LeetCode

  Invert a binary tree.
       4
     /   \
    2     7
   / \   / \
  1   3 6   9
  to
       4
     /   \
    7     2
   / \   / \
  9   6 3   1


*/

// 定义树结构
struct TreeNode {
   int val;
   struct TreeNode *left;
   struct TreeNode *right;
};

// 二叉树反转函数
struct TreeNode* invertTree(struct TreeNode* root)
{
    struct TreeNode* tmp;

    if(root == NULL)
    {
        return NULL;
    }

    tmp = root->left;
    root->left = root->right;
    root->right = tmp;

    invertTree(root->left);

    invertTree(root->right);

    return root;
}

int main()
{
    // 建立一个简单的树
    struct TreeNode l2 = {3, NULL, NULL};
    struct TreeNode l3 = {4, NULL, NULL};
    struct TreeNode r2 = {6, NULL, NULL};
    struct TreeNode r3 = {7, NULL, NULL};
    struct TreeNode l1 = {2, &l2, &l3};
    struct TreeNode r1 = {5, &r2, &r3};

    struct TreeNode tree = {1, &l1, &r1};

    printf("%d%d%d%d%d%d%d\n",tree.val, tree.left->val, tree.left->left->val, tree.left->right->val, tree.right->val, tree.right->left->val, tree.right->right->val);
    // 执行转换
    invertTree(&tree);

    printf("%d%d%d%d%d%d%d\n",tree.val, tree.left->val, tree.left->left->val, tree.left->right->val, tree.right->val, tree.right->left->val, tree.right->right->val);

    return 0;
}
