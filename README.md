# fractals

fun weekend project regarding the dragon fractal

### Inspiration

### What it is

### How I built it

#### The golang WASM app

Early-ish in quarantine 2020, I had built another goWASM app called [wonder](github.com/joshprzybyszewski/wonder). I just went about reviving that in order to get the display working, and it took a few steps:

1. Get the latest wasm_exec.js file.
   - According to [this blog](https://www.sitepen.com/blog/compiling-go-to-webassembly), I should be able to copy from `$GOROOT/misc/wasm/wasm_exec.js`. And according to [this blog](https://medium.com/swlh/getting-started-with-webassembly-and-go-by-building-an-image-to-ascii-converter-dea10bdf71f6), I was able to find it with: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .` in go1.15.
