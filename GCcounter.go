
package main

import (
	"fmt"
)

var (
	fasta	=(`>Rosalind_9454
	CCGGAGAAACAGGTTCCACCCGAGTGCACTATTCTGTTCGTAAATATGAACAGCTTCGGA
	GAACTTCGGATACATGGAATGGGACTCGTAGTAAAACATACTTTTAGCCTCAGAAGAGGC
	GACTTGGGCCAACTAGGGCCCAAGTTTAAAGCATTCCAGAATTCAAGAAAGTTGTCTGCT
	TTTACCGCTAATTTCACCCGTTATTGGAAGTTTAATTCCTAGGGAACCTTTGTGAACATG
	ACAAGGTGGCACTAGCATCGAGCCGTGCTCACTTTGGCGTGGGTGTAGAACGCGAATGGG
	AACCGATCTTACCAGGGGCATGCTCAATCGATGTCATACGTTAGTCACCCATCCGATACA
	TTGGATACTTTTAGGCAAAGTAGGCCTATACTTACCGCAACGTTCGGGATACTAACTCCA
	ACCACACTGCATATGTCTAGCTCCCGACACGGAGTAGCCATTAACTCACCCCGTCAACAG
	TCCTTCCTTATGAGAGCGTCTGATAACTAGCAGCTTTGGCGTGATCGGATAGGCCGAAGA
	ACGGCTCTGCCTCGTTATTGCTTGGATTGTTTACGCTCACCGTACAGGCATCCGCATAGG
	AACTTGTACTAGGATCGCTTTGCCATCCGGGTGTTGGGGCTCACCAGAAAGCACGACTTA
	TTGCACGACCCCATGTTTTCTGGGTGCGTCGTCAGAGATTGCTTTGAGTGCTTAAACCTC
	AATCACGTAGCTCCTTAAATGGGCTAAAAATTGCAAAAGGTCTTTCCGCTGCTCGAACCA
	TTAATAACAGTCCTAGGGCTTACTCGTGCCGTTATCTGGTTGGCGTCGGCGTCTATACCT
	TCTGACTGTATATTTAGGCCT
	>Rosalind_0200
	TGGACCCCTGCGAATCTAAGAAGAAACGGGATCCTAGAGACGTTTTAATGTTAGCAGGGA
	GCAAACTCCTTAACTTTTGAGTCGTGTTTTCAGGGGTATCTGAAGTGGGAGACCAGTTCT
	GGACGCCGTGTTGGCGACACGTCGTAACCGCCGCAACTGGTGCTCCGTGCTTCCTGAGTT
	TTGTATAGGGGATATCTCGCAAATGTCAAGGACTCAACTTCCCTCCCAACGACCCTATGG
	AGTACATCACAGCGTATCAGACCGAACTCAGAAGTTGCCTGACGATACAGATATTAAAAT
	TTATCCTACTAATTCAGCCCCGTCCCAGCAAAGGAGTAATCTCAGTAAATTTATAGTTTT
	ACGATCTCAATATAGCCATTGTACGTGTCTTCAATTGACATATGATTTCAGTAAGGTCTA
	GCTCGTTATAATCTTTGCGTAGAGCGCTGTTAATGTCTGAGACATGCGTACGCAATTTCA
	TTAAAGTCGGGAGCGCTCAATACGACAGCCTTTCAAAAGGCGCCAGAGTATGACGATTGC
	CGACCTACGCGCCGCATGTTCAACACATCAAATTAGGCCCCATCGATAACGATGGTGGTT
	ATCATGTCCATTGCAGCTGACACGCTGGTGCAACAATGCTGTATCCATAACACAGTTGAC
	CGAGTTAGGGCTCCGAGTCAGTGTAGAATCGCAGGGGCGAAGGACAACTTTGCAGTCATC
	AGCTAGCCGTAACTGGAATATGCATGGGGTGGGGTAGTCGTGACTTGCACCGGTGCTAAA
	TCCCCATTTCCCCGAGAACAGCCCGGTTGAAAAGGGAGATGGGCGGAATAGCACAACTTT
	CGTGTATGATTGGCATAAGCGGGAAACTGTATATCGAGAACGCATCGAGTAGCGCTTTCG
	GAATACTCATGCTGGTCAACCCAGTTCACTGTGCACTTATGCGGTCTG
	>Rosalind_1410
	GATTGGATAAAGCTCTCACCTGGTCCTTGAGATTCGAAGGGCGGGGAGTCGAACCCAGGC
	AAAATTTAGCTTTTGAGTAGGCTAGTTCTCTTGCAGTCCTACTACCTTGGAACGCAGCTT
	CAGCAGCGAAGTCTATGACATCAGTGACCGTGTACGGATATGCTGTGGGCCAACTGTAAT
	AGAATCATAGCGGACGACGATGACACTGGCCTGCCCACCATAGGGCATCGTGACGGTGCC
	CCTAGGCACGCGTTCTTCCTCGGCTAGTTAATCTCTTATGAGTGAGCAACGATGACGTCA
	ATTTTCCAACTAGGTTCTAAAAAAGGACTTGTACTGCGGGTCTGTTCATTTTTTGAAGAC
	GCGACGCTATGGTACGCAACTCGGTCGACGCGATAATATCTTTTTCGCACTCAGCCCGTA
	GTTCGGGTCCCCGTGGATCAGTCTTCCCAGATCGCCAGCCATAACACTGCCGAGATATGC
	CAATATGAACTTCCGTTACCCCGCCCCTTTCTCCGATTATCTACACGTCGGAGCTGGCGC
	AGTATCTCGGAAATTCTAGGGAGACTAGAACACTTAGCCTCAGCGATTCAAATTTGACGC
	ATAGCATAACGTTGTCACGAACTACTGGCCTATGGGCCCAGTTTAACGCTTCTGCGACGG
	ATCTAATCTTAACCGCAGCGCTGCTTATTACCCTCTAGACTCGCTGACATGAGGCGCAAT
	CTCATTGTCATTAGATCTCCCTCCCCCTAGGTCTTATTTCTTTCGATGTTAGTATCCGCT
	CTCACATCCTTCTACCAGGTGTTAAAGTGAGTCGCCCGTCCAACCCCCACGCAAGAATTG
	CCCTGACAAGTGTGAACTATTAGCGTCACCGGA
	>Rosalind_0695
	GTAAGAACAAAGAACAGTTGAAGTCTCGGACTCGCGGAACGCGGGCATTAAGGGCATGGA
	GGCGCGCAGAGCATAACTTCGTTAAGCCTGCCCCTTATTAGATCTTTCGTCCGTGGTTGG
	GTGAGGCTTGCTATTGCGACGTATAGCGGCTTGGCGGCGGGAAATTCGCTACTCAGAATT
	GTAGAGTTTTCGCTCAGGCGTTCTCATTCTGGCATAAGCCCCACCTGAATTTCAGGCGCT
	ATGATCTACAATCTGCCGCGAACGTGATCGGCAATGTCGACCGGTCAGTGTTCGTCAACG
	TTTACATAATCCTGATTAAAAATCGCTACTCGGCATTGCGGTATACGGAACGGAGTCTCA
	ACCAAGTGCTGAACAACGTGAGACACCTTAGAATTTTCCGTTCACCGTCGTCTGCAAGGA
	AGATGCAGAGCTAATATCATACATTAGGGTTTGTGCGCTTCCAAGAAACCAATGACCGTT
	TCGCGGTACTGCATGGAAGCCAGCCCTCCCACACCGAATTTCTGGTATTTTCGGATGCTG
	ATTACTGCAAACCAGACTCAAATCATTATCCCTCCAAGACACTTCAACACGGTCCGCTGT
	CATCGTGTACTCAGCGTCGACCTGTTCTTGGACGGTGCGCGTACCAGTTGGGTGTAAGCA
	CAATTCGGACTGCCGTAGCAGCATAATGCCCGGCCTATAACGAACAATAGAAAAGTTCGC
	CGGGCGTAGGACTAGCACCAGAGGTTCAAACTAAGTCCCAAAGGCCCCGGCGTGAAACGC
	GTACTGTACGCCAAGAGCGAAGGGGATTTAGACCCACAAATGGGTTGTCACCCGAAGGAG
	GGCCACATGTATATAACATGGTGTGGTACACCTGCTTGCTTACTCATCCTATGCCTGCAA
	CTATATAGGTCCGAA
	>Rosalind_8077
	AGCCAAGGGTGCCGGGGTGCTACGACTGGAAACAAAGTCGAGCTAGCGTGTAATCCTGGT
	ATAGGTCGTACTATTGAACTTATGCCCTAAAGTATTCCATTATCGCCGGGACAGTATGAA
	GCCCGCCCACTGTACGTCATTTCGCTGAGTCGTACTGTCTGTGCTGGGCTCTCCGTACGT
	GTTCATAAAGTTCAAGGATGATGTATCCCTATACAACGCTCAGGACGTGGTCCAAGAAAA
	GCCATACCTATAAACTTACCCAAAAGTTGATGTATCAGAAACCTCCAGGTGGCTAACTTA
	GGACCGGGCACGCGCCTTGGGACGTGTACGGCCTGGCTCGATCTGGTTAACTAAGCAGTC
	ATCTATTTAAGTTAGGTTAGATCAACGCTCCTTTTAAGGGTAGCCATACTCCGTCGGGCT
	CAAGCGTTACCTCCAGATAAAGAGTTAGTATACAAGCAAACACGTCCGAGAGCCCTGTTC
	GCCCCAAAGATAGCTATCTTCCTGAAAGACATCCTCTCTCATATCCAGTTACATGCGCTC
	ACGAACCCGAACATTGTCCATCTCACGCAGAGCCGTATGTATTGGCTGTGAGACCCATTA
	GATAAAGCTCTTTAGTGACAGGCGGGCCGCCCTAAGCCATAGTAATATGGGACGCCAAGG
	ACCAACTGCGCTCATGAAATGCTTATGGCTGGGACGGTGGTAATGGCGGACTATCTTACT
	TGGGCCTTGTTCTTGAATCTAGATAGGACTACAAGTCGTAGGTCAAGGGAGAACAAATGC
	GTCGGTCCGAAAGTAGACCCTCAGTGGGGTAGGGAGTCTTAACGCCGCTGGGGCGACAGG
	TGGAGCGTAAAATTTTTAGGAGACGGGACATTAGTTTGGTCCTGTCGTACGGGCCTCGAA
	AAGTTGTTCGGTTACATGTGTGGTCATCTGGTTGTCCCTTGCAATACGTGTTGGCAAGTT
	ATCAAACGGTGCATCATAGTACTTAAACAGCTTCTC>Rosalind_9454
	CCGGAGAAACAGGTTCCACCCGAGTGCACTATTCTGTTCGTAAATATGAACAGCTTCGGA
	GAACTTCGGATACATGGAATGGGACTCGTAGTAAAACATACTTTTAGCCTCAGAAGAGGC
	GACTTGGGCCAACTAGGGCCCAAGTTTAAAGCATTCCAGAATTCAAGAAAGTTGTCTGCT
	TTTACCGCTAATTTCACCCGTTATTGGAAGTTTAATTCCTAGGGAACCTTTGTGAACATG
	ACAAGGTGGCACTAGCATCGAGCCGTGCTCACTTTGGCGTGGGTGTAGAACGCGAATGGG
	AACCGATCTTACCAGGGGCATGCTCAATCGATGTCATACGTTAGTCACCCATCCGATACA
	TTGGATACTTTTAGGCAAAGTAGGCCTATACTTACCGCAACGTTCGGGATACTAACTCCA
	ACCACACTGCATATGTCTAGCTCCCGACACGGAGTAGCCATTAACTCACCCCGTCAACAG
	TCCTTCCTTATGAGAGCGTCTGATAACTAGCAGCTTTGGCGTGATCGGATAGGCCGAAGA
	ACGGCTCTGCCTCGTTATTGCTTGGATTGTTTACGCTCACCGTACAGGCATCCGCATAGG
	AACTTGTACTAGGATCGCTTTGCCATCCGGGTGTTGGGGCTCACCAGAAAGCACGACTTA
	TTGCACGACCCCATGTTTTCTGGGTGCGTCGTCAGAGATTGCTTTGAGTGCTTAAACCTC
	AATCACGTAGCTCCTTAAATGGGCTAAAAATTGCAAAAGGTCTTTCCGCTGCTCGAACCA
	TTAATAACAGTCCTAGGGCTTACTCGTGCCGTTATCTGGTTGGCGTCGGCGTCTATACCT
	TCTGACTGTATATTTAGGCCT
	>Rosalind_0200
	TGGACCCCTGCGAATCTAAGAAGAAACGGGATCCTAGAGACGTTTTAATGTTAGCAGGGA
	GCAAACTCCTTAACTTTTGAGTCGTGTTTTCAGGGGTATCTGAAGTGGGAGACCAGTTCT
	GGACGCCGTGTTGGCGACACGTCGTAACCGCCGCAACTGGTGCTCCGTGCTTCCTGAGTT
	TTGTATAGGGGATATCTCGCAAATGTCAAGGACTCAACTTCCCTCCCAACGACCCTATGG
	AGTACATCACAGCGTATCAGACCGAACTCAGAAGTTGCCTGACGATACAGATATTAAAAT
	TTATCCTACTAATTCAGCCCCGTCCCAGCAAAGGAGTAATCTCAGTAAATTTATAGTTTT
	ACGATCTCAATATAGCCATTGTACGTGTCTTCAATTGACATATGATTTCAGTAAGGTCTA
	GCTCGTTATAATCTTTGCGTAGAGCGCTGTTAATGTCTGAGACATGCGTACGCAATTTCA
	TTAAAGTCGGGAGCGCTCAATACGACAGCCTTTCAAAAGGCGCCAGAGTATGACGATTGC
	CGACCTACGCGCCGCATGTTCAACACATCAAATTAGGCCCCATCGATAACGATGGTGGTT
	ATCATGTCCATTGCAGCTGACACGCTGGTGCAACAATGCTGTATCCATAACACAGTTGAC
	CGAGTTAGGGCTCCGAGTCAGTGTAGAATCGCAGGGGCGAAGGACAACTTTGCAGTCATC
	AGCTAGCCGTAACTGGAATATGCATGGGGTGGGGTAGTCGTGACTTGCACCGGTGCTAAA
	TCCCCATTTCCCCGAGAACAGCCCGGTTGAAAAGGGAGATGGGCGGAATAGCACAACTTT
	CGTGTATGATTGGCATAAGCGGGAAACTGTATATCGAGAACGCATCGAGTAGCGCTTTCG
	GAATACTCATGCTGGTCAACCCAGTTCACTGTGCACTTATGCGGTCTG
	>Rosalind_1410
	GATTGGATAAAGCTCTCACCTGGTCCTTGAGATTCGAAGGGCGGGGAGTCGAACCCAGGC
	AAAATTTAGCTTTTGAGTAGGCTAGTTCTCTTGCAGTCCTACTACCTTGGAACGCAGCTT
	CAGCAGCGAAGTCTATGACATCAGTGACCGTGTACGGATATGCTGTGGGCCAACTGTAAT
	AGAATCATAGCGGACGACGATGACACTGGCCTGCCCACCATAGGGCATCGTGACGGTGCC
	CCTAGGCACGCGTTCTTCCTCGGCTAGTTAATCTCTTATGAGTGAGCAACGATGACGTCA
	ATTTTCCAACTAGGTTCTAAAAAAGGACTTGTACTGCGGGTCTGTTCATTTTTTGAAGAC
	GCGACGCTATGGTACGCAACTCGGTCGACGCGATAATATCTTTTTCGCACTCAGCCCGTA
	GTTCGGGTCCCCGTGGATCAGTCTTCCCAGATCGCCAGCCATAACACTGCCGAGATATGC
	CAATATGAACTTCCGTTACCCCGCCCCTTTCTCCGATTATCTACACGTCGGAGCTGGCGC
	AGTATCTCGGAAATTCTAGGGAGACTAGAACACTTAGCCTCAGCGATTCAAATTTGACGC
	ATAGCATAACGTTGTCACGAACTACTGGCCTATGGGCCCAGTTTAACGCTTCTGCGACGG
	ATCTAATCTTAACCGCAGCGCTGCTTATTACCCTCTAGACTCGCTGACATGAGGCGCAAT
	CTCATTGTCATTAGATCTCCCTCCCCCTAGGTCTTATTTCTTTCGATGTTAGTATCCGCT
	CTCACATCCTTCTACCAGGTGTTAAAGTGAGTCGCCCGTCCAACCCCCACGCAAGAATTG
	CCCTGACAAGTGTGAACTATTAGCGTCACCGGA
	>Rosalind_0695
	GTAAGAACAAAGAACAGTTGAAGTCTCGGACTCGCGGAACGCGGGCATTAAGGGCATGGA
	GGCGCGCAGAGCATAACTTCGTTAAGCCTGCCCCTTATTAGATCTTTCGTCCGTGGTTGG
	GTGAGGCTTGCTATTGCGACGTATAGCGGCTTGGCGGCGGGAAATTCGCTACTCAGAATT
	GTAGAGTTTTCGCTCAGGCGTTCTCATTCTGGCATAAGCCCCACCTGAATTTCAGGCGCT
	ATGATCTACAATCTGCCGCGAACGTGATCGGCAATGTCGACCGGTCAGTGTTCGTCAACG
	TTTACATAATCCTGATTAAAAATCGCTACTCGGCATTGCGGTATACGGAACGGAGTCTCA
	ACCAAGTGCTGAACAACGTGAGACACCTTAGAATTTTCCGTTCACCGTCGTCTGCAAGGA
	AGATGCAGAGCTAATATCATACATTAGGGTTTGTGCGCTTCCAAGAAACCAATGACCGTT
	TCGCGGTACTGCATGGAAGCCAGCCCTCCCACACCGAATTTCTGGTATTTTCGGATGCTG
	ATTACTGCAAACCAGACTCAAATCATTATCCCTCCAAGACACTTCAACACGGTCCGCTGT
	CATCGTGTACTCAGCGTCGACCTGTTCTTGGACGGTGCGCGTACCAGTTGGGTGTAAGCA
	CAATTCGGACTGCCGTAGCAGCATAATGCCCGGCCTATAACGAACAATAGAAAAGTTCGC
	CGGGCGTAGGACTAGCACCAGAGGTTCAAACTAAGTCCCAAAGGCCCCGGCGTGAAACGC
	GTACTGTACGCCAAGAGCGAAGGGGATTTAGACCCACAAATGGGTTGTCACCCGAAGGAG
	GGCCACATGTATATAACATGGTGTGGTACACCTGCTTGCTTACTCATCCTATGCCTGCAA
	CTATATAGGTCCGAA
	>Rosalind_8077
	AGCCAAGGGTGCCGGGGTGCTACGACTGGAAACAAAGTCGAGCTAGCGTGTAATCCTGGT
	ATAGGTCGTACTATTGAACTTATGCCCTAAAGTATTCCATTATCGCCGGGACAGTATGAA
	GCCCGCCCACTGTACGTCATTTCGCTGAGTCGTACTGTCTGTGCTGGGCTCTCCGTACGT
	GTTCATAAAGTTCAAGGATGATGTATCCCTATACAACGCTCAGGACGTGGTCCAAGAAAA
	GCCATACCTATAAACTTACCCAAAAGTTGATGTATCAGAAACCTCCAGGTGGCTAACTTA
	GGACCGGGCACGCGCCTTGGGACGTGTACGGCCTGGCTCGATCTGGTTAACTAAGCAGTC
	ATCTATTTAAGTTAGGTTAGATCAACGCTCCTTTTAAGGGTAGCCATACTCCGTCGGGCT
	CAAGCGTTACCTCCAGATAAAGAGTTAGTATACAAGCAAACACGTCCGAGAGCCCTGTTC
	GCCCCAAAGATAGCTATCTTCCTGAAAGACATCCTCTCTCATATCCAGTTACATGCGCTC
	ACGAACCCGAACATTGTCCATCTCACGCAGAGCCGTATGTATTGGCTGTGAGACCCATTA
	GATAAAGCTCTTTAGTGACAGGCGGGCCGCCCTAAGCCATAGTAATATGGGACGCCAAGG
	ACCAACTGCGCTCATGAAATGCTTATGGCTGGGACGGTGGTAATGGCGGACTATCTTACT
	TGGGCCTTGTTCTTGAATCTAGATAGGACTACAAGTCGTAGGTCAAGGGAGAACAAATGC
	GTCGGTCCGAAAGTAGACCCTCAGTGGGGTAGGGAGTCTTAACGCCGCTGGGGCGACAGG
	TGGAGCGTAAAATTTTTAGGAGACGGGACATTAGTTTGGTCCTGTCGTACGGGCCTCGAA
	AAGTTGTTCGGTTACATGTGTGGTCATCTGGTTGTCCCTTGCAATACGTGTTGGCAAGTT
	ATCAAACGGTGCATCATAGTACTTAAACAGCTTCTC`)

	countGC	float64	= 0
	count	float64 = 0
	maxGC	float64 = 0
	prcntGC float64 = 0
	crrntID = ``
	prvID   = ``
	maxID   = ``
)

func baseCount (frgmt string) {

	crrntID = frgmt[1:14]

	for j := range frgmt {
		if j > 0 {
			if frgmt[j] == '>' {
				prcntGC = 100 * countGC/count
				prvID   = crrntID
				crrntID = frgmt[j+1:j+14]
				count   = 0
				countGC = 0
				j = j + 15

			}else if frgmt[j] == 'G' || frgmt[j] == 'C' {
				countGC++
				count++

			}else if frgmt[j] == 'A' || frgmt[j] == 'T' {
				count++
			}

			if maxGC < prcntGC || maxGC == 0 {
				maxGC = prcntGC
				maxID = prvID
			}
		}

	}

	prcntGC = 100 * countGC/count

	if maxGC < prcntGC {
		maxGC = prcntGC
		maxID = crrntID
	}
}

func main() {
	baseCount(fasta)
	fmt.Println(maxID)
	fmt.Printf("%.6f", maxGC)
	fmt.Println("")
}

