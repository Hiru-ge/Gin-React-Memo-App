import { Form } from "react-router";
import type { Memo } from "~/api/memos";

export function DeleteModal({ memo, onClose }: { memo: Memo; onClose: () => void }) {

  return (
    // 背景オーバーレイ
    <div
      className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
      onClick={onClose}
    >
      {/* モーダル本体 */}
      <div
        className="bg-white rounded-lg p-6 max-w-md w-full mx-4"
        onClick={(e) => e.stopPropagation()}
      >
        <h2 className="text-xl font-bold mb-4">メモの削除</h2>
        <p className="mb-6">本当にこのメモを削除しますか？</p>

        <Form method="post" action="/memos/delete" onSubmit={onClose}>
          <input type="hidden" name="id" value={memo.id} />
          <div className="flex space-x-2 justify-end">
            <button
              type="button"
              className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-400"
              onClick={onClose}
            >
              キャンセル
            </button>
            <button
              type="submit"
              className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
            >
              削除
            </button>
          </div>
        </Form>
      </div>
    </div>
  );
}
