"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { Event } from "@/types"; // 作成した型定義をインポート

export default function Home() {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);

  // 画面が表示されたら一回だけ実行
  useEffect(() => {
    const fetchEvents = async () => {
      try {
        // GoのAPIからイベント一覧を取得
        const res = await fetch("http://localhost:8080/events");
        if (!res.ok) throw new Error("取得失敗");
        
        const data = await res.json();
        // データがnullの場合は空配列にする（Goがnullを返すことがあるため）
        setEvents(data || []);
      } catch (error) {
        console.error(error);
        alert("イベントの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchEvents();
  }, []);

  if (loading) return <div className="p-8 text-center">読み込み中...</div>;

  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <div className="max-w-4xl mx-auto">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800">バンド活動予定</h1>
          {/* イベント作成ページへのリンク（後で作ります） */}
          <Link 
            href="/events/new" 
            className="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700"
          >
            ＋ イベント作成
          </Link>
        </div>

        {events.length === 0 ? (
          <p className="text-gray-500 text-center">予定はまだありません。</p>
        ) : (
          <div className="grid gap-4">
            {events.map((event) => (
              <div key={event.ID} className="bg-white p-6 rounded-lg shadow hover:shadow-md transition">
                <div className="flex justify-between items-center">
                  <div>
                    <h2 className="text-xl font-semibold text-gray-900">{event.Name}</h2>
                    <p className="text-gray-600 mt-1">
                      {/* 日付のフォーマット整形 */}
                      📅 {new Date(event.Date).toLocaleDateString()}
                    </p>
                  </div>
                  <Link 
                    href={`/events/${event.ID}`}
                    className="text-indigo-600 hover:text-indigo-800 font-medium"
                  >
                    詳細・出欠へ &rarr;
                  </Link>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}