// base/factory.go
package base

// define interface for Class
type Class interface {
	Do()
}

var (
	// 存放註冊好的 類別工廠資訊
	factoryByName = make(map[string]func() Class)
)

// 註冊一個類別工廠
func Register(name string, factory func() Class) {
	factoryByName[name] = factory
}

// 根據name創建對應的類別
func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	}
	panic("name not found")
}
