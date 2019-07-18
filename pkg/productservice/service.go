package productservice

import (
	"context"
	"userService/pkg/common"
	productmodel "userService/pkg/model/product"
	"userService/pkg/pb"
)

type service struct{}

func (s *service) ListTransMap(ctx context.Context, in *pb.ListTransMapRequest) (*pb.ListTransMapReply, error) {
	reply := new(pb.ListTransMapReply)
	query := new(productmodel.Trans)
	if in.Item != nil {
		query.ProdCd = in.Item.ProdCd
		query.BizCd = in.Item.BizCd
		query.TransCd = in.Item.TransCd
		query.UpdateDate = in.Item.UpdateDate
		query.Description = in.Item.Description
		query.ResvFld1 = in.Item.ResvFld1
		query.ResvFld2 = in.Item.ResvFld2
		query.ResvFld3 = in.Item.ResvFld3
	}
	db := common.DB

	items, err := productmodel.ListTrans(db, query)
	if err != nil {
		return nil, err
	}

	pbItems := make([]*pb.ProductBizTransMapField, len(items))

	for i := range items {
		pbItems[i] = &pb.ProductBizTransMapField{
			ProdCd:      items[i].ProdCd,
			BizCd:       items[i].BizCd,
			TransCd:     items[i].TransCd,
			UpdateDate:  items[i].UpdateDate,
			Description: items[i].Description,
			ResvFld1:    items[i].ResvFld1,
			ResvFld2:    items[i].ResvFld2,
			ResvFld3:    items[i].ResvFld3,
		}
	}

	reply.Items = pbItems
	return reply, nil
}
