package controller

import (
	"errtrack/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type RBACController struct {
	*zap.SugaredLogger
}

func NewRBACController(l *zap.SugaredLogger) *RBACController {
	return &RBACController{
		l,
	}
}

func (c *RBACController) RegisterRoutes(app *fiber.App) {
	r := app.Group("/rbac")

	// 用户角色
	r.Post("/users/:id/roles", c.assignRolesToUser)
	r.Get("/users/:id/permissions", c.getUserPermissions)

	// 角色管理
	r.Post("/roles", c.createRole)
	r.Get("/roles", c.listRoles)
	r.Post("/roles/:id/permissions", c.assignPermissionsToRole)
	r.Get("/roles/:id/permissions", c.getRolePermissions)

	// 权限管理
	r.Post("/permissions", c.createPermission)
	r.Get("/permissions", c.listPermissions)

	// 权限校验
	r.Post("/access/check", c.checkPermission)
}

func (c *RBACController) assignRolesToUser(ctx *fiber.Ctx) error {
	userID, _ := strconv.Atoi(ctx.Params("id")) // 获取 URL 参数中的 userID
	var payload struct {
		RoleIDs []int64 `json:"role_ids"`
	}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	fmt.Println(userID)
	// 在这里调用 service 层的功能，将角色分配给用户
	// err := userService.AssignRolesToUser(userID, payload.RoleIDs)
	// if err != nil {
	//    return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *RBACController) getUserPermissions(ctx *fiber.Ctx) error {
	userID, _ := strconv.Atoi(ctx.Params("id"))
	fmt.Println(userID)
	// 调用 service 层获取用户所有权限
	// permissions, err := userService.GetUserPermissions(userID)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	permissions := []model.Permission{
		{Action: "read", Resource: "article"},
		{Action: "write", Resource: "article"},
	} // 示例权限

	return ctx.JSON(fiber.Map{
		"permissions": permissions,
	})
}

func (c *RBACController) createRole(ctx *fiber.Ctx) error {
	var payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// 调用 service 层创建角色
	// role, err := roleService.CreateRole(payload.Name, payload.Description)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	role := model.Role{ID: 1, Name: payload.Name, Description: payload.Description} // 示例角色

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"role": role,
	})
}

func (c *RBACController) listRoles(ctx *fiber.Ctx) error {
	// 调用 service 层获取所有角色
	// roles, err := roleService.GetAllRoles()
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	roles := []model.Role{
		{ID: 1, Name: "admin", Description: "Administrator role"},
		{ID: 2, Name: "editor", Description: "Editor role"},
	} // 示例角色

	return ctx.JSON(fiber.Map{
		"roles": roles,
	})
}

func (c *RBACController) assignPermissionsToRole(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("id"))
	fmt.Println(roleID)
	var payload struct {
		PermissionIDs []int64 `json:"permission_ids"`
	}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// 在这里调用 service 层将权限分配给角色
	// err := roleService.AssignPermissions(roleID, payload.PermissionIDs)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *RBACController) getRolePermissions(ctx *fiber.Ctx) error {
	roleID, _ := strconv.Atoi(ctx.Params("id"))
	fmt.Println(roleID)
	// 调用 service 层获取角色权限
	// permissions, err := roleService.GetRolePermissions(roleID)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	permissions := []model.Permission{
		{Action: "read", Resource: "article"},
		{Action: "write", Resource: "article"},
	} // 示例权限

	return ctx.JSON(fiber.Map{
		"permissions": permissions,
	})
}

func (c *RBACController) createPermission(ctx *fiber.Ctx) error {
	var payload struct {
		Action   string `json:"action"`
		Resource string `json:"resource"`
	}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// 调用 service 层创建权限
	// permission, err := permissionService.CreatePermission(payload.Action, payload.Resource)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	permission := model.Permission{ID: 1, Action: payload.Action, Resource: payload.Resource} // 示例权限

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"permission": permission,
	})
}

func (c *RBACController) listPermissions(ctx *fiber.Ctx) error {
	// 调用 service 层获取所有权限
	// permissions, err := permissionService.GetAllPermissions()
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	permissions := []model.Permission{
		{ID: 1, Action: "read", Resource: "article"},
		{ID: 2, Action: "write", Resource: "article"},
	} // 示例权限

	return ctx.JSON(fiber.Map{
		"permissions": permissions,
	})
}

func (c *RBACController) checkPermission(ctx *fiber.Ctx) error {
	var payload struct {
		UserID   int64  `json:"user_id"`
		Action   string `json:"action"`
		Resource string `json:"resource"`
	}
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// 调用 service 层检查权限
	// allowed, err := accessControl.CheckPermission(payload.UserID, payload.Action, payload.Resource)
	// if err != nil {
	//     return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	allowed := true // 示例，表示允许该操作

	return ctx.JSON(fiber.Map{
		"allowed": allowed,
	})
}
