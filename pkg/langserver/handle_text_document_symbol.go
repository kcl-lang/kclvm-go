package langserver

import (
	"context"
	"encoding/json"

	"github.com/sourcegraph/jsonrpc2"

	"kusionstack.io/kclvm-go/pkg/service"
	"kusionstack.io/kclvm-go/pkg/spec/gpyrpc"
)

func (h *langHandler) handleTextDocumentSymbol(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (result interface{}, err error) {
	if req.Params == nil {
		return nil, &jsonrpc2.Error{Code: jsonrpc2.CodeInvalidParams}
	}

	var params DocumentSymbolParams
	if err := json.Unmarshal(*req.Params, &params); err != nil {
		return nil, err
	}

	return h.symbol(params.TextDocument.URI)
}
func (h *langHandler) symbol(uri DocumentURI) ([]DocumentSymbol, error) {
	filename, f, err := h.openOrLoadFile(uri)
	if err != nil {
		return nil, err
	}
	client := service.NewKclvmServiceClient()
	resp, err := client.DocumentSymbol(&gpyrpc.DocumentSymbol_Args{
		File: filename,
		Code: f.Text,
	})
	if err != nil {
		return nil, err
	}
	symbolStr := resp.Symbol
	var result []DocumentSymbol
	err = json.Unmarshal([]byte(symbolStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var symbolKindMap = map[string]int{
	"file":          1,
	"module":        2,
	"namespace":     3,
	"package":       4,
	"class":         5,
	"method":        6,
	"property":      7,
	"field":         8,
	"constructor":   9,
	"enum":          10,
	"interface":     11,
	"function":      12,
	"variable":      13,
	"constant":      14,
	"string":        15,
	"number":        16,
	"boolean":       17,
	"array":         18,
	"object":        19,
	"key":           20,
	"null":          21,
	"enummember":    22,
	"struct":        23,
	"event":         24,
	"operator":      25,
	"typeparameter": 26,
}
