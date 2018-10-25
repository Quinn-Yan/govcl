// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    Splitter1  *vcl.TSplitter
    ToolBar1   *vcl.TToolBar
    StatusBar1 *vcl.TStatusBar
    Panel1     *vcl.TPanel
    Panel2     *vcl.TPanel
    Button1    *vcl.TButton

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(mainFormBytes, &MainForm)

var (
    mainFormBytes = []byte {
        0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00, 
        0x93, 0x48, 0x16, 0xDA, 0x5A, 0x80, 0xEB, 0x6B, 0x99, 0x7F, 0xF2, 0x56, 
        0xF7, 0xDF, 0xEA, 0x37, 0x2F, 0xD5, 0x1A, 0x0B, 0x86, 0x59, 0x00, 0x7B, 
        0x1F, 0x93, 0x88, 0xE4, 0xEA, 0x79, 0xFA, 0x90, 0x86, 0x90, 0x1F, 0xBE, 
        0x5F, 0xB4, 0x91, 0x32, 0xA5, 0xE4, 0x94, 0x98, 0xED, 0xCC, 0x6F, 0xAF, 
        0x07, 0xCF, 0x53, 0xE3, 0xBB, 0xEC, 0x55, 0xCB, 0x31, 0xC8, 0xC9, 0x5C, 
        0x52, 0x36, 0x70, 0xB1, 0x7E, 0x34, 0x77, 0x51, 0x5A, 0xD9, 0x57, 0xA5, 
        0xEA, 0x24, 0x09, 0xFD, 0x4D, 0xED, 0x2B, 0x8C, 0x7A, 0x1C, 0xE1, 0xC5, 
        0xCD, 0xED, 0xE6, 0x42, 0x74, 0x18, 0x33, 0x32, 0x53, 0xA8, 0xDC, 0xFC, 
        0xD3, 0x45, 0x76, 0xDA, 0x01, 0x97, 0xC2, 0x9F, 0xA0, 0xFC, 0xA4, 0x38, 
        0xA9, 0x00, 0x83, 0xEB, 0xDB, 0xF2, 0x88, 0x8D, 0x33, 0x76, 0xD8, 0x59, 
        0x89, 0xE6, 0xE2, 0x6C, 0x4D, 0xEF, 0x3C, 0xF9, 0xDA, 0xB3, 0xC7, 0xE0, 
        0x84, 0xC4, 0x6E, 0xDE, 0xFB, 0xA7, 0xD8, 0x90, 0xDA, 0x8E, 0x1E, 0x19, 
        0xC1, 0x08, 0x0C, 0xF5, 0xE9, 0x22, 0x81, 0x77, 0x28, 0x27, 0xDC, 0x5B, 
        0x23, 0x18, 0x64, 0xD5, 0x7F, 0x9C, 0x32, 0xD6, 0x50, 0x0B, 0x11, 0x70, 
        0x23, 0x58, 0xD2, 0x31, 0x58, 0x23, 0xD3, 0xCB, 0x6D, 0x0B, 0x54, 0x6F, 
        0x74, 0xB9, 0xE1, 0x4F, 0x8E, 0x8E, 0x45, 0xB8, 0x93, 0x8D, 0x32, 0x8F, 
        0x5A, 0x40, 0xC4, 0xD0, 0x79, 0x5A, 0x50, 0x59, 0xFF, 0x32, 0xDC, 0x09, 
        0x06, 0xCF, 0x21, 0x10, 0x42, 0x33, 0xF2, 0x1A, 0x5F, 0x3B, 0x89, 0xDE, 
        0x8F, 0x8F, 0x2B, 0x11, 0x08, 0x4E, 0x6B, 0x99, 0xCE, 0x5E, 0x5D, 0xA3, 
        0xB2, 0x1E, 0x72, 0x3E, 0x4C, 0xD2, 0x41, 0xEA, 0xED, 0xD9, 0x81, 0x93, 
        0xC2, 0xC1, 0x22, 0x2E, 0x43, 0x9F, 0x81, 0xEB, 0xF0, 0x88, 0x87, 0xC7, 
        0x81, 0xE7, 0x8D, 0xD9, 0xD0, 0x1E, 0x38, 0xDC, 0x01, 0xA9, 0x1A, 0x74, 
        0x15, 0x33, 0x8F, 0x4C, 0xEE, 0x5D, 0xB3, 0xBC, 0xF4, 0xB3, 0x88, 0x86, 
        0x03, 0x20, 0xCE, 0x00, 0x6A, 0x70, 0x07, 0x35, 0xEF, 0x49, 0x44, 0xFA, 
        0xB5, 0x77, 0x56, 0x67, 0x87, 0x43, 0x4F, 0x61, 0xAC, 0xDF, 0xE5, 0x16, 
        0x8E, 0x78, 0x60, 0xAC, 0x2D, 0x86, 0x61, 0x41, 0xBE, 0x5F, 0xCB, 0x45, 
        0xFC, 0xEC, 0xC0, 0x65, 0x83, 0x24, 0x90, 0x20, 0xBA, 0x7F, 0x2B, 0xFF, 
        0xB8, 0x4E, 0xA0, 0x77, 0x45, 0xE2, 0xA1, 0x3C, 0x26, 0xF1, 0xE1, 0xAD, 
        0x5E, 0x79, 0xE7, 0x68, 0x4F, 0x55, 0x4F, 0x20, 0x33, 0x10, 0x58, 0x02, 
        0xC5, 0xB0, 0x3A, 0xAE, 0x4A, 0xE3, 0x8F, 0xCA, 0x27, 0x0B, 0x0C, 0x88, 
        0x8E, 0x15, 0x48, 0xD8, 0x2D, 0x15, 0x2D, 0x3B, 0xA6, 0x72, 0x5E, 0xA7, 
        0x37, 0x80, 0xA2, 0xF5, 0xBE, 0x58, 0x59, 0x0F, 0xCF, 0x89, 0x1E, 0x29, 
        0x6B, 0xCA, 0x72, 0x3D, 0xF5, 0x3E, 0x99, 0xCC, 0x1D, 0x77, 0x98, 0x88, 
        0x11, 0xF9, 0x80, 0x92, 0x8C, 0x86, 0x43, 0x13, 0xE3, 0xA9, 0x13, 0x96, 
        0x5E, 0xB6, 0xA6, 0xA9, 0x3C, 0x97, 0x00, 0x67, 0xCB, 0xEB, 0x6E, 0xC3, 
        0x72, 0xAB, 0xAD, 0xA7, 0xF2, 0xAF, 0xC9, 0x8B, 0xDE, 0x11, 0xAB, 0xAD, 
        0x85, 0xEB, 0x77, 0x77, 0xE6, 0xD2, 0xC7, 0xBB, 0xFD, 0x2D, 0xFE, 0x99, 
        0x2E, 0xE3, 0xEC, 0x80, 0x32, 0xFE, 0x36, 0xF7, 0x0F, 0x73, 0x22, 0xEA, 
        0xDB, 0x7B, 0x58, 0xB3, 0xC8, 0x8D, 0x7A, 0x51, 0xEE, 0xCA, 0x7E, 0x87, 
        0x8D, 0x6A, 0xBD, 0xE7, 0xF3, 0x8A, 0x71, 0x1A, 0x9C, 0xAB, 0xDE, 0xE4, 
        0x43, 0x2B, 0x36, 0x26, 0x68, 0x1F, 0x7E, 0x16, 0xBE, 0x5B, 0xC6, 0xA9, 
        0xB2, 0x2E, 0x5F, 0x06, 0xC7, 0xDA}
)
