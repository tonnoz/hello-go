package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
) //import are automatic from Intellij

func main(){

	/*
	Any variable that begins with a capital letter means it will be exported, private otherwise.
	The same rule applies for functions and constants, no public or private keyword exists in Go.
	*/
	var ApublicVar = 1
	fmt.Println(ApublicVar)

	//1. Hello world
	var hello = "Hello World"
	fmt.Println(hello)

	//2. short assignment with type inference
	hello2 := "Hello World2"
	fmt.Println(hello2)

	//3. fixed size array
	var arr = [5]int {1,2,3,4,5}
	fmt.Println(arr)

	//4. slices: an array with no fixed size
	var slice = []int {1,2,3,4,5}
	fmt.Println(slice)
	var slice2 = slice[2:4] //create a slice from slice[2..4] n.b. index 4 excluded
	fmt.Println(slice2)
	var slice3 = slice[:4] //create a slice from slice[0..4] n.b. index 4 excluded
	fmt.Println(slice3)
	var slice4 = slice[2:] //create a slice from slice[2..n]
	fmt.Println(slice4)
	var slice5 = make([]int, 5, 10) // define a slice of int with default value 5 and max 10 elements
	fmt.Println(slice5)

	newSlice := append(slice, 6,7,8,9,10) //append
	fmt.Println(cap(slice)) //cap gets the max length of the slice
	/*
	Attention: append will change the array that slice points to, and affect other slices that point to the same array.
	Also, if there is not enough length for the slice ((cap-len) == 0), append returns a new array for this slice.
	When this happens, other slices pointing to the old array will not be affected.
	 */
	fmt.Println(newSlice)

	sliceb := make([]int, len(slice))
	copy(sliceb, slice) //copy slices
	fmt.Println("copied slice: ", sliceb)

	//5. map: a key value data structure
	aKeyValuePair := make(map[string]int)
	aKeyValuePair["akey"] = 1  //add
	aKeyValuePair["anotherKey"] = 2
	aKeyValuePair["thirdKey"] = 3
	fmt.Println(aKeyValuePair)
	fmt.Println(aKeyValuePair["akey"])
	delete(aKeyValuePair, "thirdKey") //delete
	fmt.Println(aKeyValuePair)
	fmt.Println(len(aKeyValuePair)) //map length


	//6. for loop
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}

	//7. while equivalent
	i:=0
	for i < 5 {
		fmt.Println(i)
		i++
	}


	//8. loop over range of slices/arrays
	slice = []int {1,2,3,4,5}
	for index, value := range slice {
		fmt.Println("index",index,"value", value)

	}

	//9. loop over maps
	aMap := make(map[string]string)
	aMap["a"] = "alpha"
	aMap["b"] = "beta"

	for key, value := range aMap {
		fmt.Println("key",key,"value", value)
	}

	//10. functions
	fmt.Println(sum(2,3))

	//11. multiple return values function: GO DOES NOT HAVE EXCEPTIONS
	result,err := sqrt(16) //error is not nil if x is negative

	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	//12. Variatic function (equivalent  of vararg in Java)
	fmt.Println(addThem(1,2,3,4,5,6))

	//13. use of struct
	p:= person{name:"Tony", age:31}
	fmt.Println(p.name)

	//14. pointers
	i=7
	fmt.Println(&i) //print the address of the var

	fmt.Println("before", i)//before
	inc(i)  //no effect since it's passed by value
	fmt.Println("after", i)//after

	fmt.Println("before", i)//before
	incReference(&i)  //no effect since it's passed by value
	fmt.Println("after", i)//after


	anIntPointer := new(int)
	changeValueOfPointer(anIntPointer)
	fmt.Println("anIntPointer= ", *anIntPointer)

	//15. closures
	num3:=3
	duplicateNumber := func() int{
		num3*=2
		return num3
	}
	fmt.Println(duplicateNumber()) //will print 6
	fmt.Println(duplicateNumber()) //will print 12

	//16. use of defer: defer the call of a function after the current function has terminated
	defer printTwo() //will be called after main has ended: useful for cleanup operation, flush operations, close resources etc..
	printOne()

	//17: defer and recover
	fmt.Println(saveDivision(2,0)) //error b but program doesnt quit
	fmt.Println(saveDivision(2,2)) //ok

	//18: defer and panic
	testPanic()

	//19: struct
	rect1 := Rectangle{width:10, height:10}
	fmt.Println(rect1.area()) //attached function area

	//20: interfaces
	rect := Rectangle{20,50}
	circle := Circle{4}
	fmt.Println("area rect", getArea(&rect))
	fmt.Println("area circle", getArea(&circle))

	//21: String manipulation

	sampleString := "Hello World"
	fmt.Println(strings.Contains(sampleString, "lo"))     //boolean in return
	fmt.Println(strings.Index(sampleString, "lo"))        //integer in return (first occurrence)
	fmt.Println(strings.Count(sampleString, "l"))         //integer in return (number of  occurrences)
	fmt.Println(strings.Replace(sampleString, "l","x",3)) //replace all l with x (only first 3 occurrences)
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ",")) //return a list (or array?)

	listOfLetters :=[]string{"c","b","d","a"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters: ", listOfLetters)

	listOfNums := strings.Join([]string{"hello", "world"},"! ")
	fmt.Println(listOfNums)


	//22: files IO
	//writing
	file,err := os.Create("sample.txt")
	if err!=nil{log.Fatal(err)}
	file.WriteString("Hello World!")
	file.Close()

	//reading
	stream, err2 := ioutil.ReadFile("sample.txt")
	if err2!=nil{log.Fatal(err)}
	readString := string(stream)
	fmt.Println(readString)

	//23: numbers conversion
	randInt:=5
	randFloat:=10.5
	randString:="100"
	randString2:="250.2"

	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

	//24: goroutines
	for i:=0; i < 10; i++{
		go count(i)
	}
	time.Sleep(time.Second *5)


	//25: channels (pass data to go routines)
	stringChan := make(chan string)
	for i :=0; i<3; i++{
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond*500)
	}
	time.Sleep(time.Second * 8)

	//26: HTTP WEB SERVER!
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)


}

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChannel chan string){
	pizzaNum++
	pizzaName = "pizza #"+strconv.Itoa(pizzaNum)
	fmt.Println("Make Dough for ", pizzaName, "and Send for Sauce")
	stringChannel <- pizzaName //send value to the channel
	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChannel chan string){
	pizza:= <- stringChannel
	fmt.Println("Add Sauce and Send ", pizza, " for toppings")
	stringChannel <- pizzaName //send value to the channel
	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChannel chan string){
	pizza:= <- stringChannel
	fmt.Println("Add toppings to ", pizza, " and ship")
	time.Sleep(time.Millisecond * 10)
}

func count(id int){
	for i:=0; i < 10; i++{
		fmt.Println("routine ID: ", id, ":", "Value: ", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Go!\n")
}
func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello Earth!\n")
}


func changeValueOfPointer(anIntPointer *int) {
	*anIntPointer = 100
}


func inc(i int) {
	i++
}

func incReference(i *int) {
	*i++
}
//simple function
func sum(x int, y int) int {
	return x+y
}

//multiple return value function
func sqrt(x float64) (float64, error){
	if x<0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

//struct
type person struct {
	name string
	age int
}

//variatic function
func addThem(args ...int) int{
	finalValue :=0
	for _, value := range args{
		finalValue += value
	}
	return finalValue
}

func printOne(){fmt.Println(1)}
func printTwo(){fmt.Println(2)}

//this function will execute the division but if an error will occur (eg num/0) then the program wont terminate because of recover
func saveDivision(num1,num2 int) int{
	defer func(){ //defer closure
		fmt.Println(recover()) //remove println to not print the error
	}()
	solution := num1 / num2
	return solution
}

//panic is a sort of throw exception: interrupt normal flow, call the defer if present and then quit the function
func testPanic(){
	defer func(){
		fmt.Println(recover())
	}()
	panic("an error occurred, I am panicking")
}

type Shape interface {
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}


type Circle struct {
	radius float64
}

func getArea(shape Shape) float64{
	return shape.area()
}

func (rect *Rectangle) area() float64 { //method: a function attached to a type
	return rect.height * rect.width
}

func (c *Circle) area() float64 {  //method: a function attached to a type
	return math.Pow(c.radius,2) * math.Pi
}

