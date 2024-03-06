package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	sprig "github.com/go-task/slim-sprig/v3"
	v1 "github.com/metal-stack/api/go/api/v1"
	"github.com/metal-stack/api/go/permissions"
	"github.com/metal-stack/api/go/tests/protoparser"

	_ "embed"
)

const (
	// serverReflectionInfo1alpha1 is always allowed to access to get a list of exposed services for example with grpcurl
	serverReflectionInfov1alpha1 = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
	// serverReflectionInfo is always allowed to access to get a list of exposed services for example with grpcurl
	serverReflectionInfo = "/grpc.reflection.v1.ServerReflection/ServerReflectionInfo"
)

var (
	//go:embed go_servicepermissions.tpl
	servicePermissionsTpl string
	//go:embed go_mock_client.tpl
	mockClientTpl string
	//go:embed go_client.tpl
	clientTpl string
)

type api struct {
	Name     string
	Services []string
	Path     string
}

func main() {
	fmt.Println("generating service permissions")

	perms, err := servicePermissions("../proto")
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/permissions/servicepermissions.go", servicePermissionsTpl, perms)
	if err != nil {
		panic(err)
	}

	fmt.Println("generating clients")

	svcs, err := svcs("../proto")
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/client/client.go", clientTpl, svcs)
	if err != nil {
		panic(err)
	}

	err = writeTemplate("../go/tests/mock_clients.go", mockClientTpl, svcs)
	if err != nil {
		panic(err)
	}
}

func servicePermissions(root string) (*permissions.ServicePermissions, error) {
	var (
		walk = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
		roles = permissions.Roles{
			Admin:   &permissions.Admin{},
			Tenant:  &permissions.Tenant{},
			Project: &permissions.Project{},
		}
		methods    = permissions.Methods{}
		visibility = permissions.Visibility{
			Public: map[string]bool{
				// Allow service reflection to list available methods
				serverReflectionInfov1alpha1: true,
				serverReflectionInfo:         true,
			},
			Private: map[string]bool{},
			Self:    map[string]bool{},
		}
		chargeable = permissions.Chargeable{}
		auditable  = permissions.Auditable{}
		services   = []string{}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		filename := filename
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return nil, err
		}
		for _, serviceDesc := range fd.GetService() {
			serviceDesc := serviceDesc
			services = append(services, fmt.Sprintf("%s.%s", fd.GetPackage(), serviceDesc.GetName()))
			for _, method := range serviceDesc.GetMethod() {
				method := method
				methodName := fmt.Sprintf("/%s.%s/%s", fd.GetPackage(), serviceDesc.GetName(), method.GetName())
				methodOpts := method.GetOptions().GetUninterpretedOption()
				for _, methodOpt := range methodOpts {
					methodOpt := methodOpt
					for _, namePart := range methodOpt.GetName() {
						namePart := namePart
						if !namePart.GetIsExtension() {
							continue
						}
						auditable[methodName] = true
						// Tenant
						switch methodOpt.GetIdentifierValue() {
						case v1.TenantRole_TENANT_ROLE_OWNER.String():
							roles.Tenant.Owner = append(roles.Tenant.Owner, methodName)
						case v1.TenantRole_TENANT_ROLE_EDITOR.String():
							roles.Tenant.Editor = append(roles.Tenant.Editor, methodName)
						case v1.TenantRole_TENANT_ROLE_VIEWER.String():
							roles.Tenant.Viewer = append(roles.Tenant.Viewer, methodName)
						case v1.TenantRole_TENANT_ROLE_UNSPECIFIED.String():
							// noop
						// Project
						case v1.ProjectRole_PROJECT_ROLE_OWNER.String():
							roles.Project.Owner = append(roles.Project.Owner, methodName)
						case v1.ProjectRole_PROJECT_ROLE_EDITOR.String():
							roles.Project.Editor = append(roles.Project.Editor, methodName)
						case v1.ProjectRole_PROJECT_ROLE_VIEWER.String():
							roles.Project.Viewer = append(roles.Project.Viewer, methodName)
						case v1.ProjectRole_PROJECT_ROLE_UNSPECIFIED.String():
							// noop
						// Admin
						case v1.AdminRole_ADMIN_ROLE_EDITOR.String():
							roles.Admin.Editor = append(roles.Admin.Editor, methodName)
						case v1.AdminRole_ADMIN_ROLE_VIEWER.String():
							roles.Admin.Viewer = append(roles.Admin.Viewer, methodName)
						case v1.AdminRole_ADMIN_ROLE_UNSPECIFIED.String():
							// noop
						// Visibility
						case v1.Visibility_VISIBILITY_PUBLIC.String():
							visibility.Public[methodName] = true
						case v1.Visibility_VISIBILITY_PRIVATE.String():
							visibility.Private[methodName] = true
						case v1.Visibility_VISIBILITY_SELF.String():
							visibility.Self[methodName] = true
						case v1.Visibility_VISIBILITY_UNSPECIFIED.String():
							// noop
						// Auditable
						case v1.Auditing_AUDITING_EXCLUDED.String():
							auditable[methodName] = false
						case v1.Auditing_AUDITING_INCLUDED.String():
							auditable[methodName] = true
						case v1.Auditing_AUDITING_UNSPECIFIED.String():
							auditable[methodName] = true
						// noop
						default:
							return nil, fmt.Errorf("unknown method identifier value detected:%s", methodOpt.GetIdentifierValue())

						}
						methods[methodName] = true
					}
				}
			}
		}
	}
	slices.Sort(services)
	sp := &permissions.ServicePermissions{
		Roles:      roles,
		Methods:    methods,
		Visibility: visibility,
		Chargeable: chargeable,
		Auditable:  auditable,
		Services:   services,
	}

	return sp, nil
}

func svcs(root string) (map[string]api, error) {
	var (
		result = map[string]api{}
		walk   = func(root string) ([]string, error) {
			var files []string
			err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.HasSuffix(info.Name(), ".proto") {
					files = append(files, path)
				}
				return nil
			})
			return files, err
		}
	)

	files, err := walk(root)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		fd, err := protoparser.Parse(filename)
		if err != nil {
			return nil, err
		}
		n := strings.ReplaceAll(fd.GetPackage(), ".", "")
		a, ok := result[n]
		if !ok {
			a = api{
				Name: n,
				Path: path.Dir(strings.TrimPrefix(filename, root)),
			}
		}
		for _, serviceDesc := range fd.GetService() {
			a.Services = append(a.Services, serviceDesc.GetName())
		}
		result[n] = a
	}

	return result, nil
}

func writeTemplate(dest, text string, data any) error {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(text)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fmt.Println("wrote " + dest)

	return os.WriteFile(dest, p, 0755) // nolint:gosec
}
