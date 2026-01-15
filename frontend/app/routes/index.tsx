function Memos() {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* ヘッダー */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">
            メモ一覧
          </h1>
          <p className="text-base text-gray-600">
            現在のメモ数: 2件
          </p>
        </div>

        {/* カードグリッド */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {/* カード1 */}
          <div className="bg-white rounded-lg shadow-md hover:shadow-xl transition-all duration-300 hover:-translate-y-1 cursor-pointer border border-gray-200 overflow-hidden">
            <div className="p-6">
              <h2 className="text-xl font-semibold text-gray-900 mb-3 line-clamp-2">
                メモタイトル1
              </h2>
              <p className="text-gray-700 text-sm leading-relaxed line-clamp-3">
                メモの内容がここに表示されます。
              </p>
            </div>
            <div className="px-6 py-3 bg-gray-50 border-t border-gray-100">
              <span className="text-xs text-gray-500">
                2024年1月15日
              </span>
            </div>
          </div>

          {/* カード2 */}
          <div className="bg-white rounded-lg shadow-md hover:shadow-xl transition-all duration-300 hover:-translate-y-1 cursor-pointer border border-gray-200 overflow-hidden">
            <div className="p-6">
              <h2 className="text-xl font-semibold text-gray-900 mb-3 line-clamp-2">
                メモタイトル2
              </h2>
              <p className="text-gray-700 text-sm leading-relaxed line-clamp-3">
                メモの内容がここに表示されます。
              </p>
            </div>
            <div className="px-6 py-3 bg-gray-50 border-t border-gray-100">
              <span className="text-xs text-gray-500">
                2024年1月15日
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default function Home() {
  return <Memos />;
}
