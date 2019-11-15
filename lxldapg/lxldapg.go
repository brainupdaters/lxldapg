package main

import (
	"crypto/tls"
	"fmt"

	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"gopkg.in/ldap.v3"
)

var admInput *widgets.QComboBox
var roleName *widgets.QLineEdit
var roleNameDM *widgets.QListWidget
var changeNumber *widgets.QLineEdit
var userSelect *widgets.QLineEdit
var buttonu *widgets.QPushButton
var userNameDM *widgets.QListWidget
var groupSelect *widgets.QLineEdit
var buttong *widgets.QPushButton
var groupNameDM *widgets.QListWidget
var hostSelect *widgets.QLineEdit
var hostSelectDM *widgets.QListWidget
var commandName *widgets.QLineEdit
var commandNameDM *widgets.QListWidget
var optionName *widgets.QLineEdit
var optionNameDM *widgets.QListWidget
var resume *widgets.QTextEdit
var buttonExit *widgets.QPushButton
var buttonExecute *widgets.QPushButton
var buttonh *widgets.QPushButton
var buttonc *widgets.QPushButton
var buttono *widgets.QPushButton
var runAsName *widgets.QLineEdit
var runAsGName *widgets.QLineEdit
var sudoOrder *widgets.QLineEdit
var notAfter *widgets.QDateTimeEdit
var notBefore *widgets.QDateTimeEdit

var rau string
var rag string
var sdo string
var nbf string
var naf string

func check(e error) {
	if e != nil {
		//panic(e)
		fmt.Println(e.Error())
		showWarnigMod()
	}
}

func showWarnigMod() {
	widgets.QMessageBox_Warning(nil, "Warning", "Error executing modifications!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func main() {

	InitConfig()

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(500, 200)
	window.SetWindowTitle("Brain Updaters, SL")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	layouth0 := widgets.NewQWidget(nil, 0)
	layouth0.SetLayout(widgets.NewQHBoxLayout())

	layoutv1 := widgets.NewQWidget(nil, 0)
	layoutv1.SetLayout(widgets.NewQVBoxLayout())
	layoutv2 := widgets.NewQWidget(nil, 0)
	layoutv2.SetLayout(widgets.NewQVBoxLayout())
	layoutv2.SetMaximumWidth(800)
	layouth0.Layout().AddWidget(layoutv1)
	layouth0.Layout().AddWidget(layoutv2)

	layouth1 := widgets.NewQWidget(nil, 0)
	layouth1.SetLayout(widgets.NewQHBoxLayout())
	layouth2 := widgets.NewQWidget(nil, 0)
	layouth2.SetLayout(widgets.NewQHBoxLayout())
	layoutv1.Layout().AddWidget(layouth1)
	layoutv1.Layout().AddWidget(layouth2)

	layouthu := widgets.NewQWidget(nil, 0)
	layouthu.SetLayout(widgets.NewQHBoxLayout())
	layouthg := widgets.NewQWidget(nil, 0)
	layouthg.SetLayout(widgets.NewQHBoxLayout())

	layouthh := widgets.NewQWidget(nil, 0)
	layouthh.SetLayout(widgets.NewQHBoxLayout())
	layouthc := widgets.NewQWidget(nil, 0)
	layouthc.SetLayout(widgets.NewQHBoxLayout())
	layoutho := widgets.NewQWidget(nil, 0)
	layoutho.SetLayout(widgets.NewQHBoxLayout())

	layouthra := widgets.NewQWidget(nil, 0)
	layouthra.SetLayout(widgets.NewQHBoxLayout())
	layouthrag := widgets.NewQWidget(nil, 0)
	layouthrag.SetLayout(widgets.NewQHBoxLayout())
	layouthor := widgets.NewQWidget(nil, 0)
	layouthor.SetLayout(widgets.NewQHBoxLayout())

	layoutha := widgets.NewQWidget(nil, 0)
	layoutha.SetLayout(widgets.NewQHBoxLayout())
	layouthb := widgets.NewQWidget(nil, 0)
	layouthb.SetLayout(widgets.NewQHBoxLayout())

	qboxMain := widgets.NewQGroupBox(nil)
	qboxMain.SetTitle("Main")
	qboxMain.SetLayout(widgets.NewQVBoxLayout())

	qboxUser := widgets.NewQGroupBox(nil)
	qboxUser.SetTitle("User(s)")
	qboxUser.SetLayout(widgets.NewQVBoxLayout())

	qboxGroup := widgets.NewQGroupBox(nil)
	qboxGroup.SetTitle("Group(s)")
	qboxGroup.SetLayout(widgets.NewQVBoxLayout())

	qboxHost := widgets.NewQGroupBox(nil)
	qboxHost.SetTitle("Host(s)")
	qboxHost.SetMaximumWidth(200)
	qboxHost.SetLayout(widgets.NewQVBoxLayout())

	qboxCommand := widgets.NewQGroupBox(nil)
	qboxCommand.SetTitle("Command(s)")
	qboxCommand.SetLayout(widgets.NewQVBoxLayout())

	qboxOption := widgets.NewQGroupBox(nil)
	qboxOption.SetTitle("Option(s)")
	qboxOption.SetLayout(widgets.NewQVBoxLayout())

	qboxResume := widgets.NewQGroupBox(nil)
	qboxResume.SetTitle("Resume/Info")
	qboxResume.SetLayout(widgets.NewQVBoxLayout())

	layouth1.Layout().AddWidget(qboxMain)
	layouth1.Layout().AddWidget(qboxUser)
	layouth1.Layout().AddWidget(qboxGroup)

	layouth2.Layout().AddWidget(qboxHost)
	layouth2.Layout().AddWidget(qboxCommand)
	layouth2.Layout().AddWidget(qboxOption)

	layoutv2.Layout().AddWidget(qboxResume)

	window.SetCentralWidget(layouth0)

	admInput = widgets.NewQComboBox(nil)
	admInput.AddItems([]string{"VIEW", "ADD", "DELETE", "MODIFY"})
	admInput.ConnectCurrentIndexChanged(func(index int) {
		LoadSudoers()
		switch admInput.CurrentText() {
		case "VIEW":
			clearView()
			roleName.SetDisabled(true)
			roleName.SetStyleSheet("")
			roleNameDM.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
			roleNameDM.SetEnabled(true)
			changeNumber.SetDisabled(true)
			changeNumber.SetStyleSheet("")
			userSelect.SetDisabled(true)
			groupSelect.SetDisabled(true)
			hostSelect.SetDisabled(true)
			commandName.SetDisabled(true)
			optionName.SetDisabled(true)
			runAsName.SetReadOnly(true)
			runAsGName.SetReadOnly(true)
			sudoOrder.SetReadOnly(true)
			notAfter.SetReadOnly(true)
			notBefore.SetReadOnly(true)
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())
			}
		case "ADD":
			clearView()
			roleName.SetEnabled(true)
			roleName.SetReadOnly(false)
			roleName.SetStyleSheet("border: 2px solid red")
			roleNameDM.SetDisabled(true)
			changeNumber.SetEnabled(true)
			changeNumber.SetStyleSheet("border: 2px solid red")
			userSelect.SetEnabled(true)
			groupSelect.SetEnabled(true)
			hostSelect.SetEnabled(true)
			optionName.SetEnabled(true)
			commandName.SetEnabled(true)
			runAsName.SetReadOnly(false)
			runAsGName.SetReadOnly(false)
			sudoOrder.SetReadOnly(false)
			notAfter.SetReadOnly(false)
			notBefore.SetReadOnly(false)
			LoadUsers()
			LoadGroups()
			hostSelectDM.AddItem("ALL")
			hostSelectDM.SetCurrentRow(0)
			commandNameDM.AddItem("ALL")
			commandNameDM.SetCurrentRow(0)
			optionNameDM.AddItem("!authenticate")
			optionNameDM.SetCurrentRow(0)
			notAfter.SetDateTime(core.NewQDateTime().FromString2(time.Now().AddDate(0, 0, -2).Format("20060102"), "yyyyMMdd"))
			notBefore.SetDateTime(core.NewQDateTime().FromString2(time.Now().AddDate(0, 0, -2).Format("20060102"), "yyyyMMdd"))
		case "DELETE":
			clearView()
			roleName.SetDisabled(true)
			roleName.SetStyleSheet("")
			roleNameDM.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
			roleNameDM.SetEnabled(true)
			changeNumber.SetDisabled(false)
			changeNumber.SetStyleSheet("border: 2px solid red")
			userSelect.SetDisabled(true)
			groupSelect.SetDisabled(true)
			hostSelect.SetDisabled(true)
			commandName.SetDisabled(true)
			optionName.SetDisabled(true)
			runAsName.SetReadOnly(true)
			runAsGName.SetReadOnly(true)
			sudoOrder.SetReadOnly(true)
			notAfter.SetReadOnly(true)
			notBefore.SetReadOnly(true)
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())
			}
		case "MODIFY":
			clearView()
			roleName.SetEnabled(true)
			roleName.SetReadOnly(true)
			roleName.SetStyleSheet("")
			roleNameDM.SetEnabled(true)
			changeNumber.SetEnabled(true)
			changeNumber.SetStyleSheet("border: 2px solid red")
			userSelect.SetEnabled(true)
			groupSelect.SetEnabled(true)
			hostSelect.SetEnabled(true)
			optionName.SetEnabled(true)
			commandName.SetEnabled(true)
			runAsName.SetReadOnly(false)
			runAsGName.SetReadOnly(false)
			sudoOrder.SetReadOnly(false)
			notAfter.SetReadOnly(false)
			notBefore.SetReadOnly(false)
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())
				LoadUsersSpecial()
				LoadGroupsSpecial()
				updateResumeMOD()
			}
		}
	})
	qboxMain.Layout().AddWidget(admInput)

	roleName = widgets.NewQLineEdit(nil)
	roleName.SetPlaceholderText("Role Name ...")
	roleName.ConnectTextChanged(func(text string) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
			if roleName.Text() == "" {
				roleName.SetStyleSheet("border: 2px solid red")
			} else {
				roleName.SetStyleSheet("")
			}
		}
	})
	qboxMain.Layout().AddWidget(roleName)

	roleNameDM = widgets.NewQListWidget(nil)
	roleNameDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	LoadSudoers()
	roleNameDM.ConnectItemSelectionChanged(func() {
		switch admInput.CurrentText() {
		case "VIEW":
			clearView()
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())
			}
		case "DELETE":
			clearView()
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())

			}
		case "MODIFY":
			clearView()
			if roleNameDM.CurrentItem().Text() != "" {
				roleName.SetText(roleNameDM.CurrentItem().Text())
				LoadSudo(roleNameDM.CurrentItem().Text())
				LoadGroupsSpecial()
				LoadUsersSpecial()
				updateResumeMOD()
			}
		}
	})
	qboxMain.Layout().AddWidget(roleNameDM)

	changeNumber = widgets.NewQLineEdit(nil)
	changeNumber.SetPlaceholderText("Approved Service Manager change number ...")
	changeNumber.ConnectTextChanged(func(text string) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
			if changeNumber.Text() == "" {
				changeNumber.SetStyleSheet("border: 2px solid red")
			} else {
				changeNumber.SetStyleSheet("")
			}
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
			if changeNumber.Text() == "" {
				changeNumber.SetStyleSheet("border: 2px solid red")
			} else {
				changeNumber.SetStyleSheet("")
			}
		} else if admInput.CurrentText() == "DELETE" {
			if changeNumber.Text() == "" {
				changeNumber.SetStyleSheet("border: 2px solid red")
			} else {
				changeNumber.SetStyleSheet("")
			}
		}
	})
	qboxMain.Layout().AddWidget(changeNumber)

	qboxUser.Layout().AddWidget(layouthu)
	userSelect = widgets.NewQLineEdit(nil)
	userSelect.SetPlaceholderText("User(s) to grant Role privileges ...")
	layouthu.Layout().AddWidget(userSelect)
	buttonu = widgets.NewQPushButton2("add", nil)
	buttonu.SetMaximumWidth(35)
	buttonu.ConnectClicked(func(bool) {
		userNameDM.AddItem(userSelect.Text())
		userSelect.SetText("")
	})
	layouthu.Layout().AddWidget(buttonu)
	userNameDM = widgets.NewQListWidget(nil)
	userNameDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	userNameDM.ConnectItemSelectionChanged(func() {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	qboxUser.Layout().AddWidget(userNameDM)

	qboxGroup.Layout().AddWidget(layouthg)
	groupSelect = widgets.NewQLineEdit(nil)
	groupSelect.SetPlaceholderText("Group(s) to grant Role privileges ...")
	layouthg.Layout().AddWidget(groupSelect)
	buttong = widgets.NewQPushButton2("add", nil)
	buttong.SetMaximumWidth(35)
	buttong.ConnectClicked(func(bool) {
		groupNameDM.AddItem(groupSelect.Text())
		groupSelect.SetText("")
	})
	layouthg.Layout().AddWidget(buttong)
	groupNameDM = widgets.NewQListWidget(nil)
	groupNameDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	groupNameDM.ConnectItemSelectionChanged(func() {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	qboxGroup.Layout().AddWidget(groupNameDM)

	hostSelect = widgets.NewQLineEdit(nil)
	hostSelect.SetPlaceholderText("Host(s) allowed ...")
	qboxHost.Layout().AddWidget(layouthh)
	layouthh.Layout().AddWidget(hostSelect)
	buttonh = widgets.NewQPushButton2("add", nil)
	buttonh.SetMaximumWidth(35)
	buttonh.ConnectClicked(func(bool) {
		hostSelectDM.AddItem(hostSelect.Text())
		hostSelect.SetText("")
	})
	layouthh.Layout().AddWidget(buttonh)

	hostSelectDM = widgets.NewQListWidget(nil)
	hostSelectDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	hostSelectDM.ConnectItemSelectionChanged(func() {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	qboxHost.Layout().AddWidget(hostSelectDM)

	commandName = widgets.NewQLineEdit(nil)
	commandName.SetPlaceholderText("Command(s) allowed to run ...")
	qboxCommand.Layout().AddWidget(layouthc)

	layouthc.Layout().AddWidget(commandName)
	buttonc = widgets.NewQPushButton2("add", nil)
	buttonc.SetMaximumWidth(35)
	buttonc.ConnectClicked(func(bool) {
		commandNameDM.AddItem(commandName.Text())
		commandName.SetText("")
	})
	layouthc.Layout().AddWidget(buttonc)

	commandNameDM = widgets.NewQListWidget(nil)
	commandNameDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	commandNameDM.ConnectItemSelectionChanged(func() {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	qboxCommand.Layout().AddWidget(commandNameDM)

	optionName = widgets.NewQLineEdit(nil)
	optionName.SetPlaceholderText("Option(s) ...")
	qboxOption.Layout().AddWidget(layoutho)

	layoutho.Layout().AddWidget(optionName)
	buttono = widgets.NewQPushButton2("add", nil)
	buttono.SetMaximumWidth(35)
	buttono.ConnectClicked(func(bool) {
		optionNameDM.AddItem(optionName.Text())
		optionName.SetText("")
	})
	layoutho.Layout().AddWidget(buttono)

	optionNameDM = widgets.NewQListWidget(nil)
	optionNameDM.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	optionNameDM.ConnectItemSelectionChanged(func() {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	qboxOption.Layout().AddWidget(optionNameDM)

	qboxOption.Layout().AddWidget(layouthra)
	qboxOption.Layout().AddWidget(layouthrag)
	qboxOption.Layout().AddWidget(layouthor)

	runAsLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	runAsLabel.SetText("Run As User:")
	runAsLabel.SetMaximumWidth(90)
	layouthra.Layout().AddWidget(runAsLabel)

	runAsName = widgets.NewQLineEdit(nil)
	runAsName.SetPlaceholderText("Sudo Run As User ...")
	runAsName.ConnectTextChanged(func(text string) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	layouthra.Layout().AddWidget(runAsName)

	runAsGLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	runAsGLabel.SetText("Run As Group:")
	runAsGLabel.SetMaximumWidth(95)
	layouthrag.Layout().AddWidget(runAsGLabel)

	runAsGName = widgets.NewQLineEdit(nil)
	runAsGName.SetPlaceholderText("Sudo Run As Group ...")
	runAsGName.ConnectTextChanged(func(text string) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	layouthrag.Layout().AddWidget(runAsGName)

	sudoOrderLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	sudoOrderLabel.SetText("Order:")
	sudoOrderLabel.SetMaximumWidth(70)
	layouthor.Layout().AddWidget(sudoOrderLabel)

	sudoOrder = widgets.NewQLineEdit(nil)
	sudoOrder.SetPlaceholderText("Sudo Order ...")
	sudoOrder.ConnectTextChanged(func(text string) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	layouthor.Layout().AddWidget(sudoOrder)

	qboxOption.Layout().AddWidget(layouthb)
	qboxOption.Layout().AddWidget(layoutha)

	notBeforeLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	notBeforeLabel.SetText("not Before:")
	notBeforeLabel.SetMaximumWidth(70)
	layouthb.Layout().AddWidget(notBeforeLabel)

	notBefore = widgets.NewQDateTimeEdit(nil)
	notBefore.SetDisplayFormat("dd.MM.yyyy HH:mm")
	notBefore.SetCalendarPopup(true)
	notBefore.ConnectDateChanged(func(date *core.QDate) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	layouthb.Layout().AddWidget(notBefore)

	notAfterLabel := widgets.NewQLabel(nil, core.Qt__Widget)
	notAfterLabel.SetText("not After:")
	notAfterLabel.SetMaximumWidth(70)
	layoutha.Layout().AddWidget(notAfterLabel)

	notAfter = widgets.NewQDateTimeEdit(nil)
	notAfter.SetDisplayFormat("dd.MM.yyyy HH:mm")
	notAfter.SetCalendarPopup(true)
	notAfter.ConnectDateChanged(func(date *core.QDate) {
		if admInput.CurrentText() == "ADD" {
			updateResumeADD()
		} else if admInput.CurrentText() == "MODIFY" {
			updateResumeMOD()
		}
	})
	layoutha.Layout().AddWidget(notAfter)

	resume = widgets.NewQTextEdit(nil)
	qboxResume.Layout().AddWidget(resume)

	buttonExecute = widgets.NewQPushButton2("Execute", nil)
	buttonExecute.ConnectClicked(func(bool) {
		switch admInput.CurrentText() {
		case "VIEW":
		case "ADD":
			if roleName.Text() == "" || changeNumber.Text() == "" {
				widgets.QMessageBox_Warning(nil, "Warning", "Role Name and Change Number can not be empty!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			} else {
				reply := widgets.QMessageBox_Question(nil, "ADD SUDO Role", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__No)
				if reply == widgets.QMessageBox__Yes {
					var isOk bool
					password := widgets.QInputDialog_GetText(nil, "Input password", "Password", widgets.QLineEdit__Password, "", &isOk, core.Qt__Dialog, core.Qt__ImhNone)

					if isOk {
						if _, err := os.Stat(Config.Openldap.DiffFiles); os.IsNotExist(err) {
							err = os.MkdirAll(Config.Openldap.DiffFiles, 0755)
							if err != nil {
								panic(err)
							}
						}
						f, err := os.Create(Config.Openldap.DiffFiles + "/" + time.Now().Format("20060102150405") + "-add-" + changeNumber.Text() + ".ldif")
						check(err)
						defer f.Close()
						f.WriteString(resume.ToPlainText())
						f.Sync()
						fileToRun := "ldapmodify -a -x -D " + Config.Openldap.DN + " -w " + password + " -H ldaps://" + Config.Openldap.Server + ":" + Config.Openldap.Port + " -f " + f.Name()
						test := exec.Command("/bin/bash", "-c", fileToRun)
						err = test.Run()

						if err != nil {
							fmt.Println(err.Error())
							showWarnigMod()
							os.Rename(f.Name(), f.Name()+"-ERROR")
						} else {
							admInput.SetCurrentIndex(0)
						}
					}
				} else if reply == widgets.QMessageBox__No {
					fmt.Println("no add")
				}
			}

		case "DELETE":
			if changeNumber.Text() == "" {
				widgets.QMessageBox_Warning(nil, "Warning", "Change Number can not be empty!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			} else {
				reply := widgets.QMessageBox_Question(nil, "DELETE SUDO Role", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__No)
				if reply == widgets.QMessageBox__Yes {
					var isOk bool
					password := widgets.QInputDialog_GetText(nil, "Input password", "Password", widgets.QLineEdit__Password, "", &isOk, core.Qt__Dialog, core.Qt__ImhNone)

					if isOk {
						if _, err := os.Stat(Config.Openldap.DiffFiles); os.IsNotExist(err) {
							err = os.MkdirAll(Config.Openldap.DiffFiles, 0755)
							if err != nil {
								panic(err)
							}
						}
						f, err := os.Create(Config.Openldap.DiffFiles + "/" + time.Now().Format("20060102150405") + "-del-" + changeNumber.Text() + ".ldif")
						check(err)
						defer f.Close()
						f.WriteString(resume.ToPlainText())
						f.Sync()
						fileToRun := "ldapmodify -x -D " + Config.Openldap.DN + " -w " + password + " -H ldaps://" + Config.Openldap.Server + ":" + Config.Openldap.Port + " -f " + f.Name()
						test := exec.Command("/bin/bash", "-c", fileToRun)
						err = test.Run()

						if err != nil {
							fmt.Println(err.Error())
							showWarnigMod()
							os.Rename(f.Name(), f.Name()+"-ERROR")
						} else {
							roleNameDM.CurrentItem().SetSelected(false)
							admInput.SetCurrentIndex(0)
						}
					}
				} else if reply == widgets.QMessageBox__No {
					fmt.Println("no delete")
				}
			}

		case "MODIFY":
			if changeNumber.Text() == "" {
				widgets.QMessageBox_Warning(nil, "Warning", "Change Number can not be empty!", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
			} else {
				reply := widgets.QMessageBox_Question(nil, "MODIFY SUDO Role", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__No)
				if reply == widgets.QMessageBox__Yes {
					var isOk bool
					password := widgets.QInputDialog_GetText(nil, "Input password", "Password", widgets.QLineEdit__Password, "", &isOk, core.Qt__Dialog, core.Qt__ImhNone)

					if isOk {
						if _, err := os.Stat(Config.Openldap.DiffFiles); os.IsNotExist(err) {
							err = os.MkdirAll(Config.Openldap.DiffFiles, 0755)
							if err != nil {
								panic(err)
							}
						}
						f, err := os.Create(Config.Openldap.DiffFiles + "/" + time.Now().Format("20060102150405") + "-mod-" + changeNumber.Text() + ".ldif")
						check(err)
						defer f.Close()
						f.WriteString(resume.ToPlainText())
						f.Sync()
						fileToRun := "ldapmodify -x -D " + Config.Openldap.DN + " -w " + password + " -H ldaps://" + Config.Openldap.Server + ":" + Config.Openldap.Port + " -f " + f.Name()
						test := exec.Command("/bin/bash", "-c", fileToRun)
						err = test.Run()

						if err != nil {
							fmt.Println(err.Error())
							showWarnigMod()
							os.Rename(f.Name(), f.Name()+"-ERROR")
						} else {
							admInput.SetCurrentIndex(0)
						}
					}
				} else if reply == widgets.QMessageBox__No {
					fmt.Println("no modify")
				}
			}
		}
	})
	layoutv2.Layout().AddWidget(buttonExecute)

	buttonExit = widgets.NewQPushButton2("Exit", nil)
	buttonExit.ConnectClicked(func(bool) {
		reply := widgets.QMessageBox_Question(nil, "Exit", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__No)
		if reply == widgets.QMessageBox__Yes {
			app.Quit()
		}
	})
	layoutv2.Layout().AddWidget(buttonExit)

	initGUI()

	window.Show()
	app.Exec()
}

func initGUI() {
	roleName.SetDisabled(true)
	roleNameDM.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	changeNumber.SetDisabled(true)
	userSelect.SetDisabled(true)
	groupSelect.SetDisabled(true)
	hostSelect.SetDisabled(true)
	commandName.SetDisabled(true)
	optionName.SetDisabled(true)
	runAsName.SetReadOnly(true)
	runAsGName.SetReadOnly(true)
	sudoOrder.SetReadOnly(true)
	notAfter.SetReadOnly(true)
	notBefore.SetReadOnly(true)
	resume.SetTextInteractionFlags(core.Qt__TextSelectableByMouse)
}

func clearView() {
	roleName.SetText("")
	changeNumber.SetText("")
	userSelect.SetText("")
	userNameDM.Clear()
	groupSelect.SetText("")
	groupNameDM.Clear()
	hostSelect.SetText("")
	hostSelectDM.Clear()
	commandName.SetText("")
	commandNameDM.Clear()
	optionName.SetText("")
	optionNameDM.Clear()
	runAsName.SetText("")
	runAsGName.SetText("")
	sudoOrder.SetText("")
	notAfter.SetDateTime(core.NewQDateTime().FromString2("20000101000000.000", "yyyyMMddHHmmss.zzz"))
	notBefore.SetDateTime(core.NewQDateTime().FromString2("20000101000000.000", "yyyyMMddHHmmss.zzz"))
	resume.SetText("")
}

func updateResumeMOD() {

	currentTime := time.Now()

	resume.SetHtml("<p style=\"margin:0\"><b><font color=\"DarkMagenta\">dn</font><font color=\"orange\">: </font></b> cn=" + roleName.Text() + "," + Config.Openldap.OUsudoers + "," + Config.Openldap.BaseDN + "</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">changetype</font><font color=\"orange\">: </font></b> modify</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>description</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">description</font><font color=\"orange\">: </font></b>" + currentTime.Format("200601021504") + " - mod - " + changeNumber.Text() + "</p>")

	separator := 0
	after := false
	before := false
	total := userNameDM.Count()

	for index := 0; index < total; index++ {
		if userNameDM.Item(index).Text() == "---" {
			separator = index
		} else {
			if userNameDM.Item(index).IsSelected() {
				if separator > 0 {
					after = true
				} else {
					before = true
				}
			}
		}
	}

	if before && separator > 0 {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoUser</p>")
		for i := 0; i < separator; i++ {
			if userNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>" + userNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	if after || (before && separator == 0) {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoUser</p>")
		for i := separator; i < total; i++ {
			if userNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>" + userNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	separator = 0
	after = false
	before = false
	total = groupNameDM.Count()

	for index := 0; index < total; index++ {
		if groupNameDM.Item(index).Text() == "---" {
			separator = index
		} else {
			if groupNameDM.Item(index).IsSelected() {
				if separator > 0 {
					after = true
				} else {
					before = true
				}
			}
		}
	}

	if before && separator > 0 {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoUser</p>")
		for i := 0; i < separator; i++ {
			if groupNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>" + groupNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	if after || (before && separator == 0) {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoUser</p>")
		for i := separator; i < total; i++ {
			if groupNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>" + groupNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	separator = 0
	after = false
	before = false
	total = hostSelectDM.Count()

	for index := 0; index < total; index++ {
		if hostSelectDM.Item(index).Text() == "---" {
			separator = index
		} else {
			if hostSelectDM.Item(index).IsSelected() {
				if separator > 0 {
					after = true
				} else {
					before = true
				}
			}
		}
	}

	if before && separator > 0 {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoHost</p>")
		for i := 0; i < separator; i++ {
			if hostSelectDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoHost</font><font color=\"orange\">: </font></b>" + hostSelectDM.Item(i).Text() + "</p>")
			}
		}
	}

	if after || (before && separator == 0) {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoHost</p>")
		for i := separator; i < total; i++ {
			if hostSelectDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoHost</font><font color=\"orange\">: </font></b>" + hostSelectDM.Item(i).Text() + "</p>")
			}
		}
	}

	separator = 0
	after = false
	before = false
	total = commandNameDM.Count()

	for index := 0; index < total; index++ {
		if commandNameDM.Item(index).Text() == "---" {
			separator = index
		} else {
			if commandNameDM.Item(index).IsSelected() {
				if separator > 0 {
					after = true
				} else {
					before = true
				}
			}
		}
	}

	if before && separator > 0 {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoCommand</p>")
		for i := 0; i < separator; i++ {
			if commandNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoCommand</font><font color=\"orange\">: </font></b>" + commandNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	if after || (before && separator == 0) {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoCommand</p>")
		for i := separator; i < total; i++ {
			if commandNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoCommand</font><font color=\"orange\">: </font></b>" + commandNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	separator = 0
	after = false
	before = false
	total = optionNameDM.Count()

	for index := 0; index < total; index++ {
		if optionNameDM.Item(index).Text() == "---" {
			separator = index
		} else {
			if optionNameDM.Item(index).IsSelected() {
				if separator > 0 {
					after = true
				} else {
					before = true
				}
			}
		}
	}

	if before && separator > 0 {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoOption</p>")
		for i := 0; i < separator; i++ {
			if optionNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOption</font><font color=\"orange\">: </font></b>" + optionNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	if after || (before && separator == 0) {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoOption</p>")
		for i := separator; i < total; i++ {
			if optionNameDM.Item(i).IsSelected() {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOption</font><font color=\"orange\">: </font></b>" + optionNameDM.Item(i).Text() + "</p>")
			}
		}
	}

	if rau != "" {
		if runAsName.Text() != rau {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">replace</font><font color=\"orange\">: </font></b>sudoRunAsUser</p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsUser</font><font color=\"orange\">: </font></b>" + runAsName.Text() + "</p>")
		}
	} else if runAsName.Text() != "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoRunAsUser</p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsUser</font><font color=\"orange\">: </font></b>" + runAsName.Text() + "</p>")
	}

	if rag != "" {
		if runAsGName.Text() != rag {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">replace</font><font color=\"orange\">: </font></b>sudoRunAsGroup</p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsGroup</font><font color=\"orange\">: </font></b>" + runAsGName.Text() + "</p>")
		}
	} else if runAsGName.Text() != "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoRunAsGroup</p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsGroup</font><font color=\"orange\">: </font></b>" + runAsGName.Text() + "</p>")
	}

	if sdo != "" {
		if sudoOrder.Text() != sdo {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">replace</font><font color=\"orange\">: </font></b>sudoOrder</p>")
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOrder</font><font color=\"orange\">: </font></b>" + sudoOrder.Text() + "</p>")
		}
	} else if sudoOrder.Text() != "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoOrder</p>")
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOrder</font><font color=\"orange\">: </font></b>" + sudoOrder.Text() + "</p>")
	}

	if len(nbf) > 0 && notBefore.DateTime().ToString("yyyyMMddHHmm") != nbf[:12] {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		if nbf[:12] == "200001010000" {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoNotBefore</p>")
		} else if notBefore.DateTime().ToString("yyyyMMddHHmm") == "200001010000" {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoNotBefore</p>")
		} else {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">replace</font><font color=\"orange\">: </font></b>sudoNotBefore</p>")
		}
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoNotBefore</font><font color=\"orange\">: </font></b>" + notBefore.DateTime().ToString("yyyyMMddHHmmss") + "Z</p>")
	}

	if len(naf) > 0 && notAfter.DateTime().ToString("yyyyMMddHHmm") != naf[:12] {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">-</font><font color=\"orange\"></font></b></p>")
		if naf[:12] == "200001010000" {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">add</font><font color=\"orange\">: </font></b>sudoNotAfter</p>")
		} else if notAfter.DateTime().ToString("yyyyMMddHHmm") == "200001010000" {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">delete</font><font color=\"orange\">: </font></b>sudoNotAfter</p>")
		} else {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">replace</font><font color=\"orange\">: </font></b>sudoNotAfter</p>")
		}
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoNotAfter</font><font color=\"orange\">: </font></b>" + notAfter.DateTime().ToString("yyyyMMddHHmmss") + "Z</p>")
	}
}

func updateResumeADD() {

	if len(userNameDM.SelectedItems()) > 1 || len(groupNameDM.SelectedItems()) > 0 {
		for _, entry := range userNameDM.SelectedItems() {
			if entry.Text() == "ALL" {
				entry.SetSelected(false)
			}
		}
	} else if len(userNameDM.SelectedItems()) == 0 && len(groupNameDM.SelectedItems()) == 0 {
		userNameDM.Item(0).SetSelected(true)
	}

	currentTime := time.Now()

	resume.SetHtml("<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">dn</font><font color=\"orange\">: </font></b> cn=" + roleName.Text() + "," + Config.Openldap.OUsudoers + "," + Config.Openldap.BaseDN + "</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">changetype</font><font color=\"orange\">: </font></b> add</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">objectClass</font><font color=\"orange\">: </font></b>top</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">objectClass</font><font color=\"orange\">: </font></b>sudoRole</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">description</font><font color=\"orange\">: </font></b>" + currentTime.Format("200601021504") + " - add - " + changeNumber.Text() + "</p>")
	resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">cn</font><font color=\"orange\">: </font></b>" + roleName.Text() + "</p>")

	if len(userNameDM.SelectedItems()) > 0 {
		for _, entry := range userNameDM.SelectedItems() {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>" + entry.Text() + "</p>")
		}

	}

	if len(groupNameDM.SelectedItems()) > 0 {
		for _, entry := range groupNameDM.SelectedItems() {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b>%" + entry.Text() + "</p>")
		}

	}

	if len(hostSelectDM.SelectedItems()) > 0 {
		for _, entry := range hostSelectDM.SelectedItems() {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoHost</font><font color=\"orange\">: </font></b>" + entry.Text() + "</p>")
		}

	}

	if len(commandNameDM.SelectedItems()) > 0 {
		for _, entry := range commandNameDM.SelectedItems() {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoCommand</font><font color=\"orange\">: </font></b>" + entry.Text() + "</p>")
		}

	}

	if len(optionNameDM.SelectedItems()) > 0 {
		for _, entry := range optionNameDM.SelectedItems() {
			resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOption</font><font color=\"orange\">: </font></b>" + entry.Text() + "</p>")
		}

	}

	if runAsName.Text() == "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsUser</font><font color=\"orange\">: </font></b>ALL</p>")
	} else {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsUser</font><font color=\"orange\">: </font></b>" + runAsName.Text() + "</p>")
	}

	if runAsGName.Text() != "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoRunAsGroup</font><font color=\"orange\">: </font></b>" + runAsGName.Text() + "</p>")
	}

	if sudoOrder.Text() != "" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoOrder</font><font color=\"orange\">: </font></b>" + sudoOrder.Text() + "</p>")
	}

	if notBefore.DateTime().ToString("yyyyMMddHHmm")[:8] != time.Now().AddDate(0, 0, -2).Format("200601021504")[:8] {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoNotBefore</font><font color=\"orange\">: </font></b>" + notBefore.DateTime().ToString("yyyyMMddHHmmss") + "Z</p>")
	}

	if notAfter.DateTime().ToString("yyyyMMddHHmm")[:8] != time.Now().AddDate(0, 0, -2).Format("200601021504")[:8] {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">sudoNotAfter</font><font color=\"orange\">: </font></b>" + notAfter.DateTime().ToString("yyyyMMddHHmmss") + "Z</p>")
	}
}

func LoadSudo(sudoFilter string) {

	if sudoFilter == "" {
		return
	}

	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = "cn=" + sudoFilter + "," + Config.Openldap.OUsudoers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	resume.SetHtml("<p style=\"margin:0\"><b><font color=\"DarkMagenta\">dn</font><font color=\"orange\">: </font></b>" + baseDN + "</p>")

	if admInput.CurrentText() == "DELETE" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">changetype</font><font color=\"orange\">: </font></b> delete</p>")
		return
	}

	if admInput.CurrentText() == "MODIFY" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\"><b><font color=\"DarkMagenta\">changetype</font><font color=\"orange\">: </font></b> modify</p>")
	}

	if admInput.CurrentText() == "VIEW" {
		resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">cn</font><font color=\"orange\">: </font></b> " + sudoFilter + "</p>")
	}

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"sudoUser", "sudoCommand", "sudoHost", "sudoOption", "cn", "sudoOrder", "sudoRunAs", "sudoRunAsUser", "sudoRunAsGroup", "sudoNotBefore", "sudoNotAfter", "description"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		for _, entry2 := range entry.GetAttributeValues("description") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">description</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
		}
	}

	for _, entry := range sr.Entries {
		for _, entry2 := range entry.GetAttributeValues("sudoUser") {
			if entry2[:1] == "%" {
				groupNameDM.AddItem(entry2)

				if admInput.CurrentText() == "VIEW" {
					resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
				}

				baseDN = "cn=" + entry2[1:] + "," + Config.Openldap.OUgroups + "," + Config.Openldap.BaseDN
				searchRequest2 := ldap.NewSearchRequest(
					baseDN,
					ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
					filter,
					[]string{"memberUid", "gidNumber", "cn"},
					nil,
				)

				sr2, err := l.Search(searchRequest2)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					for _, entry3 := range sr2.Entries {
						for _, entry4 := range entry3.GetAttributeValues("memberUid") {
							if entry4 != "" {
								if admInput.CurrentText() == "VIEW" {
									resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">        <b><font color=\"DarkMagenta\">member</font><font color=\"orange\">: </font></b> " + entry4 + "</p>")
								}
							}
						}
					}
				}

			} else {
				if admInput.CurrentText() == "VIEW" {
					resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoUser</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
				}
				userNameDM.AddItem(entry2)
			}
		}
	}

	if admInput.CurrentText() == "MODIFY" {
		if userNameDM.Count() > 0 {
			userNameDM.AddItem("---")
			userNameDM.Item(userNameDM.Count() - 1).SetFlags(userNameDM.Item(userNameDM.Count()-1).Flags() & ^core.Qt__ItemIsSelectable)
		}

		if groupNameDM.Count() > 0 {
			groupNameDM.AddItem("---")
			groupNameDM.Item(groupNameDM.Count() - 1).SetFlags(groupNameDM.Item(groupNameDM.Count()-1).Flags() & ^core.Qt__ItemIsSelectable)
		}
	}

	for _, entry := range sr.Entries {
		for _, entry2 := range entry.GetAttributeValues("sudoHost") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoHost</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			hostSelectDM.AddItem(entry2)
		}
	}

	if admInput.CurrentText() == "MODIFY" {
		if hostSelectDM.Count() > 0 {
			hostSelectDM.AddItem("---")
			hostSelectDM.Item(hostSelectDM.Count() - 1).SetFlags(hostSelectDM.Item(hostSelectDM.Count()-1).Flags() & ^core.Qt__ItemIsSelectable)

		}
	}

	for _, entry := range sr.Entries {
		for _, entry2 := range entry.GetAttributeValues("sudoCommand") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoCommand</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			commandNameDM.AddItem(entry2)
		}
	}

	if admInput.CurrentText() == "MODIFY" {
		if commandNameDM.Count() > 0 {
			commandNameDM.AddItem("---")
			commandNameDM.Item(commandNameDM.Count() - 1).SetFlags(commandNameDM.Item(commandNameDM.Count()-1).Flags() & ^core.Qt__ItemIsSelectable)

		}
	}

	for _, entry := range sr.Entries {
		for _, entry2 := range entry.GetAttributeValues("sudoOption") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoOption</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			optionNameDM.AddItem(entry2)
		}
	}

	if admInput.CurrentText() == "MODIFY" {
		if optionNameDM.Count() > 0 {
			optionNameDM.AddItem("---")
			optionNameDM.Item(optionNameDM.Count() - 1).SetFlags(optionNameDM.Item(optionNameDM.Count()-1).Flags() & ^core.Qt__ItemIsSelectable)

		}
	}

	for _, entry := range sr.Entries {
		rau = ""
		for _, entry2 := range entry.GetAttributeValues("sudoRunAsUser") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoRunAsUser</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			runAsName.SetText(entry2)
			rau = entry2
		}
	}

	for _, entry := range sr.Entries {
		rag = ""
		for _, entry2 := range entry.GetAttributeValues("sudoRunAsGroup") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoRunAsGrup</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			runAsGName.SetText(entry2)
			rag = entry2
		}
	}

	for _, entry := range sr.Entries {
		sdo = ""
		for _, entry2 := range entry.GetAttributeValues("sudoOrder") {
			if admInput.CurrentText() == "VIEW" {
				resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoOrder</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
			}
			sudoOrder.SetText(entry2)
			sdo = entry2
		}
	}

	for _, entry := range sr.Entries {

		if len(entry.GetAttributeValues("sudoNotBefore")) > 0 {
			for _, entry2 := range entry.GetAttributeValues("sudoNotBefore") {
				if admInput.CurrentText() == "VIEW" {
					resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoNotBefore</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
				}

				nbf = entry2

				if len(entry2) == 15 {
					notBefore.SetEnabled(true)
					notBefore.SetDateTime(core.NewQDateTime().FromString2(entry2[:14], "yyyyMMddHHmmss"))
				} else if len(entry2) == 19 {
					notBefore.SetEnabled(true)
					notBefore.SetDateTime(core.NewQDateTime().FromString2(entry2[:18], "yyyyMMddHHmmss.zzz"))
				} else {
					notBefore.SetDateTime(core.NewQDateTime().FromString2("20000101000000.000", "yyyyMMddHHmmss.zzz"))
					notBefore.SetEnabled(false)
				}
			}
		} else {
			nbf = "20000101000000.000"
		}
	}

	for _, entry := range sr.Entries {
		if len(entry.GetAttributeValues("sudoNotAfter")) > 0 {
			for _, entry2 := range entry.GetAttributeValues("sudoNotAfter") {
				if admInput.CurrentText() == "VIEW" {
					resume.SetHtml(resume.ToHtml() + "<p style=\"margin:0\">    <b><font color=\"DarkMagenta\">sudoNotAfter</font><font color=\"orange\">: </font></b> " + entry2 + "</p>")
				}

				naf = entry2

				if len(entry2) == 15 {
					notAfter.SetEnabled(true)
					notAfter.SetDateTime(core.NewQDateTime().FromString2(entry2[:14], "yyyyMMddHHmmss"))
				} else if len(entry2) == 19 {
					notAfter.SetEnabled(true)
					notAfter.SetDateTime(core.NewQDateTime().FromString2(entry2[:18], "yyyyMMddHHmmss.zzz"))
				} else {
					notAfter.SetDateTime(core.NewQDateTime().FromString2("20000101000000.000", "yyyyMMddHHmmss.zzz"))
					notAfter.SetEnabled(false)
				}
			}
		} else {
			naf = "20000101000000.000"
		}
	}
}

func LoadSudoers() {
	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUsudoers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	roleNameDM.Clear()

	for _, entry := range sr.Entries {
		roleNameDM.AddItem(entry.GetAttributeValue("cn"))
	}
}

func LoadUsers() {
	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUusers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	userNameDM.AddItem("ALL")

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("cn") != "" {
			userNameDM.AddItem(entry.GetAttributeValue("cn"))
		}
	}

	userNameDM.SetCurrentRow(0)

}

func LoadUsersSpecial() {
	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUusers + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	if userNameDM.Count() > 0 {
		index := 0
		longi := userNameDM.Count()

		for index < longi && userNameDM.Item(index).Text() != "ALL" {
			index = index + 1
		}
		if index == longi {
			userNameDM.AddItem("ALL")
		}
	} else {
		userNameDM.AddItem("ALL")
	}

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("cn") != "" {

			if userNameDM.Count() > 0 {
				index := 0
				longi := userNameDM.Count()

				for index < longi && userNameDM.Item(index).Text() != entry.GetAttributeValue("cn") {
					index = index + 1
				}
				if index == longi {
					userNameDM.AddItem(entry.GetAttributeValue("cn"))
				}
			} else {
				userNameDM.AddItem(entry.GetAttributeValue("cn"))
			}
		}
	}
}

func LoadGroups() {
	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUgroups + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("cn") != "" {
			groupNameDM.AddItem(entry.GetAttributeValue("cn"))
		}
	}
}

func LoadGroupsSpecial() {
	var ldapServer = Config.Openldap.Server
	var ldapTLSPort, _ = strconv.Atoi(Config.Openldap.Port)
	var baseDN = Config.Openldap.OUgroups + "," + Config.Openldap.BaseDN
	var filter = "(&(objectClass=*))"

	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapTLSPort), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer l.Close()

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		//log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		if entry.GetAttributeValue("cn") != "" {

			if groupNameDM.Count() > 0 {
				index := 0
				longi := groupNameDM.Count()

				for index < longi && groupNameDM.Item(index).Text() != ("%"+entry.GetAttributeValue("cn")) {
					index = index + 1
				}
				if index == longi {
					groupNameDM.AddItem("%" + entry.GetAttributeValue("cn"))
				}
			} else {
				groupNameDM.AddItem("%" + entry.GetAttributeValue("cn"))
			}
		}
	}
}

type OpenldapConfig struct {
	Server    string
	Port      string
	BaseDN    string
	OUusers   string
	OUgroups  string
	OUconfig  string
	OUsudoers string
	DN        string
	PS        string
	DiffFiles string
}

// LxldapConfig lxldap configuration data
type LxldapConfig struct {
	Openldap OpenldapConfig `mapstructure:"Openldap"`
}

// Config to storage
var Config *LxldapConfig

// InitConfig initialize the configuration
func InitConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".lxldap" (without extension).
	viper.AddConfigPath(home)
	viper.AddConfigPath("./config/")
	viper.SetConfigName(".lxldapg")

	// Enable environment variables
	// ex.: LXLDAPG_LDAP_PORT=8000
	viper.SetEnvPrefix("LXLDAPG")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic("Unable to unmarshal config")
	}
}
