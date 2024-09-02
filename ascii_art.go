package main

import (
	"fmt"
	"os"
	"strings"
)

var banners map[string]string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as an argument")
		return
	}

	LoadBannerFiles()

	input := os.Args[1]

	fmt.Println("Shadow:")
	PrintBanner(input, "shadow")

	fmt.Println("Standard:")
	PrintBanner(input, "standard")

	fmt.Println("Thinkertoy:")
	PrintBanner(input, "thinkertoy")
}

func LoadBannerFiles() {
	banners = make(map[string]string)

	shadowFile, err := os.ReadFile("shadow.txt")
	if err != nil {
		fmt.Println("Error reading shadow.txt:", err)
		return
	}
	banners["shadow"] = string(shadowFile)

	standardFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading standard.txt:", err)
		return
	}
	banners["standard"] = string(standardFile)

	thinkertoyFile, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error reading thinkertoy.txt:", err)
		return
	}
	banners["thinkertoy"] = string(thinkertoyFile)
}

func PrintBanner(text string, bannerType string) {
	banner, exists := banners[bannerType]
	if !exists {
		fmt.Println("Banner type not found:", bannerType)
		return
	}

	lines := strings.Split(banner, "\n")
	for i := 0; i < 8; i++ { // assuming each character is 8 lines tall
		for _, char := range text {
			runeChar := rune(char)
			index := (int(runeChar) - 32) * 9
			if index < 0 || index+8 >= len(lines) {
				fmt.Print("         ")
			} else {
				fmt.Print(lines[index+i])
			}
		}
		fmt.Println()
	}
}
