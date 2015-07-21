package acousticid
import (
	"testing"
	"fmt"
)

func TestFuzzyMatchString(t *testing.T) {
	fmt.Println(FuzzyMatchString("test this ", "this test"))
}
