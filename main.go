package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	auFiles, _ := filepath.Glob("./audio*.mp4")
	voFiles, _ := filepath.Glob("./video*.mp4")

	var prefixMap = make(map[string][]string)
	for _, aufile := range auFiles {
		for _, vifile := range voFiles {
			prefixAu, _ := strings.CutPrefix(aufile, "audio")
			prefixVi, _ := strings.CutPrefix(vifile, "video")
			if prefixAu == prefixVi {
				var outName, _ = strings.CutSuffix(prefixAu, ".mp4")
				prefixMap[outName] = append(prefixMap[outName], aufile)
				prefixMap[outName] = append(prefixMap[outName], vifile)
			}
		}
	}

	for prefix, fileList := range prefixMap {
		if len(fileList) == 2 {
			videoFile := ""
			audioFile := ""
			for _, file := range fileList {
				if strings.Contains(file, "video") {
					videoFile = file
				} else if strings.Contains(file, "audio") {
					audioFile = file
				}
			}

			if videoFile != "" && audioFile != "" {
				cmd := exec.Command("./ffmpeg.exe", "-i", videoFile, "-i", audioFile, "-c:v", "copy", "-c:a", "aac", "-strict", "experimental", prefix+".mp4")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}
	}
}
