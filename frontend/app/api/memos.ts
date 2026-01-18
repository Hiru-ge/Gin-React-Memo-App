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

export async function getMemoById(id: number): Promise<Memo> {
    const memo: Memo = await fetch(`http://localhost:8080/memos/${id}`).then(res => res.json());
    return memo;
}

export async function editMemo(id: number, title: string, content: string): Promise<Memo> {
    const memo: Memo = await fetch(`http://localhost:8080/memos/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title, content }),
    }).then(res => res.json());
    console.log(memo);
    return memo;
}