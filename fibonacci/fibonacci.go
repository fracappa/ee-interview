package fibonacci

import (
	"fmt"
	"strconv"
)

func ParseIndex(i string) (string, error) {
	n, err:= strconv.Atoi(i)
	if err != nil {
		return "", fmt.Errorf("error in parsing the input. (err: %v)", err)
	}
	x, y := 0, 1
	for j := 0; j <= n; j++{
		x, y = y, x + y
	}
	return strconv.Itoa(x), nil
}
