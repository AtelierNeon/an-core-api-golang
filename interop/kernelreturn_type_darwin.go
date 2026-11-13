package interop

type KernelReturnType int32

/* https://github.com/apple/darwin-xnu/blob/main/osfmk/mach/kern_return.h */
const (
	KernelReturnTypeSuccess                KernelReturnType = /* KERN_SUCCESS                  */ 0x0
	KernelReturnTypeInvalidAddress         KernelReturnType = /* KERN_INVALID_ADDRESS          */ 0x1
	KernelReturnTypeProtectionFailure      KernelReturnType = /* KERN_PROTECTION_FAILURE       */ 0x2
	KernelReturnTypeNoSpace                KernelReturnType = /* KERN_NO_SPACE                 */ 0x3
	KernelReturnTypeInvalidArgument        KernelReturnType = /* KERN_INVALID_ARGUMENT         */ 0x4
	KernelReturnTypeFailure                KernelReturnType = /* KERN_FAILURE                  */ 0x5
	KernelReturnTypeResourceShortage       KernelReturnType = /* KERN_RESOURCE_SHORTAGE        */ 0x6
	KernelReturnTypeNotReceiver            KernelReturnType = /* KERN_NOT_RECEIVER             */ 0x7
	KernelReturnTypeNoAccess               KernelReturnType = /* KERN_NO_ACCESS                */ 0x8
	KernelReturnTypeMemoryFailure          KernelReturnType = /* KERN_MEMORY_FAILURE           */ 0x9
	KernelReturnTypeMemoryError            KernelReturnType = /* KERN_MEMORY_ERROR             */ 0xa
	KernelReturnTypeAlreadyInSet           KernelReturnType = /* KERN_ALREADY_IN_SET           */ 0xb
	KernelReturnTypeNotInSet               KernelReturnType = /* KERN_NOT_IN_SET               */ 0xc
	KernelReturnTypeNameExists             KernelReturnType = /* KERN_NAME_EXISTS              */ 0xd
	KernelReturnTypeAborted                KernelReturnType = /* KERN_ABORTED                  */ 0xe
	KernelReturnTypeInvalidName            KernelReturnType = /* KERN_INVALID_NAME             */ 0xf
	KernelReturnTypeInvalidTask            KernelReturnType = /* KERN_INVALID_TASK             */ 0x10
	KernelReturnTypeInvalidRight           KernelReturnType = /* KERN_INVALID_RIGHT            */ 0x11
	KernelReturnTypeInvalidValue           KernelReturnType = /* KERN_INVALID_VALUE            */ 0x12
	KernelReturnTypeUserReferencesOverflow KernelReturnType = /* KERN_UREFS_OVERFLOW           */ 0x13
	KernelReturnTypeInvalidCapability      KernelReturnType = /* KERN_INVALID_CAPABILITY       */ 0x14
	KernelReturnTypeRightExists            KernelReturnType = /* KERN_RIGHT_EXISTS             */ 0x15
	KernelReturnTypeInvalidHost            KernelReturnType = /* KERN_INVALID_HOST             */ 0x16
	KernelReturnTypeMemoryPresent          KernelReturnType = /* KERN_MEMORY_PRESENT           */ 0x17
	KernelReturnTypeMemoryDataMoved        KernelReturnType = /* KERN_MEMORY_DATA_MOVED        */ 0x18
	KernelReturnTypeMemoryRestartCopy      KernelReturnType = /* KERN_MEMORY_RESTART_COPY      */ 0x19
	KernelReturnTypeInvalidProcessorSet    KernelReturnType = /* KERN_INVALID_PROCESSOR_SET    */ 0x1a
	KernelReturnTypePolicyLimit            KernelReturnType = /* KERN_POLICY_LIMIT             */ 0x1b
	KernelReturnTypeInvalidPolicy          KernelReturnType = /* KERN_INVALID_POLICY           */ 0x1c
	KernelReturnTypeInvalidObject          KernelReturnType = /* KERN_INVALID_OBJECT           */ 0x1d
	KernelReturnTypeAlreadyWaiting         KernelReturnType = /* KERN_ALREADY_WAITING          */ 0x1e
	KernelReturnTypeDefaultSet             KernelReturnType = /* KERN_DEFAULT_SET              */ 0x1f
	KernelReturnTypeExceptionProtected     KernelReturnType = /* KERN_EXCEPTION_PROTECTED      */ 0x20
	KernelReturnTypeInvalidLedger          KernelReturnType = /* KERN_INVALID_LEDGER           */ 0x21
	KernelReturnTypeInvalidMemoryControl   KernelReturnType = /* KERN_INVALID_MEMORY_CONTROL   */ 0x22
	KernelReturnTypeInvalidSecurity        KernelReturnType = /* KERN_INVALID_SECURITY         */ 0x23
	KernelReturnTypeNotDepressed           KernelReturnType = /* KERN_NOT_DEPRESSED            */ 0x24
	KernelReturnTypeTerminated             KernelReturnType = /* KERN_TERMINATED               */ 0x25
	KernelReturnTypeLockSetDestroyed       KernelReturnType = /* KERN_LOCK_SET_DESTROYED       */ 0x26
	KernelReturnTypeLockUnstable           KernelReturnType = /* KERN_LOCK_UNSTABLE            */ 0x27
	KernelReturnTypeLockOwned              KernelReturnType = /* KERN_LOCK_OWNED               */ 0x28
	KernelReturnTypeLockOwnedSelf          KernelReturnType = /* KERN_LOCK_OWNED_SELF          */ 0x29
	KernelReturnTypeSemaphoreDestroyed     KernelReturnType = /* KERN_SEMAPHORE_DESTROYED      */ 0x2a
	KernelReturnTypeRPCServerTerminated    KernelReturnType = /* KERN_RPC_SERVER_TERMINATED    */ 0x2b
	KernelReturnTypeRPCTerminateOrphan     KernelReturnType = /* KERN_RPC_TERMINATE_ORPHAN     */ 0x2c
	KernelReturnTypeRPCContinueOrphan      KernelReturnType = /* KERN_RPC_CONTINUE_ORPHAN      */ 0x2d
	KernelReturnTypeNotSupported           KernelReturnType = /* KERN_NOT_SUPPORTED            */ 0x2e
	KernelReturnTypeNodeDown               KernelReturnType = /* KERN_NODE_DOWN                */ 0x2f
	KernelReturnTypeNotWaiting             KernelReturnType = /* KERN_NOT_WAITING              */ 0x30
	KernelReturnTypeOperationTimedOut      KernelReturnType = /* KERN_OPERATION_TIMED_OUT      */ 0x31
	KernelReturnTypeCodesignError          KernelReturnType = /* KERN_CODESIGN_ERROR           */ 0x32
	KernelReturnTypePolicyStatic           KernelReturnType = /* KERN_POLICY_STATIC            */ 0x33
	KernelReturnTypeInsufficientBufferSize KernelReturnType = /* KERN_INSUFFICIENT_BUFFER_SIZE */ 0x34
	KernelReturnTypeDenied                 KernelReturnType = /* KERN_DENIED                   */ 0x35
	KernelReturnTypeMissingKC              KernelReturnType = /* KERN_MISSING_KC               */ 0x36
	KernelReturnTypeInvalidKC              KernelReturnType = /* KERN_INVALID_KC               */ 0x37
	KernelReturnTypeReturnMax              KernelReturnType = /* KERN_RETURN_MAX               */ 0x100
)

var kernelReturnType_typeStringMap = map[KernelReturnType]string{
	KernelReturnTypeSuccess:                "Success",
	KernelReturnTypeInvalidAddress:         "Specified address is not currently valid.",
	KernelReturnTypeProtectionFailure:      "Specified memory is valid, but does not permit the required forms of access.",
	KernelReturnTypeNoSpace:                "The address range specified is already in use, or no address range of the size specified could be found.",
	KernelReturnTypeInvalidArgument:        "The function requested was not applicable to this type of argument, or an argument is invalid.",
	KernelReturnTypeFailure:                "The function could not be performed. A catch-all.",
	KernelReturnTypeResourceShortage:       "A system resource could not be allocated to fulfill this request. This failure may not be permanent.",
	KernelReturnTypeNotReceiver:            "The task in question does not hold receive rights for the port argument.",
	KernelReturnTypeNoAccess:               "Bogus access restriction.",
	KernelReturnTypeMemoryFailure:          "During a page fault, the target address refers to a memory object that has been destroyed. This failure is permanent.",
	KernelReturnTypeMemoryError:            "During a page fault, the memory object indicated that the data could not be returned. This failure may be temporary; future attempts to access this same data may succeed, as defined by the memory object.",
	KernelReturnTypeAlreadyInSet:           "The receive right is already a member of the portset.",
	KernelReturnTypeNotInSet:               "The receive right is not a member of a port set.",
	KernelReturnTypeNameExists:             "The name already denotes a right in the task.",
	KernelReturnTypeAborted:                "The operation was aborted. Ipc code will catch this and reflect it as a message error.",
	KernelReturnTypeInvalidName:            "The name doesn’t denote a right in the task.",
	KernelReturnTypeInvalidTask:            "Target task isn’t an active task.",
	KernelReturnTypeInvalidRight:           "The name denotes a right, but not an appropriate right.",
	KernelReturnTypeInvalidValue:           "A blatant range error.",
	KernelReturnTypeUserReferencesOverflow: "Operation would overflow limit on user-references.",
	KernelReturnTypeInvalidCapability:      "The supplied (port) capability is improper.",
	KernelReturnTypeRightExists:            "The task already has send or receive rights for the port under another name.",
	KernelReturnTypeInvalidHost:            "Target host isn’t actually a host.",
	KernelReturnTypeMemoryPresent:          "An attempt was made to supply “precious” data for memory that is already present in a memory object.",
	KernelReturnTypeMemoryDataMoved:        "A page was requested of a memory manager via memory_object_data_request for an object using a MEMORY_OBJECT_COPY_CALL strategy, with the VM_PROT_WANTS_COPY flag being used to specify that the page desired is for a copy of the object, and the memory manager has detected the page was pushed into a copy of the object while the kernel was walking the shadow chain from the copy to the object. This error code is delivered via memory_object_data_error and is handled by the kernel (it forces the kernel to restart the fault). It will not be seen by users.",
	KernelReturnTypeMemoryRestartCopy:      "A strategic copy was attempted of an object upon which a quicker copy is now possible. The caller should retry the copy using vm_object_copy_quickly. This error code is seen only by the kernel.",
	KernelReturnTypeInvalidProcessorSet:    "An argument applied to assert processor set privilege was not a processor set control port.",
	KernelReturnTypePolicyLimit:            "The specified scheduling attributes exceed the thread’s limits.",
	KernelReturnTypeInvalidPolicy:          "The specified scheduling policy is not currently enabled for the processor set.",
	KernelReturnTypeInvalidObject:          "The external memory manager failed to initialize the memory object.",
	KernelReturnTypeAlreadyWaiting:         "A thread is attempting to wait for an event for which there is already a waiting thread.",
	KernelReturnTypeDefaultSet:             "An attempt was made to destroy the default processor set.",
	KernelReturnTypeExceptionProtected:     "An attempt was made to fetch an exception port that is protected, or to abort a thread while processing a protected exception.",
	KernelReturnTypeInvalidLedger:          "A ledger was required but not supplied.",
	KernelReturnTypeInvalidMemoryControl:   "The port was not a memory cache control port.",
	KernelReturnTypeInvalidSecurity:        "An argument supplied to assert security privilege was not a host security port.",
	KernelReturnTypeNotDepressed:           "thread_depress_abort was called on a thread which was not currently depressed.",
	KernelReturnTypeTerminated:             "Object has been terminated and is no longer available.",
	KernelReturnTypeLockSetDestroyed:       "Lock set has been destroyed and is no longer available.",
	KernelReturnTypeLockUnstable:           "The thread holding the lock terminated before releasing the lock.",
	KernelReturnTypeLockOwned:              "The lock is already owned by another thread.",
	KernelReturnTypeLockOwnedSelf:          "The lock is already owned by the calling thread.",
	KernelReturnTypeSemaphoreDestroyed:     "Semaphore has been destroyed and is no longer available.",
	KernelReturnTypeRPCServerTerminated:    "Return from RPC indicating the target server was terminated before it successfully replied.",
	KernelReturnTypeRPCTerminateOrphan:     "Terminate an orphaned activation.",
	KernelReturnTypeRPCContinueOrphan:      "Allow an orphaned activation to continue executing.",
	KernelReturnTypeNotSupported:           "Empty thread activation (No thread linked to it).",
	KernelReturnTypeNodeDown:               "Remote node down or inaccessible.",
	KernelReturnTypeNotWaiting:             "A signalled thread was not actually waiting.",
	KernelReturnTypeOperationTimedOut:      "Some thread-oriented operation (semaphore_wait) timed out.",
	KernelReturnTypeCodesignError:          "During a page fault, indicates that the page was rejected as a result of a signature check.",
	KernelReturnTypePolicyStatic:           "The requested property cannot be changed at this time.",
	KernelReturnTypeInsufficientBufferSize: "The provided buffer is of insufficient size for the requested data.",
	KernelReturnTypeDenied:                 "Denied by security policy.",
	KernelReturnTypeMissingKC:              "The KC on which the function is operating is missing.",
	KernelReturnTypeInvalidKC:              "The KC on which the function is operating is invalid.",
	KernelReturnTypeReturnMax:              "Maximum return value allowable.",
}

func (krct KernelReturnType) String() string {
	str, ok := kernelReturnType_typeStringMap[krct]
	if ok {
		return str
	}

	return "Unknown"
}
