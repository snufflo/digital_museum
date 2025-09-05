package fetchUtils

import (
	st "digital_museum/shared"
	"net/http"
	"io"
	"os/exec"
	"encoding/json"
	"runtime"
	"os"
	"log"
	"fmt"
)

func FetchArtInfo(url string) st.Art {
	resp, err := http.Get(url)
	logErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	logErr(err)

	var jsonData st.Art
	err_json := json.Unmarshal(body, &jsonData)
	logErr(err_json)

	return jsonData
}

func FetchImg(url string, imgName string) {
	resp, err := http.Get(url)
	fmt.Println("fetched img")
	logErr(err)
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	logErr(err)
	FilePath := "art/" + imgName

	err_img := os.WriteFile(FilePath, imgData, 0644)
	logErr(err_img)
}

func FetchTotalArtNum() (int, []int) {
	bytes, err := os.ReadFile("objIDs.json")
	if err != nil { panic(err) }

	var jsonData st.Gallery
	err_json := json.Unmarshal(bytes, &jsonData)
	logErr(err_json)

	return jsonData.Total, jsonData.ObjectIDs
}

func RefreshObjIDs() {
	resp, err := http.Get("https://collectionapi.metmuseum.org/public/collection/v1/objects")
	if err != nil { panic(err) }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil { panic(err) }

	if err := os.WriteFile("objIDs.json", body, 0644); err != nil {
		panic(err)
	}
}

func OpenImg(imgName string) {
	var cmd *exec.Cmd
	FilePath := "art/" + imgName
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rund1132", "url.dll,FileProtocolHandler", FilePath)
	default:
		cmd = exec.Command(
			"feh", 
			"--geometry", 
			"960x1080+965+0", 
			"--borderless", 
			"--image-bg",
			"black", 
			"--scale-down",
			FilePath)
	}

	cmd.Start()
}

func logErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
