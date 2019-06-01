
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TTaskDialogProgressBar struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTaskDialogProgressBar
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialogProgressBar() *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance = TaskDialogProgressBar_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogProgressBarFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TaskDialogProgressBarFromInst(inst uintptr) *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TaskDialogProgressBarFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TaskDialogProgressBarFromObj(obj IObject) *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogProgressBarFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TaskDialogProgressBarFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogProgressBar {
    t := new(TTaskDialogProgressBar)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialogProgressBar) Free() {
    if t.instance != 0 {
        TaskDialogProgressBar_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogProgressBar) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogProgressBar) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogProgressBar) IsValid() bool {
    return t.instance != 0
}

// TTaskDialogProgressBarClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogProgressBarClass() TClass {
    return TaskDialogProgressBar_StaticClassType()
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogProgressBar) Assign(Source IObject) {
    TaskDialogProgressBar_Assign(t.instance, CheckPtr(Source))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogProgressBar) GetNamePath() string {
    return TaskDialogProgressBar_GetNamePath(t.instance)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskDialogProgressBar) DisposeOf() {
    TaskDialogProgressBar_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogProgressBar) ClassType() TClass {
    return TaskDialogProgressBar_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogProgressBar) ClassName() string {
    return TaskDialogProgressBar_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogProgressBar) InstanceSize() int32 {
    return TaskDialogProgressBar_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogProgressBar) InheritsFrom(AClass TClass) bool {
    return TaskDialogProgressBar_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogProgressBar) Equals(Obj IObject) bool {
    return TaskDialogProgressBar_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogProgressBar) GetHashCode() int32 {
    return TaskDialogProgressBar_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogProgressBar) ToString() string {
    return TaskDialogProgressBar_ToString(t.instance)
}

// Max
func (t *TTaskDialogProgressBar) Max() int32 {
    return TaskDialogProgressBar_GetMax(t.instance)
}

// SetMax
func (t *TTaskDialogProgressBar) SetMax(value int32) {
    TaskDialogProgressBar_SetMax(t.instance, value)
}

// Min
func (t *TTaskDialogProgressBar) Min() int32 {
    return TaskDialogProgressBar_GetMin(t.instance)
}

// SetMin
func (t *TTaskDialogProgressBar) SetMin(value int32) {
    TaskDialogProgressBar_SetMin(t.instance, value)
}

// Position
func (t *TTaskDialogProgressBar) Position() int32 {
    return TaskDialogProgressBar_GetPosition(t.instance)
}

// SetPosition
func (t *TTaskDialogProgressBar) SetPosition(value int32) {
    TaskDialogProgressBar_SetPosition(t.instance, value)
}

// State
func (t *TTaskDialogProgressBar) State() TProgressBarState {
    return TaskDialogProgressBar_GetState(t.instance)
}

// SetState
func (t *TTaskDialogProgressBar) SetState(value TProgressBarState) {
    TaskDialogProgressBar_SetState(t.instance, value)
}

