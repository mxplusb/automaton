// Copyright © 2017 Mike Lloyd <kevin.michael.lloyd@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mxplusb/automaton/initialiser"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new workspace.",
	Long: `Initialises a new workspace at the current path. It takes the HTTPS url of a git repository.

auto init https://github.com/mxplusb/auto-demo.git -u`,
	Run: func(cmd *cobra.Command, args []string) {
		initialiser.NewWorkspace(args[0], keepUpdated)
	},
}

var (
	keepUpdated bool
)

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&keepUpdated, "update", "u", false, "Pass this flag to keep the workspace manifest up to date.")
}
