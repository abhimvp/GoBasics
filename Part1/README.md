# Part1

- `main.go` - First Go program & to run go program we do `go run main.go`
- `go.mod` - first line in it is NAME of the module or project we can imagine & second line is current version of go that the project depends on.
  - why do we need `go.mod` file: It has some use cases ->
    - If we have a project obviously we need to name it something & which we can record in this file in the first line of the file and `as convention the name should be unique & it should have a prefix with your code_hosting_provider` ( like github, gitlab , bitbucket or anything), because a GO program or project can be of two types:
      - `Executable` - a program something we can run - our go projects will be in the category of executables - we write code and run it in our own system -> that's why it's an executable.
      - `library` or `package` - these are the code that gets imported into other people's go projects example likes `"fmt"` - a package provided by the GO's standard library which we use in our program or project. In the same way we use lot of external packages for building web servers for doing validations or for managing our configs which we use when building our projects - those are called packages - because someone has built them and have distributed it in GitHub or other hosting providers - which we can import into our project which is an executable.
    - This `go.mod` file will also have all the dependencies of the project ( external packages that this project depends on )
    - we can generate it using the command - `go mod init github.com/abhimvp/GoBasics/Part1`
    - Overtime as we add external packages into our project it will be registered in go.mod file.
      - refer : `fider` - open_source_project - a feedback management platform -> which has very good code -> as they follow lof of conventions , idiomatic GO writing practices - sriniously studied it often for learning and referencing code ( we should also do that) & during this course we will refer some open source repos & so we can learn how the concepts work in real world. we see how exactly they work and how you will use them. (literally typing the words of sriniously - he's making a lot of sense to me as i didn't know how to understand things right way - thank you so much sir) - everyone should watch the video.

```bash
abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part1 (main)
$ go run main.go
Hello, 世界

abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part1 (main)
$ go mod init github.com/abhimvp/GoBasics/Part1
go: creating new go.mod: module github.com/abhimvp/GoBasics/Part1
go: to add module requirements and sums:
        go mod tidy
```

- Now let's go into `main.go` file:

  - we have a package declaration -> `package main`
    - `packages` in `GO` are there to organize our code in a meaningful way. Every GO file that ends with extension `.go` - the file starts with a package declaration -> In a big project Structure - The Directory/Folder name will often match the package name.
    - the file `main.go` belongs to package main.
    - Rule is if you're building an executable - a program that will run in our system or in cloud - it has to have one main package & it will have a main function -> why is that because this function is the entry point of the program - means - when we compile this package or run this package - the execution starts from `function main` & it calls other functions and all.
    - the directory and folder structure we will get it to later.
    - So package are a way to organize our code - divide our code in different\*2 files, so that it is easier to manage. It also offers some kind of properties like encapsulation and all, we learn about it soon.
  - we have imports -> `import "fmt"`
    - imports a package from standard library - the package "fmt" has lot of functions - all related to print or string manipulation.
  - we have a function - which name is main & using the imported package "fmt" here and calling Println function and passing a string which gets printed in our console(terminal) -> `func main() { fmt.Println("Hello, WorldInChinese")}`
    - Functions are a way to group logically related code statements or expressions - function name is main and every executable code has to have a function name main because that is the entry point. Also focus that P is capital and it is case sensitive.
    - Another thing to notice is why did we write in CHINESE in print function because to express the concept that `GO` is a language which supports `UTF-8` encoding characters - means - in previous programming languages they only supported ASCII character encoding format - which has a mechanism to express upto 127 characters which consisted of lower,upper case english letters , 0 to 9 & some special characters..etc which is very limiting & this `UTF-8` falls under unicode character encoding system & `UTF-8` is one encoding scheme which supports almost 12,000..+ characters.

- There is another way to run or build this program `main.go`:

  - we can do `go build main.go` -> it generates an executable and we can run this executable using `./main` & we can see the same output as above. Go provides a feature called cross-platform single binary executables -> we can take it and run on any operating system or we can create a executable for specific platform/OS as well ` GOOS=windows GOARCH=amd64 go build .` which generates a for-example - `module_name.exe` - we can send this binary to that system and it runs fine.

- Difference between `go run` and `go build` & running that executable -> while you're developing you should use the `go run` command - and `go build` does is goes through your code, compiles all the packages, does the linking and in the end generates the binary.But `go run` command -> goes through the code,compiles all the package and it directly runs on your system without generating a single binary - it directly runs your program with a temporary binary somewhere in the system. `go build` throws the result in garbage and starts from scratch again. `go run` saves the result somewhere in the system & if nothing is changed in file the results will be faster as it use those cached results to compile our program faster. You will `go build` when you deploy your application or distributing across different systems in your CI/CD pipelines or cloud dev environments.

- **Adding package "math/rand"**

- we have added another package called `"math/rand"` - we use this package to generate random numbers and every time we run the program we get a number between 0 to 10.
- Reason to focus on this example is that `Most packages will have these prefixes like math - the path the package can be accessible from` - the important one is the last part of that package import `rand` - which you will use that in your own program & that's how you referred `rand.Intn(10)` & the best practice to name your packages is that the - package name cannot be capital letters \_ they have to be small case - the packages names have to be as short as possible - not a rule but a convention.
- Typically package name matches the name of the folder apart from test ones. ( refer fider ).
- when we are in a different package(main) and trying to access something from another package we imported , the name will always have to start with CAPITAL Letter , because all the variables , data types that start in small case - can't be accessed outside & only those data types whose names start with a capital case are accessible from outside the package - this is one example of goes encapsulation - a technique to hide implementation or private variables. Example - `import "math" and fmt.Println(math.Pi)`

```bash
abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part1 (main)
$ go run main.go
Hello, 世界
My Favorite Number is 3
3.141592653589793
```

- **Functions**

- now let's look at custom functions in the same `main.go` file.
- There is structure a function follows:
  - first we have the keyword `func`
  - Then we have the name of the function `add`
  - then we have braces `()` inside which we'll have parameter
  - parameters will follow a structure - `name_of_parameter space type` - `x int,y int` - x,y are parameters.
  - then we will explicitly have to define the return type of the function - mandatory to define what type the function returns - `int`
  - we write statements inside `return x + y`.
  - we call this function inside main function by providing two arguments of type int.
  - Go is a statically typed language - means all the types are verified,checked during compile times.

```go
package main

import (
 "fmt"
 "math"
 "math/rand"
)

func add(x int, y int) int { // we can also write it as add(x,y int)
 return x + y
}

func main() {
 fmt.Println("Hello, 世界")
 fmt.Println("My Favorite Number is", rand.Intn(10))
 fmt.Println(math.Pi)
 fmt.Println("Adding two numbers: ", add(3, 54))
}

```

```bash
abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part1 (main)
$ go run main.go
Hello, 世界
My Favorite Number is 5
3.141592653589793
Adding two numbers:  57
```

- Go has a feature that a `function can return multiple values` from inside it.
  - we look at `swap` `func`tion example in our code file - `main.go`
  - swap takes two parameters x and y and returns two values of type string.
  - we use it inside main function & we provide it two arguments.
  - we can use this structure `a,b := (comma separated variables structure) to read the both values returned by the swap function into these new variables`.

```go
func swap(x, y string) (string, string) {
return y, x
}
<!-- Inside function -->
a, b := swap("hello", "world")
fmt.Println("Swapped values:", a, b)
// output
// Swapped values: world hello
```

- Another way to define return values in our function:

```go
func split(sum int) (x, y int) { // we have named those return values instead of just (int,int) - this way those two variables are initialized with their zero values. before the start of function two variables are implicitly declared to zero , and we're not initializing(:=) a variable inside split function we're directly assigning it(=) because of that we don't have to write return x,y.
 x = sum * 4 / 9
 y = sum - x
 return
}
// Inside function
x, y := split(17)
fmt.Println("Split values:", x, y)
// output
// Split values: 7 10
```

- suggestion is don't use the above format for large functions
