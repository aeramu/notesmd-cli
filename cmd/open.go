package cmd

import (
	"log"

	"github.com/Yakitrak/notesmd-cli/pkg/actions"
	"github.com/Yakitrak/notesmd-cli/pkg/obsidian"
	"github.com/spf13/cobra"
)

var vaultName string
var sectionName string
var OpenVaultCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Opens note in vault by note name",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault := obsidian.Vault{Name: vaultName}
		uri := obsidian.Uri{}
		noteName := args[0]

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

		params := actions.OpenParams{NoteName: noteName, Section: sectionName, UseEditor: useEditor}
		err = actions.OpenNote(&vault, &uri, params)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	OpenVaultCmd.Flags().StringVarP(&vaultName, "vault", "v", "", "vault name (not required if default is set)")
	OpenVaultCmd.Flags().StringVarP(&sectionName, "section", "s", "", "heading text to open within the note (case-sensitive)")
	OpenVaultCmd.Flags().BoolP("editor", "e", false, "open in editor instead of Obsidian")
	rootCmd.AddCommand(OpenVaultCmd)
}
