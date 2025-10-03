package tuple

type T5[F0 any, F1 any, F2 any, F3 any, F4 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
	F4 F4
}

// NewT5 creates a new T5 tuple
func NewT5[F0 any, F1 any, F2 any, F3 any, F4 any](f0 F0, f1 F1, f2 F2, f3 F3, f4 F4) T5[F0, F1, F2, F3, F4] {
	return T5[F0, F1, F2, F3, F4]{
		F0: f0,
		F1: f1,
		F2: f2,
		F3: f3,
		F4: f4,
	}
}

// First returns the first field of the tuple
func (t T5[F0, F1, F2, F3, F4]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T5[F0, F1, F2, F3, F4]) Last() F4 {
	return t.F4
}

// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T5[F0, F1, F2, F3, F4]) CutLast() T4[F0, F1, F2, F3] {
	return T4[F0, F1, F2, F3]{
		F0: t.F0,
		F1: t.F1,
		F2: t.F2,
		F3: t.F3,
	}
}

// Get returns all tuple field values
func (t T5[F0, F1, F2, F3, F4]) Get() (F0, F1, F2, F3, F4) {
	return t.F0, t.F1, t.F2, t.F3, t.F4
}
