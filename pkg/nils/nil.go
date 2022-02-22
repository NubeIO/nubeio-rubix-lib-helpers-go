package nils

func StringIsNil(b *string) string {
	if b == nil {
		return ""
	} else {
		return *b
	}
}

func Float64IsNil(b *float64) float64 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func NewInt(value int) *int {
	return &value
}

func NewFloat64(value float64) *float64 {
	return &value
}

func IntIsNil(b *int) int {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func IntNilCheck(b *int) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

func Float32IsNil(b *float32) float32 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func UnitIsNil(b *uint) uint {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func Unit16IsNil(b *uint16) uint16 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func Unit32IsNil(b *uint32) uint32 {
	if b == nil {
		return 0
	} else {
		return *b
	}
}

func Unit32NilCheck(b *uint32) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}

func FloatIsNilCheck(b *float64) bool {
	if b == nil {
		return true
	} else {
		return false
	}
}