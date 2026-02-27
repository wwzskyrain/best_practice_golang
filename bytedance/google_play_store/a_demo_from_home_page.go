package google_play_store

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/n0madic/google-play-scraper/pkg/app"
)

func LoadDetailAndPermission() {
	// es.socialpoint.dragoncity
	// com.google.android.googlequicksearchbox
	a := app.New("es.socialpoint.dragoncity", app.Options{
		Country: "us",
		// Language: "us",
	})
	err := a.LoadDetails()
	if err != nil {
		panic(err)
	}
	err = a.LoadPermissions()
	if err != nil {
		panic(err)
	}
	spew.Dump(a)
}
