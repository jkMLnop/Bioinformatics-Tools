
package main
import (
	"fmt"
)

var (
	prob float64 = 0
	pXX  float64 = 15
	pXy  float64 = 25
	pyy  float64 = 16
)

func	calcProb (k,m,n float64) float64 {

	probDmt	:=	0.0

	probDmt	 =	((k * k - k) + 2 * (k * m) + 2 * (k * n) +
			(0.75 * (m * m - m) + 2 * (0.5 * m * n))) /
			((k + m + n) * (k + m + n -1))

	return probDmt;
}

func main() {
	prob = calcProb(pXX, pXy, pyy)
	fmt.Printf("%5f", prob)
	fmt.Println("")
}
