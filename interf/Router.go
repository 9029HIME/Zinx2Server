package interf

/**
Router是用来代替HandleAPI的组件
*/
type AbstractRouter interface {
	PreHandler(request AbstractRequest)
	DoHandle(request AbstractRequest)
	PostHandler(request AbstractRequest)
}
