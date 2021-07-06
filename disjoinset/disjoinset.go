package disjoinset

type DisjoinSet interface {
	FindRoot(x int) int
	Merge(x, y int) bool
}

type (
	DjSet struct {
		parent []int
		rank   []int
	}
)

func (d DjSet) FindRoot(x int) int {
	for d.parent[x] != -1 {
		x = d.parent[x]
	}
	return x
}

func (d DjSet) Merge(x, y int) bool {
	rx := d.FindRoot(x)
	ry := d.FindRoot(y)

	if rx == ry {
		return false
	} else {
		if d.rank[rx] > d.rank[ry] {
			d.parent[ry] = rx
		} else if d.rank[rx] < d.rank[ry] {
			d.parent[rx] = ry
		} else {
			d.rank[ry] += 1
			d.parent[rx] = ry
		}

		return true
	}
}

func NewDjSet(l int) DjSet {
	parent := make([]int, l)
	for i := 0; i < l; i++ {
		parent[i] = -1
	}
	return DjSet{
		parent: parent,
		rank:   make([]int, l),
	}
}
