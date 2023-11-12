package vis

type BlockRange struct {
	TopX, TopY, BottomX, BottomY int
}

func (b BlockRange) Overlaps(other BlockRange) bool {
	return b.TopX <= other.BottomX && b.BottomX >= other.TopX && b.TopY <= other.BottomY && b.BottomY >= other.TopY
}

func (b BlockRange) Contains(x, y int) bool {
	return b.TopX <= x && b.BottomX >= x && b.TopY <= y && b.BottomY >= y
}

func (b BlockRange) ContainsX(x int) bool {
	return b.TopX <= x && b.BottomX >= x
}

func (b BlockRange) ContainsY(y int) bool {
	return b.TopY <= y && b.BottomY >= y
}
