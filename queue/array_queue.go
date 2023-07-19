package queue

type ArrayQueue struct {
	datas []interface{}
}

func (a *ArrayQueue) GetSize() int {
	return len(a.datas)
}

func (a *ArrayQueue) IsEmpty() bool {
	if len(a.datas) == 0 {
		return true
	} else {
		return false
	}
}

func (a *ArrayQueue) EnQueue(data interface{}) {
	a.datas = append(a.datas, data)
}

func (a *ArrayQueue) DeQueue() interface{} {
	if len(a.datas) == 0 {
		panic("Queue is empty")
	}
	delData := a.datas[0]
	a.datas = a.datas[1:]
	return delData
}

func (a *ArrayQueue) GetFront() interface{} {
	seeDate := a.datas[0]
	return seeDate
}
