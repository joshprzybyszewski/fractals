// +build js,wasm

package fractals

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/joshprzybyszewski/fractals/drawing"
)

type FractalApp struct {
	runCb, shutdownCb js.Func

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

	// Make sure the kill button works
	w.setupShutdownCb()
	js.Global().Get("document").
		Call("getElementById", "close").
		Call("addEventListener", "click", w.shutdownCb)

	<-w.done
	w.log("Shutting down app")
	w.runCb.Release()
	w.shutdownCb.Release()
}

// utility function to log a msg to the UI from inside a callback
func (w *FractalApp) log(msg string) {
	js.Global().Get("document").
		Call("getElementById", "status").
		Set("innerText", msg)
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

		// TODO this likely isn't the best way to make this work
		nStepsStr, err := strconv.Atoi(v.String())
		if err != nil {
			w.log(err.Error())
			nStepsStr = 16
		}
		nSteps := uint64(nStepsStr)
		w.log(fmt.Sprintf("building path with %v steps...", nSteps))

		// update the path
		path := drawing.BuildPath(nSteps)

		// find the svg and set the path
		js.Global().Get("document").
			Call("getElementById", "pathID").
			Call("setAttribute", "d", path)

		w.log(fmt.Sprintf("building path with %v steps...Complete!", nSteps))

		return nil
	})
}
