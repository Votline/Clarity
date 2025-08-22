package ui

import (
	"Clarity/internal/render"
)

type elemMesh struct {
	vao uint32
	vtq int32
}

type HomeView struct {
	elems [10]elemMesh
	pg uint32
	ofL int32
}

func CreateView(pg uint32, ofl int32) *HomeView {
	var elems [10]elemMesh
	for i := 0; i < 10; i++ {
		d := digit{num: i}
		el := d.create(0)
		vao := render.CreateVAO(el.getVtc())
		elems[i] = elemMesh{vao: vao, vtq: el.getVtq()}
	}

	return &HomeView{elems: elems, pg: pg, ofL: ofl}
}

func (hv *HomeView) Render(num int) {
	digit1 := num / 10
	digit2 := num % 10
	render.ElemRender(hv.pg, hv.ofL,
		hv.elems[digit1].vao, hv.elems[digit1].vtq, 0.0)
	render.ElemRender(hv.pg, hv.ofL,
		hv.elems[digit2].vao, hv.elems[digit2].vtq, 0.15)
}
