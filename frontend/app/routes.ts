import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/index.tsx"),
  route("memos/:id", "routes/memo.tsx"),
  route("memos/:id/edit", "routes/edit-memo.tsx"),
  route("memos/new", "routes/new-memo.tsx"),
  route("memos/delete", "routes/delete-memo.tsx"),
] satisfies RouteConfig;
