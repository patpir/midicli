package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/patpir/miditf/blocks"

	"github.com/patpir/midicli/messages"
)


func List() cli.Command {
	return cli.Command{
		Name: "list",
		Aliases: []string{"ls"},
		Usage: "list all available sources, transformations and visualizations",
		Action: performList,
		Subcommands: []cli.Command{
			{
				Name: "all",
				Aliases: []string{"a"},
				Usage: "list all available sources, transformations and visualizations",
				Action: performList,
			},
			{
				Name: "sources",
				Aliases: []string{"s"},
				Usage: "list all available sources",
				Action: performListSources,
			},
			{
				Name: "transformations",
				Aliases: []string{"t"},
				Usage: "list all available transformations",
				Action: performListTransformations,
			},
			{
				Name: "visualizations",
				Aliases: []string{"v"},
				Usage: "list all available visualizations",
				Action: performListVisualizations,
			},
		},
	}
}


func performList(c *cli.Context) error {
	if len(c.Args()) == 0 {
		performListSources(c)
		performListTransformations(c)
		performListVisualizations(c)
		return nil
	}

	messages.PrintCommandArgumentUnexpected(c.Args().First())
	cli.ShowAppHelpAndExit(c, 1)
	return nil
}

func performListSources(c *cli.Context) error {
	sources := blocks.DefaultRegistrator().Sources()
	printBlockInfoList("sources", sources)
	return nil
}

func performListTransformations(c *cli.Context) error {
	transformations := blocks.DefaultRegistrator().Transformations()
	printBlockInfoList("transformations", transformations)
	return nil
}

func performListVisualizations(c *cli.Context) error {
	visualizations := blocks.DefaultRegistrator().Visualizations()
	printBlockInfoList("visualizations", visualizations)
	return nil
}

func printBlockInfoList(category string, infoList []blocks.BlockInfo) {
	if len(infoList) == 0 {
		fmt.Printf("NO %s AVAILABLE!\n\n", strings.ToUpper(category))
	} else {
		fmt.Printf("AVAILABLE %s:\n\n", strings.ToUpper(category))
		for _, info := range infoList {
			fmt.Printf("  %s - %s\n", info.Identifier(), info.Description())
			args := info.ArgumentInfos()
			if len(args) == 0 {
				fmt.Println("     (no arguments)")
			} else {
				for _, arg := range args {
					optional := " (required)"
					if arg.IsOptional() {
						optional = " (optional)"
					}
					fmt.Printf("     -  %-10s    %s%s\n", arg.Name(), arg.Description(), optional)
				}
			}
			fmt.Println()
		}
	}
}

