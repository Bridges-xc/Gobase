package main

// 类型约束接口
// 泛型接口
type Queuep[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q
}
