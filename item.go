package main

type execWithArgs struct {
	Exec string   // `json:exec`
	Args []string // `json:args,omitempty`
}

type item struct {
	Name     string         // `json:name`
	Commands []execWithArgs // `json:commands,omitempty`
	Exec     string         // `json:exec`
	Args     []string       // `json:args,omitempty`
}
