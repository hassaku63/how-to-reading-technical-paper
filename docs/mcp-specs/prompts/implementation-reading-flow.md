# Implementation Reading Flow Prompt

## 基本情報

- **Name**: `implementation-reading-flow`
- **Title**: 実装向け読解フロー
- **Description**: ソフトウェアエンジニアが論文のアイデアを実装するための読解フロー
- **Category**: Reading Workflow

## 引数定義

```json
{
  "arguments": [
    {
      "name": "skill_goals",
      "type": "array",
      "items": {"type": "string"},
      "description": "習得したいスキル（分散システム、アルゴリズム最適化、API設計等）",
      "required": true
    },
    {
      "name": "implementation_timeline",
      "type": "string", 
      "description": "実装予定期間（例：1-3ヶ月）",
      "default": "2ヶ月"
    },
    {
      "name": "current_level",
      "type": "string",
      "enum": ["初心者", "中級者", "上級者"],
      "description": "現在の技術レベル",
      "required": true
    }
  ]
}
```

## プロンプトテンプレート

```
以下の手順で論文を実装観点から読解してください：

**Phase 0: 学習目標設定（5分）**
1. スキル獲得目標の明確化：{skill_goals}から3つ選択
2. 現在レベル：{current_level}
3. 実装予定期間：{implementation_timeline}

**Phase 1: 実装アイデア抽出（20分）**
1. システム概要理解：Figure 1を手描きで模写
2. 核心技術特定：新規技術・アルゴリズム名を抽出
3. 技術スタック決定：使用言語・ライブラリを選択
4. MVP機能選定：最小実装で必要な機能を3つ選定

**Phase 2: 実装設計（30-45分）**
1. データモデル設計：主要データ構造を5個まで特定
2. 核心アルゴリズム理解：擬似コードを5-10ステップに分解
3. 実装計画作成：3段階の段階的実装計画

**Phase 3: 実装準備（15分）**
1. 開発環境セットアップ：プロジェクト構成決定
2. 実装可能性チェック：10項目のリスク評価

各段階で実装の具体性を高め、コードレベルでの理解を促進してください。
```

## 使用例

### 入力例
```json
{
  "skill_goals": ["分散システムの設計・実装", "アルゴリズム最適化", "API設計"],
  "implementation_timeline": "3ヶ月",
  "current_level": "中級者"
}
```

### 期待される出力
- 実装に特化した読解プロセス
- 具体的な技術スタックの提案
- 段階的な実装計画
- リスク評価と対策

## 関連リソース

- `paper://guides/implementation-reading` - 実装向け読み方ガイド
- `paper://templates/implementation-plan` - 実装計画テンプレート
- `paper://checklists/implementation-readiness` - 実装準備チェックリスト

## 注意事項

- 実装可能性を常に意識する
- 現在のスキルレベルに応じた計画を立てる
- MVP（最小実装版）から始める
- 各段階で具体的なアウトプットを生成する
