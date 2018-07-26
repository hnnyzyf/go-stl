package deque

import (
	"sync"
)

const (
	MapSize    = 8
	ChunckSize = 4
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

//PushFront add a new val in the front
func (d *Deque) PushFront(val interface{}) {
	//find the new position
	d.begin.index -= 1
	if d.begin.index < 0 {
		d.begin.index = ChunckSize + d.begin.index
		d.begin.chunck -= 1
	}
	//alloc memory
	if d.mmap[d.begin.chunck] == nil {
		d.mmap[d.begin.chunck] = dequePool.Get().(chunck)
	}
	//add new val
	d.mmap[d.begin.chunck][d.begin.index] = val

	//realloc if need
	if d.begin.chunck == 0 && d.begin.index == 0 {
		d.reallocmmap()
	}
}

//PushBack add a new val in the back
func (d *Deque) PushBack(val interface{}) {
	//find the new position
	d.end.index += 1
	if d.end.index == ChunckSize {
		d.end.index = 0
		d.end.chunck += 1
	}

	//alloc memory
	if d.mmap[d.end.chunck] == nil {
		d.mmap[d.end.chunck] = dequePool.Get().(chunck)
	}

	//add new val
	d.mmap[d.end.chunck][d.end.index] = val

	//realloc if need
	if d.end.chunck == len(d.mmap)-1 && d.end.index == ChunckSize-1 {
		d.reallocmmap()
	}
}

//PopBack delete a new val in the back
func (d *Deque) PopBack() (interface{}, bool) {
	//revoke used memory if need
	defer d.reallocmmap()

	if d.Len() != 0 {
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
	//defer d.reallocmmap()

	if d.Len() != 0 {
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

//Get access element
func (d *Deque) Get(i int) interface{} {
	if i >= d.Len() {
		panic("Deque index out of range")
	}

	index := (i + d.begin.index) % ChunckSize
	nchunck := d.begin.chunck + (i+d.begin.index)/ChunckSize
	return d.mmap[nchunck][index]
}

//Len return size
func (d *Deque) Len() int {
	return ChunckSize*(d.end.chunck-d.begin.chunck) - d.begin.index + d.end.index + 1
}

//IsEmpty test whether container is empty
func (d *Deque) IsEmpty() bool {
	return d.Len() == 0
}

//reallocmmap malloc memory and revoke memory
func (d *Deque) reallocmmap() {

	//end has no space and begin has space
	if d.end.chunck == len(d.mmap)-1 && d.begin.chunck >= 1 {
		//data section
		data := d.mmap[d.begin.chunck : d.end.chunck+1]

		//cal the offset and get destionation section
		offset := d.begin.chunck/2 + d.begin.chunck%2
		d.begin.chunck -= offset
		d.end.chunck -= offset
		destionation := d.mmap[d.begin.chunck : d.end.chunck+1]
		//copy all between d.begin.chunck and d.end.chunck
		copy(destionation, data)

		//clears
		for i := 1; i <= offset; i++ {
			d.mmap[d.end.chunck+i] = nil
		}

		//begin has no space and end has space
	} else if d.begin.chunck == 0 && len(d.mmap)-d.end.chunck > 1 {
		//data section
		data := d.mmap[d.begin.chunck : d.end.chunck+1]

		//cal the offset and get destionation section
		offset := (len(d.mmap)-d.end.chunck-1)/2 + (len(d.mmap)-d.end.chunck-1)%2
		d.begin.chunck += offset
		d.end.chunck += offset
		destionation := d.mmap[d.begin.chunck : d.end.chunck+1]

		//copy all between d.begin.chunck and d.end.chunck
		copy(destionation, data)

		//clear
		for i := 0; i < offset; i++ {
			d.mmap[i] = nil
		}

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
		offset := (len(mmap) - len(d.mmap)) / 2
		d.begin.chunck += offset
		d.end.chunck += offset
		destionation := mmap[d.begin.chunck : d.end.chunck+1]

		//copy all into mmap
		copy(destionation, d.mmap)

		d.mmap = mmap
		//revoke unused memory when there are only half chuncks have been used
	} else if d.end.chunck-d.begin.chunck < len(d.mmap)/2 && len(d.mmap) > MapSize {

		//data section
		data := d.mmap[d.begin.chunck : d.end.chunck+1]

		////cal the offset and get destionation section
		mmap := make([]chunck, len(d.mmap)-len(d.mmap)/4)
		offset := (len(d.mmap) - len(mmap)) / 2
		d.end.chunck = d.end.chunck - d.begin.chunck + offset
		d.begin.chunck = offset
		destionation := mmap[d.begin.chunck : d.end.chunck+1]

		//copy
		copy(destionation, data)
		d.mmap = mmap
	} else {
		//do nothing
	}

}
