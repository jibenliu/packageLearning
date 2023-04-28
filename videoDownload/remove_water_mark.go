package main

import "os/exec"

func removeWatermark(inputFile string, outputFile string) error {
	ffmpegCmd := exec.Command("ffmpeg", "-i", inputFile, "-vf", "delogo=x=400:y=20:w=100:h=100:show=0", "-c:a", "copy", outputFile)
	return ffmpegCmd.Run()
}

func main() {
	err := removeWatermark("input.mp4", "output.mp4")
	if err != nil {
		panic(err)
	}
}
