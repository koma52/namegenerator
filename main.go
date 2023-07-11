package main

import (
	"os"

	"github.com/therecipe/qt/widgets"

	"namegenerator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewMainWindow(nil).Show()

	widgets.QApplication_Exec()
}
