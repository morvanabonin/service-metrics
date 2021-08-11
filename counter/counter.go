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

func countNumber(out chan<- int) {
	for i := 0; i < 100; i++ {
		nCounter = append(nCounter, i)
		out <- len(nCounter)
	}
	close(out)
	//return len(nCounter)
}

func MetricsCountQuery() int {
	total := make(chan int)

	go countNumber(total)
	return printer(total)
}

func printer(in <-chan int) int {
	var total int
	for v := range in {
		total = v
	}
	return total
}
