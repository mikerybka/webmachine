package types

import "strconv"

type ID int64

func (id ID) String() string {
	return strconv.FormatInt(int64(id)+1, 16)
}

func ParseID(s string) (ID, error) {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return ID(n - 1), nil
}
