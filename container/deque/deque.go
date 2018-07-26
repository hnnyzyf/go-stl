package deque

import (
	"math"
	"sync"
)

const (
	MapSize    = 8
	ChunckSize = 128
)

type chunck []interface{}

var dequePool = &sync.Pool{
	New: func() interface{} {
		return make(chunck, ChunckSize)
	},
}

type pos struct {
	chunck int
	index  int
}

type Deque struct {
	mmap []chunck

	begin *pos
	end   *pos
}

//the Deque is a double-ended queue
//the init size is MapSize,which is a const value
func New() *Deque {
	return &Deque{
		mmap: make([]chunck, MapSize),

		begin: &pos{
			chunck: (MapSize + 1) / 2,
			index:  0,
		},
		end: &pos{
			chunck: (MapSize - 1) / 2,
			index:  ChunckSize - 1,
		},
	}
}

//PushBack add a new val in the back
func (d *Deque) PushBack(val interface{}) {
	//realloc if need
	if d.end.chunck == len(d.mmap)-1 && d.end.index == ChunckSize-1 {
		d.reallocmmap()
	}

	//find the new position
	d.end.chunck, d.end.index = d.end.chunck+(d.end.index+1)/ChunckSize, (d.end.index+1)%ChunckSize

	//alloc memory
	if d.mmap[d.end.chunck] == nil {
		d.mmap[d.end.chunck] = dequePool.Get().(chunck)
	}

	//add new val
	d.mmap[d.end.chunck][d.end.index] = val
}

//PushFront add a new val in the front
func (d *Deque) PushFront(val interface{}) {
	//realloc if need
	if d.begin.chunck == 0 && d.begin.index == 0 {
		d.reallocmmap()
	}

	//find the new position
	d.begin.chunck, d.begin.index = d.begin.chunck+(d.begin.index-ChunckSize)/ChunckSize, (d.begin.index-1+ChunckSize)%ChunckSize

	//alloc memory
	if d.mmap[d.begin.chunck] == nil {
		d.mmap[d.begin.chunck] = dequePool.Get().(chunck)
	}
	//add new val
	d.mmap[d.begin.chunck][d.begin.index] = val
}

//PopBack delete a new val in the back
func (d *Deque) PopBack() (interface{}, bool) {
	//revoke used memory if need
	defer d.reallocmmap()

	if d.end.chunck >= d.begin.chunck && d.end.index >= d.begin.index {
		val := d.mmap[d.end.chunck][d.end.index]
		chunck, index := d.end.chunck+(d.end.index-ChunckSize)/ChunckSize, (d.end.index-1+ChunckSize)%ChunckSize

		//if old chunck has never been used ,try to revoke it
		if chunck != d.end.chunck {
			dequePool.Put(d.mmap[d.end.chunck])
			d.mmap[d.end.chunck] = nil
		}

		d.end.chunck = chunck
		d.end.index = index
		return val, true
	}

	return nil, false
}

//PopFront delete a new val in the front
func (d *Deque) PopFront() (interface{}, bool) {
	//revoke used memory if need
	defer d.reallocmmap()

	if d.begin.chunck <= d.end.chunck && d.begin.index <= d.end.index {
		val := d.mmap[d.begin.chunck][d.begin.index]
		chunck, index := d.begin.chunck+(d.begin.index+1)/ChunckSize, (d.begin.index+1)%ChunckSize

		//if old chunck has never been used ,try to revoke it
		if chunck != d.begin.chunck {
			dequePool.Put(d.mmap[d.begin.chunck])
			d.mmap[d.begin.chunck] = nil
		}

		d.begin.chunck = chunck
		d.begin.index = index
		return val, true
	}

	return nil, false
}

func (d *Deque) First() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	return d.mmap[d.begin.chunck][d.begin.index], true
}

func (d *Deque) Last() (interface{}, bool) {
	if d.IsEmpty() {
		return nil, false
	}
	return d.mmap[d.end.chunck][d.end.index], true
}

func (d *Deque) Get(i int) interface{} {
	if i > d.Len()-1 {
		panic("Deque index out of range")
	}
	chunck, index := d.end.chunck+(d.end.index+i)/ChunckSize, (d.end.index+i)%ChunckSize
	return d.mmap[chunck][index]
}

func (d *Deque) Len() int {
	return ChunckSize*(d.end.chunck-d.begin.chunck-1) + (ChunckSize - d.begin.index) + (d.end.index + 1)
}

func (d *Deque) IsEmpty() bool {
	if d.end.chunck < d.begin.chunck {
		return true
	} else if d.end.chunck == d.begin.chunck && d.begin.index > d.end.index {
		return true
	} else {
		return false
	}
}

//reallocmmap malloc memory and revoke memory
func (d *Deque) reallocmmap() {

	//end has no space and begin has space
	if d.end.chunck == len(d.mmap)-1 && d.begin.chunck >= 1 {
		//cal the offset
		offset := (d.begin.chunck + 1) / 2

		//copy all between d.begin.chunck and d.end.chunck
		copy(d.mmap[d.begin.chunck-offset:d.end.chunck+1-offset], d.mmap[d.begin.chunck:d.end.chunck+1])

		//set nil
		for i := range d.mmap[len(d.mmap)-offset:] {
			d.mmap[i+len(d.mmap)-offset] = nil
		}

		//reindex
		d.begin.chunck -= offset
		d.end.chunck -= offset

		//begin has no space and end has space
	} else if d.begin.chunck == 0 && d.end.chunck <= len(d.mmap)-2 {

		//cal the offset
		offset := (len(d.mmap) - d.end.chunck) / 2

		//copy all between d.begin.chunck and d.end.chunck
		copy(d.mmap[offset:d.end.chunck+offset+1], d.mmap[:d.end.chunck+1])

		//set nil
		for i := range d.mmap[:offset] {
			d.mmap[i] = nil
		}

		//reindex
		d.begin.chunck += offset
		d.end.chunck += offset

	} else if d.end.chunck == len(d.mmap)-1 && d.begin.chunck == 0 {

		//we need to relloc a new map,add two node
		var mmap []chunck

		//if size of d.mmap is smaller than 1024,we double every time
		//if size of d.mmap is bigger than 1024,we add 25% every time
		if len(d.mmap) < 1024 {
			mmap = make([]chunck, 2*len(d.mmap))
		} else {
			mmap = make([]chunck, len(d.mmap)/4+len(d.mmap))
		}

		//cal offset
		offset := float64(len(mmap)-len(d.mmap)) / 2
		d.begin.chunck = int(math.Floor(offset))
		d.end.chunck = len(d.mmap) + int(math.Ceil(offset)) - 1

		//copy all into mmap
		copy(mmap[d.begin.chunck:d.end.chunck+1], d.mmap)

		d.mmap = mmap

		//revoke unused memory when there are only half chuncks have been used
	} else if len(d.mmap) > 1024 && d.end.chunck-d.begin.chunck+1 < len(d.mmap)/2 && len(d.mmap) > MapSize {
		//new mmap
		mmap := make([]chunck, len(d.mmap)-len(d.mmap)/4)

		//cal offset
		offset := (len(d.mmap) - len(mmap)) / 2
		begin := offset
		end := offset + d.end.chunck - d.begin.chunck

		//copy
		copy(mmap[begin:end+1], d.mmap[d.begin.chunck:d.end.chunck+1])

		//reindex
		d.begin.chunck = begin
		d.end.chunck = end

		d.mmap = mmap
	} else {
		//do nothing
	}

}
