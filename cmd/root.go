package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "v2bridge",
	Short: "A bridge between V2Ray and other services.",
	Long: `
       ___   __         _     __         
 _   _|__ \ / /_  _____(_)___/ /___ ____ 
| | / /_/ // __ \/ ___/ / __  / __ '/ _ \
| |/ / __// /_/ / /  / / /_/ / /_/ /  __/
|___/____/_.___/_/  /_/\__,_/\__, /\___/ 
                            /____/

	v2bridge is a bridge between V2Ray and other services.
It provides a set of commands to interact with V2Ray.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var Server string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&Server, "server", "s", "localhost:10085", "V2Ray server address (address:port)")
}
