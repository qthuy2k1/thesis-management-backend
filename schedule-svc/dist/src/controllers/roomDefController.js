"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RoomDefController = void 0;
const roomDefModel_1 = require("../models/roomDefModel");
exports.RoomDefController = {
    async createRoomDef(req, res, next) {
        const room = req.body;
        try {
            const createdRoomDef = await roomDefModel_1.RoomDefModel.createRoomDef(room);
            if (createdRoomDef) {
                res.status(200).json({
                    room: createdRoomDef,
                    message: "RoomDef has been created",
                });
            }
        }
        catch (err) {
            if (err) {
                console.error(err);
                res.status(500);
            }
        }
    },
    async getRoomDef(req, res) {
        const id = req.params.id;
        try {
            const room = await roomDefModel_1.RoomDefModel.getRoomDef(id);
            if (!room) {
                res.status(404);
                return;
            }
            res.status(200).json(room);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async getAllRoomDef(req, res) {
        try {
            const rooms = await roomDefModel_1.RoomDefModel.getAllRoomDef();
            if (!rooms) {
                res.status(404).json("RoomDef is empty");
                return;
            }
            res.status(200).json(rooms);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async updateRoomDef(req, res) {
        const room = req.body;
        const id = req.params.id;
        try {
            await roomDefModel_1.RoomDefModel.updateRoomDef({ id, ...room });
            res.status(200).json({ id, ...room });
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
    async deleteRoomDef(req, res) {
        const id = req.params.id;
        try {
            const rooms = await roomDefModel_1.RoomDefModel.deleteRoomDef(id);
            if (!rooms) {
                res.status(404);
                return false;
            }
            return res.status(200).json(rooms);
        }
        catch (err) {
            console.error(err);
            res.status(500);
        }
    },
};
