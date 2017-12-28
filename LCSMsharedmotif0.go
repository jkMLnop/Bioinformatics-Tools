package main

import (
	"fmt"
	"strings"
)

const(
	offset		= 12			/*offset constant accounts for FACTA header lengths*/
)

var (
	input		=
`>Rosalind_1
GATTACA
>Rosalind_2
ZZZZZZA
>Rosalind_3
TACCA`

	crnt_match	= ``
	best_match	= ``
	ss_start	= offset + 1		/*initialize substring as first two base positions*/
	ss_end		= offset + 2

	starts		= make([]int, 0)	/*creates empty slice of strand start positions */
	next_start	= 0
	next_end	= 0
)

func countStrands (in string){
	for i := range in {
		if in[i] == '>' {
			i = i + offset
			starts = append(starts, (i))	/*INEFFICIENT: re-initializes slice to house every additional start position*/
		}
	}
/*MAKE SEPERATE FUNCTION*/
	if len(starts) <= 1 {
		fmt.Println("insufficient strands to compare!")
		/*ADD END PROGRAM!!*/
	}else{
		next_start = starts[1]			/*sets initial comparison strand start as start index of second strand*/

		if len(starts) == 2 {
			next_end = (len(in) - 1)	/*in 2-strand case last position of 2nd strand is end of input string - sets as such*/

		}else if  len(starts) > 2 {
			next_end = starts[2] - (offset + 1)	/*in >2-strand strngs, 2nd strand ends on '>' so can use 3rd entry in starts slice*/
		}
	}
}

func findMatch (substring string, start_nxt int, end_nxt int) {
	/*ADD IF END_NEXT = LENGTH(STARTS) && MATCH FOUND END/RETURN______*/
	fmt.Println("substring being searched is: ", substring)
	fmt.Println("next strand is: ", input[start_nxt:end_nxt])

	width := 1
				/*next_string*/
	if strings.Contains(input[start_nxt:end_nxt], substring) == true { /*MAY WANNA CLEAN THIS UP...REMOVE IT EVENTUALLY*/
		fmt.Println("in here..")		/*next_string*/
		fmt.Println("current base: ", input[start_nxt:end_nxt],"substring: ", substring)

		if input[start_nxt:(start_nxt + width)] != substring {

			if input[start_nxt:(start_nxt + width)] != "\n" || input[ss_start:ss_end] != `>` {
				/*findMatch(input[start_nxt:end_nxt], start_nxt, end_nxt)/*ADD FUNCTIONALITY - increment depth by 1*/
				fmt.Println("substring", "'", substring, "'", "found at ss_start: ", ss_start, "and ss_end: ", ss_end)
			}
		}else{
			fmt.Println("base found at: ",start_nxt,end_nxt)
			/*call self, increase depth by 1*/
		}
		/*first check next strand until there is no next strand, if all match return remains true,*/
		/*if return is true exit that level of the sub and decrement depth by 1 (repeat until depth = 1)*/
		/*once return = true and depth = 1 increment substring size by one and call self with newly longer substring input*/
		/*NOTE: BE SURE TO INCREMENT DEPTH IN CONTROLLED MANNER SO NO DEEP RECURSION OCCURS!!*/
	}else{

		fmt.Println("substring not found at ss_start: ", ss_start, "and ss_end: ", ss_end)
		/*PROB CUT THIS../CHANGE IN SOME WAY..*/
		ss_start ++
		ss_end = ss_start + 1

		if input[ss_start:ss_end] != "\n" || input[ss_start:ss_end] != `>` {
			findMatch(input[ss_start:ss_end], next_start, next_end)/*ADD FUNCTIONALITY - increment depth by 1*/
		}
	}
}

func main () {

	countStrands(input)	/*populates array */

	fmt.Println(starts)
		  /*MATCH SUSBSTRING    , NEXT STRAND START, END NEXT STRAND, **LAST STRAND(BOOL)** */
	findMatch(input[ss_start:ss_end], next_start, next_end)

	/*fmt.Println(best_match) */
}

