import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/index.tsx"),
    route("memos/:id", "routes/memo.tsx"),
    route("memos/:id/edit", "routes/edit_memo.tsx"),
    route("memos/new", "routes/new_memo.tsx"),
] satisfies RouteConfig;
