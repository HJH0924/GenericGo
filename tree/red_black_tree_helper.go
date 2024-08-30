// Package set
/**
* @Project : GenericGo
* @File    : red_black_tree_helper.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/30 14:34
**/

package tree

// nodeColor 定义了红黑树节点的颜色类型，使用布尔值表示。
type nodeColor bool

const (
	Red   nodeColor = false
	Black nodeColor = true
)

// rbNode 定义了红黑树的节点结构，使用泛型允许存储任意类型的键和值。
type rbNode[K any, V any] struct {
	color  nodeColor
	key    K
	val    V
	left   *rbNode[K, V]
	right  *rbNode[K, V]
	parent *rbNode[K, V]
}

func newRBNode[K any, V any](key K, val V) *rbNode[K, V] {
	return &rbNode[K, V]{
		color:  Red,
		key:    key,
		val:    val,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}
