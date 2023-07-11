package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/therecipe/qt/widgets"
)

var checkBoxStatus bool = false

func (c *MainWindow) init() {
	c.VowelImport.ConnectClicked(func(bool) {
		fileDialog := widgets.NewQFileDialog(nil, 0)
		fileDialog.SetNameFilter("Text files (*.txt);;All files (*)")
		fileDialog.SetFileMode(widgets.QFileDialog__AnyFile)
		fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptOpen)

		if fileDialog.Exec() == int(widgets.QDialog__Accepted) {
			filename := fileDialog.SelectedFiles()[0]
			fileContent, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println("Error reading file:", err)
			}
			vowels := readVowels(fileContent)
			c.Vowels.SetText(strings.Join(vowels, ""))
		}
	})
	c.BaseImport.ConnectClicked(func(bool) {
		fileDialog := widgets.NewQFileDialog(nil, 0)
		fileDialog.SetNameFilter("Text files (*.txt);;All files (*)")
		fileDialog.SetFileMode(widgets.QFileDialog__AnyFile)
		fileDialog.SetAcceptMode(widgets.QFileDialog__AcceptOpen)

		if fileDialog.Exec() == int(widgets.QDialog__Accepted) {
			filename := fileDialog.SelectedFiles()[0]
			fileContent, err := os.ReadFile(filename)
			if err != nil {
				fmt.Println("Error reading file:", err)
			}
			if !checkBoxStatus {
				c.Input.SetPlainText(readBases(fileContent))
			} else {
				c.Input.AppendPlainText(readBases(fileContent))
			}
		}
	})
	c.Generate.ConnectClicked(func(bool) {
		generateNames(c)
	})
	c.CheckBox.ConnectStateChanged(func(state int) {
		if state == 0 {
			checkBoxStatus = false
		} else {
			checkBoxStatus = true
		}
	})
}
