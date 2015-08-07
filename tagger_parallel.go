package acousticid
import (
	"path/filepath"
	"os"
	"strings"
)

var mp3FileCount = 0
var fingerPrintInputChan = make(chan FingerprintInput, 100)
var acousticIdInputChan = make(chan Acousticidinput, 100)
var id3TagInputChan = make(chan ID3TagInput, 100)
var processedChan = make(chan bool, 100)

func TagDirParallel(start_dir string) {
	//Setup fingerprint worker
	for w := 0; w < 3 ; w ++ {
		go FingerprintWorker(w)
	}

	//Setup acousticid worker
	for w := 0; w < 3 ; w ++ {
		go AcousticidWorker(w)
	}

	//Setup ID3 worker
	for w := 0; w < 3 ; w ++ {
		go ID3Worker(w)
	}

	//Count the number of files to tag:
	filepath.Walk(start_dir, countMp3Files)
	log.Info("Need to tag %d mp3 files..", mp3FileCount)

	//Send the files to be fingerprinted
	filepath.Walk(start_dir, tagFileParallel)

	//Ensure all files have been processed
	for i := 0; i < mp3FileCount; i++ {
		<-processedChan
	}
}

func countMp3Files(file string, info os.FileInfo, err error) error  {
	if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
		mp3FileCount ++;
	}
	return nil
}

func tagFileParallel(file string, info os.FileInfo, err error) error {
	if err == nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".mp3") {
		log.Info("Sending file " + file + " for fingerprinting" )
		fingerPrintInputChan <- FingerprintInput{file, info}
	}
	return nil
}

func FingerprintWorker(id int) {
	log.Info("Initializing FingerprintWorker...")
	for fingerprintInput := range fingerPrintInputChan {
		log.Info("[ Worker: %d - Fingerprinting %s ]", id, fingerprintInput.file)
		acousticIdInputChan <- Acousticidinput{fingerprintInput.file, fingerprintInput.info, GetFingerprint(fingerprintInput.file)}
	}
	log.Info("FingerprintWorker closed!")
}

func AcousticidWorker(id int) {
	for acousticidInput := range acousticIdInputChan {
		log.Info("[ Worker: %d - Finding acoustic id of file with duration %s ]", id, acousticidInput.fingerprint.duration)
		id3TagInputChan <- ID3TagInput{acousticidInput.file, acousticidInput.info, GetAcousticId(acousticidInput.fingerprint)}
	}
}

func ID3Worker(id int) {
	for id3TagInput := range id3TagInputChan {
		log.Info("[ Worker: %d - Setting ID3 Tag %s ]", id, id3TagInput.file)
		SetID3(id3TagInput.acousticidresponse, id3TagInput.file, id3TagInput.info)
		processedChan <- true
	}
}


//Structs to pass around the channels
type FingerprintInput struct {
	file string
	info os.FileInfo
}

type Acousticidinput struct  {
	file string
	info os.FileInfo
	fingerprint Fingerprint
}

type ID3TagInput struct  {
	file string
	info os.FileInfo
	acousticidresponse AcousticidResponse
}
