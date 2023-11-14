"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const uploadMiddleware_1 = __importDefault(require("../middlewares/uploadMiddleware"));
const uploadReportController_1 = require("../controllers/uploadReportController");
const Router = express_1.default.Router();
Router.route("/")
    .get(uploadReportController_1.UploadReportController.getAllUploadReport)
    .post(uploadMiddleware_1.default.array("attachment"), uploadReportController_1.UploadReportController.createUploadReport);
Router.route("/:id")
    .get(uploadReportController_1.UploadReportController.getUploadReport)
    .put(uploadReportController_1.UploadReportController.updateUploadReport)
    .delete(uploadReportController_1.UploadReportController.deleteUploadReport);
exports.default = Router;
