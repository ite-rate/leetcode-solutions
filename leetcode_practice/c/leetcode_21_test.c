/*
 * LeetCode #21: 题目21
 * 难度: 未知
 * 
 * 题目描述:
 * 由大模型直接生成
 * 
 * 代码骨架完整度: 80%
 */

#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

/**
 * LeetCode #21 合并两个有序链表
 * 难度：简单
 * 
 * 题目描述：
 * 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 * 
 * 示例：
 * 输入：l1 = [1,2,4], l2 = [1,3,4]
 * 输出：[1,1,2,3,4,4]
 */

// 链表节点定义
struct ListNode {
    int val;
    struct ListNode *next;
};

/**
 * 解法1：迭代法
 * 算法思路：
 * 1. 创建哑节点(dummy)作为新链表的起始点
 * 2. 使用指针遍历两个链表，比较当前节点值
 * 3. 将较小值的节点连接到新链表
 * 4. 当一个链表遍历完后，将另一个链表的剩余部分直接连接
 * 
 * 时间复杂度：O(m+n)，其中m和n分别是两个链表的长度
 * 空间复杂度：O(1)
 */
struct ListNode* mergeTwoLists_iterative(struct ListNode* l1, struct ListNode* l2) {
    // 创建哑节点简化边界处理
    struct ListNode dummy;
    struct ListNode* tail = &dummy;
    dummy.next = NULL;
    
    // TODO: 实现迭代合并逻辑
    // 提示：
    // 1. 使用while循环同时遍历l1和l2
    // 2. 比较l1和l2当前节点的值，将较小的连接到tail后面
    // 3. 移动对应链表的指针和tail指针
    // 4. 循环结束后将非空链表直接连接到tail后面
    while (l1 != NULL && l2 != NULL) {
        if (l1->val <= l2->val) {
            tail->next = l1;
            l1 = l1->next;
        } else {
            tail->next = l2;
            l2 = l2->next;
        }
        tail = tail->next;
    }

    // 连接剩余部分（其中一个可能为NULL）
    tail->next = (l1 != NULL) ? l1 : l2;
    
    return dummy.next;
}

/**
 * 解法2：递归法
 * 算法思路：
 * 1. 基准情况：如果任一链表为空，直接返回另一个链表
 * 2. 比较两个链表头节点的值
 * 3. 将较小值节点作为头节点，递归合并剩余部分
 * 4. 返回新的头节点
 * 
 * 时间复杂度：O(m+n)
 * 空间复杂度：O(m+n) - 递归调用栈
 */
struct ListNode* mergeTwoLists_recursive(struct ListNode* l1, struct ListNode* l2) {
    // TODO: 实现递归合并逻辑
    // 提示：
    // 1. 处理基准情况（l1或l2为空）
    // 2. 比较l1和l2的值
    // 3. 如果l1值较小，则l1->next = 递归调用(l1->next, l2)，返回l1
    // 4. 否则l2->next = 递归调用(l1, l2->next)，返回l2
    
    return NULL; // 占位符
}

// 创建新节点辅助函数
struct ListNode* createNode(int val) {
    struct ListNode* node = (struct ListNode*)malloc(sizeof(struct ListNode));
    node->val = val;
    node->next = NULL;
    return node;
}

// 打印链表（用于测试）
void printList(struct ListNode* head) {
    while (head) {
        printf("%d", head->val);
        if (head->next) printf("->");
        head = head->next;
    }
    printf("\n");
}

// 释放链表内存
void freeList(struct ListNode* head) {
    while (head) {
        struct ListNode* temp = head;
        head = head->next;
        free(temp);
    }
}

// 比较两个链表是否相等
int compareLists(struct ListNode* l1, struct ListNode* l2) {
    while (l1 && l2) {
        if (l1->val != l2->val) return 0;
        l1 = l1->next;
        l2 = l2->next;
    }
    return l1 == NULL && l2 == NULL;
}

int main() {
    printf("测试LeetCode #21 合并两个有序链表\n");
    
    // 测试用例1：常规情况
    printf("测试用例1: ");
    struct ListNode* l1 = createNode(1);
    l1->next = createNode(2);
    l1->next->next = createNode(4);
    
    struct ListNode* l2 = createNode(1);
    l2->next = createNode(3);
    l2->next->next = createNode(4);
    
    struct ListNode* result1 = mergeTwoLists_iterative(l1, l2);
    printf("合并结果: ");
    printList(result1);
    
    // 创建预期结果链表 [1,1,2,3,4,4]
    struct ListNode* expected = createNode(1);
    expected->next = createNode(1);
    expected->next->next = createNode(2);
    expected->next->next->next = createNode(3);
    expected->next->next->next->next = createNode(4);
    expected->next->next->next->next->next = createNode(4);
    
    assert(compareLists(result1, expected) == 1);
    printf("测试用例1通过！\n");
    
    freeList(expected);
    // 注意：result1已经包含了l1和l2的所有节点，不需要再次释放
    
    // 测试用例2：空链表情况
    printf("测试用例2: ");
    struct ListNode* l3 = NULL;
    struct ListNode* l4 = createNode(0);
    
    struct ListNode* result2 = mergeTwoLists_iterative(l3, l4);
    printf("合并结果: ");
    printList(result2);
    
    assert(result2 != NULL);
    assert(result2->val == 0);
    assert(result2->next == NULL);
    printf("测试用例2通过！\n");
    
    freeList(result2);
    
    printf("所有测试用例通过！\n");
    return 0;
}
