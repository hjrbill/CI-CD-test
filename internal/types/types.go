// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me|Eric"`
}

type Response struct {
	Message string `json:"message"`
}
