package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/takama/caldera/pkg/config"
	"github.com/takama/caldera/pkg/generator"
	"github.com/takama/caldera/pkg/helper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates new service from templates using default/config settings",
	Long: `In this mode, you'll be not asked about everything.
The configuration file will be used for all other data,
such as the host, port, etc., if you have saved it before.
Otherwise, the default settings will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := new(config.Config)
		if err := viper.Unmarshal(&cfg); err != nil {
			fmt.Println("Error parsing of configuration, used default:", err)
		}
		if !path.IsAbs(cfg.Directories.Templates) {
			if currentDir, err := os.Getwd(); err == nil {
				cfg.Directories.Templates = path.Join(currentDir, cfg.Directories.Templates)
			}
		}
		if cfg.Directories.Service == "" {
			if currentDir, err := os.Getwd(); err == nil {
				cfg.Directories.Service = path.Join(path.Dir(currentDir), cfg.Name)
			}
		}
		generator.Run(cfg)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	newCmd.PersistentFlags().String("name", "my-service", "A name of your new service")
	newCmd.PersistentFlags().String("description", "my-service description", "A description of your new service")
	RootCmd.PersistentFlags().String("github", "my-account", "A Github account name")
	RootCmd.PersistentFlags().Bool("grpc-client", false, "A gRPC client using")
	helper.LogF("Flag error", viper.BindPFlag("name", newCmd.PersistentFlags().Lookup("name")))
	helper.LogF("Flag error", viper.BindPFlag("description", newCmd.PersistentFlags().Lookup("description")))
	helper.LogF("Flag error", viper.BindPFlag("github", RootCmd.PersistentFlags().Lookup("github")))
	helper.LogF("Flag error", viper.BindPFlag("client", RootCmd.PersistentFlags().Lookup("grpc-client")))
}