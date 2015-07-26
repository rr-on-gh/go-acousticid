package acousticid
import (
	"testing"
	"os"
	"os/exec"
)

func TestTaggerParallel(t *testing.T) {
	log.Debug("Setting up test dir...")
	os.RemoveAll("_testdir_parallel")
	_, _ = exec.Command("cp", "-r", "_template", "_testdir_parallel").Output()
	//fmt.Println(string(op))

	start_dir := os.Getenv("GOPATH") + "/src/github.com/raks81/go-acousticid/_testdir_parallel"
	//fmt.Println(start_dir)
	TagDirParallel(start_dir)

	//TODO Validate if id3 tags have been added

	//TODO delete the _testdir_parallel
}