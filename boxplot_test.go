package boxplot

import "testing"

func TestBoxplot(t *testing.T) {
	arr := []float64{57, 57, 57, 58, 63, 66, 66, 67, 67, 68, 69, 70, 70, 70, 70, 72, 73, 75, 75, 76, 76, 78, 79, 81}
	bl := NewBoxPlot(len(arr))

	for _, n := range arr {
		bl.Push(n)
	}

	bl.Calculate()

	if bl.q0 != 57 {
		t.Error("misbehave: q0")
		return
	}
	if bl.q1 != 66 {
		t.Error("misbehave: q1")
		return
	}
	if bl.q2 != 70 {
		t.Error("misbehave: q2")
		return
	}
	if bl.q3 != 75 {
		t.Error("misbehave: q3")
		return
	}
	if bl.q4 != 81 {
		t.Error("misbehave: q4")
		return
	}
	if bl.Upper() != 88.5 {
		t.Error("misbehave: upper")
		return
	}
	if bl.Lower() != 52.5 {
		t.Error("misbehave: lower")
		return
	}
}

func TestBoxplotOverlap(t *testing.T) {
	arr := []float64{57, 57, 57, 58, 63, 66, 66, 67, 67, 68, 69, 70, 70, 70, 70, 72, 73, 75, 75, 76, 76, 78, 79, 81}
	bl := NewBoxPlot(len(arr))

	for _, n := range arr {
		bl.Push(n)
	}

	for _, n := range arr {
		bl.Push(n)
	}

	bl.Calculate()

	if bl.q0 != 57 {
		t.Error("misbehave: q0")
		return
	}
	if bl.q1 != 66 {
		t.Error("misbehave: q1")
		return
	}
	if bl.q2 != 70 {
		t.Error("misbehave: q2")
		return
	}
	if bl.q3 != 75 {
		t.Error("misbehave: q3")
		return
	}
	if bl.q4 != 81 {
		t.Error("misbehave: q4")
		return
	}
	if bl.Upper() != 88.5 {
		t.Error("misbehave: upper")
		return
	}
	if bl.Lower() != 52.5 {
		t.Error("misbehave: lower")
		return
	}
}
