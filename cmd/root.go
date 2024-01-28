/*
Copyright © 2024 Rhenato Francisco Baraúna <rhbarauna@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/rhbarauna/goexpert-desafio-stress-tester/internal/stresstester"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goexpert-desafio-stress-tester",
	Short: "Um CLI para teste de carga",
	Long:  `Um CLI para teste de carga: envia requisições simultâneas a uma URL, configurável com flags as flags -url, -requests e -concurrency.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		stresstester.RunTester(url, requests, concurrency)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "Escolha ping ou pong")
	rootCmd.Flags().IntP("requests", "r", 1, "Escolha ping ou pong")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Escolha ping ou pong")
	rootCmd.MarkFlagRequired("url")
}
