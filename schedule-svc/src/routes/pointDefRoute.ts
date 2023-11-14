import express from "express";
import checkAuth from "../middlewares/authorization";
import { PointDefController } from "../controllers/pointDefController";

const Router = express.Router();

Router.route("/")
  .get(PointDefController.getAllPointDef)
  .post(PointDefController.createOrUpdatePointDef);
Router.route("/:id")
  .get(PointDefController.getPointDef)
  .put(PointDefController.updatePointDef)
  .delete(PointDefController.deletePointDef);
Router.route("/student-point/:studefId&:lecId").get(
  PointDefController.getPointDefForLecturer
);

export default Router;
