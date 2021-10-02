package provider

type CmdReader interface {
	ReadCommand() (string, error)
}
