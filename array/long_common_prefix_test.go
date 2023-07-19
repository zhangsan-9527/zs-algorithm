package array

import (
	"fmt"
	"testing"
)

func TestLongComminPrefix(t *testing.T) {

	arr := []string{"zhangsan", "zhangsansisi", "zhangsansi", "zhangsansisisaa"}
	prefix := LongComminPrefix(arr)
	fmt.Println(prefix)
}
