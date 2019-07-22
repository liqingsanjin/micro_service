package staticservice

import (
	"context"
	"encoding/json"
	"net/http"
	"userService/pkg/common"
	"userService/pkg/model/static"
	"userService/pkg/pb"
	"userService/pkg/util"

	"github.com/hashicorp/consul/api"
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

//service .
type service struct {
}

func (s *service) GetUnionPayBankListByCode(ctx context.Context, in *pb.GetUnionPayBankListByCodeRequest) (*pb.GetUnionPayBankListByCodeReply, error) {
	reply := new(pb.GetUnionPayBankListByCodeReply)
	if in.Item == nil || in.Item.Code == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "code不能为空",
		}
		return reply, nil
	}
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 10
	}
	db := common.DB

	items, count, err := static.FindBankListByCode(db, in.Item.Code, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	pbItems := make([]*pb.UnionPayBankListField, len(items))
	for i := range items {
		pbItems[i] = &pb.UnionPayBankListField{
			Id:   items[i].Id,
			Code: items[i].Code,
			Name: items[i].Name,
		}
		if !items[i].UpdatedAt.IsZero() {
			pbItems[i].UpdatedAt = items[i].UpdatedAt.Format(util.TimePattern)
		}
	}
	reply.Size = in.Size
	reply.Page = in.Page
	reply.Items = pbItems
	reply.Count = count
	return reply, nil

}

//Download .
func (s *service) SyncData(ctx context.Context, in *pb.StaticSyncDataReq) (*pb.StaticSyncDataResp, error) {
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

func (s *service) GetDictionaryItem(ctx context.Context, in *pb.StaticGetDictionaryItemReq) (*pb.StaticGetDictionaryItemResp, error) {

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
	return &pb.StaticGetDictionaryItemResp{Items: items}, nil

}

func (s *service) GetDicByProdAndBiz(ctx context.Context, in *pb.StaticGetDicByProdAndBizReq) (*pb.StaticGetDicByProdAndBizResp, error) {
	if in.ProdCd == "" {
		return &pb.StaticGetDicByProdAndBizResp{Err: &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "必须传入ProdCd",
		}}, nil
	}
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
		return &pb.StaticGetDicByProdAndBizResp{Items: items}, nil
	}

	transCds := getTransCdByProdAndBiz(in.ProdCd, in.BizCd)
	if len(transCds) == 0 {
		return &pb.StaticGetDicByProdAndBizResp{}, nil
	}

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
	return &pb.StaticGetDicByProdAndBizResp{Items: items}, nil

}

func (s *service) GetDicByInsCmpCd(ctx context.Context, in *pb.StaticGetDicByInsCmpCdReq) (*pb.StaticGetDicByInsCmpCdResp, error) {
	if in.InsCompanyCd == "" {
		return &pb.StaticGetDicByInsCmpCdResp{Err: &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "必须传入InsCompanyCd",
		}}, nil
	}
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
	return &pb.StaticGetDicByInsCmpCdResp{Items: items}, nil
}

func (s *service) CheckValues(ctx context.Context, in *pb.StaticCheckValuesReq) (*pb.StaticCheckValuesResp, error) {
	if in.ProdCd == "" && in.BizCd == "" && in.TransCd == "" && in.InsCompanyCd == "" && in.FwdInsIdCd == "" {
		return &pb.StaticCheckValuesResp{Err: &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "必须传入参数",
		}}, nil
	}
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
		if includes(dicTypeArr, MyMap.dicItem[i].DicType) &&
			includes(dicNameArr, MyMap.dicItem[i].DicName) &&
			includes(dicCodeArr, MyMap.dicItem[i].DicCode) {
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

func (s *service) GetDictionaryLayerItem(ctx context.Context, in *pb.GetDictionaryLayerItemReq) (*pb.GetDictionaryLayerItemResp, error) {
	out := make([]*pb.DictionaryLayerItem, 0)
	db := common.DB
	items := static.GetDictionaryLayerItem(db, &static.DictionaryLayerItem{
		DicType:  in.DicType,
		DicPCode: in.DicPCode,
		DicCode:  in.DicCode,
		DicLevel: in.DicLevel,
	})
	for i := range items {
		out = append(out, &pb.DictionaryLayerItem{
			DicType:   items[i].DicType,
			DicCode:   items[i].DicCode,
			DicPCode:  items[i].DicPCode,
			DicLevel:  items[i].DicLevel,
			DisPOrder: items[i].DisPOrder,
			Name:      items[i].Name,
			Memo:      items[i].Memo,
		})
	}
	return &pb.GetDictionaryLayerItemResp{
		Items: out,
	}, nil
}

func (s *service) GetDictionaryItemByPk(ctx context.Context, in *pb.GetDictionaryItemByPkReq) (*pb.GetDictionaryItemByPkResp, error) {
	reply := new(pb.GetDictionaryItemByPkResp)
	if in.DicCode == "" || in.DicType == "" {
		reply.Err = &pb.Error{
			Code:        http.StatusBadRequest,
			Message:     InvalidParam,
			Description: "dicCode 和 dicType不能为空",
		}
		return reply, nil
	}
	db := common.DB

	dic := static.GetDictionaryItemByPk(db, &static.DictionaryItem{
		DicType: in.DicType,
		DicCode: in.DicCode,
	})
	if dic != nil {
		reply.Item = &pb.StaticGetDictionaryItem{
			DicType: dic.DicType,
			DicCode: dic.DicCode,
			DicName: dic.DicName,
		}
	}

	return reply, nil
}
