package interop

/* https://github.com/apple-oss-distributions/xnu/blob/main/iokit/DriverKit/IOReturn.h */
type KernelReturn DarwinInt

/* https://github.com/apple-oss-distributions/xnu/blob/main/osfmk/mach/kern_return.h */
const (
	KernelReturnSuccess                KernelReturn = /* KERN_SUCCESS                  */ 0x0
	KernelReturnInvalidAddress         KernelReturn = /* KERN_INVALID_ADDRESS          */ 0x1
	KernelReturnProtectionFailure      KernelReturn = /* KERN_PROTECTION_FAILURE       */ 0x2
	KernelReturnNoSpace                KernelReturn = /* KERN_NO_SPACE                 */ 0x3
	KernelReturnInvalidArgument        KernelReturn = /* KERN_INVALID_ARGUMENT         */ 0x4
	KernelReturnFailure                KernelReturn = /* KERN_FAILURE                  */ 0x5
	KernelReturnResourceShortage       KernelReturn = /* KERN_RESOURCE_SHORTAGE        */ 0x6
	KernelReturnNotReceiver            KernelReturn = /* KERN_NOT_RECEIVER             */ 0x7
	KernelReturnNoAccess               KernelReturn = /* KERN_NO_ACCESS                */ 0x8
	KernelReturnMemoryFailure          KernelReturn = /* KERN_MEMORY_FAILURE           */ 0x9
	KernelReturnMemoryError            KernelReturn = /* KERN_MEMORY_ERROR             */ 0xa
	KernelReturnAlreadyInSet           KernelReturn = /* KERN_ALREADY_IN_SET           */ 0xb
	KernelReturnNotInSet               KernelReturn = /* KERN_NOT_IN_SET               */ 0xc
	KernelReturnNameExists             KernelReturn = /* KERN_NAME_EXISTS              */ 0xd
	KernelReturnAborted                KernelReturn = /* KERN_ABORTED                  */ 0xe
	KernelReturnInvalidName            KernelReturn = /* KERN_INVALID_NAME             */ 0xf
	KernelReturnInvalidTask            KernelReturn = /* KERN_INVALID_TASK             */ 0x10
	KernelReturnInvalidRight           KernelReturn = /* KERN_INVALID_RIGHT            */ 0x11
	KernelReturnInvalidValue           KernelReturn = /* KERN_INVALID_VALUE            */ 0x12
	KernelReturnUserReferencesOverflow KernelReturn = /* KERN_UREFS_OVERFLOW           */ 0x13
	KernelReturnInvalidCapability      KernelReturn = /* KERN_INVALID_CAPABILITY       */ 0x14
	KernelReturnRightExists            KernelReturn = /* KERN_RIGHT_EXISTS             */ 0x15
	KernelReturnInvalidHost            KernelReturn = /* KERN_INVALID_HOST             */ 0x16
	KernelReturnMemoryPresent          KernelReturn = /* KERN_MEMORY_PRESENT           */ 0x17
	KernelReturnMemoryDataMoved        KernelReturn = /* KERN_MEMORY_DATA_MOVED        */ 0x18
	KernelReturnMemoryRestartCopy      KernelReturn = /* KERN_MEMORY_RESTART_COPY      */ 0x19
	KernelReturnInvalidProcessorSet    KernelReturn = /* KERN_INVALID_PROCESSOR_SET    */ 0x1a
	KernelReturnPolicyLimit            KernelReturn = /* KERN_POLICY_LIMIT             */ 0x1b
	KernelReturnInvalidPolicy          KernelReturn = /* KERN_INVALID_POLICY           */ 0x1c
	KernelReturnInvalidObject          KernelReturn = /* KERN_INVALID_OBJECT           */ 0x1d
	KernelReturnAlreadyWaiting         KernelReturn = /* KERN_ALREADY_WAITING          */ 0x1e
	KernelReturnDefaultSet             KernelReturn = /* KERN_DEFAULT_SET              */ 0x1f
	KernelReturnExceptionProtected     KernelReturn = /* KERN_EXCEPTION_PROTECTED      */ 0x20
	KernelReturnInvalidLedger          KernelReturn = /* KERN_INVALID_LEDGER           */ 0x21
	KernelReturnInvalidMemoryControl   KernelReturn = /* KERN_INVALID_MEMORY_CONTROL   */ 0x22
	KernelReturnInvalidSecurity        KernelReturn = /* KERN_INVALID_SECURITY         */ 0x23
	KernelReturnNotDepressed           KernelReturn = /* KERN_NOT_DEPRESSED            */ 0x24
	KernelReturnTerminated             KernelReturn = /* KERN_TERMINATED               */ 0x25
	KernelReturnLockSetDestroyed       KernelReturn = /* KERN_LOCK_SET_DESTROYED       */ 0x26
	KernelReturnLockUnstable           KernelReturn = /* KERN_LOCK_UNSTABLE            */ 0x27
	KernelReturnLockOwned              KernelReturn = /* KERN_LOCK_OWNED               */ 0x28
	KernelReturnLockOwnedSelf          KernelReturn = /* KERN_LOCK_OWNED_SELF          */ 0x29
	KernelReturnSemaphoreDestroyed     KernelReturn = /* KERN_SEMAPHORE_DESTROYED      */ 0x2a
	KernelReturnRPCServerTerminated    KernelReturn = /* KERN_RPC_SERVER_TERMINATED    */ 0x2b
	KernelReturnRPCTerminateOrphan     KernelReturn = /* KERN_RPC_TERMINATE_ORPHAN     */ 0x2c
	KernelReturnRPCContinueOrphan      KernelReturn = /* KERN_RPC_CONTINUE_ORPHAN      */ 0x2d
	KernelReturnNotSupported           KernelReturn = /* KERN_NOT_SUPPORTED            */ 0x2e
	KernelReturnNodeDown               KernelReturn = /* KERN_NODE_DOWN                */ 0x2f
	KernelReturnNotWaiting             KernelReturn = /* KERN_NOT_WAITING              */ 0x30
	KernelReturnOperationTimedOut      KernelReturn = /* KERN_OPERATION_TIMED_OUT      */ 0x31
	KernelReturnCodesignError          KernelReturn = /* KERN_CODESIGN_ERROR           */ 0x32
	KernelReturnPolicyStatic           KernelReturn = /* KERN_POLICY_STATIC            */ 0x33
	KernelReturnInsufficientBufferSize KernelReturn = /* KERN_INSUFFICIENT_BUFFER_SIZE */ 0x34
	KernelReturnDenied                 KernelReturn = /* KERN_DENIED                   */ 0x35
	KernelReturnMissingKC              KernelReturn = /* KERN_MISSING_KC               */ 0x36
	KernelReturnInvalidKC              KernelReturn = /* KERN_INVALID_KC               */ 0x37
	KernelReturnNotFound               KernelReturn = /* KERN_NOT_FOUND                */ 0x38
	KernelReturnReturnMax              KernelReturn = /* KERN_RETURN_MAX               */ 0x100
)

var kernelReturn_typeStringMap = map[KernelReturn]string{
	KernelReturnSuccess:                "Success",
	KernelReturnInvalidAddress:         "Specified address is not currently valid.",
	KernelReturnProtectionFailure:      "Specified memory is valid, but does not permit the required forms of access.",
	KernelReturnNoSpace:                "The address range specified is already in use, or no address range of the size specified could be found.",
	KernelReturnInvalidArgument:        "The function requested was not applicable to this type of argument, or an argument is invalid.",
	KernelReturnFailure:                "The function could not be performed. A catch-all.",
	KernelReturnResourceShortage:       "A system resource could not be allocated to fulfill this request. This failure may not be permanent.",
	KernelReturnNotReceiver:            "The task in question does not hold receive rights for the port argument.",
	KernelReturnNoAccess:               "Bogus access restriction.",
	KernelReturnMemoryFailure:          "During a page fault, the target address refers to a memory object that has been destroyed. This failure is permanent.",
	KernelReturnMemoryError:            "During a page fault, the memory object indicated that the data could not be returned. This failure may be temporary; future attempts to access this same data may succeed, as defined by the memory object.",
	KernelReturnAlreadyInSet:           "The receive right is already a member of the portset.",
	KernelReturnNotInSet:               "The receive right is not a member of a port set.",
	KernelReturnNameExists:             "The name already denotes a right in the task.",
	KernelReturnAborted:                "The operation was aborted. Ipc code will catch this and reflect it as a message error.",
	KernelReturnInvalidName:            "The name doesn't denote a right in the task.",
	KernelReturnInvalidTask:            "Target task isn't an active task.",
	KernelReturnInvalidRight:           "The name denotes a right, but not an appropriate right.",
	KernelReturnInvalidValue:           "A blatant range error.",
	KernelReturnUserReferencesOverflow: "Operation would overflow limit on user-references.",
	KernelReturnInvalidCapability:      "The supplied (port) capability is improper.",
	KernelReturnRightExists:            "The task already has send or receive rights for the port under another name.",
	KernelReturnInvalidHost:            "Target host isn't actually a host.",
	KernelReturnMemoryPresent:          "An attempt was made to supply “precious” data for memory that is already present in a memory object.",
	KernelReturnMemoryDataMoved:        "A page was requested of a memory manager via memory_object_data_request for an object using a MEMORY_OBJECT_COPY_CALL strategy, with the VM_PROT_WANTS_COPY flag being used to specify that the page desired is for a copy of the object, and the memory manager has detected the page was pushed into a copy of the object while the kernel was walking the shadow chain from the copy to the object. This error code is delivered via memory_object_data_error and is handled by the kernel (it forces the kernel to restart the fault). It will not be seen by users.",
	KernelReturnMemoryRestartCopy:      "A strategic copy was attempted of an object upon which a quicker copy is now possible. The caller should retry the copy using vm_object_copy_quickly. This error code is seen only by the kernel.",
	KernelReturnInvalidProcessorSet:    "An argument applied to assert processor set privilege was not a processor set control port.",
	KernelReturnPolicyLimit:            "The specified scheduling attributes exceed the thread's limits.",
	KernelReturnInvalidPolicy:          "The specified scheduling policy is not currently enabled for the processor set.",
	KernelReturnInvalidObject:          "The external memory manager failed to initialize the memory object.",
	KernelReturnAlreadyWaiting:         "A thread is attempting to wait for an event for which there is already a waiting thread.",
	KernelReturnDefaultSet:             "An attempt was made to destroy the default processor set.",
	KernelReturnExceptionProtected:     "An attempt was made to fetch an exception port that is protected, or to abort a thread while processing a protected exception.",
	KernelReturnInvalidLedger:          "A ledger was required but not supplied.",
	KernelReturnInvalidMemoryControl:   "The port was not a memory cache control port.",
	KernelReturnInvalidSecurity:        "An argument supplied to assert security privilege was not a host security port.",
	KernelReturnNotDepressed:           "thread_depress_abort was called on a thread which was not currently depressed.",
	KernelReturnTerminated:             "Object has been terminated and is no longer available.",
	KernelReturnLockSetDestroyed:       "Lock set has been destroyed and is no longer available.",
	KernelReturnLockUnstable:           "The thread holding the lock terminated before releasing the lock.",
	KernelReturnLockOwned:              "The lock is already owned by another thread.",
	KernelReturnLockOwnedSelf:          "The lock is already owned by the calling thread.",
	KernelReturnSemaphoreDestroyed:     "Semaphore has been destroyed and is no longer available.",
	KernelReturnRPCServerTerminated:    "Return from RPC indicating the target server was terminated before it successfully replied.",
	KernelReturnRPCTerminateOrphan:     "Terminate an orphaned activation.",
	KernelReturnRPCContinueOrphan:      "Allow an orphaned activation to continue executing.",
	KernelReturnNotSupported:           "Empty thread activation (No thread linked to it).",
	KernelReturnNodeDown:               "Remote node down or inaccessible.",
	KernelReturnNotWaiting:             "A signalled thread was not actually waiting.",
	KernelReturnOperationTimedOut:      "Some thread-oriented operation (semaphore_wait) timed out.",
	KernelReturnCodesignError:          "During a page fault, indicates that the page was rejected as a result of a signature check.",
	KernelReturnPolicyStatic:           "The requested property cannot be changed at this time.",
	KernelReturnInsufficientBufferSize: "The provided buffer is of insufficient size for the requested data.",
	KernelReturnDenied:                 "Denied by security policy.",
	KernelReturnMissingKC:              "The KC on which the function is operating is missing.",
	KernelReturnInvalidKC:              "The KC on which the function is operating is invalid.",
	KernelReturnNotFound:               "A search or query operation did not return a result.",
	KernelReturnReturnMax:              "Maximum return value allowable.",
}

func (krt KernelReturn) String() string {
	str, ok := kernelReturn_typeStringMap[krt]
	if ok {
		return str
	}

	return "Unknown"
}
