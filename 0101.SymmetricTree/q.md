# 101. 对称二叉树

## 问题描述

给你一个二叉树的根节点 `root` ，检查它是否是轴对称的。

## 示例

**示例 1:**

```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

```
输入：root = [1,2,2,3,4,4,3]
输出：true
```

**示例 2:**

```
    1
   / \
  2   2
   \   \
   3    3
```

```
输入：root = [1,2,2,null,3,null,3]
输出：false
```

## 提示

- 树中节点数目在范围 [1, 1000] 内
- -100 <= Node.val <= 100
