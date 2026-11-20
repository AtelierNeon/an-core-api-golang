package interop

type MibTcpState WindowsEnum

const (
	MibTcpStateClosed    MibTcpState = /* MIB_TCP_STATE_CLOSED     */ 1
	MibTcpStateListen    MibTcpState = /* MIB_TCP_STATE_LISTEN     */ 2
	MibTcpStateSynSent   MibTcpState = /* MIB_TCP_STATE_SYN_SENT   */ 3
	MibTcpStateSynRcvd   MibTcpState = /* MIB_TCP_STATE_SYN_RCVD   */ 4
	MibTcpStateEstab     MibTcpState = /* MIB_TCP_STATE_ESTAB      */ 5
	MibTcpStateFinWait1  MibTcpState = /* MIB_TCP_STATE_FIN_WAIT1  */ 6
	MibTcpStateFinWait2  MibTcpState = /* MIB_TCP_STATE_FIN_WAIT2  */ 7
	MibTcpStateCloseWait MibTcpState = /* MIB_TCP_STATE_CLOSE_WAIT */ 8
	MibTcpStateClosing   MibTcpState = /* MIB_TCP_STATE_CLOSING    */ 9
	MibTcpStateLastAck   MibTcpState = /* MIB_TCP_STATE_LAST_ACK   */ 10
	MibTcpStateTimeWait  MibTcpState = /* MIB_TCP_STATE_TIME_WAIT  */ 11
	MibTcpStateDeleteTcb MibTcpState = /* MIB_TCP_STATE_DELETE_TCB */ 12
	MibTcpStateReserved  MibTcpState = /* MIB_TCP_STATE_RESERVED   */ 100
)

var mibTcpState_typeStringMap = map[MibTcpState]string{
	MibTcpStateClosed:    "CLOSED",
	MibTcpStateListen:    "LISTEN",
	MibTcpStateSynSent:   "SYN_SENT",
	MibTcpStateSynRcvd:   "SYN_RECEIVED",
	MibTcpStateEstab:     "ESTABLISHED",
	MibTcpStateFinWait1:  "FIN_WAIT_1",
	MibTcpStateFinWait2:  "FIN_WAIT_2",
	MibTcpStateCloseWait: "CLOSE_WAIT",
	MibTcpStateClosing:   "CLOSING",
	MibTcpStateLastAck:   "LAST_ACK",
	MibTcpStateTimeWait:  "TIME_WAIT",
	MibTcpStateDeleteTcb: "DELETE",
	MibTcpStateReserved:  "UNKNOWN",
}

func (mts MibTcpState) String() string {
	str, ok := mibTcpState_typeStringMap[mts]
	if ok {
		return str
	}
	return "UNKNOWN"
}
