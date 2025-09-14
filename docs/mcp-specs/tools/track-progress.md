# trackProgress Tool

## 基本情報

- **Name**: `trackProgress`
- **Category**: Evaluation
- **Purpose**: 読解進捗と理解度を記録・管理

## MCP Schema

```json
{
  "name": "trackProgress",
  "description": "読解進捗と理解度を記録・管理",
  "inputSchema": {
    "type": "object",
    "properties": {
      "paper_id": {
        "type": "string",
        "description": "論文の一意識別子"
      },
      "phase": {
        "type": "string",
        "enum": ["screening", "understanding", "deep_reading"],
        "description": "現在の読解段階"
      },
      "completion_data": {
        "type": "object",
        "properties": {
          "checkpoints_completed": {"type": "number"},
          "time_spent": {"type": "number"},
          "understanding_score": {"type": "number"}
        },
        "description": "完了したタスクのデータ"
      }
    },
    "required": ["paper_id", "phase", "completion_data"]
  }
}
```

## パラメータ

- `paper_id`: 論文の一意識別子
- `phase`: 現在の読解段階（"screening", "understanding", "deep_reading"）
- `completion_data`: 完了したタスクのデータ

## 動作詳細

1. 各段階の完了状況を記録
2. 理解度チェックポイントの結果を保存
3. 学習時間と効率を追跡
4. 次のステップを提案

## 返却データ

```json
{
  "paper_id": "paper_001",
  "current_phase": "understanding",
  "progress": {
    "screening": {
      "completed": true,
      "time_spent": 12,
      "score": 8,
      "grade": "A"
    },
    "understanding": {
      "completed": false,
      "time_spent": 25,
      "progress_percentage": 60,
      "checkpoints_passed": 3
    },
    "deep_reading": {
      "completed": false,
      "time_spent": 0,
      "progress_percentage": 0
    }
  },
  "understanding_checkpoints": [
    {
      "checkpoint": "システム全体像の理解",
      "status": "passed",
      "score": 8
    },
    {
      "checkpoint": "アルゴリズムの理解",
      "status": "in_progress",
      "score": 6
    }
  ],
  "next_recommendations": [
    "アルゴリズムの詳細理解を完了",
    "実装計画の作成",
    "関連論文の調査"
  ],
  "efficiency_metrics": {
    "average_reading_speed": "2.5分/ページ",
    "comprehension_rate": 75,
    "time_vs_expected": "15%早い"
  }
}
```

## 進捗追跡項目

### 段階別進捗
- スクリーニング完了状況
- 構造理解の進捗
- 精読の進捗

### 理解度チェックポイント
- システム全体像の理解
- アルゴリズムの理解
- 実装可能性の評価

### 効率指標
- 読解速度（分/ページ）
- 理解率（%）
- 予想時間との比較

## 使用例

### 入力例
```json
{
  "paper_id": "paper_001",
  "phase": "understanding",
  "completion_data": {
    "checkpoints_completed": 3,
    "time_spent": 25,
    "understanding_score": 6
  }
}
```

### 期待される出力
- 現在の進捗状況
- 理解度の評価
- 次のアクション提案
- 効率性の指標

## 関連リソース

- `paper://templates/progress-tracking` - 進捗管理テンプレート
- `paper://criteria/understanding-levels` - 理解度評価基準

## 注意事項

- 進捗データは継続的に更新される
- 理解度は相対的な評価である
- 効率指標は参考値として使用
