package pr

import (
	"github.com/jenkins-x/jx-gitops/pkg/cmd/pr/push"
	"github.com/jenkins-x/jx-helpers/pkg/cobras"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/spf13/cobra"
)

// NewCmdPR creates the new command
func NewCmdPR() *cobra.Command {
	command := &cobra.Command{
		Use:     "pr",
		Short:   "Commands for working with Pull Requests",
		Aliases: []string{"pullrequest", "pullrequests"},
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	command.AddCommand(cobras.SplitCommand(push.NewCmdPullRequestPush()))
	return command
}