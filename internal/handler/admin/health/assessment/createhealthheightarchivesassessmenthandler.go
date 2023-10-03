package assessment

import (
	"net/http"

	"PowerX/internal/logic/admin/health/assessment"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateHealthHeightArchivesAssessmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ActionHealthHeightArchivesAssessmentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := assessment.NewCreateHealthHeightArchivesAssessmentLogic(r.Context(), svcCtx)
		resp, err := l.CreateHealthHeightArchivesAssessment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
