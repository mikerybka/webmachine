package web

func Cookie(key string) string {
	c, err := Request("", "").Cookie(key)
	if err != nil {
		return ""
	}
	return c.Value
}
