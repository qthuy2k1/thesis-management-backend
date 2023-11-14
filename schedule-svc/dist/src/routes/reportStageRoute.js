"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const reportStageController_1 = require("../controllers/reportStageController");
const Router = express_1.default.Router();
Router.route("/")
    .get(reportStageController_1.ReportStageController.getAllReportStage)
    .post(reportStageController_1.ReportStageController.createReportStage);
Router.route("/:id")
    .get(reportStageController_1.ReportStageController.getReportStage)
    .put(reportStageController_1.ReportStageController.updateReportStage)
    .delete(reportStageController_1.ReportStageController.deleteReportStage);
exports.default = Router;
