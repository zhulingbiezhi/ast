package code

type person struct {
	Male string
}

type men struct {
	person
	age string
}

type female struct {
	person
	weight string
	maps   map[string]*men
}

func pay(a, b string) string {
	return ""
}

func speak(p1 person) {

}

type test interface {
	speak(p1 person)
}

func (p *person) sp(a, b string) {

}
