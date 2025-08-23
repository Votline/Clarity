package ui

import (
	"Clarity/internal/render"
)

type elemMesh struct {
	vao uint32
	vtq int32
}

type HomeView struct {
	elems [11]elemMesh
	pg uint32
	ofL int32
}

func CreateView(pg uint32, ofl int32) *HomeView {
	var elems [11]elemMesh
	for i := 0; i < 11; i++ {
		d := digit{}
		if i == 10 {
			d.setNum(rune('.'))
		} else {
			d.setNum(rune(i))
		}
		el := d.create()
		vao := render.CreateVAO(el.getVtc())
		elems[i] = elemMesh{vao: vao, vtq: el.getVtq()}
	}

	return &HomeView{elems: elems, pg: pg, ofL: ofl}
}

func (hv *HomeView) Render(num string) {
	nums := []rune(num)
	offset := float32(0.0)
	for _, char := range nums {
		if char == '.' {
			render.ElemRender(hv.pg, hv.ofL,
				hv.elems[10].vao, hv.elems[10].vtq, offset)
			offset += 0.09
			continue
		}
		digit := int(char - '0')
		render.ElemRender(hv.pg, hv.ofL,
			hv.elems[digit].vao, hv.elems[digit].vtq, offset)
		offset += 0.15
	}
}
