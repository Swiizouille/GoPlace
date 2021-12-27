//Solution proposée suite à la formation "Apprendre Go" - Udemy {Swiique, Gaëtan}
package main

import (
	"fmt"
	"strings"
)

//Fonction pour trouver les mots définie dans la fonction main puis retourner leur ligne et là ils ont été remplacé
func ProcessLine(line, old, new string) (result bool, res string, occ int) {
	oldWord := strings.ToLower(old)
	newWord := strings.ToLower(new)
	res = line

	if strings.Contains(line, old) && strings.Contains(line, oldWord) {
		result = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldWord)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(line, oldWord, newWord, -1)

	}
	return result, res, occ
}

/* func FindReplaceFile(src string, dst string, old string, new string) (occ int, lines []int, err error) {

} */

func main() {
	old := "Java"
	new := "Go"

	fmt.Printf("J'ai trouvé %v, que j'ai remplacé par %v", old, new)

}
