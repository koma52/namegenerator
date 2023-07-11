package ui

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var vowels []string
var bases []string
var onset []string
var vowelString string
var inter []string
var rhyme []string
var tmpBases []string

func readVowels(f []byte) []string {
	fileContents := string(f)
	fileContents = strings.ToLower(fileContents)
	fileContents = strings.Replace(fileContents, " ", "", -1)
	fileContents = strings.Replace(fileContents, "\t", "", -1)
	fileContents = strings.Replace(fileContents, "\n", "", -1)
	fileContents = strings.Replace(fileContents, "\r", "", -1)

	var inVowels []string

	for i := 0; i < len(fileContents); i++ {
		if fileContents[i] != ' ' && !contains(inVowels, string(fileContents[i])) {
			inVowels = append(inVowels, string(fileContents[i]))
		}
	}

	return inVowels
}

func readBases(f []byte) string {
	fileContents := string(f)
	fileContents = strings.Trim(fileContents, "\n")

	return fileContents
}

func getVowels(c *MainWindow) {
	vowels = make([]string, 0)
	fromInput := c.Vowels.Text()
	for i := 0; i < len(fromInput); i++ {
		vowels = append(vowels, string(fromInput[i]))
	}
	// fmt.Println(vowels)
}

func getBases(c *MainWindow) {
	bases = make([]string, 0)
	fromInput := c.Input.ToPlainText()
	fromInput = strings.Trim(fromInput, "\n")
	bases = strings.Split(fromInput, "\n")
	for i, str := range bases {
		bases[i] = strings.ToUpper(str)
	}
	// fmt.Println(bases)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func findRhyme(s string) (string, string) {
	var substrings []string

	start := -1
	for i, r := range s {
		if strings.ContainsRune(vowelString, r) {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				substrings = append(substrings, s[start:])
				start = -1
			}
		}
	}

	if start != -1 {
		substrings = append(substrings, s[start:])
	}

	if len(substrings) == 0 {
		return s, ""
	}

	lastSubstring := substrings[len(substrings)-1]
	firstSubstring := strings.TrimRight(s, lastSubstring)

	return lastSubstring, firstSubstring
}

func findOnset(s string) (string, string) {
	var currentOnset string
	var afterOnset string
	for i, c := range s {
		if strings.ContainsRune(vowelString, c) {
			if i == 0 {
				currentOnset, afterOnset = "", s
				break
			} else {
				currentOnset, afterOnset = string(s)[0:i], string(s)[i:]
				break
			}
		} else {
			currentOnset, afterOnset = s, ""
		}
	}

	return currentOnset, afterOnset
}

func generateNames(c *MainWindow) {

	getBases(c)
	getVowels(c)

	onset = make([]string, 0)
	inter = make([]string, 0)
	rhyme = make([]string, 0)

	vowelString = strings.Join(vowels, "")
	vowelString = strings.ToUpper(vowelString)

	for i := 0; i < len(bases); i++ {
		r, f := findRhyme(bases[i])
		rhyme = append(rhyme, r)
		tmpBases = append(tmpBases, f)
	}

	// fmt.Println("tmpBases after rhyme: ", tmpBases)

	for i, v := range tmpBases {
		o, a := findOnset(v)
		onset = append(onset, o)
		tmpBases[i] = a
	}

	// fmt.Println("tmpBases after onset: ", tmpBases)

	for i, v := range tmpBases {
		tmp := v
		var ir string
		for {
			if tmp != "" {
				ir, tmp = findRhyme(tmp)
				inter = append(inter, ir)
			} else {
				break
			}
		}
		tmpBases[i] = ""
	}

	// fmt.Println("tmpBases after inter: ", tmpBases)

	// fmt.Println("Onset: ", onset)
	// fmt.Println("Inter: ", inter)
	// fmt.Println("Rhyme: ", rhyme)
	tmpBases = make([]string, 0)

	nameNumber, _ := strconv.Atoi(c.NumberOfGen.Text())
	syllableNumber, _ := strconv.Atoi(c.NumberOfSyl.Text())
	var names []string

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nameNumber; i++ {
		var tmpName string
		randomIndex := rand.Intn(len(onset))
		tmpName += onset[randomIndex]

		for j := 0; j < syllableNumber-1; j++ {
			randomIndex = rand.Intn(len(inter))
			tmpName += inter[randomIndex]
		}

		randomIndex = rand.Intn(len(rhyme))
		tmpName += rhyme[randomIndex]

		names = append(names, tmpName)
	}

	c.Output.SetPlainText("")

	for _, v := range names {
		c.Output.AppendPlainText(v)
	}
}
