package benchmark

import (
	"testing"
)

type (
	W1 struct {
		A [1]int64
	}
	W2 struct {
		A [2]int64
	}
	W3 struct {
		A [3]int64
	}
	W4 struct {
		A [4]int64
	}
	W5 struct {
		A [5]int64
	}
	W6 struct {
		A [6]int64
	}
	W7 struct {
		A [7]int64
	}
	W8 struct {
		A [8]int64
	}
	W9 struct {
		A [9]int64
	}
	W10 struct {
		A [10]int64
	}
	W11 struct {
		A [11]int64
	}
	W12 struct {
		A [12]int64
	}
	W13 struct {
		A [13]int64
	}
	W14 struct {
		A [14]int64
	}
	W15 struct {
		A [15]int64
	}
	W16 struct {
		A [16]int64
	}
	W17 struct {
		A [17]int64
	}
	W18 struct {
		A [18]int64
	}
	W19 struct {
		A [19]int64
	}
	W20 struct {
		A [20]int64
	}
)

func (w W1) ByValue() W1 {
	return w
}

func (w *W1) ByPointer() *W1 {
	return w
}

func (w W2) ByValue() W2 {
	return w
}

func (w *W2) ByPointer() *W2 {
	return w
}

func (w W3) ByValue() W3 {
	return w
}

func (w *W3) ByPointer() *W3 {
	return w
}

func (w W4) ByValue() W4 {
	return w
}

func (w *W4) ByPointer() *W4 {
	return w
}

func (w W5) ByValue() W5 {
	return w
}

func (w *W5) ByPointer() *W5 {
	return w
}

func (w W6) ByValue() W6 {
	return w
}

func (w *W6) ByPointer() *W6 {
	return w
}

func (w W7) ByValue() W7 {
	return w
}

func (w *W7) ByPointer() *W7 {
	return w
}

func (w W8) ByValue() W8 {
	return w
}

func (w *W8) ByPointer() *W8 {
	return w
}

func (w W9) ByValue() W9 {
	return w
}

func (w *W9) ByPointer() *W9 {
	return w
}

func (w W10) ByValue() W10 {
	return w
}

func (w *W10) ByPointer() *W10 {
	return w
}

func (w W11) ByValue() W11 {
	return w
}

func (w *W11) ByPointer() *W11 {
	return w
}

func (w W12) ByValue() W12 {
	return w
}

func (w *W12) ByPointer() *W12 {
	return w
}

func (w W13) ByValue() W13 {
	return w
}

func (w *W13) ByPointer() *W13 {
	return w
}

func (w W14) ByValue() W14 {
	return w
}

func (w *W14) ByPointer() *W14 {
	return w
}

func (w W15) ByValue() W15 {
	return w
}

func (w *W15) ByPointer() *W15 {
	return w
}

func (w W16) ByValue() W16 {
	return w
}

func (w *W16) ByPointer() *W16 {
	return w
}

func (w W17) ByValue() W17 {
	return w
}

func (w *W17) ByPointer() *W17 {
	return w
}

func (w W18) ByValue() W18 {
	return w
}

func (w *W18) ByPointer() *W18 {
	return w
}

func (w W19) ByValue() W19 {
	return w
}

func (w *W19) ByPointer() *W19 {
	return w
}

func (w W20) ByValue() W20 {
	return w
}

func (w *W20) ByPointer() *W20 {
	return w
}

func BenchmarkValueReceiverW1(b *testing.B) {
	w := W1{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW1(b *testing.B) {
	w := W1{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW2(b *testing.B) {
	w := W2{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW2(b *testing.B) {
	w := W2{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW3(b *testing.B) {
	w := W3{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW3(b *testing.B) {
	w := W3{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW4(b *testing.B) {
	w := W4{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW4(b *testing.B) {
	w := W4{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW5(b *testing.B) {
	w := W5{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW5(b *testing.B) {
	w := W5{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW6(b *testing.B) {
	w := W6{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW6(b *testing.B) {
	w := W6{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW7(b *testing.B) {
	w := W7{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW7(b *testing.B) {
	w := W7{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW8(b *testing.B) {
	w := W8{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW8(b *testing.B) {
	w := W8{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW9(b *testing.B) {
	w := W9{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW9(b *testing.B) {
	w := W9{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW10(b *testing.B) {
	w := W10{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW10(b *testing.B) {
	w := W10{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW11(b *testing.B) {
	w := W11{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW11(b *testing.B) {
	w := W11{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW12(b *testing.B) {
	w := W12{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW12(b *testing.B) {
	w := W12{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW13(b *testing.B) {
	w := W13{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW13(b *testing.B) {
	w := W13{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW14(b *testing.B) {
	w := W14{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW14(b *testing.B) {
	w := W14{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW15(b *testing.B) {
	w := W15{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW15(b *testing.B) {
	w := W15{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW16(b *testing.B) {
	w := W16{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW16(b *testing.B) {
	w := W16{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW17(b *testing.B) {
	w := W17{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW17(b *testing.B) {
	w := W17{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW18(b *testing.B) {
	w := W18{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW18(b *testing.B) {
	w := W18{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW19(b *testing.B) {
	w := W19{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW19(b *testing.B) {
	w := W19{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}

func BenchmarkValueReceiverW20(b *testing.B) {
	w := W20{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByValue()
	}
}

func BenchmarkPointerReceiverW20(b *testing.B) {
	w := W20{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = w.ByPointer()
	}
}
