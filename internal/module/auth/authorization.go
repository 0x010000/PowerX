package auth

import (
	"PowerX/internal/model/permission"
	"context"
)

//
// Authorization
//  @Description:
//  @param ctx
//  @return author
//
func Authorization(ctx context.Context) (author *permission.AdminAuthMetadata) {

	if ctx != nil {
		author = ctx.Value(permission.AdminAuthMetadataKey{}).(*permission.AdminAuthMetadata)
	}
	return author
}
