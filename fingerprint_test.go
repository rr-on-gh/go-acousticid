package acousticid
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os/exec"
)

func TestFingerprint(t *testing.T) {
	_, _ = exec.Command("cp", "_template/test.mp3", "./test_fp.mp3").Output()

	fingerprint := new(Fingerprint)
	fingerprint.Get("test_fp.mp3")
	assert.NotEmpty(t, fingerprint.duration, "Duration should not be empty")
	assert.NotEmpty(t, fingerprint.duration, "Duration should not be empty")
	assert.NotEmpty(t, fingerprint.fingerprint, "Fingerprint should not be empty")
	exec.Command("rm", "-rf", "test_fp.mp3").Output()

}

func TestBadMp3File(t *testing.T) {
	
}

func TestLoooongMp3File(t *testing.T) {

}

func TestShortMp3File(t *testing.T) {

}