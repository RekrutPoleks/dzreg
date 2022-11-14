package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {

	bfile, err := ioutil.ReadFile("in.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	solutionBuf := []byte{}

	reg := regexp.MustCompile(`(\d+)([\+\-\*\/\\])(\d+)=\?`)

	sub := reg.FindAllStringSubmatch(string(bfile), -1)
	fmt.Println(sub)

	if len(sub) > 0 {
		for _, s := range sub {
			bsolution, err := easySol(s)
			if err != nil {
				fmt.Println(err)
			}
			solutionBuf = append(solutionBuf, bsolution...)
		}
		if err := ioutil.WriteFile("solution.txt", solutionBuf, 775); err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Match not found.")
	}

}

func easySol(s []string) ([]byte, error) {
	fmt.Println(s)
	templ := "%s%s%s=%d\n"
	a, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(s[3])
	if err != nil {
		return nil, err
	}
	switch s[2] {
	case "+":
		return []byte(fmt.Sprintf(templ, s[1], s[2], s[3], a+b)), nil
	case "-":
		return []byte(fmt.Sprintf(templ, s[1], s[2], s[3], a-b)), nil
	case "/":
		return []byte(fmt.Sprintf(templ, s[1], s[2], s[3], a/b)), nil
	case "*":
		return []byte(fmt.Sprintf(templ, s[1], s[2], s[3], a*b)), nil
	default:
		return nil, errors.New("not a possible error")
	}

}
