package kick_user

import (
	"github.com/spf13/cobra"
)

const (
	migrateTypeAll        = "all"
	migrateTypeExceptTask = "except-task"
	migrateTypeOnlyTask   = "only-task"

	actionTypeMigrateData = "migrate-data"
	actionTypeCheckData   = "check-data"
	actionTypeRollback    = "rollback"
)

var (
	uid         int
	concurrency int
)

func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "kick",
		Short:   "kick user",
		Example: `./speedctl kick -c 100`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCommand()
		},
	}
	command.Flags().IntVarP(&uid, "userid", "u", 0, "用户ID")
	command.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "并发数")
	return command
}

func runCommand() (err error) {
	return
}
