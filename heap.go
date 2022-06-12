package Heap

import (
	"errors"
	"fmt"
)

type Type int

const (
	MAX Type = 1
	MIN Type = 2

	RIGHT_C = 3
	LEFT_C  = 4
	HALT    = 5
)

type Heap struct {
	Nodes       []int
	Type        Type
	lastNodeIdx int
}

func NewHeap(heapType Type) *Heap {
	return &Heap{
		Nodes:       make([]int, 0),
		Type:        heapType,
		lastNodeIdx: -1,
	}
}

func (h *Heap) WithElements(elements []int) *Heap {
	for _, elem := range elements {
		h.Insert(elem)
	}
	return h
}

func (h *Heap) Sort() []int {
	fmt.Println("h.lastNodeIdx", h.lastNodeIdx)
	for i := 0; i < len(h.Nodes); i++ {
		elem, err := h.Delete()
		if err != nil {
			break
		}

		h.Nodes[len(h.Nodes)-1-i] = elem
	}
	return h.Nodes
}

func (h *Heap) Insert(element int) {
	fmt.Println("Inserting ", element)
	h.Nodes = append(h.Nodes, element)
	h.lastNodeIdx++
	if h.lastNodeIdx == 0 {
		return
	}
	h.adjust()
}

func (h *Heap) adjust() {
	childIdx := h.lastNodeIdx
	parentIdx := getParent(childIdx)
	for h.adjustCondition(childIdx, parentIdx) {
		h.swap(childIdx, parentIdx)
		childIdx = parentIdx
		parentIdx = (childIdx - 1) / 2
	}
}

func (h *Heap) swap(i, j int) {
	temp := h.Nodes[i]
	h.Nodes[i] = h.Nodes[j]
	h.Nodes[j] = temp
}

func (h *Heap) Print() {
	fmt.Println(h.Nodes[:h.lastNodeIdx+1])
}

func (h *Heap) Delete() (int, error) {
	if h.lastNodeIdx == -1 {
		return -1, errors.New("heap is empty")
	}
	deletedNode := h.Nodes[0]
	h.Nodes[0] = h.Nodes[h.lastNodeIdx]
	parentIdx := 0
	for {
		swap := h.topdownAdjustCondition(parentIdx)
		if swap == RIGHT_C {
			h.swap(parentIdx, 2*parentIdx+2)
			parentIdx = 2*parentIdx + 2
		}

		if swap == LEFT_C {
			h.swap(parentIdx, 2*parentIdx+1)
			parentIdx = 2*parentIdx + 1
		}
		if swap == HALT {
			break
		}
	}
	h.lastNodeIdx--
	return deletedNode, nil
}

func (h *Heap) GetMin() int {
	return 0
}

func (h *Heap) adjustCondition(child, parent int) bool {
	if h.Type == MAX {
		return h.Nodes[child] > h.Nodes[parent]
	}
	return h.Nodes[parent] > h.Nodes[child]
}

func getParent(index int) int {
	return (index - 1) / 2
}

func getLeftChild(index int) int {
	return (2 * index) + 1
}

func getRightChild(index int) int {
	return (2 * index) + 2
}

func (h *Heap) topdownAdjustCondition(parentIdx int) int {
	rcIndex := getRightChild(parentIdx)
	if rcIndex > h.lastNodeIdx {
		return HALT
	}
	rightChild := h.Nodes[rcIndex]
	lcIndex := getLeftChild(parentIdx)
	if lcIndex > h.lastNodeIdx {
		return HALT
	}
	leftChild := h.Nodes[lcIndex]
	if h.Type == MAX {
		if rightChild > leftChild && rightChild > h.Nodes[parentIdx] {
			return RIGHT_C
		}

		if leftChild > rightChild && leftChild > h.Nodes[parentIdx] {
			return LEFT_C
		}
	}

	if h.Type == MIN {
		if rightChild < leftChild && rightChild < h.Nodes[parentIdx] {
			return RIGHT_C
		}

		if leftChild < rightChild && leftChild < h.Nodes[parentIdx] {
			return LEFT_C
		}
	}
	return HALT
}
