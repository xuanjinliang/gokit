package gotool

import (
	"reflect"
	"testing"
	"time"
)

func TestStringPtr(t *testing.T) {
	a := StringPtr("test")
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrString(t *testing.T) {
	str := "test"
	a := PtrString(&str)
	t.Log(a == str)
}

func TestStringSlicePtr(t *testing.T) {
	a := []string{"aa", "bb"}
	list := StringSlicePtr(a)
	for _, str := range list {
		if reflect.ValueOf(str).IsValid() {
			t.Logf("%v", *str)
		}
	}
}

func TestPtrStringSlice(t *testing.T) {
	a := []*string{StringPtr("aa"), StringPtr("bb")}
	list := PtrStringSlice(a)
	for _, str := range list {
		if reflect.ValueOf(str).IsValid() {
			t.Logf("%v", str)
		}
	}
}

func TestStringMapPtr(t *testing.T) {
	a := map[string]string{"k": "aa", "j": "bb"}
	mapList := StringMapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrStringMap(t *testing.T) {
	a := map[string]*string{"k": StringPtr("aa"), "j": StringPtr("bb")}
	mapList := PtrStringMap(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestBoolPtr(t *testing.T) {
	a := BoolPtr(true)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrBool(t *testing.T) {
	a := true
	b := PtrBool(&a)
	t.Log(a == b)
}

func TestBoolSlicePtr(t *testing.T) {
	a := []bool{true, false}
	list := BoolSlicePtr(a)
	for _, b := range list {
		if reflect.ValueOf(b).IsValid() {
			t.Logf("%v", *b)
		}
	}
}

func TestPtrBoolSlice(t *testing.T) {
	a := []*bool{BoolPtr(true), BoolPtr(false)}
	list := PtrBoolSlice(a)
	for _, b := range list {
		t.Logf("%v", b)
	}
}

func TestBoolMapPtr(t *testing.T) {
	a := map[string]bool{"aa": true, "bb": false}
	mapList := BoolMapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrBoolMap(t *testing.T) {
	a := map[string]*bool{"aa": BoolPtr(true), "bb": BoolPtr(false)}
	mapList := PtrBoolMap(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestIntPtr(t *testing.T) {
	a := IntPtr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrInt(t *testing.T) {
	a := 1
	b := PtrInt(&a)
	t.Logf("%v", b)
}

func TestIntSlicePtr(t *testing.T) {
	a := []int{1, 2, 3}
	list := IntSlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrIntSlice(t *testing.T) {
	a := []*int{IntPtr(1), IntPtr(2), IntPtr(3)}
	list := PtrIntSlice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestIntMapPtr(t *testing.T) {
	a := map[string]int{"k": 1, "j": 2}
	mapList := IntMapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrIntMap(t *testing.T) {
	a := map[string]*int{"k": IntPtr(1), "j": IntPtr(2)}
	mapList := PtrIntMap(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestUintPtr(t *testing.T) {
	a := UintPtr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrUint(t *testing.T) {
	a := PtrUint(UintPtr(1))
	t.Logf("%v", a)
}

func TestUintSlicePtr(t *testing.T) {
	a := []uint{1, 2, 3}
	list := UintSlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrUintSlice(t *testing.T) {
	a := []*uint{UintPtr(1), UintPtr(2), UintPtr(3)}
	list := PtrUintSlice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestUintMapPtr(t *testing.T) {
	a := map[string]uint{"k": 1, "j": 2}
	mapList := UintMapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrUintMap(t *testing.T) {
	a := map[string]*uint{"k": UintPtr(1), "j": UintPtr(2)}
	mapList := PtrUintMap(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestInt8Ptr(t *testing.T) {
	a := Int8Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrInt8(t *testing.T) {
	var a int8 = 1
	b := PtrInt8(&a)
	t.Logf("%v", a == b)
}

func TestInt8SlicePtr(t *testing.T) {
	a := []int8{1, 2, 3}
	list := Int8SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrInt8Slice(t *testing.T) {
	a := []*int8{Int8Ptr(1), Int8Ptr(2), Int8Ptr(3)}
	list := PtrInt8Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestInt8MapPtr(t *testing.T) {
	a := map[string]int8{"k": 1, "j": 2}
	mapList := Int8MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrInt8Map(t *testing.T) {
	a := map[string]*int8{"k": Int8Ptr(1), "j": Int8Ptr(2)}
	mapList := PtrInt8Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestInt16Ptr(t *testing.T) {
	a := Int16Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrInt16(t *testing.T) {
	var a int16 = 1
	b := PtrInt16(&a)
	t.Logf("%v", a == b)
}

func TestInt16SlicePtr(t *testing.T) {
	a := []int16{1, 2, 3}
	list := Int16SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrInt16Slice(t *testing.T) {
	a := []*int16{Int16Ptr(1), Int16Ptr(2), Int16Ptr(3)}
	list := PtrInt16Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestInt16MapPtr(t *testing.T) {
	a := map[string]int16{"k": 1, "j": 2}
	mapList := Int16MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrInt16Map(t *testing.T) {
	a := map[string]*int16{"k": Int16Ptr(1), "j": Int16Ptr(2)}
	mapList := PtrInt16Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestInt32Ptr(t *testing.T) {
	a := Int32Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrInt32(t *testing.T) {
	var a int32 = 1
	b := PtrInt32(&a)
	t.Logf("%v", a == b)
}

func TestInt32SlicePtr(t *testing.T) {
	a := []int32{1, 2, 3}
	list := Int32SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrInt32Slice(t *testing.T) {
	a := []*int32{Int32Ptr(1), Int32Ptr(2), Int32Ptr(3)}
	list := PtrInt32Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestInt32MapPtr(t *testing.T) {
	a := map[string]int32{"k": 1, "j": 2}
	mapList := Int32MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrInt32Map(t *testing.T) {
	a := map[string]*int32{"k": Int32Ptr(1), "j": Int32Ptr(2)}
	mapList := PtrInt32Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestInt64Ptr(t *testing.T) {
	a := Int64Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrInt64(t *testing.T) {
	var a int64 = 1
	b := PtrInt64(&a)
	t.Logf("%v", a == b)
}

func TestInt64SlicePtr(t *testing.T) {
	a := []int64{1, 2, 3}
	list := Int64SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrInt64Slice(t *testing.T) {
	a := []*int64{Int64Ptr(1), Int64Ptr(2), Int64Ptr(3)}
	list := PtrInt64Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestInt64MapPtr(t *testing.T) {
	a := map[string]int64{"k": 1, "j": 2}
	mapList := Int64MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrInt64Map(t *testing.T) {
	a := map[string]*int64{"k": Int64Ptr(1), "j": Int64Ptr(2)}
	mapList := PtrInt64Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestUint8Ptr(t *testing.T) {
	a := Uint8Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrUint8(t *testing.T) {
	a := PtrUint8(Uint8Ptr(1))
	t.Logf("%v", a)
}

func TestUint8SlicePtr(t *testing.T) {
	a := []uint8{1, 2, 3}
	list := Uint8SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrUint8Slice(t *testing.T) {
	a := []*uint8{Uint8Ptr(1), Uint8Ptr(2), Uint8Ptr(3)}
	list := PtrUint8Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestUint8MapPtr(t *testing.T) {
	a := map[string]uint8{"k": 1, "j": 2}
	mapList := Uint8MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrUint8Map(t *testing.T) {
	a := map[string]*uint8{"k": Uint8Ptr(1), "j": Uint8Ptr(2)}
	mapList := PtrUint8Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestUint16Ptr(t *testing.T) {
	a := Uint16Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrUint16(t *testing.T) {
	a := PtrUint16(Uint16Ptr(1))
	t.Logf("%v", a)
}

func TestUint16SlicePtr(t *testing.T) {
	a := []uint16{1, 2, 3}
	list := Uint16SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrUint16Slice(t *testing.T) {
	a := []*uint16{Uint16Ptr(1), Uint16Ptr(2), Uint16Ptr(3)}
	list := PtrUint16Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestUint16MapPtr(t *testing.T) {
	a := map[string]uint16{"k": 1, "j": 2}
	mapList := Uint16MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrUint16Map(t *testing.T) {
	a := map[string]*uint16{"k": Uint16Ptr(1), "j": Uint16Ptr(2)}
	mapList := PtrUint16Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestUint32Ptr(t *testing.T) {
	a := Uint32Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrUint32(t *testing.T) {
	a := PtrUint32(Uint32Ptr(1))
	t.Logf("%v", a)
}

func TestUint32SlicePtr(t *testing.T) {
	a := []uint32{1, 2, 3}
	list := Uint32SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrUint32Slice(t *testing.T) {
	a := []*uint32{Uint32Ptr(1), Uint32Ptr(2), Uint32Ptr(3)}
	list := PtrUint32Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestUint32MapPtr(t *testing.T) {
	a := map[string]uint32{"k": 1, "j": 2}
	mapList := Uint32MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrUint32Map(t *testing.T) {
	a := map[string]*uint32{"k": Uint32Ptr(1), "j": Uint32Ptr(2)}
	mapList := PtrUint32Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestUint64Ptr(t *testing.T) {
	a := Uint64Ptr(1)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrUint64(t *testing.T) {
	a := PtrUint64(Uint64Ptr(1))
	t.Logf("%v", a)
}

func TestUint64SlicePtr(t *testing.T) {
	a := []uint64{1, 2, 3}
	list := Uint64SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrUint64Slice(t *testing.T) {
	a := []*uint64{Uint64Ptr(1), Uint64Ptr(2), Uint64Ptr(3)}
	list := PtrUint64Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestUint64MapPtr(t *testing.T) {
	a := map[string]uint64{"k": 1, "j": 2}
	mapList := Uint64MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrUint64Map(t *testing.T) {
	a := map[string]*uint64{"k": Uint64Ptr(1), "j": Uint64Ptr(2)}
	mapList := PtrUint64Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestFloat32Ptr(t *testing.T) {
	a := Float32Ptr(1.0)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrFloat32(t *testing.T) {
	a := PtrFloat32(Float32Ptr(1.0))
	t.Logf("%v", a)
}

func TestFloat32SlicePtr(t *testing.T) {
	a := []float32{1.0, 2.0, 3.0}
	list := Float32SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrFloat32Slice(t *testing.T) {
	a := []*float32{Float32Ptr(1.0), Float32Ptr(2.0), Float32Ptr(3.0)}
	list := PtrFloat32Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestFloat32MapPtr(t *testing.T) {
	a := map[string]float32{"k": 1.0, "j": 2.0}
	mapList := Float32MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrFloat32Map(t *testing.T) {
	a := map[string]*float32{"k": Float32Ptr(1.0), "j": Float32Ptr(2.0)}
	mapList := PtrFloat32Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestFloat64Ptr(t *testing.T) {
	a := Float64Ptr(1.0)
	if reflect.ValueOf(a).IsValid() {
		t.Logf("%v", *a)
	}
}

func TestPtrFloat64(t *testing.T) {
	a := PtrFloat64(Float64Ptr(1.0))
	t.Logf("%v", a)
}

func TestFloat64SlicePtr(t *testing.T) {
	a := []float64{1.0, 2.0, 3.0}
	list := Float64SlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrFloat64Slice(t *testing.T) {
	a := []*float64{Float64Ptr(1.0), Float64Ptr(2.0), Float64Ptr(3.0)}
	list := PtrFloat64Slice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestFloat64MapPtr(t *testing.T) {
	a := map[string]float64{"k": 1.0, "j": 2.0}
	mapList := Float64MapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrFloat64Map(t *testing.T) {
	a := map[string]*float64{"k": Float64Ptr(1.0), "j": Float64Ptr(2.0)}
	mapList := PtrFloat64Map(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}

func TestTimePtr(t *testing.T) {
	now := TimePtr(time.Now())
	if reflect.ValueOf(now).IsValid() {
		t.Logf("%v", *now)
	}
}

func TestPtrTime(t *testing.T) {
	a := PtrTime(TimePtr(time.Now()))
	t.Logf("%v", a)
}

func TestTimeSlicePtr(t *testing.T) {
	now := time.Now()
	a := []time.Time{now.Add(1 * time.Minute), now.Add(2 * time.Minute), now.Add(3 * time.Minute)}
	list := TimeSlicePtr(a)
	for _, v := range list {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("%v", *v)
		}
	}
}

func TestPtrTimeSlice(t *testing.T) {
	now := time.Now()
	a := []*time.Time{TimePtr(now.Add(1 * time.Minute)), TimePtr(now.Add(2 * time.Minute)), TimePtr(now.Add(3 * time.Minute))}
	list := PtrTimeSlice(a)
	for _, v := range list {
		t.Logf("%v", v)
	}
}

func TestTimeMapPtr(t *testing.T) {
	now := time.Now()
	a := map[string]time.Time{"k": now.Add(1 * time.Minute), "j": now.Add(2 * time.Minute)}
	mapList := TimeMapPtr(a)
	for k, v := range mapList {
		if reflect.ValueOf(v).IsValid() {
			t.Logf("key --> %v, value --> %v", k, *v)
		}
	}
}

func TestPtrTimeMap(t *testing.T) {
	now := time.Now()
	a := map[string]*time.Time{"k": TimePtr(now.Add(1 * time.Minute)), "j": TimePtr(now.Add(2 * time.Minute))}
	mapList := PtrTimeMap(a)
	for k, v := range mapList {
		t.Logf("key --> %v, value --> %v", k, v)
	}
}