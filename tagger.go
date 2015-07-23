package acousticid

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/op/go-logging"
	"strconv"
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
	//fmt.Println(start_dir)
	filepath.Walk(start_dir, tagFile)
}

func tagFile(file string, info os.FileInfo, err error) error {
	//fmt.Println("---------->", file, err)
	if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
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
		if acousticidresponse.Status == "ok" && len(acousticidresponse.Results) > 0  {
			matches := []Match{}
			for _, result := range acousticidresponse.Results {
				//log.Info(reflect.TypeOf(result).Kind().String())
				if result.Score > 0.7 && len(result.Recordings) > 0 &&  &result.Recordings[0].Title != nil && len(result.Recordings[0].Artists) > 0 && &result.Recordings[0].Artists[0].Name != nil {

					match := &Match{}
					match.filename = info.Name()
					match.result = result
					matches = append(matches, *match)
					//fmt.Println("match 0 ", matches[0].filename)
					//break
				} else {
					log.Debug("[" + info.Name() + "] --> [Ignoring a acousticid result.]")
				}
			}

			//Find the best match
			//Steps:
			// Check score of the match
			// Check score of the filename match
			// If file name matches >
			bestMatchIndex := 0
			if len(matches) > 1 {
				for i := 0; i < len(matches); i++ {
					for j := 1; j < len(matches); j++ {
						if matches[i].compare(matches[j]) > 0 {
							bestMatchIndex = i
						} else {
							bestMatchIndex = j
						}
					}
				}
				bestMatch := matches[bestMatchIndex].result
				for _, match := range matches {
					log.Info("Possible Matches: " + "[ " + info.Name() + "] --> [" + match.result.Recordings[0].Artists[0].Name + " - " + match.result.Recordings[0].Title + "] [Score:", strconv.FormatFloat(match.result.Score, 'f', 6, 64), "]")
				}
				log.Info("Best Match: " + "[ " + info.Name() + "] --> [" + bestMatch.Recordings[0].Artists[0].Name + " - " + bestMatch.Recordings[0].Title + "] [Score:", strconv.FormatFloat(bestMatch.Score, 'f', 6, 64), "]")

			} else if len(matches) == 1 {
				bestMatch := matches[bestMatchIndex].result
				for _, match := range matches {
					log.Info("Possible Matches: " + "[" + info.Name() + "] --> [" + match.result.Recordings[0].Artists[0].Name + " - " + match.result.Recordings[0].Title + "] [Score:", strconv.FormatFloat(match.result.Score, 'f', 6, 64), "]")
				}
				log.Info("Best Match: " + "[" + info.Name() + "] --> [" + bestMatch.Recordings[0].Artists[0].Name + " - " + bestMatch.Recordings[0].Title + "] [Score:", strconv.FormatFloat(bestMatch.Score, 'f', 6, 64), "]")

			} else {
				//todo handle no/bad matches
				log.Error("[" + info.Name() + "] --> [No match found!!]")
			}
			//Set the id3 tag
//			log.Info("[" + info.Name() + "] --> [" + result.Recordings[0].Artists[0].Name + " - " + result.Recordings[0].Title + "] [Score:", result.Score,"]")
//			id3Tag := new(Id3)
//			id3Tag.Artist = result.Recordings[0].Artists[0].Name
//			id3Tag.Title = result.Recordings[0].Title
//			id3Tag.set(file)
		} else {
			//todo handle no/bad matches
			log.Error("[" + info.Name() + "] --> [No match found!]")
		}
	}
	log.Info("-----------------------------------------------------------------------------------------------------------------------------------")
	return nil
}
//
type Match struct {
	filename string
	result Result
	fileNameMatchScore float64
	weightedScore float64
}
//
func (m *Match) compare(m2 Match) int {
	//fmt.Println(len(m.result.Recordings))
	//log.Info(m.result.Recordings[0].Artists[0].Name)
	if FuzzyMatchString(m.filename, m.result.Recordings[0].Title + " - " + m.result.Recordings[0].Artists[0].Name) >
		FuzzyMatchString(m2.filename, m2.result.Recordings[0].Title + " - " + m2.result.Recordings[0].Artists[0].Name) {
		return 1
	} else if  FuzzyMatchString(m.filename, m.result.Recordings[0].Title + " - " + m.result.Recordings[0].Artists[0].Name) <
	FuzzyMatchString(m2.filename, m2.result.Recordings[0].Title + " - " + m2.result.Recordings[0].Artists[0].Name) {
		return -1
	} else {
		if m.result.Score > m2.result.Score {
			return 1
		} else if m.result.Score < m2.result.Score {
			return -1
		} else {
			return 0
		}
	}
}