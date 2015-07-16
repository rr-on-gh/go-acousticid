package acousticid

import (
	"io/ioutil"
	"fmt"
	"strings"
)

func TagDir(start_dir string) {

	files, err := ioutil.ReadDir(start_dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp3") {
			fmt.Println(file.Name())
			//Get the fingerprint
			fingerprint := new(Fingerprint)
			fingerprint.Get(file.Name())

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
				id3Tag.set(file.Name())
			}


			//Create an id3 tag and set the id3 tag

		}
	}
}
