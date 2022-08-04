package queue

type Queue []interface{}

func (q *Queue) Push(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
