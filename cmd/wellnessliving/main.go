package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tekkamanendless/wellnessliving"
)

func main() {
	ctx := context.Background()

	client := wellnessliving.Client{}

	var verbose bool
	rootCommand := &cobra.Command{
		Use: "wellnessliving",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
	}
	rootCommand.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose logging.")

	{
		cmd := &cobra.Command{
			Use:  "raw <method> <path> [key=value [...]]",
			Args: cobra.MinimumNArgs(2),
			Run: func(cmd *cobra.Command, args []string) {
				method := args[0]
				path := args[1]
				values := url.Values{}
				for _, v := range args[2:] {
					if !strings.Contains(v, "=") {
						logrus.WithContext(ctx).Errorf("Invalid syntax for variable %q; expected '='.", v)
						os.Exit(1)
					}
					parts := strings.SplitN(v, "=", 2)
					values.Set(parts[0], parts[1])
				}

				contents, err := client.Raw(ctx, method, path, values)
				if err != nil {
					logrus.WithContext(ctx).Errorf("Could not perform request: [%T] %v", err, err)
					os.Exit(1)
				}
				fmt.Printf("%s\n", contents)
			},
		}
		rootCommand.AddCommand(cmd)
	}

	{
		cmd := &cobra.Command{
			Use:  "list-events [key=value [...]]",
			Args: cobra.MinimumNArgs(0),
			Run: func(cmd *cobra.Command, args []string) {
				values := url.Values{}
				for _, v := range args {
					if !strings.Contains(v, "=") {
						logrus.WithContext(ctx).Errorf("Invalid syntax for variable %q; expected '='.", v)
						os.Exit(1)
					}
					parts := strings.SplitN(v, "=", 2)
					values.Set(parts[0], parts[1])
				}

				var eventListResponse wellnessliving.EventListResponse
				err := client.Request(ctx, http.MethodGet, "/Wl/Event/EventList.json", values, &eventListResponse)
				if err != nil {
					logrus.WithContext(ctx).Errorf("Could not perform request: [%T] %v", err, err)
					os.Exit(1)
				}
				spew.Dump(eventListResponse)
			},
		}
		rootCommand.AddCommand(cmd)
	}

	err := rootCommand.Execute()
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error: [%T] %v", err, err)
		os.Exit(1)
	}
}
