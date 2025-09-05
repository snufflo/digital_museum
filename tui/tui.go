package tui

import (
	"log"
	"fmt"
	st "digital_museum/shared"
	"github.com/gdamore/tcell/v2"
	"strconv"
)

func TuiGallery(artInfo st.Art) {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}
	defer screen.Fini()

	artInfo = normalizeInfo(artInfo)
	showArt := false
	input := ""
	_, height := screen.Size()
	screenMiddle := height / 2

//	scroll := 0

	for {
		title := fmt.Sprintf("%s [ %s ]", artInfo.Title, artInfo.Medium)
		location := fmt.Sprintf("Location: %s, %s", artInfo.Country, artInfo.City)
		date := "Date of Creation: " + artInfo.ObjectDate
		author := fmt.Sprintf("Artist: %s, [ %s ]", artInfo.ArtistDisplayName, artInfo.ArtistDisplayBio)
		department := "Department: " + artInfo.Department
		objID := "Object ID: " + strconv.Itoa(artInfo.ObjectID)
		objUrl := "Object info: " + artInfo.ObjectURL
		artistUrl := "Artist info: " + artInfo.ArtistWikidata_URL

		totalDisplayHeight := 22

		l := getStartLine(screen, totalDisplayHeight)

		screen.Clear()

		if showArt {
			printLine(screen, l, title)
			printLine(screen, l+8, location)
			printLine(screen, l+10, date)
			printLine(screen, l+12, author)
			printLine(screen, l+14, department)
			printLine(screen, l+16, objID)
			printLine(screen, l+20, objUrl)
			printLine(screen, l+22, artistUrl)
		} else {
			printLine(screen, screenMiddle, "Press Spacebar to show further Information...")
		}

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			case tcell.KeyRune:
				if ev.Rune() == 'q' {
					return
				} else if ev.Rune() == ' ' {
					showArt = true
				}
			}
			/*		case tcell.KeyUp:
			if scroll > 0 {
				scroll--
			}
		case tcell.KeyDown:
			if scroll < len(lines)-height {
				scroll++
			}*/
		}
	}
}

func printLine(s tcell.Screen, y int, text string) {
	width, _ := s.Size()
	runes := []rune(text)

	for i := 0; i < len(runes); i += width {
		// wrap text
		end := i + width
		if end > len(runes) {
			end = len(runes)
		}
		line := string(runes[i:end])

		style := tcell.StyleDefault
		for x, r := range line {
			s.SetContent(x, y, r, nil, style)
		}
		y++
	}
}

func getStartLine(s tcell.Screen, displayHeight int) int {
	_, height := s.Size()

	var screenMiddle int
	screenMiddle = height / 2

	var displayHeightMiddle int
	displayHeightMiddle = displayHeight / 2

	startLocation := screenMiddle - displayHeightMiddle
	if startLocation < 0 {
		panic("totalDisplay too long!")
	}
	
	return startLocation
}

func printLineHighlight(s tcell.Screen, y int, text string, offset int, fg, bg tcell.Color) {
	width, _ := s.Size()
	style := tcell.StyleDefault.Foreground(fg).Background(bg)

	for x := offset; x < width; x++ {
		var r rune = ' '
		if x-offset < len(text) {
			r = rune(text[x-offset])
		}
		s.SetContent(x, y, r, nil, style)
	}
}

func normalizeInfo(a st.Art) st.Art {
	u := "_____" 
	if a.Title == "" { a.Title = u }
	if a.Medium == "" { a.Medium = u }
	if a.Country == "" { a.Country = u }
	if a.City == "" {a.City = u }
	if a.ObjectDate == "" {a.ObjectDate = u }
	if a.ArtistDisplayName == "" { a.ArtistDisplayName = u }
	if a.ArtistDisplayBio == "" { a.ArtistDisplayBio = u }
	if a.Department == "" { a.Department = u }
	if a.ObjectURL == "" { a.ObjectURL = u }
	if a.ArtistWikidata_URL == "" { a.ArtistWikidata_URL = u }

	return a
}
