package util

// Enqueue appends an int to the queue
func Enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

// Dequeue pops the first element of the queue
func Dequeue(queue []int) (int, []int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		tmp := []int{}
		return element, tmp
	}
	return element, queue[1:] // Slice off the element once it is dequeued.
}

func Tail(queue []int) int {
	l := len(queue)
	if l == 0 {
		return -1
	} else {
		return l - 1
	}
}
