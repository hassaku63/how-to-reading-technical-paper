package resources

import (
	"context"
	"io/fs"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	appassets "github.com/hassaku63/how-to-reading-technical-paper/assets"
)

type assetSpec struct {
	URI      string
	Name     string
	Title    string
	Desc     string
	Path     string // path within embedded FS
	MIMEType string
}

func registerAsset(s *mcp.Server, a assetSpec) {
	res := &mcp.Resource{
		Name:        a.Name,
		Title:       a.Title,
		Description: a.Desc,
		URI:         a.URI,
		MIMEType:    a.MIMEType,
	}
	handler := func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		if req.Params.URI != a.URI {
			return nil, mcp.ResourceNotFoundError(req.Params.URI)
		}
		data, err := fs.ReadFile(appassets.FS, a.Path)
		if err != nil {
			return nil, err
		}
		return &mcp.ReadResourceResult{Contents: []*mcp.ResourceContents{
			{URI: a.URI, MIMEType: a.MIMEType, Blob: data},
		}}, nil
	}
	s.AddResource(res, handler)
}

// RegisterInitial registers the first resource from docs/mcp-specs for review.
// We add one at a time per the requested review flow.
func RegisterInitial(s *mcp.Server) {
	// 1) templates/screening-checklist (asset-driven, embedded JSON)
	registerAsset(s, assetSpec{
		URI:      "paper-reading://templates/screening-checklist",
		Name:     "templates/screening-checklist",
		Title:    "スクリーニング用チェックリスト",
		Desc:     "論文スクリーニングのための標準化されたチェックリスト",
		Path:     "resources/templates/screening-checklist.json",
		MIMEType: "application/json",
	})
}
