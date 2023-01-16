package SkyLine

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	SourceFile       = flag.String("source", "", "Load source code file into SkyLine")
	Bnn              = flag.Bool("bout", false, "If true will output the SkyLine banner when running a code file")
	Server           = flag.Bool("server", false, "If true will load the SkyLine local server")
	CompileWithGo    = flag.Bool("compile-with-go", false, "Compile with the interpreter but rather take the input of a source code file and compile it with the embedded interpreter")
	SettingsFileYAML = flag.String("settings-fileYAML", "", "Will load settings for SkyLine's interpreter from a YAML file")
	SettingsFileJSON = flag.String("settings-fileJSON", "", "Will load settings for SkyLine's interpreter from a JSON file")
)

func Banner() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("	 \033[38;5;51m┏━┓\x1b[0m")
		fmt.Println("	\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
		fmt.Println("    \033[38;5;56m━━━━┛\x1b[0m")
		fmt.Println("	\033[38;5;249mSky Line Interpreter| V 0.0.2")
		fmt.Print("\n\n\033[39m")
	default:
		fmt.Println("\t\t\t\t	 \u001b[38;5;51m┏━┓\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;56m┃\u001b[38;5;51m┃ ┃\u001b[0m")
		fmt.Println("\t\t\t\t    \u001b[38;5;56m━━━━┛\u001b[0m")
		fmt.Println("\t\t\t\t	\u001b[38;5;249mSky Line Interpreter| V 0.0.2")

	}
}