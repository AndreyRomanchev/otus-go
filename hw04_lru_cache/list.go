package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value      interface{}
	Next, Prev *listItem
}

type list struct {
	front, back *listItem
	len         int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *listItem {
	return l.front
}

func (l list) Back() *listItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *listItem {
	i := listItem{Value: v}

	if l.front == nil {
		l.back = &i
	} else {
		i.Prev = l.front
		l.front.Next = &i
	}

	l.front = &i
	l.len++

	return &i
}

func (l *list) PushBack(v interface{}) *listItem {
	i := listItem{Value: v}

	if l.back == nil {
		l.front = &i
	} else {
		i.Next = l.back
		l.back.Prev = &i
	}

	l.back = &i
	l.len++

	return &i
}

func (l *list) Remove(i *listItem) {
	l.len--

	if i.Next == nil {
		l.front = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	if i.Prev == nil {
		l.back = i.Next
	} else {
		i.Prev.Next = i.Next
	}
}

func (l *list) MoveToFront(i *listItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
