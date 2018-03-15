package main

import (
  "github.com/andlabs/ui"
)

func main() {
  err := ui.Main(func() {
    frequency :=  ui.NewLabel("")
    input := ui.NewEntry()
    freqButton := ui.NewButton("Set Freq")
    freqUp := ui.NewButton("+")
    freqDown := ui.NewButton("-")
    vhfButton := ui.NewButton("2m")
    uhfButton := ui.NewButton("70cm")
    ssbButton := ui.NewButton("SSB")
    cwButton := ui.NewButton("CW")
    box := ui.NewVerticalBox()
    box.Append(ui.NewLabel("Frequency"), false)
    box.Append(frequency, false)
    box.Append(input, false)
    box.Append(freqButton, false)
    box.Append(freqUp, false)
    box.Append(freqDown, false)
    box.Append(ui.NewLabel("Radio Control"), false)
    box.Append(vhfButton, false)
    box.Append(uhfButton, false)
    box.Append(cwButton, false)
    box.Append(ssbButton, false)
    freqButton.OnClicked(func(*ui.Button) {
      frequency.SetText(input.Text())
    })
    window := ui.NewWindow("Rig Dashboard", 200, 100, false)
    window.SetMargined(true)
    window.SetChild(box)
    window.OnClosing(func(*ui.Window) bool {
        ui.Quit()
        return true
    })
    window.Show()
    })
    if err != nil {
      panic(err)
    }
}


