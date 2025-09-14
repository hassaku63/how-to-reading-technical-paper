# evaluateComplexity Tool

## 基本情報

- **Name**: `evaluateComplexity`
- **Category**: Evaluation
- **Purpose**: 実装難易度を多面的に評価

## シグネチャ

```typescript
evaluateComplexity(paper_data: object, target_skills: string[])
```

## パラメータ

- `paper_data`: analyzePaper の出力データ
- `target_skills`: 習得目標スキルリスト

## 動作詳細

1. アルゴリズムの複雑度を分析
2. 必要な前提知識を評価
3. 実装環境の要求を特定
4. 学習コストを推定

## 返却データ

```json
{
  "overall_difficulty": "medium",
  "complexity_breakdown": {
    "algorithmic": "medium",
    "mathematical": "low",
    "system_design": "high",
    "implementation": "medium"
  },
  "prerequisite_skills": [
    "分散システムの基礎",
    "ハッシュ関数の理解",
    "ネットワークプログラミング"
  ],
  "learning_curve": {
    "beginner_friendly": false,
    "estimated_study_time": "2-3ヶ月",
    "recommended_approach": "段階的実装"
  },
  "risk_factors": [
    "分散システムの複雑さ",
    "ネットワーク分断の処理",
    "データ一貫性の保証"
  ]
}
```

## 複雑度評価項目

### アルゴリズム複雑度
- 時間計算量（O(n), O(log n), O(n²)等）
- 空間計算量（メモリ使用量）
- 実装の複雑さ

### 数学的複雑度
- 数式の理解難易度
- 証明の複雑さ
- 統計・確率の知識要求

### システム設計複雑度
- アーキテクチャの複雑さ
- コンポーネント間の相互作用
- スケーラビリティの考慮

### 実装複雑度
- プログラミングの難易度
- デバッグの困難さ
- テストの複雑さ

## 使用例

### 入力例
```json
{
  "paper_data": {
    "sections": {
      "method": {"algorithms": ["ConsistentHash"]},
      "figures": [{"type": "architecture"}]
    }
  },
  "target_skills": ["分散システム", "アルゴリズム最適化"]
}
```

### 期待される出力
- 総合的な難易度評価
- 各観点での複雑度分析
- 前提知識の特定
- 学習戦略の提案

## 関連リソース

- `paper://criteria/implementation-feasibility` - 実装可能性評価基準
- `paper://patterns/algorithm-complexity` - アルゴリズム複雑度パターン

## 注意事項

- 複雑度は相対的な評価である
- 習得目標スキルによって評価が変動
- 段階的学習を推奨
