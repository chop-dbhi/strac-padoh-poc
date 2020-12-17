package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/chop-dbhi/strac/converter"

	pa "github.com/chop-dbhi/strac/states/pa"
)

var (
	buildVersion string
)

var states = map[string][]*converter.Column{
	"pa": pa.Columns,
}

func main() {
	log.SetFlags(0)

	validateFlags := validateCmd.Flags()
	validateFlags.String("state", "", "State the validation applies to.")
	viper.BindPFlag("validate.state", validateFlags.Lookup("state"))

	convertFlags := convertCmd.Flags()
	convertFlags.String("state", "", "State the convert will be performed for.")
	viper.BindPFlag("convert.state", convertFlags.Lookup("state"))

	httpFlags := httpCmd.Flags()
	httpFlags.String("addr", "", "HTTP address.")
	viper.BindPFlag("http.addr", httpFlags.Lookup("addr"))

	rootCmd.AddCommand(
		versionCmd,
		validateCmd,
		convertCmd,
		httpCmd,
	)

	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use: "strac",
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildVersion)
	},
}

var validateCmd = &cobra.Command{
	Use: "validate",
	RunE: func(cmd *cobra.Command, args []string) error {
		state := viper.GetString("validate.state")
		// Supports both input and paths paths, input path (and STDOUT),
		// or no paths (STDIN and STDOUT).
		var (
			input     io.Reader
			inputPath string
		)

		switch len(args) {
		case 0:
		case 1:
			inputPath = args[0]
		case 2:
			inputPath = args[0]
		}

		columns, ok := states[strings.ToLower(state)]
		if !ok {
			return fmt.Errorf("state not registered: %s", state)
		}

		if inputPath != "" {
			f, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("open input file: %w", err)
			}
			defer f.Close()
			input = f
		} else {
			input = os.Stdin
		}

		return converter.Convert(input, ioutil.Discard, columns)
	},
}

var convertCmd = &cobra.Command{
	Use: "convert",
	RunE: func(cmd *cobra.Command, args []string) error {
		state := viper.GetString("convert.state")

		// Supports both input and paths paths, input path (and STDOUT),
		// or no paths (STDIN and STDOUT).
		var (
			input      io.Reader
			output     io.Writer
			inputPath  string
			outputPath string
		)

		switch len(args) {
		case 0:
		case 1:
			inputPath = args[0]
		case 2:
			inputPath = args[0]
			outputPath = args[1]
		}

		columns, ok := states[strings.ToLower(state)]
		if !ok {
			return fmt.Errorf("state not registered: %s", state)
		}

		if inputPath != "" {
			f, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("open input file: %w", err)
			}
			defer f.Close()
			input = f
		} else {
			input = os.Stdin
		}

		if outputPath != "" {
			f, err := os.Create(args[0])
			if err != nil {
				return fmt.Errorf("create output file: %w", err)
			}
			defer f.Close()
			output = f
		} else {
			output = os.Stdout
		}

		return converter.Convert(input, output, columns)
	},
}

var httpCmd = &cobra.Command{
	Use: "http",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := viper.GetString("http.addr")

		mux := echo.New()
		mux.GET("/", func(c echo.Context) error {
			return nil
		})

		mux.POST("/", func(c echo.Context) error {
			return nil
		})

		return mux.Start(addr)
	},
}
