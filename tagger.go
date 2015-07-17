package acousticid

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func TagDir(start_dir string) {
	filepath.Walk(start_dir, tagFile)
}

func tagFile(file string, info os.FileInfo, err error) error {

	if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
		fmt.Println("Fingerprinting", file,"...")
		//Get the fingerprint
		fingerprint := new(Fingerprint)
		fingerprint.Get(file)

		//Get the acoustic id
		acousticid := new(AcousticidRequest)
		acousticid.Fingerprint = fingerprint.fingerprint
		acousticid.Duration = fingerprint.duration
		acousticid.Apikey = "ULjKruIh"
		acousticid.Metadata = "recordings"
		acousticidresponse := acousticid.Request()

		//Is the match good enough?
		if acousticidresponse.Results[0].Score > 0.9 {
			id3Tag := new(Id3)
			id3Tag.Artist = acousticidresponse.Results[0].Recordings[0].Artists[0].Name
			id3Tag.Title = acousticidresponse.Results[0].Recordings[0].Title
			id3Tag.set(file)
		}
	}
	return nil
}
