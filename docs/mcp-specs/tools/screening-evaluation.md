# screeningEvaluation Tool

## 基本情報

- **Name**: `screeningEvaluation`
- **Category**: Evaluation
- **Purpose**: 15分スクリーニングのための自動採点と判定

## シグネチャ

```typescript
screeningEvaluation(paper_data: object, user_keywords: string[], criteria?: object)
```

## パラメータ

- `paper_data`: analyzePaper の出力データ
- `user_keywords`: ユーザーの研究キーワードリスト
- `criteria`: カスタム評価基準（オプション）

## 動作詳細

1. ユーザーの研究キーワードとのマッチング度を計算
2. 改善率（%）を Abstract/Results から自動抽出
3. コード公開、データセット規模などを点数化
4. 新規性、実用性、実装可能性を評価
5. 15分タイマーで自動判定

## 返却データ

```json
{
  "total_score": 8,
  "breakdown": {
    "relevance": 3,
    "impact": 2,
    "practicality": 2,
    "quality": 1
  },
  "grade": "A",
  "recommendation": "精読推奨",
  "estimated_reading_time": "45-60分",
  "key_factors": [
    "研究キーワードが4個マッチ",
    "25%の性能改善を達成",
    "実装コードが公開済み"
  ],
  "next_steps": [
    "詳細な技術理解",
    "実装検討",
    "関連研究調査"
  ]
}
```

## 評価基準

### 関連性（3点）
- 研究キーワードが3個以上含まれる（+3点）
- 自分の研究領域と一致する（+2点）

### 影響度（2点）
- 既存手法より20%以上改善（+2点）
- 新しいアプローチを提案（+1点）

### 実用性（2点）
- 実装コードが公開されている（+2点）
- 再現可能な実験設定（+1点）

### 品質（1点）
- 大規模データセットで評価（+1点）
- 複数の既存手法と比較（+1点）

## 使用例

### 入力例
```json
{
  "paper_data": {
    "metadata": {"title": "Consistent Hashing..."},
    "sections": {"abstract": {"text": "..."}}
  },
  "user_keywords": ["分散システム", "一貫性", "ハッシュ"],
  "criteria": {"min_score": 6}
}
```

### 期待される出力
- 定量的な採点結果
- A/B/C級の判定
- 精読推奨時間
- 次のアクション提案

## 関連リソース

- `paper://criteria/paper-evaluation` - 論文評価基準
- `paper://templates/screening-checklist` - スクリーニングチェックリスト

## 注意事項

- 評価基準は相対的なものであり、絶対的な指標ではない
- ユーザーの研究分野によって評価結果が変動する
- 15分の時間制限内での判定であることに注意
