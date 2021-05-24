package commando

type Command struct {
	Name     string `json:"name"`
	Args     []byte `json:"args,omitempty"`
	Encoding string `json:"encoding,omitempty"`
}
