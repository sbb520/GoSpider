package core

type persistenter struct {

}

func NewPersistenter() *persistenter {
	p := new(persistenter)
	return p
}

func (p *persistenter) SaveToSQL () {
	
}