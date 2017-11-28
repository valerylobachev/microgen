package template

import (
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/devimteam/microgen/util"
)

type httpServerTemplate struct {
	Info *GenerationInfo
}

func NewHttpServerTemplate(info *GenerationInfo) Template {
	return &httpServerTemplate{
		Info: info.Copy(),
	}
}

func (t *httpServerTemplate) DefaultPath() string {
	return "./transport/http/server.go"
}

func (t *httpServerTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	if err := util.StatFile(t.Info.AbsOutPath, t.DefaultPath()); !t.Info.Force && err == nil {
		return nil, nil
	}
	return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}

func (t *httpServerTemplate) Prepare() error {
	tags := util.FetchTags(t.Info.Iface.Docs, TagMark+ForceTag)
	if util.IsInStringSlice("http", tags) || util.IsInStringSlice("http-server", tags) {
		t.Info.Force = true
	}
	return nil
}

// Render http server constructor.
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package transporthttp
//
//		import (
//			svc "github.com/devimteam/microgen/example/svc"
//			http2 "github.com/devimteam/microgen/example/svc/transport/converter/http"
//			http "github.com/go-kit/kit/transport/http"
//			http1 "net/http"
//		)
//
//		func NewHTTPHandler(endpoints *svc.Endpoints, opts ...http.ServerOption) http1.Handler {
//			handler := http1.NewServeMux()
//			handler.Handle("/test_case", http.NewServer(
//				endpoints.TestCaseEndpoint,
//				http2.DecodeHTTPTestCaseRequest,
//				http2.EncodeHTTPTestCaseResponse,
//				opts...))
//			handler.Handle("/empty_req", http.NewServer(
//				endpoints.EmptyReqEndpoint,
//				http2.DecodeHTTPEmptyReqRequest,
//				http2.EncodeHTTPEmptyReqResponse,
//				opts...))
//			handler.Handle("/empty_resp", http.NewServer(
//				endpoints.EmptyRespEndpoint,
//				http2.DecodeHTTPEmptyRespRequest,
//				http2.EncodeHTTPEmptyRespResponse,
//				opts...))
//			return handler
//		}
//
func (t *httpServerTemplate) Render() write_strategy.Renderer {
	f := NewFile("transporthttp")
	f.PackageComment(FileHeader)
	f.PackageComment(`Please, do not edit.`)

	f.Func().Id("NewHTTPHandler").Params(
		Id("endpoints").Op("*").Qual(t.Info.ServiceImportPath, "Endpoints"),
		Id("opts").Op("...").Qual(PackagePathGoKitTransportHTTP, "ServerOption"),
	).Params(
		Qual(PackagePathHttp, "Handler"),
	).BlockFunc(func(g *Group) {
		g.Id("handler").Op(":=").Qual(PackagePathHttp, "NewServeMux").Call()
		for _, fn := range t.Info.Iface.Methods {
			g.Id("handler").Dot("Handle").Call(
				Lit("/"+util.ToURLSnakeCase(fn.Name)),
				Qual(PackagePathGoKitTransportHTTP, "NewServer").Call(
					Line().Id("endpoints").Dot(endpointStructName(fn.Name)),
					Line().Qual(pathToHttpConverter(t.Info.ServiceImportPath), httpDecodeRequestName(fn)),
					Line().Qual(pathToHttpConverter(t.Info.ServiceImportPath), httpEncodeResponseName(fn)),
					Line().Id("opts").Op("..."),
				),
			)
		}
		g.Return(Id("handler"))
	})

	return f
}

func pathToHttpConverter(servicePath string) string {
	return filepath.Join(servicePath, "transport/converter/http")
}
