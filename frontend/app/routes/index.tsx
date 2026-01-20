import { getMemos } from "~/api/memos";
import type { Route } from "./+types";
import type { Memo } from "~/api/memos";
import { useNavigate } from "react-router";
import { useState } from "react";
import { DeleteModal } from "~/components/DeleteModal";

export async function loader() {
  const memos: Memo[] = await getMemos();
  return { memos };
}

export default function Memos({ loaderData }: Route.ComponentProps) {
  const navigate = useNavigate();
  const [deletingMemo, setDeletingMemo] = useState<Memo | null>(null);
  const { memos } = loaderData;
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* ヘッダー */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">メモ一覧</h1>
          <button
            className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 mb-4"
            onClick={() => {
              navigate("/memos/new");
            }}
          >
            新規メモ作成
          </button>
          <p className="text-base text-gray-600">
            現在のメモ数: {memos.length}件
          </p>
        </div>

        {/* カードグリッド */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {memos.map((memo) => (
            <div
              key={memo.id}
              className="bg-white rounded-lg shadow p-6"
              onClick={() => navigate(`/memos/${memo.id}`)}
            >
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
                <button
                  className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                  onClick={(e) => {
                    e.stopPropagation();
                    navigate(`/memos/${memo.id}/edit`);
                  }}
                >
                  編集
                </button>
                <button
                  className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
                  onClick={(e) => {
                    e.stopPropagation();
                    setDeletingMemo(memo);
                  }}
                >
                  削除
                </button>
              </div>
            </div>
          ))}
        </div>

        {/* 削除モーダル */}
        {deletingMemo && (
          <DeleteModal
            memo={deletingMemo}
            onClose={() => setDeletingMemo(null)}
          />
        )}
      </div>
    </div>
  );
}
