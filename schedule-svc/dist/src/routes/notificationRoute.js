"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const notificationController_1 = require("../controllers/notificationController");
const Router = express_1.default.Router();
Router.route("/").post(notificationController_1.notificationController.createNotification);
Router.route("/:id").get(notificationController_1.notificationController.getAllNotification);
exports.default = Router;
