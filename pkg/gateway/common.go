package gateway

import "time"

var (
	rpcRetryTimes = 3
	rpcTimeOut    = 5 * time.Second
)

const (
	scanBreaker        = "scanBreaker"
	staticBreaker      = "staticBreaker"
	institutionBreaker = "institutionBreaker"
	userBreaker        = "userBreaker"
	merchantBreaker    = "merchantBreaker"
	termBreaker        = "termBreaker"
	workflowBreaker    = "workflowBeaker"
)
