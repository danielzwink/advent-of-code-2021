package types

type Player struct {
	Score        int
	winningScore int
	position     int
}

func NewPlayer(winning int, start int) Player {
	return Player{Score: 0, winningScore: winning, position: start}
}

func (p *Player) Move(value int) {
	p.position += value
	p.position = p.position % 10
	if p.position == 0 {
		p.position = 10
	}
	p.Score += p.position
}

func (p *Player) Won() bool {
	return p.Score >= p.winningScore
}

type DeterministicDie struct {
	Rolls int
	value int
}

func (d *DeterministicDie) RollThreeTimes() int {
	step1 := d.roll()
	step2 := d.roll()
	step3 := d.roll()
	return step1 + step2 + step3
}

func (d *DeterministicDie) roll() int {
	d.Rolls++
	d.value++

	if d.value == 101 {
		d.value = 1
	}
	return d.value
}

type QuantumDie struct {
}
