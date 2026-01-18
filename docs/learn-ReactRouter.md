# React Router 学習ノート

このドキュメントでは、React Router v7を使ったSPA開発の流れに沿って、必要な概念を順に学んでいきます。

## 目次

1. [React Routerとは](#react-routerとは)
2. [プロジェクトの始め方](#プロジェクトの始め方)
3. [ファイル構造を理解する](#ファイル構造を理解する)
4. [ステップ1：複数のページを作る](#ステップ1複数のページを作る)
5. [ステップ2：各ページでデータを取得する](#ステップ2各ページでデータを取得する)
6. [ステップ3：データを送信する](#ステップ3データを送信する)
7. [ステップ4：ページ間を移動する](#ステップ4ページ間を移動する)
8. [ステップ5：URLからパラメータを取得する](#ステップ5urlからパラメータを取得する)
9. [高度なUI制御](#高度なui制御)

---

## React Routerとは

**問題：Reactだけでは単一ページしか作れない**

通常のReactアプリは1つのページ（コンポーネント）しか表示できません。しかし、実際のWebアプリケーションでは：
- トップページ（`/`）
- 商品一覧ページ（`/products`）
- 商品詳細ページ（`/products/123`）

のように、複数のページが必要です。

**解決：React Router**

React RouterはReactでページ遷移（ルーティング）を実現するライブラリです。URLに応じて異なるコンポーネントを表示し、ページ全体をリロードせずに画面を切り替えます（SPA: Single Page Application）。

**React Router v7の特徴:**
- **ファイルベースルーティング**: routes.tsでルートを一元管理
- **データフェッチ統合**: loader関数でページ表示前にデータ取得
- **フォーム処理統合**: action関数でデータ送信を簡潔に記述
- **型安全**: TypeScriptで型が自動生成される

---

## プロジェクトの始め方

React Router v7を使ったプロジェクトは、Viteテンプレートから始めるのが一般的です。

### 新規プロジェクトの作成

```bash
npm create vite@latest
```

実行後、以下の質問に答えます：

1. **Project name**: プロジェクト名を入力（例：`my-app`）
2. **Select a framework**: `React`を選択
3. **Select a variant**: `React Router v7 ↗`を選択

これでReact + React Router v7のプロジェクトが作成されます。

### プロジェクトのセットアップ

```bash
cd my-app
npm install
npm run dev
```

開発サーバーが起動し、`http://localhost:5173`でアプリが表示されます。

---

## ファイル構造を理解する

React Router v7プロジェクトの典型的なファイル構造を理解しましょう。

### 典型的なディレクトリ構成

```
my-app/
├─ app/
│  ├─ root.tsx           ... 全体レイアウト・共通UI・Outlet配置
│  ├─ routes.ts          ... ルーティング設定
│  ├─ routes/            ... 各ページコンポーネント
│  │  ├─ index.tsx       ... トップページ (/)
│  │  ├─ about.tsx       ... /about
│  │  └─ contact.tsx     ... /contact
│  ├─ layouts/           ... レイアウトコンポーネント（任意）
│  ├─ api/               ... API呼び出し関数（任意）
│  └─ app.css            ... グローバルスタイル
├─ public/               ... 静的ファイル（画像など）
├─ package.json          ... 依存管理
└─ vite.config.ts        ... ビルド設定
```

### 各ファイルの役割

**1. root.tsx - アプリ全体のレイアウト**

すべてのページに共通するレイアウト（ヘッダー、フッターなど）を定義します。

```tsx
// app/root.tsx
import { Outlet } from "react-router";

export default function Root() {
  return (
    <html lang="ja">
      <body>
        <header>
          <nav>{/* ナビゲーション */}</nav>
        </header>
        <main>
          <Outlet /> {/* 各ページがここに表示される */}
        </main>
        <footer>{/* フッター */}</footer>
      </body>
    </html>
  );
}
```

`<Outlet />`が重要です。ここに各ページのコンポーネントが表示されます。

**2. routes.ts - ルーティング設定**

URLとコンポーネントの対応関係を定義します。

```tsx
// app/routes.ts
import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/index.tsx"),
  route("about", "routes/about.tsx"),
] satisfies RouteConfig;
```

**3. routes/ - 各ページコンポーネント**

URLに対応する実際のページを作成します。

```tsx
// app/routes/index.tsx
export default function Home() {
  return <h1>ホームページ</h1>;
}
```

**4. layouts/ - レイアウトコンポーネント（任意）**

複数ページで共有するレイアウト（サイドバーなど）を定義します。

**5. api/ - API呼び出し関数（任意）**

バックエンドAPIとの通信をまとめます。

```tsx
// app/api/memos.ts
export async function getMemos() {
  return await fetch('/api/memos').then(res => res.json());
}
```

### Outletとネストされたレイアウト

**Outlet**は「子ルートの表示場所」を示すプレースホルダーです。

**使用例：サイドバー付きレイアウト**

```tsx
// app/layouts/sidebar.tsx
import { Outlet } from "react-router";

export default function SidebarLayout() {
  return (
    <div className="flex">
      <aside className="sidebar">
        {/* サイドバーの内容 */}
        <nav>
          <a href="/">ホーム</a>
          <a href="/profile">プロフィール</a>
        </nav>
      </aside>
      <div className="content">
        <Outlet /> {/* ここに各ページが表示される */}
      </div>
    </div>
  );
}
```

**routes.tsでレイアウトを適用:**

```tsx
import { type RouteConfig, index, route, layout } from "@react-router/dev/routes";

export default [
  layout("layouts/sidebar.tsx", [
    index("routes/index.tsx"),           // サイドバー付き
    route("profile", "routes/profile.tsx"), // サイドバー付き
  ]),
  route("about", "routes/about.tsx"),    // サイドバーなし
] satisfies RouteConfig;
```

**動作:**
- `/` と `/profile` はサイドバーレイアウトが適用される
- `/about` はレイアウトなし（root.tsxのみ）

---

## ステップ1：複数のページを作る

最初のステップは、URLに応じて異なるページ（コンポーネント）を表示することです。

### ルート定義ファイル（routes.ts）

React Router v7では、`app/routes.ts`ですべてのルートを定義します。

```tsx
// app/routes.ts
import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/index.tsx"),              // "/" → index.tsx
    route("about", "routes/about.tsx"),     // "/about" → about.tsx
    route("contact", "routes/contact.tsx"), // "/contact" → contact.tsx
] satisfies RouteConfig;
```

### ルートコンポーネント

各ルートに対応するコンポーネントファイルを作成します。

```tsx
// app/routes/index.tsx
export default function Home() {
  return (
    <div>
      <h1>ホームページ</h1>
      <p>ようこそ！</p>
    </div>
  );
}
```

```tsx
// app/routes/about.tsx
export default function About() {
  return (
    <div>
      <h1>Aboutページ</h1>
      <p>このサイトについて</p>
    </div>
  );
}
```

これだけで、`/`にアクセスするとHomeコンポーネントが、`/about`にアクセスするとAboutコンポーネントが表示されます。

### ルート定義の種類

**1. index() - ルートパス**
```tsx
index("routes/index.tsx")
// URL: "/"
```

**2. route() - 通常のルート**
```tsx
route("about", "routes/about.tsx")
// URL: "/about"

route("users", "routes/users.tsx")
// URL: "/users"
```

**3. 動的ルート**
```tsx
route("products/:id", "routes/product.tsx")
// URL: "/products/123" → idは"123"
// URL: "/products/456" → idは"456"
```

**4. ネストされたルート**
```tsx
route("users", "routes/users.tsx", [
    index("routes/users/index.tsx"),      // "/users"
    route(":id", "routes/users/detail.tsx"), // "/users/123"
])
```

### 使用例：メモアプリのルート

```tsx
// app/routes.ts
export default [
    index("routes/index.tsx"),                    // "/" - メモ一覧
    route("memos/:id", "routes/memo.tsx"),        // "/memos/1" - メモ詳細
    route("memos/:id/edit", "routes/edit_memo.tsx"), // "/memos/1/edit" - メモ編集
] satisfies RouteConfig;
```

| URL | 表示されるページ | ファイル |
|-----|--------------|---------|
| `/` | メモ一覧 | `routes/index.tsx` |
| `/memos/1` | ID=1のメモ詳細 | `routes/memo.tsx` |
| `/memos/1/edit` | ID=1のメモ編集 | `routes/edit_memo.tsx` |

---

## ステップ2：各ページでデータを取得する

ページができたら、次は各ページで必要なデータを取得します。

### 問題：従来の方法（useEffect）は面倒

従来のReactでは、useEffectとuseStateを使ってデータを取得していました。

```tsx
// ❌ 従来の方法：useEffectでデータ取得
export default function MemoList() {
  const [memos, setMemos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('/api/memos')
      .then(res => res.json())
      .then(data => {
        setMemos(data);
        setLoading(false);
      })
      .catch(err => {
        setError(err);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>読み込み中...</div>;
  if (error) return <div>エラー: {error.message}</div>;

  return (
    <div>
      {memos.map(memo => <div key={memo.id}>{memo.title}</div>)}
    </div>
  );
}
```

**問題点:**
- ローディング状態を手動で管理
- エラーハンドリングを手動で管理
- データがない時の処理を毎回書く必要がある

### 解決：loader関数

React Router v7では、`loader`関数でデータ取得を定義します。
loader関数によって「サーバーとのやり取りに関する」useStateやuseEffectは不要になる。
(純粋なローカルでのUI状態管理のためにはuseState等が必要)

```tsx
// app/routes/index.tsx
import type { Route } from "./+types";

// ページ表示前にデータを取得
export async function loader() {
  const memos = await fetch('/api/memos').then(res => res.json());
  return { memos };
}

// loaderDataでデータを受け取る
export default function MemoList({ loaderData }: Route.ComponentProps) {
  const { memos } = loaderData;

  return (
    <div>
      {memos.map(memo => <div key={memo.id}>{memo.title}</div>)}
    </div>
  );
}
```

**メリット:**
- ローディング状態は自動管理（データが揃ってからレンダリング）
- コンポーネントがシンプルになる
- 型が自動推論される

### loader関数の基本パターン

**1. 単一のデータ取得**
```tsx
export async function loader() {
  const users = await fetchUsers();
  return { users };
}

export default function Users({ loaderData }: Route.ComponentProps) {
  const { users } = loaderData;
  return <div>{/* users表示 */}</div>;
}
```

**2. 複数のデータを並列取得**
```tsx
export async function loader() {
  // Promise.allで並列実行
  const [memos, categories] = await Promise.all([
    fetchMemos(),
    fetchCategories(),
  ]);
  return { memos, categories };
}

export default function MemoList({ loaderData }: Route.ComponentProps) {
  const { memos, categories } = loaderData;
  return <div>{/* ... */}</div>;
}
```

**3. オブジェクトをそのまま返す**
```tsx
export async function loader() {
  const user = await fetchUser();
  return user; // { name: "太郎", age: 25 }
}

export default function Profile({ loaderData }: Route.ComponentProps) {
  // loaderDataが直接userオブジェクト
  return <div>{loaderData.name}</div>;
}
```

### 使用例：メモ一覧の取得

```tsx
// app/routes/index.tsx
import { getMemos } from "~/api/memos";
import type { Route } from "./+types";

export async function loader() {
  const memos = await getMemos();
  return { memos };
}

export default function MemoList({ loaderData }: Route.ComponentProps) {
  const { memos } = loaderData;

  return (
    <div>
      <h1>メモ一覧 ({memos.length}件)</h1>
      {memos.map(memo => (
        <div key={memo.id}>
          <h2>{memo.title}</h2>
          <p>{memo.content}</p>
        </div>
      ))}
    </div>
  );
}
```

### 型安全性：自動生成される型

React Router v7の大きな特徴は、**型が自動生成される**ことです。

**+types ファイル**

各ルートファイルに対して、`+types.ts`という型定義ファイルが自動生成されます。

```tsx
// app/routes/index.tsx から自動生成される型
// app/routes/+types/index.d.ts

export interface Route {
  // loader関数の引数の型
  LoaderArgs: {
    request: Request;
    params: Record<string, string>;
    context?: any;
  };

  // action関数の引数の型
  ActionArgs: {
    request: Request;
    params: Record<string, string>;
    context?: any;
  };

  // コンポーネントのPropsの型
  ComponentProps: {
    loaderData: Awaited<ReturnType<typeof loader>>;
    actionData?: Awaited<ReturnType<typeof action>>;
    params: Record<string, string>;
  };
}
```

**型の使い方:**

```tsx
import type { Route } from "./+types";

// loader関数：LoaderArgsを使う
export async function loader({ request, params }: Route.LoaderArgs) {
  // requestとparamsの型が自動的にわかる
  const url = new URL(request.url);
  const id = params.id;
  return { data: await fetchData(id) };
}

// action関数：ActionArgsを使う
export async function action({ request, params }: Route.ActionArgs) {
  const formData = await request.formData();
  return { success: true };
}

// コンポーネント：ComponentPropsを使う
export default function MyPage({ loaderData, actionData, params }: Route.ComponentProps) {
  // loaderDataの型が自動推論される！
  // loaderの戻り値の型がそのまま使える
  return <div>{loaderData.data}</div>;
}
```

**メリット:**
1. **手動で型を書く必要がない**: loader/actionの戻り値から自動推論
2. **タイプセーフ**: loaderDataへのアクセスでIDEが補完してくれる
3. **リファクタリングが安全**: loader関数の戻り値を変えると、コンポーネント側でエラーになる

### データフローと型の流れ

```
1. ユーザーがURLにアクセス
   ↓
2. loader関数が実行される
   export async function loader({ params }: Route.LoaderArgs) {
     const memo = await getMemoById(params.id);
     return { memo }; // この型が自動で推論される
   }
   ↓
3. loaderDataとして自動的に型付けされる
   ↓
4. コンポーネントで型安全に使える
   export default function Memo({ loaderData }: Route.ComponentProps) {
     const { memo } = loaderData; // memoの型が自動的にわかる！
     return <h1>{memo.title}</h1>;
   }
```

---

## ステップ3：データを送信する

ページの表示とデータ取得ができたら、次はフォームからデータを送信する機能を作ります。

### 問題：従来の方法は複雑

```tsx
// ❌ 従来の方法：onSubmitで手動処理
export default function CreateMemo() {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      await fetch('/api/memos', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, content }),
      });
      // 成功したらリダイレクト
      window.location.href = '/';
    } catch (error) {
      alert('エラーが発生しました');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input value={title} onChange={e => setTitle(e.target.value)} />
      <textarea value={content} onChange={e => setContent(e.target.value)} />
      <button disabled={loading}>
        {loading ? '送信中...' : '作成'}
      </button>
    </form>
  );
}
```

### 解決：Form + action関数

React Router v7では、`Form`コンポーネントと`action`関数を使います。

```tsx
import { Form, redirect } from "react-router";
import type { Route } from "./+types";

// フォーム送信時に実行される
export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;

  await createMemo({ title, content }); // 外部に書いてあるcreateMemo関数によってバックエンドAPIをたたいてもらう想定

  // リダイレクト
  return redirect("/");
}

// コンポーネントはシンプルに
export default function CreateMemo() {
  return (
    <Form method="post">
      <input name="title" />
      <textarea name="content" />
      <button type="submit">作成</button>
    </Form>
  );
}
```

**メリット:**
- コンポーネントがシンプル（ローディング状態不要）
- フォーム送信後、自動的にloaderが再実行される（データ再取得）
- プログレッシブエンハンスメント（JavaScriptなしでも動作）

### Formコンポーネントの基本

**1. method属性でHTTPメソッドを指定**
```tsx
<Form method="post">   {/* POST */}
<Form method="put">    {/* PUT */}
<Form method="delete"> {/* DELETE */}
```

**2. name属性は必須**
```tsx
// ❌ name属性がないとformDataで取得できない
<input type="text" />

// ✅ name属性をつける
<input name="title" type="text" />
```

**3. defaultValueで初期値を設定**
```tsx
// 非制御コンポーネント（推奨）
<input name="title" defaultValue={memo.title} />
<textarea name="content" defaultValue={memo.content} />
```

### FormDataの取得

```tsx
export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();

  // 方法1: 個別に取得
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;

  // 方法2: 一括で取得
  const data = Object.fromEntries(formData);
  // { title: "...", content: "..." }

  return null;
}
```

### 使用例：メモ編集

```tsx
// app/routes/edit_memo.tsx
import { Form, redirect } from "react-router";
import type { Route } from "./+types";
import { getMemoById, editMemo } from "~/api/memos";

// 編集画面用のデータ取得
export async function loader({ params }: Route.LoaderArgs) {
  const memo = await getMemoById(Number(params.id));
  return memo;
}

// フォーム送信時の処理
export async function action({ params, request }: Route.ActionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;

  await editMemo(Number(params.id), title, content);

  // 詳細ページにリダイレクト
  return redirect(`/memos/${params.id}`);
}

export default function EditMemo({ loaderData }: Route.ComponentProps) {
  const memo = loaderData;

  return (
    <Form method="put">
      <input name="title" defaultValue={memo.title} />
      <textarea name="content" defaultValue={memo.content} />
      <button type="submit">保存</button>
    </Form>
  );
}
```

### redirectとjson

action関数から返す値の種類：

**1. redirect() - 別のページに遷移**
```tsx
export async function action() {
  await saveData();
  return redirect("/success");
}
```

**2. json() - JSONレスポンスを返す（エラー時など）**
```tsx
export async function action() {
  try {
    await saveData();
    return redirect("/");
  } catch (error) {
    return json(
      { error: "保存に失敗しました" },
      { status: 400 }
    );
  }
}
```

**3. null - 何も返さない（同じページに留まる）**
```tsx
export async function action() {
  await saveData();
  return null; // 現在のページに留まる
}
```

### 自動再検証：データの自動更新

React Router v7の重要な機能の1つが**自動再検証**です。

**問題：データ送信後、画面を手動で更新する必要がある？**

従来の方法では、フォーム送信後に手動でデータを再取得する必要がありました。

**解決：React Routerが自動でloaderを再実行**

`Form`コンポーネントで送信すると、自動的にloader関数が再実行されます。

```tsx
// データが自動で最新になる！
export async function loader() {
  const memos = await getMemos();
  return { memos };
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  await createMemo(formData);
  return redirect("/"); // リダイレクト先のloaderが自動実行される
}

export default function MemoList({ loaderData }: Route.ComponentProps) {
  const { memos } = loaderData; // 常に最新のデータ

  return (
    <div>
      <Form method="post" action="/create">
        {/* フォーム */}
      </Form>
      {/* メモ一覧は自動更新される */}
      {memos.map(memo => <div key={memo.id}>{memo.title}</div>)}
    </div>
  );
}
```

**自動再検証のルール:**

| HTTPメソッド | 再検証 | 理由 |
|------------|-------|------|
| POST | ✅ する | データを作成（サーバーの状態が変わる） |
| PUT | ✅ する | データを更新（サーバーの状態が変わる） |
| DELETE | ✅ する | データを削除（サーバーの状態が変わる） |
| GET | ❌ しない | データを取得するだけ（サーバーの状態は変わらない） |

**メリット:**
- 手動でデータを再取得する必要がない
- 画面とサーバーのデータが常に同期される
- コードがシンプルになる

### action属性のパス解決

`Form`コンポーネントの`action`属性は、相対パスとして解決されます。

**基本的な動作:**

```tsx
// 現在のURL: /memos/123
<Form action="edit" method="post">
  {/* 送信先: /memos/123/edit */}
</Form>

<Form action="delete" method="delete">
  {/* 送信先: /memos/123/delete */}
</Form>
```

**action属性を省略した場合:**

```tsx
// 現在のURL: /memos/123
<Form method="post">
  {/* 送信先: /memos/123（現在のURL） */}
</Form>
```

**絶対パスを指定:**

```tsx
<Form action="/memos/create" method="post">
  {/* 送信先: /memos/create（どのページからでも） */}
</Form>
```

**利点:**
- コンポーネントの再利用性が高まる
- `memoId`のような動的な値を意識せず、静的に`action="edit"`と書ける
- コードがシンプルになる

### Tips：Formのaction属性とバックエンドAPIの関係

**混乱しやすいポイント：action属性は何を指定するのか？**

React Routerには2つのレイヤーがあります：

```
[Form action属性] → [action関数] → [バックエンドAPI]
      ↑                ↑                ↑
 どのルート?      FormData処理      HTTPリクエスト
```

**具体例：**

```tsx
// routes/edit_memo.tsx - ルーティング層
export async function action({ params, request }: Route.ActionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;

  // API層の関数を呼び出す
  await editMemo(Number(params.id), title, content);

  return redirect(`/memos/${params.id}`);
}

// api/memos.ts - API通信層
export async function editMemo(id: number, title: string, content: string) {
  return await fetch(`http://localhost:8080/memos/${id}`, {
    method: 'PUT',
    body: JSON.stringify({ title, content }),
  }).then(res => res.json());
}
```

**ポイント：**
- **`<Form action="...">`**: フロントエンド内のどのaction関数を呼ぶか（省略時は現在のルート）
- **API関数（`editMemo`等）**: 実際のバックエンドAPIへのHTTPリクエスト
- **この分離が推奨設計**：再利用性、テスト容易性、関心の分離

### useFetcher：ナビゲーションなしのデータ送信

**問題：お気に入りボタンを押した時、ページ遷移したくない**

通常の`Form`は送信後にページ遷移します。しかし、「データだけ更新してページには留まりたい」場合があります。

**解決：useFetcher**

```tsx
import { useFetcher } from "react-router";

function FavoriteButton({ contact }: { contact: Contact }) {
  const fetcher = useFetcher();

  return (
    <fetcher.Form method="post" action="/contacts/favorite">
      <input type="hidden" name="contactId" value={contact.id} />
      <input type="hidden" name="favorite" value={contact.favorite ? "false" : "true"} />
      <button type="submit">
        {contact.favorite ? "★" : "☆"}
      </button>
    </fetcher.Form>
  );
}
```

**通常のFormとuseFetcherの違い:**

| 機能 | Form | useFetcher |
|------|------|-----------|
| ページ遷移 | する | **しない** |
| URLの変更 | する | **しない** |
| データ更新 | する | する |
| 履歴への追加 | する | **しない** |
| ローディング状態 | navigation.state | **fetcher.state** |

**使い分けの基準:**

| シチュエーション | 使用するもの |
|---------------|-----------|
| フォーム送信後に別ページへ | `Form` |
| データ更新だけしたい | `useFetcher` |
| お気に入りボタン | `useFetcher` |
| いいねボタン | `useFetcher` |
| 削除ボタン（モーダル内） | `useFetcher` |

### Optimistic UI：楽観的UI更新

**問題：ネットワークが遅いと、ボタンを押しても反応が遅い**

通常、サーバーの応答を待ってからUIを更新します。これだとネットワークが遅い時に反応が遅く感じます。

**解決：Optimistic UI（楽観的UI）**

「サーバーの処理は成功する」と楽観的に考え、**先にUIを更新**します。

```tsx
import { useFetcher } from "react-router";

function FavoriteButton({ contact }: { contact: Contact }) {
  const fetcher = useFetcher();

  // fetcher.formDataから「送信予定の値」を取得
  const favorite = fetcher.formData
    ? fetcher.formData.get("favorite") === "true"
    : contact.favorite; // formDataがない場合は現在の値

  return (
    <fetcher.Form method="post">
      <button
        name="favorite"
        value={favorite ? "false" : "true"}
      >
        {favorite ? "★" : "☆"} {/* 即座に切り替わる！ */}
      </button>
    </fetcher.Form>
  );
}
```

**動作:**
1. ユーザーがボタンをクリック
2. `fetcher.formData`に「送信予定のデータ」が入る
3. UIが即座に更新される（★ ↔ ☆が切り替わる）
4. 裏でサーバーに送信される
5. サーバーの応答が返ってきたら、最終的なデータで更新

**メリット:**
- ユーザー体験が向上（即座に反応する）
- ネットワーク遅延の影響を受けにくい

**注意:**
- サーバーの処理が失敗した場合は、元に戻す必要がある
- 楽観的な更新が適切かどうか考える（削除など重要な操作には慎重に）

---

## ステップ4：ページ間を移動する

データの表示・送信ができたら、次はボタンやリンクでページを移動できるようにします。

### 問題：どうやってページ遷移する？

SPAでは、通常の`<a>`タグを使うとページ全体がリロードされてしまいます。React Routerを使って、リロードせずに画面を切り替える必要があります。

### 解決策1：Linkコンポーネント

通常のリンク（クリックでページ遷移）には`Link`を使います。

```tsx
import { Link } from "react-router";

function Navigation() {
  return (
    <nav>
      <Link to="/">ホーム</Link>
      <Link to="/about">About</Link>
      <Link to="/contact">お問い合わせ</Link>
    </nav>
  );
}
```

**動的なリンク:**
```tsx
function MemoList({ memos }: { memos: Memo[] }) {
  return (
    <div>
      {memos.map(memo => (
        <Link key={memo.id} to={`/memos/${memo.id}`}>
          {memo.title}
        </Link>
      ))}
    </div>
  );
}
```

### 解決策2：useNavigate フック

JavaScriptで動的にページ遷移したい場合は`useNavigate`を使います。

```tsx
import { useNavigate } from "react-router";

function MemoCard({ memo }: { memo: Memo }) {
  const navigate = useNavigate();

  return (
    <div onClick={() => navigate(`/memos/${memo.id}`)}>
      <h2>{memo.title}</h2>
      <p>{memo.content}</p>
    </div>
  );
}
```

**よく使うパターン:**
```tsx
function MyComponent() {
  const navigate = useNavigate();

  return (
    <>
      {/* 指定URLに遷移 */}
      <button onClick={() => navigate("/")}>
        ホームへ
      </button>

      {/* 前のページに戻る */}
      <button onClick={() => navigate(-1)}>
        戻る
      </button>

      {/* 次のページに進む */}
      <button onClick={() => navigate(1)}>
        進む
      </button>

      {/* 条件付き遷移 */}
      <button onClick={() => {
        if (isValid) {
          navigate("/success");
        } else {
          navigate("/error");
        }
      }}>
        送信
      </button>
    </>
  );
}
```

### Link vs useNavigate

| 用途 | 使用するもの | 理由 |
|------|-----------|------|
| ナビゲーションメニュー | `Link` | SEO、アクセシビリティ |
| リスト項目のリンク | `Link` | 右クリックで新しいタブで開ける |
| ボタンクリックで遷移 | `useNavigate` | ボタンとして機能 |
| 条件付き遷移 | `useNavigate` | JavaScriptロジックが必要 |
| カード全体をクリック可能に | `useNavigate` | クリック範囲が広い |

### 使用例：イベント伝播の制御

カード全体がクリック可能だが、中のボタンは別の処理をしたい場合：

```tsx
function MemoCard({ memo }: { memo: Memo }) {
  const navigate = useNavigate();

  return (
    <div onClick={() => navigate(`/memos/${memo.id}`)}>
      <h2>{memo.title}</h2>
      <p>{memo.content}</p>

      {/* 編集ボタン：カードのクリックを止める */}
      <button onClick={(e) => {
        e.stopPropagation(); // 親のonClickを止める
        navigate(`/memos/${memo.id}/edit`);
      }}>
        編集
      </button>

      {/* 削除ボタン */}
      <button onClick={(e) => {
        e.stopPropagation();
        if (confirm('削除しますか？')) {
          deleteMemo(memo.id);
        }
      }}>
        削除
      </button>
    </div>
  );
}
```

### NavLink：アクティブなリンクのスタイリング

**問題：現在のページのナビゲーションリンクをハイライトしたい**

サイドバーやヘッダーのナビゲーションで、「今どのページにいるか」を視覚的に示したいことがあります。

**解決：NavLink**

`NavLink`は、現在のURLと一致する時に特別なスタイルを適用できます。

```tsx
import { NavLink } from "react-router";

function Navigation() {
  return (
    <nav>
      <NavLink
        to="/"
        className={({ isActive, isPending }) =>
          isActive ? "active" : isPending ? "pending" : ""
        }
      >
        ホーム
      </NavLink>

      <NavLink
        to="/about"
        className={({ isActive }) =>
          isActive ? "text-blue-600 font-bold" : "text-gray-600"
        }
      >
        About
      </NavLink>
    </nav>
  );
}
```

**状態の種類:**
- `isActive`: 現在のURLと一致している
- `isPending`: 遷移中（データ読み込み中）

**スタイルの適用方法:**

**1. className（関数）**
```tsx
<NavLink
  to="/profile"
  className={({ isActive }) => isActive ? "active" : ""}
>
  プロフィール
</NavLink>
```

**2. style（関数）**
```tsx
<NavLink
  to="/profile"
  style={({ isActive }) => ({
    color: isActive ? "blue" : "black",
    fontWeight: isActive ? "bold" : "normal",
  })}
>
  プロフィール
</NavLink>
```

**3. children（関数）**
```tsx
<NavLink to="/profile">
  {({ isActive }) => (
    <span className={isActive ? "active" : ""}>
      {isActive && "→ "}
      プロフィール
    </span>
  )}
</NavLink>
```

**使用例：サイドバーナビゲーション**
```tsx
function Sidebar({ contacts }: { contacts: Contact[] }) {
  return (
    <nav>
      {contacts.map(contact => (
        <NavLink
          key={contact.id}
          to={`/contacts/${contact.id}`}
          className={({ isActive, isPending }) =>
            isActive
              ? "bg-blue-500 text-white"
              : isPending
              ? "bg-gray-200 text-gray-600"
              : "text-gray-700"
          }
        >
          {contact.name}
        </NavLink>
      ))}
    </nav>
  );
}
```

### 履歴スタックの管理

**問題：リアルタイム検索で「戻る」ボタンを何度も押す必要がある**

リアルタイム検索では、文字を入力するたびにURLが変わります。すると、履歴スタックに大量のエントリーが積まれてしまいます。

```
検索前: /
"a"入力: /?q=a
"ap"入力: /?q=ap
"app"入力: /?q=app
"appl"入力: /?q=appl
"apple"入力: /?q=apple

→ 戻るボタンを5回押さないと検索前に戻れない！
```

**解決：replace オプション**

`replace`オプションを使うと、履歴に新しいエントリーを追加せず、現在のエントリーを置き換えます。

```tsx
import { Form, useSubmit } from "react-router";

function SearchForm({ q }: { q: string | null }) {
  const submit = useSubmit();

  return (
    <Form onChange={(event) => {
      const isFirstSearch = (q === null);
      submit(event.currentTarget, {
        replace: !isFirstSearch, // 初回以外は置き換え
      });
    }}>
      <input name="q" defaultValue={q || ""} />
    </Form>
  );
}
```

**動作:**
1. **初回検索**: `replace: false` → 新しい履歴エントリーを追加
2. **継続検索**: `replace: true` → 現在のエントリーを置き換え

**結果:**
```
検索前: /
"apple"入力: /?q=apple

→ 戻るボタン1回で検索前に戻れる！
```

**replaceのその他の用途:**

**1. useNavigateでの使用**
```tsx
const navigate = useNavigate();

// 履歴に追加せず置き換え
navigate("/login", { replace: true });
```

**2. redirectでの使用**
```tsx
export async function action() {
  // ログイン成功後、ログインページに戻れないようにする
  return redirect("/dashboard", { replace: true });
}
```

---

## ステップ5：URLからパラメータを取得する

最後に、URLに含まれる動的なパラメータ（IDなど）を取得する方法を学びます。

### 問題：URLのIDを使いたい

例えば、`/memos/123`というURLにアクセスした時、`123`という値を使ってデータを取得したいです。

### URLパラメータの定義

routes.tsで`:id`のように定義します。

```tsx
// app/routes.ts
export default [
    route("memos/:id", "routes/memo.tsx"),
    // :id が動的パラメータ
] satisfies RouteConfig;
```

### loader関数でパラメータを取得

```tsx
// app/routes/memo.tsx
import type { Route } from "./+types";
import { getMemoById } from "~/api/memos";

export async function loader({ params }: Route.LoaderArgs) {
  // params.id でURLのパラメータを取得
  const id = Number(params.id);
  const memo = await getMemoById(id);
  return memo;
}

export default function MemoDetail({ loaderData }: Route.ComponentProps) {
  const memo = loaderData;
  return (
    <div>
      <h1>{memo.title}</h1>
      <p>{memo.content}</p>
    </div>
  );
}
```

### action関数でパラメータを取得

```tsx
export async function action({ params, request }: Route.ActionArgs) {
  const id = Number(params.id);
  const formData = await request.formData();

  // パラメータとフォームデータの両方を使う
  await updateMemo(id, {
    title: formData.get("title") as string,
    content: formData.get("content") as string,
  });

  return redirect(`/memos/${id}`);
}
```

### パラメータの型

URLパラメータは**常に文字列**です。数値として使う場合は変換が必要です。

```tsx
export async function loader({ params }: Route.LoaderArgs) {
  // params.id は "123"（文字列）

  // 数値に変換
  const id = Number(params.id);

  // より安全な変換
  const id = parseInt(params.id, 10);
  if (isNaN(id)) {
    throw new Response("Invalid ID", { status: 400 });
  }

  return await getMemoById(id);
}
```

### 複数のパラメータ

```tsx
// routes.ts
route("users/:userId/posts/:postId", "routes/post.tsx")

// loader関数
export async function loader({ params }: Route.LoaderArgs) {
  const userId = Number(params.userId);
  const postId = Number(params.postId);

  const post = await getPost(userId, postId);
  return post;
}
```

### 使用例：完全な CRUD

メモアプリの完全な例：

**routes.ts:**
```tsx
export default [
    index("routes/index.tsx"),                    // 一覧
    route("memos/new", "routes/new_memo.tsx"),    // 新規作成
    route("memos/:id", "routes/memo.tsx"),        // 詳細
    route("memos/:id/edit", "routes/edit_memo.tsx"), // 編集
] satisfies RouteConfig;
```

**詳細ページ（routes/memo.tsx）:**
```tsx
import { useNavigate } from "react-router";
import type { Route } from "./+types";
import { getMemoById } from "~/api/memos";

export async function loader({ params }: Route.LoaderArgs) {
  const memo = await getMemoById(Number(params.id));
  return memo;
}

export default function MemoDetail({ loaderData }: Route.ComponentProps) {
  const navigate = useNavigate();
  const memo = loaderData;

  return (
    <div>
      <h1>{memo.title}</h1>
      <p>{memo.content}</p>

      <button onClick={() => navigate(-1)}>戻る</button>
      <button onClick={() => navigate(`/memos/${memo.id}/edit`)}>
        編集
      </button>
    </div>
  );
}
```

**編集ページ（routes/edit_memo.tsx）:**
```tsx
import { Form, redirect } from "react-router";
import type { Route } from "./+types";
import { getMemoById, editMemo } from "~/api/memos";

export async function loader({ params }: Route.LoaderArgs) {
  const memo = await getMemoById(Number(params.id));
  return memo;
}

export async function action({ params, request }: Route.ActionArgs) {
  const formData = await request.formData();
  await editMemo(Number(params.id), {
    title: formData.get("title") as string,
    content: formData.get("content") as string,
  });
  return redirect(`/memos/${params.id}`);
}

export default function EditMemo({ loaderData }: Route.ComponentProps) {
  const memo = loaderData;

  return (
    <Form method="put">
      <input name="title" defaultValue={memo.title} />
      <textarea name="content" defaultValue={memo.content} />
      <button type="submit">保存</button>
    </Form>
  );
}
```

---

## 高度なUI制御

ここまでで基本的なSPAが作れるようになりました。最後に、より良いユーザー体験のための高度なUI制御を学びます。

### useNavigation：グローバルなローディング状態

**問題：ページ遷移中にローディング表示をしたい**

ページ遷移中（loader関数の実行中）に、ユーザーにフィードバックを表示したいことがあります。

**解決：useNavigation**

```tsx
import { useNavigation, Outlet } from "react-router";

export default function Root() {
  const navigation = useNavigation();

  return (
    <div>
      <nav>{/* ナビゲーション */}</nav>

      {/* ローディング中はフェードアウト */}
      <div className={navigation.state === "loading" ? "opacity-50" : ""}>
        <Outlet />
      </div>

      {/* グローバルローディングインジケーター */}
      {navigation.state === "loading" && (
        <div className="loading-bar">読み込み中...</div>
      )}
    </div>
  );
}
```

**navigation.stateの値:**

| 状態 | 意味 | 表示すべきもの |
|------|------|-------------|
| `"idle"` | 通常状態 | 何もしない |
| `"loading"` | ページ遷移中（loaderが実行中） | ローディング表示 |
| `"submitting"` | フォーム送信中（actionが実行中） | 送信中表示 |

**より高度な制御：検索中は除外**

リアルタイム検索などで、検索中はメインコンテンツをフェードアウトさせたくない場合：

```tsx
import { useNavigation, Outlet } from "react-router";

export default function Root() {
  const navigation = useNavigation();

  // 検索中かどうかを判定
  const searching =
    navigation.location &&
    new URLSearchParams(navigation.location.search).has("q");

  return (
    <div>
      {/* 検索中はフェードアウトしない */}
      <div
        className={
          navigation.state === "loading" && !searching
            ? "opacity-50"
            : ""
        }
      >
        <Outlet />
      </div>
    </div>
  );
}
```

### useNavigation：フォーム送信中の制御

```tsx
import { useNavigation, Form } from "react-router";

function CreateMemo() {
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";

  return (
    <Form method="post">
      <input name="title" disabled={isSubmitting} />
      <textarea name="content" disabled={isSubmitting} />
      <button type="submit" disabled={isSubmitting}>
        {isSubmitting ? "送信中..." : "作成"}
      </button>
    </Form>
  );
}
```

### useNavigation.formDataで送信内容を取得

送信中のフォームデータを取得して、Optimistic UIに使えます。

```tsx
import { useNavigation, Form } from "react-router";

function TodoList({ todos }: { todos: Todo[] }) {
  const navigation = useNavigation();

  // 送信中の新しいTodoを取得
  const newTodo = navigation.formData?.get("todo") as string | undefined;

  return (
    <div>
      <ul>
        {todos.map(todo => <li key={todo.id}>{todo.text}</li>)}

        {/* 送信中の新しいTodoを先に表示（Optimistic UI） */}
        {newTodo && (
          <li className="opacity-50">{newTodo}</li>
        )}
      </ul>

      <Form method="post">
        <input name="todo" />
        <button type="submit">追加</button>
      </Form>
    </div>
  );
}
```

### ローディング状態の使い分け

| フック | スコープ | 用途 |
|--------|---------|------|
| `useNavigation` | **グローバル** | アプリ全体のローディング表示（root.tsx） |
| `useFetcher` | **ローカル** | 特定のコンポーネントのローディング（いいねボタンなど） |

**useNavigation（グローバル）:**
```tsx
// root.tsx
const navigation = useNavigation();

// アプリ全体のローディングバーを表示
{navigation.state === "loading" && <LoadingBar />}
```

**useFetcher（ローカル）:**
```tsx
// 特定のコンポーネント
const fetcher = useFetcher();

// このボタンだけのローディング状態
<button disabled={fetcher.state === "submitting"}>
  {fetcher.state === "submitting" ? "送信中..." : "いいね"}
</button>
```

### ベストプラクティス

**1. グローバルローディングはroot.tsxに**
```tsx
// app/root.tsx
export default function Root() {
  const navigation = useNavigation();

  return (
    <>
      {/* グローバルローディングバー */}
      {navigation.state !== "idle" && <ProgressBar />}
      <Outlet />
    </>
  );
}
```

**2. フォームのボタンは送信中disable**
```tsx
<button type="submit" disabled={navigation.state === "submitting"}>
  {navigation.state === "submitting" ? "送信中..." : "送信"}
</button>
```

**3. Optimistic UIで即座にフィードバック**
```tsx
// ユーザーの操作に即座に反応
const optimisticValue = fetcher.formData
  ? fetcher.formData.get("value")
  : currentValue;
```

---

## まとめ：React Router v7の開発フロー

React Router v7を使ったSPA開発は、以下の流れで進めます：

### 1. ルート定義（routes.ts）
```tsx
export default [
    index("routes/index.tsx"),
    route("items/:id", "routes/item.tsx"),
] satisfies RouteConfig;
```

### 2. データ取得（loader）
```tsx
export async function loader({ params }: Route.LoaderArgs) {
  const data = await fetchData(params.id);
  return data;
}
```

### 3. データ送信（Form + action）
```tsx
export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  await saveData(formData);
  return redirect("/");
}
```

### 4. ページ遷移（useNavigate / Link）
```tsx
const navigate = useNavigate();
<button onClick={() => navigate("/path")}>移動</button>
```

### 従来の方法との比較

| 機能 | React Router v7 | 従来の方法 |
|------|----------------|----------|
| データ取得 | loader関数 | useEffect + useState |
| ローディング管理 | 自動 | 手動でstate管理 |
| フォーム送信 | Form + action | onSubmit + fetch |
| データ再検証 | 自動 | 手動でリロード |
| 型安全性 | 自動生成 | 手動で定義 |

React Router v7は、従来の方法と比べて**宣言的で簡潔**なコードを書くことができます。useEffectやuseStateを使う機会が大幅に減り、コンポーネントがシンプルになります。
