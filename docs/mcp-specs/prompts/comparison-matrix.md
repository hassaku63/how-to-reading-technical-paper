# Comparison Matrix Prompt

## 基本情報

- **Name**: `comparison-matrix`
- **Title**: 論文比較分析マトリクス
- **Description**: 複数の論文を体系的に比較し、各手法の特徴と優劣を整理するフロー
- **Category**: Analysis Workflow

## 引数定義

```json
{
  "arguments": [
    {
      "name": "comparison_criteria",
      "type": "array",
      "items": {"type": "string"},
      "description": "比較基準（性能、実装難易度、スケーラビリティ等）",
      "required": true
    },
    {
      "name": "paper_count",
      "type": "number",
      "description": "比較する論文数",
      "default": 3
    }
  ]
}
```

## プロンプトテンプレート

```
以下のマトリクス形式で論文を比較分析してください：

**比較マトリクス**
| 論文 | 問題設定 | 手法 | 性能 | 実装難易度 | スケーラビリティ | 限界 |
|------|----------|------|------|------------|------------------|------|
| 論文A | | | | | | |
| 論文B | | | | | | |
| 論文C | | | | | | |

**比較基準：{comparison_criteria}**

**各論文の特徴**
1. **論文A**: 強み・弱み・適用場面
2. **論文B**: 強み・弱み・適用場面
3. **論文C**: 強み・弱み・適用場面

**総合評価**
- 最も優れている手法：
- 実装に適している手法：
- 研究の方向性：
- 残された課題：

**次のアクション**
- 深く読むべき論文：
- 実装を試すべき手法：
- 調査すべき関連研究：
```

## 使用例

### 入力例
```json
{
  "comparison_criteria": ["性能", "実装難易度", "スケーラビリティ", "メモリ効率"],
  "paper_count": 3
}
```

### 期待される出力
- 構造化された比較マトリクス
- 各論文の特徴分析
- 総合的な評価と判断
- 具体的な次のアクション

## 関連リソース

- `paper://templates/comparison-matrix` - 比較マトリクステンプレート
- `paper://criteria/paper-evaluation` - 論文評価基準
- `paper://patterns/methodology-comparison` - 手法比較パターン

## 注意事項

- 客観的な基準で比較する
- 各論文の文脈を考慮する
- 実用性も含めて評価する
- 次のアクションを明確にする
