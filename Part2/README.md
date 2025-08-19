# Part2

- Before getting started i will create a go.mod file for Part2 learning using the following command.

```bash
abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part2 (main)
$ go mod init github.com/abhimvp/GoBasics/Part2
go: creating new go.mod: module github.com/abhimvp/GoBasics/Part2
```

---

## variables

- Now let's understand about variables.
- This is how we declare variables in GO
  - outside the function we can write keyword `var c, python, java bool` - c,python,java are name of the variables and all of them are same type `bool` - all these variables are initialized to their zero values , because in GO a variable is never in a uninitialized state/ undefined state / null state. since it's data type is bool it will be initialized to its `ZERO STATE` which is `false`.
  - If we are inside a function, we have couple of ways to do it:
    - First type is `var i int` - will initialize i value to Zero State of integer - that is `0`.
      - output `$ go run main.go -> 0 false false false`
    - Another type and used mostly -> `i := 0` if you already know the initialization value of a variable , you will use this.
      - this `:=` colon equals to format is only available inside the function.
    - Another way is `var i = 10` - using the initializer it interprets that i is type integer.
      - the `:=` is equal to writing `var i = 10`
- now to get some more idea of how people use it in real system application - we can refer `fider` - search the repo with `var` or `:=` you will see the initialization outside & inside the functions in package level..etc.

## data types

- let's look at various data types declared - `datatype.go`
  - `%T` verb is used to print the type of variable.
  - `%v` verb is used to print the value of variable.
  - `%s` to print value of a string
  - `%q` to print value of string in double quotes.

```bash
abhis@Tinku MINGW64 ~/Desktop/GoBasics/Part2 (main)
$ go run .
0 false false false
Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
```

- Go has basically four types of data types.

  - [basic](https://go.dev/tour/basics/11) -> number, string, boolean.
  - Aggregate -> array, struct -> collection of same and different data-types.
  - Reference -> Pointers, Slices, Functions, channels, Maps.
  - Interfaces

- While working on backend systems and when dealing with user data ( not sure about it's size) - then we can safely use `int`.
- `byte` also used to represent strings inside GO. // alias for unit8
- `rune` - alias for int32 - used to represent UTF-8 characters inside a go program.

- Zero values : Variables declared without an explicit initial value are given their zero value.
- The zero value is:

  - 0 for numeric types,
  - false for the boolean type, and
  - "" (the empty string) for strings.

- Type conversions : The expression T(v) converts the value v to the type T.In Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

- Type inference: When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

- Constants : Constants are declared like variables, but with the `const` keyword.
  - Constants can be character, string, boolean, or numeric values.
  - Constants cannot be declared using the `:=` syntax.
  - Numeric Constants:
    - Numeric constants are high-precision values.
    - An untyped constant takes the type needed by its context.
    - Try printing needInt(Big) too.
    - (An int can store at maximum a 64-bit integer, and sometimes less.)

## loops

- **For**

Go has only one looping construct, the for loop.

- The basic for loop has three components separated by semicolons:
  - the init statement: executed before the first iteration
  - the condition expression: evaluated before every iteration
  - the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

- To visualize the loop , we can use a `debugger`. add a breakpoint at sum+=i and debugger is basically mean you can investigate the flow of the program using this tool called `debugger`

  - VSCode has in-built debugger for GO.

- **For continued** : The init and post statements are optional.

- For is Go's "while": At that point you can drop the semicolons: C's while is spelled for in Go.

- Also we have this infinite loops - we use them especially during concurrency chapters - where we are constantly listening for signals - there we use this pattern of infinite loops.

## conditions

- **If**

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

- If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

- If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow return their results before the call to fmt.Println in main begins.)

## switch

- Switch

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

- Switch evaluation order : Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
- Switch with no condition : Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.

## defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

- Stacking defers

  - Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

- `Why do we need differs? use cases? `1 is when we're reading a lot of files and there can be memory leaks & every time you open a file , you have to explicitly close it. as a human error you may forget to close the file or database connection which we need to explicitly close while we are returning from the program otherwise memory leak happens. or during the execution of a program , the program is broken and throws error/exception - panic in go terminology - before you called close , the program panicked - stopped the flow of execution - now you never got to close - now you have a memory leak problem in your app - so for these use cases - go given the feature `defer` , now you don't have to remember to close the files/db connection at the end of function. what we can do is - we can open the file and in next line we can say defer(close file) - as we know the defer function is not called instantly and the moment an error occurs - go programming in a configured in a way it starts executing all the functions that were in the defer stack & starts executing them, it's a construct provided by go functionality. `- GREAT EXAMPLE`.
