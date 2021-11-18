# Computer Graphics and Multi Media Practical programs in Go<img src="res/golang-48.png">
CGMM Practical lab program written in Go

This repo contains variety of intersting agorithms I found in my Computer Graphics course I took in 5<sup>th</sup> semester (2021). 
For the purpose of plotting I am using [raylib-go](https://github.com/gen2brain/raylib-go); Go bindings for original [Raylib](http://www.raylib.com/) library written in C.

## To run the application follow these Intstructions:
Make sure you have installed go, and specified `GOPATH`.
1. You need C compiler as raylib depends on that, like [Mingw-w64](https://mingw-w64.org/) or [TDM-GCC](http://tdm-gcc.tdragon.net/). You can also build binary in [MSYS2](https://msys2.github.io/) shell.
2. Install raylib-go by typing in terminal.
  
    ```go
    go get -v -u github.com/gen2brain/raylib-go/raylib
    ```
3. Run Test functions written in `CGMM_test.go` using your text editor or using command line:
  
    ```go
    go test -run <test_function_name> CGMM_programs
    ```
    The command Line invocation may only respond with Pass/Fail if the test function don't initialize any GUI, ignoring any print statement. 
    
    In that case resort to running tests using your text editor or create a main function and invoke the test function.

## Algorithms covered:
- 📜<code>**line.go**</code>
  
    1. DDA (Digital Differential Analyzer)
    2. Breseham algorithm

- 📜<code>**clipping.go**</code>
    
    1. Cohen Sutherland clipping algorithm
- 📜<code>**circle.go**</code>
  
    1. Bresenham circle drawing algorithm
    2. Mid point circle drawing algorithm
- 📜<code>**fill.go**</code>
    1. Flood fill algorithm 
    2. Boundary fill algorithm

The `CGMM_test` also contains a test animation function (a mini-game).

## Other files:
- <code>**utils.go**</code>: Contains Utility functions used across the project.
- <code>**config.go**</code>: Conatins global variables like max height and width of application window.
- <code>**go.mod**</code>: Module file listing all other libraries the project depends upon.
