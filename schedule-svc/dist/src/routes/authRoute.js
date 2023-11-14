"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const authController_1 = require("../controllers/authController");
const Router = express_1.default.Router();
Router.route("/")
    .get(authController_1.AuthController.getAllAuth)
    .post(authController_1.AuthController.createAuth);
Router.route("/lecturer").get(authController_1.AuthController.getAllLecturer); // Lấy những auth có role là "lecturer"
Router.route("/:id")
    .get(authController_1.AuthController.getAuth)
    .put(authController_1.AuthController.updateAuth)
    .delete(authController_1.AuthController.deleteAuth);
Router.route("/check-subscribe/:id").get(authController_1.AuthController.checkStatusSubscribe);
Router.route("/un-subscribe/:id").get(authController_1.AuthController.unsubscribeState);
exports.default = Router;
