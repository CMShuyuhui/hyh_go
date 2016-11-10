package hello

import (
	"errors"
	"fmt"
	"os"
)

var (
	v1 int    = 123
	v2 string = "你好，我是胡宇辉"
	v3 [10]int
	v4 []int
	v5 struct {
		f int
	}
	v6 *int
	v7 map[string]int
	v8 func(a int) int
)

const pi float64 = 3014159265358979323846
const zero = 0.0
const (
	size int64 = 1024
	eof        = 1
)

const (
	a = 1 << iota
	b = 1 << iota
	c = 1 << iota
)

type PersonInfo struct {
	Id      string
	Name    string
	Address string
}

// func main() {
// 	Test_hello()
// }

func Test_hello() {

	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["1"] = PersonInfo{"1", "Huyuhui", "ShenZhen"}
	personDB["2"] = PersonInfo{"2", "Huqingqing", "HengYang"}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("Id:", person.Id, " ,Name:", person.Name, "Address:", person.Address)
	} else {
		fmt.Println("cannot find the person id ")
	}

	before := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("before", before)
	for i, j := 0, len(before)-1; i < j; i, j = i+1, j-1 {
		before[i], before[j] = before[j], before[i]
	}
	fmt.Println("after", before)

	result, err := add(5, 10)
	if err == nil {
		fmt.Println("add result = ", result)
	} else {
		fmt.Println("error: ", err)
	}

	myfunc(1, 2, 3, 4)
	MyPrintf(v1)
	MyPrintf(v2)
	MyPrintf(v3)

	myclosefun := func() func() {
		return func() {
			fmt.Println("akjdflajdlkf")
		}
	}()
	myclosefun()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover")
		}
	}()
	var err_test myPathError = myPathError{"test_string ", errors.New("errors.new()")}
	fmt.Println(err_test.Error())

	defer func() {
		// do the things you want before main return
	}()
	// var myarray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println("")

	// for _, v := range myarray {
	// 	fmt.Print(v, " ")
	// }
	// fmt.Println("")

	// var myslice []int = myarray[:5]
	// for _, v := range myslice {
	// 	fmt.Print(v, " ")
	// }
	// fmt.Println("")

	// var slice2 = make([]int, 5)
	// for i := 0; i < len(slice2); i++ {
	// 	fmt.Printf("hello world %d ：%s\n", v1, v2)
	// 	firstname, lastname, nickname := GetName()
	// 	fmt.Printf("firstname: %s, lastname: %s, nickname: %s\n", firstname, lastname, nickname)
	// 	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
	// 	fmt.Println("firstname: ", firstname, "lastname: ", lastname, " nickname: ", nickname)
	// }
	writeValuestoFile("log.txt", "hello world", "你好，我是胡宇辉", "1827398217947")
}

func writeValuestoFile(outfile string, values ...interface{}) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		file.WriteString(value.(string) + "\n") //通过interface{}.(string) 转化成string类型
	}
	return nil
}

type myPathError struct {
	err string
	Err error
}

func (e *myPathError) Error() string {
	return e.err + " Error: " + e.Err.Error()
}

func add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers")
		return
	}
	return a + b, nil
}

func GetName() (firstname, lastname, nickname string) {
	return "Hu", "Yuhui", "maomao"
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is int")
		case string:
			fmt.Println(arg, "is string")
		default:
			fmt.Println(arg, " unkowned type")
		}
	}
}
