package assessment

import (
	"net/http"

	"PowerX/internal/logic/admin/health/assessment"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHealthHeightArchivesAssessmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHealthHeightArchivesAssessmentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := assessment.NewListHealthHeightArchivesAssessmentLogic(r.Context(), svcCtx)
		resp, err := l.ListHealthHeightArchivesAssessment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
