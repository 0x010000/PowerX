package archives

import (
	"net/http"

	"PowerX/internal/logic/admin/health/archives"
	"PowerX/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHealthHeightArchivesOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := archives.NewListHealthHeightArchivesOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListHealthHeightArchivesOption()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
