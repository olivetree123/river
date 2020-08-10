package pocket

import (
	"github.com/olivetree123/river/utils"
	"sync"
)

// pocket: 蓝胖子的口袋，啥都能装

var DataList *dataListStruct
var ConsumingList *consumingListStruct
var GarbageList *garbageListStruct
var FreeList *freeListStruct
var ConsumingPool *consumingPoolStruct


type DataNode struct {
	ID string
	Data interface{}
	Prev *DataNode
}

type PtrNode struct {
	Next *DataNode
}

type dataListStruct struct {
	Head *PtrNode
	Tail *PtrNode
	Persist *PtrNode
}

//type consumingListStruct struct {
//	Head *PtrNode
//	Tail *PtrNode
//}

type consumingPoolStruct struct {
	data map[string]*DataNode
}

type garbageListStruct struct {
	Head *PtrNode
	Tail *PtrNode
}

type freeListStruct struct {
	Head *PtrNode
	Tail *PtrNode
}

func init() {
	DataList = &dataListStruct{
		Head: &PtrNode{nil},
		Tail: &PtrNode{nil},
		Persist: &PtrNode{nil},
	}
	ConsumingPool = &consumingPoolStruct{
		data: make(map[string]*DataNode),
	}
	//ConsumingList = &ConsumingListStruct{
	//	Head: &PtrNode{nil},
	//	Tail: &PtrNode{nil},
	//}
	GarbageList = &garbageListStruct{
		Head: &PtrNode{nil},
		Tail: &PtrNode{nil},
	}
	FreeList = &freeListStruct{
		Head: &PtrNode{nil},
		Tail: &PtrNode{nil},
	}
}

func NewDataNode(data interface{}) *DataNode {
	node := &DataNode{
		ID: utils.NewUUID(),
		Data: data,
	}
	return node
}

func (d *dataListStruct) Push(data interface{}) {
	node := NewDataNode(data)
	if d.Tail == nil {
		d.Head.Next = node
		d.Tail.Next = node
		return
	}
	node.Prev = d.Tail.Next
	d.Tail.Next = node
}

func (d *dataListStruct) Pop() *DataNode {
	if d.Head.Next == nil {
		return nil
	}
	node := d.Head.Next
	d.Head.Next = node.Prev
	return node
}

func (d *dataListStruct) Persistence() {

}

func (p *consumingPoolStruct) Store(node *DataNode) {
	p.data[node.ID] = node
}

func (p *consumingPoolStruct) Load(nodeID string) *DataNode {
	node, ok := p.data[nodeID]
	if !ok {
		return nil
	}
	delete(p.data, nodeID)
	return node
}

func (g *garbageListStruct) Push(node *DataNode) {
	if g.Tail == nil {
		g.Head.Next = node
		g.Tail.Next = node
		return
	}
	node.Prev = g.Tail.Next
	g.Tail.Next = node
}

func (g *garbageListStruct) Pop() *DataNode {
	if g.Head.Next == nil {
		return nil
	}
	node := g.Head.Next
	g.Head.Next = node.Prev
	return node
}

func (f *freeListStruct) Push(node *DataNode) {
	if f.Tail == nil {
		f.Head.Next = node
		f.Tail.Next = node
		return
	}
	node.Prev = f.Tail.Next
	f.Tail.Next = node
}

func (f *freeListStruct) Pop() *DataNode {
	if f.Head.Next == nil {
		return nil
	}
	node := f.Head.Next
	f.Head.Next = node.Prev
	return node
}