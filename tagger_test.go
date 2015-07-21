package acousticid
import (
	"testing"
	"os/exec"
	"os"
)

func TestTagger(t *testing.T) {
	log.Debug("Setting up test dir...")
	os.RemoveAll("_testdir")
	_, _ = exec.Command("cp", "-r", "_template", "_testdir").Output()
	//fmt.Println(string(op))

	start_dir := os.Getenv("GOPATH") + "/src/github.com/raks81/go-acousticid/_testdir"
	//fmt.Println(start_dir)
	TagDir(start_dir)

	//TODO Validate if id3 tags have been added

	//TODO delete the _testdir
}
