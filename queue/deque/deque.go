package deque

import (
	"fmt"
	"log"
	"strings"
)

// 双端队列
type deque struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func New(cap int) Deque {
	return &deque{
		data: make([]interface{}, cap),
	}
}

func Init() Deque {
	return New(10)
}

func (d *deque) resize(n int) {
	newData := make([]interface{}, n)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.front+i)%len(d.data)]
	}
	d.data = newData
	d.front = 0
	d.tail = d.size
}

func (d *deque) AddLast(e interface{}) {
	if d.size == d.Cap() {
		d.resize(d.Cap() * 2)
	}
	d.data[d.tail] = e
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

func (d *deque) AddFront(e interface{}) {
	if d.size == d.Cap() {
		d.resize(d.Cap() * 2)
	}
	// 处理队首指针
	if d.front == 0 {
		d.front = len(d.data) - 1
	} else {
		d.front--
	}
	d.data[d.front] = e
	d.size++
}

func (d *deque) RemoveFront() interface{} {
	if d.IsEmpty() {
		log.Fatal("deque is empty")
	}
	res := d.data[d.front]
	d.data[d.front] = nil
	d.front = (d.front + 1) % len(d.data)
	d.size--
	if d.size == d.Cap()/4 && d.Cap()/2 != 0 {
		d.resize(d.Cap() / 2)
	}
	return res
}

func (d *deque) RemoveLast() interface{} {
	if d.IsEmpty() {
		log.Fatal("deque is empty")
	}
	if d.tail == 0 {
		d.tail = len(d.data) - 1
	} else {
		d.tail--
	}
	res := d.data[d.tail]
	d.data[d.tail] = nil
	d.size--
	if d.size == d.Cap()/4 && d.Cap()/2 != 0 {
		d.resize(d.Cap() / 2)
	}
	return res
}

func (d *deque) GetFront() interface{} {
	if d.IsEmpty() {
		log.Fatal("deque is empty")
	}
	return d.data[d.front]
}

func (d *deque) GetLast() interface{} {
	if d.IsEmpty() {
		log.Fatal("deque is empty")
	}
	var i int
	if d.tail == 0 {
		i = len(d.data) - 1
	} else {
		i = d.tail - 1
	}
	return d.data[i]
}

func (d *deque) Len() int {
	return d.size
}

func (d *deque) Cap() int {
	return len(d.data)
}

func (d *deque) IsEmpty() bool {
	return d.size == 0
}

func (d *deque) String() string {
	var res strings.Builder
	res.WriteString(fmt.Sprintf("Deque: size: %d, capacity: %d\n", d.size, d.Cap()))
	res.WriteString("front [")
	for i := 0; i < d.size; i++ {
		res.WriteString(fmt.Sprintf("%v", d.data[(d.front+i)%len(d.data)]))
		if i != d.size-1 {
			res.WriteString(", ")
		}
	}
	res.WriteString("] tail")
	return res.String()
}
