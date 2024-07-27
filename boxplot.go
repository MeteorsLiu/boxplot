package boxplot

import (
	"slices"
	"sort"
)

type Boxplot struct {
	size  int
	stack []float64

	q0, q1, q2, q3, q4 float64
}

// Qk = xk + (percent * (n+1) - k)*(xk+1 - xk)
func quartile(arr []float64, percent float64) float64 {
	k := float64(len(arr)) * percent
	pos := int(k - 1)
	// xk+1
	xk1 := arr[pos+1]
	// xk
	xk := arr[pos]

	return xk + (percent*float64(len(arr)+1)-float64(k))*(xk1-xk)
}

func NewBoxPlot(size int) *Boxplot {
	return &Boxplot{
		size: size,
	}
}

func (b *Boxplot) Push(val float64) {
	if len(b.stack) >= b.size {
		// 淘汰掉最先入栈的
		_ = slices.Delete(b.stack, 0, 1)
		b.stack[len(b.stack)-1] = val
		return
	}
	b.stack = append(b.stack, val)
}

func (b *Boxplot) copyStack() []float64 {
	return append([]float64(nil), b.stack...)
}

func (b *Boxplot) Calculate() {
	if len(b.stack) == 0 {
		return
	}
	// sort会在原数组上进行排序操作，但我们需要保证原数组的顺序性
	// 所以需要重新拷贝一份
	copied := b.copyStack()
	sort.Float64s(copied)

	b.q0 = copied[0]
	b.q1 = quartile(copied, 0.25)
	b.q2 = quartile(copied, 0.5)
	b.q3 = quartile(copied, 0.75)
	b.q4 = copied[len(copied)-1]
}

func (b *Boxplot) Max() float64 {
	return b.q4
}

func (b *Boxplot) Min() float64 {
	return b.q0
}

func (b *Boxplot) IQR() float64 {
	return b.q3 - b.q1
}

func (b *Boxplot) Upper() float64 {
	return b.q3 + 1.5*b.IQR()
}

func (b *Boxplot) Lower() float64 {
	return b.q1 - 1.5*b.IQR()
}
