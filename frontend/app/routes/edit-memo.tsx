import { getMemoById, updateMemo } from "~/api/memos";
import type { Route } from "./+types/edit-memo";
import type { Memo } from "~/api/memos";
import { redirect } from "react-router";
import MemoForm from "~/components/MemoForm";

export async function loader({ params }: Route.LoaderArgs) {
  const memo: Memo = await getMemoById(Number(params.id));
  return memo;
}

export async function action({ params, request }: Route.ActionArgs) {
  const formData = await request.formData();
  const title = formData.get("title") as string;
  const content = formData.get("content") as string;
  await updateMemo(Number(params.id), title, content);
  return redirect(`/memos/${params.id}`);
}

export default function MemoEdit({ loaderData }: Route.ComponentProps) {
  const memo = loaderData;
  return (
    <MemoForm
      pageTitle="メモ編集"
      method="put"
      defaultTitle={memo.title}
      defaultContent={memo.content}
    />
  );
}
