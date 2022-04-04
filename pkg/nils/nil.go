package nils

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"math/rand"
	"time"
)

func NewString(value string) *string {
	return &value
}

func StringIsNil(b *string) string {
	if b == nil {
		return ""
	} else {
		return *b
	}
}
func StringNilCheck(b *string) bool {
	if b == nil {
		return true
	} else {
		return false
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

func NewUint16(value uint16) *uint16 {
	return &value
}

func NewUint32(value uint32) *uint32 {
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

//RandomString will make a random string that can be used for naming as an example a unique name
func RandomString() string {
	u, _ := uuid.MakeUUID()
	return fmt.Sprintf("n_%s", truncateString(u, 8))

}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func truncateString(str string, num int) string {
	ret := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		ret = str[0:num] + ""
	}
	return ret
}
