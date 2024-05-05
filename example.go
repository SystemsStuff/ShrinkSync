// This is an example program that uses the ShrinkSync framework
// To try this, replace the contents of cli/main.go with this file's and run `make clean datagrid infra`

package main

import "github.com/SystemsStuff/ShrinkSync/core"

func main() {
	shrinkSyncJob := core.NewShrinkSyncJob()
	shrinkSyncJob.SetInputMetadataLocation("wordCount_metadata")
	shrinkSyncJob.Start();
}