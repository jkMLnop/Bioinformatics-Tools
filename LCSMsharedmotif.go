package main

import (
	"fmt"
)

const(
	offset		= 12			/*offset constant accounts for FACTA header lengths*/
)

var (
	input		=
`>Rosalind_1
GATTACA
>Rosalind_2
KQDLZZA
>Rosalind_3
GCAAT`

	crnt_match	= ``
	best_match	= ``
	ss_start	= offset + 1		/*initialize substring as first two base positions*/
	ss_end		= offset + 2

	starts		= make([]int, 0)	/*creates empty slice of strand start positions */
	ends		= make([]int, 0)	/*creates empty slice of strand end positions */
	next_start	= 0
	next_end	= 0
	width		= 1			/*TODO implement width - frame expands after a single match found in all strands*/
	count		= 0
	last		= 0			/*last possible position of substring (must be within first strand) being searched for*/
						/*FOR BASE CASE WHERE NO MATCH FOUND FOR ALL OTHER SS POSITIONS!!*/
)

func countStrands (in string){
	for i := range in {
		if in[i] == '>' {	/*Append start positions of all but last strands*/

			if i > offset {	/*when the current index is not at the first element*/
				ends = append(ends,(i - 1))
			}

			i = i + offset
			starts = append(starts, (i))
		}
		if i == len(in) - 1 {	/*Append last base position of last strand*/
			ends = append(ends, (i))
		}
	}

	if len(starts) <= 1 {
		fmt.Println("insufficient strands to compare!")
	}else{
		count      = 1
		next_start = starts[count]			/*sets initial comparison strand start as start index of second strand*/
		next_end   = ends[count]
	}
}

func findMatch (start_sbstr, end_sbstr, start_nxt, end_nxt int) {

	/*width := 1*/
	/*TODO FIX END_NXT(end_mtch) LOGIC SOMEHOW.. make end_mtch = start_mtch + width? */
	fmt.Println("comparison strand: ", input[start_nxt:end_nxt],"substring: ", input[start_sbstr:end_sbstr])

	if input[start_sbstr:end_sbstr] != input[start_nxt:end_nxt] && end_nxt <= len(input){	/*no match on current position of next strand*/
		fmt.Println("current comparision base: ", input[start_nxt:end_nxt],"moving next base: ",input[(start_nxt + 1):(end_nxt + 1)])
	}else{	/*match on current position of next strand*/

	}
}

func main () {

	countStrands(input)	/*populates array */

	fmt.Println("ends: ", ends)	/*TODO last element of this array is not 5 positions from first, is that correct?!*/

	findMatch(ss_start, ss_end, next_start, next_end)

}

