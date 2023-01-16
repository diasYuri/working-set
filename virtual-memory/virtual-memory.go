package virtual_memory

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type VirtualMemory struct {
	pageReplacement    IPageReplacement
	nextPageId         int
	memory             *Memory
	pageMap            map[int]map[int]int
	memoryProcessIdMap map[int]int
	lastAccess         map[int]time.Time
}

func NewVirtualMemory(pr IPageReplacement, numPages int) *VirtualMemory {
	return &VirtualMemory{
		pageReplacement:    pr,
		pageMap:            make(map[int]map[int]int),
		memoryProcessIdMap: make(map[int]int),
		lastAccess:         make(map[int]time.Time),
		memory:             NewMemory(numPages),
		nextPageId:         1,
	}
}

func (vm *VirtualMemory) GetNextPageId() int {
	nextId := vm.nextPageId
	vm.nextPageId++
	return nextId
}

func (vm *VirtualMemory) AccessPage(processId int, processPageId int) *[]byte {
	processMemoryMap, existsProcessMap := vm.pageMap[processId]
	if !existsProcessMap {
		vm.pageMap[processId] = make(map[int]int)
		processMemoryMap = vm.pageMap[processId]
	}

	pageId, existPage := processMemoryMap[processPageId]
	if !existPage {
		newPageId := vm.GetNextPageId()
		processMemoryMap[processPageId] = newPageId
		pageId = newPageId
	}

	vm.memoryProcessIdMap[pageId] = processId
	pageData := vm.loadPageId(pageId)
	vm.lastAccess[processId] = time.Now()
	return pageData
}

func (vm *VirtualMemory) loadPage(idx int, page *Page) {
	vm.memory.pages[idx] = page
}

func (vm *VirtualMemory) loadPageId(pageId int) *[]byte {
	hasEmptySlot := false
	idxEmptySlot := -1

	for i, page := range vm.memory.pages {
		if page != nil && page.id == pageId {
			log.Infoln("page found in memory")
			return page.Access()
		}
		if page == nil && !hasEmptySlot {
			hasEmptySlot = true
			idxEmptySlot = i
		}
	}

	page := NewPage(pageId)

	if hasEmptySlot {
		vm.loadPage(idxEmptySlot, page)
		return page.Access()
	}

	log.Warnf("page fault")
	pr := vm.pageReplacement
	idxToReplace := pr.GetPageToReplace(vm)
	vm.loadPage(idxToReplace, page)
	return page.Access()
}
