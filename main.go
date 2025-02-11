package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	// フラグ定義
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// --help フラグ処理
	if *helpFlag {
		printHelp()
		os.Exit(0)
	}

	// --version フラグ処理
	if *versionFlag {
		fmt.Printf("codegen version %s\n", version)
		os.Exit(0)
	}

	// 残りの引数を処理
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No input file specified")
		printHelp()
		os.Exit(1)
	}

	// ここにOCODE処理の実装を追加
}

func printHelp() {
	fmt.Println("Usage: codegen [options] <input-file>")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExamples:")
	fmt.Println("  codegen program.ocode")
	fmt.Println("  codegen --version")
}
