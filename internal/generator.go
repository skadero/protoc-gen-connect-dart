package internal

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed client.tmpl
var apiTemplate string

type Service struct {
	Name    string
	Package string
	Methods []ServiceMethod
}

type ServiceMethod struct {
	Name       string
	Path       string
	InputArg   string
	InputType  string
	OutputType string
}

func NewAPIContext(filename string) APIContext {
	ctx := APIContext{FileName: filename}

	return ctx
}

type APIContext struct {
	Services []*Service
	Imports  []Import
	FileName string
}

type Import struct {
	Path string
}

func dartPbFilename(name string) string {
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		base := path.Base(name)
		name = base[:len(base)-len(path.Ext(base))]
	}

	name += ".pb.dart"
	return name
}

func connectFilename(fullPath string) string {
	name := ""
	if ext := path.Ext(fullPath); ext == ".proto" || ext == ".protodevel" {
		base := path.Base(fullPath)
		name = base[:len(base)-len(path.Ext(base))]
	}
	name += ".connect.dart"
	return path.Join(path.Dir(fullPath), name)
}

func (ctx *APIContext) ApplyImports(f *protogen.File) {
	var deps []Import

	if len(ctx.Services) > 0 {
		deps = append(
			deps,
			Import{"dart:async"},
			Import{"package:ska_connect/ska_connect.dart"},
		)
	}
	deps = append(deps, Import{"dart:convert"})

	baseName := filepath.Base(f.Proto.GetName())
	baseName = strings.TrimSuffix(baseName, ".proto")

	deps = append(deps, Import{fmt.Sprintf("./%s.pb.dart", baseName)})

	for _, dep := range f.Proto.Dependency {
		if dep == "google/protobuf/timestamp.proto" {
			continue
		}
		importPath := path.Dir(dep)
		sourceDir := path.Dir(f.Proto.GetName())
		sourceComponents := strings.Split(sourceDir, fmt.Sprintf("%c", os.PathSeparator))
		distanceFromRoot := len(sourceComponents)
		for _, pathComponent := range sourceComponents {
			if strings.HasPrefix(importPath, pathComponent) {
				importPath = strings.TrimPrefix(importPath, pathComponent)
				distanceFromRoot--
			}
		}
		deps = append(deps, Import{fullImportPath(dartPbFilename(dep), importPath, distanceFromRoot)})
	}
	ctx.Imports = deps
}

func fullImportPath(fileName string, importPath string, distanceFromRoot int) string {
	fullPath := fileName
	fullPath = path.Join(importPath, fullPath)
	if distanceFromRoot > 0 {
		for i := 0; i < distanceFromRoot; i++ {
			fullPath = path.Join("..", fullPath)
		}
	}

	return fullPath
}

func CreateClientAPI(p *protogen.Plugin, f *protogen.File) error {
	ctx := NewAPIContext(*f.Proto.Name)
	for _, s := range f.Services {
		service := &Service{
			Name:    string(s.Desc.Name()),
			Package: f.Proto.GetPackage(),
		}

		for _, m := range s.Methods {
			methodPath := string(m.Desc.Name())
			methodName := strings.ToLower(methodPath[0:1]) + methodPath[1:]
			in := string(m.Input.Desc.Name())
			arg := strings.ToLower(in[0:1]) + in[1:]
			service.Methods = append(service.Methods, ServiceMethod{
				Name:       methodName,
				Path:       methodPath,
				InputArg:   arg,
				InputType:  in,
				OutputType: string(m.Output.Desc.Name()),
			})
		}

		ctx.Services = append(ctx.Services, service)
	}

	ctx.ApplyImports(f)

	t, err := template.New("client_api").Parse(apiTemplate)
	if err != nil {
		return err
	}

	b := bytes.NewBufferString("")
	err = t.Execute(b, ctx)
	if err != nil {
		return err
	}

	ff := p.NewGeneratedFile(connectFilename(f.Proto.GetName()), "")
	_, err = ff.Write(b.Bytes())

	return err
}
