// 系统常量
// author xiaoRui

package constant

const (
	ContextKeyUserObj      = "authedUserObj"
	LOGIN_CODE             = "login_code:"
	INVALID_PARAMS         = 400
	GROUP_EXIST            = 415
	ECS_AUTH_CREATE_FAILED = 416
	ECS_AUTH_UPDATE_FAILED = 417
	ECS_AUTH_DELETE_FAILED = 418
	ECS_AUTH_NAME_EXISTS   = 419
	ECS_AUTH_NOT_FOUND     = 420

	// CMDB主机相关错误码
	CMDB_HOST_CREATE_FAILED = 421
	CMDB_HOST_UPDATE_FAILED = 422
	CMDB_HOST_DELETE_FAILED = 423
	CMDB_HOST_NAME_EXISTS   = 424
	CMDB_HOST_NOT_FOUND     = 425
	CMDB_HOST_AUTH_NOT_FOUND = 426
	
	// Kubernetes集群相关常量
	KUBE_CLUSTER_CODE         = "kube_cluster:"
	KUBE_CLUSTER_CACHE_CODE   = "kube_cluster_cache:"
)
