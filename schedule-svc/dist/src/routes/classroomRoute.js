"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const classroomController_1 = require("../controllers/classroomController");
const Router = express_1.default.Router();
Router.route("/")
    .get(classroomController_1.ClassroomController.getAllClassroom)
    .post(classroomController_1.ClassroomController.createClassroom);
Router.route("/:id")
    .get(classroomController_1.ClassroomController.getClassroom) // id (id của lecturer) get classroom trong đó có chứa lecturer
    .put(classroomController_1.ClassroomController.updateClassroom)
    .delete(classroomController_1.ClassroomController.deleteClassroom);
exports.default = Router;
