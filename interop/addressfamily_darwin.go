package interop

type AddressFamily DarwinInt

/* https://github.com/apple/darwin-xnu/blob/main/bsd/sys/proc_info.h */
const (
	AddressFamilyUnspec         AddressFamily = /* AF_UNSPEC          */ 0
	AddressFamilyUnix           AddressFamily = /* AF_UNIX            */ 1
	AddressFamilyLocal          AddressFamily = /* AF_LOCAL           */ AddressFamilyUnix
	AddressFamilyInet           AddressFamily = /* AF_INET            */ 2
	AddressFamilyImpLink        AddressFamily = /* AF_IMPLINK         */ 3
	AddressFamilyPup            AddressFamily = /* AF_PUP             */ 4
	AddressFamilyChaos          AddressFamily = /* AF_CHAOS           */ 5
	AddressFamilyNs             AddressFamily = /* AF_NS              */ 6
	AddressFamilyIso            AddressFamily = /* AF_ISO             */ 7
	AddressFamilyOsi            AddressFamily = /* AF_OSI             */ AddressFamilyIso
	AddressFamilyEcma           AddressFamily = /* AF_ECMA            */ 8
	AddressFamilyDataKit        AddressFamily = /* AF_DATAKIT         */ 9
	AddressFamilyCcitt          AddressFamily = /* AF_CCITT           */ 10
	AddressFamilySna            AddressFamily = /* AF_SNA             */ 11
	AddressFamilyDecNet         AddressFamily = /* AF_DECnet          */ 12
	AddressFamilyDli            AddressFamily = /* AF_DLI             */ 13
	AddressFamilyLat            AddressFamily = /* AF_LAT             */ 14
	AddressFamilyHyLink         AddressFamily = /* AF_HYLINK          */ 15
	AddressFamilyAppleTalk      AddressFamily = /* AF_APPLETALK       */ 16
	AddressFamilyRoute          AddressFamily = /* AF_ROUTE           */ 17
	AddressFamilyLink           AddressFamily = /* AF_LINK            */ 18
	PseudoAddressFamilyXtp      AddressFamily = /* pseudo_AF_XTP      */ 19
	AddressFamilyCoip           AddressFamily = /* AF_COIP            */ 20
	AddressFamilyCnt            AddressFamily = /* AF_CNT             */ 21
	PseudoAddressFamilyRtip     AddressFamily = /* pseudo_AF_RTIP     */ 22
	AddressFamilyIpx            AddressFamily = /* AF_IPX             */ 23
	AddressFamilySip            AddressFamily = /* AF_SIP             */ 24
	PseudoAddressFamilyPip      AddressFamily = /* pseudo_AF_PIP      */ 25
	AddressFamilyNdrv           AddressFamily = /* AF_NDRV            */ 27
	AddressFamilyIsdn           AddressFamily = /* AF_ISDN            */ 28
	AddressFamilyE164           AddressFamily = /* AF_E164            */ AddressFamilyIsdn
	PseudoAddressFamilyKey      AddressFamily = /* pseudo_AF_KEY      */ 29
	AddressFamilyInet6          AddressFamily = /* AF_INET6           */ 30
	AddressFamilyNatm           AddressFamily = /* AF_NATM            */ 31
	AddressFamilySystem         AddressFamily = /* AF_SYSTEM          */ 32
	AddressFamilyNetBios        AddressFamily = /* AF_NETBIOS         */ 33
	AddressFamilyPpp            AddressFamily = /* AF_PPP             */ 34
	PseudoAddressFamilyHdrcmplt AddressFamily = /* pseudo_AF_HDRCMPLT */ 35
	AddressFamilyAfp            AddressFamily = /* AF_AFP             */ 36
	AddressFamilyIeee80211      AddressFamily = /* AF_IEEE80211       */ 37
	AddressFamilyUtun           AddressFamily = /* AF_UTUN            */ 38
	AddressFamilyMultiPath      AddressFamily = /* AF_MULTIPATH       */ 39
	AddressFamilyVsock          AddressFamily = /* AF_VSOCK           */ 40
	AddressFamilyMax            AddressFamily = /* AF_MAX             */ 41
)
