#include "textflag.h"

// int proc_listpids(int pid, void* buffer, uint32_t buffersize)
GLOBL ·libc_proc_listpids_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_listpids_trampoline_addr(SB)/8, $libc_proc_listpids_trampoline<>(SB)

TEXT libc_proc_listpids_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_listpids(SB)

// int proc_pidfdinfo(int pid, int fd, int flavor, void* buffer, int buffersize)
GLOBL ·libc_proc_pidfdinfo_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_pidfdinfo_trampoline_addr(SB)/8, $libc_proc_pidfdinfo_trampoline<>(SB)

TEXT libc_proc_pidfdinfo_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_pidfdinfo(SB)

// int proc_pidinfo(int pid, int flavor, uint64_t arg, void* buffer, int buffersize)
GLOBL ·libc_proc_pidinfo_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_pidinfo_trampoline_addr(SB)/8, $libc_proc_pidinfo_trampoline<>(SB)

TEXT libc_proc_pidinfo_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_pidinfo(SB)

// int proc_pidpath(int pid, void* buffer, uint32_t buffersize)
GLOBL ·libc_proc_pidpath_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_pidpath_trampoline_addr(SB)/8, $libc_proc_pidpath_trampoline<>(SB)

TEXT libc_proc_pidpath_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_pidpath(SB)
