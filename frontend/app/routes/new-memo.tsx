import { createMemo } from "~/api/memos";
import type { Route } from "./+types/new-memo";
import { redirect } from "react-router";
import MemoForm from "~/components/MemoForm";

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;
  const id: number = await createMemo(title, content);
  return redirect(`/memos/${id}`);
}

export default function MemoCreate() {
  return <MemoForm pageTitle="メモ作成" method="post" />;
}
