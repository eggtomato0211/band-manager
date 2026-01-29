FROM node:22-alpine

WORKDIR /app

# 依存関係のインストール
COPY package*.json ./
RUN npm install

# ソースコードのコピー
COPY . .

# ポート公開
EXPOSE 3000

# 起動コマンド (開発用)
CMD ["npm", "run", "dev"]