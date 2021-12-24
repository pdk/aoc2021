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

	bits := getBits(sourceMessage)
	// bits := getBits("8A004A801A8002F478")
	// bits := getBits("620080001611562C8802118E34")
	// bits := getBits("C0015000016115A2E0802F182340")
	// bits := getBits("A0016C880162017C3686B18A3D4780")

	parsePackets(math.MaxInt, bits)

	log.Printf("sum of versions is %d", versionSum)

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

func parsePackets(maxPkts int, bits string) string {
	for i := 0; len(bits) > 0 && i < maxPkts; i++ {
		bits = parseOnePacket(bits)
	}

	return bits
}

func parseOnePacket(bits string) string {

	log.Printf("parseOnePacket length is %d", len(bits))

	if len(bits) < 8 {
		return ""
	}

	versionBits, bits := chew(3, bits)
	version := parseBits(versionBits, "versionBits")
	versionSum += version
	log.Printf("packet version %d", version)

	typeIDBits, bits := chew(3, bits)
	typeID := parseBits(typeIDBits, "typeIDBits")
	// log.Printf("package type ID %d", typeID)

	if typeID == 4 {
		return parseLiteral(bits)
	}

	return parseOperator(bits)
}

func parseOperator(bits string) string {

	lengthTypeID, bits := chew(1, bits)
	// log.Printf("lengthTypeID %s", lengthTypeID)

	if lengthTypeID == "0" {
		lengthBits, bits := chew(15, bits)
		return parseSubPacketsByBitLength(parseBits(lengthBits, "length type 0"), bits)
	}

	lengthBits, bits := chew(11, bits)
	return parseSubPacketsByPktCount(parseBits(lengthBits, "length type 1"), bits)
}

func parseSubPacketsByBitLength(length int, bits string) string {

	subPackets, remainBits := chew(length, bits)

	parsePackets(math.MaxInt, subPackets)

	return remainBits
}

func parseSubPacketsByPktCount(count int, bits string) string {
	return parsePackets(count, bits)
}

func parseLiteral(bits string) string {
	var bite string
	for {
		bite, bits = chew(5, bits)
		if bite[0] == '0' {
			return bits
		}
		// todo accumulate bites to build literal value
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
