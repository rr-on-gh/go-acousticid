package acousticid
import (
	"path/filepath"
	"os"
	"strings"
	"strconv"
)

var filesChan = make(chan string, 100)
var fingerprintsChan = make(chan Fingerprint, 100)
var acousticIdsChan = make(chan AcousticidResponse, 100)
var id3TagInputChan = make(chan ID3TagInput, 100)

func TagDirParallel(start_dir string) {
	//Setup fingerprint worker
	for w := 0; w < 3 ; w ++ {
		go FingerprintWorker(w, filesChan, fingerprintsChan)
	}

	//Setup acousticid worker
	for w := 0; w < 3 ; w ++ {
		go AcousticidWorker(w, fingerprintsChan, acousticIdsChan)
	}

	//Setup ID3 worker
	for w := 0; w < 3 ; w ++ {
		go ID3Worker(w, id3TagInputChan)
	}

	filepath.Walk(start_dir, tagFileParallel)
}

func tagFileParallel(file string, info os.FileInfo, err error) error {

	log.Info("-----------------------------------------------------------------------------------------------------------------------------------")
	if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {

		//Get fingerprint
		//fingerprint := fil
		filesChan <- file

		//Get acoustic Id
		fingerprintsChan <- (<-fingerprintsChan)

		//Set ID3 tag
		id3TagInputChan <- ID3TagInput{<-acousticIdsChan, file, info}

		log.Info("-----------------------------------------------------------------------------------------------------------------------------------")
	}
	return nil
}

func FingerprintWorker(id int, files <-chan string, fingerPrints chan <- Fingerprint) {
	for file := range files {
		log.Info("[ Worker: " + strconv.FormatInt(int64(id),10) + " - Fingerprinting " +  file + " ]")
		fingerPrints <- GetFingerprint(file)
	}
}

func AcousticidWorker(id int, fingerprints <-chan Fingerprint, results chan <- AcousticidResponse) {
	for fingerprint := range fingerprints{
		log.Info("[ Worker: " + strconv.FormatInt(int64(id),10) + " - Finding acoustic id of file with duration " + strconv.FormatInt(int64(fingerprint.duration),10) + " ]")
		results <- GetAcousticId(fingerprint)
	}
}

func ID3Worker(id int, id3TagInputs <-chan ID3TagInput) {
	for id3TagInput := range id3TagInputs {
		log.Info("[ Worker: " + strconv.FormatInt(int64(id),10) + " - Setting ID3 Tag " + id3TagInput.file + " ]")
		SetID3(id3TagInput.acousticidresponse, id3TagInput.file, id3TagInput.info)
	}
}

type ID3TagInput struct {
	acousticidresponse AcousticidResponse
	file string
	info os.FileInfo
}