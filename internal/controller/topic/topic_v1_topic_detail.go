package topic

import (
	"context"

	"focus-single/api/topic/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
)

func (c *ControllerV1) TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error) {
	out, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		service.View().Render404(ctx)
		return
	}
	err = service.Content().AddViewCount(ctx, req.Id, 1)
	service.View().Render(ctx, model.View{
		ContentType: consts.ContentTypeTopic,
		Data:        out,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
			CurrentName: out.Content.Title,
		}),
		BreadCrumb: service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   out.Content.Id,
			ContentType: out.Content.Type,
			CategoryId:  out.Content.CategoryId,
		}),
	})
	return
}
