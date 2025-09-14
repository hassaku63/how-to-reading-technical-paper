# extractKeyInsights Tool

## 基本情報

- **Name**: `extractKeyInsights`
- **Category**: Paper Analysis
- **Purpose**: Abstract/図表から重要情報を自動抽出

## MCP Schema

```json
{
  "name": "extractKeyInsights",
  "description": "Abstract/図表から重要情報を自動抽出",
  "inputSchema": {
    "type": "object",
    "properties": {
      "paper_data": {
        "type": "object",
        "description": "analyzePaper の出力データ"
      },
      "focus_areas": {
        "type": "array",
        "items": {
          "type": "string",
          "enum": ["problem", "method", "results", "limitations"]
        },
        "description": "抽出したい観点"
      }
    },
    "required": ["paper_data", "focus_areas"]
  }
}
```

## パラメータ

- `paper_data`: analyzePaper の出力データ
- `focus_areas`: 抽出したい観点（["problem", "method", "results", "limitations"]）

## 動作詳細

1. Abstractから問題設定、手法、結果を抽出
2. 図表から視覚的な情報を解析
3. 数値データ（改善率、性能指標）を特定
4. 制約・限界を特定

## 返却データ

```json
{
  "problem_statement": "既存の分散システムは一貫性と可用性のトレードオフに課題がある",
  "proposed_method": "Consistent Hashing with Virtual Nodes",
  "key_results": {
    "improvement": "25%の性能向上",
    "metrics": ["スループット", "レイテンシ", "可用性"]
  },
  "limitations": [
    "ノード数が少ない場合の性能劣化",
    "メモリ使用量の増加"
  ],
  "figures_insights": [
    {
      "figure_id": "Figure 1",
      "key_message": "システム全体のアーキテクチャ",
      "important_elements": ["ハッシュリング", "仮想ノード", "データレプリカ"]
    }
  ]
}
```

## 抽出対象

### 問題設定
- 解決しようとする課題
- 既存手法の限界
- 研究の動機

### 提案手法
- 核心技術の名前
- アプローチの概要
- 新規性のポイント

### 主要結果
- 性能改善の数値
- 評価指標
- 実験結果の要約

### 制約・限界
- 適用できない条件
- 今後の課題
- 実用上の制約

## 使用例

### 入力例
```json
{
  "paper_data": {
    "sections": {
      "abstract": {"text": "We present a novel approach..."},
      "figures": [{"id": "Figure 1", "caption": "System architecture"}]
    }
  },
  "focus_areas": ["problem", "method", "results", "limitations"]
}
```

### 期待される出力
- 構造化された重要情報
- 数値データの抽出
- 図表からの洞察
- 制約事項の特定

## 関連リソース

- `paper://templates/key-insights` - 重要情報抽出テンプレート
- `paper://patterns/paper-structure` - 論文構造パターン

## 注意事項

- 抽出精度は論文の品質に依存する
- 数値データの解釈には注意が必要
- 図表の複雑さによって解析精度が変動する
