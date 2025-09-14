# generateReadingPlan Tool

## 基本情報

- **Name**: `generateReadingPlan`
- **Category**: Generation
- **Purpose**: 目的に応じた読解計画を自動生成

## MCP Schema

```json
{
  "name": "generateReadingPlan",
  "description": "目的に応じた読解計画を自動生成",
  "inputSchema": {
    "type": "object",
    "properties": {
      "paper_data": {
        "type": "object",
        "description": "analyzePaper の出力データ"
      },
      "purpose": {
        "type": "string",
        "enum": ["survey", "implementation"],
        "description": "読解の目的（サーベイ or 実装）"
      },
      "user_profile": {
        "type": "object",
        "properties": {
          "skill_level": {
            "type": "string",
            "enum": ["beginner", "intermediate", "advanced"]
          },
          "target_skills": {
            "type": "array",
            "items": {"type": "string"}
          },
          "time_budget": {
            "type": "string"
          }
        },
        "description": "ユーザーのスキルレベルと目標"
      }
    },
    "required": ["paper_data", "purpose", "user_profile"]
  }
}
```

## パラメータ

- `paper_data`: analyzePaper の出力データ
- `purpose`: 読解の目的（サーベイ or 実装）
- `user_profile`: ユーザーのスキルレベルと目標

## 動作詳細

1. 論文の複雑度とユーザープロファイルを分析
2. 目的に応じた読解段階を設計
3. 各段階の時間配分を最適化
4. チェックポイントと成功基準を設定

## 返却データ

```json
{
  "plan_id": "plan_001",
  "purpose": "implementation",
  "estimated_total_time": "90分",
  "phases": [
    {
      "phase": 1,
      "name": "初動スクリーニング",
      "duration": "15分",
      "objectives": [
        "問題設定の理解",
        "手法の概要把握",
        "実装可能性の評価"
      ],
      "checkpoints": [
        "3行要約の作成",
        "A/B/C級判定"
      ]
    },
    {
      "phase": 2,
      "name": "構造理解",
      "duration": "45分",
      "objectives": [
        "アルゴリズムの詳細理解",
        "データ構造の把握",
        "実装要件の特定"
      ],
      "checkpoints": [
        "システム図の模写",
        "アルゴリズムの分解",
        "実装計画の作成"
      ]
    },
    {
      "phase": 3,
      "name": "実装準備",
      "duration": "30分",
      "objectives": [
        "技術スタックの決定",
        "開発環境の準備",
        "テスト戦略の立案"
      ],
      "checkpoints": [
        "プロジェクト構成の決定",
        "依存関係の確認",
        "MVP機能の選定"
      ]
    }
  ],
  "success_criteria": [
    "システム全体像の理解",
    "核心アルゴリズムの把握",
    "実装可能な計画の作成"
  ],
  "risk_mitigation": [
    "複雑な部分は段階的に理解",
    "不明な用語は事前に調べる",
    "時間超過時は優先順位を調整"
  ]
}
```

## 計画生成ロジック

### サーベイ向け計画
- 初動スクリーニング（10-15分）
- 構造理解（25-40分）
- 精読（必要時のみ、45-90分）

### 実装向け計画
- 学習目標設定（5分）
- 実装アイデア抽出（20分）
- 実装設計（30-45分）
- 実装準備（15分）

## 使用例

### 入力例
```json
{
  "paper_data": {
    "metadata": {"title": "Consistent Hashing..."},
    "sections": {"abstract": {"text": "..."}}
  },
  "purpose": "implementation",
  "user_profile": {
    "skill_level": "intermediate",
    "target_skills": ["分散システム", "アルゴリズム"],
    "time_budget": "2時間"
  }
}
```

### 期待される出力
- 段階的な読解計画
- 各段階の時間配分
- チェックポイントと成功基準
- リスク軽減策

## 関連リソース

- `paper://methodology/survey-guide` - サーベイ向け方法論
- `paper://methodology/implementation-guide` - 実装向け方法論
- `paper://templates/reading-plan` - 読解計画テンプレート

## 注意事項

- 計画は柔軟に調整可能
- ユーザーのスキルレベルに応じて難易度を調整
- 時間制限を考慮した現実的な計画を生成
