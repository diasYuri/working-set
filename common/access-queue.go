package common

type AccessQueue struct {
	processId int
	size      int
	head      int
	tail      int
	array     []int
}

func NewAccessQueue(processId int, size int) *AccessQueue {
	aq := &AccessQueue{
		processId: processId,
		size:      size,
		head:      size - 1,
		tail:      size - 1,
		array:     make([]int, size),
	}

	for i := 0; i < size; i++ {
		aq.array[i] = -1
	}

	return aq
}

func (aq *AccessQueue) recalculateTail() {
	if aq.head == aq.tail {
		if aq.tail == 0 {
			aq.tail = aq.size - 1
		} else {
			aq.tail = aq.tail - 1
		}
	}
}

func (aq *AccessQueue) recalculateHead() {
	if aq.head == 0 {
		aq.head = aq.size - 1
	} else {
		aq.head = aq.head - 1
	}
}

func (aq *AccessQueue) Enqueue(pageId int) {
	aq.array[aq.head] = pageId
	aq.recalculateHead()
	aq.recalculateTail()
}

func (aq *AccessQueue) GetHead() int {
	return aq.array[aq.head]
}
