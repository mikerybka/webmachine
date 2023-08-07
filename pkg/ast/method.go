package myast

type Method struct {
	Name    string   `json:"name"`
	Params  []*Field `json:"params"`
	Results []*Field `json:"results"`
}
