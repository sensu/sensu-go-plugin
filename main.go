package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
)

var (
	foo   string
	stdin *os.File
)

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-CHANGEME",
		Short: "The Sensu Go CHANGEME for x",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&foo,
		"foo",
		"f",
		"",
		"example")

	_ = cmd.MarkFlagRequired("foo")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return fmt.Errorf("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}

	eventJSON, err := ioutil.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %s", err)
	}

	event := &types.Event{}
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal stdin data: %s", err)
	}

	if err = event.Validate(); err != nil {
		return fmt.Errorf("failed to validate event: %s", err)
	}

	if !event.HasCheck() {
		return fmt.Errorf("event does not contain check")
	}

	return exampleAction(event)
}

func exampleAction(event *types.Event) error {
	fmt.Printf("hello world: %s\n", foo)

	fmt.Printf("Event Key: %s-%s\n", event.Entity.Name, event.Check.Name)

	return nil
}
