package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"

	"github.com/patpir/miditf/blocks"
)


func ListCmd() cli.Command {
	return cli.Command{
		Name: "list",
		Usage: "list all available sources, transformations and visualizations",
		Action: performList,
		Subcommands: []cli.Command{
			{
				Name: "all",
				Usage: "list all available sources, transformations and visualizations",
				Action: performList,
			},
			{
				Name: "sources",
				Usage: "list all available sources",
				Action: performListSources,
			},
			{
				Name: "transformations",
				Usage: "list all available transformations",
				Action: performListTransformations,
			},
			{
				Name: "visualizations",
				Usage: "list all available visualizations",
				Action: performListVisualizations,
			},
		},
	}
}


func performList(c *cli.Context) error {
	performListSources(c)
	performListTransformations(c)
	performListVisualizations(c)
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

