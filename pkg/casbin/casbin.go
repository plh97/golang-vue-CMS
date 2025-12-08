package casbinPkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// NewEnforcer 初始化 Casbin Enforcer
func NewEnforcer(db *gorm.DB) *casbin.CachedEnforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Println("创建 Casbin 适配器失败:", err)
		return nil
	}
	cwd, _ := os.Getwd()
	m, err := model.NewModelFromFile(filepath.Join(cwd, "config", "model.conf"))
	if err != nil {
		fmt.Println("创建 NewModelFromFile 失败:", err)
		return nil
	}
	enforcer, err := casbin.NewCachedEnforcer(m, a)
	if err != nil {
		fmt.Println("创建 Casbin Enforcer 失败:", err)
		return nil
	}
	if enforcer == nil {
		fmt.Println("Casbin Enforcer 初始化失败，返回 nil")
		return nil
	}
	if err := enforcer.LoadPolicy(); err != nil {
		fmt.Println("加载 Casbin 策略失败:", err)
		return nil
	}
	return enforcer
}
