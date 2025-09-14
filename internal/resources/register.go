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

// RegisterTemplatesReadingNotes registers the reading-notes template (Markdown).
func RegisterTemplatesReadingNotes(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://templates/reading-notes",
        Name:     "templates/reading-notes",
        Title:    "読書メモテンプレート",
        Desc:     "サーベイ/実装向けの標準化された読書メモテンプレート",
        Path:     "resources/templates/reading-notes.json",
        MIMEType: "application/json",
    })
}

// RegisterPatternsCommonStructures registers the common paper structure patterns (JSON).
func RegisterPatternsCommonStructures(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://patterns/common-structures",
        Name:     "patterns/common-structures",
        Title:    "論文の一般的構造パターン",
        Desc:     "システム/アルゴリズム/評価論文の構造とキーパート",
        Path:     "resources/patterns/common-structures.json",
        MIMEType: "application/json",
    })
}

// RegisterCriteriaPaperEvaluation registers the paper evaluation criteria (JSON).
func RegisterCriteriaPaperEvaluation(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://criteria/paper-evaluation",
        Name:     "criteria/paper-evaluation",
        Title:    "論文評価基準",
        Desc:     "論文の品質と価値を評価するための基準",
        Path:     "resources/criteria/paper-evaluation.json",
        MIMEType: "application/json",
    })
}

// RegisterMethodologySurveyGuide registers the survey methodology guide (JSON).
func RegisterMethodologySurveyGuide(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://methodology/survey-guide",
        Name:     "methodology/survey-guide",
        Title:    "サーベイ向け読解方法論",
        Desc:     "大量の論文を効率的にサーベイするための3段階アプローチ",
        Path:     "resources/methodology/survey-guide.json",
        MIMEType: "application/json",
    })
}

// RegisterMethodologyImplementationGuide registers the implementation methodology guide (JSON).
func RegisterMethodologyImplementationGuide(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://methodology/implementation-guide",
        Name:     "methodology/implementation-guide",
        Title:    "実装向け読解方法論",
        Desc:     "論文アイデアを実装するための読解手法（段階的アプローチ）",
        Path:     "resources/methodology/implementation-guide.json",
        MIMEType: "application/json",
    })
}

// RegisterActionableSurveySteps registers the actionable steps for survey (JSON).
func RegisterActionableSurveySteps(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://actionable/survey-steps",
        Name:     "actionable/survey-steps",
        Title:    "サーベイ向け実践手順",
        Desc:     "高速サーベイのための段階的手順とテンプレート",
        Path:     "resources/actionable/survey-steps.json",
        MIMEType: "application/json",
    })
}

// RegisterActionableImplementationSteps registers the actionable steps for implementation (JSON).
func RegisterActionableImplementationSteps(s *mcp.Server) {
    registerAsset(s, assetSpec{
        URI:      "paper-reading://actionable/implementation-steps",
        Name:     "actionable/implementation-steps",
        Title:    "実装向け実践手順",
        Desc:     "論文アイデアを実装するための段階的手順とリスク低減策",
        Path:     "resources/actionable/implementation-steps.json",
        MIMEType: "application/json",
    })
}
