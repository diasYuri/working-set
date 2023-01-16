package virtual_memory

import (
	log "github.com/sirupsen/logrus"
	"time"
)

const ticket = time.Nanosecond

type WorkingSet struct {
	timeWindow time.Duration
}

func NewWorkingSet(timeWindow time.Duration) *WorkingSet {
	return &WorkingSet{
		timeWindow: timeWindow * time.Nanosecond,
	}
}

func (ws *WorkingSet) GetPageToReplace(vm *VirtualMemory) int {
	olderTime := time.Duration(0)
	idxOlderPage := 0

	for i, page := range vm.memory.pages {
		processId := vm.memoryProcessIdMap[page.id]
		lastAccess := vm.lastAccess[processId]

		timeWindows := lastAccess.Sub(page.lastAccess)

		if timeWindows > ws.timeWindow {
			log.Warnf("replacing page %d of process %d\n", page.id, processId)
			return i
		}

		if timeWindows > olderTime {
			olderTime = timeWindows
			idxOlderPage = i
		}
	}

	log.Warnf("replacing index %d", idxOlderPage)
	return idxOlderPage
}
