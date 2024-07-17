// Code generated by protoc-gen-protorpc. DO NOT EDIT.
//
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-plugin
// plugin: https://github.com/chai2010/protorpc/tree/master/protoc-gen-protorpc
//
// source: gpyrpc.proto

package gpyrpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/chai2010/protorpc"
	"github.com/golang/protobuf/proto"
)

var (
	_ = fmt.Sprint
	_ = io.Reader(nil)
	_ = log.Print
	_ = net.Addr(nil)
	_ = rpc.Call{}
	_ = time.Second

	_ = proto.String
	_ = protorpc.Dial
)

type PROTORPC_BuiltinService interface {
	Ping(in *Ping_Args, out *Ping_Result) error
	ListMethod(in *ListMethod_Args, out *ListMethod_Result) error
}

// PROTORPC_AcceptBuiltinServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func PROTORPC_AcceptBuiltinServiceClient(lis net.Listener, x PROTORPC_BuiltinService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_RegisterBuiltinService publish the given PROTORPC_BuiltinService implementation on the server.
func PROTORPC_RegisterBuiltinService(srv *rpc.Server, x PROTORPC_BuiltinService) error {
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		return err
	}
	return nil
}

// PROTORPC_NewBuiltinServiceServer returns a new PROTORPC_BuiltinService Server.
func PROTORPC_NewBuiltinServiceServer(x PROTORPC_BuiltinService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// PROTORPC_ListenAndServeBuiltinService listen announces on the local network address laddr
// and serves the given BuiltinService implementation.
func PROTORPC_ListenAndServeBuiltinService(network, addr string, x PROTORPC_BuiltinService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_ServeBuiltinService serves the given PROTORPC_BuiltinService implementation.
func PROTORPC_ServeBuiltinService(conn io.ReadWriteCloser, x PROTORPC_BuiltinService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("BuiltinService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeCodec(protorpc.NewServerCodec(conn))
}

type PROTORPC_BuiltinServiceClient struct {
	*rpc.Client
}

// PROTORPC_NewBuiltinServiceClient returns a PROTORPC_BuiltinService stub to handle
// requests to the set of PROTORPC_BuiltinService at the other end of the connection.
func PROTORPC_NewBuiltinServiceClient(conn io.ReadWriteCloser) *PROTORPC_BuiltinServiceClient {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &PROTORPC_BuiltinServiceClient{c}
}

func (c *PROTORPC_BuiltinServiceClient) Ping(in *Ping_Args) (out *Ping_Result, err error) {
	if in == nil {
		in = new(Ping_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Ping_Result)
	if err = c.Call("BuiltinService.Ping", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_BuiltinServiceClient) AsyncPing(in *Ping_Args, out *Ping_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Ping_Args)
	}
	return c.Go(
		"BuiltinService.Ping",
		in, out,
		done,
	)
}

func (c *PROTORPC_BuiltinServiceClient) ListMethod(in *ListMethod_Args) (out *ListMethod_Result, err error) {
	if in == nil {
		in = new(ListMethod_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListMethod_Result)
	if err = c.Call("BuiltinService.ListMethod", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_BuiltinServiceClient) AsyncListMethod(in *ListMethod_Args, out *ListMethod_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListMethod_Args)
	}
	return c.Go(
		"BuiltinService.ListMethod",
		in, out,
		done,
	)
}

// PROTORPC_DialBuiltinService connects to an PROTORPC_BuiltinService at the specified network address.
func PROTORPC_DialBuiltinService(network, addr string) (*PROTORPC_BuiltinServiceClient, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_BuiltinServiceClient{c}, nil
}

// PROTORPC_DialBuiltinServiceTimeout connects to an PROTORPC_BuiltinService at the specified network address.
func PROTORPC_DialBuiltinServiceTimeout(network, addr string, timeout time.Duration) (*PROTORPC_BuiltinServiceClient, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_BuiltinServiceClient{c}, nil
}

type PROTORPC_KclvmService interface {
	Ping(in *Ping_Args, out *Ping_Result) error
	ExecProgram(in *ExecProgram_Args, out *ExecProgram_Result) error
	BuildProgram(in *BuildProgram_Args, out *BuildProgram_Result) error
	ExecArtifact(in *ExecArtifact_Args, out *ExecProgram_Result) error
	ParseProgram(in *ParseProgram_Args, out *ParseProgram_Result) error
	ListOptions(in *ParseProgram_Args, out *ListOptions_Result) error
	ListVariables(in *ListVariables_Args, out *ListVariables_Result) error
	LoadPackage(in *LoadPackage_Args, out *LoadPackage_Result) error
	FormatCode(in *FormatCode_Args, out *FormatCode_Result) error
	FormatPath(in *FormatPath_Args, out *FormatPath_Result) error
	LintPath(in *LintPath_Args, out *LintPath_Result) error
	OverrideFile(in *OverrideFile_Args, out *OverrideFile_Result) error
	GetSchemaTypeMapping(in *GetSchemaTypeMapping_Args, out *GetSchemaTypeMapping_Result) error
	ValidateCode(in *ValidateCode_Args, out *ValidateCode_Result) error
	ListDepFiles(in *ListDepFiles_Args, out *ListDepFiles_Result) error
	LoadSettingsFiles(in *LoadSettingsFiles_Args, out *LoadSettingsFiles_Result) error
	Rename(in *Rename_Args, out *Rename_Result) error
	RenameCode(in *RenameCode_Args, out *RenameCode_Result) error
	Test(in *Test_Args, out *Test_Result) error
	UpdateDependencies(in *UpdateDependencies_Args, out *UpdateDependencies_Result) error
	GetVersion(in *GetVersion_Args, out *GetVersion_Result) error
}

// PROTORPC_AcceptKclvmServiceClient accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func PROTORPC_AcceptKclvmServiceClient(lis net.Listener, x PROTORPC_KclvmService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_RegisterKclvmService publish the given PROTORPC_KclvmService implementation on the server.
func PROTORPC_RegisterKclvmService(srv *rpc.Server, x PROTORPC_KclvmService) error {
	if err := srv.RegisterName("KclvmService", x); err != nil {
		return err
	}
	return nil
}

// PROTORPC_NewKclvmServiceServer returns a new PROTORPC_KclvmService Server.
func PROTORPC_NewKclvmServiceServer(x PROTORPC_KclvmService) *rpc.Server {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}
	return srv
}

// PROTORPC_ListenAndServeKclvmService listen announces on the local network address laddr
// and serves the given KclvmService implementation.
func PROTORPC_ListenAndServeKclvmService(network, addr string, x PROTORPC_KclvmService) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}

// PROTORPC_ServeKclvmService serves the given PROTORPC_KclvmService implementation.
func PROTORPC_ServeKclvmService(conn io.ReadWriteCloser, x PROTORPC_KclvmService) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("KclvmService", x); err != nil {
		log.Fatal(err)
	}
	srv.ServeCodec(protorpc.NewServerCodec(conn))
}

type PROTORPC_KclvmServiceClient struct {
	*rpc.Client
}

// PROTORPC_NewKclvmServiceClient returns a PROTORPC_KclvmService stub to handle
// requests to the set of PROTORPC_KclvmService at the other end of the connection.
func PROTORPC_NewKclvmServiceClient(conn io.ReadWriteCloser) *PROTORPC_KclvmServiceClient {
	c := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
	return &PROTORPC_KclvmServiceClient{c}
}

func (c *PROTORPC_KclvmServiceClient) Ping(in *Ping_Args) (out *Ping_Result, err error) {
	if in == nil {
		in = new(Ping_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Ping_Result)
	if err = c.Call("KclvmService.Ping", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncPing(in *Ping_Args, out *Ping_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Ping_Args)
	}
	return c.Go(
		"KclvmService.Ping",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ExecProgram(in *ExecProgram_Args) (out *ExecProgram_Result, err error) {
	if in == nil {
		in = new(ExecProgram_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ExecProgram_Result)
	if err = c.Call("KclvmService.ExecProgram", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncExecProgram(in *ExecProgram_Args, out *ExecProgram_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ExecProgram_Args)
	}
	return c.Go(
		"KclvmService.ExecProgram",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) BuildProgram(in *BuildProgram_Args) (out *BuildProgram_Result, err error) {
	if in == nil {
		in = new(BuildProgram_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(BuildProgram_Result)
	if err = c.Call("KclvmService.BuildProgram", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncBuildProgram(in *BuildProgram_Args, out *BuildProgram_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(BuildProgram_Args)
	}
	return c.Go(
		"KclvmService.BuildProgram",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ExecArtifact(in *ExecArtifact_Args) (out *ExecProgram_Result, err error) {
	if in == nil {
		in = new(ExecArtifact_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ExecProgram_Result)
	if err = c.Call("KclvmService.ExecArtifact", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncExecArtifact(in *ExecArtifact_Args, out *ExecProgram_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ExecArtifact_Args)
	}
	return c.Go(
		"KclvmService.ExecArtifact",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ParseFile(in *ParseFile_Args) (out *ParseFile_Result, err error) {
	if in == nil {
		in = new(ParseFile_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ParseFile_Result)
	if err = c.Call("KclvmService.ParseFile", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncParseFile(in *ParseFile_Args, out *ParseFile_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseFile_Args)
	}
	return c.Go(
		"KclvmService.ParseFile",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ParseProgram(in *ParseProgram_Args) (out *ParseProgram_Result, err error) {
	if in == nil {
		in = new(ParseProgram_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ParseProgram_Result)
	if err = c.Call("KclvmService.ParseProgram", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncParseProgram(in *ParseProgram_Args, out *ParseProgram_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseProgram_Args)
	}
	return c.Go(
		"KclvmService.ParseProgram",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListOptions(in *ParseProgram_Args) (out *ListOptions_Result, err error) {
	if in == nil {
		in = new(ParseProgram_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListOptions_Result)
	if err = c.Call("KclvmService.ListOptions", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListOptions(in *ParseProgram_Args, out *ListOptions_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ParseProgram_Args)
	}
	return c.Go(
		"KclvmService.ListOptions",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListVariables(in *ListVariables_Args) (out *ListVariables_Result, err error) {
	if in == nil {
		in = new(ListVariables_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListVariables_Result)
	if err = c.Call("KclvmService.ListVariables", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListVariables(in *ListVariables_Args, out *ListVariables_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListVariables_Args)
	}
	return c.Go(
		"KclvmService.ListVariables",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) LoadPackage(in *LoadPackage_Args) (out *LoadPackage_Result, err error) {
	if in == nil {
		in = new(LoadPackage_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(LoadPackage_Result)
	if err = c.Call("KclvmService.LoadPackage", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncLoadPackage(in *LoadPackage_Args, out *LoadPackage_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(LoadPackage_Args)
	}
	return c.Go(
		"KclvmService.LoadPackage",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) FormatCode(in *FormatCode_Args) (out *FormatCode_Result, err error) {
	if in == nil {
		in = new(FormatCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(FormatCode_Result)
	if err = c.Call("KclvmService.FormatCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncFormatCode(in *FormatCode_Args, out *FormatCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(FormatCode_Args)
	}
	return c.Go(
		"KclvmService.FormatCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) FormatPath(in *FormatPath_Args) (out *FormatPath_Result, err error) {
	if in == nil {
		in = new(FormatPath_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(FormatPath_Result)
	if err = c.Call("KclvmService.FormatPath", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncFormatPath(in *FormatPath_Args, out *FormatPath_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(FormatPath_Args)
	}
	return c.Go(
		"KclvmService.FormatPath",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) LintPath(in *LintPath_Args) (out *LintPath_Result, err error) {
	if in == nil {
		in = new(LintPath_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(LintPath_Result)
	if err = c.Call("KclvmService.LintPath", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncLintPath(in *LintPath_Args, out *LintPath_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(LintPath_Args)
	}
	return c.Go(
		"KclvmService.LintPath",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) OverrideFile(in *OverrideFile_Args) (out *OverrideFile_Result, err error) {
	if in == nil {
		in = new(OverrideFile_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(OverrideFile_Result)
	if err = c.Call("KclvmService.OverrideFile", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncOverrideFile(in *OverrideFile_Args, out *OverrideFile_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(OverrideFile_Args)
	}
	return c.Go(
		"KclvmService.OverrideFile",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) GetSchemaTypeMapping(in *GetSchemaTypeMapping_Args) (out *GetSchemaTypeMapping_Result, err error) {
	if in == nil {
		in = new(GetSchemaTypeMapping_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(GetSchemaTypeMapping_Result)
	if err = c.Call("KclvmService.GetSchemaTypeMapping", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncGetSchemaTypeMapping(in *GetSchemaTypeMapping_Args, out *GetSchemaTypeMapping_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GetSchemaTypeMapping_Args)
	}
	return c.Go(
		"KclvmService.GetSchemaTypeMapping",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ValidateCode(in *ValidateCode_Args) (out *ValidateCode_Result, err error) {
	if in == nil {
		in = new(ValidateCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ValidateCode_Result)
	if err = c.Call("KclvmService.ValidateCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncValidateCode(in *ValidateCode_Args, out *ValidateCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ValidateCode_Args)
	}
	return c.Go(
		"KclvmService.ValidateCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) ListDepFiles(in *ListDepFiles_Args) (out *ListDepFiles_Result, err error) {
	if in == nil {
		in = new(ListDepFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(ListDepFiles_Result)
	if err = c.Call("KclvmService.ListDepFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncListDepFiles(in *ListDepFiles_Args, out *ListDepFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(ListDepFiles_Args)
	}
	return c.Go(
		"KclvmService.ListDepFiles",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) LoadSettingsFiles(in *LoadSettingsFiles_Args) (out *LoadSettingsFiles_Result, err error) {
	if in == nil {
		in = new(LoadSettingsFiles_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(LoadSettingsFiles_Result)
	if err = c.Call("KclvmService.LoadSettingsFiles", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncLoadSettingsFiles(in *LoadSettingsFiles_Args, out *LoadSettingsFiles_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(LoadSettingsFiles_Args)
	}
	return c.Go(
		"KclvmService.LoadSettingsFiles",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) Rename(in *Rename_Args) (out *Rename_Result, err error) {
	if in == nil {
		in = new(Rename_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Rename_Result)
	if err = c.Call("KclvmService.Rename", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncRename(in *Rename_Args, out *Rename_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Rename_Args)
	}
	return c.Go(
		"KclvmService.RenameCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) RenameCode(in *RenameCode_Args) (out *RenameCode_Result, err error) {
	if in == nil {
		in = new(RenameCode_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(RenameCode_Result)
	if err = c.Call("KclvmService.RenameCode", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncRenameCode(in *RenameCode_Args, out *RenameCode_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(RenameCode_Args)
	}
	return c.Go(
		"KclvmService.RenameCode",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) Test(in *Test_Args) (out *Test_Result, err error) {
	if in == nil {
		in = new(Test_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(Test_Result)
	if err = c.Call("KclvmService.Test", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncTest(in *Test_Args, out *Test_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(Test_Args)
	}
	return c.Go(
		"KclvmService.Test",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) UpdateDependencies(in *UpdateDependencies_Args) (out *UpdateDependencies_Result, err error) {
	if in == nil {
		in = new(UpdateDependencies_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(UpdateDependencies_Result)
	if err = c.Call("KclvmService.UpdateDependencies", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncUpdateDependencies(in *UpdateDependencies_Args, out *UpdateDependencies_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(UpdateDependencies_Args)
	}
	return c.Go(
		"KclvmService.UpdateDependencies",
		in, out,
		done,
	)
}

func (c *PROTORPC_KclvmServiceClient) GetVersion(in *GetVersion_Args) (out *GetVersion_Result, err error) {
	if in == nil {
		in = new(GetVersion_Args)
	}

	type Validator interface {
		Validate() error
	}
	if x, ok := proto.Message(in).(Validator); ok {
		if err := x.Validate(); err != nil {
			return nil, err
		}
	}

	out = new(GetVersion_Result)
	if err = c.Call("KclvmService.GetVersion", in, out); err != nil {
		return nil, err
	}

	if x, ok := proto.Message(out).(Validator); ok {
		if err := x.Validate(); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *PROTORPC_KclvmServiceClient) AsyncGetVersion(in *GetVersion_Args, out *GetVersion_Result, done chan *rpc.Call) *rpc.Call {
	if in == nil {
		in = new(GetVersion_Args)
	}
	return c.Go(
		"KclvmService.GetVersion",
		in, out,
		done,
	)
}

// PROTORPC_DialKclvmService connects to an PROTORPC_KclvmService at the specified network address.
func PROTORPC_DialKclvmService(network, addr string) (*PROTORPC_KclvmServiceClient, error) {
	c, err := protorpc.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_KclvmServiceClient{c}, nil
}

// PROTORPC_DialKclvmServiceTimeout connects to an PROTORPC_KclvmService at the specified network address.
func PROTORPC_DialKclvmServiceTimeout(network, addr string, timeout time.Duration) (*PROTORPC_KclvmServiceClient, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}
	return &PROTORPC_KclvmServiceClient{c}, nil
}
