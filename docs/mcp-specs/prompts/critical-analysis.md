# Critical Analysis Prompt

## 基本情報

- **Name**: `critical-analysis`
- **Title**: 批判的論文分析
- **Description**: 論文の主張を批判的に検証し、限界と問題点を特定するフロー
- **Category**: Analysis Workflow

## MCP Schema

```json
{
  "name": "critical-analysis",
  "title": "批判的論文分析",
  "description": "論文の主張を批判的に検証し、限界と問題点を特定するフロー",
  "arguments": [
    {
      "name": "analysis_depth",
      "type": "string",
      "description": "分析の深さ（浅い、中程度、深い）",
      "required": false
    },
    {
      "name": "focus_area",
      "type": "string",
      "description": "重点分析領域（実験設計、理論的基盤、実装詳細等）",
      "required": false
    }
  ]
}
```

## プロンプトテンプレート

```
以下の観点で論文を批判的に分析してください：

**実験の妥当性チェック**
- ベースラインは適切か？
- データセットは現実的か？
- 評価指標は目的に合致するか？
- アブレーションスタディは十分か？

**理論的基盤の検証**
- 証明は完全か？
- 仮定は現実的か？
- 計算量の分析は正確か？
- エッジケースは考慮されているか？

**実装・再現性の評価**
- 実装詳細は十分に記述されているか？
- ハイパーパラメータは公開されているか？
- コードは利用可能か？
- 再現に必要な情報は揃っているか？

**実用性の検討**
- スケーラビリティは十分か？
- 計算コストは現実的か？
- 運用・保守は容易か？
- セキュリティ面での考慮はあるか？

各項目について具体的な根拠を示し、改善提案も含めて分析してください。
```

## 使用例

### 入力例
```json
{
  "analysis_depth": "中程度",
  "focus_area": "実験設計"
}
```

### 期待される出力
- 各観点での具体的な問題点
- 改善提案
- 論文の信頼性評価
- 実用性の総合判断

## 関連リソース

- `paper-reading://criteria/paper-evaluation` - 論文評価基準
- `paper-reading://templates/critical-analysis` - 批判的分析テンプレート
- `paper-reading://patterns/common-issues` - よくある問題パターン

## 注意事項

- 客観的な根拠に基づいて分析する
- 建設的な批判を心がける
- 改善提案を含める
- 論文の価値も適切に評価する

## セキュリティ考慮事項

> [MCP Specification 2025-06-18](https://modelcontextprotocol.io/specification/2025-06-18/server/prompts#security) より:
> 実装は「すべてのプロンプト入出力を慎重に検証し、インジェクション攻撃やリソースへの不正アクセスを防ぐ必要がある」

### 入力検証
- `analysis_depth` は事前定義されたレベルのみ受け入れる
- `focus_area` に悪意のあるスクリプトやコマンドが含まれていないか検証
- 批判分析パラメータの範囲を適切に制限

### インジェクション攻撃対策
- 批判分析テンプレートへのユーザー入力埋め込み時のエスケープ処理
- プロンプトインジェクションや指示の上書きを防止

### リソースアクセス制御
- 分析対象の論文リソースへのアクセス権限を確認
- 許可された `paper-reading://` リソースのみへのアクセスを許可
