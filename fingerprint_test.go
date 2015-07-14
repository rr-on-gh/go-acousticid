package acousticid
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFingerprint(t *testing.T) {
	fingerprint := new(Fingerprint)
	fingerprint.Get("test.mp3")
	assert.NotEmpty(t, fingerprint.duration, "Duration should not be empty")
	assert.NotEmpty(t, fingerprint.duration, "Duration should not be empty")
	assert.NotEmpty(t, fingerprint.fingerprint, "Fingerprint should not be empty")

}

func TestBadMp3File(t *testing.T) {
	
}

func TestLoooongMp3File(t *testing.T) {

}

func TestShortMp3File(t *testing.T) {

}