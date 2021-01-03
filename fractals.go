// +build js,wasm

package fractals

import (
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

type FractalApp struct {
	runCb, runBenchmarkCb, shutdownCb js.Func

	console js.Value
	done    chan struct{}
}

// New returns a new instance of shimmer
func New() *FractalApp {
	return &FractalApp{
		console: js.Global().Get("console"),
		done:    make(chan struct{}),
	}
}

// Start sets up all the callbacks and waits for the close signal
// to be sent from the browser.
func (w *FractalApp) Start() {
	// Make sure the run button works
	w.setupRunCb()
	js.Global().Get("document").
		Call("getElementById", "run").
		Call("addEventListener", "click", w.runCb)

	// Make sure we can run benchmarks
	w.setupRunBenchmarkCb()
	js.Global().Get("document").
		Call("getElementById", "runBenchmark").
		Call("addEventListener", "click", w.runBenchmarkCb)

	// Make sure the kill button works
	w.setupShutdownCb()
	js.Global().Get("document").
		Call("getElementById", "close").
		Call("addEventListener", "click", w.shutdownCb)

	w.rebuildDragon(10)

	<-w.done
	w.log("Shutting down app")
	w.runCb.Release()
	w.runBenchmarkCb.Release()
	w.shutdownCb.Release()
}

// utility function to log a msg to the UI from inside a callback
func (w *FractalApp) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
}

func (w *FractalApp) consoleLog(str string) {
	w.console.Call("log", str)
}

func (w *FractalApp) setupShutdownCb() {
	w.shutdownCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		w.done <- struct{}{}
		return nil
	})
}

func (w *FractalApp) setupRunCb() {
	w.runCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Find out how many steps we'd like to generate
		v := js.Global().Get("document").
			Call("getElementById", "numSteps").
			Get("value")

		nSteps, err := strconv.Atoi(v.String())
		if err != nil {
			w.log(err.Error())
			nSteps = 4
		}
		w.rebuildDragon(uint64(nSteps))

		return nil
	})
}

func (w *FractalApp) setupRunBenchmarkCb() {
	w.runBenchmarkCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		v := js.Global().Get("document").
			Call("getElementById", "benchmarkPower").
			Get("value")

		power, err := strconv.Atoi(v.String())
		if err != nil {
			w.log(err.Error())
			power = 4
		}

		v = js.Global().Get("document").
			Call("getElementById", "benchmarkRepeats").
			Get("value")

		numRepeats, err := strconv.Atoi(v.String())
		if err != nil {
			w.log(err.Error())
			numRepeats = 5
		}
		w.log(`running benchmarks (check console)`)
		w.runBenchmarks(uint64(power), numRepeats)

		return nil
	})
}

func (w *FractalApp) rebuildDragon(n uint64) {

	w.log(fmt.Sprintf("building path with 2^%v steps...", n))

	path, vb, dur := getPathAndViewBoxForDragonWithDuration(n)

	js.Global().Get("document").
		Call("getElementById", "svgID").
		Call("setAttribute", "viewBox", vb)

	// find the svg and set the path
	js.Global().Get("document").
		Call("getElementById", "pathID").
		Call("setAttribute", "d", path)

	w.log(fmt.Sprintf("building path with 2^%v steps...Completed in %s!", n, dur.String()))
}

func getPathAndViewBoxForDragonWithDuration(n uint64) (string, string, time.Duration) {
	t0 := time.Now()
	path, vb := getPathAndViewBoxForDragon(n)

	return path, vb, time.Since(t0)
}

func (w *FractalApp) runBenchmarks(maxPower uint64, numRepeats int) {
	for power := maxPower; power > 0; power -= 1 {
		for i := 0; i < numRepeats; i++ {
			_, _, dur := getPathAndViewBoxForDragonWithDuration(power)
			w.consoleLog(fmt.Sprintf("power = %d; avgDur = %s\n", power, dur.String()))
		}
	}
}
