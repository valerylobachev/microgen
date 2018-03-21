package template

import (
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/devimteam/microgen/util"
	"github.com/vetcher/godecl/types"
)

type gRPCServerTemplate struct {
	Info    *GenerationInfo
	tracing bool
}

func NewGRPCServerTemplate(info *GenerationInfo) Template {
	return &gRPCServerTemplate{
		Info: info.Copy(),
	}
}

func serverStructName(iface *types.Interface) string {
	return iface.Name + "Server"
}

func privateServerStructName(iface *types.Interface) string {
	return util.ToLower(iface.Name) + "Server"
}

func pathToConverter(servicePath string) string {
	return filepath.Join(servicePath, "transport/converter/protobuf")
}

// Render whole grpc server file.
//
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package transportgrpc
//
//		import (
//			svc "github.com/devimteam/microgen/example/svc"
//			protobuf "github.com/devimteam/microgen/example/svc/transport/converter/protobuf"
//			grpc "github.com/go-kit/kit/transport/grpc"
//			stringsvc "gitlab.devim.team/protobuf/stringsvc"
//			context "golang.org/x/net/context"
//		)
//
//		type stringServiceServer struct {
//			count grpc.Handler
//		}
//
//		func NewGRPCServer(endpoints *svc.Endpoints, opts ...grpc.ServerOption) stringsvc.StringServiceServer {
//			return &stringServiceServer{count: grpc.NewServer(
//				endpoints.CountEndpoint,
//				protobuf.DecodeCountRequest,
//				protobuf.EncodeCountResponse,
//				opts...,
//			)}
//		}
//
//		func (s *stringServiceServer) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//			_, resp, err := s.count.ServeGRPC(ctx, req)
//			if err != nil {
//				return nil, err
//			}
//			return resp.(*stringsvc.CountResponse), nil
//		}
//
func (t *gRPCServerTemplate) Render() write_strategy.Renderer {
	f := NewFile("transportgrpc")
	f.PackageComment(t.Info.FileHeader)
	f.PackageComment(`Please, do not edit.`)

	f.Type().Id(privateServerStructName(t.Info.Iface)).StructFunc(func(g *Group) {
		for _, method := range t.Info.Iface.Methods {
			g.Id(util.ToLowerFirst(method.Name)).Qual(PackagePathGoKitTransportGRPC, "Handler")
		}
	}).Line()

	f.Func().Id("NewGRPCServer").
		ParamsFunc(func(p *Group) {
			p.Id("endpoints").Op("*").Qual(t.Info.ServiceImportPath, "Endpoints")
			if t.tracing {
				p.Id("logger").Qual(PackagePathGoKitLog, "Logger")
			}
			if t.tracing {
				p.Id("tracer").Qual(PackagePathOpenTracingGo, "Tracer")
			}
			p.Id("opts").Op("...").Qual(PackagePathGoKitTransportGRPC, "ServerOption")
		}).Params(
		Qual(t.Info.ProtobufPackage, serverStructName(t.Info.Iface)),
	).
		Block(
			Return().Op("&").Id(privateServerStructName(t.Info.Iface)).Values(DictFunc(func(g Dict) {
				for _, m := range t.Info.Iface.Methods {
					g[(&Statement{}).Id(util.ToLowerFirst(m.Name))] = Qual(PackagePathGoKitTransportGRPC, "NewServer").
						Call(
							Line().Id("endpoints").Dot(endpointStructName(m.Name)),
							Line().Qual(pathToConverter(t.Info.ServiceImportPath), decodeRequestName(m)),
							Line().Qual(pathToConverter(t.Info.ServiceImportPath), encodeResponseName(m)),
							Line().Add(t.serverOpts(m)).Op("...").Line(),
						)
				}
			}),
			),
		)
	f.Line()

	for _, signature := range t.Info.Iface.Methods {
		f.Line()
		f.Add(t.grpcServerFunc(signature, t.Info.Iface)).Line()
	}

	return f
}

func (gRPCServerTemplate) DefaultPath() string {
	return "./transport/grpc/server.go"
}

func (t *gRPCServerTemplate) Prepare() error {
	if t.Info.ProtobufPackage == "" {
		return ProtobufEmptyError
	}

	tags := util.FetchTags(t.Info.Iface.Docs, TagMark+ForceTag)
	if util.IsInStringSlice("grpc", tags) || util.IsInStringSlice("grpc-server", tags) {
		t.Info.Force = true
	}
	tags = util.FetchTags(t.Info.Iface.Docs, TagMark+MicrogenMainTag)
	for _, tag := range tags {
		switch tag {
		case TracingTag:
			t.tracing = true
		}
	}
	return nil
}

func (t *gRPCServerTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	if err := util.StatFile(t.Info.AbsOutPath, t.DefaultPath()); !t.Info.Force && err == nil {
		return nil, nil
	}
	return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}

// Render service interface method for grpc server.
//
//		func (s *stringServiceServer) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//			_, resp, err := s.count.ServeGRPC(ctx, req)
//			if err != nil {
//				return nil, err
//			}
//			return resp.(*stringsvc.CountResponse), nil
//		}
//
func (t *gRPCServerTemplate) grpcServerFunc(signature *types.Function, i *types.Interface) *Statement {
	return Func().
		Params(Id(util.LastUpperOrFirst(privateServerStructName(i))).Op("*").Id(privateServerStructName(i))).
		Id(signature.Name).
		Call(Id("ctx").Qual(PackagePathNetContext, "Context"), Id("req").Add(t.grpcServerReqStruct(signature))).
		Params(t.grpcServerRespStruct(signature), Error()).
		BlockFunc(t.grpcServerFuncBody(signature, i))
}

// Special case for empty request
// Render
//		*empty.Empty
// or
//		*stringsvc.CountRequest
func (t *gRPCServerTemplate) grpcServerReqStruct(fn *types.Function) *Statement {
	args := RemoveContextIfFirst(fn.Args)
	if len(args) == 0 {
		return Op("*").Qual(PackagePathEmptyProtobuf, "Empty")
	}
	if len(args) == 1 {
		sp := specialTypeConverter(args[0].Type)
		if sp != nil {
			return sp
		}
	}
	return Op("*").Qual(t.Info.ProtobufPackage, requestStructName(fn))
}

// Special case for empty response
// Render
//		*empty.Empty
// or
//		*stringsvc.CountResponse
func (t *gRPCServerTemplate) grpcServerRespStruct(fn *types.Function) *Statement {
	results := removeErrorIfLast(fn.Results)
	if len(results) == 0 {
		return Op("*").Qual(PackagePathEmptyProtobuf, "Empty")
	}
	if len(results) == 1 {
		sp := specialTypeConverter(results[0].Type)
		if sp != nil {
			return sp
		}
	}
	return Op("*").Qual(t.Info.ProtobufPackage, responseStructName(fn))
}

// Render service method body for grpc server.
//
//		_, resp, err := s.count.ServeGRPC(ctx, req)
//		if err != nil {
//			return nil, err
//		}
//		return resp.(*stringsvc.CountResponse), nil
//
func (t *gRPCServerTemplate) grpcServerFuncBody(signature *types.Function, i *types.Interface) func(g *Group) {
	return func(g *Group) {
		g.List(Id("_"), Id("resp"), Err()).
			Op(":=").
			Id(util.LastUpperOrFirst(privateServerStructName(i))).Dot(util.ToLowerFirst(signature.Name)).Dot("ServeGRPC").Call(Id("ctx"), Id("req"))

		g.If(Err().Op("!=").Nil()).Block(
			Return().List(Nil(), Err()),
		)

		g.Return().List(Id("resp").Assert(t.grpcServerRespStruct(signature)), Nil())
	}
}

func (t *gRPCServerTemplate) serverOpts(fn *types.Function) *Statement {
	s := &Statement{}
	if t.tracing {
		s.Op("append(")
		defer s.Op(")")
	}
	s.Id("opts")
	if t.tracing {
		s.Op(",").Qual(PackagePathGoKitTransportGRPC, "ServerBefore").Call(
			Line().Qual(PackagePathGoKitTracing, "GRPCToContext").Call(Id("tracer"), Lit(fn.Name), Id("logger")),
		)
	}
	return s
}
