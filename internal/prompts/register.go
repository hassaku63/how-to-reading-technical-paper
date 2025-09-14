package prompts

import (
	"context"
	"fmt"
	"io/fs"
	"regexp"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	appassets "github.com/hassaku63/how-to-reading-technical-paper/assets"
)

// simpleTemplate fills {var} placeholders with values from args.
var placeholderRe = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)

func fillTemplate(tmpl string, args map[string]string) string {
	return placeholderRe.ReplaceAllStringFunc(tmpl, func(m string) string {
		key := strings.Trim(m, "{}")
		if v, ok := args[key]; ok {
			return v
		}
		// leave placeholder if missing to make it visible to the caller
		return m
	})
}

func registerPromptFromAsset(s *mcp.Server, p *mcp.Prompt, assetPath string) {
	handler := func(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		if req.Params.Name != p.Name {
			return nil, fmt.Errorf("unknown prompt: %s", req.Params.Name)
		}
		// Load template text
		data, err := fs.ReadFile(appassets.FS, assetPath)
		if err != nil {
			return nil, err
		}
		text := string(data)
		// Fill placeholders
		filled := fillTemplate(text, req.Params.Arguments)
		return &mcp.GetPromptResult{
			Description: p.Description,
			Messages: []*mcp.PromptMessage{
				{Role: "user", Content: &mcp.TextContent{Text: filled}},
			},
		}, nil
	}
	s.AddPrompt(p, handler)
}

// RegisterSurveyReadingFlow registers the survey-reading-flow prompt.
func RegisterSurveyReadingFlow(s *mcp.Server) {
	p := &mcp.Prompt{
		Name:        "survey-reading-flow",
		Title:       "高速サーベイ読解フロー",
		Description: "研究者が大量の論文を効率的にサーベイするための構造化フロー",
		Arguments: []*mcp.PromptArgument{
			{Name: "research_domain", Title: "研究分野", Description: "例：分散システム、機械学習、データベース", Required: true},
			{Name: "survey_goal", Title: "サーベイ目的", Description: "例：最新動向把握、手法比較、研究ギャップ特定", Required: true},
			{Name: "time_budget", Title: "時間予算(分)", Description: "1論文あたりの時間予算（分）", Required: false},
		},
	}
	registerPromptFromAsset(s, p, "prompts/survey-reading-flow.txt")
}

// RegisterImplementationReadingFlow registers the implementation-reading-flow prompt.
func RegisterImplementationReadingFlow(s *mcp.Server) {
	p := &mcp.Prompt{
		Name:        "implementation-reading-flow",
		Title:       "実装向け読解フロー",
		Description: "ソフトウェアエンジニアが論文のアイデアを実装するための読解フロー",
		Arguments: []*mcp.PromptArgument{
			{Name: "skill_goals", Title: "スキル目標", Description: "カンマ区切り推奨：例）分散システム, アルゴリズム最適化, API設計", Required: true},
			{Name: "implementation_timeline", Title: "実装期間", Description: "例：1-3ヶ月", Required: false},
			{Name: "current_level", Title: "現在レベル", Description: "初心者/中級者/上級者", Required: true},
		},
	}
	registerPromptFromAsset(s, p, "prompts/implementation-reading-flow.txt")
}

// RegisterOutputGeneration registers the output-generation prompt.
func RegisterOutputGeneration(s *mcp.Server) {
	p := &mcp.Prompt{
		Name:        "output-generation",
		Title:       "構造化アウトプット生成",
		Description: "読解結果を目的に応じた形式で構造化してアウトプットするフロー",
		Arguments: []*mcp.PromptArgument{
			{Name: "output_format", Title: "出力形式", Description: "研究メモ/技術ブログ/実装計画/プレゼン/レポート", Required: true},
			{Name: "target_audience", Title: "対象読者", Description: "研究者/エンジニア/一般読者など", Required: true},
			{Name: "key_insights", Title: "強調ポイント", Description: "カンマ区切りで列挙", Required: false},
		},
	}
	registerPromptFromAsset(s, p, "prompts/output-generation.txt")
}

// RegisterCriticalAnalysis registers the critical-analysis prompt.
func RegisterCriticalAnalysis(s *mcp.Server) {
	p := &mcp.Prompt{
		Name:        "critical-analysis",
		Title:       "批判的論文分析",
		Description: "論文の主張を批判的に検証し、限界と問題点を特定するフロー",
		Arguments: []*mcp.PromptArgument{
			{Name: "analysis_depth", Title: "分析の深さ", Description: "浅い/中程度/深い", Required: false},
			{Name: "focus_area", Title: "重点分析領域", Description: "実験設計/理論的基盤/実装詳細 など", Required: false},
		},
	}
	registerPromptFromAsset(s, p, "prompts/critical-analysis.txt")
}
