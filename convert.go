package gotool

import "time"

// StringPtr returns a pointer to the string value passed in.
func StringPtr(v string) *string {
	return &v
}

// PtrString returns the value of the string pointer passed in or
// "" if the pointer is nil.
func PtrString(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlicePtr converts a slice of string values into a slice of
// string pointers
func StringSlicePtr(src []string) []*string {
	dst := make([]*string, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrStringSlice converts a slice of string pointers into a slice of
// string values
func PtrStringSlice(src []*string) []string {
	dst := make([]string, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// StringMapPtr converts a string map of string values into a string
// map of string pointers
func StringMapPtr(src map[string]string) map[string]*string {
	dst := make(map[string]*string)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrStringMap converts a string map of string pointers into a string
// map of string values
func PtrStringMap(src map[string]*string) map[string]string {
	dst := make(map[string]string)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// BoolPtr returns a pointer to the bool value passed in.
func BoolPtr(v bool) *bool {
	return &v
}

// PtrBool returns the value of the bool pointer passed in or
// false if the pointer is nil.
func PtrBool(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlicePtr converts a slice of bool values into a slice of
// bool pointers
func BoolSlicePtr(src []bool) []*bool {
	dst := make([]*bool, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrBoolValueSlice converts a slice of bool pointers into a slice of
// bool values
func PtrBoolSlice(src []*bool) []bool {
	dst := make([]bool, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// BoolMapPtr converts a string map of bool values into a string
// map of bool pointers
func BoolMapPtr(src map[string]bool) map[string]*bool {
	dst := make(map[string]*bool)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrBoolMap converts a string map of bool pointers into a string
// map of bool values
func PtrBoolMap(src map[string]*bool) map[string]bool {
	dst := make(map[string]bool)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// IntPtr returns a pointer to the int value passed in.
func IntPtr(v int) *int {
	return &v
}

// PtrInt returns the value of the int pointer passed in or
// 0 if the pointer is nil.
func PtrInt(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntSlicePtr converts a slice of int values into a slice of
// int pointers
func IntSlicePtr(src []int) []*int {
	dst := make([]*int, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrIntSlice converts a slice of int pointers into a slice of
// int values
func PtrIntSlice(src []*int) []int {
	dst := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// IntMapPtr converts a string map of int values into a string
// map of int pointers
func IntMapPtr(src map[string]int) map[string]*int {
	dst := make(map[string]*int)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrIntMap converts a string map of int pointers into a string
// map of int values
func PtrIntMap(src map[string]*int) map[string]int {
	dst := make(map[string]int)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// UintPtr returns a pointer to the uint value passed in.
func UintPtr(v uint) *uint {
	return &v
}

// PtrUint returns the value of the uint pointer passed in or
// 0 if the pointer is nil.
func PtrUint(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintSlicePtr converts a slice of uint values uinto a slice of
// uint pointers
func UintSlicePtr(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrUintSlice converts a slice of uint pointers uinto a slice of
// uint values
func PtrUintSlice(src []*uint) []uint {
	dst := make([]uint, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// UintMapPtr converts a string map of uint values uinto a string
// map of uint pointers
func UintMapPtr(src map[string]uint) map[string]*uint {
	dst := make(map[string]*uint)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrUintMap converts a string map of uint pointers uinto a string
// map of uint values
func PtrUintMap(src map[string]*uint) map[string]uint {
	dst := make(map[string]uint)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int8Ptr returns a pointer to the int8 value passed in.
func Int8Ptr(v int8) *int8 {
	return &v
}

// PtrInt8 returns the value of the int8 pointer passed in or
// 0 if the pointer is nil.
func PtrInt8(v *int8) int8 {
	if v != nil {
		return *v
	}
	return 0
}

// Int8SlicePtr converts a slice of int8 values into a slice of
// int8 pointers
func Int8SlicePtr(src []int8) []*int8 {
	dst := make([]*int8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrInt8Slice converts a slice of int8 pointers into a slice of
// int8 values
func PtrInt8Slice(src []*int8) []int8 {
	dst := make([]int8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int8MapPtr converts a string map of int8 values into a string
// map of int8 pointers
func Int8MapPtr(src map[string]int8) map[string]*int8 {
	dst := make(map[string]*int8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrInt8Map converts a string map of int8 pointers into a string
// map of int8 values
func PtrInt8Map(src map[string]*int8) map[string]int8 {
	dst := make(map[string]int8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int16 returns a pointer to the int16 value passed in.
func Int16Ptr(v int16) *int16 {
	return &v
}

// Int16Value returns the value of the int16 pointer passed in or
// 0 if the pointer is nil.
func PtrInt16(v *int16) int16 {
	if v != nil {
		return *v
	}
	return 0
}

// Int16SlicePtr converts a slice of int16 values into a slice of
// int16 pointers
func Int16SlicePtr(src []int16) []*int16 {
	dst := make([]*int16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrInt16Slice converts a slice of int16 pointers into a slice of
// int16 values
func PtrInt16Slice(src []*int16) []int16 {
	dst := make([]int16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int16MapPtr converts a string map of int16 values into a string
// map of int16 pointers
func Int16MapPtr(src map[string]int16) map[string]*int16 {
	dst := make(map[string]*int16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrInt16Map converts a string map of int16 pointers into a string
// map of int16 values
func PtrInt16Map(src map[string]*int16) map[string]int16 {
	dst := make(map[string]int16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int32Ptr returns a pointer to the int32 value passed in.
func Int32Ptr(v int32) *int32 {
	return &v
}

// PtrInt32 returns the value of the int32 pointer passed in or
// 0 if the pointer is nil.
func PtrInt32(v *int32) int32 {
	if v != nil {
		return *v
	}
	return 0
}

// Int32SlicePtr converts a slice of int32 values into a slice of
// int32 pointers
func Int32SlicePtr(src []int32) []*int32 {
	dst := make([]*int32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrInt32Slice converts a slice of int32 pointers into a slice of
// int32 values
func PtrInt32Slice(src []*int32) []int32 {
	dst := make([]int32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int32MapPtr converts a string map of int32 values into a string
// map of int32 pointers
func Int32MapPtr(src map[string]int32) map[string]*int32 {
	dst := make(map[string]*int32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrInt32Map converts a string map of int32 pointers into a string
// map of int32 values
func PtrInt32Map(src map[string]*int32) map[string]int32 {
	dst := make(map[string]int32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Int64Ptr returns a pointer to the int64 value passed in.
func Int64Ptr(v int64) *int64 {
	return &v
}

// PtrInt64 returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func PtrInt64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64SlicePtr converts a slice of int64 values into a slice of
// int64 pointers
func Int64SlicePtr(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrInt64Slice converts a slice of int64 pointers into a slice of
// int64 values
func PtrInt64Slice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Int64MapPtr converts a string map of int64 values into a string
// map of int64 pointers
func Int64MapPtr(src map[string]int64) map[string]*int64 {
	dst := make(map[string]*int64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrInt64Map converts a string map of int64 pointers into a string
// map of int64 values
func PtrInt64Map(src map[string]*int64) map[string]int64 {
	dst := make(map[string]int64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint8Ptr returns a pointer to the uint8 value passed in.
func Uint8Ptr(v uint8) *uint8 {
	return &v
}

// PtrUint8 returns the value of the uint8 pointer passed in or
// 0 if the pointer is nil.
func PtrUint8(v *uint8) uint8 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint8SlicePtr converts a slice of uint8 values into a slice of
// uint8 pointers
func Uint8SlicePtr(src []uint8) []*uint8 {
	dst := make([]*uint8, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrUint8Slice converts a slice of uint8 pointers into a slice of
// uint8 values
func PtrUint8Slice(src []*uint8) []uint8 {
	dst := make([]uint8, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint8MapPtr converts a string map of uint8 values into a string
// map of uint8 pointers
func Uint8MapPtr(src map[string]uint8) map[string]*uint8 {
	dst := make(map[string]*uint8)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrUint8Map converts a string map of uint8 pointers into a string
// map of uint8 values
func PtrUint8Map(src map[string]*uint8) map[string]uint8 {
	dst := make(map[string]uint8)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint16Ptr returns a pointer to the uint16 value passed in.
func Uint16Ptr(v uint16) *uint16 {
	return &v
}

// PtrUint16 returns the value of the uint16 pointer passed in or
// 0 if the pointer is nil.
func PtrUint16(v *uint16) uint16 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint16SlicePtr converts a slice of uint16 values into a slice of
// uint16 pointers
func Uint16SlicePtr(src []uint16) []*uint16 {
	dst := make([]*uint16, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrUint16Slice converts a slice of uint16 pointers into a slice of
// uint16 values
func PtrUint16Slice(src []*uint16) []uint16 {
	dst := make([]uint16, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint16MapPtr converts a string map of uint16 values into a string
// map of uint16 pointers
func Uint16MapPtr(src map[string]uint16) map[string]*uint16 {
	dst := make(map[string]*uint16)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrUint16Map converts a string map of uint16 pointers into a string
// map of uint16 values
func PtrUint16Map(src map[string]*uint16) map[string]uint16 {
	dst := make(map[string]uint16)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint32Ptr returns a pointer to the uint32 value passed in.
func Uint32Ptr(v uint32) *uint32 {
	return &v
}

// PtrUint32 returns the value of the uint32 pointer passed in or
// 0 if the pointer is nil.
func PtrUint32(v *uint32) uint32 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint32Slice converts a slice of uint32 values into a slice of
// uint32 pointers
func Uint32SlicePtr(src []uint32) []*uint32 {
	dst := make([]*uint32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrUint32Slice converts a slice of uint32 pointers into a slice of
// uint32 values
func PtrUint32Slice(src []*uint32) []uint32 {
	dst := make([]uint32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint32MapPtr converts a string map of uint32 values into a string
// map of uint32 pointers
func Uint32MapPtr(src map[string]uint32) map[string]*uint32 {
	dst := make(map[string]*uint32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrUint32Map converts a string map of uint32 pointers into a string
// map of uint32 values
func PtrUint32Map(src map[string]*uint32) map[string]uint32 {
	dst := make(map[string]uint32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Uint64Ptr returns a pointer to the uint64 value passed in.
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// PtrUint64 returns the value of the uint64 pointer passed in or
// 0 if the pointer is nil.
func PtrUint64(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64SlicePtr converts a slice of uint64 values into a slice of
// uint64 pointers
func Uint64SlicePtr(src []uint64) []*uint64 {
	dst := make([]*uint64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrUint64Slice converts a slice of uint64 pointers into a slice of
// uint64 values
func PtrUint64Slice(src []*uint64) []uint64 {
	dst := make([]uint64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Uint64MapPtr converts a string map of uint64 values into a string
// map of uint64 pointers
func Uint64MapPtr(src map[string]uint64) map[string]*uint64 {
	dst := make(map[string]*uint64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrUint64Map converts a string map of uint64 pointers into a string
// map of uint64 values
func PtrUint64Map(src map[string]*uint64) map[string]uint64 {
	dst := make(map[string]uint64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Float32Ptr returns a pointer to the float32 value passed in.
func Float32Ptr(v float32) *float32 {
	return &v
}

// PtrFloat32 returns the value of the float32 pointer passed in or
// 0 if the pointer is nil.
func PtrFloat32(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0
}

// Float32SlicePtr converts a slice of float32 values into a slice of
// float32 pointers
func Float32SlicePtr(src []float32) []*float32 {
	dst := make([]*float32, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrFloat32Slice converts a slice of float32 pointers into a slice of
// float32 values
func PtrFloat32Slice(src []*float32) []float32 {
	dst := make([]float32, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float32MapPtr converts a string map of float32 values into a string
// map of float32 pointers
func Float32MapPtr(src map[string]float32) map[string]*float32 {
	dst := make(map[string]*float32)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrFloat32Map converts a string map of float32 pointers into a string
// map of float32 values
func PtrFloat32Map(src map[string]*float32) map[string]float32 {
	dst := make(map[string]float32)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// Float64Ptr returns a pointer to the float64 value passed in.
func Float64Ptr(v float64) *float64 {
	return &v
}

// PtrFloat64 returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func PtrFloat64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64SlicePtr converts a slice of float64 values into a slice of
// float64 pointers
func Float64SlicePtr(src []float64) []*float64 {
	dst := make([]*float64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrFloat64Slice converts a slice of float64 pointers into a slice of
// float64 values
func PtrFloat64Slice(src []*float64) []float64 {
	dst := make([]float64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// Float64MapPtr converts a string map of float64 values into a string
// map of float64 pointers
func Float64MapPtr(src map[string]float64) map[string]*float64 {
	dst := make(map[string]*float64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrFloat64Map converts a string map of float64 pointers into a string
// map of float64 values
func PtrFloat64Map(src map[string]*float64) map[string]float64 {
	dst := make(map[string]float64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}

// TimePtr returns a pointer to the time.Time value passed in.
func TimePtr(v time.Time) *time.Time {
	return &v
}

// PtrTime returns the value of the time.Time pointer passed in or
// time.Time{} if the pointer is nil.
func PtrTime(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// TimeSlicePtr converts a slice of time.Time values into a slice of
// time.Time pointers
func TimeSlicePtr(src []time.Time) []*time.Time {
	dst := make([]*time.Time, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}
	return dst
}

// PtrTimeSlice converts a slice of time.Time pointers into a slice of
// time.Time values
func PtrTimeSlice(src []*time.Time) []time.Time {
	dst := make([]time.Time, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}
	return dst
}

// TimeMapPtr converts a string map of time.Time values into a string
// map of time.Time pointers
func TimeMapPtr(src map[string]time.Time) map[string]*time.Time {
	dst := make(map[string]*time.Time)
	for k, val := range src {
		v := val
		dst[k] = &v
	}
	return dst
}

// PtrTimeMap converts a string map of time.Time pointers into a string
// map of time.Time values
func PtrTimeMap(src map[string]*time.Time) map[string]time.Time {
	dst := make(map[string]time.Time)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}
	return dst
}