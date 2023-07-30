package product

import (
	"PowerX/internal/logic/admin/product/category"
	"PowerX/internal/logic/admin/product/pricebookentry"
	"PowerX/internal/model"
	"PowerX/internal/model/media"
	"PowerX/internal/model/product"
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/uc/powerx"
	product2 "PowerX/internal/uc/powerx/product"
	"context"
	"encoding/json"
	"github.com/golang-module/carbon/v2"
	"gorm.io/datatypes"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductRequest) (resp *types.CreateProductReply, err error) {

	mdlProduct := TransformProductRequestToProduct(&(req.Product))

	if len(req.SalesChannelsItemIds) > 0 {
		salesChannelsItems, err := l.svcCtx.PowerX.DataDictionary.FindAllDictionaryItems(l.ctx, &powerx.FindManyDataDictItemOption{
			Ids: req.SalesChannelsItemIds,
		})
		if err != nil {
			return nil, err
		}
		mdlProduct.PivotSalesChannels, err = (&model.PivotDataDictionaryToObject{}).MakeMorphPivotsFromObjectToDDs(mdlProduct, salesChannelsItems)
	}

	if len(req.PromoteChannelsItemIds) > 0 {
		promoteChannelsItems, err := l.svcCtx.PowerX.DataDictionary.FindAllDictionaryItems(l.ctx, &powerx.FindManyDataDictItemOption{
			Ids: req.PromoteChannelsItemIds,
		})
		if err != nil {
			return nil, err
		}
		mdlProduct.PivotPromoteChannels, err = (&model.PivotDataDictionaryToObject{}).MakeMorphPivotsFromObjectToDDs(mdlProduct, promoteChannelsItems)
	}

	if len(req.CategoryIds) > 0 {
		productCategories := l.svcCtx.PowerX.ProductCategory.FindAllProductCategories(l.ctx, &product2.FindProductCategoryOption{
			Ids: req.CategoryIds,
		})
		mdlProduct.ProductCategories = productCategories
	}

	if len(req.CoverImageIds) > 0 {
		mediaResources, err := l.svcCtx.PowerX.MediaResource.FindAllMediaResources(l.ctx, &powerx.FindManyMediaResourcesOption{
			Ids: req.CoverImageIds,
		})
		if err != nil {
			return nil, err
		}
		mdlProduct.PivotCoverImages, err = (&media.PivotMediaResourceToObject{}).MakeMorphPivotsFromObjectToMediaResources(mdlProduct, mediaResources, media.MediaUsageCover)
	}

	if len(req.DetailImageIds) > 0 {
		mediaResources, err := l.svcCtx.PowerX.MediaResource.FindAllMediaResources(l.ctx, &powerx.FindManyMediaResourcesOption{
			Ids: req.DetailImageIds,
		})
		if err != nil {
			return nil, err
		}
		mdlProduct.PivotDetailImages, err = (&media.PivotMediaResourceToObject{}).MakeMorphPivotsFromObjectToMediaResources(mdlProduct, mediaResources, media.MediaUsageDetail)
	}

	err = l.svcCtx.PowerX.Product.CreateProduct(l.ctx, mdlProduct)

	return &types.CreateProductReply{
		mdlProduct.Id,
	}, err
}

func TransformProductRequestToProduct(productRequest *types.Product) (mdlProduct *product.Product) {

	saleStartDate := carbon.Parse(productRequest.SaleStartDate)
	saleEndDate := carbon.Parse(productRequest.SaleEndDate)
	mdlProduct = &product.Product{
		Name:                productRequest.Name,
		SPU:                 productRequest.SPU,
		Type:                productRequest.Type,
		Plan:                productRequest.Plan,
		AccountingCategory:  productRequest.AccountingCategory,
		CanSellOnline:       productRequest.CanSellOnline,
		CanUseForDeduct:     productRequest.CanUseForDeduct,
		ApprovalStatus:      productRequest.ApprovalStatus,
		IsActivated:         productRequest.IsActivated,
		Description:         productRequest.Description,
		AllowedSellQuantity: productRequest.AllowedSellQuantity,
		ValidityPeriodDays:  productRequest.ValidityPeriodDays,
		SaleStartDate:       saleStartDate.ToStdTime(),
		SaleEndDate:         saleEndDate.ToStdTime(),
		ProductAttribute: product.ProductAttribute{
			Inventory: productRequest.Inventory,
			Weight:    productRequest.Weight,
			Volume:    productRequest.Volume,
			Encode:    productRequest.Encode,
			BarCode:   productRequest.BarCode,
			Extra:     datatypes.JSON(productRequest.Extra),
		},
	}

	return mdlProduct

}

func TransformProductToProductReply(mdlProduct *product.Product) (productReply *types.Product) {

	return &types.Product{
		Id:                  mdlProduct.Id,
		Name:                mdlProduct.Name,
		SPU:                 mdlProduct.SPU,
		Type:                mdlProduct.Type,
		Plan:                mdlProduct.Plan,
		AccountingCategory:  mdlProduct.AccountingCategory,
		CanSellOnline:       mdlProduct.CanSellOnline,
		CanUseForDeduct:     mdlProduct.CanUseForDeduct,
		ApprovalStatus:      mdlProduct.ApprovalStatus,
		IsActivated:         mdlProduct.IsActivated,
		Description:         mdlProduct.Description,
		AllowedSellQuantity: mdlProduct.AllowedSellQuantity,
		ValidityPeriodDays:  mdlProduct.ValidityPeriodDays,
		SaleStartDate:       mdlProduct.SaleStartDate.String(),
		SaleEndDate:         mdlProduct.SaleEndDate.String(),
		//PivotSalesChannels:   TransformDDsToDDsReply(mdlProduct.PivotSalesChannels),
		//PivotPromoteChannels: TransformDDsToDDsReply(mdlProduct.PivotPromoteChannels),
		ProductCategories:      category.TransformProductCategoriesToProductCategoriesReply(mdlProduct.ProductCategories),
		SalesChannelsItemIds:   model.GetItemIds(mdlProduct.PivotSalesChannels),
		PromoteChannelsItemIds: model.GetItemIds(mdlProduct.PivotPromoteChannels),
		CategoryIds:            product.GetCategoryIds(mdlProduct.ProductCategories),
		ProductSpecifics:       TransformSpecificsToSpecificsReply(mdlProduct.ProductSpecifics),
		ActivePriceEntry:       pricebookentry.TransformPriceEntriesToActivePriceEntryReply(mdlProduct.PriceBookEntries),
		PriceBookEntries:       pricebookentry.TransformPriceBookEntriesToPriceBookEntriesReply(mdlProduct.PriceBookEntries),
		SKUs:                   TransformSkusToSkusReply(mdlProduct.SKUs),
		CoverImageIds:          media.GetImageIds(mdlProduct.PivotCoverImages),
		DetailImageIds:         media.GetImageIds(mdlProduct.PivotDetailImages),
		CoverImages:            TransformProductImagesToImagesReply(mdlProduct.PivotCoverImages),
		DetailImages:           TransformProductImagesToImagesReply(mdlProduct.PivotDetailImages),
	}

}

func TransformProductImagesToImagesReply(pivots []*media.PivotMediaResourceToObject) (imagesReply []*types.ProductImage) {

	imagesReply = []*types.ProductImage{}
	for _, pivot := range pivots {
		imageReply := TransformProductImageToImageReply(pivot.MediaResource)
		imagesReply = append(imagesReply, imageReply)
	}
	return imagesReply
}

func TransformProductImageToImageReply(resource *media.MediaResource) (imagesReply *types.ProductImage) {
	if resource == nil {
		return nil
	}
	return &types.ProductImage{
		Id:            resource.Id,
		Url:           resource.Url,
		IsLocalStored: resource.IsLocalStored,
		Filename:      resource.Filename,
	}
}

func TransformSkusToSkusReply(skus []*product.SKU) (skusReply []*types.SKU) {

	skusReply = []*types.SKU{}
	for _, sku := range skus {
		skuReply := TransformSkuToSkuReply(sku)
		skusReply = append(skusReply, skuReply)
	}
	return skusReply
}

func TransformSkuToSkuReply(sku *product.SKU) (skuReply *types.SKU) {
	if sku == nil {
		return nil
	}

	unitPrice := 0.0
	listPrice := 0.0
	isActive := false
	if sku.PriceBookEntry != nil {
		unitPrice = sku.PriceBookEntry.UnitPrice
		listPrice = sku.PriceBookEntry.ListPrice
		isActive = sku.PriceBookEntry.IsActive
	}

	optionsIds := []int64{}
	_ = json.Unmarshal(sku.OptionIds, &optionsIds)

	return &types.SKU{
		Id:         sku.Id,
		UniqueId:   sku.UniqueID.String,
		SkuNo:      sku.SkuNo,
		ProductId:  sku.ProductId,
		Inventory:  sku.Inventory,
		UnitPrice:  unitPrice,
		ListPrice:  listPrice,
		IsActive:   isActive,
		OptionsIds: optionsIds,
	}
}

func TransformSpecificsToSpecificsReply(specifics []*product.ProductSpecific) (specificReplies []*types.ProductSpecific) {

	specificReplies = []*types.ProductSpecific{}
	for _, specific := range specifics {
		specificReply := TransformSpecificToSpecificReply(specific)
		specificReplies = append(specificReplies, specificReply)
	}
	return specificReplies
}

func TransformSpecificToSpecificReply(specific *product.ProductSpecific) (imagesReply *types.ProductSpecific) {
	if specific == nil {
		return nil
	}
	return &types.ProductSpecific{
		Id:              specific.Id,
		ProductId:       specific.ProductId,
		Name:            specific.Name,
		SpecificOptions: TransformSpecificOptionsToSpecificOptionsReply(specific.Options),
	}
}

func TransformSpecificOptionsToSpecificOptionsReply(options []*product.SpecificOption) (optionsReply []*types.SpecificOption) {

	optionsReply = []*types.SpecificOption{}
	for _, option := range options {
		specificReply := TransformSpecificOptionToSpecificOptionReply(option)
		optionsReply = append(optionsReply, specificReply)
	}
	return optionsReply
}

func TransformSpecificOptionToSpecificOptionReply(option *product.SpecificOption) (imagesReply *types.SpecificOption) {
	if option == nil {
		return nil
	}
	return &types.SpecificOption{
		Id:                option.Id,
		ProductSpecificId: option.ProductSpecificId,
		Name:              option.Name,
		IsActivated:       option.IsActivated,
	}
}
