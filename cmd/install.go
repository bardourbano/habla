package cmd

import (
	"bufio"
	"errors"
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
		fiveLetters := "five-letters-words.txt"

		downloadFile(url, fileName)

		fiveLettersFile, err := os.Create(fiveLetters)

		if err != nil {
			fmt.Println("Não foi possível tratar o arquivo")
		}

		defer fiveLettersFile.Close()
		writer := bufio.NewWriter(fiveLettersFile)

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			word := scanner.Text()

			fmt.Println("word: ", word)

			if len(word) == 5 {
				fmt.Println("letters: ", 5)
				_, _ = writer.WriteString(word)
				_, _ = writer.WriteString("\n")
			}

			if scanner.Err() != nil {
				fmt.Println("error: ", scanner.Err())
			}
		}

		return
	},
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("unable to download file")
	}

	defer resp.Body.Close()

	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	return nil
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
