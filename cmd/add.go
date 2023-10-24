package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"rwietter/goc/repository"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add your local repository",
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Erro ao obter o diretório atual:", err)
			return
		}

		repo := repository.Repository{
			Path:   currentDir,
			Name:   "Teste",
			GitHub: "",
		}

		fmt.Println("Adicionando o repositório", repo.Name, "no diretório", repo.Path)

		repository.AddRepository(repo)
	},
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		return errors.New("requires at least one arg")
	// 	}
	// 	return fmt.Errorf("invalid color specified: %s", args[0])
	// },
}

func init() {
	rootCmd.AddCommand(addCmd)
}
