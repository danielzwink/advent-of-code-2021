package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	binary := readHexToBinary()

	part1Result, part1Duration := part1(binary)
	fmt.Printf("Part 1: %12d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(binary)
	fmt.Printf("Part 2: %12d (duration: %s)\n", part2Result, part2Duration)
}

func part1(binary string) (int, time.Duration) {
	start := time.Now()
	packet, _ := evaluatePacket(binary)
	return packet.versions(), time.Since(start)
}

func part2(binary string) (int, time.Duration) {
	start := time.Now()
	packet, _ := evaluatePacket(binary)
	return packet.value(), time.Since(start)
}

func evaluatePacket(binary string) (Packet, int) {
	typeID := convertBinaryToDecimal(binary[3:6])

	if typeID == 4 {
		return readLiteralValuePacket(binary)
	} else {
		return readOperatorPackets(binary)
	}
}

func readLiteralValuePacket(binary string) (Packet, int) {
	version := convertBinaryToDecimal(binary[:3])
	typeID := convertBinaryToDecimal(binary[3:6])

	position := 6
	builder := strings.Builder{}
	for {
		builder.WriteString(binary[position+1 : position+5])

		if binary[position:position+1] == "0" {
			position += 5
			break
		} else {
			position += 5
		}
	}

	number := convertBinaryToDecimal(builder.String())
	packet := NewLiteralValuePacket(version, typeID, number)
	return packet, position
}

func readOperatorPackets(binary string) (Packet, int) {
	version := convertBinaryToDecimal(binary[:3])
	typeID := convertBinaryToDecimal(binary[3:6])
	packet := NewOperatorPacket(version, typeID)

	position := 6
	if binary[position:position+1] == "0" {
		position += 1
		packetsLength := convertBinaryToDecimal(binary[position : position+15])
		position += 15

		for packetsLength > 0 {
			innerPacket, length := evaluatePacket(binary[position:])
			packet.addPacket(innerPacket)
			packetsLength -= length
			position += length
		}

	} else if binary[position:position+1] == "1" {
		position += 1
		packetsCount := convertBinaryToDecimal(binary[position : position+11])
		position += 11

		for i := 0; i < packetsCount; i++ {
			innerPacket, length := evaluatePacket(binary[position:])
			packet.addPacket(innerPacket)
			position += length
		}
	}

	return packet, position
}

func convertBinaryToDecimal(binary string) int {
	i, _ := strconv.ParseUint(binary, 2, len(binary))
	return int(i)
}

func readHexToBinary() string {
	line := util.ReadFile("16")[0]

	binary := strings.Builder{}
	binary.Grow(len(line) * 4)

	for _, char := range line {
		switch char {
		case '0':
			binary.WriteString("0000")
		case '1':
			binary.WriteString("0001")
		case '2':
			binary.WriteString("0010")
		case '3':
			binary.WriteString("0011")
		case '4':
			binary.WriteString("0100")
		case '5':
			binary.WriteString("0101")
		case '6':
			binary.WriteString("0110")
		case '7':
			binary.WriteString("0111")
		case '8':
			binary.WriteString("1000")
		case '9':
			binary.WriteString("1001")
		case 'A':
			binary.WriteString("1010")
		case 'B':
			binary.WriteString("1011")
		case 'C':
			binary.WriteString("1100")
		case 'D':
			binary.WriteString("1101")
		case 'E':
			binary.WriteString("1110")
		case 'F':
			binary.WriteString("1111")
		}
	}
	return binary.String()
}

type Packet struct {
	version int
	typeID  int
	number  int
	packets []Packet
}

func NewLiteralValuePacket(version, typeID, number int) Packet {
	return Packet{version: version, typeID: typeID, number: number}
}

func NewOperatorPacket(version, typeID int) Packet {
	return Packet{version: version, typeID: typeID, packets: make([]Packet, 0)}
}

func (p *Packet) addPacket(packet Packet) {
	p.packets = append(p.packets, packet)
}

func (p *Packet) versions() int {
	result := p.version
	for _, packet := range p.packets {
		result += packet.versions()
	}

	return result
}

func (p *Packet) value() int {
	switch p.typeID {
	case 0:
		sum := 0
		for _, packet := range p.packets {
			sum += packet.value()
		}
		return sum
	case 1:
		product := p.packets[0].value()
		for _, packet := range p.packets[1:] {
			product *= packet.value()
		}
		return product
	case 2:
		min := p.packets[0].value()
		for _, packet := range p.packets[1:] {
			if min > packet.value() {
				min = packet.value()
			}
		}
		return min
	case 3:
		max := p.packets[0].value()
		for _, packet := range p.packets[1:] {
			if max < packet.value() {
				max = packet.value()
			}
		}
		return max
	case 4:
		return p.number
	case 5:
		if p.packets[0].value() > p.packets[1].value() {
			return 1
		}
		return 0
	case 6:
		if p.packets[0].value() < p.packets[1].value() {
			return 1
		}
		return 0
	case 7:
		if p.packets[0].value() == p.packets[1].value() {
			return 1
		}
		return 0
	}
	return 0
}
