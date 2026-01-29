// ユーザー
export type User = {
  ID: string;
  Name: string;
  Email: string;
};

// イベント
export type Event = {
  ID: number;
  Name: string;
  Date: string;
  Attendances?: EventAttendance[];
};

// 出欠
export type EventAttendance = {
  ID: number;
  EventID: number;
  UserID: string;
  User?: User;
  Status: number; // 1:参加, 2:不参加
  Comment: string;
};