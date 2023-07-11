package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __mainwindow struct{}

func (*__mainwindow) init() {}

type MainWindow struct {
	*__mainwindow
	*widgets.QMainWindow
	ActionEnglish_Last    *widgets.QAction
	ActionHungarian_First *widgets.QAction
	ActionMale            *widgets.QAction
	ActionFemale          *widgets.QAction
	ActionLast_Name       *widgets.QAction
	ActionMale_2          *widgets.QAction
	ActionFemale_2        *widgets.QAction
	ActionLast_Name_2     *widgets.QAction
	Centralwidget         *widgets.QWidget
	GridLayout            *widgets.QGridLayout
	VerticalLayout        *widgets.QVBoxLayout
	HorizontalLayout      *widgets.QHBoxLayout
	VerticalLayout_3      *widgets.QVBoxLayout
	HorizontalLayout_3    *widgets.QHBoxLayout
	ListOfVowels          *widgets.QLabel
	VowelImport           *widgets.QPushButton
	Vowels                *widgets.QLineEdit
	HorizontalLayout_4    *widgets.QHBoxLayout
	ListOfbases           *widgets.QLabel
	HorizontalLayout_7    *widgets.QHBoxLayout
	CheckBox              *widgets.QCheckBox
	BaseImport            *widgets.QPushButton
	VerticalLayout_2      *widgets.QVBoxLayout
	HorizontalLayout_6    *widgets.QHBoxLayout
	SylNum                *widgets.QLabel
	NumberOfSyl           *widgets.QSpinBox
	HorizontalLayout_5    *widgets.QHBoxLayout
	GenNumber             *widgets.QLabel
	NumberOfGen           *widgets.QSpinBox
	HorizontalLayout_2    *widgets.QHBoxLayout
	Input                 *widgets.QPlainTextEdit
	Output                *widgets.QPlainTextEdit
	HorizontalLayout_8    *widgets.QHBoxLayout
	HorizontalSpacer      *widgets.QSpacerItem
	Generate              *widgets.QPushButton
	Save                  *widgets.QPushButton
	MenuBar               *widgets.QMenuBar
	MenuChoose_Language   *widgets.QMenu
	MenuEnglish           *widgets.QMenu
	MenuFirst_name        *widgets.QMenu
	MenuHungarian         *widgets.QMenu
	MenuFirst_Name        *widgets.QMenu
}

func NewMainWindow(p widgets.QWidget_ITF) *MainWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &MainWindow{QMainWindow: widgets.NewQMainWindow(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *MainWindow) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("MainWindow")
	}
	w.Resize2(1087, 746)
	w.ActionEnglish_Last = widgets.NewQAction(w)
	w.ActionEnglish_Last.SetObjectName("actionEnglish_Last")
	w.ActionHungarian_First = widgets.NewQAction(w)
	w.ActionHungarian_First.SetObjectName("actionHungarian_First")
	w.ActionMale = widgets.NewQAction(w)
	w.ActionMale.SetObjectName("actionMale")
	w.ActionFemale = widgets.NewQAction(w)
	w.ActionFemale.SetObjectName("actionFemale")
	w.ActionLast_Name = widgets.NewQAction(w)
	w.ActionLast_Name.SetObjectName("actionLast_Name")
	w.ActionMale_2 = widgets.NewQAction(w)
	w.ActionMale_2.SetObjectName("actionMale_2")
	w.ActionFemale_2 = widgets.NewQAction(w)
	w.ActionFemale_2.SetObjectName("actionFemale_2")
	w.ActionLast_Name_2 = widgets.NewQAction(w)
	w.ActionLast_Name_2.SetObjectName("actionLast_Name_2")
	w.Centralwidget = widgets.NewQWidget(w, 0)
	w.Centralwidget.SetObjectName("centralwidget")
	w.GridLayout = widgets.NewQGridLayout(w.Centralwidget)
	w.GridLayout.SetObjectName("gridLayout")
	w.VerticalLayout = widgets.NewQVBoxLayout()
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.HorizontalLayout = widgets.NewQHBoxLayout()
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.VerticalLayout_3 = widgets.NewQVBoxLayout()
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.ListOfVowels = widgets.NewQLabel(w.Centralwidget, 0)
	w.ListOfVowels.SetObjectName("listOfVowels")
	w.HorizontalLayout_3.QLayout.AddWidget(w.ListOfVowels)
	w.VowelImport = widgets.NewQPushButton(w.Centralwidget)
	w.VowelImport.SetObjectName("vowelImport")
	w.HorizontalLayout_3.AddWidget(w.VowelImport, 0, core.Qt__AlignRight)
	w.VerticalLayout_3.AddLayout(w.HorizontalLayout_3, 0)
	w.Vowels = widgets.NewQLineEdit(w.Centralwidget)
	w.Vowels.SetObjectName("vowels")
	sizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Ignored, widgets.QSizePolicy__Fixed, 0)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(w.Vowels.SizePolicy().HasHeightForWidth())
	w.Vowels.SetSizePolicy(sizePolicy)
	w.VerticalLayout_3.QLayout.AddWidget(w.Vowels)
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.ListOfbases = widgets.NewQLabel(w.Centralwidget, 0)
	w.ListOfbases.SetObjectName("listOfbases")
	w.HorizontalLayout_4.QLayout.AddWidget(w.ListOfbases)
	w.HorizontalLayout_7 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_7.SetObjectName("horizontalLayout_7")
	w.CheckBox = widgets.NewQCheckBox(w.Centralwidget)
	w.CheckBox.SetObjectName("checkBox")
	w.CheckBox.SetLayoutDirection(core.Qt__RightToLeft)
	w.HorizontalLayout_7.AddWidget(w.CheckBox, 0, core.Qt__AlignLeft)
	w.BaseImport = widgets.NewQPushButton(w.Centralwidget)
	w.BaseImport.SetObjectName("baseImport")
	w.BaseImport.SetEnabled(true)
	sizePolicy1 := widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, 0)
	sizePolicy1.SetHorizontalStretch(0)
	sizePolicy1.SetVerticalStretch(0)
	sizePolicy1.SetHeightForWidth(w.BaseImport.SizePolicy().HasHeightForWidth())
	w.BaseImport.SetSizePolicy(sizePolicy1)
	w.HorizontalLayout_7.QLayout.AddWidget(w.BaseImport)
	w.HorizontalLayout_4.AddLayout(w.HorizontalLayout_7, 0)
	w.VerticalLayout_3.AddLayout(w.HorizontalLayout_4, 0)
	w.HorizontalLayout.AddLayout(w.VerticalLayout_3, 0)
	w.VerticalLayout_2 = widgets.NewQVBoxLayout()
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.HorizontalLayout_6 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_6.SetObjectName("horizontalLayout_6")
	w.SylNum = widgets.NewQLabel(w.Centralwidget, 0)
	w.SylNum.SetObjectName("sylNum")
	w.HorizontalLayout_6.QLayout.AddWidget(w.SylNum)
	w.NumberOfSyl = widgets.NewQSpinBox(w.Centralwidget)
	w.NumberOfSyl.SetObjectName("numberOfSyl")
	w.NumberOfSyl.SetMinimum(1)
	w.NumberOfSyl.SetMaximum(999)
	w.HorizontalLayout_6.QLayout.AddWidget(w.NumberOfSyl)
	w.VerticalLayout_2.AddLayout(w.HorizontalLayout_6, 0)
	w.HorizontalLayout_5 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_5.SetObjectName("horizontalLayout_5")
	w.GenNumber = widgets.NewQLabel(w.Centralwidget, 0)
	w.GenNumber.SetObjectName("genNumber")
	w.HorizontalLayout_5.QLayout.AddWidget(w.GenNumber)
	w.NumberOfGen = widgets.NewQSpinBox(w.Centralwidget)
	w.NumberOfGen.SetObjectName("numberOfGen")
	w.NumberOfGen.SetMinimum(1)
	w.NumberOfGen.SetMaximum(999)
	w.HorizontalLayout_5.QLayout.AddWidget(w.NumberOfGen)
	w.VerticalLayout_2.AddLayout(w.HorizontalLayout_5, 0)
	w.HorizontalLayout.AddLayout(w.VerticalLayout_2, 0)
	w.VerticalLayout.AddLayout(w.HorizontalLayout, 0)
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.Input = widgets.NewQPlainTextEdit(w.Centralwidget)
	w.Input.SetObjectName("input")
	w.Input.SetLineWrapMode(widgets.QPlainTextEdit__NoWrap)
	w.HorizontalLayout_2.QLayout.AddWidget(w.Input)
	w.Output = widgets.NewQPlainTextEdit(w.Centralwidget)
	w.Output.SetObjectName("output")
	w.Output.SetLineWrapMode(widgets.QPlainTextEdit__NoWrap)
	w.Output.SetTextInteractionFlags(core.Qt__TextSelectableByKeyboard | core.Qt__TextSelectableByMouse)
	w.HorizontalLayout_2.QLayout.AddWidget(w.Output)
	w.VerticalLayout.AddLayout(w.HorizontalLayout_2, 0)
	w.GridLayout.AddLayout2(w.VerticalLayout, 0, 0, 1, 1, 0)
	w.HorizontalLayout_8 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_8.SetObjectName("horizontalLayout_8")
	w.HorizontalSpacer = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_8.AddItem(w.HorizontalSpacer)
	w.Generate = widgets.NewQPushButton(w.Centralwidget)
	w.Generate.SetObjectName("generate")
	w.HorizontalLayout_8.QLayout.AddWidget(w.Generate)
	w.Save = widgets.NewQPushButton(w.Centralwidget)
	w.Save.SetObjectName("save")
	w.HorizontalLayout_8.QLayout.AddWidget(w.Save)
	w.GridLayout.AddLayout2(w.HorizontalLayout_8, 1, 0, 1, 1, 0)
	w.SetCentralWidget(w.Centralwidget)
	w.MenuBar = widgets.NewQMenuBar(w)
	w.MenuBar.SetObjectName("menuBar")
	w.MenuBar.SetGeometry(core.NewQRect4(0, 0, 1087, 28))
	w.MenuChoose_Language = widgets.NewQMenu(w.MenuBar)
	w.MenuChoose_Language.SetObjectName("menuChoose_Language")
	w.MenuEnglish = widgets.NewQMenu(w.MenuChoose_Language)
	w.MenuEnglish.SetObjectName("menuEnglish")
	w.MenuFirst_name = widgets.NewQMenu(w.MenuEnglish)
	w.MenuFirst_name.SetObjectName("menuFirst_name")
	w.MenuHungarian = widgets.NewQMenu(w.MenuChoose_Language)
	w.MenuHungarian.SetObjectName("menuHungarian")
	w.MenuFirst_Name = widgets.NewQMenu(w.MenuHungarian)
	w.MenuFirst_Name.SetObjectName("menuFirst_Name")
	w.SetMenuBar(w.MenuBar)
	w.MenuBar.AddActions([]*widgets.QAction{w.MenuChoose_Language.MenuAction()})
	w.MenuChoose_Language.AddActions([]*widgets.QAction{w.MenuEnglish.MenuAction()})
	w.MenuChoose_Language.AddActions([]*widgets.QAction{w.MenuHungarian.MenuAction()})
	w.MenuEnglish.AddActions([]*widgets.QAction{w.MenuFirst_name.MenuAction()})
	w.MenuEnglish.AddActions([]*widgets.QAction{w.ActionLast_Name})
	w.MenuFirst_name.AddActions([]*widgets.QAction{w.ActionMale})
	w.MenuFirst_name.AddActions([]*widgets.QAction{w.ActionFemale})
	w.MenuHungarian.AddActions([]*widgets.QAction{w.MenuFirst_Name.MenuAction()})
	w.MenuHungarian.AddActions([]*widgets.QAction{w.ActionLast_Name_2})
	w.MenuFirst_Name.AddActions([]*widgets.QAction{w.ActionMale_2})
	w.MenuFirst_Name.AddActions([]*widgets.QAction{w.ActionFemale_2})
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *MainWindow) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("MainWindow", "Chawurgon", "", 0))
	w.ActionEnglish_Last.SetText(core.QCoreApplication_Translate("MainWindow", "English Last Name", "", 0))
	w.ActionHungarian_First.SetText(core.QCoreApplication_Translate("MainWindow", "Hungarian First Name", "", 0))
	w.ActionMale.SetText(core.QCoreApplication_Translate("MainWindow", "Male", "", 0))
	w.ActionFemale.SetText(core.QCoreApplication_Translate("MainWindow", "Female", "", 0))
	w.ActionLast_Name.SetText(core.QCoreApplication_Translate("MainWindow", "Last Name", "", 0))
	w.ActionMale_2.SetText(core.QCoreApplication_Translate("MainWindow", "Male", "", 0))
	w.ActionFemale_2.SetText(core.QCoreApplication_Translate("MainWindow", "Female", "", 0))
	w.ActionLast_Name_2.SetText(core.QCoreApplication_Translate("MainWindow", "Last Name", "", 0))
	w.ListOfVowels.SetText(core.QCoreApplication_Translate("MainWindow", "List of Vowels:", "", 0))
	w.VowelImport.SetText(core.QCoreApplication_Translate("MainWindow", "Import", "", 0))
	w.Vowels.SetText(core.QCoreApplication_Translate("MainWindow", "aeiou", "", 0))
	w.ListOfbases.SetText(core.QCoreApplication_Translate("MainWindow", "List of Bases:", "", 0))
	w.CheckBox.SetText(core.QCoreApplication_Translate("MainWindow", "Keep current list", "", 0))
	w.BaseImport.SetText(core.QCoreApplication_Translate("MainWindow", "Import", "", 0))
	w.SylNum.SetText(core.QCoreApplication_Translate("MainWindow", "Number of Syllables:", "", 0))
	w.GenNumber.SetText(core.QCoreApplication_Translate("MainWindow", "Number of Generations:", "", 0))
	w.Generate.SetText(core.QCoreApplication_Translate("MainWindow", "Generate", "", 0))
	w.Save.SetText(core.QCoreApplication_Translate("MainWindow", "Save", "", 0))
	w.MenuChoose_Language.SetTitle(core.QCoreApplication_Translate("MainWindow", "Select Preset", "", 0))
	w.MenuEnglish.SetTitle(core.QCoreApplication_Translate("MainWindow", "English", "", 0))
	w.MenuFirst_name.SetTitle(core.QCoreApplication_Translate("MainWindow", "First Name", "", 0))
	w.MenuHungarian.SetTitle(core.QCoreApplication_Translate("MainWindow", "Hungarian", "", 0))
	w.MenuFirst_Name.SetTitle(core.QCoreApplication_Translate("MainWindow", "First Name", "", 0))

}
