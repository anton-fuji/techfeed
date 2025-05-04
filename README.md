#  👽 Tech Feed
## Usage
### Step.1 Feed設定
- `config/feed-setup.yaml`にある`urls`をご自身のQiita, ZennなどのURLに変更してください。
- その他の、プロパティについて
    - `limit`: 記事取得上限
    - `sorted`: 取得記事をソートするかどうか

```yaml
- urls:
    - https://qiita.com/your_name/feed.atom
  template: templates/feeds.tmpl
  limit: 5
  sorted: true

- urls:
    - https://zenn.dev/your_name/feed
  template: templates/feeds.tmpl
  limit: 5
  sorted: false
```
### Step.2 README.mdへ出力
- プレースホルダーをご自身のREADME.mdに加えてください([*mdmo.md*](./memo.md)を参照)
- 以下は、FujiのQiita, Zennを取得したものになります

<!-- feeds:start -->

* [![](./image/qiita.png) サクッとGoで AI エージェントを構築してみる](https://qiita.com/fujifuji1414/items/fc259d51de4aaf1bc75e)

* [![](./image/qiita.png) TypeScriptのコンパイラをGoに移植｜10倍高速になった tsgo とは](https://qiita.com/fujifuji1414/items/98ddf083995f4e03ff32)

* [![](./image/qiita.png) ブラウザでWebサイトが表示されるまでの仕組みを整理してみた](https://qiita.com/fujifuji1414/items/f9c53b451fa4890b8bfc)

* [![](./image/qiita.png) 【Go言語】Qiitaの投稿をGitHubのプロフィールに反映させてみた](https://qiita.com/fujifuji1414/items/f9606bb184951d4a3fb2)



* [![](./image/zenn.png) Dockerイメージの安全性を高める10のセキュリティハック](https://zenn.dev/fuuji/articles/3909c8d444eaa9)

* [![](./image/zenn.png) Dockerイメージ軽量化のアーキテクチャ設計を考える](https://zenn.dev/fuuji/articles/9eb7f2aefcd6c5)

<!-- feeds:end -->