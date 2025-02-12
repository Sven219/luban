package gva

import (
	"github.com/mitchellh/go-homedir"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gva",
	Short: "小飞猪运维平台",
	Long: `欢迎大家加入我们,一起共创社区。
        _             
        (_)            
  _ __   _   __ _  ___ 
 | '_ \ | | / _'  |/ __|
 | |_) || || (_| |\__ \
 | .__/ |_| \__, ||___/
 | |         __/ |     
 |_|        |___/      

`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Warn.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gva.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			color.Warn.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gva" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gva")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		color.Warn.Println("Using config file:", viper.ConfigFileUsed())
	}
}
