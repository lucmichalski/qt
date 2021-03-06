package qt

import (
	"runtime"
	"strconv"
	"strings"
	"sync"

	_ "github.com/therecipe/qt/internal/android"
)

var (
	signals      = make(map[string]interface{})
	signalsMutex = new(sync.Mutex)
	ids          int
)

func init() { runtime.LockOSThread() }

func ConnectSignal(name, signal string, function interface{}) {
	signalsMutex.Lock()
	signals[name+":"+signal] = function
	signalsMutex.Unlock()
}

func GetSignal(name, signal string) interface{} {
	if signal == "destroyed" || signal == "deleteLater" {
		defer DisconnectAllSignals(name)
	}
	var s interface{}
	signalsMutex.Lock()
	s = signals[name+":"+signal]
	signalsMutex.Unlock()
	return s
}

func DisconnectSignal(name, signal string) {
	signalsMutex.Lock()
	delete(signals, name+":"+signal)
	signalsMutex.Unlock()
}

func DisconnectAllSignals(name string) {
	for entry := range signals {
		if strings.Contains(entry, name) {
			signalsMutex.Lock()
			delete(signals, entry)
			signalsMutex.Unlock()
		}
	}
}

func Identifier() string {
	ids++
	return strconv.Itoa(ids)
}

func DumpSignals() {
	println("##############################\tDUMP_SIGNALTABLE_START\t##############################")
	for entry := range signals {
		println(entry)
	}
	println("##############################\tDUMP_SIGNALTABLE_END\t##############################")
}

func GoBoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Recovering(fn string) {
	if recover() != nil {
		println("recovered in:", fn)
	}
}
