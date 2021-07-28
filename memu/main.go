package main

import "fmt"

/*
                                                     4  2  0
            |                        |        | +----+--+--+
mov   rt,rs | R[rt] <- R[rs]         | R-type | |0000|rt|rs|
            |                        |        | +----+--+--+
            |                        |        | +----+--+--+
add   rt,rs | R[rt] <- R[rs] + R[rt] | R-type | |0001|rt|rs|
            |                        |        | +----+--+--+
            |                        |        | +----+--+--+
load  addr  | R[0] <- M[addr]        | M-type | |1110| addr|
            |                        |        | +----+--+--+
            |                        |        | +----+--+--+
store addr  | M[addr] <- R[0]        | M-type | |1111| addr|
            |                        |        | +----+--+--+
*/

// NREG 寄存器数量
const NREG = 4

// NMEM 内存大小
const NMEM = 16

var PC uint8 = 0
var R [NREG]uint8
var M = [NMEM]uint8{
	0b11100110, // load 6#       | R[0] = 48
	0b00000100, // move r1, r0   | R[1] = R[0]
	0b11100101, // load 5#       | R[0] = 33
	0b00010001, // add  r0, r1   | R[0] = R[0] + R[1] = 81
	0b11110111, // store 7       | M[7] = R[0] = 81
	0b00110000, // x = 48
	0b00100001, // y = 33
	0b00000000, // z = 0
}
func getRTAndRS(instruction uint8) (rt, rs uint8) {
	rt = (instruction >> 2) & 0b000000011
	rs = instruction & 0b00000011
	return
}

func getOp(instruction uint8) uint8 {
	return instruction >> 4
}

func getMemoryAddr(instruction uint8) uint8 {
	return instruction & 0b00001111
}

var halt = false

func exec_once() {
	instruction := M[PC]
	switch getOp(instruction) {
	case 0b0000:
		rt, rs := getRTAndRS(instruction)
		R[rt] = R[rs]
	case 0b0001:
		rt, rs := getRTAndRS(instruction)
		R[rt] += R[rs]
	case 0b1110:
		R[0] = M[getMemoryAddr(instruction)]
	case 0b1111:
		M[getMemoryAddr(instruction)] = R[0]
	default:
		halt = true
	}

	PC += 1
}

func main() {
	for {
		exec_once()
		if halt {
			break
		}
	}
	fmt.Printf("16 + 33 = %d\n", M[7])
}
