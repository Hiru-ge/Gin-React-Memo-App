import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/index.tsx"),
    route("memos/:id", "routes/memo.tsx"),
] satisfies RouteConfig;
