package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// testChew()

	if err := run(os.Args, os.Stdout); err != nil {
		log.Fatalf("program failed: %s\n", err)
	}

	fmt.Printf("fini\n")
}

var versionSum int

func run(args []string, stdout io.Writer) error {

	var bits string

	bits = getBits(sourceMessage)
	// bits = getBits("8A004A801A8002F478")
	// bits = getBits("620080001611562C8802118E34")
	// bits = getBits("C0015000016115A2E0802F182340")
	// bits = getBits("A0016C880162017C3686B18A3D4780")

	// bits = getBits("C200B40A82") // finds the sum of 1 and 2, resulting in the value 3.
	// bits = getBits("04005AC33890") // finds the product of 6 and 9, resulting in the value 54.
	// bits = getBits("880086C3E88112") // finds the minimum of 7, 8, and 9, resulting in the value 7.
	// bits = getBits("CE00C43D881120") // finds the maximum of 7, 8, and 9, resulting in the value 9.
	// bits = getBits("D8005AC2A8F0") // produces 1, because 5 is less than 15.
	// bits = getBits("F600BC2D8F") // produces 0, because 5 is not greater than 15.
	// bits = getBits("9C005AC2F8F0") // produces 0, because 5 is not equal to 15.
	// bits = getBits("9C0141080250320F1802104A08") // produces 1, because 1 + 3 = 2 * 2.

	_, val := parsePackets(math.MaxInt, bits)

	log.Printf("end result val is %d", val)

	return nil
}

func getBits(message string) string {
	s := strings.Builder{}
	for _, e := range message {
		i, err := strconv.ParseInt(string(e), 16, 0)
		if err != nil {
			log.Fatalf("failed to parse hex %s: %v", string(e), err)
		}
		b := fmt.Sprintf("%04b", i)
		s.WriteString(b)
	}
	return s.String()
}

func parsePackets(maxPkts int, bits string) (string, []int) {
	var vals []int

	for i := 0; len(bits) > 0 && i < maxPkts; i++ {
		var eachVals []int
		bits, eachVals = parseOnePacket(bits)
		vals = append(vals, eachVals...)
	}

	return bits, vals
}

func parseOnePacket(bits string) (string, []int) {

	log.Printf("parseOnePacket length is %d", len(bits))

	if len(bits) < 8 {
		return "", []int{}
	}

	versionBits, bits := chew(3, bits)
	version := parseBits(versionBits, "versionBits")
	versionSum += version
	log.Printf("packet version %d", version)

	typeIDBits, bits := chew(3, bits)
	typeID := parseBits(typeIDBits, "typeIDBits")
	// log.Printf("package type ID %d", typeID)

	switch typeID {
	case 0:
		return sum(parseOperands(bits))
	case 1:
		return product(parseOperands(bits))
	case 2:
		return minimum(parseOperands(bits))
	case 3:
		return maximum(parseOperands(bits))
	case 5:
		return greaterThan(parseOperands(bits))
	case 6:
		return lessThan(parseOperands(bits))
	case 7:
		return equals(parseOperands(bits))
	}

	// case 4
	return parseLiteral(bits)
}

func sum(remain string, vals []int) (string, []int) {
	sum := 0
	for _, each := range vals {
		sum += each
	}
	return remain, []int{sum}
}

func product(remain string, vals []int) (string, []int) {
	product := 1
	for _, each := range vals {
		product *= each
	}
	return remain, []int{product}
}

func minimum(remain string, vals []int) (string, []int) {
	minimum := vals[0]
	for _, each := range vals {
		if each < minimum {
			minimum = each
		}
	}
	return remain, []int{minimum}
}

func maximum(remain string, vals []int) (string, []int) {
	maximum := vals[0]
	for _, each := range vals {
		if each > maximum {
			maximum = each
		}
	}
	return remain, []int{maximum}
}

func greaterThan(remain string, vals []int) (string, []int) {
	if vals[0] > vals[1] {
		return remain, []int{1}
	}
	return remain, []int{0}
}

func lessThan(remain string, vals []int) (string, []int) {
	if vals[0] < vals[1] {
		return remain, []int{1}
	}
	return remain, []int{0}
}

func equals(remain string, vals []int) (string, []int) {
	if vals[0] == vals[1] {
		return remain, []int{1}
	}
	return remain, []int{0}
}

func parseOperands(bits string) (string, []int) {

	lengthTypeID, bits := chew(1, bits)
	// log.Printf("lengthTypeID %s", lengthTypeID)

	if lengthTypeID == "0" {
		lengthBits, bits := chew(15, bits)
		return parseSubPacketsByBitLength(parseBits(lengthBits, "length type 0"), bits)
	}

	lengthBits, bits := chew(11, bits)
	return parseSubPacketsByPktCount(parseBits(lengthBits, "length type 1"), bits)
}

func parseSubPacketsByBitLength(length int, bits string) (string, []int) {

	subPackets, remainBits := chew(length, bits)

	_, vals := parsePackets(math.MaxInt, subPackets)

	return remainBits, vals
}

func parseSubPacketsByPktCount(count int, bits string) (string, []int) {
	return parsePackets(count, bits)
}

func parseLiteral(bits string) (string, []int) {
	var bite string
	var accum string

	for {
		bite, bits = chew(5, bits)
		accum += bite[1:5]
		if bite[0] == '0' {
			return bits, []int{parseBits(accum, "literal")}
		}
	}
}

func chew(n int, bits string) (string, string) {

	if n > len(bits) {
		log.Printf("chew(%d, ...) on len %d", n, len(bits))
		return bits, ""
	}

	return bits[0:n], bits[n:]
}

func testChew() {

	s := "1234567890"

	for i := 1; i <= 10; i++ {
		a, b := chew(i, s)
		log.Printf("chew(%d) %#v %#v", i, a, b)
	}

}

func parseBits(s string, errKey string) int {
	val, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		log.Fatalf("failed to parse bits %s at %s: %v", s, errKey, err)
	}

	return int(val)
}
