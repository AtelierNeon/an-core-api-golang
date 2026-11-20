package interop

type TcpState uint32

const (
	TcpStateEstablished   TcpState = /* TCP_ESTABLISHED    */ 1
	TcpStateSynSent       TcpState = /* TCP_SYN_SENT       */ 2
	TcpStateSynRecv       TcpState = /* TCP_SYN_RECV       */ 3
	TcpStateFinWait1      TcpState = /* TCP_FIN_WAIT1      */ 4
	TcpStateFinWait2      TcpState = /* TCP_FIN_WAIT2      */ 5
	TcpStateTimeWait      TcpState = /* TCP_TIME_WAIT      */ 6
	TcpStateClose         TcpState = /* TCP_CLOSE          */ 7
	TcpStateCloseWait     TcpState = /* TCP_CLOSE_WAIT     */ 8
	TcpStateLastAck       TcpState = /* TCP_LAST_ACK       */ 9
	TcpStateListen        TcpState = /* TCP_LISTEN         */ 10
	TcpStateClosing       TcpState = /* TCP_CLOSING        */ 11
	TcpStateNewSynRecv    TcpState = /* TCP_NEW_SYN_RECV   */ 12
	TcpStateBoundInactive TcpState = /* TCP_BOUND_INACTIVE */ 13
	TcpStateMaxStates     TcpState = /* TCP_MAX_STATES     */ 14
)

var tcpState_typeStringMap = map[TcpState]string{
	TcpStateEstablished:   "ESTABLISHED",
	TcpStateSynSent:       "SYN_SENT",
	TcpStateSynRecv:       "SYN_RECV",
	TcpStateFinWait1:      "FIN_WAIT1",
	TcpStateFinWait2:      "FIN_WAIT2",
	TcpStateTimeWait:      "TIME_WAIT",
	TcpStateClose:         "CLOSE",
	TcpStateCloseWait:     "CLOSE_WAIT",
	TcpStateLastAck:       "LAST_ACK",
	TcpStateListen:        "LISTEN",
	TcpStateClosing:       "CLOSING",
	TcpStateNewSynRecv:    "NEW_SYN_RECV",
	TcpStateBoundInactive: "BOUND_INACTIVE",
	TcpStateMaxStates:     "UNKNOWN",
}

func (pts TcpState) String() string {
	str, ok := tcpState_typeStringMap[pts]
	if ok {
		return str
	}

	return "UNKNOWN"
}
