"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function CreateEventPage() {
  const router = useRouter();
  
  const [name, setName] = useState("");
  const [date, setDate] = useState(""); // datetime-localの入力値(文字列)
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      // 日付の変換: 入力された時間をISO形式(UTC)に変換してGoに送る
      // 例: "2025-06-01T19:00" -> "2025-06-01T10:00:00.000Z"
      const isoDate = new Date(date).toISOString();

      const res = await fetch("http://localhost:8080/events", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: name,
          date: isoDate,
        }),
      });

      if (!res.ok) {
        throw new Error("イベント作成に失敗しました");
      }

      alert("イベントを作成しました！");
      router.push("/"); // トップページに戻る
      
    } catch (error) {
      console.error(error);
      alert("エラーが発生しました");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 p-4">
      <div className="bg-white p-8 rounded shadow-md w-full max-w-md">
        <h1 className="text-2xl font-bold mb-6 text-center text-gray-800">新規イベント作成</h1>
        
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* イベント名 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">イベント名</label>
            <input
              type="text"
              required
              placeholder="例: 定期練習、ライブ"
              className="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-indigo-500 outline-none"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </div>
          
          {/* 日時 */}
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">開催日時</label>
            <input
              type="datetime-local"
              required
              className="w-full border border-gray-300 rounded p-2 focus:ring-2 focus:ring-indigo-500 outline-none"
              value={date}
              onChange={(e) => setDate(e.target.value)}
            />
          </div>

          {/* ボタンエリア */}
          <div className="flex gap-4">
            <button
              type="button"
              onClick={() => router.back()} // キャンセルボタン
              className="w-1/3 bg-gray-200 text-gray-700 py-2 rounded hover:bg-gray-300 transition"
            >
              戻る
            </button>
            <button
              type="submit"
              disabled={loading}
              className="w-2/3 bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700 transition disabled:opacity-50"
            >
              {loading ? "作成中..." : "作成する"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}