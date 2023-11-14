"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const pointDefController_1 = require("../controllers/pointDefController");
const Router = express_1.default.Router();
Router.route("/")
    .get(pointDefController_1.PointDefController.getAllPointDef)
    .post(pointDefController_1.PointDefController.createOrUpdatePointDef);
Router.route("/:id")
    .get(pointDefController_1.PointDefController.getPointDef)
    .put(pointDefController_1.PointDefController.updatePointDef)
    .delete(pointDefController_1.PointDefController.deletePointDef);
Router.route("/student-point/:studefId&:lecId").get(pointDefController_1.PointDefController.getPointDefForLecturer);
exports.default = Router;
