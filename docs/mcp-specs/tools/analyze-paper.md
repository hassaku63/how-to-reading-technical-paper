# analyzePaper Tool

## 基本情報

- **Name**: `analyzePaper`
- **Category**: Paper Analysis
- **Purpose**: 論文PDFを解析し、構造化データを抽出する

## シグネチャ

```typescript
analyzePaper(pdf_path: string, analysis_type: "survey" | "implementation")
```

## パラメータ

- `pdf_path`: 解析対象の論文PDFファイルパス
- `analysis_type`: 解析の観点（サーベイ向け or 実装向け）

## 動作詳細

1. PDFをパースして論文の構造要素を抽出
2. Abstract、Introduction、Method、Results、Conclusion を分離
3. 図表（Figure/Table）のキャプションと位置を記録
4. 引用文献数とコード/データセットのURLを抽出
5. 数式、アルゴリズムブロックを識別

## 返却データ

```json
{
  "metadata": {
    "title": "Consistent Hashing and Random Trees",
    "authors": ["David Karger", "Eric Lehman"],
    "year": 2024,
    "venue": "SOSP"
  },
  "sections": {
    "abstract": {
      "text": "We present a novel approach to...",
      "keywords": ["分散システム", "一貫性", "ハッシュ"]
    },
    "introduction": {
      "text": "...",
      "problem_statement": "既存の分散システムは..."
    },
    "method": {
      "text": "...",
      "algorithms": ["Algorithm 1: ConsistentHash"]
    }
  },
  "figures": [
    {
      "id": "Figure 1",
      "caption": "System architecture overview",
      "page": 2,
      "type": "architecture"
    }
  ],
  "tables": [
    {
      "id": "Table 1",
      "caption": "Performance comparison",
      "page": 8
    }
  ],
  "citations_count": 45,
  "code_available": true,
  "dataset_url": "https://github.com/...",
  "equations_count": 12
}
```

## 使用例

### 入力例
```json
{
  "pdf_path": "/path/to/paper.pdf",
  "analysis_type": "implementation"
}
```

### 期待される出力
- 論文の構造化されたメタデータ
- 各セクションの内容とキーワード
- 図表の情報と位置
- 実装関連の情報（コード、データセット）

## 関連リソース

- `paper://templates/paper-analysis` - 論文解析テンプレート
- `paper://patterns/paper-structure` - 論文構造パターン

## 注意事項

- PDFの品質によって解析精度が変動する
- 数式や図表の複雑さに応じて処理時間が増加する
- 実装観点での解析時は技術的詳細を重視する
