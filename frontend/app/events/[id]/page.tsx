"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { Event, User } from "@/types";

export default function EventDetailPage() {
  const params = useParams(); // URLからIDを取得 (例: /events/1 → id: "1")
  const router = useRouter();
  const id = params.id;

  const [event, setEvent] = useState<Event | null>(null);
  const [users, setUsers] = useState<User[]>([]); // ユーザー選択用
  const [loading, setLoading] = useState(true);

  // フォーム入力用
  const [selectedUserId, setSelectedUserId] = useState("");
  const [status, setStatus] = useState("1"); // 1:参加, 2:不参加
  const [comment, setComment] = useState("");

  // 初期データ取得 (イベント詳細 & ユーザー一覧)
  useEffect(() => {
    const fetchData = async () => {
      try {
        // 1. イベント詳細を取得
        const eventRes = await fetch(`http://localhost:8080/events/${id}`);
        if (!eventRes.ok) throw new Error("イベント取得失敗");
        const eventData = await eventRes.json();
        setEvent(eventData);

        // 2. ユーザー一覧を取得（「私は誰？」を選択するため）
        const usersRes = await fetch("http://localhost:8080/users");
        if (usersRes.ok) {
          const usersData = await usersRes.json();
          setUsers(usersData || []);
        }
      } catch (error) {
        console.error(error);
        alert("データの読み込みに失敗しました");
      } finally {
        setLoading(false);
      }
    };

    if (id) fetchData();
  }, [id]);

  // 出欠送信処理
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedUserId) {
      alert("名前を選択してください");
      return;
    }

    try {
      const res = await fetch("http://localhost:8080/attendances", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          event_id: Number(id),
          user_id: selectedUserId,
          status: Number(status),
          comment: comment,
        }),
      });

      if (!res.ok) throw new Error("送信失敗");

      alert("回答しました！");
      window.location.reload(); // 画面を更新して最新の状態を表示
    } catch (error) {
      console.error(error);
      alert("エラーが発生しました");
    }
  };

  if (loading) return <div className="p-8">読み込み中...</div>;
  if (!event) return <div className="p-8">イベントが見つかりません</div>;

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-2xl mx-auto bg-white p-8 rounded shadow">
        
        {/* ヘッダー部分 */}
        <button onClick={() => router.push("/")} className="text-gray-500 mb-4 hover:underline">
          &larr; 一覧に戻る
        </button>
        <h1 className="text-3xl font-bold mb-2">{event.Name}</h1>
        <p className="text-gray-600 mb-8">
          📅 {new Date(event.Date).toLocaleString()}
        </p>

        {/* 出欠リスト表示 */}
        <div className="mb-10">
          <h2 className="text-xl font-semibold border-b pb-2 mb-4">みんなの回答</h2>
          {(!event.Attendances || event.Attendances.length === 0) ? (
            <p className="text-gray-500">まだ回答はありません。</p>
          ) : (
            <ul className="space-y-3">
              {event.Attendances.map((att) => (
                <li key={att.ID} className="flex justify-between items-center bg-gray-100 p-3 rounded">
                  <div>
                    <span className="font-bold mr-2">
                      {att.User ? att.User.Name : "不明なユーザー"}
                    </span>
                    <span className={`px-2 py-1 rounded text-sm ${att.Status === 1 ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800"}`}>
                      {att.Status === 1 ? "⭕️ 参加" : "❌ 不参加"}
                    </span>
                  </div>
                  <span className="text-gray-600 text-sm">{att.Comment}</span>
                </li>
              ))}
            </ul>
          )}
        </div>

        {/* 出欠入力フォーム */}
        <div className="bg-blue-50 p-6 rounded-lg">
          <h2 className="text-lg font-bold mb-4">出欠を回答する</h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            
            {/* 1. ユーザー選択（認証がないので仮実装） */}
            <div>
              <label className="block text-sm font-medium mb-1">あなたの名前</label>
              <select 
                className="w-full border rounded p-2"
                value={selectedUserId}
                onChange={(e) => setSelectedUserId(e.target.value)}
                required
              >
                <option value="">選択してください</option>
                {users.map((u) => (
                  <option key={u.ID} value={u.ID}>{u.Name}</option>
                ))}
              </select>
            </div>

            {/* 2. 参加/不参加 */}
            <div>
              <label className="block text-sm font-medium mb-1">回答</label>
              <div className="flex gap-4">
                <label className="flex items-center gap-2 cursor-pointer">
                  <input 
                    type="radio" 
                    name="status" 
                    value="1" 
                    checked={status === "1"}
                    onChange={(e) => setStatus(e.target.value)}
                  />
                  <span>⭕️ 参加</span>
                </label>
                <label className="flex items-center gap-2 cursor-pointer">
                  <input 
                    type="radio" 
                    name="status" 
                    value="2" 
                    checked={status === "2"}
                    onChange={(e) => setStatus(e.target.value)}
                  />
                  <span>❌ 不参加</span>
                </label>
              </div>
            </div>

            {/* 3. コメント */}
            <div>
              <label className="block text-sm font-medium mb-1">一言コメント</label>
              <input 
                type="text" 
                className="w-full border rounded p-2"
                placeholder="遅れます、など"
                value={comment}
                onChange={(e) => setComment(e.target.value)}
              />
            </div>

            <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">
              回答を送信
            </button>
          </form>
        </div>

      </div>
    </div>
  );
}