// Package list
/**
* @Project : GenericGo
* @File    : linked_list.go
* @IDE     : GoLand
* @Author  : Tvux
* @Date    : 2024/8/16 10:05
**/

package list

import "github.com/HJH0924/GenericGo/errs"

var (
	_ List[any] = &LinkedList[any]{}
)

// node 定义双向循环链表的节点结构
type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

// LinkedList 定义双向循环链表结构
type LinkedList[T any] struct {
	head   *node[T]
	tail   *node[T]
	length int
}

// Append 在 LinkedList 末尾追加一个或多个元素。
func (ll *LinkedList[T]) Append(vals ...T) {
	for _, val := range vals {
		newNode := &node[T]{
			prev: ll.tail.prev,
			next: ll.tail,
			val:  val,
		}
		ll.tail.prev = newNode
		newNode.prev.next = newNode
		ll.length++
	}
}

// Add 在特定下标处增加一个新元素。
// 如果下标超出合法范围，返回错误。
// 如果 idx 等于 LinkedList 长度，则表示往 LinkedList 末端增加元素。
func (ll *LinkedList[T]) Add(idx int, val T) error {
	if idx < 0 || idx > ll.length {
		return errs.NewErrIndexOutOfRange(ll.length, idx)
	}
	if idx == ll.length {
		ll.Append(val)
		return nil
	}
	p := ll.getNodeAt(idx)
	newNode := &node[T]{
		prev: p.prev,
		next: p,
		val:  val,
	}
	p.prev.next = newNode
	p.prev = newNode
	ll.length++
	return nil
}

// Delete 删除指定下标的元素，并返回被删除的元素。
// 如果下标超出合法范围，返回错误。
func (ll *LinkedList[T]) Delete(idx int) (deletedVal T, err error) {
	if idx < 0 || idx >= ll.length {
		return deletedVal, errs.NewErrIndexOutOfRange(ll.length, idx)
	}
	p := ll.getNodeAt(idx)
	p.prev.next = p.next
	p.next.prev = p.prev
	p.prev, p.next = nil, nil
	ll.length--
	return p.val, nil
}

// Set 重置指定下标位置的元素为 val。
// 如果下标超出合法范围，返回错误。
func (ll *LinkedList[T]) Set(idx int, val T) error {
	if idx < 0 || idx >= ll.length {
		return errs.NewErrIndexOutOfRange(ll.length, idx)
	}
	p := ll.getNodeAt(idx)
	p.val = val
	return nil
}

// Get 返回对应下标的元素。
// 如果下标超出合法范围，返回错误。
func (ll *LinkedList[T]) Get(idx int) (val T, err error) {
	if idx < 0 || idx >= ll.length {
		return val, errs.NewErrIndexOutOfRange(ll.length, idx)
	}
	p := ll.getNodeAt(idx)
	return p.val, nil
}

// Len 返回 LinkedList 中元素的数量。
func (ll *LinkedList[T]) Len() int {
	return ll.length
}

// Cap 方法返回 LinkedList 的容量。
// 默认情况下，链表的容量等于其长度。
func (ll *LinkedList[T]) Cap() int {
	return ll.Len()
}

// Range 遍历 LinkedList 的所有元素，并使用给定的函数访问每个元素。
func (ll *LinkedList[T]) Range(onVal func(idx int, val T) error) error {
	for p, i := ll.head.next, 0; i < ll.length; i++ {
		err := onVal(i, p.val)
		if err != nil {
			return err
		}
		p = p.next
	}
	return nil
}

// AsSlice 将 LinkedList 转化为一个新切片，即使 LinkedList 为空，也返回一个长度和容量都为0的切片。
func (ll *LinkedList[T]) AsSlice() []T {
	res := make([]T, ll.length)
	// 使用 i < ll.length 可以防止索引越界
	// 如果使用 p != ll.tail 则无法检查出索引越界的错误，导致 panic
	for p, i := ll.head.next, 0; i < ll.length; i++ {
		res[i] = p.val
		p = p.next
	}
	return res
}

// NewLinkedList 创建并返回一个新的LinkedList实例。
func NewLinkedList[T any]() *LinkedList[T] {
	// 创建头尾节点
	head := &node[T]{}
	tail := &node[T]{}

	// 相互指向
	head.prev = tail
	head.next = tail

	tail.prev = head
	tail.next = head

	return &LinkedList[T]{
		head: head,
		tail: tail,
	}
}

// NewLinkedListOf 创建一个新的LinkedList实例，并使用提供的切片vals作为初始元素。
// 注意，如果vals包含任何元素，这些元素将被添加到链表的尾部。
func NewLinkedListOf[T any](vals []T) *LinkedList[T] {
	ll := NewLinkedList[T]()
	ll.Append(vals...)
	return ll
}

// getNodeAt 返回双向循环链表中索引为 i 的节点。
// 因为该函数只供内部使用，所以在调用该函数之前已经确保了索引合法
func (ll *LinkedList[T]) getNodeAt(idx int) *node[T] {
	var p *node[T]
	if idx < ll.length/2 {
		// 从前往后找
		p = ll.head // 对应索引 -1
		for i := -1; i < idx; i++ {
			p = p.next
		}
		return p
	} else {
		// 从后往前找
		p = ll.tail // 对应索引 ll.length
		for i := ll.length; i > idx; i-- {
			p = p.prev
		}
		return p
	}
}
