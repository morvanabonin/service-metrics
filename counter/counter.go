package counter

type (
	CounterService interface {
		CounterNumber() int
	}

	Counter interface {
		CounterService
	}
)

var (
	nCounter = make([]int, 0)
)

func CountNumber() int {
	for i := 0; i < 84; i++ {
		nCounter = append(nCounter, i)
	}
	return len(nCounter)
}
