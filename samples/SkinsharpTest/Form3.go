// 由GOVCL UI设计器自动生成，不要编辑。
// Automatically generated by the GOVCL UI designer, do not edit.
package main

import (
    "gitee.com/ying32/govcl/vcl"
)

type TForm3 struct {
    *vcl.TForm
    ListBox1      *vcl.TListBox
    CheckListBox1 *vcl.TCheckListBox
    ScrollBox1    *vcl.TScrollBox
    ListView1     *vcl.TListView
    TreeView1     *vcl.TTreeView
    ColorBox1     *vcl.TColorBox
    Button1       *vcl.TButton
}

var Form3 *TForm3




// 以字节形式加载
// Loaded in bytes.
// vcl.Application.CreateFormFromBytes(form3Bytes, &Form3)

var (
    form3Bytes = []byte {
        0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00, 
        0x93, 0x48, 0x46, 0x1E, 0x5A, 0x60, 0xA3, 0x1B, 0x99, 0xA6, 0x0F, 0x52, 
        0xF0, 0xA0, 0xF6, 0x15, 0xB9, 0x71, 0x65, 0x9E, 0xCC, 0x41, 0xA9, 0x36, 
        0x75, 0xC3, 0xA3, 0xA3, 0x9D, 0x49, 0x18, 0xDE, 0xC0, 0x72, 0x5E, 0xF9, 
        0xB5, 0x68, 0x23, 0x0B, 0xB9, 0xD3, 0xEE, 0x68, 0x40, 0x9F, 0xD0, 0x38, 
        0x1D, 0x0A, 0xD0, 0x56, 0x37, 0x41, 0xEF, 0x5C, 0xE1, 0x39, 0xCA, 0xBD, 
        0xDC, 0xB9, 0x79, 0xEE, 0x91, 0x3C, 0x8D, 0xA8, 0x6C, 0x78, 0x87, 0x42, 
        0xFF, 0xCE, 0x49, 0xD2, 0x52, 0x2F, 0xC3, 0x34, 0x2D, 0xCC, 0x68, 0x63, 
        0xA6, 0xB4, 0x50, 0x1C, 0xDA, 0x4C, 0xA6, 0x22, 0x11, 0x09, 0xD5, 0x71, 
        0x68, 0x78, 0x08, 0x23, 0x75, 0x6E, 0x77, 0xF9, 0xE2, 0x7F, 0x95, 0x39, 
        0x11, 0x9E, 0x31, 0x9F, 0x0C, 0x12, 0xB5, 0x8A, 0x98, 0xA0, 0x72, 0xA2, 
        0x1C, 0xDB, 0x3E, 0x89, 0xA9, 0xE0, 0x0F, 0x32, 0xB6, 0x38, 0x58, 0x82, 
        0x45, 0x13, 0x2A, 0x34, 0xE5, 0x96, 0xFE, 0x23, 0x29, 0x6A, 0x49, 0xF8, 
        0x5C, 0x99, 0xAC, 0x08, 0xD3, 0x70, 0x0E, 0x61, 0x92, 0x27, 0x91, 0x4D, 
        0x2B, 0x0E, 0x58, 0x4F, 0x38, 0x1F, 0x56, 0xD1, 0x36, 0xDC, 0xED, 0x83, 
        0x6F, 0xC0, 0xC6, 0x8C, 0x08, 0x8F, 0x28, 0x53, 0xE7, 0x60, 0xB6, 0x9D, 
        0xEF, 0x8F, 0x06, 0xDE, 0x76, 0x11, 0xEB, 0x3E, 0xF3, 0x4D, 0x8C, 0xF5, 
        0xD9, 0x98, 0x88, 0xE7, 0xDB, 0x64, 0x80, 0x7B, 0x7E, 0x96, 0x85, 0x90, 
        0x4C, 0xF7, 0xF7, 0x06, 0x64, 0x98, 0x11, 0x2E, 0x29, 0x8B, 0xC4, 0x11, 
        0xE4, 0x56, 0x43, 0x43, 0x2A, 0xDD, 0x3C, 0x83, 0x1A, 0x42, 0xA5, 0xF7, 
        0xC2, 0x2D, 0xB7, 0xD2, 0x26, 0x5F, 0x7F, 0xE3, 0x97, 0x5F, 0xED, 0x3A, 
        0x93, 0x8F, 0xDA, 0xA6, 0xDA, 0x14, 0x57, 0x76, 0x07, 0xEF, 0x3A, 0xBB, 
        0x19, 0x4E, 0xEC, 0xF6, 0xC7, 0x3A, 0xF8, 0x5B, 0x15, 0x58, 0x79, 0xA3, 
        0xB9, 0x67, 0xDE, 0xBC, 0x33, 0x5F, 0x67, 0x50, 0xF5, 0x95, 0x3E, 0x44, 
        0xFE, 0x17, 0x52, 0x3C, 0x21, 0x2B, 0x8E, 0xBC, 0xB7, 0x6B, 0xB4, 0x56, 
        0x8C, 0xBC, 0xB3, 0x61, 0x3E, 0xE9, 0x7E, 0x55, 0x12, 0x81, 0x43, 0x57, 
        0x60, 0x89, 0x7A, 0xCD, 0x5F, 0xE8, 0x93, 0x9C, 0x9B, 0xD3, 0x7E, 0x90, 
        0x52, 0xDC, 0x9A, 0x84, 0x15, 0xD4, 0x4D, 0x38, 0x84, 0x84, 0x8D, 0x55, 
        0xFC, 0xC3, 0xEA, 0x4D, 0xC2, 0x8D, 0xB9, 0x64, 0x68, 0xCB, 0xCC, 0xD5, 
        0x7C, 0x71, 0xAE, 0xDE, 0x61, 0x57, 0x02, 0xCB, 0x63, 0x74, 0x19, 0xC8, 
        0x0A, 0xAD, 0x14, 0x7F, 0x60, 0x33, 0x52, 0x49, 0x3F, 0xC8, 0x5A, 0xB2, 
        0xEA, 0x6F, 0x22, 0x45, 0xE5, 0x1D, 0x9C, 0x9B, 0x1A, 0x4E, 0x48, 0x75, 
        0x9E, 0xEF, 0x73, 0x9A, 0xCE, 0xE0, 0xD9, 0x14, 0x41, 0x52, 0xA5, 0x4A, 
        0x8F, 0xAE, 0xD8, 0xC4, 0x22, 0xB6, 0x7E, 0x9D, 0x5D, 0x2E, 0x50, 0x45, 
        0x2F, 0xF3, 0x06, 0x10, 0xA2, 0x86, 0xAA, 0xDB, 0x89, 0x5E, 0xDA, 0xC9, 
        0x72, 0x58, 0x8F, 0xE9, 0x14, 0x14, 0x79, 0x1D, 0x1B, 0x49, 0xF0, 0xFF, 
        0x64, 0x5E, 0xE0, 0xD5, 0x95, 0x4E, 0xEA, 0xB4, 0xD4, 0x14, 0xD8, 0x92, 
        0x9C, 0x9B, 0x34, 0xD2, 0x27, 0x39, 0xED, 0x81, 0x16, 0x6A, 0xD9, 0x75, 
        0xEB, 0xEA, 0xFE, 0xA1, 0x6B, 0xEE, 0xBE, 0xC9, 0x7F, 0x34, 0x0A, 0xD2, 
        0x20, 0x36, 0xED, 0xF1, 0x25, 0xB2, 0x26, 0xBE, 0x9F, 0x28, 0xD9, 0x0E, 
        0x1D, 0xAD, 0x83, 0x63, 0x88, 0xBA, 0x27, 0x84, 0xE1, 0x3D, 0x3A, 0x08, 
        0x21, 0xF1, 0xC9, 0xF1, 0x33, 0x53, 0x16, 0xB2, 0x1D, 0xF3, 0x95, 0x37, 
        0xA3, 0x4F, 0x14, 0x18, 0xB4, 0xFA, 0x8A, 0x5C, 0x7D, 0x47, 0x68, 0xF9, 
        0x15, 0x7F, 0x80, 0xF9, 0x88, 0xE3, 0x94, 0x04, 0xAF, 0xBF, 0x7F, 0xDB, 
        0x3F, 0x20, 0x34, 0x44, 0x28, 0x00, 0x01, 0x80, 0xF1, 0xC7, 0xE8, 0x6A, 
        0x5B, 0xDB, 0xEA, 0x8B, 0xDD, 0x1D, 0xF7, 0x51, 0x56, 0x00, 0xD3, 0x59, 
        0xED, 0xEC, 0xFC, 0x2E, 0x2A, 0xDF, 0xFF, 0xE9, 0xAF, 0x8D, 0x57, 0x5F, 
        0x89, 0x99, 0xE5, 0x0C, 0x6A, 0xD9, 0xEA, 0xCD, 0x5C, 0x3D, 0x73, 0x42, 
        0xFE, 0xF2, 0xAB, 0xCD, 0xEA, 0x0C, 0x02, 0xEA, 0x18, 0xC2, 0xFB, 0xAA, 
        0x5B, 0x55, 0xF8, 0x80, 0x2A, 0x31, 0x95, 0x5E, 0xB5, 0xD2, 0xDE, 0x25, 
        0xE0, 0x7B, 0xAA, 0xD4, 0x30, 0xD3, 0x0F, 0x04, 0x70, 0x91, 0xD6, 0xE1, 
        0xF2, 0x38, 0x99, 0xA5, 0xF1, 0xBF, 0x32, 0x86, 0xD0, 0x2B, 0xD0, 0x4D, 
        0x9B, 0x2E, 0xD2, 0xE4, 0x71, 0x6F, 0x0C, 0x2E, 0x09, 0x39, 0x84, 0xDC, 
        0x11, 0xFB, 0x68, 0x53, 0x5F, 0xD5, 0xD7, 0x3E, 0x57, 0xCD, 0x4D, 0x3A, 
        0xD0, 0x26, 0xBC, 0x55, 0x4E, 0xB7}
)
