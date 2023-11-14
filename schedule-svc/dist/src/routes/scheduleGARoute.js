"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const GAController_1 = require("../controllers/GAController");
const Router = express_1.default.Router();
Router.route("/").post(GAController_1.GAScheduleController.scheduleGAThesisDefense);
Router.route("/").get(GAController_1.GAScheduleController.getSchedule);
Router.route("/student/:id").get(GAController_1.GAScheduleController.getCouncilInSchedule);
Router.route("/student-schedule/:id").get(GAController_1.GAScheduleController.getScheduleForStudent);
Router.route("/lecturer/:id").get(GAController_1.GAScheduleController.getCouncilInSchedule);
Router.route("/lecturer-schedule/:id").get(GAController_1.GAScheduleController.getScheduleForLecturer);
exports.default = Router;
