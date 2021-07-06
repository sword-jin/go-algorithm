package hanoi

type Frame struct {
	pc, n         int
	from, to, via string
}

func Simunator(n int, from, to, via string, move MoveFunc) {
	var (
		top    = 0
		stacks = make([]*Frame, 64)

		call = func(n int, from, to, via string) {
			stacks[top] = &Frame{0, n, from, to, via}
			top += 1
		}
		ret = func() {
			top -= 1
		}
	)

	call(n, from, to, via)

START:
	for top > 0 {
		frame := stacks[top-1]

		switch frame.pc {
		case 0:
			if frame.n == 1 {
				move(frame.from, frame.to)
				frame.pc = 4 // or ret()
				goto START
			}
		case 1:
			call(frame.n-1, frame.from, frame.via, frame.to)
		case 2:
			call(1, frame.from, frame.to, frame.via)
		case 3:
			call(frame.n-1, frame.via, frame.to, frame.from)
		case 4:
			ret()
		}
		frame.pc++
	}
}
