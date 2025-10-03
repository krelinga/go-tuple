package tuple

type T3[F0 any, F1 any, F2 any] struct {
	F0 F0
	F1 F1
	F2 F2
}

// NewT3 creates a new T3 tuple
func NewT3[F0 any, F1 any, F2 any](f0 F0, f1 F1, f2 F2) T3[F0, F1, F2] {
	return T3[F0, F1, F2]{
		F0: f0,
		F1: f1,
		F2: f2,
	}
}

// First returns the first field of the tuple
func (t T3[F0, F1, F2]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T3[F0, F1, F2]) Last() F2 {
	return t.F2
}

// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T3[F0, F1, F2]) CutLast() T2[F0, F1] {
	return T2[F0, F1]{
		F0: t.F0,
		F1: t.F1,
	}
}

// Get returns all tuple field values
func (t T3[F0, F1, F2]) Get() (F0, F1, F2) {
	return t.F0, t.F1, t.F2
}
