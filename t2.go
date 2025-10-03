package tuple

type T2[F0 any, F1 any] struct {
	F0 F0
	F1 F1
}

// NewT2 creates a new T2 tuple
func NewT2[F0 any, F1 any](f0 F0, f1 F1) T2[F0, F1] {
	return T2[F0, F1]{
		F0: f0,
		F1: f1,
	}
}

// First returns the first field of the tuple
func (t T2[F0, F1]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T2[F0, F1]) Last() F1 {
	return t.F1
}

// CutLast returns the first field (omitting the last field)
func (t T2[F0, F1]) CutLast() F0 {
	return t.F0
}

// Get returns all tuple field values
func (t T2[F0, F1]) Get() (F0, F1) {
	return t.F0, t.F1
}
