# Output Generation Prompt

## 基本情報

- **Name**: `output-generation`
- **Title**: 構造化アウトプット生成
- **Description**: 読解結果を目的に応じた形式で構造化してアウトプットするフロー
- **Category**: Output Workflow

## 引数定義

```json
{
  "arguments": [
    {
      "name": "output_format",
      "type": "string",
      "enum": ["研究メモ", "技術ブログ", "実装計画", "プレゼン資料", "レポート"],
      "description": "出力形式",
      "required": true
    },
    {
      "name": "target_audience",
      "type": "string",
      "description": "対象読者（研究者、エンジニア、一般読者等）",
      "required": true
    },
    {
      "name": "key_insights",
      "type": "array",
      "items": {"type": "string"},
      "description": "強調したい洞察・ポイント",
      "required": false
    }
  ]
}
```

## プロンプトテンプレート

```
以下の形式で{output_format}を生成してください：

**対象読者**: {target_audience}
**強調ポイント**: {key_insights}

**基本構造**
1. **エグゼクティブサマリー**（3行以内）
2. **問題設定と背景**
3. **提案手法の概要**
4. **主要な結果・洞察**
5. **実装・応用の可能性**
6. **限界と今後の課題**
7. **次のアクション**

**各セクションの記述方針**
- 専門用語は適切に説明
- 具体的な数値・例を含める
- 図表があれば参照
- 実装観点も含める
- 批判的視点も織り込む

**品質チェック**
- 論理的な流れになっているか？
- 対象読者に適したレベルか？
- 実用的な情報が含まれているか？
- 次のアクションが明確か？

このテンプレートに従って、読解内容を価値あるアウトプットに変換してください。
```

## 使用例

### 入力例
```json
{
  "output_format": "技術ブログ",
  "target_audience": "ソフトウェアエンジニア",
  "key_insights": ["分散システムの一貫性保証", "実装の複雑さ", "性能の改善効果"]
}
```

### 期待される出力
- 対象読者に適した内容レベル
- 構造化された読みやすい形式
- 実用的な洞察と次のアクション
- 専門用語の適切な説明

## 関連リソース

- `paper://templates/output-formats` - 出力形式テンプレート集
- `paper://guides/writing-technical-content` - 技術文書作成ガイド
- `paper://patterns/effective-communication` - 効果的なコミュニケーションパターン

## 注意事項

- 対象読者のレベルに合わせる
- 実用的な情報を含める
- 次のアクションを明確にする
- 批判的視点も含める
