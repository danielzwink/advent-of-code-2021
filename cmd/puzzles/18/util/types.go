package types

import (
	"math"
	"strconv"
	"strings"
)

type Pair struct {
	Left   *Element
	Right  *Element
	Parent *Pair
}

func (p *Pair) Magnitude() int {
	return 3*p.Left.magnitude() + 2*p.Right.magnitude()
}

func (p *Pair) Add(other *Pair) *Pair {
	sum := &Pair{Left: &Element{Pair: p}, Right: &Element{Pair: other}}
	p.Parent = sum
	other.Parent = sum
	return sum
}

func (p *Pair) Reduce() *Pair {
	performed := true
	for performed {
		_, exploded, _ := p.explode(0)
		performed = exploded || p.split()
	}
	return p
}

func (p *Pair) Copy() *Pair {
	return &Pair{Left: p.Left.Copy(), Right: p.Right.Copy()}
}

func (p *Pair) explode(depth int) (*Pair, bool, bool) {
	depth++

	if p.Left.IsPair() {
		up, found, remove := p.Left.Pair.explode(depth)
		if found {
			if remove {
				p.Left.Pair = nil
				p.Left.Number = 0

				if p.Right.IsPair() {
					p.Right.Pair.addLeftMost(up.Right.Number)
				} else {
					p.Right.Number += up.Right.Number
				}

				current := p
				parent := p.Parent
				for parent != nil {
					if parent.Right.IsPair() && parent.Right.Pair == current {
						if parent.Left.IsPair() {
							parent.Left.Pair.addRightMost(up.Left.Number)
						} else {
							parent.Left.Number += up.Left.Number
						}
						break
					}
					current = parent
					parent = parent.Parent
				}
			}
			return up, found, false
		}
	}
	if p.Right.IsPair() {
		up, found, remove := p.Right.Pair.explode(depth)
		if found {
			if remove {
				p.Right.Pair = nil
				p.Right.Number = 0

				if p.Left.IsPair() {
					p.Left.Pair.addRightMost(up.Left.Number)
				} else {
					p.Left.Number += up.Left.Number
				}

				current := p
				parent := p.Parent
				for parent != nil {
					if parent.Left.IsPair() && parent.Left.Pair == current {
						if parent.Right.IsPair() {
							parent.Right.Pair.addLeftMost(up.Right.Number)
						} else {
							parent.Right.Number += up.Right.Number
						}
						break
					}
					current = parent
					parent = parent.Parent
				}
			}
			return up, found, false
		}
	}

	return p, depth > 4, depth > 4
}

func (p *Pair) addLeftMost(number int) {
	if p.Left.IsPair() {
		p.Left.Pair.addLeftMost(number)
	} else {
		p.Left.Number += number
	}
	return
}

func (p *Pair) addRightMost(number int) {
	if p.Right.IsPair() {
		p.Right.Pair.addRightMost(number)
	} else {
		p.Right.Number += number
	}
	return
}

func (p *Pair) split() bool {
	if p.Left.IsPair() && p.Left.Pair.split() {
		return true
	} else if p.Left.Number > 9 {
		p.Left.split(p)
		return true
	}
	if p.Right.IsPair() && p.Right.Pair.split() {
		return true
	} else if p.Right.Number > 9 {
		p.Right.split(p)
		return true
	}
	return false
}

func (p *Pair) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	builder.WriteString(p.Left.String())
	builder.WriteString(",")
	builder.WriteString(p.Right.String())
	builder.WriteString("]")
	return builder.String()
}

type Element struct {
	Pair   *Pair
	Number int
}

func (e *Element) IsPair() bool {
	return e.Pair != nil
}

func (e *Element) Copy() *Element {
	if e.IsPair() {
		return &Element{Pair: e.Pair.Copy()}
	} else {
		return &Element{Number: e.Number}
	}
}

func (e *Element) magnitude() int {
	if e.IsPair() {
		return e.Pair.Magnitude()
	} else {
		return e.Number
	}
}

func (e *Element) split(parent *Pair) {
	result := float64(e.Number) / 2.0
	round := math.Round(result)
	e.Pair = &Pair{Left: &Element{Number: int(result)}, Right: &Element{Number: int(round)}, Parent: parent}
	e.Number = 0
}

func (e *Element) String() string {
	builder := strings.Builder{}
	if e.IsPair() {
		builder.WriteString(e.Pair.String())
	} else {
		builder.WriteString(strconv.Itoa(e.Number))
	}
	return builder.String()
}
