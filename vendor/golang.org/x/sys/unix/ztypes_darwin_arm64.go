// cgo -godefs types_darwin.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build arm64 && darwin
// +build arm64,darwin

package unix

const (
	SizeofPtr      = 0x8
	SizeofShort    = 0x2
	SizeofInt      = 0x4
	SizeofLong     = 0x8
	SizeofLongLong = 0x8
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int64
	_C_long_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int32
	_    [4]byte
}

type Timeval32 struct {
	Sec  int32
	Usec int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Dev     int32
	Mode    uint16
	Nlink   uint16
	Ino     uint64
	Uid     uint32
	Gid     uint32
	Rdev    int32
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	Btim    Timespec
	Size    int64
	Blocks  int64
	Blksize int32
	Flags   uint32
	Gen     uint32
	Lspare  int32
	Qspare  [2]int64
}

type Statfs_t struct {
	Bsize       uint32
	Iosize      int32
	Blocks      uint64
	Bfree       uint64
	Bavail      uint64
	Files       uint64
	Ffree       uint64
	Fsid        Fsid
	Owner       uint32
	Type        uint32
	Flags       uint32
	Fssubtype   uint32
	Fstypename  [16]byte
	Mntonname   [1024]byte
	Mntfromname [1024]byte
	Flags_ext   uint32
	Reserved    [7]uint32
}

type Flock_t struct {
	Start  int64
	Len    int64
	Pid    int32
	Type   int16
	Whence int16
}

type Fstore_t struct {
	Flags      uint32
	Posmode    int32
	Offset     int64
	Length     int64
	Bytesalloc int64
}

type Radvisory_t struct {
	Offset int64
	Count  int32
	_      [4]byte
}

type Fbootstraptransfer_t struct {
	Offset int64
	Length uint64
	Buffer *byte
}

type Log2phys_t struct {
	Flags uint32
	_     [16]byte
}

type Fsid struct {
	Val [2]int32
}

type Dirent struct {
	Ino     uint64
	Seekoff uint64
	Reclen  uint16
	Namlen  uint16
	Type    uint8
	Name    [1024]int8
	_       [3]byte
}

const (
	PathMax = 0x400
)

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [104]int8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [92]int8
}

type RawSockaddrCtl struct {
	Sc_len      uint8
	Sc_family   uint8
	Ss_sysaddr  uint16
	Sc_id       uint32
	Sc_unit     uint32
	Sc_reserved [5]uint32
}

type _Socklen uint32

type Xucred struct {
	Version uint32
	Uid     uint32
	Ngroups int16
	Groups  [16]uint32
}

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet4Pktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Filt [8]uint32
}

const (
	SizeofSockaddrInet4    = 0x10
	SizeofSockaddrInet6    = 0x1c
	SizeofSockaddrAny      = 0x6c
	SizeofSockaddrUnix     = 0x6a
	SizeofSockaddrDatalink = 0x14
	SizeofSockaddrCtl      = 0x20
	SizeofXucred           = 0x4c
	SizeofLinger           = 0x8
	SizeofIovec            = 0x10
	SizeofIPMreq           = 0x8
	SizeofIPMreqn          = 0xc
	SizeofIPv6Mreq         = 0x14
	SizeofMsghdr           = 0x30
	SizeofCmsghdr          = 0xc
	SizeofInet4Pktinfo     = 0xc
	SizeofInet6Pktinfo     = 0x14
	SizeofIPv6MTUInfo      = 0x20
	SizeofICMPv6Filter     = 0x20
)

const (
	PTRACE_TRACEME = 0x0
	PTRACE_CONT    = 0x7
	PTRACE_KILL    = 0x8
)

type Kevent_t struct {
	Ident  uint64
	Filter int16
	Flags  uint16
	Fflags uint32
	Data   int64
	Udata  *byte
}

type FdSet struct {
	Bits [32]int32
}

const (
	SizeofIfMsghdr    = 0x70
	SizeofIfData      = 0x60
	SizeofIfaMsghdr   = 0x14
	SizeofIfmaMsghdr  = 0x10
	SizeofIfmaMsghdr2 = 0x14
	SizeofRtMsghdr    = 0x5c
	SizeofRtMetrics   = 0x38
)

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Data    IfData
}

type IfData struct {
	Type       uint8
	Typelen    uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Recvquota  uint8
	Xmitquota  uint8
	Unused1    uint8
	Mtu        uint32
	Metric     uint32
	Baudrate   uint32
	Ipackets   uint32
	Ierrors    uint32
	Opackets   uint32
	Oerrors    uint32
	Collisions uint32
	Ibytes     uint32
	Obytes     uint32
	Imcasts    uint32
	Omcasts    uint32
	Iqdrops    uint32
	Noproto    uint32
	Recvtiming uint32
	Xmittiming uint32
	Lastchange Timeval32
	Unused2    uint32
	Hwassist   uint32
	Reserved1  uint32
	Reserved2  uint32
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Metric  int32
}

type IfmaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       [2]byte
}

type IfmaMsghdr2 struct {
	Msglen   uint16
	Version  uint8
	Type     uint8
	Addrs    int32
	Flags    int32
	Index    uint16
	Refcount int32
}

type RtMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Flags   int32
	Addrs   int32
	Pid     int32
	Seq     int32
	Errno   int32
	Use     int32
	Inits   uint32
	Rmx     RtMetrics
}

type RtMetrics struct {
	Locks    uint32
	Mtu      uint32
	Hopcount uint32
	Expire   int32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pksent   uint32
	State    uint32
	Filler   [3]uint32
}

const (
	SizeofBpfVersion = 0x4
	SizeofBpfStat    = 0x8
	SizeofBpfProgram = 0x10
	SizeofBpfInsn    = 0x8
	SizeofBpfHdr     = 0x14
)

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint32
	Drop uint32
}

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfInsn struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type BpfHdr struct {
	Tstamp  Timeval32
	Caplen  uint32
	Datalen uint32
	Hdrlen  uint16
	_       [2]byte
}

type Termios struct {
	Iflag  uint64
	Oflag  uint64
	Cflag  uint64
	Lflag  uint64
	Cc     [20]uint8
	Ispeed uint64
	Ospeed uint64
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

const (
	AT_FDCWD            = -0x2
	AT_REMOVEDIR        = 0x80
	AT_SYMLINK_FOLLOW   = 0x40
	AT_SYMLINK_NOFOLLOW = 0x20
)

type PollFd struct {
	Fd      int32
	Events  int16
	Revents int16
}

const (
	POLLERR    = 0x8
	POLLHUP    = 0x10
	POLLIN     = 0x1
	POLLNVAL   = 0x20
	POLLOUT    = 0x4
	POLLPRI    = 0x2
	POLLRDBAND = 0x80
	POLLRDNORM = 0x40
	POLLWRBAND = 0x100
	POLLWRNORM = 0x4
)

type Utsname struct {
	Sysname  [256]byte
	Nodename [256]byte
	Release  [256]byte
	Version  [256]byte
	Machine  [256]byte
}

const SizeofClockinfo = 0x14

type Clockinfo struct {
	Hz      int32
	Tick    int32
	Tickadj int32
	Stathz  int32
	Profhz  int32
}

type CtlInfo struct {
	Id   uint32
	Name [96]byte
}

const SizeofKinfoProc = 0x288

type Eproc struct {
	Paddr   uintptr
	Sess    uintptr
	Pcred   Pcred
	Ucred   Ucred
	Vm      Vmspace
	Ppid    int32
	Pgid    int32
	Jobc    int16
	Tdev    int32
	Tpgid   int32
	Tsess   uintptr
	Wmesg   [8]int8
	Xsize   int32
	Xrssize int16
	Xccount int16
	Xswrss  int16
	Flag    int32
	Login   [12]int8
	Spare   [4]int32
	_       [4]byte
}

type ExternProc struct {
	P_starttime Timeval
	P_vmspace   *Vmspace
	P_sigacts   uintptr
	P_flag      int32
	P_stat      int8
	P_pid       int32
	P_oppid     int32
	P_dupfd     int32
	User_stack  *int8
	Exit_thread *byte
	P_debugger  int32
	Sigwait     int32
	P_estcpu    uint32
	P_cpticks   int32
	P_pctcpu    uint32
	P_wchan     *byte
	P_wmesg     *int8
	P_swtime    uint32
	P_slptime   uint32
	P_realtimer Itimerval
	P_rtime     Timeval
	P_uticks    uint64
	P_sticks    uint64
	P_iticks    uint64
	P_traceflag int32
	P_tracep    uintptr
	P_siglist   int32
	P_textvp    uintptr
	P_holdcnt   int32
	P_sigmask   uint32
	P_sigignore uint32
	P_sigcatch  uint32
	P_priority  uint8
	P_usrpri    uint8
	P_nice      int8
	P_comm      [17]int8
	P_pgrp      uintptr
	P_addr      uintptr
	P_xstat     uint16
	P_acflag    uint16
	P_ru        *Rusage
}

type Itimerval struct {
	Interval Timeval
	Value    Timeval
}

type KinfoProc struct {
	Proc  ExternProc
	Eproc Eproc
}

type Vmspace struct {
	Dummy  int32
	Dummy2 *int8
	Dummy3 [5]int32
	Dummy4 [3]*int8
}

type Pcred struct {
	Pc_lock  [72]int8
	Pc_ucred uintptr
	P_ruid   uint32
	P_svuid  uint32
	P_rgid   uint32
	P_svgid  uint32
	P_refcnt int32
	_        [4]byte
}

type Ucred struct {
	Ref     int32
	Uid     uint32
	Ngroups int16
	Groups  [16]uint32
}
