package load

import "log"


// A weighted approach to balance loads
type ListenerPod struct {
	NCname string `json:"ncname"`
	PodId  string `json:"podid"`
	Weight int    `json:"weight"`
}


//List of Weighted ListenerPods for load control.
var ListenerPods []ListenerPod


// Imp: We need both the cached one and ListenerPods balancing, pod will be selected from below array, cache will help in the process.
func selectListener(ncname string) string {
	var finalPodId string
	for i := 0; i < len(ListenerPods); i++ {
		if ListenerPods[i].Weight == 0 && ListenerPods[i].NCname == ncname {
			ListenerPods[i].Weight = ListenerPods[i].Weight + 1
			finalPodId = ListenerPods[i].PodId
			return finalPodId
		}
		for j := i + 1; j < len(ListenerPods); j++ {
			if ListenerPods[j].Weight < ListenerPods[i].Weight && ListenerPods[j].NCname == ncname {
				log.Println("-----------balancing---------", i))
				ListenerPods[j].Weight = ListenerPods[j].Weight + 1
				finalPodId = ListenerPods[j].PodId
				return finalPodId
			}
		}
		if ListenerPods[i].NCname == ncname {
			logger.Debug("already balanced :", zap.Int("i: ", i))
			ListenerPods[i].Weight = ListenerPods[i].Weight + 1
			finalPodId = ListenerPods[i].PodId
			return finalPodId
		}

	}
	return finalPodId
}


func caller() {
giveMeListenerId :=  selectListener("demo")
  // TODO: select pg_notify("demo", {}&payload)

}
