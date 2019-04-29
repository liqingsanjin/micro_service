package staticservice

import (
	"encoding/json"
	"userService/pkg/common"
	"userService/pkg/model/static"
	"userService/pkg/pb"

	"github.com/hashicorp/consul/api"

	"golang.org/x/net/context"
)

var (
	DictionaryConsulKey          = "static/dictionaryItem"
	InsProdBizFeeMapInfConsulKey = "static/insProdBizFeeMapInf"
	ProdBizTransMapConsulKey     = "static/prodBizTransMap"
	InsInfConsulKey              = "static/insInf"
)

type StaticMapData struct {
	dicItem             []*static.DictionaryItem
	insProdBizFeeMapInf []*static.InsProdBizFeeMapInf
	prodBizTransMap     []*static.ProdBizTransMap
	insInf              []*static.InsInf
}

var MyMap = StaticMapData{
	dicItem:             make([]*static.DictionaryItem, 0),
	insProdBizFeeMapInf: make([]*static.InsProdBizFeeMapInf, 0),
	prodBizTransMap:     make([]*static.ProdBizTransMap, 0),
	insInf:              make([]*static.InsInf, 0),
}

//setService .
type setService struct {
}

//NewSetService return institution service with grpc registry type.
func NewSetService() pb.StaticServer {
	return &setService{}
}

//Download .
func (s *setService) SyncData(ctx context.Context, in *pb.StaticSyncDataReq) (*pb.StaticSyncDataResp, error) {
	dicItems := static.GetDictionaryItem(common.DB)
	insProdBizFeeMapInfs := static.GetInsProdBizFeeMapInf(common.DB)
	prodBizTransMaps := static.GetProdBizTransMap(common.DB)
	insInfs := static.GetInsInf(common.DB)

	val1, err := json.Marshal(dicItems)
	if err != nil {
		return nil, err
	}
	kv1 := &api.KVPair{
		Key:   DictionaryConsulKey,
		Flags: 0,
		Value: val1,
	}

	val2, err := json.Marshal(insProdBizFeeMapInfs)
	if err != nil {
		return nil, err
	}
	kv2 := &api.KVPair{
		Key:   InsProdBizFeeMapInfConsulKey,
		Flags: 0,
		Value: val2,
	}

	val3, err := json.Marshal(prodBizTransMaps)
	if err != nil {
		return nil, err
	}
	kv3 := &api.KVPair{
		Key:   ProdBizTransMapConsulKey,
		Flags: 0,
		Value: val3,
	}

	val4, err := json.Marshal(insInfs)
	if err != nil {
		return nil, err
	}
	kv4 := &api.KVPair{
		Key:   InsInfConsulKey,
		Flags: 0,
		Value: val4,
	}

	common.ConsulClient.KV().Put(kv1, nil)
	common.ConsulClient.KV().Put(kv2, nil)
	common.ConsulClient.KV().Put(kv3, nil)
	common.ConsulClient.KV().Put(kv4, nil)

	return &pb.StaticSyncDataResp{Result: true}, nil
}

func (s *setService) GetDictionaryItem(ctx context.Context, in *pb.StaticGetDictionaryItemReq) (*pb.StaticGetDictionaryItemResp, error) {

	dicTypeCondition := make([]string, 0)
	dicNameCondition := make([]string, 0)
	dicCodeCondition := make([]string, 0)

	if in.DicType != "" {
		dicTypeCondition = append(dicTypeCondition, in.DicType)
	}
	if in.DicName != "" {
		dicNameCondition = append(dicNameCondition, in.DicName)
	}
	if in.DicCode != "" {
		dicCodeCondition = append(dicCodeCondition, in.DicCode)
	}

	items := make([]*pb.StaticGetDictionaryItem, 0)
	results := getDicItemByCondition(dicTypeCondition, dicNameCondition, dicCodeCondition)
	for i := 0; i < len(results); i++ {
		items = append(items, &pb.StaticGetDictionaryItem{
			DicCode: results[i].DicCode,
			DicName: results[i].DicName,
			DicType: results[i].DicType,
		})
	}
	return &pb.StaticGetDictionaryItemResp{GetDictionaryItems: items}, nil

}

func (s *setService) GetDicByProdAndBiz(ctx context.Context, in *pb.StaticGetDicByProdAndBizReq) (*pb.StaticGetDicByProdAndBizResp, error) {
	bizCdArr := make([]string, 0)
	prodCdArr := make([]string, 0)
	dicTypeCondition := make([]string, 0)
	dicNameCondition := make([]string, 0)

	if in.BizCd == "" {
		prodCdArr = append(prodCdArr, in.ProdCd)
		bizCdArr = getBizCdByProdCd(prodCdArr)
		if len(bizCdArr) == 0 {
			return &pb.StaticGetDicByProdAndBizResp{}, nil
		}

		dicTypeCondition = append(dicTypeCondition, "BIZ_CD")
		results := getDicItemByCondition(dicTypeCondition, dicNameCondition, bizCdArr)
		if len(results) == 0 {
			return &pb.StaticGetDicByProdAndBizResp{}, nil
		}

		items := make([]*pb.StaticGetDictionaryItem, 0)
		for i := 0; i < len(results); i++ {
			items = append(items, &pb.StaticGetDictionaryItem{
				DicCode: results[i].DicCode,
				DicName: results[i].DicName,
				DicType: results[i].DicType,
			})
		}
		return &pb.StaticGetDicByProdAndBizResp{GetDictionaryItems: items}, nil
	}

	transCds := getTransCdByProdAndBiz(in.ProdCd, in.BizCd)
	dicTypeCondition = append(dicTypeCondition, "TRANS_CD")
	results := getDicItemByCondition(dicTypeCondition, dicNameCondition, transCds)

	items := make([]*pb.StaticGetDictionaryItem, 0)
	for i := 0; i < len(results); i++ {
		items = append(items, &pb.StaticGetDictionaryItem{
			DicCode: results[i].DicCode,
			DicName: results[i].DicName,
			DicType: results[i].DicType,
		})
	}
	return &pb.StaticGetDicByProdAndBizResp{GetDictionaryItems: items}, nil

}

func (s *setService) GetDicByInsCmpCd(ctx context.Context, in *pb.StaticGetDicByInsCmpCdReq) (*pb.StaticGetDicByInsCmpCdResp, error) {
	dicTypeCondition := []string{"INS_ID_CD"}
	dicNameCondition := make([]string, 0)
	insCds := getDicByInsCmpCd(in.InsCompanyCd)
	if len(insCds) == 0 {
		return &pb.StaticGetDicByInsCmpCdResp{}, nil
	}

	results := getDicItemByCondition(dicTypeCondition, dicNameCondition, insCds)
	items := make([]*pb.StaticGetDictionaryItem, 0)
	for i := 0; i < len(results); i++ {
		items = append(items, &pb.StaticGetDictionaryItem{
			DicCode: results[i].DicCode,
			DicName: results[i].DicName,
			DicType: results[i].DicType,
		})
	}
	return &pb.StaticGetDicByInsCmpCdResp{GetDictionaryItems: items}, nil
}

func (s *setService) CheckValues(ctx context.Context, in *pb.StaticCheckValuesReq) (*pb.StaticCheckValuesResp, error) {
	result := checkProdBizAndTrans(in.ProdCd, in.BizCd, in.TransCd) && checkCpyCdAndFwd(in.InsCompanyCd, in.FwdInsIdCd)

	return &pb.StaticCheckValuesResp{Result: result}, nil
}

func checkCpyCdAndFwd(insCompanyCd, fwdInsIdCd string) bool {
	if insCompanyCd == "" && fwdInsIdCd == "" {
		return true
	}

	dicCode := make([]string, 0)
	if insCompanyCd != "" {
		dicCode = append(dicCode, insCompanyCd)
	}
	if fwdInsIdCd != "" {
		dicCode = append(dicCode, fwdInsIdCd)
	}

	results := getDicItemByCondition([]string{}, []string{}, dicCode)
	if len(results) != len(dicCode) {
		return false
	}
	if insCompanyCd != "" && fwdInsIdCd != "" {
		cpyCdArr := getDicByInsCmpCd(insCompanyCd)
		return inArr(cpyCdArr, fwdInsIdCd)
	}
	return true
}

func checkProdBizAndTrans(prodCd, bizCd, transCd string) bool {
	if prodCd == "" && bizCd == "" && transCd == "" {
		return true
	}

	dicCode := make([]string, 0)
	if prodCd != "" {
		dicCode = append(dicCode, prodCd)
	}
	if bizCd != "" {
		dicCode = append(dicCode, bizCd)
	}
	if transCd != "" {
		dicCode = append(dicCode, transCd)
	}

	results := getDicItemByCondition([]string{}, []string{}, dicCode)
	if len(results) != len(dicCode) {
		return false
	}

	if bizCd == "" || prodCd == "" {
		return true
	}

	if transCd == "" {
		bizArr := getBizCdByProdCd([]string{prodCd})
		return inArr(bizArr, bizCd)
	}

	tranArr := getTransCdByProdAndBiz(prodCd, bizCd)
	return inArr(tranArr, transCd)
}

func inArr(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func getDicItemByCondition(dicTypeArr, dicNameArr, dicCodeArr []string) []*static.DictionaryItem {
	returnDic := make([]*static.DictionaryItem, 0)
	for i := 0; i < len(MyMap.dicItem); i++ {
		if includes(dicTypeArr, (*MyMap.dicItem[i]).DicType) &&
			includes(dicNameArr, (*MyMap.dicItem[i]).DicName) &&
			includes(dicCodeArr, (*MyMap.dicItem[i]).DicCode) {
			returnDic = append(returnDic, MyMap.dicItem[i])
		}
	}
	return returnDic
}

func getBizCdByProdCd(prodCd []string) []string {
	bizCdArr := make([]string, 0)
	for i := 0; i < len(MyMap.insProdBizFeeMapInf); i++ {
		if includes(prodCd, (*MyMap.insProdBizFeeMapInf[i]).ProdCd) {
			bizCdArr = append(bizCdArr, (*MyMap.insProdBizFeeMapInf[i]).BizCd)
		}
	}
	return bizCdArr
}

func getTransCdByProdAndBiz(prodCd, bizCd string) []string {
	transCdArr := make([]string, 0)
	for i := 0; i < len(MyMap.prodBizTransMap); i++ {
		if (*MyMap.prodBizTransMap[i]).ProdCd == prodCd &&
			(*MyMap.prodBizTransMap[i]).BizCd == bizCd {
			transCdArr = append(transCdArr, (*MyMap.prodBizTransMap[i]).TransCd)
		}
	}
	return transCdArr
}

func getDicByInsCmpCd(insCmpCd string) []string {
	insIds := make([]string, 0)
	for i := 0; i < len(MyMap.insInf); i++ {
		if (*MyMap.insInf[i]).InsCompanyCd == insCmpCd {
			insIds = append(insIds, (*MyMap.insInf[i]).InsIDCd)
		}
	}
	return insIds
}

func includes(arr []string, val string) bool {
	if len(arr) == 0 {
		return true
	}
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
