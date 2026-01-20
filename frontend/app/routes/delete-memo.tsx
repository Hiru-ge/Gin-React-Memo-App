import { deleteMemo } from "~/api/memos";
import type { Route } from "./+types/delete-memo";
import { redirect } from "react-router";

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const id = formData.get("id");
  await deleteMemo(Number(id));
  return redirect("/");
}
