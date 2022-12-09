package template

import (
	"context"

	. "github.com/dave/jennifer/jen"
	mstrings "github.com/valerylobachev/microgen/generator/strings"
	"github.com/valerylobachev/microgen/generator/write_strategy"
	"github.com/vetcher/go-astra/types"
)

type gRPCServerTemplate struct {
	info *GenerationInfo
}

func NewGRPCServerTemplate(info *GenerationInfo) Template {
	return &gRPCServerTemplate{
		info: info,
	}
}

func serverStructName(iface *types.Interface) string {
	return iface.Name + "Server"
}

func privateServerStructName(iface *types.Interface) string {
	return mstrings.ToLower(iface.Name) + "Server"
}

// Render whole grpc server file.
//
//	// This file was automatically generated by "microgen" utility.
//	// DO NOT EDIT.
//	package transportgrpc
//
//	import (
//		svc "github.com/valerylobachev/microgen/examples/svc"
//		protobuf "github.com/valerylobachev/microgen/examples/svc/transport/converter/protobuf"
//		grpc "github.com/go-kit/kit/transport/grpc"
//		stringsvc "gitlab.devim.team/protobuf/stringsvc"
//		context "golang.org/x/net/context"
//	)
//
//	type stringServiceServer struct {
//		count grpc.Handler
//	}
//
//	func NewGRPCServer(endpoints *svc.Endpoints, opts ...grpc.ServerOption) stringsvc.StringServiceServer {
//		return &stringServiceServer{count: grpc.NewServer(
//			endpoints.CountEndpoint,
//			protobuf.DecodeCountRequest,
//			protobuf.EncodeCountResponse,
//			opts...,
//		)}
//	}
//
//	func (s *stringServiceServer) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//		_, resp, err := s.count.ServeGRPC(ctx, req)
//		if err != nil {
//			return nil, err
//		}
//		return resp.(*stringsvc.CountResponse), nil
//	}
func (t *gRPCServerTemplate) Render(ctx context.Context) write_strategy.Renderer {
	f := NewFile("transportgrpc")
	f.ImportAlias(t.info.ProtobufPackageImport, "pb")
	f.ImportAlias(t.info.SourcePackageImport, serviceAlias)
	f.HeaderComment(t.info.FileHeader)
	f.PackageComment(`DO NOT EDIT.`)

	f.Type().Id(privateServerStructName(t.info.Iface)).StructFunc(func(g *Group) {
		for _, method := range t.info.Iface.Methods {
			if !t.info.AllowedMethods[method.Name] {
				continue
			}
			g.Id(mstrings.ToLowerFirst(method.Name)).Qual(PackagePathGoKitTransportGRPC, "Handler")
		}
	}).Line()

	f.Func().Id("NewGRPCServer").
		ParamsFunc(func(p *Group) {
			p.Id("endpoints").Op("*").Qual(t.info.OutputPackageImport+"/transport", EndpointsSetName)
			if Tags(ctx).Has(TracingMiddlewareTag) {
				p.Id("logger").Qual(PackagePathGoKitLog, "Logger")
			}
			if Tags(ctx).Has(TracingMiddlewareTag) {
				p.Id("tracer").Qual(PackagePathOpenTracingGo, "Tracer")
			}
			p.Id("opts").Op("...").Qual(PackagePathGoKitTransportGRPC, "ServerOption")
		}).Params(
		Qual(t.info.ProtobufPackageImport, serverStructName(t.info.Iface)),
	).
		Block(
			Return().Op("&").Id(privateServerStructName(t.info.Iface)).Values(DictFunc(func(g Dict) {
				for _, m := range t.info.Iface.Methods {
					if !t.info.AllowedMethods[m.Name] {
						continue
					}
					g[(&Statement{}).Id(mstrings.ToLowerFirst(m.Name))] = Qual(PackagePathGoKitTransportGRPC, "NewServer").
						Call(
							Line().Id("endpoints").Dot(endpointsStructFieldName(m.Name)),
							Line().Id(decodeRequestName(m)),
							Line().Id(encodeResponseName(m)),
							Line().Add(t.serverOpts(ctx, m)).Op("...").Line(),
						)
				}
			}),
			),
		)
	f.Line()

	for _, signature := range t.info.Iface.Methods {
		if !t.info.AllowedMethods[signature.Name] {
			continue
		}
		f.Add(t.grpcServerFunc(signature, t.info.Iface)).Line()
	}

	return f
}

func (gRPCServerTemplate) DefaultPath() string {
	return filenameBuilder(PathTransport, "grpc", "server")
}

func (t *gRPCServerTemplate) Prepare(ctx context.Context) error {
	if t.info.ProtobufPackageImport == "" {
		return ErrProtobufEmpty
	}
	return nil
}

func (t *gRPCServerTemplate) ChooseStrategy(ctx context.Context) (write_strategy.Strategy, error) {
	return write_strategy.NewCreateFileStrategy(t.info.OutputFilePath, t.DefaultPath()), nil
}

// Render service interface method for grpc server.
//
//	func (s *stringServiceServer) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//		_, resp, err := s.count.ServeGRPC(ctx, req)
//		if err != nil {
//			return nil, err
//		}
//		return resp.(*stringsvc.CountResponse), nil
//	}
func (t *gRPCServerTemplate) grpcServerFunc(signature *types.Function, i *types.Interface) *Statement {
	return Func().
		Params(Id(rec(privateServerStructName(i))).Op("*").Id(privateServerStructName(i))).
		Id(signature.Name).
		Call(Id("ctx").Qual(PackagePathNetContext, "Context"), Id("req").Add(t.grpcServerReqStruct(signature))).
		Params(t.grpcServerRespStruct(signature), Error()).
		BlockFunc(t.grpcServerFuncBody(signature, i))
}

// Special case for empty request
// Render
//
//	*empty.Empty
//
// or
//
//	*stringsvc.CountRequest
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
	return Op("*").Qual(t.info.ProtobufPackageImport, requestStructName(fn))
}

// Special case for empty response
// Render
//
//	*empty.Empty
//
// or
//
//	*stringsvc.CountResponse
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
	return Op("*").Qual(t.info.ProtobufPackageImport, responseStructName(fn))
}

// Render service method body for grpc server.
//
//	_, resp, err := s.count.ServeGRPC(ctx, req)
//	if err != nil {
//		return nil, err
//	}
//	return resp.(*stringsvc.CountResponse), nil
func (t *gRPCServerTemplate) grpcServerFuncBody(signature *types.Function, i *types.Interface) func(g *Group) {
	return func(g *Group) {
		g.List(Id("_"), Id("resp"), Err()).
			Op(":=").
			Id(rec(privateServerStructName(i))).Dot(mstrings.ToLowerFirst(signature.Name)).Dot("ServeGRPC").Call(Id("ctx"), Id("req"))

		g.If(Err().Op("!=").Nil()).Block(
			Return().List(Nil(), Err()),
		)

		g.Return().List(Id("resp").Assert(t.grpcServerRespStruct(signature)), Nil())
	}
}

func (t *gRPCServerTemplate) serverOpts(ctx context.Context, fn *types.Function) *Statement {
	s := &Statement{}
	if Tags(ctx).Has(TracingMiddlewareTag) {
		s.Op("append(")
		defer s.Op(")")
	}
	s.Id("opts")
	if Tags(ctx).Has(TracingMiddlewareTag) {
		s.Op(",").Qual(PackagePathGoKitTransportGRPC, "ServerBefore").Call(
			Line().Qual(PackagePathGoKitTracing, "GRPCToContext").Call(Id("tracer"), Lit(fn.Name), Id("logger")),
		)
	}
	return s
}
