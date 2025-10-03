package tuple

type T7[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any, F6 any] struct {
	F0 F0
	F1 F1
	F2 F2
	F3 F3
	F4 F4
	F5 F5
	F6 F6
}

// NewT7 creates a new T7 tuple
func NewT7[F0 any, F1 any, F2 any, F3 any, F4 any, F5 any, F6 any](f0 F0, f1 F1, f2 F2, f3 F3, f4 F4, f5 F5, f6 F6) T7[F0, F1, F2, F3, F4, F5, F6] {
	return T7[F0, F1, F2, F3, F4, F5, F6]{
		F0: f0,
		F1: f1,
		F2: f2,
		F3: f3,
		F4: f4,
		F5: f5,
		F6: f6,
	}
}

// First returns the first field of the tuple
func (t T7[F0, F1, F2, F3, F4, F5, F6]) First() F0 {
	return t.F0
}

// Last returns the last field of the tuple
func (t T7[F0, F1, F2, F3, F4, F5, F6]) Last() F6 {
	return t.F6
}

// CutLast returns a tuple with one fewer field (omitting the last field)
func (t T7[F0, F1, F2, F3, F4, F5, F6]) CutLast() T6[F0, F1, F2, F3, F4, F5] {
	return T6[F0, F1, F2, F3, F4, F5]{
		F0: t.F0,
		F1: t.F1,
		F2: t.F2,
		F3: t.F3,
		F4: t.F4,
		F5: t.F5,
	}
}

// Get returns all tuple field values
func (t T7[F0, F1, F2, F3, F4, F5, F6]) Get() (F0, F1, F2, F3, F4, F5, F6) {
	return t.F0, t.F1, t.F2, t.F3, t.F4, t.F5, t.F6
}
