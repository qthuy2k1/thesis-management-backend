"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const studentDefController_1 = require("../controllers/studentDefController");
const Router = express_1.default.Router();
Router.route("/")
    .get(studentDefController_1.StudentDefController.getAllStudentDef)
    .post(studentDefController_1.StudentDefController.createStudentDef)
    .delete(studentDefController_1.StudentDefController.deleteAllStudentDef);
Router.route("/:id")
    .get(studentDefController_1.StudentDefController.getStudentDef)
    .put(studentDefController_1.StudentDefController.updateStudentDef)
    .delete(studentDefController_1.StudentDefController.deleteStudentDef);
Router.route("/list-studef/:id").get(studentDefController_1.StudentDefController.getAllStudentDefPag);
exports.default = Router;
