package utils

import (
	"strings"
)

type PrefixTrieNode struct {
	isEnd    bool // 表示该节点是否为结尾
	pattern  string
	children map[string]*PrefixTrieNode // 核心！维护一个哈希表来表示孩子节点(节点value -> 节点所在位置)
}

func NewPatternTrieNodeForOriginal(p string) *PrefixTrieNode {
	return &PrefixTrieNode{
		isEnd:    false,
		pattern:  p,
		children: make(map[string]*PrefixTrieNode),
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

func (t *PrefixTrieNode) Insert(fullPattern string) {
	curr := t
	// 将完整的Pattern解析成数组传入遍历进行插入
	patterns := t.SplitPattern(fullPattern)
	for _, v := range patterns {
		// 若不存在则添加
		if _, ok := curr.children[v]; !ok {
			curr.children[v] = NewPatternTrieNode()
		}
		// 如果存在，则切换对应的孩子节点上去
		curr = curr.children[v]
	}
	// 遍历完word后，确认结尾并设置end flag
	curr.isEnd = true
}

func (t *PrefixTrieNode) Search(fullPattern string) bool {
	patterns := t.SplitPattern(fullPattern)
	curr := t
	for _, p := range patterns {
		if _, ok := curr.children[p]; !ok {
			return false
		}
		curr = curr.children[p]
	}
	return curr.isEnd

}

// 前缀pattern匹配路由
func (t *PrefixTrieNode) SearchWithPrefix(prefix string) []string {
	patternResult := make([]string, 0)
	curr := t
	patterns := t.SplitPattern(prefix)
	// 判断前缀是否属于pattern
	for _, v := range patterns {

	}
	if curr.pattern == prefix && curr.isEnd {
		patternResult = append(patternResult, prefix)
	}
	// 递归检查孩子节点中匹配的路由

}
