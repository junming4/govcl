
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

type TTaskDialogButtons struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewTaskDialogButtons
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewTaskDialogButtons() *TTaskDialogButtons {
    t := new(TTaskDialogButtons)
    t.instance = TaskDialogButtons_Create()
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogButtonsFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func TaskDialogButtonsFromInst(inst uintptr) *TTaskDialogButtons {
    t := new(TTaskDialogButtons)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// TaskDialogButtonsFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func TaskDialogButtonsFromObj(obj IObject) *TTaskDialogButtons {
    t := new(TTaskDialogButtons)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// TaskDialogButtonsFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func TaskDialogButtonsFromUnsafePointer(ptr unsafe.Pointer) *TTaskDialogButtons {
    t := new(TTaskDialogButtons)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TTaskDialogButtons) Free() {
    if t.instance != 0 {
        TaskDialogButtons_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TTaskDialogButtons) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TTaskDialogButtons) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TTaskDialogButtons) IsValid() bool {
    return t.instance != 0
}

// TTaskDialogButtonsClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TTaskDialogButtonsClass() TClass {
    return TaskDialogButtons_StaticClassType()
}

// Add
func (t *TTaskDialogButtons) Add() *TTaskDialogBaseButtonItem {
    return TaskDialogBaseButtonItemFromInst(TaskDialogButtons_Add(t.instance))
}

// Buttons
func (t *TTaskDialogButtons) Buttons() uintptr {
    return TaskDialogButtons_Buttons(t.instance)
}

// Owner
// CN: 组件所有者。
// EN: component owner.
func (t *TTaskDialogButtons) Owner() *TObject {
    return ObjectFromInst(TaskDialogButtons_Owner(t.instance))
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (t *TTaskDialogButtons) Assign(Source IObject) {
    TaskDialogButtons_Assign(t.instance, CheckPtr(Source))
}

// BeginUpdate
func (t *TTaskDialogButtons) BeginUpdate() {
    TaskDialogButtons_BeginUpdate(t.instance)
}

// Clear
// CN: 清除。
// EN: .
func (t *TTaskDialogButtons) Clear() {
    TaskDialogButtons_Clear(t.instance)
}

// ClearAndResetID
func (t *TTaskDialogButtons) ClearAndResetID() {
    TaskDialogButtons_ClearAndResetID(t.instance)
}

// Delete
func (t *TTaskDialogButtons) Delete(Index int32) {
    TaskDialogButtons_Delete(t.instance, Index)
}

// EndUpdate
func (t *TTaskDialogButtons) EndUpdate() {
    TaskDialogButtons_EndUpdate(t.instance)
}

// FindItemID
func (t *TTaskDialogButtons) FindItemID(ID int32) *TCollectionItem {
    return CollectionItemFromInst(TaskDialogButtons_FindItemID(t.instance, ID))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (t *TTaskDialogButtons) GetNamePath() string {
    return TaskDialogButtons_GetNamePath(t.instance)
}

// Insert
func (t *TTaskDialogButtons) Insert(Index int32) *TCollectionItem {
    return CollectionItemFromInst(TaskDialogButtons_Insert(t.instance, Index))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TTaskDialogButtons) DisposeOf() {
    TaskDialogButtons_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TTaskDialogButtons) ClassType() TClass {
    return TaskDialogButtons_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TTaskDialogButtons) ClassName() string {
    return TaskDialogButtons_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TTaskDialogButtons) InstanceSize() int32 {
    return TaskDialogButtons_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TTaskDialogButtons) InheritsFrom(AClass TClass) bool {
    return TaskDialogButtons_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TTaskDialogButtons) Equals(Obj IObject) bool {
    return TaskDialogButtons_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TTaskDialogButtons) GetHashCode() int32 {
    return TaskDialogButtons_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TTaskDialogButtons) ToString() string {
    return TaskDialogButtons_ToString(t.instance)
}

// DefaultButton
func (t *TTaskDialogButtons) DefaultButton() *TTaskDialogBaseButtonItem {
    return TaskDialogBaseButtonItemFromInst(TaskDialogButtons_GetDefaultButton(t.instance))
}

// SetDefaultButton
func (t *TTaskDialogButtons) SetDefaultButton(value *TTaskDialogBaseButtonItem) {
    TaskDialogButtons_SetDefaultButton(t.instance, CheckPtr(value))
}

// Capacity
func (t *TTaskDialogButtons) Capacity() int32 {
    return TaskDialogButtons_GetCapacity(t.instance)
}

// SetCapacity
func (t *TTaskDialogButtons) SetCapacity(value int32) {
    TaskDialogButtons_SetCapacity(t.instance, value)
}

// Count
func (t *TTaskDialogButtons) Count() int32 {
    return TaskDialogButtons_GetCount(t.instance)
}

// Items
func (t *TTaskDialogButtons) Items(Index int32) *TTaskDialogBaseButtonItem {
    return TaskDialogBaseButtonItemFromInst(TaskDialogButtons_GetItems(t.instance, Index))
}

// Items
func (t *TTaskDialogButtons) SetItems(Index int32, value *TTaskDialogBaseButtonItem) {
    TaskDialogButtons_SetItems(t.instance, Index, CheckPtr(value))
}

