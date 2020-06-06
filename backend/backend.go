package backend

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/wailsapp/wails"
)

// Backend struct is used to execute processes behind the scenes
type Backend struct {
	File       string
	contents   string
	runtime    *wails.Runtime
	log        *wails.CustomLogger
	fullscreen bool
}

// WailsInit intializes the Backend
func (a *Backend) WailsInit(runtime *wails.Runtime) error {
	a.log = runtime.Log.New("Backend")
	a.runtime = runtime

	runtime.Events.On("file-changed", a.SaveFile)

	return nil
}

// SaveFile is invoked whtne the file-save event triggers
func (a *Backend) SaveFile(data ...interface{}) {
	if a.File == "" {
		return
	}

	contents := fmt.Sprintf("%v", data[0])
	decodedContents, err := base64.StdEncoding.DecodeString(contents)
	if err != nil {
		a.ErrorNotification("unable to decode data: " + err.Error())
		return
	}

	err = ioutil.WriteFile(a.File, []byte(decodedContents), 0644)
	if err != nil {
		a.ErrorNotification("unable to write file: " + err.Error())
		return
	}

}

// ToggleFullScreen toggles fullscreen
func (a *Backend) ToggleFullScreen() {
	if a.fullscreen {
		a.runtime.Window.UnFullscreen()
		a.fullscreen = false
	} else {
		a.runtime.Window.Fullscreen()
		a.fullscreen = true
	}
}

// SelectFile opens the selected file
func (a *Backend) SelectFile() {

	// Get a file
	a.File = a.runtime.Dialog.SelectFile()
	if a.File == "" {
		return
	}
	a.OpenFile()
}

// OpenFile opens the selected file
func (a *Backend) OpenFile() {
	// Read the file
	contents, err := ioutil.ReadFile(a.File)
	if err != nil {
		a.ErrorNotification("unable to open file: " + err.Error())
		return
	}

	encodedContents := base64.StdEncoding.EncodeToString(contents)

	// Open the file in the editor
	a.runtime.Window.SetTitle(filepath.Base(a.File))
	a.runtime.Events.Emit("file-opened", string(encodedContents))
}

// GetDefaultContent returns the inital content in the editor
func (a *Backend) GetDefaultContent() string {
	if a.File == "" {
		return base64.StdEncoding.EncodeToString(defaultContents)
	} else {
		a.OpenFile()
		return ""
	}
}

// OpenURL opens a URL in the browser
func (a *Backend) OpenURL(url string) {
	a.runtime.Browser.OpenURL(url)
}

// ErrorNotification emits a viewable error to the frontend app
func (a *Backend) ErrorNotification(s string) {
	a.log.Error(s)
	a.runtime.Events.Emit("error-notification", s)
}
