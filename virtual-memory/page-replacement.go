package virtual_memory

type IPageReplacement interface {
	GetPageToReplace(vm *VirtualMemory) int
}
