import express from "express";
import { GAScheduleController } from "../controllers/GAController";

const Router = express.Router();

Router.route("/").post(GAScheduleController.scheduleGAThesisDefense);
Router.route("/").get(GAScheduleController.getSchedule);
Router.route("/student/:id").get(GAScheduleController.getCouncilInSchedule);
Router.route("/student-schedule/:id").get(
  GAScheduleController.getScheduleForStudent
);
Router.route("/lecturer/:id").get(GAScheduleController.getCouncilInSchedule);
Router.route("/lecturer-schedule/:id").get(
  GAScheduleController.getScheduleForLecturer
);

export default Router;
