// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	adminauth "PowerX/internal/handler/admin/auth"
	adminbusinessopportunity "PowerX/internal/handler/admin/business/opportunity"
	admincommon "PowerX/internal/handler/admin/common"
	admincontractway "PowerX/internal/handler/admin/contractway"
	admincustomerdomaincustomer "PowerX/internal/handler/admin/customerdomain/customer"
	admincustomerdomainleader "PowerX/internal/handler/admin/customerdomain/leader"
	admindepartment "PowerX/internal/handler/admin/department"
	admindictionary "PowerX/internal/handler/admin/dictionary"
	adminemployee "PowerX/internal/handler/admin/employee"
	adminmarketmedia "PowerX/internal/handler/admin/market/media"
	adminmedia "PowerX/internal/handler/admin/media"
	adminpermission "PowerX/internal/handler/admin/permission"
	adminproduct "PowerX/internal/handler/admin/product"
	adminproductartisan "PowerX/internal/handler/admin/product/artisan"
	adminproductcategory "PowerX/internal/handler/admin/product/category"
	adminproductpricebook "PowerX/internal/handler/admin/product/pricebook"
	adminproductstore "PowerX/internal/handler/admin/product/store"
	adminscrmcontact "PowerX/internal/handler/admin/scrm/contact"
	adminscrmcustomer "PowerX/internal/handler/admin/scrm/customer"
	admintradeaddressbilling "PowerX/internal/handler/admin/trade/address/billing"
	admintradeaddressdelivery "PowerX/internal/handler/admin/trade/address/delivery"
	admintradeaddressshipping "PowerX/internal/handler/admin/trade/address/shipping"
	admintradeorder "PowerX/internal/handler/admin/trade/order"
	admintradepayment "PowerX/internal/handler/admin/trade/payment"
	admintradewarehouse "PowerX/internal/handler/admin/trade/warehouse"
	adminuserinfo "PowerX/internal/handler/admin/userinfo"
	mpcustomerauth "PowerX/internal/handler/mp/customer/auth"
	mpdictionary "PowerX/internal/handler/mp/dictionary"
	mpproduct "PowerX/internal/handler/mp/product"
	mpproductstore "PowerX/internal/handler/mp/product/store"
	mptradeaddressbilling "PowerX/internal/handler/mp/trade/address/billing"
	mptradeaddressdelivery "PowerX/internal/handler/mp/trade/address/delivery"
	mptradeaddressshipping "PowerX/internal/handler/mp/trade/address/shipping"
	mptradecart "PowerX/internal/handler/mp/trade/cart"
	mptradeorder "PowerX/internal/handler/mp/trade/order"
	"PowerX/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: GetHomeHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeNoPermJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/options/employees",
					Handler: admincommon.GetEmployeeOptionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/options/employee-query",
					Handler: admincommon.GetEmployeeQueryOptionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/options/departments",
					Handler: admincommon.GetDepartmentOptionsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/common"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/department-tree/:depId",
					Handler: admindepartment.GetDepartmentTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/departments/:id",
					Handler: admindepartment.GetDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/departments",
					Handler: admindepartment.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/departments/:id",
					Handler: admindepartment.PatchDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/departments/:id",
					Handler: admindepartment.DeleteDepartmentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/department"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/employees/actions/sync",
					Handler: adminemployee.SyncEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees/:id",
					Handler: adminemployee.GetEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees",
					Handler: adminemployee.ListEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/employees",
					Handler: adminemployee.CreateEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/employees/:id",
					Handler: adminemployee.UpdateEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/employees/:id",
					Handler: adminemployee.DeleteEmployeeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/employees/actions/reset-password",
					Handler: adminemployee.ResetPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/employee"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/roles",
					Handler: adminpermission.ListRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles",
					Handler: adminpermission.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/roles/:roleCode",
					Handler: adminpermission.GetRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/roles/:roleCode",
					Handler: adminpermission.PatchRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/roles/:roleCode/users",
					Handler: adminpermission.GetRoleEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles/:roleCode/actions/set-permissions",
					Handler: adminpermission.SetRolePermissionsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api-list",
					Handler: adminpermission.ListAPIHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/roles/:roleCode/actions/set-employees",
					Handler: adminpermission.SetRoleEmployeesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/:userId/actions/set-roles",
					Handler: adminpermission.SetUserRolesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/permission"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/access/actions/basic-login",
				Handler: adminauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access/actions/exchange-token",
				Handler: adminauth.ExchangeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/admin/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/types/page-list",
					Handler: admindictionary.ListDictionaryPageTypesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/types/:type",
					Handler: admindictionary.GetDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/types",
					Handler: admindictionary.CreateDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/types/:type",
					Handler: admindictionary.UpdateDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/types/:type",
					Handler: admindictionary.DeleteDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items",
					Handler: admindictionary.ListDictionaryItemsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items/:type/:key",
					Handler: admindictionary.GetDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/items",
					Handler: admindictionary.CreateDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/items/:type/:key",
					Handler: admindictionary.UpdateDictionaryItemHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/items/:type/:key",
					Handler: admindictionary.DeleteDictionaryItemHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/dictionary"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user-info",
					Handler: adminuserinfo.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu-roles",
					Handler: adminuserinfo.GetMenuRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/users/actions/modify-password",
					Handler: adminuserinfo.ModifyUserPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/user-center"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/media-resources/page-list",
					Handler: adminmedia.ListMediaResourcesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/media-resources",
					Handler: adminmedia.CreateMediaResourceHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/media-resources/:id",
					Handler: adminmedia.GetMediaResourceHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/media-resources/:id",
					Handler: adminmedia.DeleteMediaResourceHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/media"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/leads/:id",
					Handler: admincustomerdomainleader.GetLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/leads/page-list",
					Handler: admincustomerdomainleader.ListLeadsPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/leads",
					Handler: admincustomerdomainleader.CreateLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/leads/:id",
					Handler: admincustomerdomainleader.PutLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/leads/:id",
					Handler: admincustomerdomainleader.PatchLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/leads/:id",
					Handler: admincustomerdomainleader.DeleteLeadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/leads/:id/actions/employees",
					Handler: admincustomerdomainleader.AssignLeadToEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/customerdomain"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/customers/:id",
					Handler: admincustomerdomaincustomer.GetCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customers/page-list",
					Handler: admincustomerdomaincustomer.ListCustomersPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers",
					Handler: admincustomerdomaincustomer.CreateCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/customers/:id",
					Handler: admincustomerdomaincustomer.PutCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/customers/:id",
					Handler: admincustomerdomaincustomer.PatchCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/customers/:id",
					Handler: admincustomerdomaincustomer.DeleteCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers/:id/actions/employees",
					Handler: admincustomerdomaincustomer.AssignCustomerToEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/customerdomain"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/medias",
					Handler: adminmarketmedia.ListMediasPageHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/medias/actions/create-upload-url",
					Handler: adminmarketmedia.CreateMediaUploadRequestHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/medias/:mediaKey",
					Handler: adminmarketmedia.CreateOrUpdateMediaHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/medias/:key",
					Handler: adminmarketmedia.GetMediaByKeyHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/medias/:key",
					Handler: adminmarketmedia.DeleteMediaHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/market"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/opportunities",
					Handler: adminbusinessopportunity.GetOpportunityListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/opportunities",
					Handler: adminbusinessopportunity.CreateOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/opportunities/:id/assign-employee",
					Handler: adminbusinessopportunity.AssignEmployeeToOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/opportunities/:id",
					Handler: adminbusinessopportunity.UpdateOpportunityHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/opportunities/:id",
					Handler: adminbusinessopportunity.DeleteOpportunityHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/business"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/products/page-list",
					Handler: adminproduct.ListProductsPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/products/:id",
					Handler: adminproduct.GetProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/products",
					Handler: adminproduct.CreateProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/products/:id",
					Handler: adminproduct.PutProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/products/:id",
					Handler: adminproduct.PatchProductHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/products/:id",
					Handler: adminproduct.DeleteProductHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/products/:id/actions/assign-to-product-categroy",
					Handler: adminproduct.AssignProductToProductCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/product-category-tree",
					Handler: adminproductcategory.ListProductCategoryTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/product-categories/:id",
					Handler: adminproductcategory.GetProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/product-categories",
					Handler: adminproductcategory.UpsertProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/product-categories/:id",
					Handler: adminproductcategory.PatchProductCategoryHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/product-categories/:id",
					Handler: adminproductcategory.DeleteProductCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/price-books/page-list",
					Handler: adminproductpricebook.ListPriceBooksHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/price-books/:id",
					Handler: adminproductpricebook.GetPriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/price-books",
					Handler: adminproductpricebook.UpsertPriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/price-books/:id",
					Handler: adminproductpricebook.DeletePriceBookHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/price-book-entries",
					Handler: adminproductpricebook.ConfigPriceBookHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/stores/page-list",
					Handler: adminproductstore.ListStoresPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/stores/:id",
					Handler: adminproductstore.GetStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/stores",
					Handler: adminproductstore.CreateStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/stores/:id",
					Handler: adminproductstore.PutStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/stores/:id",
					Handler: adminproductstore.PatchStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/stores/:id",
					Handler: adminproductstore.DeleteStoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/stores/:id/actions/assign-to-store-categroy",
					Handler: adminproductstore.AssignStoreToStoreManagerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/artisans",
					Handler: adminproductartisan.ListArtisansPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/artisans/:id",
					Handler: adminproductartisan.GetArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/artisans",
					Handler: adminproductartisan.CreateArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/artisans/:id",
					Handler: adminproductartisan.PutArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/artisans/:id",
					Handler: adminproductartisan.PatchArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/artisans/:id",
					Handler: adminproductartisan.DeleteArtisanHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/artisans/:id/actions/assign-to-artisan-categroy",
					Handler: adminproductartisan.AssignArtisanToArtisanCategoryHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/shipping/page-list",
					Handler: admintradeaddressshipping.ListShippingAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/shipping/:id",
					Handler: admintradeaddressshipping.GetShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shipping",
					Handler: admintradeaddressshipping.CreateShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/shipping/:id",
					Handler: admintradeaddressshipping.PutShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/shipping/:id",
					Handler: admintradeaddressshipping.PatchShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/shipping/:id",
					Handler: admintradeaddressshipping.DeleteShippingAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade/address"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/billing/page-list",
					Handler: admintradeaddressbilling.ListBillingAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/billing/:id",
					Handler: admintradeaddressbilling.GetBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/billing",
					Handler: admintradeaddressbilling.CreateBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/billing/:id",
					Handler: admintradeaddressbilling.PutBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/billing/:id",
					Handler: admintradeaddressbilling.PatchBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/billing/:id",
					Handler: admintradeaddressbilling.DeleteBillingAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade/address"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/delivery/page-list",
					Handler: admintradeaddressdelivery.ListDeliveryAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delivery/:id",
					Handler: admintradeaddressdelivery.GetDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delivery",
					Handler: admintradeaddressdelivery.CreateDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/delivery/:id",
					Handler: admintradeaddressdelivery.PutDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/delivery/:id",
					Handler: admintradeaddressdelivery.PatchDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delivery/:id",
					Handler: admintradeaddressdelivery.DeleteDeliveryAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade/address"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/warehouses",
					Handler: admintradewarehouse.ListWarehousesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/warehouses/:id",
					Handler: admintradewarehouse.GetWarehouseHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/warehouses",
					Handler: admintradewarehouse.CreateWarehouseHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/warehouses/:id",
					Handler: admintradewarehouse.UpdateWarehouseHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/warehouses/:id",
					Handler: admintradewarehouse.PatchWarehouseHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/warehouses/:id",
					Handler: admintradewarehouse.DeleteWarehouseHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/group-tree",
					Handler: admincontractway.GetContractWayGroupTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/groups",
					Handler: admincontractway.GetContractWayGroupListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: admincontractway.GetContractWaysHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: admincontractway.CreateContractWayHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/:id",
					Handler: admincontractway.UpdateContractWayHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: admincontractway.DeleteContractWayHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/contract-way"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/customers/:id",
					Handler: adminscrmcustomer.GetWeWorkCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/customers",
					Handler: adminscrmcustomer.ListWeWorkCustomersHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/customers/:id",
					Handler: adminscrmcustomer.PatchWeWorkCustomerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/customers/actions/sync",
					Handler: adminscrmcustomer.SyncWeWorkCustomerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/scrm/customer"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/contacts/actions/sync",
					Handler: adminscrmcontact.SyncWeWorkContactHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/employees",
					Handler: adminscrmcontact.ListWeWorkEmployeeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/scrm/contact"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/types/page-list",
					Handler: mpdictionary.ListDictionaryPageTypesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/types/:type",
					Handler: mpdictionary.GetDictionaryTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items",
					Handler: mpdictionary.ListDictionaryItemsHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/items/:type/:key",
					Handler: mpdictionary.GetDictionaryItemHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/dictionary"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: mpcustomerauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/authByPhone",
				Handler: mpcustomerauth.AuthByPhoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/authByProfile",
				Handler: mpcustomerauth.AuthByProfileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/mp/customer"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/stores/page-list",
					Handler: mpproductstore.ListStoresPageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/products/page-list",
					Handler: mpproduct.ListProductsPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/products/:id",
					Handler: mpproduct.GetProductHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/product-category-tree",
					Handler: mpproduct.ListProductCategoryTreeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/product"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth, serverCtx.MPCustomerGet},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/cart/items/page-list",
					Handler: mptradecart.ListCartItemsPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/cart/:cartId",
					Handler: mptradecart.GetCartHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cart/items",
					Handler: mptradecart.AddToCartHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/cart/items/:itemId",
					Handler: mptradecart.UpdateCartItemQuantityHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/cart/items/:itemId",
					Handler: mptradecart.RemoveCartItemHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/cart/items/clear",
					Handler: mptradecart.ClearCartItemsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/trade"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth, serverCtx.MPCustomerGet},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/orders/page-list",
					Handler: mptradeorder.ListOrdersPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/orders/:id",
					Handler: mptradeorder.GetOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/orders",
					Handler: mptradeorder.CreateOrderByCartItemsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/orders/cancel/:id",
					Handler: mptradeorder.CancelOrderHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/trade"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/orders/page-list",
					Handler: admintradeorder.ListOrdersPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/orders/:id",
					Handler: admintradeorder.GetOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/orders",
					Handler: admintradeorder.CreateOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/orders/:id",
					Handler: admintradeorder.PutOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/orders/:id",
					Handler: admintradeorder.PatchOrderHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/orders/:id",
					Handler: admintradeorder.DeleteOrderHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.EmployeeJWTAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/payment/page-list",
					Handler: admintradepayment.ListPaymentsPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/payment/:id",
					Handler: admintradepayment.GetPaymentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/payment",
					Handler: admintradepayment.CreatePaymentHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/payment/:id",
					Handler: admintradepayment.PutPaymentHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/payment/:id",
					Handler: admintradepayment.PatchPaymentHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/payment/:id",
					Handler: admintradepayment.DeletePaymentHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/trade"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth, serverCtx.MPCustomerGet},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/shipping/page-list",
					Handler: mptradeaddressshipping.ListShippingAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/shipping/:id",
					Handler: mptradeaddressshipping.GetShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shipping",
					Handler: mptradeaddressshipping.CreateShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/shipping/:id",
					Handler: mptradeaddressshipping.PutShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/shipping/:id",
					Handler: mptradeaddressshipping.PatchShippingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/shipping/:id",
					Handler: mptradeaddressshipping.DeleteShippingAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/trade/address"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth, serverCtx.MPCustomerGet},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/delivery/page-list",
					Handler: mptradeaddressdelivery.ListDeliveryAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delivery/:id",
					Handler: mptradeaddressdelivery.GetDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delivery",
					Handler: mptradeaddressdelivery.CreateDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/delivery/:id",
					Handler: mptradeaddressdelivery.PutDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/delivery/:id",
					Handler: mptradeaddressdelivery.PatchDeliveryAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delivery/:id",
					Handler: mptradeaddressdelivery.DeleteDeliveryAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/trade/address"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.MPCustomerJWTAuth, serverCtx.MPCustomerGet},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/billing/page-list",
					Handler: mptradeaddressbilling.ListBillingAddressesPageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/billing/:id",
					Handler: mptradeaddressbilling.GetBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/billing",
					Handler: mptradeaddressbilling.CreateBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/billing/:id",
					Handler: mptradeaddressbilling.PutBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodPatch,
					Path:    "/billing/:id",
					Handler: mptradeaddressbilling.PatchBillingAddressHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/billing/:id",
					Handler: mptradeaddressbilling.DeleteBillingAddressHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/mp/trade/address"),
	)
}
