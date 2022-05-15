# 动态规划

- 状态
- 状态转移方程

## 有、无后效性

**无后效性**： 前边状态的结果是确定的，后边状态能从前边状态推导出新的结果
**有后效性**： 前边状态结果不确定，后边状态结果无法直接从前边推导出来

解决「有后效性」的办法是固定住需要分类讨论的地方，记录下更多的结果。在代码层面上表现为：
状态数组增加维度，例如：「力扣」的股票系列问题；
把状态定义得更细致、准确，例如：前天推送的第 124 题：状态定义只解决路径来自左右子树的其中一个子树。