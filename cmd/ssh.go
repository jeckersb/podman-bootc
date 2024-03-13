package cmd

import (
	"podman-bootc/pkg/config"
	"podman-bootc/pkg/user"
	"podman-bootc/pkg/vm"

	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use:   "ssh <ID>",
	Short: "SSH into an existing OS Container machine",
	Long:  "SSH into an existing OS Container machine",
	Args:  cobra.MinimumNArgs(1),
	RunE:  doSsh,
}
var sshUser string

func init() {
	RootCmd.AddCommand(sshCmd)
	sshCmd.Flags().StringVarP(&sshUser, "user", "u", "root", "--user <user name> (default: root)")
}

func doSsh(_ *cobra.Command, args []string) error {
	user, err := user.NewUser()
	if err != nil {
		return err
	}

	id := args[0]

	vm, err := vm.NewVM(vm.NewVMParameters{
		ImageID:    id,
		User:       user,
		LibvirtUri: config.LibvirtUri,
	})

	if err != nil {
		return err
	}

	err = vm.SetUser(sshUser)
	if err != nil {
		return err
	}

	cmd := make([]string, 0)
	if len(args) > 1 {
		cmd = args[1:]
	}
	return vm.RunSSH(cmd)
}
