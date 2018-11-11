package main

import "fmt"

type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s *Stack) pop() {
	//poppedValue := *s[len(*s)-1]



	//*s = *s[0, 1]


	//return poppedValue
}

func (s *Stack) size() int {
	return len(*s)
}

func main()  {

	var st = Stack([]int{})

	st.push(15)
	st.push(32)

	fmt.Println(st)
	fmt.Println(st.size())


	//fmt.Println(st[0,1])
	//fmt.Println(st.pop())

	//fmt.Println(st)
	//fmt.Println(st.size())





	//fmt.Println(st)
	//st[len(st)-1] = 0
	//fmt.Println(st[len(st)-1])
	//fmt.Println(st)

	//fmt.Println(st.pop())
	//fmt.Println(st.pop())

}
