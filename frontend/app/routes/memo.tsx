import { getMemoById } from "~/api/memos";
import type { Route } from "./+types/memo.tsx";
import type { Memo } from "~/api/memos";
import { useNavigate } from "react-router";
import { useState } from 'react';
import { DeleteModal } from "~/components/DeleteModal";

export async function loader({ params }: Route.LoaderArgs) {
  const memo: Memo = await getMemoById(Number(params.id));
  return memo;
}

export default function MemoDetail({ loaderData }: Route.ComponentProps) {
  const navigate = useNavigate();
  const [deletingMemo, setDeletingMemo] = useState<Memo | null>(null);
  const memo = loaderData;
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* ヘッダー */}
        <div className="mb-8 flex items-center">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            メモ詳細
          </h1>
          <button className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600 ml-6" onClick={() => { navigate('/'); }}>
            メモ一覧に戻る
          </button>
        </div>

        <div className="bg-white rounded-lg shadow p-6">
          <h2 className="text-xl font-semibold text-gray-800 mb-4">
            {memo.title}
          </h2>
          <p className="text-gray-700 whitespace-pre-wrap">
            {memo.content}
          </p>
          <p className="text-sm text-gray-500 mt-4">
            作成日: {new Date(memo.created_at).toLocaleDateString()}
          </p>
          <div className="mt-4 flex space-x-2">
            <button className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600" onClick={() => { navigate(`/memos/${memo.id}/edit`); }}>
              編集
            </button>
            <button className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600" onClick={() => setDeletingMemo(memo)}>
              削除
            </button>
          </div>
        </div>
      </div>
      {deletingMemo && (
        <DeleteModal memo={deletingMemo} onClose={() => setDeletingMemo(null)} />
      )}
    </div>
  );
}
