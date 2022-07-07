package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var listCmd = &cobra.Command{
	Use:   "list tasks",
	Short: "list task",
	Run: func(cmd *cobra.Command, args []string) {
		parse()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func parse() {

	yfile, err := ioutil.ReadFile("./data/tasks.yml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		log.Fatal(err2)
	}

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
