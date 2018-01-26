package template

import (
	"errors"

	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/devimteam/microgen/util"
	"github.com/vetcher/godecl/types"
)

var (
	GRPCAddrEmptyError = errors.New("grpc server address is empty")
	ProtobufEmptyError = errors.New("protobuf package is empty")
)

type gRPCClientTemplate struct {
	Info *GenerationInfo
}

func NewGRPCClientTemplate(info *GenerationInfo) Template {
	return &gRPCClientTemplate{
		Info: info.Copy(),
	}
}

func (t *gRPCClientTemplate) grpcConverterPackagePath() string {
	return t.Info.ServiceImportPath + "/transport/converter/protobuf"
}

// Render whole grpc client file.
//
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package transportgrpc
//
//		import (
//			svc "github.com/devimteam/microgen/example/svc"
//			protobuf "github.com/devimteam/microgen/example/svc/transport/converter/protobuf"
//			grpc1 "github.com/go-kit/kit/transport/grpc"
//			stringsvc "gitlab.devim.team/protobuf/stringsvc"
//			grpc "google.golang.org/grpc"
//		)
//
//		func NewGRPCClient(conn *grpc.ClientConn, opts ...grpc1.ClientOption) svc.StringService {
//			return &svc.Endpoints{CountEndpoint: grpc1.NewClient(
//				conn,
//				"devim.string.protobuf.StringService",
//				"Count",
//				protobuf.EncodeCountRequest,
//				protobuf.DecodeCountResponse,
//				stringsvc.CountResponse{},
//				opts...,
//			).Endpoint()}
//		}
//
func (t *gRPCClientTemplate) Render() write_strategy.Renderer {
	f := NewFile("transportgrpc")
	f.PackageComment(t.Info.FileHeader)
	f.PackageComment(`Please, do not edit.`)

	f.Func().Id("NewGRPCClient").
		Params(
			Id("conn").Op("*").Qual(PackagePathGoogleGRPC, "ClientConn"),
			Id("opts").Op("...").Qual(PackagePathGoKitTransportGRPC, "ClientOption"),
		).Qual(t.Info.ServiceImportPath, t.Info.Iface.Name).
		BlockFunc(func(g *Group) {
			g.Return().Op("&").Qual(t.Info.ServiceImportPath, "Endpoints").Values(DictFunc(func(d Dict) {
				for _, m := range t.Info.Iface.Methods {
					d[Id(endpointStructName(m.Name))] = Qual(PackagePathGoKitTransportGRPC, "NewClient").Call(
						Line().Id("conn"),
						Line().Lit(t.Info.GRPCRegAddr),
						Line().Lit(m.Name),
						Line().Qual(pathToConverter(t.Info.ServiceImportPath), requestEncodeName(m)),
						Line().Qual(pathToConverter(t.Info.ServiceImportPath), responseDecodeName(m)),
						Line().Add(t.replyType(m)),
						Line().Id("opts").Op("...").Line(),
					).Dot("Endpoint").Call()
				}
			}))
		})
	return f
}

// Renders reply type argument
// 		stringsvc.CountResponse{}
func (t *gRPCClientTemplate) replyType(signature *types.Function) *Statement {
	results := removeErrorIfLast(signature.Results)
	if len(results) == 0 {
		return Qual(PackagePathEmptyProtobuf, "Empty").Values()
	}
	if len(results) == 1 {
		sp := specialReplyType(results[0].Type)
		if sp != nil {
			return sp
		}
	}
	return Qual(t.Info.ProtobufPackage, responseStructName(signature)).Values()
}

func specialReplyType(p types.Type) *Statement {
	name := types.TypeName(p)
	imp := types.TypeImport(p)
	// *string -> *wrappers.StringValue
	if name != nil && *name == "string" && imp == nil && p.TypeOf() == types.T_Pointer {
		return (&Statement{}).Qual(GolangProtobufWrappers, "StringValue").Values()
	}
	return nil
}

func (gRPCClientTemplate) DefaultPath() string {
	return "./transport/grpc/client.go"
}

func (t *gRPCClientTemplate) Prepare() error {
	if t.Info.GRPCRegAddr == "" {
		return GRPCAddrEmptyError
	}
	if t.Info.ProtobufPackage == "" {
		return ProtobufEmptyError
	}

	tags := util.FetchTags(t.Info.Iface.Docs, TagMark+ForceTag)
	if util.IsInStringSlice("grpc", tags) || util.IsInStringSlice("grpc-client", tags) {
		t.Info.Force = true
	}
	return nil
}

func (t *gRPCClientTemplate) ChooseStrategy() (write_strategy.Strategy, error) {
	if err := util.StatFile(t.Info.AbsOutPath, t.DefaultPath()); !t.Info.Force && err == nil {
		return nil, nil
	}
	return write_strategy.NewCreateFileStrategy(t.Info.AbsOutPath, t.DefaultPath()), nil
}
