package multichoose

type MultiChoose struct {
	selected uint64
	length   int
	limit    int
}

func New(length int) *MultiChoose {
	return &MultiChoose{
		selected: 0,
		length:   length,
		limit:    -1,
	}
}

func (mc MultiChoose) Count() int {
	count := 0
	for i := 0; i <= mc.length; i++ {
		if mc.selected>>i&1 == 1 {
			count += 1
		}
	}
	return count
}

func (mc *MultiChoose) Select(index int) {
	if mc.limit >= 0 && mc.Count() >= mc.limit {
		return
	}

	curr := uint64(1) << index
	mc.selected = mc.selected | curr
}

func (mc *MultiChoose) Deselect(index int) {
	curr := uint64(1) << index
	mc.selected = mc.selected & (^curr)
}

func (mc MultiChoose) IsSelected(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (mc.selected & curr) != 0 {
		return true
	} else {
		return false
	}
}
