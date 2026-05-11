import { Context } from "hono";
import { sendSuccessResponse } from "./response";

// webhook subcription handler
export const subscribe = (c: Context) => {
  return sendSuccessResponse(201, "subcribed", c);
};
//
// create webhook handler
export const createWebhook = (c: Context) => {
  return sendSuccessResponse(201, "webhook created", c);
};
//
// get last 10 processed webhook handlers
export const webhookList = (c: Context) => {
  return sendSuccessResponse(200, "get successful", c);
};
//
// get webhook by id handler
export const webhookById = (c: Context) => {
  return sendSuccessResponse(200, "get successful", c);
};

export const login = (c: Context) => {
  // check that username exists
  //
  // create a session
  //
  // return response
  return sendSuccessResponse(201, "login successful", c);
};

export const register = async (c: Context) => {
  // check that username doesn't already exists

  // create username and session attached

  // return response
  return sendSuccessResponse(201, "registration successful", c);
};
