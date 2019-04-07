package datastruct

type Heap struct {
	Nodes []int
}

func (h *Heap) Add(newValue int) {
	newNodes := append(h.Nodes, newValue)

	// first position is last.
	newPosition := len(newNodes)

	for newPosition > 1 {
		parentPosition := newPosition / 2

		// index = position - 1
		if newNodes[parentPosition-1] < newNodes[newPosition-1] {
			break
		}

		tmp := newNodes[parentPosition-1]
		newNodes[parentPosition-1] = newNodes[newPosition-1]
		newNodes[newPosition-1] = tmp

		newPosition = parentPosition
	}

	h.Nodes = newNodes
}

func (h *Heap) Put() int {
	length := len(h.Nodes)
	if length < 1 {
		// nodes is empty.
		return -1
	}

	value := h.Nodes[0]
	newNodes := append([]int{}, h.Nodes...)

	// put last node valude to first
	newNodes[0] = newNodes[length-1]
	newNodes = newNodes[:length-1]

	newPosition := 1
	for true {
		childPosition := getChildPosition(newNodes, newPosition)
		if childPosition < 0 {
			break
		}

		if newNodes[childPosition-1] < newNodes[newPosition-1] {
			tmp := newNodes[childPosition-1]
			newNodes[childPosition-1] = newNodes[newPosition-1]
			newNodes[newPosition-1] = tmp
			newPosition = childPosition
			continue
		}

		// both child value is larger than newPosition value
		break
	}

	h.Nodes = newNodes

	return value
}

func getChildPosition(targetNodes []int, targetPosition int) int {
	childPositionLeft := targetPosition * 2
	if childPositionLeft > len(targetNodes) {
		return -1
	}

	childPositionRight := childPositionLeft + 1
	if childPositionRight > len(targetNodes) {
		return childPositionLeft
	}

	if targetNodes[childPositionLeft-1] < targetNodes[childPositionRight-1] {
		return childPositionLeft
	}

	return childPositionRight
}
