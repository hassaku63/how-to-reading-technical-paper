# criteria/paper-evaluation Resource

## 基本情報

- **URI**: `paper-reading://criteria/paper-evaluation`
- **Purpose**: 論文の品質と価値を評価するための基準
- **Category**: Criteria

## データ構造

```json
{
  "evaluation_criteria": {
    "technical_quality": {
      "name": "技術的品質",
      "weight": 0.3,
      "subcriteria": [
        {
          "name": "新規性",
          "weight": 0.4,
          "description": "既存手法との明確な差別化",
          "evaluation_points": [
            "全く新しいアプローチ（3点）",
            "既存手法の改良（2点）",
            "既存手法の組み合わせ（1点）",
            "既存手法の再実装（0点）"
          ]
        },
        {
          "name": "技術的深度",
          "weight": 0.3,
          "description": "技術的な複雑さと完成度",
          "evaluation_points": [
            "理論的基盤が堅実（3点）",
            "実装詳細が明確（2点）",
            "アルゴリズムが理解可能（1点）",
            "技術的詳細が不明（0点）"
          ]
        },
        {
          "name": "実装可能性",
          "weight": 0.3,
          "description": "実際のシステムへの適用可能性",
          "evaluation_points": [
            "実装コードが公開済み（3点）",
            "実装詳細が十分（2点）",
            "実装方針が明確（1点）",
            "実装が困難（0点）"
          ]
        }
      ]
    },
    "experimental_rigor": {
      "name": "実験の厳密性",
      "weight": 0.25,
      "subcriteria": [
        {
          "name": "データセット",
          "weight": 0.4,
          "description": "使用データセットの適切性",
          "evaluation_points": [
            "大規模・多様なデータセット（3点）",
            "標準ベンチマーク使用（2点）",
            "小規模データセット（1点）",
            "合成データのみ（0点）"
          ]
        },
        {
          "name": "比較手法",
          "weight": 0.3,
          "description": "ベースラインとの比較",
          "evaluation_points": [
            "複数の最新手法と比較（3点）",
            "適切なベースラインと比較（2点）",
            "限定的な比較（1点）",
            "比較が不十分（0点）"
          ]
        },
        {
          "name": "評価指標",
          "weight": 0.3,
          "description": "評価指標の適切性",
          "evaluation_points": [
            "複数の指標で包括的評価（3点）",
            "適切な指標で評価（2点）",
            "限定的な指標（1点）",
            "指標が不適切（0点）"
          ]
        }
      ]
    }
  },
  "scoring_system": {
    "total_max": 100,
    "grade_thresholds": {
      "A": {"min": 80, "description": "高品質、精読推奨"},
      "B": {"min": 60, "description": "良好、条件付き精読"},
      "C": {"min": 40, "description": "普通、要点のみ"},
      "D": {"min": 20, "description": "低品質、スキップ可"},
      "F": {"min": 0, "description": "不適切、スキップ"}
    }
  }
}
```

## 使用方法

### 評価手順
1. 各サブクライテリアについて評価
2. 重み付けして総合点数を計算
3. グレード閾値に基づいて判定

### 評価のコツ
- 客観的な根拠に基づいて評価
- 複数の観点から総合的に判断
- 自分の研究分野との関連性も考慮

## 関連ツール

- `screeningEvaluation` - 自動採点ツール
- `analyzePaper` - 論文解析ツール

## 注意事項

- 評価基準は相対的なものであり、絶対的な指標ではない
- ユーザーの研究分野によって評価結果が変動する
- 定期的な評価基準の見直しと更新が必要
