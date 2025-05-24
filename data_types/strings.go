package main
import (
	"fmt"
	"strings"
)

func main(){
	s:="I am string"
	fmt.Printf("s=%v\n",s)
	fmt.Printf("End with letter 'string' = %v\n",strings.HasSuffix(s,"string"))



}

