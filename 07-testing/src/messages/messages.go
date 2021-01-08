package messages

type Value int64

// var someValue int
func (v *Value) SumTen() int64 {
	return int64(*v) + 10
}

// func main() {
// 	var v := 100
// 	fmt.Printf("Hello Marcos %d", v.sumTen())
// }
