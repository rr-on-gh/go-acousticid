package acousticid
import (
	"testing"
	"os/exec"
	"github.com/mikkyang/id3-go"
	"github.com/stretchr/testify/assert"
)

func TestAddId3ToFile(t *testing.T) {
	_, err := exec.Command("cp", "_template/test.mp3", "test_id3_write.mp3").Output()
	if err != nil {
		panic(err)
	}
	id3Tag := new(Id3)
	id3Tag.Album = "Album Name"
	id3Tag.Artist = "Test Artist"
	id3Tag.Title = "Test Title"
	id3Tag.set("test_id3_write.mp3")

	mp3File, err := id3.Open("test_id3_write.mp3")
	defer mp3File.Close()
	assert.Equal(t, mp3File.Title(), "Test Title")
	assert.Equal(t, mp3File.Album(), "Album Name")
	assert.Equal(t, mp3File.Artist(), "Test Artist")

	exec.Command("rm", "test_id3_write.mp3").Output()
}