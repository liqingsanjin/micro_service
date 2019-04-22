package staticservice

import (
	"encoding/json"
	"net/http"
	"sync"
	"userService/pkg/common"
	"userService/pkg/model/static"
)

//ConsulIndex .
type ConsulIndex struct {
	Lock sync.Mutex
	Map  map[string]string
}

func (c *ConsulIndex) setIndex(key, ind string) bool {
	c.Lock.Lock()
	defer c.Lock.Unlock()
	for k, v := range c.Map {
		if k == key && ind == v {
			return false
		}
	}
	c.Map[key] = ind
	return true
}

var consulIndex = ConsulIndex{
	Map: make(map[string]string),
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	CIndex := r.Header.Get("X-Consul-Index")
	str := r.Header.Get("x-static")
	result := consulIndex.setIndex(str, CIndex)
	if result {
		str := r.Header.Get("x-static")

		if str == DictionaryConsulKey {
			retPair, _, err := common.ConsulClient.KV().Get(DictionaryConsulKey, nil)
			if err != nil {
				return
			}
			returnDic := make([]*static.DictionaryItem, 0)
			err = json.Unmarshal([]byte(retPair.Value), &returnDic)
			MyMap.dicItem = returnDic
		}

		if str == InsProdBizFeeMapInfConsulKey {
			retPair, _, err := common.ConsulClient.KV().Get(InsProdBizFeeMapInfConsulKey, nil)
			if err != nil {
				return
			}
			ret := make([]*static.InsProdBizFeeMapInf, 0)
			err = json.Unmarshal([]byte(retPair.Value), &ret)
			MyMap.insProdBizFeeMapInf = ret
		}

		if str == ProdBizTransMapConsulKey {
			retPair, _, err := common.ConsulClient.KV().Get(ProdBizTransMapConsulKey, nil)
			if err != nil {
				return
			}
			ret := make([]*static.ProdBizTransMap, 0)
			err = json.Unmarshal([]byte(retPair.Value), &ret)
			MyMap.prodBizTransMap = ret
		}

		if str == InsInfConsulKey {
			retPair, _, err := common.ConsulClient.KV().Get(InsInfConsulKey, nil)
			if err != nil {
				return
			}
			ret := make([]*static.InsInf, 0)
			err = json.Unmarshal([]byte(retPair.Value), &ret)
			MyMap.insInf = ret
		}
	}
}

//StartServer .
func StartServer(addr string, chanErr chan error) {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		chanErr <- err
	}
}
