package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Instala os arquivos de palavras",
	Long: `Instala os arquivos de palavras.

Habla já vem com um conjunto pronto, mas não é garantido que esteja atualizado com a fonte original.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://www.ime.usp.br/~pf/dicios/br-utf8.txt"
		fileName := "base-words.txt"

		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Não foi possível baixar o repositório de palavras")
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("Não foi possível baixar o repositório de palavras")
			return
		}

		file, err := os.Create(fileName)

		if err != nil {
			fmt.Println("Não foi possível salvar o arquivo base-words.txt")
			return
		}

		defer file.Close()

		_, err = io.Copy(file, resp.Body)

		if err != nil {
			fmt.Println("Não foi possível salvar o arquivo base-words.txt")
			return
		}

		return
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
