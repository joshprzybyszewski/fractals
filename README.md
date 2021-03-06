# fractals

fun weekend project regarding the dragon fractal

### Demo

![dragonfractal10](./demos/dragonfractal10.gif)

### Inspiration

My brother recently asked what got me into programming. I remembered that one of my first personal programming projects was to build a fractal UI in Java (I found the original source code and added it in the [oldjava directory](./oldjava/README.md)). Since I've been programming for 7 years since then, I thought it would be a fun exercise to do that same project again with new tools.

### What it is

This is called the "dragon fractal" (AKA the Jurassic Park fractal). You can google more about it (here's a [quick link](https://www.mathworks.com/matlabcentral/mlc-downloads/downloads/submissions/11069/versions/2/previews/html/dragon_curve.html#:~:text=The%20Dragon%20Curve%20is%20a,a%20novel%20by%20Michael%20Crichton.&text=The%20user%20is%20encouraged%20to,Number%20of%20fractal%20iterations)). I've always been fascinated with this fractal because it's a model of folding a piece of paper in half "n" times and then unfolding it so that each crease is at 90 degrees. What's fascinating and unexpected is that the paper will never attempt to cross-over itself; instead it may meet at an intersection but then parts ways.

### How I built it

This time around I used golang and WebAssembly. WASM is "the way of the future" for client-side applications, and golang is a super fun and fast backend language that I've recently enjoyed. Why not mash them together and see how they work?

#### Concept

In my golang implementation, I chose to draw the fractal by building an SVG path that uses a recursive function to decide which way the line should turn. That is, I've conceptualized the fractal as a path that we walk. We will always move in a straight line for a distance of `delta` either horizontally or vertically, and then we make either a left-handed or right-handed turn 90 degrees. In order to determine which way the path turns, we use a recursive function (`IsLeftTurn`) that accepts as input the current number of path segments we've walked. The path will turn left if we've walked a "power of two" segments, or if we're an even number of inverses away from a power of two. See:

| segments walked (n) | turn (left/right) | mirror relationship | distance from power |
| ------------------: | ----------------- | ------------------- | ------------------: |
|                   1 | L                 |                     |               (2^0) |
|                   2 | L                 |                     |               (2^1) |
|                   3 | R                 | !seg(1)             |              !(2^0) |
|                   4 | L                 |                     |               (2^2) |
|                   5 | L                 | !seg(3)             |             !!(2^0) |
|                   6 | R                 | !seg(2)             |              !(2^1) |
|                   7 | R                 | !seg(1)             |              !(2^0) |
|                   8 | L                 |                     |               (2^3) |
|                   9 | L                 | !seg(7)             |             !!(2^0) |
|                  10 | L                 | !seg(6)             |             !!(2^1) |
|                  11 | R                 | !seg(5)             |            !!!(2^0) |
|                  12 | R                 | !seg(4)             |              !(2^2) |
|                  13 | L                 | !seg(3)             |             !!(2^0) |
|                  14 | R                 | !seg(2)             |              !(2^1) |
|                  15 | R                 | !seg(1)             |              !(2^0) |
|                  16 | L                 |                     |               (2^4) |

Notice that if there is an even number of `!`s, that corresponds to a left turn. If there is an odd number, then we make a right turn.

#### The golang WASM app

Early-ish in quarantine 2020, I had built another goWASM app called [wonder](github.com/joshprzybyszewski/wonder). I just went about reviving that in order to get the display working, and the thing I needed to remember is:

1. Get the latest wasm_exec.js file.
   - According to [this blog](https://www.sitepen.com/blog/compiling-go-to-webassembly), I should be able to copy from `$GOROOT/misc/wasm/wasm_exec.js`. And according to [this blog](https://medium.com/swlh/getting-started-with-webassembly-and-go-by-building-an-image-to-ascii-converter-dea10bdf71f6), I was able to find it with: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` in go1.15.

### Run it yourself

1. Clone this repo.
1. `make vendor` for dependencies.
1. `make build` for WASM output.
1. `make serve` to stand up the simple file server locally.
1. Navigate to http://localhost:3434
1. Check it out!
