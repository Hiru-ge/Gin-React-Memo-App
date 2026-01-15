export type Memo = {
    id: number;
    title: string;
    content: string;
    created_at: Date;
}

export async function getMemos(): Promise<Memo[]> {
    const memos: Memo[] = await fetch('http://localhost:8080/memos').then(res => res.json());
    return memos;
}