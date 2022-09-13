package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root != nil {
		r := inorderTraversal(root.Left)
		r = append(r, root.Val)
		r = append(r, inorderTraversal(root.Right)...)
		return r
	} else {
		return []int{}
	}
}

// 递归
var MAX_LEVEL = 10

//func recursion(level, param1, param2 int){
//	//递归终止条件
//	if level > MAX_LEVEL{
//		result = append(result, ) //result添加和处理
//		return
//	}
//	//当前层的逻辑处理部分
//
//	//...
//
//	//进入下一层
//	recursion(level+1, param1, param2)
//
//	//返回当前层状态
//	return
//}

//分治

//func divide_conquer(problem , param1, param2, ...,int) int {
//	//分治终止条件
//	if problem == nil:
//	return value
//	//当前层的逻辑处理部分
//	data := prepare_data(problem)
//	subs := split_problem(problem, data)
//	//进入下一层计算子问题
//	sub1 := divide_conquer(subs[0], param1, param2, ...)
//	sub2 := divide_conquer(subs[1], param1, param2, ...)
//	sub3 := divide_conquer(subs[2], param1, param2, ...)
//	...
//	//最终结果计算
//	result := process_result(sub1, sub2, sub3, ...)
//	return	result
//}

//回溯
//func backtracking(item int) bool {
//	#回溯终止条件
//	if item == len(LIST):
//	return true
//	#当前层的逻辑处理部分
//	index := prepare_index(item)
//	#回溯
//	for f in OUTSIDE_LIST(index):
//	#改变result
//	change_result(result)
//	#进入下一层
//	if recall(item+1):
//	return True
//	#回复刚刚改过的result
//	reply_result(result)
//	#返回当前层状态
//	return	false
//}

//深度优先搜素
//def dfs(node *TreeNode, visited map[*TreeNode]int):
//if _,ok:=visited[node];ok{ //终止条件
////这个节点已经拜访了
//return
//}
////记录拜访节点
//visited[node] = 0
////当前层的逻辑处理部分
//...
//for _,next_node := range node.Children{
//if _,ok:=visited[next_node];!ok{
//dfs(next_node, visited)
//}
//}

//广度优先搜索
////第一种写法
//func bfs(root *TreeNode){
//	queue := []*TreeNode{}
//	queue = append(queue,root)
//	for len(queue) != 0{
//		node := queue[0]
//		queue = queue[1:]
//		process(node)
//		nodes := node.Children
//		queue = append(queue,nodes...)
//	}
//	//其他工作
//}
////第二种写法
//func bfs(root *TreeNode){
//	queue := []*TreeNode{}
//	queue = append(queue,root)
//	for len(queue) != 0{
//		child := []*TreeNode{}
//		for _,node :=range queue{
//			process(node)
//			nodes := node.Children
//			child = append(child,nodes...)
//		}
//		queue = child
//	}
//	//其他工作
//}
