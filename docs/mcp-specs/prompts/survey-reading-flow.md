# Survey Reading Flow Prompt

## 基本情報

- **Name**: `survey-reading-flow`
- **Title**: 高速サーベイ読解フロー
- **Description**: 研究者が大量の論文を効率的にサーベイするための構造化フロー
- **Category**: Reading Workflow

## MCP Schema

```json
{
  "name": "survey-reading-flow",
  "title": "高速サーベイ読解フロー",
  "description": "研究者が大量の論文を効率的にサーベイするための構造化フロー",
  "arguments": [
    {
      "name": "research_domain",
      "type": "string",
      "description": "研究分野（例：分散システム、機械学習、データベース）",
      "required": true
    },
    {
      "name": "survey_goal", 
      "type": "string",
      "description": "サーベイの目的（例：最新動向把握、手法比較、研究ギャップ特定）",
      "required": true
    },
    {
      "name": "time_budget",
      "type": "number",
      "description": "1論文あたりの時間予算（分）",
      "required": false
    }
  ]
}
```

## プロンプトテンプレート

```
以下の手順で論文を効率的にサーベイしてください：

**Phase 1: 初動スクリーニング（{time_budget}分）**
1. 目的を明文化：{survey_goal}
2. Abstract精査：問題・手法・結果を各1行で抽出
3. 図表スキャン：最大サイズの図を30秒で確認
4. Introduction冒頭：既存手法の問題点を1文でメモ
5. Conclusion冒頭：主要成果を1文でメモ
6. 採否判定：A級（精読）/B級（要点のみ）/C級（終了）

**Phase 2: 構造理解（A級論文のみ）**
1. セクション構造把握：目次とページ数確認
2. Related Work差分抽出：既存手法との比較表作成
3. コア技術理解：アルゴリズム/アーキテクチャの入出力
4. 実験設定確認：データセット、ベースライン、評価指標

**Phase 3: アウトプット生成**
1. 3行サマリー作成
2. 研究ギャップの特定
3. 次に読むべき論文の特定

各段階で理解度チェックポイントを設け、人間の思考を外在化してください。
```

## 使用例

### 入力例
```json
{
  "research_domain": "分散システム",
  "survey_goal": "最新の一貫性保証手法を把握したい",
  "time_budget": 15
}
```

### 期待される出力
- 構造化された読解プロセス
- 各段階での具体的なアクション
- 理解度チェックポイント
- アウトプットの形式指定

## 関連リソース

- `paper-reading://guides/survey-reading` - サーベイ向け読み方ガイド
- `paper-reading://templates/reading-notes` - 読書メモテンプレート
- `paper-reading://criteria/paper-evaluation` - 論文評価基準

## 注意事項

- 時間制限を厳守し、深入りを避ける
- 各段階で必ずアウトプットを生成する
- 理解度チェックポイントで人間の思考を確認する

## セキュリティ考慮事項

> [MCP Specification 2025-06-18](https://modelcontextprotocol.io/specification/2025-06-18/server/prompts#security) より:
> 実装は「すべてのプロンプト入出力を慎重に検証し、インジェクション攻撃やリソースへの不正アクセスを防ぐ必要がある」

### 入力検証
- `research_domain` は事前定義された研究分野のみ受け入れる
- `survey_goal` にシステムコマンドや悪意のあるスクリプトが含まれていないか検証
- `time_budget` は適切な範囲（5-60分等）に制限

### インジェクション攻撃対策
- サーベイフローテンプレートへのユーザー入力埋め込み時のエスケープ処理
- プロンプトインジェクションや指示の上書きを防止

### リソースアクセス制御
- サーベイ対象の論文リソースへのアクセス権限を確認
- `paper-reading://` スキームのリソースのみへのアクセスを許可
