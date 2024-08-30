// Package set
/**
* @Project : GenericGo
* @File    : red_black_tree.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/29 14:12
**/

package tree

import genericgo "github.com/HJH0924/GenericGo"

// RBTree 定义了红黑树的结构
type RBTree[K any, V any] struct {
	root    *rbNode[K, V]           // 指向树的根节点
	compare genericgo.Comparator[K] // 用于比较键的比较器
	size    int                     // 树中节点的数量
}

func NewRBTree[K any, V any](compare genericgo.Comparator[K]) *RBTree[K, V] {
	return &RBTree[K, V]{
		root:    nil,
		compare: compare,
		size:    0,
	}
}

func (Self *RBTree[K, V]) Size() int {
	if Self.root == nil {
		return 0
	}
	return Self.size
}

func (Self *RBTree[K, V]) Size() int {
	if Self.root == nil {
		return 0
	}
	return Self.size
}
