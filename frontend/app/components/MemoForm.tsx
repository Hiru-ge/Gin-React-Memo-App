import { Form, useNavigate } from "react-router";

type MemoFormProps = {
  pageTitle: string;
  method: "post" | "put";
  defaultTitle?: string;
  defaultContent?: string;
};

export default function MemoForm({
  pageTitle,
  method,
  defaultTitle = "",
  defaultContent = "",
}: MemoFormProps) {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">{pageTitle}</h1>
        <Form method={method} className="space-y-6">
          <div>
            <label
              htmlFor="title"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              タイトル
            </label>
            <input
              id="title"
              name="title"
              type="text"
              required
              placeholder="メモのタイトルを入力してください"
              defaultValue={defaultTitle}
              className="w-full px-3 py-2 border border-gray-300 rounded-md"
            />
          </div>
          <div>
            <label
              htmlFor="content"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              内容
            </label>
            <textarea
              id="content"
              name="content"
              required
              placeholder="メモの内容を入力してください"
              defaultValue={defaultContent}
              className="w-full px-3 py-2 border border-gray-300 rounded-md h-64"
            />
          </div>
          <div className="flex space-x-4">
            <button
              type="submit"
              className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
            >
              保存
            </button>
            <button
              type="button"
              className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600"
              onClick={() => navigate(-1)}
            >
              キャンセル
            </button>
          </div>
        </Form>
      </div>
    </div>
  );
}
