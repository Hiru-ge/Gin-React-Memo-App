import { getMemoById, updateMemo } from "~/api/memos";
import type { Route } from "./+types/memo.tsx";
import type { Memo } from "~/api/memos";
import { useNavigate, Form, redirect } from "react-router";

export async function loader({params}: Route.LoaderArgs) {
  const memo: Memo = await getMemoById(Number(params.id));
  return memo;
}

export async function action({params, request}: Route.ActionArgs) {
    const formData = await request.formData();
    const title = formData.get("title") as string;
    const content = formData.get("content") as string;
    await updateMemo(Number(params.id), title, content);
    return redirect(`/memos/${params.id}`);
}

export default function MemoEdit({loaderData}: Route.ComponentProps) {
  const navigate = useNavigate();
  const memo = loaderData;
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">
          メモ編集
        </h1>
        <Form method="put" className="space-y-6">
            <input name="title" defaultValue={memo.title} className="w-full px-3 py-2 border border-gray-300 rounded-md" />
            <textarea name="content" defaultValue={memo.content} className="w-full px-3 py-2 border border-gray-300 rounded-md h-64"></textarea>
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