package acousticid
import (
	"testing"
	"os"
	"fmt"
	"os/exec"
)

func TestTagger(t *testing.T) {
	fmt.Println("Setting up test dir...")
	op, _ := exec.Command("cp", "_testdir", "_testdir_1").Output()
	fmt.Println(string(op))

	start_dir := os.Getenv("GOPATH") + "/src/github.com/raks81/go-acousticid"

	TagDir(start_dir)
}
