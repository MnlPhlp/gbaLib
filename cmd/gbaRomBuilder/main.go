package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type params struct {
	OutFile,
	Source,
	Title,
	GameCode,
	MakerCode string
	Version int
}

func getParams() params {
	params := params{}
	flag.StringVar(&params.Title, "t", "<input>", "Title of the Game")
	flag.StringVar(&params.GameCode, "gc", "", "game code (four characters)")
	flag.StringVar(&params.MakerCode, "mc", "", "maker code (two characters)")
	flag.IntVar(&params.Version, "v", 0, "version")
	flag.StringVar(&params.OutFile, "o", "<title>.gba", "output filename")
	flag.Parse()
	println(flag.Args())
	if flag.NArg() == 0 {
		println("Missing source file")
		os.Exit(2)
	}
	params.Source = flag.Args()[flag.NArg()-1]
	if params.Title == "<input>" {
		params.Title = strings.TrimSuffix(params.Source, ".go")
	}
	if params.OutFile == "<title>.gba" {
		params.OutFile = params.Title + ".gba"
	}
	fmt.Printf("params: %v\n", func() string {
		data, _ := json.MarshalIndent(params, "", "  ")
		return string(data)
	}())
	return params
}

type dependencies struct {
	tinygo,
	gbafix,
	objcopy string
}

func getDependencies() dependencies {
	tinygo, err := exec.LookPath("tinygo")
	if err != nil {
		println("tinygo not found. Please install it first.")
	}
	gbafix, err := exec.LookPath("gbafix")
	if err != nil {
		println("gbafix not found. Please install it first.")
	}
	objcopy, err := exec.LookPath("arm-none-eabi-objcopy")
	if err != nil {
		println("arm-none-eabi-objcopy not found. Please install it first.")
	}
	return dependencies{
		tinygo:  tinygo,
		objcopy: objcopy,
		gbafix:  gbafix,
	}

}

func buildSrc(tinygo, src, dest string) {
	tmp, _ := ioutil.TempFile(".", "")
	cmd := exec.Cmd{
		Path: tinygo,
		Args: []string{
			tinygo, "build",
			"-target", "gameboy-advance",
			"-o", tmp.Name(),
			src,
		},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	err := cmd.Run()
	if err != nil {
		os.Remove(tmp.Name())
		log.Fatalln("Error: ", err)
	}
	os.Rename(tmp.Name(), dest)
}

func toBin(objcopy, path string) {
	cmd := exec.Cmd{
		Path: objcopy,
		Args: []string{
			objcopy,
			"-O", "binary",
			path, path,
		},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

func gbafix(gbafix string, params params) {
	args := []string{
		gbafix,
		params.OutFile,
		"-t", params.Title,
	}
	if len(params.MakerCode) > 0 {
		args = append(args, "-m"+params.MakerCode)
	}
	if len(params.GameCode) > 0 {
		args = append(args, "-c"+params.GameCode)
	}
	if params.Version > 0 {
		args = append(args, "-r"+strconv.Itoa(params.Version))
	}
	cmd := exec.Cmd{
		Path:   gbafix,
		Args:   args,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

func main() {
	dependencies := getDependencies()
	params := getParams()
	// build the program
	buildSrc(dependencies.tinygo, params.Source, params.OutFile)
	println("Build Complete")
	toBin(dependencies.objcopy, params.OutFile)
	println("Converted to binary")
	gbafix(dependencies.gbafix, params)
}
