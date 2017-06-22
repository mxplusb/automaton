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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/mxplusb/automaton/initialiser"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Creates or updates a new config file in your home directory with the appropriate configs.",
	Long: `
Creates or updates a new config file in your home directory with the appropriate configs.

auto init user -u "Mike Lloyd" -e "mike.lloyd@r3t.io"`,
	Run: func(cmd *cobra.Command, args []string) {
		initialiser.UpdateUser(user, email)
	},
}

var (
	user, email string
)

func init() {
	initCmd.AddCommand(userCmd)
	initCmd.Flags().StringVarP(&user, "user", "u", "", "User name for automaton. i.e., 'Mike Lloyd'")
	initCmd.Flags().StringVarP(&email, "email", "e", "", "Email for automaton.")
}
