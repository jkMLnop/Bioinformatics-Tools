
package main

import(
	"fmt"
	"bytes"
)

var(
	codons = map[string]string{	"UUA":`L`,"UUG":`L`,"CUU":`L`,"CUC":`L`,"CUA":`L`,"CUG":`L`,
					"UCU":`S`,"UCC":`S`,"UCA":`S`,"UCG":`S`,"AGU":`S`,"AGC":`S`,
					"CGU":`R`,"CGC":`R`,"CGA":`R`,"CGG":`R`,"AGA":`R`,"AGG":`R`,
					"GGU":`G`,"GGC":`G`,"GGA":`G`,"GGG":`G`,
					"GUU":`V`,"GUC":`V`,"GUA":`V`,"GUG":`V`,
					"CCU":`P`,"CCC":`P`,"CCA":`P`,"CCG":`P`,
					"ACU":`T`,"ACC":`T`,"ACA":`T`,"ACG":`T`,
					"GCU":`A`,"GCC":`A`,"GCA":`A`,"GCG":`A`,
					"AUU":`I`,"AUC":`I`,"AUA":`I`,
					"UUU":`F`,"UUC":`F`,
					"UAU":`Y`,"UAC":`Y`,
					"CAU":`H`,"CAC":`H`,
					"CAA":`Q`,"CAG":`Q`,
					"AAU":`N`,"AAC":`N`,
					"AAA":`K`,"AAG":`K`,
					"GAU":`D`,"GAC":`D`,
					"GAA":`E`,"GAG":`E`,
					"UGU":`C`,"UGC":`C`,
					"AUG":`M`,
					"UGG":`W`}

	stops = map[string]string{"UAA":`*`,"UAG":`*`,"UGA":`*`}	/*should use a simple array for this perhaps?*/

	rnastring = (`AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA`)

	protstring bytes.Buffer
	orf = false
)

func translateProt (rna string){
	for i := 0; i <= (len(rna) -3); i++ {

		/*find a start codon and begin translation..*/
		if rna[i:i+3] == `AUG` {
			protstring.WriteString(codons[`AUG`])
			orf = true
			i = i + 3
		}

		if orf == true {
			/*check for stop codons - if true break*/
			if _, exists := stops[rna[i:i+3]]; exists{
				break

			}else if _, exists := codons[rna[i:i+3]]; exists{
				protstring.WriteString(codons[rna[i:i+3]])
				i = i + 2	/*iterate to one position before next codon (for loop iterates by one too!)*/

			}else {
				fmt.Println("Invalid codon encountered, process ended")
				break
			}
		}
	}
}

func main() {
	translateProt(rnastring)
	fmt.Println(protstring.String())

}
