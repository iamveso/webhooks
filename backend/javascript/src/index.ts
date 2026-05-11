import { Hono } from "hono";
import { requestId } from "hono/request-id";
import { logger } from "hono/logger";
import { sendErrorResponse, sendSuccessResponse } from "./response";
import {
  createWebhook,
  login,
  register,
  subscribe,
  webhookById,
  webhookList,
} from "./handlers";

const app = new Hono();

app.onError((err, c) => {
  console.error(err);
  return sendErrorResponse(500, "something went wrong", c);
});

app.use("*", requestId());
app.use("*", logger());

app.get("/health", (c) => {
  return c.json({ message: "ok" });
});

app.post("/webhooks/subscribe", subscribe);

app.post("/webhooks", createWebhook);

app.get("/webhooks", webhookList);

app.get("webhooks/:id", webhookById);

app.post("/login", login);
app.post("/register", register);

export default app;
