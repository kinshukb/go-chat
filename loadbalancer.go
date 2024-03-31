package load

import "log"


// A weighted approach to balance loads
type Listener struct {
	NCname string `json:"ncname"`
	Id     string `json:"id"`
	Weight int    `json:"weight"`
}


//List of Weighted Listener for load control.
var Listeners []Listener


// Imp: We need both the cached one and Listeners balancing, will be selected from below array, cache will help in the process.
func selectListener(ncname string) string {
	var final string
	for i := 0; i < len(Listeners); i++ {
		if Listeners[i].Weight == 0 && Listeners[i].NCname == ncname {
			Listeners[i].Weight = Listeners[i].Weight + 1
			final = Listeners[i].Id
			return final
		}
		for j := i + 1; j < len(Listeners); j++ {
			if Listeners[j].Weight < Listeners[i].Weight && Listeners[j].NCname == ncname {
				log.Println("-----------balancing---------", i))
				Listeners[j].Weight = Listeners[j].Weight + 1
				final = Listeners[j].Id
				return final
			}
		}
		if Listeners[i].NCname == ncname {
			logger.Debug("already balanced :", zap.Int("i: ", i))
			Listeners[i].Weight = Listeners[i].Weight + 1
			final = Listeners[i].Id
			return final
		}

	}
	return final
}


func caller() {
giveMeListenerId :=  selectListener("demo")
  // TODO: select pg_notify("demo", {}&payload)

}
