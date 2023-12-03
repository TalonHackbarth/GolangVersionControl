package vcs

import (
	"GoVCS/pkg/Logging"
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var Log = Logging.Log{}

var TrackedItems = make([]TrackedItem, 0)
var PreviousItems = make([]TrackedItem, 0)
var CurrentCommit = make([]CommitItem, 0)

type TrackedItem struct {
	Name     string
	Contents [][]byte
}

type CommitItem struct {
	Name    string
	Changes map[int][]byte
}

func InitRepo(path string) {
	Log.Info("Initializing Repository")

	_, existErr := os.Stat(filepath.Join(path, ".gvc"))
	if existErr != nil {
		dirErr := os.Mkdir(filepath.Join(path, ".gvc"), os.ModeDir)
		if dirErr != nil {
			Log.Error(dirErr.Error())
		}
		Log.Info("Created Directory: " + filepath.Join(path, ".gvc"))
	}

}

func AddItems(path string, files ...string) {

	if files == nil {
		Log.Info("Adding . files")
		err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				f := strings.TrimPrefix(p, path)
				if strings.HasPrefix(f, "\\.gvc") {
					return nil
				}
				Track(path, f, &TrackedItems)
			}
			return nil
		})
		if err != nil {
			return
		}
	} else {
		for _, f := range files {
			Log.Trace("Tracking: " + f)
			Track(path, f, &TrackedItems)
		}
	}

	Write(path, "STAGING", TrackedItems)

}

func Commit(path string, tracked *[]TrackedItem) {
	// Add stuff for first commit
	if _, existErr := os.Stat(filepath.Join(path, "\\.gvc\\BASE.json")); existErr != nil {
		Log.Info("Setting up for first commit")
		LoadStaged(path, &PreviousItems)
		Write(path, "BASE", PreviousItems)
		RemoveStaged(path)
		return
	}

	LoadPrevious(path, &PreviousItems)

	for _, baseItem := range PreviousItems {
		var changes = make(map[int][]byte)
		var t, found = SearchTracked(*tracked, baseItem.Name)
		if found != true {
			Log.Info("Item Not Found in current tracked items. Deletion Assumed")
			return
		}

		for i, line := range t.Contents {
			if bytes.Equal(line, baseItem.Contents[i]) {
				continue
			}
			changes[i] = line
		}
		if len(changes) != 0 {
			CurrentCommit = append(CurrentCommit, CommitItem{t.Name, changes})
		}

	}

	if len(CurrentCommit) == 0 {
		fmt.Println("Current Branch up to date. Skipping Commit...")
		return
	}

	PreviewCommit()

	var commitMessage string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Commit Message: ")
	if scanner.Scan() {
		commitMessage = scanner.Text()
	}
	if len(commitMessage) == 0 {
		fmt.Println("Empty Commit Message. Aborting Commit")
		return
	}
	Log.Info("Committing : " + commitMessage)

	// Write Commit to File

	RemoveStaged(path)
}

func PreviewCommit() {
	for _, item := range CurrentCommit {
		fmt.Println("\033[36m", "Changes In <"+item.Name+">:", "\033[0m")
		for i, change := range item.Changes {
			prev, _ := SearchTracked(PreviousItems, item.Name)
			fmt.Println(" - [i]: ", "\033[31m", string(prev.Contents[i]), " -> ", "\033[32m", string(change), "\033[0m")
		}
	}
}
