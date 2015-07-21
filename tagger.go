package acousticid

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/op/go-logging"
)

var log = initLogger()

func initLogger() *logging.Logger {
	logging.MustGetLogger("Tagger")
	var format = logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.INFO, "")
	// Set the backends to be used.
	logging.SetBackend(backendLeveled)
	return logging.MustGetLogger("Tagger")
}



func TagDir(start_dir string) {
	filepath.Walk(start_dir, tagFile)
}

func tagFile(file string, info os.FileInfo, err error) error {

	if !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
		log.Debug("Fingerprinting" + file + "...")
		//Get the fingerprint
		fingerprint := new(Fingerprint)
		fingerprint.Get(file)

		//Get the acoustic id
		acousticid := new(AcousticidRequest)
		acousticid.Fingerprint = fingerprint.fingerprint
		acousticid.Duration = fingerprint.duration
		acousticid.Apikey = "JhYlJ8T1"
		acousticid.Metadata = "recordings"
		acousticidresponse := acousticid.Request()


		//log.Info("[" + info.Name()+"] --> ["+acousticidresponse.Results[0].Recordings[0].Artists[0].Name + " - " + acousticidresponse.Results[0].Recordings[0].Title+"]")

		//Is there any valid response?
		//0th element of the slice has the best match
		if acousticidresponse.Status == "ok" && len(acousticidresponse.Results) > 0  {
			for _, result := range acousticidresponse.Results {
				if result.Score > 0.7 && len(result.Recordings) > 0 && len(result.Recordings[0].Artists) > 0 {
					log.Info("[" + info.Name() + "] --> [" + result.Recordings[0].Artists[0].Name + " - " + result.Recordings[0].Title + "] [Score:", result.Score,"]")
					id3Tag := new(Id3)
					id3Tag.Artist = result.Recordings[0].Artists[0].Name
					id3Tag.Title = result.Recordings[0].Title
					id3Tag.set(file)
					//break
				} else {
					log.Error("[" + info.Name() + "] --> [No match found!]")
				}
			}
		} else {
			//todo handle no/bad matches
			log.Error("[" + info.Name() + "] --> [No match found!]")
		}
	}
	return nil
}

func match(aidresponse AcousticidResponse, info os.FileInfo, id3 *Id3)  {
	//TODO lot of works needs to go here for a good match
}