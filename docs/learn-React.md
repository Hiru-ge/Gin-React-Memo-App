# React 学習ノート

このドキュメントでは、Reactを使った開発の流れに沿って、必要な概念を順に学んでいきます。

## 目次

1. [Reactとは](#reactとは)
2. [最初のステップ：静的なページを作る](#最初のステップ静的なページを作る)
3. [データを表示する](#データを表示する)
4. [ユーザー操作に反応する](#ユーザー操作に反応する)
5. [動的な値を扱う（State）](#動的な値を扱うstate)
6. [外部データを扱う（useEffect）](#外部データを扱うuseeffect)
7. [プロジェクト構造とコンポーネントの配置](#プロジェクト構造とコンポーネントの配置)

---

## Reactとは

ReactはFacebook（Meta）が開発したJavaScriptライブラリで、Webアプリケーションのユーザーインターフェース（UI）を構築するために使用されます。

**Reactの3つの特徴:**
- **コンポーネントベース**: UIを独立した再利用可能な部品（コンポーネント）に分割できる
- **宣言的**: 「どう見えるべきか」を書くだけで、Reactが効率的に更新してくれる
- **仮想DOM**: 実際のDOMを直接操作せず、効率的に画面を更新する仕組み

---

## 最初のステップ：静的なページを作る

Reactでアプリを作る第一歩は、HTMLのようにページの見た目を定義することです。

### コンポーネント：UIの部品

Reactでは、UIを「コンポーネント」という関数で定義します。

```tsx
// 最もシンプルなコンポーネント
function Welcome() {
  return <h1>Hello, React!</h1>;
}
```

コンポーネントは通常の関数ですが、JSX（後述）を返します。

**命名規則:**
- コンポーネント名は**必ず大文字で始める**（`Welcome`, `UserProfile`, `Button`）
- 小文字で始めると、HTMLタグとして認識されてしまう

```tsx
function welcome() {  // ❌ 小文字で始めるとエラー
  return <h1>Hello</h1>;
}

function Welcome() {  // ✅ 大文字で始める
  return <h1>Hello</h1>;
}
```

### JSX：JavaScriptの中にHTMLを書く

JSXは、JavaScriptの中にHTMLのような構文を書くことができる拡張構文です。

```tsx
function UserProfile() {
  return (
    <div>
      <h1>ユーザープロフィール</h1>
      <p>ここにプロフィール情報を表示します</p>
    </div>
  );
}
```

**JSXの基本ルール:**

**1. 単一のルート要素を返す**
```tsx
// ❌ 複数のルート要素はエラー
function Bad() {
  return (
    <h1>Title</h1>
    <p>Content</p>
  );
}

// ✅ divでラップする
function Good() {
  return (
    <div>
      <h1>Title</h1>
      <p>Content</p>
    </div>
  );
}

// ✅ Fragment（<>）を使うと余分なDOMを作らない
function Better() {
  return (
    <>
      <h1>Title</h1>
      <p>Content</p>
    </>
  );
}
```

**2. すべてのタグを閉じる**
```tsx
// ❌ HTMLでは<img>と書けるが、JSXではエラー
<img src="image.png">

// ✅ 自己閉じタグにする
<img src="image.png" />
<input type="text" />
<br />
```

**3. 属性名はキャメルケース**

HTMLとJSXで属性名が異なることに注意：

```tsx
// HTML → JSX
// class → className
// onclick → onClick
// tabindex → tabIndex
// for → htmlFor

<div className="container" onClick={handleClick}>
  <label htmlFor="name">名前</label>
  <input id="name" type="text" />
</div>
```

### JSX内でJavaScriptを使う

`{}` の中にJavaScript式を書けます。これがJSXの強力な機能です。

```tsx
function Greeting() {
  const name = "太郎";
  const age = 25;

  return (
    <div>
      <h1>こんにちは、{name}さん</h1>
      <p>あなたは{age}歳です</p>
      <p>来年は{age + 1}歳になります</p>
    </div>
  );
}
```

**使用例:**
```tsx
function MemoCard() {
  const memo = {
    title: "買い物リスト",
    createdAt: new Date("2024-01-15"),
  };

  return (
    <div>
      <h2>{memo.title}</h2>
      <p>作成日: {memo.createdAt.toLocaleDateString()}</p>
    </div>
  );
}
```

---

## データを表示する

静的なページが作れたら、次は動的なデータを表示します。

### Props：親から子へデータを渡す

コンポーネントを再利用可能にするには、外部からデータを受け取る必要があります。それが**Props**です。

**問題：同じカードを何度も書くのは面倒**
```tsx
function MemoList() {
  return (
    <div>
      <div className="card">
        <h2>買い物リスト</h2>
        <p>牛乳、卵、パンを買う</p>
      </div>
      <div className="card">
        <h2>TODO</h2>
        <p>レポートを提出する</p>
      </div>
      <div className="card">
        <h2>アイデア</h2>
        <p>新しいアプリのアイデア</p>
      </div>
    </div>
  );
}
```

**解決：Propsでカードコンポーネントを再利用**
```tsx
// カードコンポーネント：Propsでデータを受け取る
function MemoCard({ title, content }: { title: string; content: string }) {
  return (
    <div className="card">
      <h2>{title}</h2>
      <p>{content}</p>
    </div>
  );
}

// 使う側：Propsでデータを渡す
function MemoList() {
  return (
    <div>
      <MemoCard title="買い物リスト" content="牛乳、卵、パンを買う" />
      <MemoCard title="TODO" content="レポートを提出する" />
      <MemoCard title="アイデア" content="新しいアプリのアイデア" />
    </div>
  );
}
```

**TypeScriptでの型定義（推奨）:**
```tsx
// 型を別で定義すると読みやすい
type MemoCardProps = {
  title: string;
  content: string;
  createdAt?: Date;  // ?は省略可能
};

function MemoCard({ title, content, createdAt }: MemoCardProps) {
  return (
    <div className="card">
      <h2>{title}</h2>
      <p>{content}</p>
      {createdAt && <p>{createdAt.toLocaleDateString()}</p>}
    </div>
  );
}
```

### リストのレンダリング：配列データを表示する

実際のアプリでは、配列のデータを表示することが多いです。

**問題：APIから取得したデータ（配列）を表示したい**
```tsx
function MemoList() {
  // APIから取得したデータ（配列）
  const memos = [
    { id: 1, title: '買い物リスト', content: '牛乳、卵' },
    { id: 2, title: 'TODO', content: 'レポート提出' },
    { id: 3, title: 'アイデア', content: 'アプリ開発' },
  ];

  // どうやって表示する？
}
```

**解決：map()を使う**
```tsx
function MemoList() {
  const memos = [
    { id: 1, title: '買い物リスト', content: '牛乳、卵' },
    { id: 2, title: 'TODO', content: 'レポート提出' },
    { id: 3, title: 'アイデア', content: 'アプリ開発' },
  ];

  return (
    <div>
      {memos.map((memo) => (
        <MemoCard
          key={memo.id}
          title={memo.title}
          content={memo.content}
        />
      ))}
    </div>
  );
}
```

**key属性が必須な理由:**

Reactは`key`を使って、どの要素が変更/追加/削除されたかを識別します。

```tsx
// ❌ keyがないと警告が出る
{memos.map((memo) => (
  <MemoCard title={memo.title} content={memo.content} />
))}

// ❌ インデックスをkeyにする（順序が変わる可能性がある場合は避ける）
{memos.map((memo, index) => (
  <MemoCard key={index} title={memo.title} content={memo.content} />
))}

// ✅ 一意のIDをkeyにする
{memos.map((memo) => (
  <MemoCard key={memo.id} title={memo.title} content={memo.content} />
))}
```

### 条件付きレンダリング：条件に応じて表示を変える

**問題：データがない時に「メモがありません」と表示したい**

**解決策1：三項演算子**
```tsx
function MemoList({ memos }: { memos: Memo[] }) {
  return (
    <div>
      {memos.length === 0 ? (
        <p>メモがありません</p>
      ) : (
        <div>
          {memos.map(memo => <MemoCard key={memo.id} {...memo} />)}
        </div>
      )}
    </div>
  );
}
```

**解決策2：&& 演算子（条件がtrueの時だけ表示）**
```tsx
function MemoList({ memos }: { memos: Memo[] }) {
  return (
    <div>
      {memos.length === 0 && <p>メモがありません</p>}
      {memos.length > 0 && (
        <div>
          {memos.map(memo => <MemoCard key={memo.id} {...memo} />)}
        </div>
      )}
    </div>
  );
}
```

**解決策3：Early Return（推奨：読みやすい）**
```tsx
function MemoList({ memos }: { memos: Memo[] }) {
  // 空の場合は早期リターン
  if (memos.length === 0) {
    return <p>メモがありません</p>;
  }

  // 通常の表示
  return (
    <div>
      {memos.map(memo => <MemoCard key={memo.id} {...memo} />)}
    </div>
  );
}
```

---

## ユーザー操作に反応する

静的なデータを表示できたら、次はユーザーの操作に反応するようにします。

### イベントハンドリング

**問題：ボタンをクリックした時に何か処理をしたい**

**解決：onClickイベントハンドラ**
```tsx
function DeleteButton() {
  // イベントハンドラ関数を定義
  const handleClick = () => {
    console.log('削除ボタンがクリックされました');
  };

  return (
    <button onClick={handleClick}>
      削除
    </button>
  );
}
```

**インラインでも書ける:**
```tsx
function DeleteButton() {
  return (
    <button onClick={() => console.log('削除')}>
      削除
    </button>
  );
}
```

**注意：関数を呼び出さないこと**
```tsx
// ❌ これは間違い：レンダリング時に実行されてしまう
<button onClick={handleClick()}>削除</button>

// ✅ 正しい：関数への参照を渡す
<button onClick={handleClick}>削除</button>

// ✅ アロー関数で包む
<button onClick={() => handleClick()}>削除</button>
```

### イベントオブジェクト

**問題：入力フォームの値を取得したい**

```tsx
function SearchForm() {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log('入力値:', e.target.value);
  };

  return <input type="text" onChange={handleChange} />;
}
```

### イベントの伝播を制御する

**問題：カード全体がクリック可能だが、その中のボタンをクリックした時は別の処理をしたい**

```tsx
function MemoCard({ memo }: { memo: Memo }) {
  return (
    <div onClick={() => console.log('カードをクリック')}>
      <h2>{memo.title}</h2>
      <button onClick={(e) => {
        e.stopPropagation(); // 親のonClickを止める
        console.log('削除ボタンをクリック');
      }}>
        削除
      </button>
    </div>
  );
}
```

**よく使うイベントメソッド:**
- `e.stopPropagation()`: イベントの伝播を止める（親要素のイベントを発火させない）
- `e.preventDefault()`: ブラウザのデフォルト動作を止める（例：フォーム送信、リンククリック）

### 主なイベント一覧

| イベント | 発火タイミング | 用途 |
|---------|-------------|------|
| `onClick` | クリック時 | ボタン、カードのクリック |
| `onChange` | 入力値変更時 | input, textarea, select |
| `onSubmit` | フォーム送信時 | form要素 |
| `onMouseEnter` | マウスが要素に入った時 | ホバー効果 |
| `onMouseLeave` | マウスが要素から出た時 | ホバー効果 |
| `onFocus` | フォーカス時 | input要素 |
| `onBlur` | フォーカスが外れた時 | バリデーション |

### 命名規則とベストプラクティス

Reactでは、イベント関連の命名規則が慣習的に決まっています。

**イベント関連の命名:**
```tsx
function MyComponent() {
  // ハンドラ関数は handle で始める
  const handleClick = () => {
    console.log('clicked');
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('submitted');
  };

  return (
    <>
      {/* イベントpropsは on で始める */}
      <button onClick={handleClick}>クリック</button>
      <form onSubmit={handleSubmit}>...</form>
    </>
  );
}
```

**カスタムコンポーネントでのイベント:**
```tsx
// 子コンポーネント：onSomething という名前で受け取る
function CustomButton({ onClick }: { onClick: () => void }) {
  return <button onClick={onClick}>ボタン</button>;
}

// 親コンポーネント：handleSomething という名前で定義
function Parent() {
  const handleButtonClick = () => {
    console.log('clicked');
  };

  return <CustomButton onClick={handleButtonClick} />;
}
```

**命名規則まとめ:**
- **イベントprops**: `onSomething`（onClick, onSubmit, onChange など）
- **イベントハンドラ**: `handleSomething`（handleClick, handleSubmit, handleChange など）
- データの流れは常に「親→子」（Props経由）
- 子から親へは「Propsで渡された関数を呼ぶ」ことでイベントを伝える

---

## 動的な値を扱う（State）

これまでは固定のデータを表示していましたが、実際のアプリではユーザー操作によって値が変化します。

### なぜStateが必要か

**問題：通常の変数では画面が更新されない**
```tsx
function Counter() {
  let count = 0; // 通常の変数

  const increment = () => {
    count = count + 1;
    console.log(count); // コンソールには表示されるが...
  };

  return (
    <div>
      <p>カウント: {count}</p> {/* 画面は更新されない！ */}
      <button onClick={increment}>増やす</button>
    </div>
  );
}
```

**解決：useState を使う**
```tsx
import { useState } from 'react';

function Counter() {
  // [現在の値, 値を更新する関数] = useState(初期値)
  const [count, setCount] = useState(0);

  const increment = () => {
    setCount(count + 1); // これで画面が更新される！
  };

  return (
    <div>
      <p>カウント: {count}</p>
      <button onClick={increment}>増やす</button>
    </div>
  );
}
```

### useStateの基本パターン

```tsx
function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>カウント: {count}</p>
      <button onClick={() => setCount(count + 1)}>+1</button>
      <button onClick={() => setCount(count - 1)}>-1</button>
      <button onClick={() => setCount(0)}>リセット</button>
    </div>
  );
}
```

### 複数のStateを管理する

```tsx
function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <form>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit">ログイン</button>
    </form>
  );
}
```

### オブジェクトのState

**問題：関連する複数の値をまとめて管理したい**

```tsx
function UserProfile() {
  const [user, setUser] = useState({
    name: '',
    age: 0,
    email: '',
  });

  // オブジェクトを更新する時は、スプレッド構文で既存の値をコピー
  const updateName = (name: string) => {
    setUser({ ...user, name });
    // または
    setUser(prevUser => ({ ...prevUser, name }));
  };

  return (
    <input
      value={user.name}
      onChange={(e) => updateName(e.target.value)}
    />
  );
}
```

**重要：Stateは直接変更しない（イミュータビリティ）**
```tsx
// ❌ 直接変更はNG
user.name = '太郎';
setUser(user);

// ✅ 新しいオブジェクトを作る
setUser({ ...user, name: '太郎' });
```

### イミュータビリティ（不変性）の重要性

Reactでは、Stateを直接変更せず、**常に新しい値を作って更新**する必要があります。これを「イミュータビリティ（不変性）」と呼びます。

**なぜイミュータビリティが必要か:**
1. **正しい再レンダリング**: Reactは参照の変更で再レンダリングを判断する
2. **履歴管理**: 過去の状態を保持できる（Undo/Redo機能）
3. **パフォーマンス最適化**: 変更検出が高速
4. **予測可能性**: 意図しない副作用を防ぐ

**配列の不変更新パターン:**
```tsx
function TodoList() {
  const [items, setItems] = useState(['りんご', 'バナナ']);

  // ❌ 直接変更（pushは元の配列を変更する）
  const addItemBad = (item: string) => {
    items.push(item);
    setItems(items); // 参照が同じなので再レンダリングされない！
  };

  // ✅ 新しい配列を作る
  const addItemGood = (item: string) => {
    setItems([...items, item]); // スプレッド演算子で新しい配列
  };

  // ✅ 削除：filter で新しい配列を作る
  const removeItem = (index: number) => {
    setItems(items.filter((_, i) => i !== index));
  };

  // ✅ 更新：map で新しい配列を作る
  const updateItem = (index: number, newValue: string) => {
    setItems(items.map((item, i) => i === index ? newValue : item));
  };

  return <div>{/* ... */}</div>;
}
```

**オブジェクトの不変更新パターン:**
```tsx
function UserProfile() {
  const [user, setUser] = useState({
    name: '太郎',
    age: 25,
    address: {
      city: '東京',
      zip: '100-0001'
    }
  });

  // ❌ 直接変更
  const updateNameBad = () => {
    user.name = '花子';
    setUser(user); // 参照が同じなので再レンダリングされない！
  };

  // ✅ スプレッド演算子で新しいオブジェクトを作る
  const updateName = (name: string) => {
    setUser({ ...user, name });
  };

  // ✅ ネストしたオブジェクトも新しく作る
  const updateCity = (city: string) => {
    setUser({
      ...user,
      address: {
        ...user.address,
        city
      }
    });
  };

  return <div>{/* ... */}</div>;
}
```

**便利な不変更新ヘルパー:**
```tsx
// 配列の特定要素を更新
const updateAtIndex = (array: T[], index: number, newValue: T) => {
  return array.map((item, i) => i === index ? newValue : item);
};

// 配列から特定要素を削除
const removeAtIndex = (array: T[], index: number) => {
  return array.filter((_, i) => i !== index);
};

// オブジェクトの一部を更新
const updateObject = (obj: T, updates: Partial<T>) => {
  return { ...obj, ...updates };
};
```

### 配列のState

```tsx
function TodoList() {
  const [todos, setTodos] = useState<string[]>([]);

  // 配列に追加
  const addTodo = (todo: string) => {
    setTodos([...todos, todo]);
  };

  // 配列から削除
  const removeTodo = (index: number) => {
    setTodos(todos.filter((_, i) => i !== index));
  };

  // 配列を更新
  const updateTodo = (index: number, newTodo: string) => {
    setTodos(todos.map((todo, i) => i === index ? newTodo : todo));
  };

  return <div>{/* ... */}</div>;
}
```

### UIの状態管理

Stateの典型的な使い道：

```tsx
function MemoDeletModal() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button onClick={() => setIsOpen(true)}>削除</button>

      {isOpen && (
        <div className="modal">
          <p>本当に削除しますか？</p>
          <button onClick={() => {
            // 削除処理
            setIsOpen(false);
          }}>
            はい
          </button>
          <button onClick={() => setIsOpen(false)}>
            いいえ
          </button>
        </div>
      )}
    </>
  );
}
```

### モーダルの実装パターン

モーダル（ダイアログ）は、UIの状態管理の典型的な例です。段階的に実装方法を学びます。

#### パターン1: 基本的なモーダル（開閉のみ）

**問題：ボタンをクリックしたらモーダルを表示し、閉じるボタンで非表示にしたい**

```tsx
function ConfirmDialog() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <button onClick={() => setIsOpen(true)}>
        確認ダイアログを開く
      </button>

      {isOpen && (
        <div className="modal">
          <p>この操作を実行しますか？</p>
          <button onClick={() => setIsOpen(false)}>
            キャンセル
          </button>
          <button onClick={() => {
            // 何か処理を実行
            console.log('実行されました');
            setIsOpen(false);
          }}>
            実行
          </button>
        </div>
      )}
    </>
  );
}
```

**ポイント:**
- `isOpen` という boolean のStateで開閉を管理
- `isOpen && <Modal>` で条件付きレンダリング
- `true` の時だけモーダルが表示される

#### パターン2: データを渡すモーダル（コンポーネント分離）

**問題：複数の場所から同じモーダルを使いたい。また、削除対象のアイテムを渡したい**

**解決：モーダルを別コンポーネントに分離し、propsでデータと閉じる関数を渡す**

```tsx
// モーダルコンポーネント
type DeleteModalProps = {
  item: { id: number; name: string };
  onClose: () => void;
  onConfirm: () => void;
};

function DeleteModal({ item, onClose, onConfirm }: DeleteModalProps) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>アイテムの削除</h2>
        <p>{item.name} を削除しますか？</p>
        <div className="modal-buttons">
          <button onClick={onClose}>
            キャンセル
          </button>
          <button onClick={onConfirm}>
            削除
          </button>
        </div>
      </div>
    </div>
  );
}

// 使う側のコンポーネント
function ItemList() {
  const [deletingItem, setDeletingItem] = useState<{ id: number; name: string } | null>(null);

  const handleDelete = () => {
    if (deletingItem) {
      console.log(`削除: ${deletingItem.id}`);
      // 削除処理...
      setDeletingItem(null); // モーダルを閉じる
    }
  };

  return (
    <div>
      <button onClick={() => setDeletingItem({ id: 1, name: 'アイテム1' })}>
        アイテム1を削除
      </button>
      <button onClick={() => setDeletingItem({ id: 2, name: 'アイテム2' })}>
        アイテム2を削除
      </button>

      {/* deletingItemがnullでない時だけモーダルを表示 */}
      {deletingItem && (
        <DeleteModal
          item={deletingItem}
          onClose={() => setDeletingItem(null)}
          onConfirm={handleDelete}
        />
      )}
    </div>
  );
}
```

**ポイント:**
- `deletingItem` を `null` または `オブジェクト` で管理
  - `null` → モーダル非表示
  - `オブジェクト` → モーダル表示 + データを保持
- **条件付きレンダリングの仕組み:**
  ```tsx
  {deletingItem && <Modal item={deletingItem} />}
  ```
  - `deletingItem` が `null` → falsy → 何も表示されない
  - `deletingItem` が `オブジェクト` → truthy → `<Modal>` が表示される
- `onClick={(e) => e.stopPropagation()}` でモーダル内クリックが背景に伝わらない

#### パターン3: Form送信時に自動で閉じる（onSubmit使用）

**問題：フォーム送信ボタンを押したら、自動的にモーダルを閉じたい**

**解決：FormのonSubmitイベントでonCloseを呼ぶ**

```tsx
import { Form } from 'react-router';

type DeleteModalProps = {
  item: { id: number; name: string };
  onClose: () => void;
};

function DeleteModal({ item, onClose }: DeleteModalProps) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>アイテムの削除</h2>
        <p>{item.name} を削除しますか？</p>

        {/* onSubmitでモーダルを閉じる */}
        <Form method="post" action="/items/delete" onSubmit={onClose}>
          <input type="hidden" name="id" value={item.id} />
          <button type="button" onClick={onClose}>
            キャンセル
          </button>
          <button type="submit">
            削除
          </button>
        </Form>
      </div>
    </div>
  );
}
```

**ポイント:**
- **Form送信の流れ:**
  1. ユーザーが削除ボタンをクリック
  2. `onSubmit`イベントが発火 → `onClose()`が実行される → **モーダルが閉じる**
  3. Form送信処理が継続される
  4. actionが実行される（削除処理）
  5. redirect()でページ遷移（または同じページの再読み込み）

- **メリット:**
  - シンプルで理解しやすい
  - useEffectやuseFetcherが不要
  - フォーム送信と同時にモーダルが閉じるのでUXが良い

- **注意点:**
  - 同じページにredirectする場合（例: `redirect("/")`で同じページに戻る）でも、loaderが再実行されてデータが更新されるため、親コンポーネントの状態もリセットされる
  - エラーハンドリングが必要な場合は、後述のパターン4を検討する

#### モーダルのオーバーレイとクリック制御

モーダルの背景（オーバーレイ）をクリックしたら閉じる、という挙動の実装：

```tsx
function Modal({ onClose }: { onClose: () => void }) {
  return (
    // 背景オーバーレイ：クリックで閉じる
    <div
      className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center"
      onClick={onClose}
    >
      {/* モーダル本体：クリックイベントを止める */}
      <div
        className="bg-white rounded-lg p-6 max-w-md"
        onClick={(e) => e.stopPropagation()}
      >
        <h2>モーダル</h2>
        <button onClick={onClose}>閉じる</button>
      </div>
    </div>
  );
}
```

**イベント伝播の仕組み:**
1. モーダル内部をクリック → `onClick={(e) => e.stopPropagation()}` でイベント伝播を止める
2. イベントが親（背景オーバーレイ）に伝わらない → `onClose` は呼ばれない
3. 背景をクリック → 直接 `onClick={onClose}` が呼ばれる → モーダルが閉じる

**まとめ:**
- モーダルの開閉は `useState` で管理
- `null` または `オブジェクト` で状態を管理すると、データも一緒に保持できる
- 条件付きレンダリング `{state && <Modal>}` で表示/非表示を切り替え
- Form送信時に自動で閉じるには `onSubmit={onClose}` を使う（シンプルで推奨）
- `e.stopPropagation()` でクリックイベントの伝播を制御

---

## 外部データを扱う（useEffect）

Stateで値を管理できるようになりましたが、コンポーネントの外部（API、ブラウザAPIなど）とやり取りする必要があります。

### なぜuseEffectが必要か

**問題：コンポーネントがレンダリングされた後に何か処理をしたい**

- APIからデータを取得
- ブラウザのタイトルを変更
- タイマーを設定
- イベントリスナーを登録

これらは「副作用（Side Effect）」と呼ばれ、レンダリングの外で実行する必要があります。

### useEffectの基本

```tsx
import { useEffect } from 'react';

function Example() {
  useEffect(() => {
    console.log('コンポーネントがレンダリングされました');
  });

  return <div>Example</div>;
}
```

### 依存配列：いつ実行するか制御する

```tsx
function UserProfile({ userId }: { userId: number }) {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // userIdが変わるたびに実行される
    fetch(`/api/users/${userId}`)
      .then(res => res.json())
      .then(data => setUser(data));
  }, [userId]); // 依存配列

  return <div>{user?.name}</div>;
}
```

**依存配列のパターン:**

```tsx
// 1. 依存配列なし：毎回のレンダリング後に実行（通常は避ける）
useEffect(() => {
  console.log('毎回実行される');
});

// 2. 空の依存配列：マウント時のみ実行
useEffect(() => {
  console.log('最初の1回だけ実行される');
}, []);

// 3. 値を指定：その値が変わった時のみ実行
useEffect(() => {
  console.log(`userIdが${userId}に変わりました`);
}, [userId]);
```

### クリーンアップ関数

**問題：コンポーネントがアンマウントされる時に後片付けをしたい**

```tsx
function Timer() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    // タイマーを設定
    const timer = setInterval(() => {
      setCount(c => c + 1);
    }, 1000);

    // クリーンアップ関数：コンポーネントがアンマウントされる時に実行
    return () => {
      clearInterval(timer); // タイマーを停止
    };
  }, []);

  return <div>{count}秒経過</div>;
}
```

### 実用例

**1. ドキュメントタイトルの変更**
```tsx
function MemoDetail({ memo }: { memo: Memo }) {
  useEffect(() => {
    document.title = memo.title;

    // クリーンアップ：元に戻す
    return () => {
      document.title = 'メモアプリ';
    };
  }, [memo.title]);

  return <div>{memo.content}</div>;
}
```

**2. ローカルストレージとの同期**
```tsx
function Settings() {
  const [theme, setTheme] = useState(() => {
    // 初期値をlocalStorageから取得
    return localStorage.getItem('theme') || 'light';
  });

  useEffect(() => {
    // themeが変わるたびにlocalStorageに保存
    localStorage.setItem('theme', theme);
  }, [theme]);

  return (
    <select value={theme} onChange={e => setTheme(e.target.value)}>
      <option value="light">ライト</option>
      <option value="dark">ダーク</option>
    </select>
  );
}
```

**3. イベントリスナーの登録**
```tsx
function WindowSize() {
  const [width, setWidth] = useState(window.innerWidth);

  useEffect(() => {
    const handleResize = () => {
      setWidth(window.innerWidth);
    };

    // イベントリスナーを登録
    window.addEventListener('resize', handleResize);

    // クリーンアップ：イベントリスナーを削除
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return <div>ウィンドウ幅: {width}px</div>;
}
```

### useEffectとdefaultValue

**問題：フォームの初期値が更新されない**

Reactのフォーム要素で`defaultValue`を使った場合、初回レンダリング時のみ値が設定されます。URLパラメータなどが変わっても、`defaultValue`は更新されません。

```tsx
function SearchForm({ initialQuery }: { initialQuery: string }) {
  // URLパラメータが変わっても、defaultValueは更新されない
  return <input name="q" defaultValue={initialQuery} />;
}
```

**解決：useEffectでDOM操作**

useEffectを使って、依存する値が変わった時にDOMを直接更新します。

```tsx
import { useEffect } from 'react';

function SearchForm({ query }: { query: string | null }) {
  useEffect(() => {
    const searchField = document.getElementById("q");
    if (searchField instanceof HTMLInputElement) {
      searchField.value = query || "";
    }
  }, [query]); // queryが変わるたびに実行

  return <input id="q" name="q" defaultValue={query || ""} />;
}
```

**使用例：検索フォームとブラウザバック**
- ユーザーが検索して別ページに移動
- ブラウザの「戻る」ボタンで検索ページに戻る
- URLパラメータは更新されるが、`defaultValue`だけでは入力欄が更新されない
- useEffectでDOM操作することで、入力欄にクエリを復元できる

### useEffectを使うべき場面・避けるべき場面

**使うべき場面:**
- **DOM操作**: ブラウザAPIを直接使う必要がある時
- **タイマー**: setInterval、setTimeoutの設定・クリーンアップ
- **グローバルイベント**: window.addEventListener など
- **外部ライブラリの初期化**: チャートライブラリ、地図ライブラリなど

**避けるべき場面:**
- **データ取得**: React RouterのloaderやReact Queryを使う（詳しくは後述）
- **状態管理**: useStateで十分
- **計算処理**: useMemoやuseMemo不要な単純な計算

**まとめ:**
- useEffectは「副作用」（レンダリング以外の処理）を明示的に分離するためのもの
- 依存配列でパフォーマンス最適化や無限ループ防止
- データ取得や状態管理はより適切な手段を優先

### 注意：React Router v7ではuseEffectでデータフェッチしない

**従来の方法（useEffectでAPI呼び出し）:**
```tsx
// ❌ React Router v7では非推奨
function MemoList() {
  const [memos, setMemos] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch('/api/memos')
      .then(res => res.json())
      .then(data => {
        setMemos(data);
        setLoading(false);
      });
  }, []);

  if (loading) return <div>Loading...</div>;
  return <div>{/* memos表示 */}</div>;
}
```

**React Router v7の方法（loader関数）:**
```tsx
// ✅ React Router v7ではloader関数を使う
export async function loader() {
  const memos = await fetch('/api/memos').then(res => res.json());
  return { memos };
}

export default function MemoList({ loaderData }: Route.ComponentProps) {
  const { memos } = loaderData; // データが必ず揃っている
  return <div>{/* memos表示 */}</div>;
}
```

詳しくは`learn-ReactRouter.md`を参照してください。

---

---

## プロジェクト構造とコンポーネントの配置

プロジェクトが大きくなると、すべてのコンポーネントを1つのファイルや同じディレクトリに置くのは管理が難しくなります。適切にファイルを整理することで、コードの可読性・保守性が向上します。

### なぜコンポーネントを分離するのか

**問題：すべてのコンポーネントを1つのファイルに書くと...**

```tsx
// app/root.tsx（肥大化した例）
export default function Root() {
  return <Outlet />;
}

export function ErrorBoundary() {
  // エラー表示コンポーネント
}

export function DeleteModal({ memo, onClose }: { memo: Memo; onClose: () => void }) {
  // モーダルコンポーネント
}

export function ConfirmDialog({ message, onConfirm }: ConfirmDialogProps) {
  // 確認ダイアログコンポーネント
}

// ...他にもたくさんのコンポーネント
```

**問題点:**
- ファイルが長すぎて読みにくい
- 特定のコンポーネントを探しにくい
- 複数人で同じファイルを編集すると競合しやすい
- 関連性の低いコードが混在する

**解決：コンポーネントを別ファイルに分離**

```tsx
// app/root.tsx（本来の役割だけに集中）
export default function Root() {
  return <Outlet />;
}

export function ErrorBoundary() {
  // エラー表示コンポーネント
}

// app/components/DeleteModal.tsx（分離）
export function DeleteModal({ memo, onClose }: { memo: Memo; onClose: () => void }) {
  // モーダルコンポーネント
}

// app/components/ConfirmDialog.tsx（分離）
export function ConfirmDialog({ message, onConfirm }: ConfirmDialogProps) {
  // 確認ダイアログコンポーネント
}
```

**メリット:**
- **単一責任の原則**: 1ファイル1コンポーネント（または関連する小さなコンポーネント群）
- **見通しが良い**: ファイル構造を見れば、どんなコンポーネントがあるか一目瞭然
- **再利用しやすい**: 必要なコンポーネントだけインポートできる
- **テストしやすい**: コンポーネント単位でテストを書ける

### 典型的なディレクトリ構造

Reactプロジェクトでは、以下のようなディレクトリ構造が一般的です。

```
app/
├── root.tsx              ... アプリ全体のルート（Outlet配置、エラーハンドリング）
├── routes.ts             ... ルーティング設定（React Router）
├── routes/               ... 各ページコンポーネント
│   ├── index.tsx         ... トップページ
│   ├── memo.tsx          ... メモ詳細ページ
│   └── edit_memo.tsx     ... メモ編集ページ
├── components/           ... 再利用可能なUIコンポーネント
│   ├── DeleteModal.tsx   ... 削除確認モーダル
│   ├── Button.tsx        ... ボタンコンポーネント
│   └── Card.tsx          ... カードコンポーネント
├── api/                  ... API通信関数
│   └── memos.ts          ... メモ関連のAPI
└── app.css               ... グローバルスタイル
```

**各ディレクトリの役割:**

| ディレクトリ | 役割 | 配置するもの |
|------------|------|------------|
| `routes/` | ページコンポーネント | URLに対応するページ全体（loader/action含む） |
| `components/` | 再利用可能なUI部品 | ボタン、モーダル、カードなど汎用コンポーネント |
| `api/` | バックエンド通信 | fetch関数、API呼び出しロジック |
| `layouts/` | レイアウトコンポーネント | サイドバー、ヘッダーなど共通レイアウト |
| `hooks/` | カスタムフック | 再利用可能なロジック（useState/useEffectを使う） |
| `utils/` | ユーティリティ関数 | 日付フォーマット、バリデーションなど |

### コンポーネントの分離：実例

**リファクタリング前:**

```tsx
// app/root.tsx
export function DeleteModal({ memo, onClose }: { memo: Memo; onClose: () => void }) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>メモの削除</h2>
        <p>本当にこのメモを削除しますか？</p>
        <Form method="post" action="/memos/delete" onSubmit={onClose}>
          <input type="hidden" name="id" value={memo.id} />
          <button type="button" onClick={onClose}>キャンセル</button>
          <button type="submit">削除</button>
        </Form>
      </div>
    </div>
  );
}

// app/routes/index.tsx
import { DeleteModal } from "~/root"; // root.tsxからインポート
```

**リファクタリング後:**

```tsx
// app/components/DeleteModal.tsx（新規作成）
import { Form } from "react-router";
import type { Memo } from "~/api/memos";

export function DeleteModal({ memo, onClose }: { memo: Memo; onClose: () => void }) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <h2>メモの削除</h2>
        <p>本当にこのメモを削除しますか？</p>
        <Form method="post" action="/memos/delete" onSubmit={onClose}>
          <input type="hidden" name="id" value={memo.id} />
          <button type="button" onClick={onClose}>キャンセル</button>
          <button type="submit">削除</button>
        </Form>
      </div>
    </div>
  );
}

// app/root.tsx（DeleteModalを削除）
export default function Root() {
  return <Outlet />;
}

// app/routes/index.tsx（インポート先を変更）
import { DeleteModal } from "~/components/DeleteModal";
```

**変更点:**
1. `components/`ディレクトリを作成
2. `DeleteModal`を`components/DeleteModal.tsx`に移動
3. 必要なインポート（Form, Memo型）を追加
4. 使う側のインポートパスを`~/root`から`~/components/DeleteModal`に変更

### インポートパスの書き方

Reactプロジェクトでは、`~`（チルダ）をプロジェクトルートのエイリアスとして使うことが多いです。

**相対パスと絶対パス:**

```tsx
// ❌ 相対パス：ファイル移動時に壊れやすい
import { DeleteModal } from "../../../components/DeleteModal";

// ✅ 絶対パス（~エイリアス）：ファイル移動に強い
import { DeleteModal } from "~/components/DeleteModal";
```

**~エイリアスの設定（React Router v7）:**

React Router v7では、デフォルトで`~`が`app/`ディレクトリを指すように設定されています。

```tsx
// app/routes/index.tsx から
import { DeleteModal } from "~/components/DeleteModal";
// → app/components/DeleteModal.tsx を参照

import { getMemos } from "~/api/memos";
// → app/api/memos.ts を参照
```

### コンポーネント分離の判断基準

**いつコンポーネントを分離すべきか:**

| 状況 | 分離すべきか | 理由 |
|------|------------|------|
| 複数の場所で使う | ✅ する | 再利用性が高まる |
| 50行以上の長いコンポーネント | ✅ する | 可読性が向上する |
| 独立したUI要素（モーダル、ボタンなど） | ✅ する | 単一責任の原則 |
| 1箇所でしか使わない小さなコンポーネント | ❌ しない | 過度な分離は逆に複雑化 |
| 親コンポーネントと強く結合している | ❌ しない | 分離しても再利用できない |

**例：モーダルコンポーネント**

```tsx
// ✅ 分離すべき：複数ページで使う可能性がある
// app/components/DeleteModal.tsx
export function DeleteModal({ item, onClose }: DeleteModalProps) {
  // ...
}

// ✅ 分離すべき：汎用的なボタンコンポーネント
// app/components/Button.tsx
export function Button({ children, onClick, variant }: ButtonProps) {
  // ...
}

// ❌ 分離不要：特定ページでしか使わない小さな部品
// app/routes/index.tsx内に定義
function MemoListHeader() {
  return <h1>メモ一覧</h1>;
}
```

### 段階的なリファクタリング

プロジェクトの成長に応じて、段階的にファイル構造を整理します。

**ステップ1: すべて1ファイル（学習初期）**
```
app/
└── routes/
    └── index.tsx  // すべてここに書く
```

**ステップ2: コンポーネントを分離**
```
app/
├── routes/
│   └── index.tsx
└── components/
    └── MemoCard.tsx  // 再利用可能なコンポーネントを分離
```

**ステップ3: API通信を分離**
```
app/
├── routes/
│   └── index.tsx
├── components/
│   └── MemoCard.tsx
└── api/
    └── memos.ts  // fetch関数をまとめる
```

**ステップ4: さらに細分化（大規模プロジェクト）**
```
app/
├── routes/
├── components/
│   ├── common/       // 汎用コンポーネント
│   │   ├── Button.tsx
│   │   └── Modal.tsx
│   └── memo/         // メモ関連コンポーネント
│       ├── MemoCard.tsx
│       └── MemoForm.tsx
├── api/
├── hooks/            // カスタムフック
└── utils/            // ユーティリティ関数
```

### まとめ

**コンポーネント配置のベストプラクティス:**

1. **root.tsx**: アプリ全体の設定（Outlet、エラーハンドリング）のみ
2. **routes/**: ページ全体の責務（loader/action + レイアウト）
3. **components/**: 再利用可能なUI部品
4. **api/**: バックエンド通信ロジック
5. **~エイリアス**: 絶対パスで明確にインポート
6. **段階的に整理**: 最初から完璧を目指さず、必要に応じてリファクタリング

適切なファイル構造は、チーム開発や将来の自分のためのドキュメントにもなります。

---

## 開発環境とツール

Reactアプリの開発には、いくつかのツールやファイルが関わります。

### package.json と package-lock.json

**package.json** - プロジェクトの設計図
```json
{
  "name": "my-app",
  "version": "1.0.0",
  "dependencies": {
    "react": "^19.0.0",
    "react-router": "^7.0.0"
  },
  "scripts": {
    "dev": "vite",
    "build": "vite build"
  }
}
```

- プロジェクトの基本情報（名前、バージョン）を記述
- 使いたいライブラリとそのバージョン範囲を指定
- よく使うコマンドをscriptsに定義

**package-lock.json** - 実際にインストールされたパッケージの記録

- 実際にインストールされたすべてのパッケージとその正確なバージョンを記録
- 依存関係の依存関係（間接的な依存）も含む
- チーム全員が同じバージョンを使うために重要

**例え:**
- package.json = 「レシピ」（材料と手順）
- package-lock.json = 「実際に使った材料の記録」（ブランド、ロット番号まで記録）

### Reactアプリに起動コマンドが多い理由

Reactプロジェクトには、用途別に複数のコマンドがあります。

**典型的なコマンド:**
```json
{
  "scripts": {
    "dev": "vite",              // 開発サーバー起動
    "build": "vite build",      // 本番用ビルド
    "start": "vite preview",    // ビルド結果をプレビュー
    "typecheck": "tsc",         // 型チェック
    "lint": "eslint ."          // コード品質チェック
  }
}
```

**なぜ複数あるのか:**
1. **開発用と本番用の違い**: 開発時は高速なホットリロード、本番は最適化されたファイル
2. **ビルドツール**: Vite、Webpackなどツール固有のコマンド
3. **React Router独自**: React Router v7は独自のビルドシステムを持つ
4. **品質管理**: 型チェック、リント、テストなど複数のツール

**それぞれの役割:**
- `npm run dev`: 開発中に使う（ホットリロード付き）
- `npm run build`: デプロイ前に本番用ファイルを生成
- `npm run start`: ビルド結果をローカルで確認

### Viteとは

**Vite**は、Reactなどのフロントエンド開発で使われる「開発サーバー」と「ビルドツール」です。

**Viteの役割:**
1. **開発サーバー**: コードを変更すると即座にブラウザに反映（ホットリロード）
2. **ビルドツール**: 本番用に最適化されたファイルを生成

**Viteの特徴:**
- **超高速な起動**: 従来のツール（Webpack）より圧倒的に速い
- **高速なホットリロード**: ファイル保存から反映まで瞬時
- **モダンな技術**: ES Modulesを活用

**Viteは必要か:**
- 小規模・学習用なら無しでもReactは動く（CDNでReactを読み込むなど）
- しかし、実際の開発では以下の理由でViteが推奨される：
  - TypeScriptのコンパイル
  - JSXの変換
  - 複数ファイルの結合
  - 本番用の最適化（圧縮、コード分割）

### SSR と CSR

Reactアプリには、レンダリング方法によって2つのアプローチがあります。

**CSR（Client-Side Rendering）**

- HTMLはほぼ空っぽの状態でブラウザに送られる
- JavaScriptがダウンロードされ、実行されてからUIが表示される

```html
<!-- サーバーから送られるHTML -->
<div id="root"></div>
<script src="/bundle.js"></script>

<!-- JavaScriptが実行されて初めて内容が表示される -->
```

**メリット:**
- インタラクティブ性が高い（SPA）
- サーバー負荷が低い
- PWA（オフライン対応）に向いている

**デメリット:**
- 初期表示が遅い（JavaScriptのダウンロード・実行が必要）
- SEOに弱い（検索エンジンがJavaScript実行前のHTMLを見る）

**SSR（Server-Side Rendering）**

- サーバー側で完全なHTMLを生成してブラウザに送る
- ブラウザは即座に内容を表示できる

```html
<!-- サーバーから送られるHTML（すでに内容がある） -->
<div id="root">
  <h1>ようこそ</h1>
  <p>これはSSRで生成されたページです</p>
</div>
<script src="/bundle.js"></script>
```

**メリット:**
- 初期表示が速い（HTMLがすでに完成している）
- SEOに強い（検索エンジンが内容を読める）
- アクセシビリティに有利

**デメリット:**
- サーバー負荷が高い（毎回HTMLを生成）
- 実装が複雑になる

**使い分けの基準:**

| 用途 | 推奨 | 理由 |
|------|------|------|
| 管理画面 | CSR | SEO不要、インタラクティブ性重視 |
| ブログ・コンテンツサイト | SSR | SEO重要、初期表示速度重視 |
| ECサイト | SSR or ハイブリッド | SEOと速度の両立 |
| リアルタイムアプリ | CSR | インタラクション重視 |

**ハイブリッドアプローチ:**
- **SSG（Static Site Generation）**: ビルド時にHTMLを生成（Next.jsなど）
- **ISR（Incremental Static Regeneration）**: 一定時間ごとに静的HTMLを再生成
- **Hydration**: SSRで生成したHTMLに、クライアント側でイベントハンドラを追加

---

## まとめ：開発の流れ

Reactアプリの開発は、通常以下の流れで進めます：

1. **静的なページを作る**
   - コンポーネントとJSXで見た目を定義

2. **データを表示する**
   - Propsで親から子へデータを渡す
   - map()で配列データをリスト表示
   - 条件付きレンダリングで状態に応じた表示

3. **ユーザー操作に反応する**
   - イベントハンドラでクリックや入力に反応

4. **動的な値を扱う**
   - useStateで変化する値を管理
   - UIの状態（モーダル開閉など）を管理

5. **外部とやり取りする**
   - useEffectで副作用を処理
   - ただしReact Router v7ではloader関数を使うことが多い

6. **プロジェクト構造を整理する**
   - コンポーネントを適切に配置する
   - ファイル構造で可読性・保守性を向上

この順序で理解を深めていけば、Reactの基礎は習得できます。
