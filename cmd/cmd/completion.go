/*
copyright 2020 the Escli authors

licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at

    http://www.apache.org/licenses/license-2.0

unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	longDescription = `
	Outputs shell completion for the given shell (bash or zsh)

	This depends on the bash-completion binary.  Example installation instructions:
	OS X:
		$ brew install bash-completion
		$ source $(brew --prefix)/etc/bash_completion
		$ escli completion bash > ~/.escli-completion  # for bash users
		$ escli completion zsh > ~/.escli-completion   # for zsh users
		$ source ~/.escli-completion
	Ubuntu:
		$ apt-get install bash-completion
		$ source /etc/bash-completion
		$ source <(escli completion bash) # for bash users
		$ source <(escli completion zsh)  # for zsh users

	Additionally, you may want to output the completion to a file and source in your .bashrc
`

	zshCompdef = "\ncompdef _escli escli\n"
)

func completion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		rootCmd(cmd).GenBashCompletion(os.Stdout)
	case "zsh":
		runCompletionZsh(cmd, os.Stdout)
	}
}

// NewCompletionCommand returns the cobra command that outputs shell completion code
func NewCompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use: "completion SHELL",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires 1 arg, found %d", len(args))
			}
			return cobra.OnlyValidArgs(cmd, args)
		},
		ValidArgs: []string{"bash", "zsh"},
		Short:     "Output shell completion for the given shell (bash or zsh)",
		Long:      longDescription,
		Run:       completion,
	}
}

func runCompletionZsh(cmd *cobra.Command, out io.Writer) {
	rootCmd(cmd).GenZshCompletion(out)
	io.WriteString(out, zshCompdef)
}

func rootCmd(cmd *cobra.Command) *cobra.Command {
	parent := cmd
	for parent.HasParent() {
		parent = parent.Parent()
	}
	return parent
}
