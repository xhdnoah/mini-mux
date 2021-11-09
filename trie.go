package mux

import "strings"

type node struct {
	children   []*node
	part       string // 路由片段
	pattern    string // 待匹配路由
	isWildcard bool   // 含有 ：或 * 为 true
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWildcard {
			return child
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWildcard {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 注册路由时插入：递归查找每一层节点，没有匹配当前 part 就新建一个
// 匹配结束时，用 n.pattern != "" 判断成功
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWildcard: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 为空值说明遇到通配符但还没到叶子节点
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		// 递归查询
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
