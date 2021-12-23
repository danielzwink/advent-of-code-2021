package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Vector struct {
	x, y, z int
}

func NewOriginVector() Vector {
	return Vector{x: 0, y: 0, z: 0}
}

func (v Vector) ManhattanDistance(o Vector) int {
	x := int(math.Abs(float64(v.x - o.x)))
	y := int(math.Abs(float64(v.y - o.y)))
	z := int(math.Abs(float64(v.z - o.z)))
	return x + y + z
}

func (b Beacon) String() string {
	return fmt.Sprintf("%+04d,%+04d,%+04d", b.x, b.y, b.z)
}

type Beacon struct {
	Vector
}

func NewBeacon(line string) Beacon {
	numbers := strings.Split(line, ",")
	x, _ := strconv.Atoi(numbers[0])
	y, _ := strconv.Atoi(numbers[1])
	z, _ := strconv.Atoi(numbers[2])
	return Beacon{Vector{x: x, y: y, z: z}}
}

func NewBeaconMoved(x, y, z int, move Vector) Beacon {
	return Beacon{Vector{x: x + move.x, y: y + move.y, z: z + move.z}}
}

func (b Beacon) Equal(o Beacon) bool {
	return b.x == o.x && b.y == o.y && b.z == o.z
}

func (b Beacon) Diff(o Beacon) Vector {
	return Vector{x: b.x - o.x, y: b.y - o.y, z: b.z - o.z}
}

type Universe struct {
	scanners []*Scanner
}

func NewUniverse() *Universe {
	return &Universe{make([]*Scanner, 0, 25)}
}

func (u *Universe) Get(i int) *Scanner {
	return u.scanners[i]
}

func (u *Universe) Add(s *Scanner) {
	u.scanners = append(u.scanners, s)
}

func (u *Universe) Del(i int) {
	l := len(u.scanners)

	if i == 0 {
		u.scanners = u.scanners[1:]
	} else if i == l-1 {
		u.scanners = u.scanners[:l-1]
	} else {
		before := u.scanners[0:i]
		after := u.scanners[i+1 : l]
		u.scanners = append(before, after...)
	}
}

func (u *Universe) Pop() *Scanner {
	l := len(u.scanners)
	last := u.scanners[l-1]
	u.scanners = u.scanners[:l-1]
	return last
}

func (u *Universe) Top() *Scanner {
	first := u.scanners[0]
	u.scanners = u.scanners[1:]
	return first
}

func (u *Universe) Len() int {
	return len(u.scanners)
}

const (
	XpYpZp = iota
	XpYmZm
	XmYmZp
	XmYpZm
	XpZpYm
	XmZmYm
	XmZpYp
	XpZmYp
	ZpXpYp
	ZpXmYm
	ZmXmYp
	ZmXpYm
	YpXpZm
	YmXmZm
	YmXpZp
	YpXmZp
	YpZpXp
	YpZmXm
	YmZmXp
	YmZpXm
	ZpYpXm
	ZmYmXm
	ZmYpXp
	ZpYmXp
)

type Scanner struct {
	Location    Vector
	Orientation int
	beacons     []Beacon
}

func NewScanner() *Scanner {
	return &Scanner{Location: NewOriginVector(), Orientation: XpYpZp, beacons: make([]Beacon, 0)}
}

func (s *Scanner) Add(beacon Beacon) {
	s.beacons = append(s.beacons, beacon)
}

func (s *Scanner) Len() int {
	return len(s.beacons)
}

func (s *Scanner) Calibrate(beaconIndex, orientation int, target Beacon) {
	s.Location = NewOriginVector()
	s.Orientation = orientation
	beacon := s.NormalizedBeacon(beaconIndex)
	currentDiff := target.Diff(beacon)
	s.Location = currentDiff
}

func (s *Scanner) NormalizedBeacons() []Beacon {
	beacons := make([]Beacon, s.Len())
	for b := 0; b < s.Len(); b++ {
		beacons[b] = s.NormalizedBeacon(b)
	}
	return beacons
}

func (s *Scanner) NormalizedBeacon(i int) Beacon {
	b := s.beacons[i]

	switch s.Orientation {
	case XpYpZp:
		return NewBeaconMoved(b.x, b.y, b.z, s.Location)
	case XpYmZm:
		return NewBeaconMoved(b.x, -b.y, -b.z, s.Location)
	case XmYmZp:
		return NewBeaconMoved(-b.x, -b.y, b.z, s.Location)
	case XmYpZm:
		return NewBeaconMoved(-b.x, b.y, -b.z, s.Location)
	case XpZpYm:
		return NewBeaconMoved(b.x, b.z, -b.y, s.Location)
	case XmZmYm:
		return NewBeaconMoved(-b.x, -b.z, -b.y, s.Location)
	case XmZpYp:
		return NewBeaconMoved(-b.x, b.z, b.y, s.Location)
	case XpZmYp:
		return NewBeaconMoved(b.x, -b.z, b.y, s.Location)
	case ZpXpYp:
		return NewBeaconMoved(b.z, b.x, b.y, s.Location)
	case ZpXmYm:
		return NewBeaconMoved(b.z, -b.x, -b.y, s.Location)
	case ZmXmYp:
		return NewBeaconMoved(-b.z, -b.x, b.y, s.Location)
	case ZmXpYm:
		return NewBeaconMoved(-b.z, b.x, -b.y, s.Location)
	case YpXpZm:
		return NewBeaconMoved(b.y, b.x, -b.z, s.Location)
	case YmXmZm:
		return NewBeaconMoved(-b.y, -b.x, -b.z, s.Location)
	case YmXpZp:
		return NewBeaconMoved(-b.y, b.x, b.z, s.Location)
	case YpXmZp:
		return NewBeaconMoved(b.y, -b.x, b.z, s.Location)
	case YpZpXp:
		return NewBeaconMoved(b.y, b.z, b.x, s.Location)
	case YpZmXm:
		return NewBeaconMoved(b.y, -b.z, -b.x, s.Location)
	case YmZmXp:
		return NewBeaconMoved(-b.y, -b.z, b.x, s.Location)
	case YmZpXm:
		return NewBeaconMoved(-b.y, b.z, -b.x, s.Location)
	case ZpYpXm:
		return NewBeaconMoved(b.z, b.y, -b.x, s.Location)
	case ZmYmXm:
		return NewBeaconMoved(-b.z, -b.y, -b.x, s.Location)
	case ZmYpXp:
		return NewBeaconMoved(-b.z, b.y, b.x, s.Location)
	case ZpYmXp:
		return NewBeaconMoved(b.z, -b.y, b.x, s.Location)
	}
	panic("unknown position value: " + strconv.Itoa(s.Orientation))
}
