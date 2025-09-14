# patterns/common-structures Resource

## 基本情報

- **URI**: `paper-reading://patterns/common-structures`
- **Purpose**: 論文の一般的構造パターン
- **Category**: Pattern

## データ構造

```json
{
  "paper_patterns": {
    "system_papers": {
      "name": "システム論文",
      "description": "新しいシステムやアーキテクチャを提案する論文",
      "common_structure": [
        "Abstract: システムの概要と主要貢献",
        "Introduction: 問題設定と既存システムの限界",
        "System Design: 提案システムのアーキテクチャ",
        "Implementation: 実装詳細と技術的選択",
        "Evaluation: 性能評価と比較実験",
        "Related Work: 既存システムとの比較",
        "Conclusion: 成果と今後の課題"
      ],
      "key_sections": {
        "system_architecture": "Figure 1にシステム全体図",
        "performance_metrics": "スループット、レイテンシ、スケーラビリティ",
        "comparison_baselines": "既存システムとの性能比較",
        "implementation_details": "実装の技術的選択と理由"
      }
    },
    "algorithm_papers": {
      "name": "アルゴリズム論文",
      "description": "新しいアルゴリズムや最適化手法を提案する論文",
      "common_structure": [
        "Abstract: アルゴリズムの概要と性能改善",
        "Introduction: 問題定義と既存アルゴリズムの限界",
        "Algorithm: 提案アルゴリズムの詳細",
        "Analysis: 計算量と理論的保証",
        "Experiments: ベンチマークでの性能評価",
        "Related Work: 既存アルゴリズムとの比較",
        "Conclusion: 成果と理論的貢献"
      ],
      "key_sections": {
        "algorithm_description": "擬似コードまたは詳細な手順",
        "complexity_analysis": "時間・空間計算量の分析",
        "theoretical_guarantees": "理論的保証と証明",
        "benchmark_results": "標準ベンチマークでの性能"
      }
    },
    "evaluation_papers": {
      "name": "評価論文",
      "description": "既存手法の包括的な評価や比較を行う論文",
      "common_structure": [
        "Abstract: 評価の範囲と主要発見",
        "Introduction: 評価の動機と重要性",
        "Methodology: 評価手法と実験設計",
        "Results: 詳細な評価結果",
        "Analysis: 結果の解釈と洞察",
        "Related Work: 既存の評価研究",
        "Conclusion: 発見と今後の方向性"
      ],
      "key_sections": {
        "evaluation_methodology": "評価手法の詳細",
        "dataset_description": "使用データセットの説明",
        "comparison_matrix": "手法間の比較表",
        "statistical_analysis": "統計的有意性の検証"
      }
    }
  },
  "section_patterns": {
    "abstract_patterns": [
      "問題設定（第1文）",
      "既存手法の限界（第2-3文）",
      "提案手法の概要（第4-5文）",
      "主要結果（第6-7文）",
      "今後の展望（最終文）"
    ],
    "introduction_patterns": [
      "研究分野の重要性",
      "既存手法の限界",
      "本論文の貢献",
      "論文の構成"
    ],
    "related_work_patterns": [
      "関連研究の分類",
      "各手法の特徴と限界",
      "本論文との差分",
      "研究の位置づけ"
    ]
  },
  "figure_patterns": {
    "system_architecture": "システム全体の構成図",
    "algorithm_flow": "アルゴリズムの処理フロー",
    "performance_comparison": "性能比較グラフ",
    "experimental_setup": "実験設定の図"
  }
}
```

## 使用方法

### 論文タイプの識別
1. AbstractとIntroductionを読んで論文タイプを判定
2. 該当するパターンの構造を確認
3. 重要セクションに重点を置いて読解

### 効率的な読解
- パターンに基づいて重要箇所を特定
- 各セクションの目的を理解して読解
- 図表の種類から内容を予測

## 関連ツール

- `analyzePaper` - 論文構造解析ツール
- `extractKeyInsights` - 重要情報抽出ツール

## 注意事項

- パターンは一般的な傾向であり、例外もある
- 分野によって構造が異なる場合がある
- 新しい論文形式に対応するため定期的な更新が必要
