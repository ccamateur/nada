package ConstansParser

import(
	VmObject "vm/object"
)
/**
常量池对象 int
 */
type ParserMethod struct {
	val int
}


func (me *ParserMethod) Parser(idx int) VmObject.IFObject {
	return nil;
}


