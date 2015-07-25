package main
//use this to transpose delimited files - must fit in memory
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var inDelim, outDelim string
func main() {
	flag.StringVar(&inDelim, "d", "\t", "input delimiter")
	flag.StringVar(&outDelim, "D", "\t", "output delimiter")
	flag.Parse()
	testFile := flag.Arg(0)
	inTest := os.Stdin
	stat, err := os.Stdin.Stat()
	check(err)
	if (stat.Mode() & os.ModeCharDevice) == 0 { //user is piping STDIN to program
		testFile = "STDIN"
	} else { //user is NOT piping STDIN to program
		inTest, err = os.Open(testFile)
		check(err)
	}

	out := make ([][]string, 1000)
	test := bufio.NewScanner(inTest)
	rows := 0
	for test.Scan() {
		cols := strings.Split(test.Text(), inDelim)
		out[rows] = cols
		rows++
		if rows>=len(out) {
			out = Expand(out)
		}
	}
	if err := test.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	length := len(out[0]) //assumes all rows in text file have same number of columns as first row
	for j:=0; j < length; j++ {
		for i:= 0; i < rows; i++ {
			fmt.Print(out[i][j])
			if i < rows-1 {
				fmt.Print(outDelim)
			}
		}
		fmt.Println()
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Expand(slice [][]string) [][]string {
	n := len(slice)
	newSize := n*3/2 + 1
	newSlice := make([][]string, newSize)
	copy(newSlice, slice)
	slice = newSlice
	return slice
}
