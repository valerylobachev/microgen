package template

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	mstrings "github.com/valerylobachev/microgen/generator/strings"
	"github.com/valerylobachev/microgen/generator/write_strategy"
	"github.com/vetcher/go-astra/types"
)

type stubInterfaceTemplate struct {
	info *GenerationInfo

	alreadyRenderedMethods []string
	isStructExist          bool
	isConstructorExist     bool
}

func NewStubInterfaceTemplate(info *GenerationInfo) Template {
	return &stubInterfaceTemplate{
		info: info,
	}
}

// Renders stub code for service, its methods and constructor, that implements service interface.
//
//	// Generated by "microgen" tool.
//	// Structure stringService implements StringService interface.
//	type stringService struct {
//	}
//
//	func NewStringService() StringService {
//		panic("constructor not provided")
//	}
//
//	func (s *stringService) Count(ctx context.Context, text string, symbol string) (count int, positions []int) {
//		panic("method not provided")
//	}
func (t *stubInterfaceTemplate) Render(ctx context.Context) write_strategy.Renderer {
	f := &Statement{}

	if !t.isStructExist {
		f.Comment(`Generated by "microgen" tool.`).Line().
			Commentf(`Struct %s implements %s interface.`, mstrings.ToLower(t.info.Iface.Name), t.info.Iface.Name).Line().
			Type().Id(mstrings.ToLower(t.info.Iface.Name)).Struct(Line()).Line()
	}

	if !t.isConstructorExist {
		f.Func().Id(constructorName(t.info.Iface)).Params().Id(t.info.Iface.Name).Block(
			Panic(Lit("constructor not provided")).Comment("// TODO: provide constructor"),
		).Line()
	}

	for _, signature := range t.info.Iface.Methods {
		if !mstrings.IsInStringSlice(signature.Name, t.alreadyRenderedMethods) {
			f.Line().Add(methodDefinition(ctx, mstrings.ToLower(t.info.Iface.Name), signature)).Block(
				Panic(Lit("method not provided")).Comment("// TODO: provide method"),
			).Line()
		}
	}
	return f
}

func (stubInterfaceTemplate) DefaultPath() string {
	return "."
}

func (t *stubInterfaceTemplate) Prepare(ctx context.Context) error {
	if err := statFile(t.info.SourceFilePath, t.DefaultPath()); os.IsNotExist(err) {
		fmt.Println("warning:", err)
		return nil
	}
	file, err := parsePackage(filepath.Join(t.info.SourceFilePath, t.DefaultPath()))
	if err != nil {
		return err
	}

	// Remove already provided service methods
	for i := range file.Methods {
		name := types.TypeName(file.Methods[i].Receiver.Type)
		if name != nil && *name == mstrings.ToLowerFirst(t.info.Iface.Name) && types.TypeImport(file.Methods[i].Receiver.Type) == nil {
			t.alreadyRenderedMethods = append(t.alreadyRenderedMethods, file.Methods[i].Name)
		}
	}

	// Remove already provided service structure
	for i := range file.Structures {
		if file.Structures[i].Name == mstrings.ToLowerFirst(t.info.Iface.Name) {
			t.isStructExist = true
			break
		}
	}

	// Remove already provided service constructor
	for i := range file.Functions {
		if file.Functions[i].Name == constructorName(t.info.Iface) {
			t.isConstructorExist = true
			break
		}
	}

	return nil
}

func (t *stubInterfaceTemplate) ChooseStrategy(ctx context.Context) (write_strategy.Strategy, error) {
	return write_strategy.NewAppendToFileStrategy(t.info.SourceFilePath, t.DefaultPath()), nil
}

func constructorName(p *types.Interface) string {
	return "New" + p.Name
}
