package restful

const (
	AgentLoginErr = 501
	AgentNoLogin  = 500

	UserErr  = 101
	AuthErr  = 100
	NOdata   = 4
	SQLERROR = 3
	ERROR    = 2
	SUCCESS  = 0
)

var Errmsg = map[int]string{
	100: "error massage",
}
