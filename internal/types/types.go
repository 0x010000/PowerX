// Code generated by goctl. DO NOT EDIT.
package types

type DepartmentNode struct {
	Id          int64            `json:"id"`
	DepName     string           `json:"depName"`
	LeaderIds   []int64          `json:"leaderIds"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Remark      string           `json:"remark"`
	Children    []DepartmentNode `json:"children"`
}

type GetDepartmentTreeRequest struct {
	DepId int64 `path:"depId"`
}

type GetDepartmentTreeReply struct {
	DepTree DepartmentNode `json:"depTree"`
}

type CreateDepartmentRequest struct {
	DepName     string  `json:"depName"`
	LeaderIds   []int64 `json:"leaderIds"`
	PId         int64   `json:"pId"`
	Desc        string  `json:"desc,optional"`
	PhoneNumber string  `json:"phoneNumber,optional"`
	Email       string  `json:"email,optional"`
	Remark      string  `json:"remark,optional"`
}

type CreateDepartmentReply struct {
	Id int64 `json:"id"`
}

type DeleteDepartmentRequest struct {
	Id int64 `path:"id"`
}

type DeleteDepartmentReply struct {
	Id int64 `json:"id"`
}

type Department struct {
	Id          int64   `json:"id"`
	DepName     string  `json:"depName"`
	LeaderIds   []int64 `json:"leaderIds"`
	PhoneNumber string  `json:"phoneNumber"`
	Email       string  `json:"email"`
	Remark      string  `json:"remark"`
}

type GetDepartmentRequest struct {
	Id int64 `path:"id"`
}

type GetDepartmentReply struct {
	*Department
}

type GetEmployeeRequest struct {
	Id int64 `path:"id"`
}

type GetEmployeeReply struct {
	*Employee
}

type ListEmployeesRequest struct {
	Ids             []int64  `form:"ids,optional"`
	LikeName        string   `form:"likeName,optional"`
	LikeEmail       string   `form:"likeEmail,optional"`
	DepIds          []int64  `form:"depIds,optional"`
	Positions       []string `form:"positions,optional"`
	LikePhoneNumber string   `form:"likePhoneNumber,optional"`
	RoleCodes       []string `form:"roleCodes,optional"`
	IsEnabled       *bool    `form:"isEnable,optional"`
	PageIndex       int      `form:"pageIndex,optional"`
	PageSize        int      `form:"pageSize,optional"`
}

type Employee struct {
	Id            int64    `json:"id"`
	Account       string   `json:"account"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	MobilePhone   string   `json:"mobilePhone"`
	Gender        int8     `json:"gender"`
	NickName      string   `json:"nickName,optional"`
	Desc          string   `json:"desc,optional"`
	Avatar        string   `json:"avatar,optional"`
	ExternalEmail string   `json:"externalEmail,optional"`
	DepIds        []int64  `json:"depIds"`
	Roles         []string `json:"roles"`
	Position      string   `json:"position"`
	JobTitle      string   `json:"jobTitle"`
	IsEnabled     bool     `json:"isEnabled"`
	CreatedAt     string   `json:"createdAt"`
}

type ListEmployeesReply struct {
	List      []Employee `json:"list"`
	PageIndex int        `json:"pageIndex"`
	PageSize  int        `json:"pageSize"`
	Total     int64      `json:"total"`
}

type SyncEmployeesRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type SyncEmployeesReply struct {
	Status bool `json:"status"`
}

type CreateEmployeeRequest struct {
	Account       string  `json:"account"`
	Name          string  `json:"name"`
	NickName      string  `json:"nickName,optional"`
	Desc          string  `json:"desc,optional"`
	Email         string  `json:"email"`
	Avatar        string  `json:"avatar,optional"`
	ExternalEmail string  `json:"externalEmail,optional"`
	MobilePhone   string  `json:"mobilePhone,optional"`
	Gender        *int8   `json:"gender,optional"`
	DepIds        []int64 `json:"depIds"`
	Position      string  `json:"position,optional"`
	JobTitle      string  `json:"jobTitle,optional"`
	Password      string  `json:"password,optional"`
}

type CreateEmployeeReply struct {
	Id int64 `json:"id"`
}

type UpdateEmployeeRequest struct {
	Id            int64   `path:"id"`
	Name          string  `json:"name,optional"`
	NickName      string  `json:"nickName,optional"`
	Desc          string  `json:"desc,optional"`
	Email         string  `json:"email,optional"`
	Avatar        string  `json:"avatar,optional"`
	ExternalEmail string  `json:"externalEmail,optional"`
	MobilePhone   string  `json:"mobilePhone,optional"`
	Gender        *int8   `json:"gender,optional"`
	DepIds        []int64 `json:"depIds,optional"`
	Position      string  `json:"position,optional"`
	JobTitle      string  `json:"jobTitle,optional"`
	Password      string  `json:"password,optional"`
	Status        *int8   `json:"status,optional"`
}

type UpdateEmployeeReply struct {
	*Employee
}

type GetEmployeeOptionsRequest struct {
	Scopes []string `form:"scopes,options=position|role|department,optional"`
}

type RoleOption struct {
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type DepartmentOption struct {
	DepartmentId   int64  `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
}

type GetEmployeeOptionsReply struct {
	Positions   []string           `json:"positions"`
	Roles       []RoleOption       `json:"roles"`
	Departments []DepartmentOption `json:"departments"`
}

type DeleteEmployeeRequest struct {
	Id int64 `path:"id"`
}

type DeleteEmployeeReply struct {
	Id int64 `json:"id"`
}

type ResetPasswordRequest struct {
	UserId int64 `json:"userId"`
}

type ResetPasswordReply struct {
	Status string `json:"status"`
}

type AuthRole struct {
	RoleCode   string   `json:"roleCode"`
	Name       string   `json:"name"`
	Desc       string   `json:"desc"`
	IsReserved bool     `json:"isReserved"`
	ActIds     []int64  `json:"actIds"`
	MenuNames  []string `json:"menuNames"`
}

type ListRolesReply struct {
	List []AuthRole `json:"list"`
}

type CreateRoleRequest struct {
	RoleCode string `json:"roleCode"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
}

type CreateRoleReply struct {
	RoleCode string `json:"roleCode"`
}

type GetRoleRequest struct {
	RoleCode string `path:"roleCode"`
}

type GetRoleReply struct {
	*AuthRole
}

type PutRoleReqeust struct {
	RoleCode  string   `path:"roleCode"`
	Name      string   `json:"name"`
	Desc      string   `json:"desc"`
	ActIds    []int64  `json:"actIds"`
	MenuNames []string `json:"menuNames"`
}

type PutRoleReply struct {
	*AuthRole
}

type AuthResAct struct {
	Id       int64  `json:"id"`
	Version  string `json:"version"`
	RestPath string `json:"restPath"`
	Action   string `json:"action"`
	Desc     string `json:"desc"`
}

type AuthRes struct {
	Id      int64        `json:"id"`
	ResCode string       `json:"resCode"`
	ResName string       `json:"resName"`
	Type    string       `json:"type"`
	Desc    string       `json:"desc"`
	Acts    []AuthResAct `json:"acts"`
}

type ListRecoursesReply struct {
	List []AuthRes `json:"list"`
}

type GetRoleEmployeeIdsReqeust struct {
	RoleCode string `path:"roleCode"`
}

type GetRoleEmployeeIdsReply struct {
	EmployeeIds []int64 `json:"employeeIds"`
}

type LoginRequest struct {
	UserName    string `json:"userName,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Password    string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ExchangeRequest struct {
	Type string `path:"type,optional=wechat"`
	Code string `json:"code"`
}

type ExchangeReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type Clue struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Source      string `json:"source"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
}

type ListCluesRequest struct {
	LikeTitle       string   `json:"likeTitle"`
	LikePhoneNumber string   `json:"likePhoneNumber"`
	Sources         []string `json:"sources"`
	Statuses        []string `json:"statuses"`
}

type ListCluesReply struct {
	List      []Clue `json:"list"`
	PageIndex int32  `json:"pageIndex"`
	PageSize  int32  `json:"pageSize"`
	Total     int64  `json:"total"`
}

type CreateCluesRequest struct {
	List []Clue `json:"list"`
}

type CreateCluesReply struct {
	List      []Clue `json:"list"`
	PageIndex int32  `json:"pageIndex"`
	PageSize  int32  `json:"pageSize"`
	Total     int64  `json:"total"`
}

type PatchClueRequest struct {
	Id          int64  `path:"id"`
	Title       string `json:"title,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Source      string `json:"source,optional"`
	Status      string `json:"status,optional"`
}

type PatchClueReply struct {
	Clue
}

type DeleteClueReply struct {
	Id int64 `json:"id"`
}