package model

//Galinha representa uma ave tipo galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave tipo pato
type Pata interface {
	Quack() string
}

//Ave é um tipo de pássaro
type Ave struct {
	Nome string
}

//Cacareja é o som da galinha
func (a Ave) Cacareja() string {

	return "Cocoricó"

}

//Quack é o som do pato
func (a Ave) Quack() string {

	return "Quack, quack"

}
