package interop

type TcpState DarwinInt

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	TcpStateClosed      TcpState = /* TSI_S_CLOSED       */ 0
	TcpStateListen      TcpState = /* TSI_S_LISTEN       */ 1
	TcpStateSynSent     TcpState = /* TSI_S_SYN_SENT     */ 2
	TcpStateSynReceived TcpState = /* TSI_S_SYN_RECEIVED */ 3
	TcpStateEstablished TcpState = /* TSI_S_ESTABLISHED  */ 4
	TcpStateCloseWait   TcpState = /* TSI_S__CLOSE_WAIT  */ 5
	TcpStateFinWait1    TcpState = /* TSI_S_FIN_WAIT_1   */ 6
	TcpStateClosing     TcpState = /* TSI_S_CLOSING      */ 7
	TcpStateLastAck     TcpState = /* TSI_S_LAST_ACK     */ 8
	TcpStateFinWait2    TcpState = /* TSI_S_FIN_WAIT_2   */ 9
	TcpStateTimeWait    TcpState = /* TSI_S_TIME_WAIT    */ 10
	TcpStateReserved    TcpState = /* TSI_S_RESERVED     */ 11
)

var tcpState_typeStringMap = map[TcpState]string{
	TcpStateClosed:      "CLOSED",
	TcpStateListen:      "LISTEN",
	TcpStateSynSent:     "SYN_SENT",
	TcpStateSynReceived: "SYN_RECEIVED",
	TcpStateEstablished: "ESTABLISHED",
	TcpStateCloseWait:   "CLOSE_WAIT",
	TcpStateFinWait1:    "FIN_WAIT_1",
	TcpStateClosing:     "CLOSING",
	TcpStateLastAck:     "LAST_ACK",
	TcpStateFinWait2:    "FIN_WAIT_2",
	TcpStateTimeWait:    "TIME_WAIT",
	TcpStateReserved:    "UNKNOWN",
}

func (ts TcpState) String() string {
	str, ok := tcpState_typeStringMap[ts]
	if ok {
		return str
	}
	return "UNKNOWN"
}
