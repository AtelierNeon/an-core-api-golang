package interop

type ProcfsTcpState uint32

const (
	ProcfsTcpStateEstablished   ProcfsTcpState = /* TCP_ESTABLISHED    */ 1
	ProcfsTcpStateSynSent       ProcfsTcpState = /* TCP_SYN_SENT       */ 2
	ProcfsTcpStateSynRecv       ProcfsTcpState = /* TCP_SYN_RECV       */ 3
	ProcfsTcpStateFinWait1      ProcfsTcpState = /* TCP_FIN_WAIT1      */ 4
	ProcfsTcpStateFinWait2      ProcfsTcpState = /* TCP_FIN_WAIT2      */ 5
	ProcfsTcpStateTimeWait      ProcfsTcpState = /* TCP_TIME_WAIT      */ 6
	ProcfsTcpStateClose         ProcfsTcpState = /* TCP_CLOSE          */ 7
	ProcfsTcpStateCloseWait     ProcfsTcpState = /* TCP_CLOSE_WAIT     */ 8
	ProcfsTcpStateLastAck       ProcfsTcpState = /* TCP_LAST_ACK       */ 9
	ProcfsTcpStateListen        ProcfsTcpState = /* TCP_LISTEN         */ 10
	ProcfsTcpStateClosing       ProcfsTcpState = /* TCP_CLOSING        */ 11
	ProcfsTcpStateNewSynRecv    ProcfsTcpState = /* TCP_NEW_SYN_RECV   */ 12
	ProcfsTcpStateBoundInactive ProcfsTcpState = /* TCP_BOUND_INACTIVE */ 13
	ProcfsTcpStateMaxStates     ProcfsTcpState = /* TCP_MAX_STATES     */ 14
)

var procfsTcpState_typeStringMap = map[ProcfsTcpState]string{
	ProcfsTcpStateEstablished:   "ESTABLISHED",
	ProcfsTcpStateSynSent:       "SYN_SENT",
	ProcfsTcpStateSynRecv:       "SYN_RECV",
	ProcfsTcpStateFinWait1:      "FIN_WAIT1",
	ProcfsTcpStateFinWait2:      "FIN_WAIT2",
	ProcfsTcpStateTimeWait:      "TIME_WAIT",
	ProcfsTcpStateClose:         "CLOSE",
	ProcfsTcpStateCloseWait:     "CLOSE_WAIT",
	ProcfsTcpStateLastAck:       "LAST_ACK",
	ProcfsTcpStateListen:        "LISTEN",
	ProcfsTcpStateClosing:       "CLOSING",
	ProcfsTcpStateNewSynRecv:    "NEW_SYN_RECV",
	ProcfsTcpStateBoundInactive: "BOUND_INACTIVE",
	ProcfsTcpStateMaxStates:     "UNKNOWN",
}

func (pts ProcfsTcpState) String() string {
	str, ok := procfsTcpState_typeStringMap[pts]
	if ok {
		return str
	}

	return "UNKNOWN"
}
