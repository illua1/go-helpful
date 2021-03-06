package helpful_sort

import (
	"github.com/illua1/go-helpful"
)

// MaxId will find index of max element in input slise.
func MaxId[T helpful.Values](in ...T) int {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if in[i] > in[id] {
				id = i
			}
		}
		return id
	} else {
		if len(in) == 2 {
			return MaxIdF[T](in[0], in[1])
		} else {
			if len(in) == 1 {
				return 0
			} else {
				return -1
			}
		}
	}
	return -1
}

// MaxIdF will find index of max element at input dual values.
func MaxIdF[T helpful.Values](a, b T) int {
	if a > b {
		return 0
	}
	return 1
}

// Max will find max element in input slise and return him, if slise size is zero, Max will return T{}.
func Max[T helpful.Values](in ...T) T {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if in[i] > in[id] {
				id = i
			}
		}
		return in[id]
	} else {
		if len(in) == 2 {
			return MaxF[T](in[0], in[1])
		} else {
			if len(in) == 1 {
				return in[0]
			} else {
				var NulValue T
				return NulValue
			}
		}
	}
	var NulValue T
	return NulValue
}

// MaxF will find max element at input dual values.
func MaxF[T helpful.Values](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// MinId will find index of min element in input slise.
func MinId[T helpful.Values](in ...T) int {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if in[i] < in[id] {
				id = i
			}
		}
		return id
	} else {
		if len(in) == 2 {
			return MinIdF[T](in[0], in[1])
		} else {
			if len(in) == 1 {
				return 0
			} else {
				return -1
			}
		}
	}
	return -1
}

// MinIdF will find index of min element at input dual values.
func MinIdF[T helpful.Values](a, b T) int {
	if a < b {
		return 0
	}
	return 1
}

// Min will find min element in input slise and return him, if slise size is zero, Min will return T{}.
func Min[T helpful.Values](in ...T) T {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if in[i] < in[id] {
				id = i
			}
		}
		return in[id]
	} else {
		if len(in) == 2 {
			return MinF[T](in[0], in[1])
		} else {
			if len(in) == 1 {
				return in[0]
			} else {
				var NulValue T
				return NulValue
			}
		}
	}
	var NulValue T
	return NulValue
}

// MinF will find min element at input dual values.
func MinF[T helpful.Values](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func PickAny[T any](condition func(a, b T) bool, in ...T) T {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if condition(in[i], in[id]) {
				id = i
			}
		}
		return in[id]
	} else {
		if len(in) == 2 {
			if condition(in[0], in[1]) {
				return in[0]
			} else {
				return in[1]
			}
		} else {
			if len(in) == 1 {
				return in[0]
			} else {
				var NulValue T
				return NulValue
			}
		}
	}
	var NulValue T
	return NulValue
}

func PickAnyId[T any](condition func(a, b T) bool, in ...T) int {
	if len(in) > 2 {
		var id = 0
		for i := 1; i < len(in); i++ {
			if condition(in[i], in[id]) {
				id = i
			}
		}
		return id
	} else {
		if len(in) == 2 {
			if condition(in[0], in[1]) {
				return 0
			} else {
				return 1
			}
		} else {
			if len(in) == 1 {
				return 0
			} else {
				return -1
			}
		}
	}
	return -1
}
