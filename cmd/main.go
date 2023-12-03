package main

import (
	"GoVCS/pkg/Logging"
	"GoVCS/pkg/vcs"
)

var Log = Logging.Log{LogLevel: Logging.Info}

func main() { // TODO: CLI Setup
	Log.Info("GoVCS Started")
	var testPath = "C:\\dev\\samples\\vcs_test"
	// vcs.InitRepo(testPath)
	// vcs.AddItems("C:\\dev\\samples\\vcs_test", "docA.txt")

	vcs.AddItems(testPath)

	// TODO: Commit file ~ Tracks a diff for each file for each commit
	vcs.Commit(testPath, &vcs.TrackedItems)

	// TODO: Config that stores branch name/info.
	/*
			---
			config:
			- activeBranch: main
			  mergeMethod: overwrite
			  ...
			  branches:
			    - random_branch
			    - dev


		In each branch dir
			---
			info:
			  - branchName: dev
			    divergentCommit: ag43d7l
	*/
	// TODO: Shift away from BASE.json and make first commit a normal commit using main branch

	// TODO: Revert Commits

	// TODO: List Commits

	// TODO: Branching

}
