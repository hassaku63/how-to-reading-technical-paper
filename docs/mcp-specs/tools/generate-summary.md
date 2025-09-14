# generateSummary Tool

## 基本情報

- **Name**: `generateSummary`
- **Category**: Generation
- **Purpose**: 構造化されたサマリーを生成

## シグネチャ

```typescript
generateSummary(paper_data: object, summary_type: "survey" | "implementation", format: "brief" | "detailed")
```

## パラメータ

- `paper_data`: analyzePaper の出力データ
- `summary_type`: サマリーの種類（サーベイ向け or 実装向け）
- `format`: 詳細度（簡潔 or 詳細）

## 動作詳細

1. 論文の核心情報を抽出
2. 目的に応じた観点で整理
3. 構造化された形式で要約
4. 次のアクションを提案

## 返却データ

```json
{
  "summary_type": "implementation",
  "format": "detailed",
  "executive_summary": "Consistent Hashing手法により分散システムの一貫性と可用性を両立する新アプローチを提案。25%の性能向上を達成。",
  "key_insights": {
    "problem": "既存の分散システムは一貫性と可用性のトレードオフに課題",
    "solution": "Consistent Hashing with Virtual Nodes",
    "results": "25%のスループット向上、99.9%の可用性達成",
    "innovation": "仮想ノードによる負荷分散の最適化"
  },
  "technical_details": {
    "algorithms": ["ConsistentHash", "VirtualNodeAssignment", "DataReplication"],
    "data_structures": ["HashRing", "VirtualNodeMap", "ReplicaSet"],
    "complexity": {
      "time": "O(log n) for lookup",
      "space": "O(n) for storage"
    }
  },
  "implementation_guidance": {
    "difficulty": "medium",
    "prerequisites": ["分散システム基礎", "ハッシュ関数", "ネットワークプログラミング"],
    "estimated_effort": "2-3ヶ月",
    "recommended_approach": "段階的実装（MVP → 完全版）"
  },
  "next_actions": [
    "GitHubリポジトリの調査",
    "関連論文の読解",
    "プロトタイプ実装の開始"
  ],
  "limitations": [
    "ノード数が少ない場合の性能劣化",
    "メモリ使用量の増加",
    "ネットワーク分断時の復旧時間"
  ]
}
```

## サマリー形式

### サーベイ向けサマリー
- 研究の位置づけと貢献
- 既存研究との差分
- 技術的革新点
- 今後の研究方向

### 実装向けサマリー
- システム全体像
- 核心技術の詳細
- 実装要件と制約
- 実装計画の提案

## 使用例

### 入力例
```json
{
  "paper_data": {
    "metadata": {"title": "Consistent Hashing..."},
    "sections": {"abstract": {"text": "..."}}
  },
  "summary_type": "implementation",
  "format": "detailed"
}
```

### 期待される出力
- 目的に応じた構造化サマリー
- 技術的詳細の整理
- 実装ガイダンス
- 次のアクション提案

## 関連リソース

- `paper://templates/summary-formats` - サマリー形式テンプレート
- `paper://patterns/paper-structure` - 論文構造パターン

## 注意事項

- サマリーは目的に応じてカスタマイズされる
- 技術的詳細のレベルは対象読者に応じて調整
- 次のアクションは実現可能なものを提案
