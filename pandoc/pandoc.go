package pandoc

import (
	"fmt"
	"log"
	"os/exec"

	"gopkg.in/fsnotify.v1"
)

// RunPandocListener takes a directory to listen to, then should print the
// different changes that occur within it
func RunPandocListener(directory string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Could not make a watcher")
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				handleFileChange(ev)
			case err := <-watcher.Errors:
				log.Fatal("Error", err)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		log.Fatal("Could not watch the directory: " + directory)
	}
	<-done
}

// handFileChange handles the file listener event, checks if its on a *.md file
// and is the final change (mac makes various different file changes when rewriting
// a file using MacVim.
func handleFileChange(event fsnotify.Event) {
	fmt.Println(event)
}

// CompileAndRefresh recompiles the given *.md file into *.pdf, refocuses
// Preview.app so that it picks up the changes on disk, then refocuses
// MacVim.app to continue editing.
func CompileAndRefresh(baseFilename string) {
	var err error
	err = compileMarkdownToPdf(baseFilename)
	if err != nil {
		log.Fatal("Could not compile markdown to pdf using base filename: " + baseFilename)
	}
	err = openPreview(baseFilename)
	if err != nil {
		log.Fatal("Could not open Preview")
	}
	err = openMacVim()
	if err != nil {
		log.Fatal("Could not open MacVim")
	}
}

// FindFile is given a filename, then it attempts to find where that file is
// to return a full path. It first tries just the filename, then the current
// working directory plus the filename. If the file can't be found, return
// an error
func FindFile(baseFilename string) error {
	return nil
}

// compileMarkdownToPdf takes in the baseFilename, then compiles the *.md
// file to *.pdf
func compileMarkdownToPdf(baseFilename string) error {
	pandocPath, err := exec.LookPath("pandoc")
	if err != nil {
		log.Fatal("Could not find an installation of pandoc")
	}
	input := fmt.Sprintf("%s.md", baseFilename)
	output := fmt.Sprintf("%s.pdf", baseFilename)
	cmd := exec.Command(pandocPath, input, "-o", output)
	return cmd.Run()

}

// openPreview uses mac's command open to refocus Preview.app
// (or open the file if its not open)
func openPreview(baseFilename string) error {
	openPath, err := exec.LookPath("open")
	if err != nil {
		log.Fatal("Could not find an installation of open")
	}
	file := fmt.Sprintf("%s.pdf", baseFilename)
	cmd := exec.Command(openPath, file)
	return cmd.Run()
}

// openMacVim uses mac's command open to refocus MacVim.app
func openMacVim() error {
	openPath, err := exec.LookPath("open")
	if err != nil {
		log.Fatal("Could not find an installation of open")
	}
	cmd := exec.Command(openPath, "-a", "MacVim")
	return cmd.Run()
}
