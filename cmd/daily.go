package cmd

import (
	"log"

	"github.com/Yakitrak/notesmd-cli/pkg/actions"
	"github.com/Yakitrak/notesmd-cli/pkg/obsidian"
	"github.com/spf13/cobra"
)

var DailyCmd = &cobra.Command{
	Use:     "daily",
	Aliases: []string{"d"},
	Short:   "Creates or opens daily note in vault",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		vault := obsidian.Vault{Name: vaultName}
		uri := obsidian.Uri{}

		useEditor, err := cmd.Flags().GetBool("editor")
		if err != nil {
			log.Fatalf("Failed to parse --editor flag: %v", err)
		}
		if !cmd.Flags().Changed("editor") {
			defaultOpenType, configErr := vault.DefaultOpenType()
			if configErr == nil && defaultOpenType == "editor" {
				useEditor = true
			}
		}

		err = actions.DailyNote(&vault, &uri, actions.DailyParams{
			UseEditor: useEditor,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	DailyCmd.Flags().StringVarP(&vaultName, "vault", "v", "", "vault name (not required if default is set)")
	DailyCmd.Flags().BoolP("editor", "e", false, "open in editor instead of Obsidian")
	rootCmd.AddCommand(DailyCmd)
}
