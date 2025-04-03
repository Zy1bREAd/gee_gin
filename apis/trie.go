package apis

import (
	"fmt"
	"strings"
)

type PrefixTrieNode struct {
	isWild   bool   // 是否精准匹配，若:和*则是动态匹配的。
	pattern  string // 当前路径下的完整pattern
	part     string // 当前节点对应的路由路径
	children []*PrefixTrieNode
}

func NewPatternTrieNodeForOriginal(p string) *PrefixTrieNode {
	return &PrefixTrieNode{
		isWild:   false,
		pattern:  p,
		children: make([]*PrefixTrieNode, 0),
	}
}

// 新建并初始化根目录前缀树节点
func NewPatternTrieRoot() *PrefixTrieNode {
	return NewPatternTrieNodeForOriginal("/")
}
func NewPatternTrieNode() *PrefixTrieNode {
	return NewPatternTrieNodeForOriginal("")
}

func (t *PrefixTrieNode) SplitPattern(fullPattern string) []string {
	splitResult := strings.Split(fullPattern, "/")
	return splitResult
}

func (t *PrefixTrieNode) Insert(fullPattern string, parts []string, height int) {
	// 拆分路由
	curr := t
	fmt.Println("parts=", parts)
	// 递归终止条件：当层级高度等同于待插入路由的树高度。到达叶子节点（通常叶子节点存储完整路径，表示这是完整的pattern）
	if len(parts) == height {
		curr.pattern = fullPattern
		return
	}
	// 插入核心逻辑（本质上判断是否有对应节点，没有就插入）
	part := parts[height]          // 获取当前层级的路由路径
	child := curr.matchChild(part) // 判断当前孩子节点中是否有对应的路由路径，没有则插入
	fmt.Println("your child=", child)
	if child == nil {
		child = &PrefixTrieNode{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':', // 判断是否为动态路由参数
		}
		// 添加孩子节点
		curr.children = append(curr.children, child)
	}

	// 递归操作下一层级
	curr.Insert(fullPattern, parts, height+1)
}

func (t *PrefixTrieNode) Search(parts []string, height int) *PrefixTrieNode {
	curr := t
	fmt.Println(curr.debugTrie())
	// 终止条件：当遍历到叶子节点时仍然没有找到返回nil
	if len(parts) == height || strings.HasPrefix(curr.part, "*") {
		// 当pattern为空代表不是一个完整的url，所以返回nil
		if curr.pattern == "" {
			return nil
		}
		return curr
	}
	// 获取当前遍历的层级
	part := parts[height]
	// 返回所有匹配的孩子节点
	childrens := curr.matchChildren(part)
	// 子节点匹配
	for _, child := range childrens {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil

}

// 从孩子节点中找到一个匹配的节点，用于插入
func (t *PrefixTrieNode) matchChild(part string) *PrefixTrieNode {
	curr := t
	// 终止条件
	for _, child := range curr.children {
		// 检查孩子节点中是否存在我这个pattern
		// 为什么判断条件要加入isWild？因为给定isWild属性是含有动态匹配项则设置会true，例如pattern中含有:或*
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 查找孩子节点中所有匹配的，用于查找
func (t *PrefixTrieNode) matchChildren(part string) []*PrefixTrieNode {
	result := make([]*PrefixTrieNode, 0)
	curr := t
	for _, child := range curr.children {
		if child.part == part || child.isWild {
			result = append(result, child)
		}
	}
	return result
}

// 遍历路由树，收集所有有效路由节点。
func (t *PrefixTrieNode) travel(list *[]*PrefixTrieNode) {
	curr := t
	if t.pattern != "" {
		// 如果当前节点是有效的pattern，将其加入数组
		*list = append(*list, curr)
	}
	for _, child := range curr.children {
		child.travel(list)
	}
}

func (t *PrefixTrieNode) debugTrie() string {
	return fmt.Sprintf("<DEBUG> - pattern=%s - part=%s - childrents=%v - isWild=%t", t.pattern, t.part, t.children, t.isWild)
}
