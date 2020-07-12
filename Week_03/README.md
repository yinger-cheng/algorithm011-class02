学习笔记
## 第三周课程

### 递归

- 思维要点

	- 1、不要人肉递归(最大误区)
	- 找到最近最简方法，将其拆解成功可重复子问题

		- Why：程序指令只包括if else/for/while loop以及递归调用

	- 数学归纳法思维

- 模板

	- recursion terminator 递归终止条件及return
	- process处理当前逻辑
	- drill down下探到下一层(level、参数)
	- reverse states清扫当前层

### BST

- 中序遍历是递增的

### 分治

### 回溯
回溯公式
result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}
