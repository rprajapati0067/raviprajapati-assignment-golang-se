package main
import (
	"fmt"
)

type Stack []rune

func (s *Stack) IsEmpty() bool{
	return len(*s) == 0
}

func (s *Stack) Push(char rune) bool{
	*s = append(*s, char)
}

func (s *Stack) Pop() (rune, bool){
	if s.IsEmpty(){
		return "",false
	}else{
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) IsBalanced(str string) string{
	if len(str) % 2 != 0{
		return "NO"
	}
	for i:=0; 
}

func main(){


}