package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"lab3/internal/common"
	"math"
	"strconv"
)

var (
	integral      = common.Integral{}
	integralValue = binding.NewString()
	n             = binding.NewString()
	nilFunction   *canvas.Text
	nilMethod     *canvas.Text
	answer1       *widget.Label
	answer2       *widget.Label
	progressBar   *widget.ProgressBarInfinite
	solveButton   *widget.Button
	stopButton    *widget.Button
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() {
	integral.SetAccuracy(1)
	myApp := fyneApp.New()
	window := myApp.NewWindow("Lab3")
	window.Resize(fyne.NewSize(800, 400))
	window.SetFixedSize(true)
	window.CenterOnScreen()

	line := canvas.NewLine(color.White)
	line.Position1 = fyne.NewPos(400, 10)
	line.Position2 = fyne.NewPos(400, 370)

	var (
		a float64
		b float64
	)

	dataA := binding.BindFloat(&a)
	dataB := binding.BindFloat(&b)

	solveButton = widget.NewButton("SOLVE", solve)
	solveButton.Move(fyne.NewPos(150, 300))
	solveButton.Resize(fyne.NewSize(100, 50))

	stopButton = widget.NewButton("STOP", func() {
		common.Quit <- true
	})
	stopButton.Move(fyne.NewPos(550, 300))
	stopButton.Resize(fyne.NewSize(100, 50))
	stopButton.DisableableWidget.Disable()
	stopButton.Hidden = true

	aLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(dataA, "%.2f"))
	bLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(dataB, "%.2f"))
	l := container.NewHBox(
		widget.NewLabel("A: "),
		aLabel,
		widget.NewLabel("B: "),
		bLabel)
	l.Move(fyne.NewPos(50, 80))
	aInput, bInput := segmentInput(dataA, dataB)

	nilFunction = canvas.NewText("Choose a function", color.RGBA{
		R: 214,
		G: 0,
		B: 0,
		A: 1,
	})
	nilMethod = canvas.NewText("Choose a method", color.RGBA{
		R: 214,
		G: 0,
		B: 0,
		A: 1,
	})
	nilFunction.Move(fyne.NewPos(58, 60))
	nilMethod.Move(fyne.NewPos(58, 210))
	nilFunction.Hidden = true
	nilMethod.Hidden = true

	input := container.NewWithoutLayout(
		line,
		functionSelect(),
		l,
		aInput,
		bInput,
		accuracySpinner(),
		methodSelect(),
		solveButton,
		nilMethod,
		nilFunction)
	input.Resize(fyne.NewSize(400, 400))

	answer1 = widget.NewLabelWithData(integralValue)
	answer2 = widget.NewLabelWithData(n)
	answer1.Move(fyne.NewPos(450, 150))
	answer2.Move(fyne.NewPos(450, 200))
	answer1.Wrapping = 0
	progressBar = widget.NewProgressBarInfinite()
	progressBar.Resize(fyne.NewSize(300, 50))
	progressBar.Move(fyne.NewPos(450, 150))
	progressBar.Hidden = true
	progressBar.Stop()

	window.SetContent(container.NewWithoutLayout(input, answer1, answer2, progressBar, stopButton))
	window.ShowAndRun()
}

func functionSelect() *widget.Select {
	functionSelection := widget.NewSelect(common.FunctionNames, func(s string) {
		integral.SetFunction(s)
		integral.SetPrimitive(s)
	})
	functionSelection.PlaceHolder = "Choose a function"
	functionSelection.Resize(fyne.NewSize(300, 30))
	functionSelection.Move(fyne.NewPos(50, 20))
	return functionSelection
}

func methodSelect() *widget.Select {
	methodSelection := widget.NewSelect(common.MethodNames, integral.SetMethod)
	methodSelection.PlaceHolder = "Choose a method"
	methodSelection.Resize(fyne.NewSize(300, 30))
	methodSelection.Move(fyne.NewPos(50, 170))
	return methodSelection
}

func segmentInput(aData binding.Float, bData binding.Float) (*widget.Slider, *widget.Slider) {
	aInput := widget.NewSlider(-5, 5)
	bInput := widget.NewSlider(-5, 5)
	aInput.Step = 0.05
	bInput.Step = 0.05
	aInput.OnChanged = func(a float64) {
		if a > bInput.Value {
			aInput.SetValue(bInput.Value)
		}
		integral.SetA(aInput.Value)
		aData.Set(aInput.Value)
	}
	bInput.OnChanged = func(b float64) {
		if b < aInput.Value {
			bInput.SetValue(aInput.Value)
		}
		integral.SetB(bInput.Value)
		bData.Set(bInput.Value)
	}
	aInput.Resize(fyne.NewSize(300, 10))
	bInput.Resize(fyne.NewSize(300, 10))
	aInput.Move(fyne.NewPos(50, 110))
	bInput.Move(fyne.NewPos(50, 140))
	return aInput, bInput
}

func accuracySpinner() *fyne.Container {
	signs := 0
	data := binding.BindInt(&signs)
	signLabel := widget.NewLabelWithData(binding.IntToString(data))
	upButton := widget.NewButtonWithIcon("", theme.MoveUpIcon(), func() {
		v, _ := data.Get()
		v += 1
		data.Set(v)
		integral.SetAccuracy(math.Pow(10, -float64(v)))
	})
	downButton := widget.NewButtonWithIcon("", theme.MoveDownIcon(), func() {
		v, _ := data.Get()
		if v > 0 {
			v -= 1
		}
		data.Set(v)
		integral.SetAccuracy(math.Pow(10, -float64(v)))
	})
	downButton.Resize(fyne.NewSize(20, 20))
	upButton.Resize(fyne.NewSize(20, 20))
	downButton.Move(fyne.NewPos(220, 27))
	upButton.Move(fyne.NewPos(220, 3))
	labels := container.NewHBox(widget.NewLabel("Accuracy (signs)"), signLabel)
	labels = container.NewVBox(layout.NewSpacer(), labels, layout.NewSpacer())
	labels.Resize(fyne.NewSize(260, 50))
	c := container.NewWithoutLayout(labels, upButton, downButton)
	c.Move(fyne.NewPos(50, 220))
	c.Resize(fyne.NewSize(300, 50))
	return c
}

func solve() {
	go func() {
		if !(integral.GetFunction() == nil || integral.GetMethod() == nil) {
			stopButton.DisableableWidget.Enable()
			stopButton.Hidden = false
			solveButton.DisableableWidget.Disable()
			answer1.Hidden = true
			answer2.Hidden = true
			progressBar.Hidden = false
			progressBar.Start()
			defer func() {
				stopButton.Hidden = true
				progressBar.Hidden = true
				progressBar.Stop()
				stopButton.DisableableWidget.Disable()
				solveButton.DisableableWidget.Enable()
			}()
			nilFunction.Hidden = true
			nilMethod.Hidden = true
			integral.Solve()
			result := <-common.Results
			if result != nil {
				if math.IsNaN(result.IntegralValue) {
					integralValue.Set("Impossible to calculate the integral\non a given segment")
					n.Set("")
				} else {
					signs := -math.Log10(integral.Accuracy())
					s := "%." + strconv.FormatInt(int64(signs), 8) + "f"
					integralValue.Set(fmt.Sprintf("Integral value: "+s, result.IntegralValue))
					n.Set(fmt.Sprintf("Number of partitions: %v", result.N))
				}
				answer1.Hidden = false
				answer2.Hidden = false
			}
		} else {
			nilMethod.Hidden = integral.GetMethod() != nil
			nilFunction.Hidden = integral.GetFunction() != nil
		}
	}()
}
