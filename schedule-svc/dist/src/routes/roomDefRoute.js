"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const roomDefController_1 = require("../controllers/roomDefController");
const Router = express_1.default.Router();
Router.route("/")
    .get(roomDefController_1.RoomDefController.getAllRoomDef)
    .post(roomDefController_1.RoomDefController.createRoomDef);
Router.route("/:id")
    .get(roomDefController_1.RoomDefController.getRoomDef)
    .put(roomDefController_1.RoomDefController.updateRoomDef)
    .delete(roomDefController_1.RoomDefController.deleteRoomDef);
exports.default = Router;
