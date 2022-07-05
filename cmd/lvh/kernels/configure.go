package kernels

import (
	"context"

	"github.com/cilium/little-vm-helper/pkg/kernels"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func configureCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "configure <kernel>",
		Short: "configure kernel",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			log := logrus.New()
			kd, err := kernels.LoadDir(dirName)
			if err != nil {
				log.Fatal(err)
			}

			kname := args[0]
			kcfg := kd.KernelConfig(kname)
			if kcfg == nil {
				log.Fatalf("kernel `%s` not found", kname)
			}

			err = kcfg.Configure(context.Background(), log, dirName)
			if err != nil {
				log.Fatal(err)
			}

		},
	}
}