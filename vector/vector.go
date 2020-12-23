package vector

type (
	VectorInterface interface {
		Size() int
		Empty() bool
	}

	Vector struct {
		list []interface{}
	}
)
