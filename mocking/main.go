package countdown

import (
	"os"
)

func main() {
	myDefault := &DefaultSleeper{}

	Countdown(os.Stdout, myDefault)
}
