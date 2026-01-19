import { createMemo } from "~/api/memos";
import type { Route } from "./+types/memo.tsx";
import { useNavigate, Form, redirect } from "react-router";

export async function action({request}: Route.ActionArgs) {
    const formData = await request.formData();
    const title = formData.get("title") as string;
    const content = formData.get("content") as string;
    const id: number = await createMemo(title, content);
    return redirect(`/memos/${id}`);
}

export default function MemoCreate() {
  const navigate = useNavigate();
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">
          メモ作成
        </h1>
        <Form method="post" className="space-y-6">
            <input name="title" className="w-full px-3 py-2 border border-gray-300 rounded-md" />
            <textarea name="content" className="w-full px-3 py-2 border border-gray-300 rounded-md h-64"></textarea>
            <div className="flex space-x-4">
                <button type="submit" className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600">
                    保存
                </button>
                <button type="button" className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600" onClick={() => navigate(-1)}>
                    キャンセル
                </button>
            </div>
        </Form>
      </div>
    </div>
  );
}