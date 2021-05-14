package util

import "strconv"

func StringToInt64(value string) (int64, error) {
	int64Val, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return int64(int64Val), nil
}
