package tuple

type T4[F0 any, F1 any, F2 any, F3 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
}

// NewT4 creates a new T4 tuple
func NewT4[F0 any, F1 any, F2 any, F3 any](f0 F0, f1 F1, f2 F2, f3 F3) T4[F0, F1, F2, F3] {
	return T4[F0, F1, F2, F3]{
		F0: f0,
		F1: f1,
		F2: f2,
		F3: f3,
	}
}

// First returns the first field of the tuple
func (t T4[F0, F1, F2, F3]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T4[F0, F1, F2, F3]) Last() F3 {
	return t.F3
}

// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T4[F0, F1, F2, F3]) CutLast() T3[F0, F1, F2] {
	return T3[F0, F1, F2]{
		F0: t.F0,
		F1: t.F1,
		F2: t.F2,
	}
}

// Get returns all tuple field values
func (t T4[F0, F1, F2, F3]) Get() (F0, F1, F2, F3) {
	return t.F0, t.F1, t.F2, t.F3
}
