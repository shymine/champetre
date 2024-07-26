package champetre

type repository struct {
	elements []Model
}

func (rp *repository) Add(model Model) {
	rp.elements = append(rp.elements, model)
}

func (rp *repository) Replace(model Model) {
	for i, el := range rp.elements {
		if el.Id() == model.Id() {
			rp.elements[i].Modify(model)
			break
		}
	}
}