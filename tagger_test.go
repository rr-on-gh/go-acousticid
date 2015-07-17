package acousticid
import (
	"testing"
	"fmt"
	"os/exec"
	"os"
)

func TestTagger(t *testing.T) {
	fmt.Println("Setting up test dir...")
	//op, err := exec.Command("rm",  "-rf", "_testdir").Output()
	err := os.RemoveAll("_testdir")

	op, _ := exec.Command("ls",  "-l", "_testdir").Output()
	fmt.Println(op,err)
	op, _ = exec.Command("cp", "-r", "_template", "_testdir").Output()
	fmt.Println(string(op))


	start_dir := os.Getenv("GOPATH") + "/src/github.com/raks81/go-acousticid/_testdir"
	fmt.Println(start_dir)
	TagDir(start_dir)



}
