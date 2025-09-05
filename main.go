package main

import (
	"digital_museum/fetchUtils"
	st "digital_museum/shared"
	"digital_museum/tui"
	"strconv"
	"math/rand/v2"
	"fmt"
)

func main() {
	fmt.Println("getting gallery...")
	total, objIDs := fetchUtils.FetchTotalArtNum()
	fmt.Println("done")

	var artData st.Art
	for {
		fmt.Println("generating random num...")
		randNum := rand.IntN(total)
		fmt.Println("done")

		objID := strconv.Itoa(objIDs[randNum])
		url := "https://collectionapi.metmuseum.org/public/collection/v1/objects/" + objID

		fmt.Println("fetching art...")
		artData = fetchUtils.FetchArtInfo(url)
		fmt.Println("done")
		fmt.Println("title: " + artData.Title)

		if artData.PrimaryImage != "" {
			break
		}
		fmt.Println("weird pic objID: " + objID)
	}

	var imgName string
	imgName = artData.Title
	fmt.Println("fetching img...")
	fmt.Println("img url:" + artData.PrimaryImage)
	fetchUtils.FetchImg(artData.PrimaryImage, imgName) // saves "temporary" file in system
	fmt.Println("done")

	fmt.Println("opening img...")
	fetchUtils.OpenImg(imgName)
	fmt.Println("done")
	//fetchUtils.RefreshObjIDs()

	tui.TuiGallery(artData)
}
