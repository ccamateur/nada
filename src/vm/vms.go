package Vm


type Vms struct {

	//所有的Klass元类
	metaKlass *FMap

	//项目的class目录 默认当前的lib目录
	classPath string

	//所有的解释器

}

func (me *Vms ) start()  {


	//初始化参数 klass等

	//str :=  NewFString("test")
	//call(str,"print")
	//msg :=   Vm.NewFString("奥术大师多")


	var fileName = "E:\\AAAA_CODE\\new-eclipse-workspace\\far-dev\\demo\\lib\\Test.class"
	classN  := NewClassByteFile()
	classN.Load(fileName).Parser();
	main:=me.metaKlass.get("Test").(FMethod)
	interp:=new(Interpreter)
	interp.setFrame()

}




func (me *Vms ) Build() *Vms{
	me.metaKlass = NewFMap()
	me.metaKlass.set(BUILTN.NADA_TYPE,   	new(TypeKlass).Init())
	me.metaKlass.set(BUILTN.NADA_NATIVE, 	new(NativeMethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_VIRTUAL,	new(VirtualMethodKlass).Init())
	me.metaKlass.set(BUILTN.NADA_STRING,    new(StringKlass).Init())
	return me
}



