package controller

import (
	"gin-api/api/system/model"
	"gin-api/api/system/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var sysDept model.SysDept

// 查询部门列表
// @Summary 查询部门列表接口
// @Produce json
// @Description 查询部门列表接口
// @Param deptName query string false "部门名称"
// @Param deptStatus query string false "部门状态"
// @Success 200 {object} result.Result
// @router /api/dept/list [get]
// @Security ApiKeyAuth
func GetSysDeptList(c *gin.Context) {
	DeptName := c.Query("deptName")
	DeptStatus := c.Query("deptStatus")
	service.SysDeptService().GetSysDeptList(c, DeptName, DeptStatus)
}

// 新增部门
// @Summary 新增部门接口
// @Produce json
// @Description 新增部门接口
// @Param data body model.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/add [post]
// @Security ApiKeyAuth
func CreateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().CreateSysDept(c, sysDept)
}

// 根据id查询部门
// @Summary 根据id查询部门接口
// @Produce json
// @Description 根据id查询部门接口
// @Param id query int true "ID"
// @Success 200 {object} result.Result
// @router /api/dept/info [get]
// @Security ApiKeyAuth
func GetSysDeptById(c *gin.Context) {
	Id, _ := strconv.Atoi(c.Query("id"))
	service.SysDeptService().GetSysDeptById(c, Id)
}

// 修改部门
// @Summary 修改部门接口
// @Produce json
// @Description 修改部门接口
// @Param data body model.SysDept true "data"
// @Success 200 {object} result.Result
// @router /api/dept/update [put]
// @Security ApiKeyAuth
func UpdateSysDept(c *gin.Context) {
	_ = c.BindJSON(&sysDept)
	service.SysDeptService().UpdateSysDept(c, sysDept)
}

// 根据id删除部门
// @Summary 根据id删除部门接口
// @Produce json
// @Description 根据id删除部门接口
// @Param data body model.SysDeptIdDto true "data"
// @Success 200 {object} result.Result
// @router /api/dept/delete [delete]
// @Security ApiKeyAuth
func DeleteSysDeptById(c *gin.Context) {
	var dto model.SysDeptIdDto
	_ = c.BindJSON(&dto)
	service.SysDeptService().DeleteSysDeptById(c, dto)
}

// 部门下拉列表
// @Summary 部门下拉列表接口
// @Produce json
// @Description 部门下拉列表接口
// @Success 200 {object} result.Result
// @router /api/dept/vo/list [get]
// @Security ApiKeyAuth
func QuerySysDeptVoList(c *gin.Context) {
	service.SysDeptService().QuerySysDeptVoList(c)
}

// 获取某部门下的所有用户
// @Summary 获取某部门下的所有用户接口
// @Produce json
// @Description 获取某部门下的所有用户
// @Param deptId query int true "部门ID"
// @Success 200 {object} result.Result
// @router /api/dept/users [get]
// @Security ApiKeyAuth
func GetDeptUsers(c *gin.Context) {
	deptId, _ := strconv.Atoi(c.Query("deptId"))
	service.SysDeptService().GetDeptUsers(c, deptId)
}
