package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/og11423074s/creacion-empaq-aplic/pokemon"
	"github.com/spf13/cobra"
)

var pkm string

func init() {
	rootCmd.Flags().StringVarP(&pkm, "pokemon", "p", "", "Name of the Pokémon to get")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pkm",
	Short: "A CLI tool for managing Pokemon data",
	Long:  `A CLI tool for managing Pokemon data`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getPokemon(pkm)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getPokemon(name string) error {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		return fmt.Errorf("failed to get pokemon: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var pokemon pokemon.Pokemon
	json.Unmarshal(body, &pokemon)
	fmt.Printf("Name: %s, %+v\n", pokemon.Name, pokemon)

	return nil
}
