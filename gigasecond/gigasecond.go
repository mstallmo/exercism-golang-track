//gigasecond package has one function AddGigasecond that takes in a time and adds a gigasecond to that time
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond takes in a time struct as an argument and returns the passed argument plus one gigasecond
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
