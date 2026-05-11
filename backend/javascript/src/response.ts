import { Context } from "hono";
import { StatusCode } from "hono/utils/http-status";

type MetaResponse = {
  request_id: string;
};

type AppResponse = {
  status: string;
  message: string;
  data?: any;
  meta: MetaResponse;
};

export const sendErrorResponse = (
  code: StatusCode,
  message: string,
  c: Context,
) => {
  const response: AppResponse = {
    status: "error",
    message: message,
    meta: { request_id: c.get("requestId") },
  };
  c.status(code);
  return c.json(response);
};

export const sendSuccessResponse = (
  code: StatusCode,
  data: any,
  c: Context,
) => {
  const response: AppResponse = {
    status: "success",
    message: "success",
    data,
    meta: { request_id: c.get("requestId") },
  };
  c.status(code);
  return c.json(response);
};
