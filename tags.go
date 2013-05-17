// Am Laher 2013
// Based on code by the Go Authors - their license:

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"
)

var cmdTags = &Command{
	UsageLine: "tags [-repo <repo>] [package]",
	Short:     "Check available tags for a package",
	Long: `
Tags lists the tags found in a repository.

The -repo flag instructs get to download the package from the specified repo, as
opposed to the one suggested by its 'import path'
	`,
}

var tagsRepo = cmdTags.Flag.String("repo", "", "")

func init() {
	addBuildFlags(cmdTags)
	cmdTags.Run = runTags // break init loop
}
func runTags(cmd *Command, args []string) {

	var stk importStack
	for _, arg := range downloadPaths(args) {
		p := loadPackage(arg, &stk)
		listPackageTags(p)
	}
}
func listPackageTags(p *Package) error {
	var (
		vcs            *vcsCmd
		repo, rootPath string
		err            error
	)
	if p.build.SrcRoot != "" {
		// Directory exists.  Look for checkout along path to src.
		vcs, rootPath, err = vcsForDir(p)
		if err != nil {
			return err
		}
		repo = "<local>" // should be unused; make distinctive
	} else if p.RepoPath != "" {
		// Analyze the import path to determine the version control system,
		// repository, and the import path for the root of the repository.
		rr, err := repoRootForImportPath(p.RepoPath, p.ImportPath)
		if err != nil {
			return err
		}
		logf("Using RepoPath")
		vcs, repo, rootPath = rr.vcs, rr.repo, rr.root
	} else {

		// Analyze the import path to determine the version control system,
		// repository, and the import path for the root of the repository.
		rr, err := repoRootForImportPath(p.ImportPath, p.ImportPath)
		if err != nil {
			return err
		}
		logf("Using ImportPath")
		vcs, repo, rootPath = rr.vcs, rr.repo, rr.root
	}
	root := filepath.Join(p.build.SrcRoot, rootPath)
	tags, err := vcs.tags(root)
	if err != nil {
		errorf("Error getting tags (repo %s): %v", repo, err)
	} else {
		logf("Tags:")
		for _, t := range tags {
			logf(" %s", t)
		}
	}
	return err
}
