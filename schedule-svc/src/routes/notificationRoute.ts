import express from "express";
import { notificationController } from "../controllers/notificationController";
import checkAuth from "../middlewares/authorization";

const Router = express.Router();
Router.route("/").post(notificationController.createNotification);
Router.route("/:id").get(notificationController.getAllNotification);

export default Router;
