# Go Learning Roadmap

## **1: Go Setup & Syntax Basics**

### **Introduction to Go**

Go (also called Golang) is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and concurrency. Go is known for its fast compilation, garbage collection, and built-in support for concurrent programming.

### **Setting Up Go**

Before writing Go code, you need to set up your environment correctly.

1. **Download & Install Go**: Download the official Go distribution from [https://go.dev/dl/](https://go.dev/dl/) and install it according to your operating system.
2. **Verify Installation**: Open a terminal or command prompt and run:
   ```sh
   go version
   ```
   This should display the installed Go version, confirming a successful installation.
3. **Set Environment Variables**:
   - `$GOROOT`: The directory where Go is installed.
   - `$GOPATH`: Your workspace where Go projects are stored.
   - `$PATH`: Add `$GOPATH/bin` to system paths to access Go binaries globally.
4. **Running Go Code**:
   - **Run directly**:  `go run main.go` (Executes without compiling a separate binary)
   - **Compile & Run**:  `go build main.go && ./main` (Creates an executable binary)
   - **Format Code**:  `go fmt main.go` (Auto-formats code to maintain Go’s style guidelines)

### **Go Syntax Basics**

Go has a simple syntax with a focus on readability and performance. Let’s cover some core concepts:

#### **Variables & Constants**

Variables store data and can be changed, while constants store fixed values.
```go
package main
import "fmt"

func main() {
    var name string = "Alice"
    age := 25 // Short-hand declaration
    const pi = 3.14
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Pi:", pi)
}
```

#### **Data Types**
Go is a statically typed language, meaning variables must have a specific type:
- `int`: Whole numbers
- `float64`: Decimal numbers
- `bool`: Boolean values (`true` or `false`)
- `string`: Text values

#### **Control Flow Statements**

##### **Conditional Statements (if-else)**
```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

##### **Looping with `for`**

###### **Basic `for` loop**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

###### **For loop as a while loop**
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

###### **Infinite loop**
```go
for {
    fmt.Println("This will run forever")
}
```

###### **Looping through a slice or array using `range`**
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

##### **Switch Case**
```go
package main
import "fmt"

func main() {
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week!")
    case "Friday":
        fmt.Println("Weekend is near!")
    default:
        fmt.Println("A regular day")
    }
}
```

### **Error Handling in Go**

Go does not use traditional exception handling like other languages. Instead, errors are treated as values that can be returned and handled explicitly.

```go
package main
import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

### **Practice Task**
#### **Build a CLI Calculator**
Create a simple command-line calculator in Go that can perform addition, subtraction, multiplication, and division. The program should:
1. Ask the user for two numbers.
2. Ask the user for an operation (+, -, *, /).
3. Perform the operation and display the result.
4. Handle division by zero gracefully.

```go
package main
import (
    "fmt"
    "errors"
)

func calculate(a float64, b float64, operator string) (float64, error) {
    switch operator {
    case "+":
        return a + b, nil
    case "-":
        return a - b, nil
    case "*":
        return a * b, nil
    case "/":
        if b == 0 {
            return 0, errors.New("division by zero is not allowed")
        }
        return a / b, nil
    default:
        return 0, errors.New("invalid operator")
    }
}

func main() {
    var a, b float64
    var operator string
    
    fmt.Print("Enter first number: ")
    fmt.Scanln(&a)
    
    fmt.Print("Enter an operator (+, -, *, /): ")
    fmt.Scanln(&operator)
    
    fmt.Print("Enter second number: ")
    fmt.Scanln(&b)
    
    result, err := calculate(a, b, operator)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
}
```

### **Conclusion**
1 covered the fundamental setup and syntax of Go, including variables, constants, data types, control flow, and error handling. The hands-on task provided a practical way to apply these concepts. Moving forward, we will explore functions, packages, and data structures in Go.



## ** 2: Functions, Packages, and Data Structures**

### **Introduction**

Functions, packages, and data structures are fundamental to Go programming. Functions help in organizing code into reusable blocks, packages allow modular programming, and data structures like arrays, slices, and maps provide efficient ways to store and manipulate data.

### **Functions in Go**

Functions in Go help break down a program into smaller, reusable pieces. Go functions have the following characteristics:

- They take zero or more parameters.
- They can return multiple values.
- They can have named return values.
- They support variadic arguments.

#### **Defining and Calling Functions**

```go
package main
import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("Sum:", result)
}
```

#### **Multiple Return Values**

```go
package main
import "fmt"

func divide(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Cannot divide by zero"
    }
    return a / b, "Success"
}

func main() {
    result, msg := divide(10, 2)
    fmt.Println("Result:", result, "Message:", msg)
}
```

#### **Named Return Values**

```go
package main
import "fmt"

func rectangleDimensions(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return // Named return values eliminate the need to specify them again
}

func main() {
    area, perimeter := rectangleDimensions(10, 5)
    fmt.Println("Area:", area, "Perimeter:", perimeter)
}
```

#### **Variadic Functions**

A variadic function can accept multiple arguments of the same type.

```go
package main
import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}
```

---

### **Packages in Go**

Go organizes code into packages. Every Go program starts execution from the `main` package.

#### **Importing Packages**

Go provides built-in packages like `fmt`, `math`, and `time`.

```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Square Root of 16:", math.Sqrt(16))
}
```

#### **Creating Custom Packages**

1. Create a folder and a new file for the package:

```
myapp/
├── main.go
├── mathutils/
│   ├── mathutils.go
```

2. Define the package in `mathutils.go`:

```go
package mathutils

func Multiply(a, b int) int {
    return a * b
}
```

3. Use it in `main.go`:

```go
package main
import (
    "fmt"
    "myapp/mathutils"
)

func main() {
    result := mathutils.Multiply(3, 4)
    fmt.Println("Multiplication Result:", result)
}
```

Compile and run:
```
go run main.go
```

---

### **Data Structures in Go**

#### **Arrays**

Arrays have a fixed size and store elements of the same type.

```go
package main
import "fmt"

func main() {
    var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(numbers)
}
```

#### **Slices**

Slices are dynamic and more flexible than arrays.

```go
package main
import "fmt"

func main() {
    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4, 5)
    fmt.Println(numbers)
}
```

#### **Maps**

Maps store key-value pairs and are similar to dictionaries in Python.

```go
package main
import "fmt"

func main() {
    person := map[string]string{
        "name": "Alice",
        "city": "New York",
    }
    fmt.Println(person["name"], "lives in", person["city"])
}
```

---

### **Practice Task**

#### **Build a CSV Processor**

Create a Go program that reads a CSV file, calculates summary statistics, and displays results.

1. Read a CSV file.
2. Process numeric data.
3. Calculate total and average values.

```go
package main
import (
    "encoding/csv"
    "fmt"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading CSV:", err)
        return
    }

    total := 0.0
    count := 0

    for _, row := range records[1:] { // Skip header
        value, err := strconv.ParseFloat(row[1], 64)
        if err == nil {
            total += value
            count++
        }
    }

    average := total / float64(count)
    fmt.Printf("Total: %.2f, Average: %.2f\n", total, average)
}
```

---

### **Conclusion**

2 covered functions, packages, and data structures in Go. Functions allow code reuse, packages help organize code, and data structures efficiently store and manipulate data. The practice task provided hands-on experience in file processing. Next, we will explore pointers and concurrency basics.


## **3: Pointers and Concurrency Basics**

### **Introduction**

In Go, pointers and concurrency are crucial concepts. Pointers allow us to reference memory locations directly, which is useful for optimizing memory usage. Concurrency enables a program to perform multiple tasks simultaneously, improving efficiency.

---

## **Pointers in Go**

A pointer is a variable that stores the memory address of another variable. Pointers help in optimizing memory usage and modifying values efficiently.

### **Declaring and Using Pointers**

- The `&` operator is used to get the address of a variable.
- The `*` operator is used to dereference a pointer (access the value stored at a memory address).

```go
package main
import "fmt"

func main() {
    var x int = 10
    var p *int = &x // Pointer to x

    fmt.Println("Value of x:", x)
    fmt.Println("Memory address of x:", &x)
    fmt.Println("Pointer p stores:", p)
    fmt.Println("Value at pointer p:", *p) // Dereferencing
}
```

### **Passing Pointers to Functions**

Passing a pointer allows us to modify the original value inside a function.

```go
package main
import "fmt"

func updateValue(p *int) {
    *p = *p + 10
}

func main() {
    num := 20
    fmt.Println("Before update:", num)
    updateValue(&num) // Passing pointer
    fmt.Println("After update:", num)
}
```

### **Pointer vs. Value Passing**

- When passing by value, a copy of the variable is made.
- When passing by reference (using pointers), the function modifies the original value.

```go
package main
import "fmt"

func modifyValue(val int) {
    val = val * 2
}

func modifyPointer(val *int) {
    *val = *val * 2
}

func main() {
    num := 10

    modifyValue(num)
    fmt.Println("After modifyValue:", num) // Unchanged

    modifyPointer(&num)
    fmt.Println("After modifyPointer:", num) // Changed
}
```

---

## **Basic Concurrency in Go**

Go has built-in support for concurrency using Goroutines and Channels.

### **Goroutines**

A Goroutine is a lightweight thread managed by the Go runtime. It allows functions to run concurrently.

```go
package main
import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 0; i < 5; i++ {
        fmt.Println(msg)
        time.Sleep(time.Millisecond * 500) // Simulate delay
    }
}

func main() {
    go printMessage("Hello from Goroutine")
    printMessage("Hello from Main")
}
```

In the example above, `go printMessage("Hello from Goroutine")` runs the function concurrently.

### **WaitGroup for Synchronization**

A `sync.WaitGroup` helps manage multiple Goroutines by waiting for them to complete execution before the program terminates.

```go
package main
import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrement counter when done
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait() // Wait for all Goroutines to complete
    fmt.Println("All workers finished")
}
```

### **Channels**

Channels allow Goroutines to communicate safely. They help avoid race conditions.

```go
package main
import "fmt"

func sum(numbers []int, result chan int) {
    total := 0
    for _, num := range numbers {
        total += num
    }
    result <- total // Send result to channel
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    result := make(chan int)
    go sum(numbers, result)
    fmt.Println("Sum:", <-result) // Receive result
}
```

### **Buffered vs. Unbuffered Channels**

- **Unbuffered Channel**: Blocks until data is received.
- **Buffered Channel**: Allows storing multiple values before blocking.

#### **Buffered Channel Example**
```go
package main
import "fmt"

func main() {
    ch := make(chan string, 2) // Buffered channel with capacity 2
    ch <- "Message 1"
    ch <- "Message 2"
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

### **Select Statement for Multiple Channels**

The `select` statement allows a Goroutine to wait on multiple channels.

```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch1 <- "Message from ch1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Message from ch2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    }
}
```

---

### **Practice Task**

#### **Build a Concurrent Sum Calculator**

Create a Go program that calculates the sum of multiple slices concurrently using Goroutines, Channels, and a WaitGroup.

```go
package main
import (
    "fmt"
    "sync"
)

func sum(numbers []int, result chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    total := 0
    for _, num := range numbers {
        total += num
    }
    result <- total
}

func main() {
    numbers1 := []int{1, 2, 3, 4, 5}
    numbers2 := []int{6, 7, 8, 9, 10}
    
    result := make(chan int, 2)
    var wg sync.WaitGroup
    
    wg.Add(2)
    go sum(numbers1, result, &wg)
    go sum(numbers2, result, &wg)
    
    wg.Wait()
    close(result)
    
    totalSum := 0
    for res := range result {
        totalSum += res
    }
    fmt.Println("Total Sum:", totalSum)
}
```

---

### **Conclusion**

3 covered pointers and concurrency in Go. Pointers help in managing memory efficiently, while concurrency enables multiple tasks to run in parallel. The practice task provided hands-on experience in using Goroutines, Channels, and WaitGroups. Next, we will explore Structs, Methods, and Interfaces.


## ** 4: Structs, Methods, and Interfaces**

### **Introduction**

In Go, Structs, Methods, and Interfaces are essential for structuring and organizing code. Structs allow us to define custom data types, Methods help associate behavior with struct instances, and Interfaces enable polymorphism and abstraction.

---

## **Structs in Go**

A **struct** in Go is a composite data type that groups multiple fields together. Unlike classes in object-oriented languages, Go uses structs to define complex data types.

### **Defining and Using Structs**

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Creating a struct instance
    p1 := Person{Name: "Alice", Age: 25}
    fmt.Println("Person:", p1)
    
    // Accessing struct fields
    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
}
```

### **Using Struct Pointers**

When passing a struct to a function, it is often more efficient to use a pointer to avoid copying the entire struct.

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func updateAge(p *Person, newAge int) {
    p.Age = newAge
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    updateAge(&p, 30)
    fmt.Println("Updated Age:", p.Age)
}
```

---

## **Methods in Go**

Go allows attaching methods to structs, providing a way to associate behavior with data.

### **Defining Methods for Structs**

A method is a function with a **receiver**, which allows it to be called on a struct instance.

```go
package main
import "fmt"

type Rectangle struct {
    Width, Height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
}
```

### **Pointer Receiver vs Value Receiver**

- **Value Receiver**: Creates a copy of the struct, so modifications inside the method do not affect the original struct.
- **Pointer Receiver**: Modifies the original struct, preventing unnecessary copying.

#### **Example of Pointer Receiver**

```go
package main
import "fmt"

type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}

func main() {
    c := Counter{Value: 0}
    c.Increment()
    fmt.Println("Updated Value:", c.Value)
}
```

---

## **Interfaces in Go**

An **interface** defines a set of method signatures. Any type that implements those methods is said to satisfy the interface.

### **Defining and Implementing Interfaces**

```go
package main
import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

type Rectangle struct {
    Width, Height float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
}

func printArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    c := Circle{Radius: 5}
    r := Rectangle{Width: 4, Height: 6}

    printArea(c)
    printArea(r)
}
```

### **Empty Interface (`interface{}`) and Type Assertion**

The empty interface (`interface{}`) can hold any type, making it useful for generic programming.

```go
package main
import "fmt"

func describe(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
    describe(42)
    describe("Hello")
    describe(3.14)
}
```

### **Type Assertion**

Type assertions allow extracting the underlying type from an interface.

```go
package main
import "fmt"

func main() {
    var x interface{} = "Go Programming"
    str, ok := x.(string) // Type assertion
    if ok {
        fmt.Println("Extracted string:", str)
    } else {
        fmt.Println("Type assertion failed")
    }
}
```

---

## **Polymorphism in Go**

Polymorphism allows different types to be used interchangeably through interfaces.

```go
package main
import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}
type Cat struct{}

type Human struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

func (c Cat) Speak() string {
    return "Meow!"
}

func (h Human) Speak() string {
    return "Hello!"
}

func main() {
    animals := []Animal{Dog{}, Cat{}, Human{}}
    for _, a := range animals {
        fmt.Println(a.Speak())
    }
}
```

---

### **Practice Task**

#### **Build a Shape Calculator**

Write a Go program that uses structs and interfaces to define different shapes and calculate their area and perimeter.

```go
package main
import "fmt"

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 4, Height: 5},
        Circle{Radius: 3},
    }
    for _, s := range shapes {
        fmt.Println("Area:", s.Area())
        fmt.Println("Perimeter:", s.Perimeter())
    }
}
```

---

### **Conclusion**

4 covered Structs, Methods, and Interfaces, which are crucial for defining and organizing data and behavior in Go. You learned how to create and use Structs, attach Methods, and leverage Interfaces for polymorphism. Next, we will explore Error Handling and Advanced Concurrency in Go.



