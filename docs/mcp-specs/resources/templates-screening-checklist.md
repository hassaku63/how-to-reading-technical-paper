# templates/screening-checklist Resource

## 基本情報

- **URI**: `paper-reading://templates/screening-checklist`
- **Purpose**: 論文スクリーニングのための標準化されたチェックリスト
- **Category**: Template

## データ構造

```json
{
  "checklist": {
    "relevance": [
      {
        "id": "keyword_match",
        "question": "研究キーワードが3個以上含まれるか",
        "points": 3,
        "evaluation_hint": "タイトル、Abstract、Introductionでキーワード出現をチェック"
      },
      {
        "id": "domain_fit",
        "question": "自分の研究領域と一致するか",
        "points": 2,
        "evaluation_hint": "会議/ジャーナルの分野、Referenced papers を確認"
      }
    ],
    "impact": [
      {
        "id": "improvement",
        "question": "既存手法より20%以上改善しているか",
        "points": 2,
        "evaluation_hint": "Results, Tables で数値を確認"
      },
      {
        "id": "novelty",
        "question": "新しいアプローチを提案しているか",
        "points": 1,
        "evaluation_hint": "Contributions, 既存研究との差分を確認"
      }
    ],
    "practicality": [
      {
        "id": "code_available",
        "question": "実装コードが公開されているか",
        "points": 2,
        "evaluation_hint": "GitHub/GitLab URL、Supplementary material を探す"
      },
      {
        "id": "reproducible",
        "question": "再現可能な実験設定か",
        "points": 1,
        "evaluation_hint": "パラメータ、データセット、環境の記載を確認"
      }
    ],
    "quality": [
      {
        "id": "evaluation",
        "question": "大規模データセットで評価されているか",
        "points": 1,
        "evaluation_hint": "標準ベンチマーク使用、データサイズを確認"
      },
      {
        "id": "comparison",
        "question": "複数の既存手法と比較されているか",
        "points": 1,
        "evaluation_hint": "Baseline methods の数と新しさを確認"
      }
    ]
  },
  "scoring": {
    "A": {
      "min": 7,
      "max": 10,
      "action": "精読推奨",
      "time": "45-60分",
      "next_steps": ["詳細な技術理解", "実装検討", "関連研究調査"]
    },
    "B": {
      "min": 4,
      "max": 6,
      "action": "条件付き精読",
      "time": "20-30分",
      "next_steps": ["特定セクションのみ精読", "アイデア抽出"]
    },
    "C": {
      "min": 0,
      "max": 3,
      "action": "スキップ可",
      "time": "記録のみ",
      "next_steps": ["メタデータ記録", "将来の参照用に保存"]
    }
  }
}
```

## 使用方法

### スクリーニング手順
1. 各項目について質問に回答
2. 該当する項目の点数を合計
3. 総合点数に基づいてA/B/C級を判定
4. 判定結果に応じて次のアクションを決定

### 評価のコツ
- まずFigure 1とTable 1を見る
- Abstractの第1文と最終文を重視
- 数値データ（改善率、性能指標）を確認
- コード公開状況を必ずチェック

## 関連ツール

- `screeningEvaluation` - 自動採点ツール
- `analyzePaper` - 論文解析ツール

## 注意事項

- 評価は15分以内に完了する
- 不明な項目は0点として扱う
- 判定結果は参考値として使用
