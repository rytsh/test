package variadic

import "testing"

var testSize = 1000000

func BenchmarkVariadic(b *testing.B) {
	datas := RandomDataGen(testSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CallVariadic(datas...)
	}
}

func BenchmarkSlice(b *testing.B) {
	datas := RandomDataGen(testSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CallSlice(datas)
	}
}

func BenchmarkOne(b *testing.B) {
	datas := RandomDataGen(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CallVariadic(datas[0])
	}
}
