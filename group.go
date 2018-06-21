package httprouter

import (
	"strings"
)

// RouterNode 路由分层节点
type RouterNode struct {
	prefix   string //url
	children []*RouterNode
	isroot   bool   //是否是根节点
	method   string //请求方式
	handle   Handle
}

// Request 请求数
type Request struct {
	URL    string `json:"url"`
	Method string `json:"method"`
	Handle Handle
}

// AddGroups 添加到组
func AddGroups(node *RouterNode) *Router {
	ns := routerNodes(node)
	r := New()
	for _, n := range ns {
		r.Handle(strings.ToUpper(n.Method), n.URL, n.Handle)
	}
	return r
}

// RouterNodes 添加路由
func routerNodes(node *RouterNode) []*Request {
	if node.isroot {
		return []*Request{&Request{URL: node.prefix, Method: node.method, Handle: node.handle}}
	}
	requsets := make([]*Request, 0)
	top := node.prefix
	for _, n := range node.childrens() {
		if n.isroot {
			requsets = append(requsets, &Request{URL: top + n.prefix, Method: n.method, Handle: n.handle})
			continue
		}
		nulrs := routerNodes(n)
		for _, nurl := range nulrs {
			nurl.URL = top + nurl.URL
		}
		requsets = append(requsets, nulrs...)
	}
	return requsets
}

// AddChildren 添加子节点
func (node *RouterNode) addChildren(no *RouterNode) {
	if node.children == nil || len(node.children) <= 0 {
		node.children = []*RouterNode{no}
		return
	}
	node.children = append(node.children, no)
}

// Childrens 获取子节点
func (node *RouterNode) childrens() []*RouterNode {
	return node.children
}

// LinkGroup 以func的方式添加到路由组
type LinkGroup func(*RouterNode)

// NewGroup 是当前的父节点
// prefix 是前缀
func NewGroup(prefix string, gs ...LinkGroup) *RouterNode {
	no := &RouterNode{
		prefix: prefix,
	}
	for _, f := range gs {
		f(no)
	}
	return no
}

// NSGroup 添加路由的方式
// prefix 前缀
func NSGroup(prefix string, gs ...LinkGroup) LinkGroup {
	return func(n *RouterNode) {
		no := NewGroup(prefix, gs...)
		//构造子节点
		n.addChild(no)

	}
}

// AddChild 构造子节点
func (node *RouterNode) addChild(ns ...*RouterNode) {
	for _, n := range ns {
		node.addChildren(n)
	}
}

// NSRouter 底层路由参数
func NSRouter(rootpath string, method string, handle Handle) LinkGroup {
	return func(no *RouterNode) {
		c := &RouterNode{
			prefix: rootpath,
			isroot: true,
			method: method,
			handle: handle,
		}
		no.addChildren(c)
	}
}
