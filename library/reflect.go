package library

import (
	"fmt"
	"reflect"
)

/*
变量包含类型信息和值信息
 	var arr [10]int
	arr[0] = 10
类型信息：是静态的元信息，是预先定义好的
值信息：是程序运行过程中动态改变的

获取类型信息：reflect.TypeOf，是静态的
获取值信息：reflect.ValueOf，是动态的
*/

// 反射可以在运行时动态获取程序的各种详细信息
// 反射获取interface类型信息
func TestReflectInterface() {
	var x float64 = 3.14
	fmt.Println("****反射获取类型****")
	reflectType(x)
	fmt.Println("****反射获取值****")
	reflectValue(x)
	fmt.Println("****反射修改值****")
	reflectUpdateValue(&x)
}

func TestReflectStruct() {

	c := Country{"中国", "北京", 9600000.00}
	getStructOfReflect(c)

	fmt.Printf("反射修改之前: %#v\n", c)
	setStructOfReflect(&c)
	fmt.Printf("反射修改之后: %#v\n", c)

	invokeMethodOfReflect(&c)
}

func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是: ", t)
	k := t.Kind()
	fmt.Println("类型是: ", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64")
	case reflect.String:
		fmt.Println("a is string")
	}
}

func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	fmt.Println("类型是: ", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64, a = ", v.Float())
	case reflect.String:
		fmt.Println("a is string")
	}
}

func reflectUpdateValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	fmt.Println("类型是: ", k)
	switch k {
	case reflect.Float64: // 会出错，得通过下面方式使用指针修改
		v.SetFloat(6.88)
		fmt.Println("a is float64, a = ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(23.43)
		fmt.Println("case:", v.Elem().Float())
		fmt.Println(v.Pointer())
	}
}

type Country struct {
	// 使用反射，字段名首字母一定要大写，代表public访问
	Name    string
	Capital string
	Area    float64
}

func (c Country) ToString() {
	fmt.Printf("%#v", c)
}

func (c *Country) SetCapitalAndArea(Capitial string, Area float64) {
	c.Capital = Capitial
	c.Area = Area
}

func getStructOfReflect(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("类型是: ", t)
	fmt.Println("字符串类型: ", t.Name())

	v := reflect.ValueOf(i)
	fmt.Println("get value = ", v)

	fmt.Println("*****获取结构体所有属性*****")
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i)
		fieldValue := v.Field(i).Interface()
		fmt.Printf("%s : %v : %v \n", fieldName.Name, fieldName.Type, fieldValue)
	}
	fmt.Println("*****获取结构体所有方法*****")
	for i := 0; i < t.NumMethod(); i++ {
		methodName := t.Method(i)
		fmt.Printf("%s : %v\n", methodName.Name, methodName.Type)
	}
}

func setStructOfReflect(i interface{}) {
	v := reflect.ValueOf(i)
	// 获取指针
	v = v.Elem()

	field := v.FieldByName("Name")
	if field.Kind() == reflect.String {
		field.SetString("China")
	}
}

func invokeMethodOfReflect(i interface{}) {
	v := reflect.ValueOf(i)
	// 有参数方法
	fmt.Println("调用SetCapitalAndArea方法")
	setCapitalAndAreaMethod := v.MethodByName("SetCapitalAndArea")
	args := []reflect.Value{
		reflect.ValueOf("beijing"),
		reflect.ValueOf(980.000),
	}
	setCapitalAndAreaMethod.Call(args)

	// 无参数方法
	fmt.Println("调用ToString方法")
	toStringMethod := v.MethodByName("ToString")
	var notArgs []reflect.Value
	toStringMethod.Call(notArgs)
}
