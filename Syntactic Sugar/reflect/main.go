package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// ============================= Go反射核心知识 ====================

// 1. 基本类型反射 - 反射的入口
func basicReflection() {
	fmt.Println("// ============================= 1. 基本类型反射 ====================")

	var num int = 42
	str := "hello"
	slice := []int{1, 2, 3}

	// 获取类型信息
	typeOfNum := reflect.TypeOf(num)
	typeOfStr := reflect.TypeOf(str)
	typeOfSlice := reflect.TypeOf(slice)

	fmt.Printf("num类型: %v, Kind: %v\n", typeOfNum, typeOfNum.Kind())
	fmt.Printf("str类型: %v, Kind: %v\n", typeOfStr, typeOfStr.Kind())
	fmt.Printf("slice类型: %v, Kind: %v, 元素类型: %v\n",
		typeOfSlice, typeOfSlice.Kind(), typeOfSlice.Elem())

	// 获取值信息
	valueOfNum := reflect.ValueOf(num)
	valueOfStr := reflect.ValueOf(str)

	fmt.Printf("num值: %v, 可设置: %v\n", valueOfNum.Interface(), valueOfNum.CanSet())
	fmt.Printf("str值: %v, 可设置: %v\n", valueOfStr.Interface(), valueOfStr.CanSet())
}

// 2. 反射三大定律演示
func reflectionLaws() {
	fmt.Println("\n// ============================= 2. 反射三大定律 ====================")

	// 第一定律：接口值转反射对象
	var x interface{} = "hello reflection"
	reflectType := reflect.TypeOf(x)
	reflectValue := reflect.ValueOf(x)
	fmt.Printf("第一定律 - 接口→反射对象:\n")
	fmt.Printf("  类型: %v, 值: %v\n", reflectType, reflectValue.Interface())

	// 第二定律：反射对象转接口值
	originalValue := reflectValue.Interface()
	fmt.Printf("第二定律 - 反射对象→接口值: %v\n", originalValue)

	// 第三定律：要修改反射对象，值必须可设置
	fmt.Printf("第三定律 - 可设置性检查:\n")
	var y int = 100
	valueY := reflect.ValueOf(y)
	fmt.Printf("  直接值可设置: %v\n", valueY.CanSet())

	valueYPtr := reflect.ValueOf(&y).Elem()
	fmt.Printf("  指针元素可设置: %v\n", valueYPtr.CanSet())
}

// 3. 修改反射值 - 核心难点
func modifyReflection() {
	fmt.Println("\n// ============================= 3. 修改反射值 ====================")

	// 3.1 基本类型修改
	num := 42
	fmt.Printf("修改前: %d\n", num)

	valuePtr := reflect.ValueOf(&num).Elem() // 关键：获取可设置的Value
	valuePtr.SetInt(100)
	fmt.Printf("修改后: %d\n", num)

	// 3.2 字符串修改
	str := "hello"
	strPtr := reflect.ValueOf(&str).Elem()
	strPtr.SetString("world")
	fmt.Printf("字符串修改: %s\n", str)
}

// 4. 结构体反射 - 最常用场景
type Person struct {
	Name    string `json:"name" db:"username" validate:"required"`
	Age     int    `json:"age" db:"user_age"`
	Address string `json:"address"`
	salary  int    // 私有字段
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d years old", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) privateMethod() string {
	return "这是私有方法"
}

func structReflection() {
	fmt.Println("\n// ============================= 4. 结构体反射 ====================")

	person := Person{Name: "Alice", Age: 25, salary: 5000}

	// 4.1 类型信息
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)

	fmt.Printf("结构体: %v\n", t)
	fmt.Printf("字段数: %d, 方法数: %d\n", t.NumField(), t.NumMethod())

	// 4.2 遍历字段信息
	fmt.Println("\n字段详情:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		fmt.Printf("  [%d] %s %v ", i, field.Name, field.Type)
		fmt.Printf("可导出: %v ", field.IsExported())

		// 安全获取值（私有字段会panic）
		if field.IsExported() {
			fmt.Printf("值: %v", fieldValue.Interface())
		}
		fmt.Println()

		// 处理Tag
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			fmt.Printf("    JSON Tag: %s\n", jsonTag)
		}
		if dbTag := field.Tag.Get("db"); dbTag != "" {
			fmt.Printf("    DB Tag: %s\n", dbTag)
		}
	}

	// 4.3 方法信息
	fmt.Println("\n方法详情:")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("  [%d] %s 类型: %v\n", i, method.Name, method.Type)
	}
}

// 5. 高级结构体操作
func advancedStructOperations() {
	fmt.Println("\n// ============================= 5. 高级结构体操作 ====================")

	person := &Person{Name: "Bob", Age: 30, salary: 8000}
	v := reflect.ValueOf(person).Elem()

	// 5.1 修改公共字段
	if nameField := v.FieldByName("Name"); nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("Bob Updated")
		fmt.Printf("修改Name后: %+v\n", person)
	}

	// 5.2 修改私有字段（需要unsafe）
	if salaryField := v.FieldByName("salary"); salaryField.IsValid() {
		// 通过unsafe修改私有字段
		salaryPtr := reflect.NewAt(salaryField.Type(), unsafe.Pointer(salaryField.UnsafeAddr())).Elem()
		salaryPtr.SetInt(10000)
		fmt.Printf("修改salary后: %+v\n", person) // 注意：打印时不会显示私有字段
	}

	// 5.3 调用方法
	greetResult := reflect.ValueOf(person).MethodByName("Greet").Call(nil)
	fmt.Printf("方法调用结果: %v\n", greetResult[0].Interface())
}

// 6. 动态创建对象
func createInstances() {
	fmt.Println("\n// ============================= 6. 动态创建对象 ====================")

	// 6.1 创建切片
	sliceType := reflect.TypeOf([]int{})
	newSlice := reflect.MakeSlice(sliceType, 3, 5)
	for i := 0; i < newSlice.Len(); i++ {
		newSlice.Index(i).SetInt(int64(i * 10))
	}
	fmt.Printf("创建的切片: %v, 长度: %d, 容量: %d\n",
		newSlice.Interface(), newSlice.Len(), newSlice.Cap())

	// 6.2 创建映射
	mapType := reflect.TypeOf(map[string]int{})
	newMap := reflect.MakeMapWithSize(mapType, 5)
	newMap.SetMapIndex(reflect.ValueOf("key1"), reflect.ValueOf(100))
	newMap.SetMapIndex(reflect.ValueOf("key2"), reflect.ValueOf(200))
	fmt.Printf("创建的map: %v\n", newMap.Interface())

	// 6.3 创建结构体
	personType := reflect.TypeOf(Person{})
	newPerson := reflect.New(personType).Elem()
	newPerson.FieldByName("Name").SetString("Charlie")
	newPerson.FieldByName("Age").SetInt(28)
	fmt.Printf("创建的结构体: %+v\n", newPerson.Interface())
}

// 7. 类型判断和转换
func typeOperations() {
	fmt.Println("\n// ============================= 7. 类型判断和转换 ====================")

	values := []interface{}{42, "hello", 3.14, Person{}}

	for _, val := range values {
		rType := reflect.TypeOf(val)
		rValue := reflect.ValueOf(val)

		fmt.Printf("值: %v, 类型: %v, Kind: %v\n",
			val, rType, rType.Kind())

		// 类型转换检查
		if rValue.Kind() == reflect.Int {
			fmt.Printf("  → 可以调用Int(): %d\n", rValue.Int())
		}

		// 类型转换能力
		stringType := reflect.TypeOf("")
		if rType.ConvertibleTo(stringType) {
			fmt.Printf("  → 可以转换为字符串\n")
		}
	}
}

// 8. 深度相等比较
func deepEqualComparison() {
	fmt.Println("\n// ============================= 8. 深度相等比较 ====================")

	// 切片比较
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}

	fmt.Printf("slice1 == slice2: %v\n", reflect.DeepEqual(slice1, slice2))
	fmt.Printf("slice1 == slice3: %v\n", reflect.DeepEqual(slice1, slice3))

	// 结构体比较
	p1 := Person{Name: "Tom", Age: 20}
	p2 := Person{Name: "Tom", Age: 20}
	p3 := Person{Name: "Tom", Age: 21}

	fmt.Printf("p1 == p2: %v\n", reflect.DeepEqual(p1, p2))
	fmt.Printf("p1 == p3: %v\n", reflect.DeepEqual(p1, p3))

	// 嵌套结构体比较
	type Company struct {
		Name string
		CEO  Person
	}

	c1 := Company{Name: "ABC", CEO: p1}
	c2 := Company{Name: "ABC", CEO: p1}
	fmt.Printf("嵌套结构体比较: %v\n", reflect.DeepEqual(c1, c2))
}

func main() {
	basicReflection()
	reflectionLaws()
	modifyReflection()
	structReflection()
	advancedStructOperations()
	createInstances()
	typeOperations()
	deepEqualComparison()
}

// ============================= 总结知识点 ====================
/*
核心知识点总结：

1. 反射三大定律（核心思想）：
   - 接口值 → 反射对象 (TypeOf/ValueOf)
   - 反射对象 → 接口值 (Interface())
   - 修改反射对象必须可设置 (通过指针.Elem())

2. 关键类型和方法：
   - reflect.Type: 类型元数据
     - Kind(): 基础类型分类
     - Elem(): 获取元素类型
     - NumField()/Field(): 结构体字段
     - NumMethod()/Method(): 方法信息

   - reflect.Value: 值操作
     - Interface(): 获取原始值
     - SetXXX(): 设置值
     - Call(): 调用方法
     - CanSet(): 可设置性检查

3. 常用场景技巧：
   - 结构体字段遍历和Tag解析
   - 动态创建对象 (MakeSlice, MakeMap, New)
   - 私有字段修改 (unsafe操作)
   - 类型判断和转换

4. 注意事项：
   - 性能开销：反射比直接代码慢，避免在热点路径使用
   - 类型安全：运行时才能发现类型错误
   - 可读性：反射代码较难理解，要加注释
   - 私有成员：需要unsafe包，可能破坏封装

5. 最佳实践：
   - 明确使用场景（序列化、ORM、配置解析等）
   - 封装反射逻辑，提供类型安全接口
   - 做好错误处理和边界检查
*/
