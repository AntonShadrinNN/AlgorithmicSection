package structures

type list struct {
	Priority int
	data     TagStat
}

type priorityQueue struct {
	lists    []list
	heapSize int
}

//func NewPriorityQueue() *priorityQueue {
//	return &priorityQueue{}
//}

func NewPriorityQueue(arr []TagStat) *priorityQueue {
	lists := make([]list, len(arr))
	for i, val := range arr {
		l := list{Priority: val.OccurrenceCount, data: val}
		lists[i] = l
	}

	q := priorityQueue{lists: lists, heapSize: len(lists)}

	for i := q.heapSize / 2; i >= 0; i-- {
		q.heapify(i)
	}
	return &q
}

func (h *priorityQueue) IsEmpty() bool {
	return h.heapSize == 0
}

func (h *priorityQueue) AddList(value TagStat) {
	h.lists = append(h.lists, list{value.OccurrenceCount, value})
	i := h.heapSize
	parent := (i - 1) / 2
	for i > 0 && h.lists[parent].Priority < h.lists[i].Priority {
		temp := h.lists[i]
		h.lists[i] = h.lists[parent]
		h.lists[parent] = temp

		i = parent
		parent = (i - 1) / 2
	}
	h.heapSize += 1
}

func (h *priorityQueue) Pop() (maxElement TagStat) {
	maxElement = h.lists[0].data
	h.lists[0] = h.lists[h.heapSize-1]
	h.lists = h.lists[:h.heapSize-1]
	h.heapSize--
	h.heapify(0)
	return

}

func (h *priorityQueue) Poop(n int) []TagStat {
	t := make([]TagStat, n)
	for i := 0; i < n; i++ {
		t[i] = h.Pop()
	}
	return t
}

func (h *priorityQueue) heapify(i int) {
	var l, r, largest int
	for {
		l = 2*i + 1
		r = 2*i + 2
		largest = i
		if l < h.heapSize && h.lists[l].Priority > h.lists[largest].Priority {
			largest = l
		}
		if r < h.heapSize && h.lists[r].Priority > h.lists[largest].Priority {
			largest = r
		}

		if largest == i {
			break
		}
		h.lists[i], h.lists[largest] = h.lists[largest], h.lists[i]
		i = largest
	}
}
