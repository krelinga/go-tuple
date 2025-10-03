package tuple

type T1[F0 any] struct {
	F0 F0
}

// NewT1 creates a new T1 tuple
func NewT1[F0 any](f0 F0) T1[F0] {
	return T1[F0]{
		F0: f0,
	}
}

// First returns the first field of the tuple
func (t T1[F0]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T1[F0]) Last() F0 {
	return t.F0
}

// Get returns all tuple field values
func (t T1[F0]) Get() (F0) {
	return t.F0
}
