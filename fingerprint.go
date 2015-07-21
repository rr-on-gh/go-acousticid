package acousticid
import (
	"os/exec"
	"strings"
	"strconv"
	"os"
)

type Fingerprint struct {
	fingerprint string
	duration int
}

func (f *Fingerprint) Get(file string) {

	fpcalc := "./fpcalc"
	if os.Getenv("FPCALC_BINARY_PATH") == "" {
		log.Debug("Environment variable FPCALC_BINARY_PATH not set. Will use fpcalc in PATH as the default fingerprinting application")
	} else {
		fpcalc = os.Getenv("FPCALC_BINARY_PATH")
	}
	out, err := exec.Command(fpcalc, file).Output()
	if err != nil {
		panic(err)
	}
	outstrs := strings.Split(string(out), "\n")

	for _, s := range outstrs  {
		if strings.Index(s, "DURATION=") == 0 {
			ds := strings.Split(s, "=")[1]
			f.duration, _ = strconv.Atoi(ds)
		} else if strings.Index(s, "FINGERPRINT=") == 0 {
			f.fingerprint = strings.Split(s, "=")[1]
		}
	}
}
